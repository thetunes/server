package config

var allowedOrigins = []string{
	"https://thetunes.github.io/",
	"https://tunes.herobuxx.me/",
}

// GetAllowedOrigins returns the list of allowed origins
func GetAllowedOrigins() []string {
	return allowedOrigins
}
