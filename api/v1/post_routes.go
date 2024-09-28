package v1

import (
	controller "go-arepas/domain/post/controller"

	"github.com/julienschmidt/httprouter"
)

func NewPostRoutes(router *httprouter.Router, postController *controller.PostController) {
	router.GET("/api/v1/post", postController.FindAll)
	router.GET("/api/v1/post/:postId", postController.FindById)
	router.POST("/api/v1/post", postController.Create)
	router.PATCH("/api/v1/post/:postId", postController.Update)
	router.DELETE("/api/v1/post/:postId", postController.Delete)
}
