// routes/routes.go

package routes

import (
	"github.com/asterfy/tis-clinic/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	especialidades := router.Group("/especialidades")
	{
		especialidades.GET("/", controllers.GetAllEspecialidades)
	}
	asegurados := router.Group("/asegurados")
	{
		asegurados.GET("/", controllers.GetAllAsegurados)
	}
	citas := router.Group("/citas")
	{
		citas.GET("/", controllers.GetAllCitas)
	}
	seguros := router.Group("/seguros")
	{
		seguros.GET("/", controllers.GetAllSeguros)
	}
	servicios := router.Group("/servicios")
	{
		servicios.POST("/reservarCita/", controllers.ReservarCitaService)
		servicios.POST("/asegurarPersonal/", controllers.AsegurarPersonalService)
	}

}
