package conf

type ServConf struct {
	Server struct {
		Host            string `yaml:"host"`
		Port            int64  `yaml:"port"`
		TimeoutInMilSec int64  `yaml:"timeoutInMilSec"`
	} `yaml:"server"`
}
