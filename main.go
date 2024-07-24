package main

import (
	"os"

	"github.com/asterfy/tis-clinic/initializers"
	"github.com/asterfy/tis-clinic/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	// Inicializar la conexión a la base de datos
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Crear un enrutador Gin
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true // If your frontend sends credentials
	router.Use(cors.New(config))

	// Configurar las rutas
	routes.SetupRoutes(router)

	// Iniciar el servidor Gin
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}

// REQUISITOS FUNCIONALES
// 	Mantenimiento (médico, paciente, cita, asegurado, receta)
// 	Proceso (Reserva de cita, Historial clínico)
