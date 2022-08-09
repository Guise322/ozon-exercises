package server

type ServConf struct {
	Server struct {
		Host          string `yaml:"host"`
		Port          int64  `yaml:"port"`
		TimeoutInMils int64  `yaml:"timeoutInMils"`
	} `yaml:"server"`
}
