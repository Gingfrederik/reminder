package config

type Config struct {
	Slack   SlackConfig
	Line    LineConfig
	Project []Project
	Notice  []Notice
}

type SlackConfig struct {
	Token   string
	Channel string
}

type LineConfig struct {
	Secret string
	Token  string
	Group  string
}

type Project struct {
	Name   string
	ID     int
	Branch []string
}

type Notice struct {
	Times   string
	User    string
	Message string
}
