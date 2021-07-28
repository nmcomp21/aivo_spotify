package server

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	router := gin.New()

	router.RedirectFixedPath = true
	router.RedirectTrailingSlash = false

	mapping := newMapping()
	mapping.mapUrlsToControllers(router)

	return router
}
