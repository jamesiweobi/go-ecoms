package config

import (
	"sync"
	"time"
)

type Config interface {
	build()
	isSet(key string) bool
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetIntSlice(key string) []int
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
}

var once sync.Once

var instance (Config)

func Default() Config {
	once.Do(func() {
		instance = &ViperConfig{}
		instance.build()
	})
	return instance
}
