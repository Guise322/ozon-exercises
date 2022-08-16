package conf

type GRPCConf struct {
	Server struct {
		Host    string `yaml:"host"`
		Port    int64  `yaml:"port"`
		Timeout int64  `yaml:"timeoutInMilSec"`
	} `yaml:"grpcServer"`
}
