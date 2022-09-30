package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	CacheSize int64
	DbUrl     string
}

func NewConfig() (*Config, error) {
	cacheSize := os.Getenv("CACHE_SIZE")
	if cacheSize == "" {
		return nil, errors.New("empty CACHE_SIZE")
	}
	cs, err := strconv.Atoi(cacheSize)
	if err != nil {
		return nil, err
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, errors.New("empty DB_URL")
	}

	cfg := &Config{
		CacheSize: int64(cs),
		DbUrl:     dbUrl,
	}
	return cfg, nil
}
