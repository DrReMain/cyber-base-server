package config

type Redis struct {
	Address        string   `mapstructure:"address" json:"address" yaml:"address"`
	Password       string   `mapstructure:"password" json:"password" yaml:"password"`
	DB             int      `mapstructure:"db" json:"db" yaml:"db"`
	Cluster        bool     `mapstructure:"cluster" json:"cluster" yaml:"cluster"`
	ClusterAddress []string `mapstructure:"cluster-address" json:"cluster-address" yaml:"cluster-address"`
}
