package server

type Config struct {
	Address     string
	LogLevel    string
	DatabaseUrl string
	Key         []byte
	LifeSpan    int
}
