package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// E is the variable containing the environment configuration,
// using the config types defined in the environment struct.
var E *e = new(e)

// Contains the environment variables
// that are required on the system
var shouldLoad = []string{
	"DATABASE_DSN",
	"SERVER_ADDR",
}

// Checks if the required environment variables are set.
// Otherwise, it panics.
func checkRequired() error {
	for _, key := range shouldLoad {
		if _, exists := os.LookupEnv(key); !exists {
			return fmt.Errorf("config load err: %s environment variable is required", key)
		}
	}
	return nil
}

// Function to load the OS environment configuration
// into the environment struct.
func loadConfig() {
	// database configuration
	E.Database.DSN = os.Getenv("DATABASE_DSN")

	// server configuration
	E.Server.Addr = os.Getenv("SERVER_ADDR")
}

// LoadEnv loads the environment configuration manually.
// It can receive a list of filenames to load the environment variables from.
// If no filenames are provided, it will assume the variables are already in the system.
// IMPORTANT: Should be called at the beginning of the main function,
// as other packages may depend on the environment configuration.
func LoadEnv(filenames ...string) error {
	if len(filenames) > 0 {
		err := godotenv.Overload(filenames...)
		if err != nil {
			return fmt.Errorf("dotenv files err: %s", err.Error())
		}
	}

	if err := checkRequired(); err != nil {
		return err
	}

	// Load the environment configuration
	loadConfig()

	return nil
}
