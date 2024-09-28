package post

import (
	"context"
	dto "go-arepas/domain/post/dto"
)

type PostService interface {
	Create(ctx context.Context, request dto.PostCreate)
	Update(ctx context.Context, request dto.PostUpdate)
	Delete(ctx context.Context, postId string)
	FindById(ctx context.Context, postId string) dto.PostResponse
	FindAll(ctx context.Context) []dto.PostResponse
}
