package post

import (
	"go-arepas/common"
	dto "go-arepas/domain/post/dto"
	service "go-arepas/domain/post/service"
	"go-arepas/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{PostService: postService}
}

func (controller *PostController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postCreateRequest := dto.PostCreate{}
	helper.ReadRequestBody(requests, &postCreateRequest)

	if postCreateRequest.Title == "" {
		helper.WriteErrorResponse(writer, "Title is required", http.StatusBadRequest)
		return
	}

	controller.PostService.Create(requests.Context(), postCreateRequest)
	webResponse := common.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PostController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postUpdateRequest := dto.PostUpdate{}
	helper.ReadRequestBody(requests, &postUpdateRequest)

	postId := params.ByName("postId")
	postUpdateRequest.Id = postId

	controller.PostService.Update(requests.Context(), postUpdateRequest)
	webResponse := common.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PostController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.PostService.FindAll(requests.Context())
	webResponse := common.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PostController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	result := controller.PostService.FindById(requests.Context(), postId)
	webResponse := common.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PostController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	controller.PostService.Delete(requests.Context(), postId)
	webResponse := common.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}
