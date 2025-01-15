package config

type Config struct {
	Server Server `mapstructure:"server"`
	Infra  Infra  `mapstructure:"infra"`
}

type Server struct {
	Port    string `mapstructure:"port"`
	Profile string `mapstructure:"profile"`
}

type Infra struct {
	Db Db `mapstructure:"db"`
}

type Db struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	DBName   string `mapstructure:"dbName"`
	Port     string `mapstructure:"port"`
}
