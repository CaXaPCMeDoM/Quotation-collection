package config

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	Config struct {
		HTTP HTTP
	}

	HTTP struct {
		Port            string
		ReadTimeout     time.Duration
		WriteTimeout    time.Duration
		IdleTimeout     time.Duration
		MaxHeaderBytes  int
		ShutdownTimeout time.Duration
	}
)

func MustLoad() *Config {
	loadDotEnv()

	cfg := &Config{}

	// HTTP
	cfg.HTTP.Port = getEnv("HTTP_PORT", "")
	cfg.HTTP.ReadTimeout = parseDuration(getEnv("HTTP_READ_TIMEOUT", "30s"))
	cfg.HTTP.WriteTimeout = parseDuration(getEnv("HTTP_WRITE_TIMEOUT", "30s"))
	cfg.HTTP.IdleTimeout = parseDuration(getEnv("HTTP_IDLE_TIMEOUT", "60s"))
	cfg.HTTP.MaxHeaderBytes = parseInt(getEnv("HTTP_MAX_HEADER_BYTES", "1048576"))
	cfg.HTTP.ShutdownTimeout = parseDuration(getEnv("HTTP_SHUTDOWN_TIMEOUT", "5s"))

	if cfg.HTTP.ReadTimeout < 0 {
		log.Fatal("HTTP_READ_TIMEOUT cannot be negative")
	}
	if cfg.HTTP.WriteTimeout < 0 {
		log.Fatal("HTTP_WRITE_TIMEOUT cannot be negative")
	}

	return cfg
}

func loadDotEnv() {
	f, err := os.Open(".env")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatalf("error opening .env file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Fatalf("invalid .env line: %q", line)
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		val = strings.Trim(val, `"'`)
		os.Setenv(key, val)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading .env file: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue != "" {
			return defaultValue
		}
		log.Fatalf("Environment variable %s is required", key)
	}
	return value
}

func parseDuration(value string) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Fatalf("Invalid duration format for '%s': %v", value, err)
	}
	return duration
}

func parseInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Invalid integer format for '%s': %v", value, err)
	}
	return intValue
}
