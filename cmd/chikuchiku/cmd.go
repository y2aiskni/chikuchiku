package main

type Cmd struct {
	Post CmdPost `cmd:""`
}

type CmdPost struct {
	Discord CmdPostDiscord `cmd:"" help:"Post chikuchiku to Discord"`
}

type CmdPostDiscord struct {
}
