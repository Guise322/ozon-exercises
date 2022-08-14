package sub_client

type SubClientConf struct {
	SubClient struct {
		Host string `yaml:"host"`
		Port int64  `yaml:"port"`
	} `yaml:"subClient"`
}
