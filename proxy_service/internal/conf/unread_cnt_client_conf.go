package conf

type UnreadCntClientConf struct {
	UnreadCntClient struct {
		Host string `yaml:"host"`
		Port int64  `yaml:"port"`
	} `yaml:"unreadCntClient"`
}
