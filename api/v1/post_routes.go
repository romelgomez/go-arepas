package v1

import (
	controller "go-arepas/domain/post/controller"
	"go-arepas/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewPostRoutes(router *httprouter.Router, postController *controller.PostController, jwtMiddleware *middleware.JWTMiddleware) {
	// Public routes
	router.GET("/api/v1/post", postController.FindAll)
	router.GET("/api/v1/post/:postId", postController.FindById)

	// Privated routes
	router.POST("/api/v1/post", jwtMiddleware.ValidateJWT(postController.Create))
	router.PATCH("/api/v1/post/:postId", jwtMiddleware.ValidateJWT(postController.Update))
	router.DELETE("/api/v1/post/:postId", jwtMiddleware.ValidateJWT(postController.Delete))
}
