package main

type config struct {
	Discord configDiscord `yaml:"discord"`
	Mysql   configMysql   `yaml:"mysql"`
}

type configDiscord struct {
	PostURL string `yaml:"post_url"`
}

type configMysql struct {
	Address  string `yaml:"address"`
	Port     uint   `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
