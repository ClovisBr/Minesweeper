package config

import "time"

type Config struct {
	Rows  int
	Cols  int
	Mines int

	Seed int64
}

func Default() Config {
	return Config{
		Rows:  9,
		Cols:  9,
		Mines: 10,
		Seed:  time.Now().UnixNano(),
	}
}
