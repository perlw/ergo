package main

import (
	"bufio"
	"bytes"
	"errors"
	"html/template"
	"log"
	"net"
	"net/http"
	"strings"
)

var (
	// ErrUnrecognizedCommand is returned when an unrecognized command is input to the interpreter
	ErrUnrecognizedCommand = errors.New("unrecognized command")
)

var gamePrompt = []byte("\r\n> ")

type GameRoom struct {
	Brief       string
	Description string
	Exits       map[string]*GameRoom
}

type GameCommand struct {
	Cmd     string
	Aliases []string
	Action  func(action string, arguments []string) string
}

type GameState struct {
	Commands []GameCommand
	Rooms    []GameRoom

	CurrentRoom *GameRoom
}

func NewGameState() *GameState {
	state := &GameState{}

	state.Commands = []GameCommand{
		{
			Cmd:     "help",
			Aliases: []string{"h", "?"},
			Action: func(action string, arguments []string) string {
				return "TBD"
			},
		},
		{
			Cmd:     "look",
			Aliases: []string{},
			Action: func(action string, arguments []string) string {
				return state.CurrentRoom.Description
			},
		},
	}

	state.Rooms = []GameRoom{
		{
			Brief:       "You're in a small jail cell.",
			Description: "long description",
			// Items, doors, monsters?
			Exits: map[string]*GameRoom{},
		},
	}

	state.CurrentRoom = &state.Rooms[0]

	return state
}

func (s *GameState) InterpretCommand(action string, arguments []string) (string, error) {
	for _, cmd := range s.Commands {
		matched := false
		if cmd.Cmd == action {
			matched = true
		} else {
			for _, alias := range cmd.Aliases {
				if alias == action {
					matched = true
					break
				}
			}
		}

		if matched {
			return cmd.Action(action, arguments), nil
		}
	}

	return "", ErrUnrecognizedCommand
}

func (s *GameState) WriteCurrentRoomBrief(w *bufio.Writer) {
	w.Write([]byte(s.CurrentRoom.Brief))
	// Write exits
	w.Flush()
}

func spawnGame(conn net.Conn) {
	defer conn.Close()

	state := NewGameState()

	w := bufio.NewWriter(conn)
	r := bufio.NewReader(conn)

	w.Write([]byte("The Dungeon\r\n===========\r\n\r\n"))
	state.WriteCurrentRoomBrief(w)
	w.Write(gamePrompt)
	w.Flush()

	lineEnd := false
	buffer := make([]byte, 1)
	line := bytes.Buffer{}
	for {
		n, err := r.Read(buffer)
		if err != nil {
			log.Println("could not read from client,", err)
			return
		} else if n <= 0 {
			continue
		}

		// Telnet commands
		if buffer[0] == 0xff {
			cmd := make([]byte, 2)
			n, err := r.Read(cmd)
			if n <= 0 || err != nil {
				log.Println("could not read cmd from client,", err)
				return
			}
			log.Println("cmd: ", cmd)
		} else if buffer[0] == '\n' || buffer[0] == '\r' {
			if line.Len() == 0 {
				if !lineEnd {
					lineEnd = true
					continue
				}
			}
			lineEnd = false

			msg := strings.TrimRight(line.String(), "\r\n")
			// Interpret the command
			if msg == "quit" {
				break
			} else {
				words := strings.Split(msg, " ")
				action := strings.ToLower(words[0])
				arguments := words[1:]

				output, err := state.InterpretCommand(action, arguments)
				if err != nil {
					log.Printf("could not interpret \"%s\", %s", action, err)

					if err == ErrUnrecognizedCommand {
						w.Write([]byte("Sorry, I don't know how to '" + action + "'..."))
					}
				} else {
					w.Write([]byte(output))
				}
				w.Write([]byte("\r\n"))
			}

			line.Reset()
			w.Write(gamePrompt)
			w.Flush()
		} else {
			line.WriteByte(buffer[0])
		}
	}

	w.Write([]byte("BYE"))
	w.Flush()
}

func serveGame() {
	ln, err := net.Listen("tcp", ":23")
	if err != nil {
		panic(err)
	}

	for {
		conn, _ := ln.Accept()
		go spawnGame(conn)
	}
}

func serveWeb() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("could not start server,", err)
	}
}

func main() {
	log.Println("Launching... well.. me")

	go serveGame()
	go serveWeb()

	var forever chan int = nil
	<-forever
}
