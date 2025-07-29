package httpserve

import (
	"fmt"
	"net/http"
	"time"

	router "github.com/build-smile/CRUD_APIs_Arise/httpserve/router"
	"github.com/build-smile/CRUD_APIs_Arise/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	r := gin.Default()
	m := middleware.NewMiddleware(viper.GetString("jwt.secret-key"))
	r.Use(m.Recovery())
	r.Use(middleware.ErrorHandlerMiddleware)
	r.Use(m.ValidateOptionMethod())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allows all origins (You can specify the allowed origins)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache preflight request
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	//member
	router.BindGetMembers(r)
	router.BindGetMember(r)
	router.BindCreateMember(r)
	router.BindUpdateMember(r)
	router.BindDeleteMember(r)

	//Manually handle OPTIONS request for debugging
	port := viper.GetString("app.port")
	if port == "" {
		port = "3000"
	}
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
