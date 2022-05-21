package routers

import (
	"github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"

	"github.com/usernamedt/doctor-service/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	//pprof.Register(r)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(limits.RequestSizeLimiter(2000000))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/jobs/:id", v1.GetJob)
		apiv1.PUT("/jobs", v1.SubmitJob)
	}

	return r
}
