package infra

type SubClientConf struct {
	SubClient struct {
		Host string `yaml:"host"`
		Port int64  `yaml:"port"`
	} `yaml:"subClient"`
}
