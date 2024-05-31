package config

type System struct {
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	IpLimitCount int    `mapstructure:"ip-limit-count" json:"ip-limit-count" yaml:"ip-limit-count"`
	IpLimitTime  int    `mapstructure:"ip-limit-time" json:"ip-limit-time" yaml:"ip-limit-time"`
	DB           string `mapstructure:"db" json:"db" yaml:"db"`
	Redis        bool   `mapstructure:"redis" json:"redis" yaml:"redis"`
}
