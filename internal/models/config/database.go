package modelconfig

type Postgresql struct {
	Address           string `mapstructure:"address"`
	Schema            string `mapstructure:"schema"`
	MaxOpenConnection int    `mapstructure:"max_open_connection"`
	MaxIdleConnection int    `mapstructure:"max_idle_connection"`
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Db       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
}
