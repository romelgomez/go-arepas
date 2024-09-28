package post

import (
	"context"
	dto "go-arepas/domain/post/dto"
	model "go-arepas/domain/post/model"
	repository "go-arepas/domain/post/repository"
	"go-arepas/helper"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

func NewPostServiceImpl(postRepository repository.PostRepository) PostService {
	return &PostServiceImpl{PostRepository: postRepository}
}

func (p *PostServiceImpl) Create(ctx context.Context, request dto.PostCreate) {

	var subTitle *string
	if request.SubTitle != nil {
		subTitle = request.SubTitle
	} else {
		subTitle = nil
	}

	postData := model.Post{
		Title:       request.Title,
		SubTitle:    subTitle,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Save(ctx, postData)
}

func (p *PostServiceImpl) Delete(ctx context.Context, postId string) {
	post, err := p.PostRepository.FindById(ctx, postId)
	helper.ErrorPanic(err)
	p.PostRepository.Delete(ctx, post.Id)
}

func (p *PostServiceImpl) FindAll(ctx context.Context) []dto.PostResponse {
	posts := p.PostRepository.FindAll(ctx)

	var postResp []dto.PostResponse

	for _, value := range posts {
		post := dto.PostResponse{
			Id:          value.Id,
			Title:       value.Title,
			SubTitle:    *value.SubTitle,
			Published:   value.Published,
			Description: value.Description,
		}
		postResp = append(postResp, post)
	}

	return postResp
}

func (p *PostServiceImpl) FindById(ctx context.Context, postId string) dto.PostResponse {
	post, err := p.PostRepository.FindById(ctx, postId)
	helper.ErrorPanic(err)

	postResponse := dto.PostResponse{
		Id:          post.Id,
		Title:       post.Title,
		SubTitle:    *post.SubTitle,
		Published:   post.Published,
		Description: post.Description,
	}

	return postResponse
}

func (p *PostServiceImpl) Update(ctx context.Context, request dto.PostUpdate) {
	postData := model.Post{
		Id:          request.Id,
		Title:       request.Title,
		SubTitle:    request.SubTitle,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Update(ctx, postData)
}
