package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// if os.Getenv("ENVIRONMENT") == "development" {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar archivo .env")
	}
	// }
}
