package main

type Cmd struct {
	Post CmdPost `cmd:""`
}

type CmdPost struct {
	Discord CmdPostDiscord `cmd:"" help:"Post chikuchiku to Discord"`

	Config string `short:"c" long:"config" description:"Path to config file" default:"./config.yaml"`
	Date   string `short:"d" long:"date" description:"Date to post (format: YYYY-MM-DD)"`
}

type CmdPostDiscord struct {
}
