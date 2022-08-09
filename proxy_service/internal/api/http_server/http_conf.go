package http_server

type httpConf struct {
	Server struct {
		Host string `yaml:"host"`
		Port int64  `yaml:"port"`
	} `yaml:"httpServer"`
}
