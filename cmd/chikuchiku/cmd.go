package main

type Cmd struct {
	Post CmdPost `cmd:""`
}

type CmdPost struct {
	Discord CmdPostDiscord `cmd:"" help:"Post chikuchiku to Discord"`

	Config string `short:"c" long:"config" description:"Path to config file" default:"./config.yaml"`
}

type CmdPostDiscord struct {
}
