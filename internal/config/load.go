package config

import (
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
}

// Checks if the required environment variables are set.
// Otherwise, it panics.
func checkRequired() {
	for _, key := range shouldLoad {
		if _, exists := os.LookupEnv(key); !exists {
			panic(key + " environment variable is required")
		}
	}
}

// LoadEnv loads the environment configuration manually.
// It can receive a list of filenames to load the environment variables from.
// If no filenames are provided, it will assume the variables are already in the system.
// IMPORTANT: Should be called at the beginning of the main function,
// as other packages may depend on the environment configuration.
func LoadEnv(filenames ...string) {
	if len(filenames) > 0 {
		err := godotenv.Overload(filenames...)
		if err != nil {
			panic(err)
		}
	}

	checkRequired()

	// database configuration
	E.Database.DSN = os.Getenv("DATABASE_DSN")
}
