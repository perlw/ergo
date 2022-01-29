So I'm leaving vim after 11 or so years.
It's got nothing to do with how vim works, or what it can't do that something else does better. This is all about choices I've made due to the ever changing circumstances of life.
So, there, for most of you that's probably all you needed to hear. Thanks for reading.

---

For the few that are actually interested in the why, thanks for staying.

I started using vim, or rather, a vim keybind plugin for Sublime Text, back around 2010. I'd read that using the **Awsome power of Vim** would make you a better developer. You'd be able to put your thoughts into code that much faster, navigating the code with ease, zipping around your code at the speed of light, impress your coworkers with your hacker eliteness.
I jest, but I honestly believed that using vim would make me a better developer.

So I got started, painfully getting used to the concept of modal editing, slowly but steadily building competence with the obscure language of commands and methods to edit your code, building up my config, my plugin-usage, etc.
And it was good. It was awesome even. I could *feel* my performance grow steadily, day by day. I really was editing code so much faster, I really did feel like I code zip around the codebase with ease.
For every passing day I learnt something new, I mastered another feature. I got used to the *hjkl* navigation to the degree that arrow keys started feeling unatural. The motion system of specifying for example *10j* to go down 10 lines or deleting lines with *dd* or rest of lines with *D* became second nature. I learnt how to use macros effectively and felt like I leveled up. I remapped Caps Lock to Escape for ease of use, and got so used to it that I constantly activate Casp Lock on accident on unfamiliar systems, even when not using vim.
Over time I simply stopped using graphical editors and IDEs and went hardcore, using terminal emulators with vim (and later neovim), to get the full, pure, experience.
I turned down using several different editors simply for not having vim-plugins, even when a dedicated IDE would really help.

It got to the point that I started to feel slow, hampered, when **not** having the vim command mode available in other applications. 
Browsers? Of course I used plugins when available, vimvixen for Firefox for example.
IDEs (when I was forced)? Any and all kinds of plugins to replicate the feeling of vim, which most of the time left me wanting for more, for the full vim experience.
CLI? Always vi-mode active on nix-like systems, even managed to get vi-mode to work on Powershell with a downloadable module.

And it was good. I felt like an elite developer. Just look how fast I can edit text! Living and breathing the command line, huffing it like glue.

This went on for a time, spewing the benefits of using vim to everyone who would listen, scoffing at using modern IDEs (my colleagues are all awesome people and simply took it as a minor quirk with me, bless them). I mean, you need to use your mouse in those! A mouse! No thank you, I'm comfortable on my home row.
What do you mean you use a specific note-taking application, I'm fine with my markdown files and git!
And so on.

After doing this for several years, I started to get annoyed that I *couldn't* use modern IDEs anymore. I *couldn't* try out new editors since there were most often not a vim plugin for it. I would detest using even plain text fields since I *couldn't* use my comfortable vim commands to move around.
Writing a lot of C++ at the time, I got interested in a new editor called 4coder. It looked brilliant for my use-case, and I was having some problems (quite few, but they annoyed me) with my C/C++ setup in vim, and 4coder seemed to solve it all.
So I bought my way into the Patreon (at the time) and gave it a shot.
Immediate frustration set in, it did not come with vim keybinds. It would be easy enough to add it in, 4coder has a comprehensive layer system so you can modify and extend the editor with ease, but I felt that I would probably need to dedicate quite some time to make it work just like I wanted it. And I simply could not find the time or motivation when my current setup worked to 99% just like I wanted it.
So I left 4coder behind, quite sadly I might add.

So, what is all this leading up to, you might ask. Well, besides me deciding to stop using vim, unlearning habits I've built up over years and years, and having to relearn even basic motions that are pretty much standardised over the standard desktop environments of today.
Well, quite frankly, I've gotten sick of it. Sick of it all. I've spent so much time dedicating myself to a way of work, that while I genuinly believe it is good, has made me so dependant on it that everything else feels bad.
I'm tired of passing up on the innovation that's constantly happening in the dev space, simply because I don't have my precious vim keybinds. I'm tired of having to find workarounds for things that only affects me at my workplace (liveshare in vscode as an example).

Since April I joined a new company, new team, new codebase, and for the first time in ages I've actually felt somewhat hindered by insisting on vim. The codebase is big, convoluted, and generally hard to navigate. I've had vscode with vim plugin on the side to help when navigating the codebase at least, and also used vscode during our PR reviews to ease walking through the code for the sake of others, seeing how much of a "magic mess" vim can become when you're not used to looking at it. Like it or not, it is *not* easy to understand what you're looking at when a vimmer is hopping around several files, splits, or tabs.
And it's been bugging me. That I've found it annoying to not being able to always use vim, to having to make concessions on account of "others". To actually feeling contempt over the situation.
It got me thinking. Real hard.
So, I left vim. Just up and quit, just days ago before writing this.

Vim has been so ingrained in me that using (for example on Mac) `⌥+←/→` to jump between words was something I only just learnt the day before writing this post.
Even now writing this post I'm struggling to not instantly reach for Escape after each edit or line written.

Now don't get me wrong. I still love vim. I still use it for working on remote systems via SSH or when editing local configuration files. It just won't be my daily driver any more.
I've removed the vim plugin from vscode. I've turned off vim mode in my terminals. 
I'm feeling slow, having to yet again methodically move around in the code using arrow keys and mouse. I'm missing a lot of my old workflow. But, I'm getting there. It's not at all been as slow of a process that I've thought it would be, reading, learning, referencing vscode cheatsheets and documentation looking for similar features to what I've been used to. Jumping between words and lines, inserting new lines above and below, navigating the code, both on a global stage and in the local text, I've managed to find replacements for most of the general codenavigation features I've relied on in vim.
Of course this has led me into using strange new incantations on the keyboard, but when thinking about it it's turned out to be no different than using the magic that is vim's modal editing. It's simply different.

Not feeling "constrained" to using vim for *everything* I want to write anymore has lead me into exploring more specialized alternatives, such as Obsidian for keeping my thoughts, which in turn prompted me to write this post.

I've been living in a bit of a crisis when it comes to development. I've felt stagnant. Friends and colleagues seemingly blasting past me in knowledge, energy, and motivation. All the while I don't feel in me what I see in them.
Gradually it's dawned on me that it's not the choice of editor or workflow that makes a great developer. It's obvious really, but being able to edit my code at the speed of sound is not what makes me fast. There's so much more to developing than simply having quick fingers. Often it's not unusual to spend more time simply thinking than you'll ever do actually writing code.

So, things have to change. I'm just about to turn the big 4-0 in a year and I feel left behind.
To be clear, I'm not blaming anyone. Or anything. This is all a me-problem. But things have to change.
So, as part of me reflecting back on myself, my goals, my choices, my hurts and my annoyances, I'm slowly changing things. Changing how I approach tasks and problems.

Changing me.

Who knows, maybe I'll end up back in vim again. Or maybe I won't.
But I hope that, whatever the future brings, I'll at least have a new perspective on things.

FIN
