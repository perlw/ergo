#!/bin/sh

cwd=$(pwd)

declare -A known_files

inotifywait -mr --timefmt '%s' --format '%T %w %f' -e close_write ./cmd |
  while read -r time dir file; do
    filepath=${dir}${file}

    if [[ $filepath =~ "~" ]]; then
      continue
    fi

    if [ ! -z ${known_files[$filepath]} ]; then
      if [ ! $time -gt ${known_files[$filepath]} ]; then
        continue
      fi
    fi
    known_files[$filepath]=$time

    echo "Updated: $filepath"
    pkill ergo
    echo "Ergo stopped!"
    if [[ $filepath =~ "go" ]]; then
      echo "Rebuilding Ergo..."
      go build -o ./bin/ergo ./cmd/ergo
    fi
    echo "Starting Ergo..."
    ./bin/ergo -port 1337 -web-base-dir ./web &
  done
