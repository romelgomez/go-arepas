package router

import (
	"fmt"
	v1 "go-arepas/api/v1"

	post_controller "go-arepas/domain/post/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	postController *post_controller.PostController,
) *httprouter.Router {
	router := httprouter.New()

	// Home route
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	// Register v1 routes
	v1.NewPostRoutes(router, postController)

	// Register v2 routes (if needed)
	// v2.NewAccountRoutes(router, accountController)

	return router
}
