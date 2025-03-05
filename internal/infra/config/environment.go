package config

// e is a struct that contains the environment configuration
// and should be used by all packages inside infrastructure layer.
type e struct {
	Database database
	Server   server
}

// Database config type containing most common database configuration
type database struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	Path         string
	URL          string
	DSN          string
}

type server struct {
	Addr   string
	Secure bool
	Host   string
	Port   int
}
