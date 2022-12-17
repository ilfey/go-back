package config

type Config struct {
	Address      string
	LogLevel     string
	DatabaseUrl  string
	DatabaseFile string
	Key          []byte
	LifeSpan     int
}
