package infra

type NotifClientConf struct {
	NotifClient struct {
		Host string `yaml:"host"`
		Port int64  `yaml:"port"`
	} `yaml:"notifClient"`
}
