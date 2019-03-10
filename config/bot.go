package config

type Config struct {
	Slack   slackConfig
	Line    lineConfig
	Project []project
	Notice  []notice
}

type slackConfig struct {
	Token   string
	Channel string
}

type lineConfig struct {
	Secret string
	Token  string
	Group  string
}

type project struct {
	Name   string
	ID     int
	Branch []string
}

type notice struct {
	Times   string
	Message string
}
