package http_server

type HTTPConf struct {
	Server struct {
		Host string `yaml:"host"`
		Port int64  `yaml:"port"`
	} `yaml:"httpServer"`
}
