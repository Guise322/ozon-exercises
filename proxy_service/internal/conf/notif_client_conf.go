package conf

type NotifClientConf struct {
	NotifClient struct {
		URL string `yaml:"url"`
	} `yaml:"notifClient"`
}
