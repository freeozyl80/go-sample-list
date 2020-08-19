package config

import (
	"pattern/abstract/plugin"
)

type Config struct {
	Input plugin.Config
	Filter plugin.Config
	Output plugin.Config
}