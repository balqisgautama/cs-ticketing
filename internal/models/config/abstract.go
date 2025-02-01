package modelconfig

type Config struct {
	Server     Server     `mapstructure:"server"`
	Postgresql Postgresql `mapstructure:"postgresql"`
}
