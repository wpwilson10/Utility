package setup

import (
	"path/filepath"

	"github.com/joho/godotenv"
)

// EnvironmentConfig loads the .env file for the whole program.
// Use os.Getenv("LABEL_NAME") to access.
func EnvironmentConfig() {
	// get .env filepath
	absPath, err := filepath.Abs("./configs/.env")
	if err != nil {
		LogCommon(err).Fatal("Config filepath")
	}

	// get .env variables
	err = godotenv.Load(absPath)
	if err != nil {
		LogCommon(err).Fatal("Loading .env file")
	}
}
