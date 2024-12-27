package config

import "os"

type Config struct {
	SupabaseURL string
	SupabaseKey string
	SupabaseJWT string
}

func LoadConfig() *Config {
	return &Config{
		SupabaseURL: os.Getenv("SUPABASE_URL"),
		SupabaseKey: os.Getenv("SUPABASE_KEY"),
		SupabaseJWT: os.Getenv("SUPABASE_JWT"),
	}
}
