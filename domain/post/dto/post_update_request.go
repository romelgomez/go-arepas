package post

type PostUpdate struct {
	Id          string  `json:"id"`
	Title       string  `validate:"required min=1,max=100" json:"title"`
	SubTitle    *string `validate:"min=1,max=100" json:"sub_title"`
	Published   bool    `json:"published"`
	Description string  `json:"description"`
}
