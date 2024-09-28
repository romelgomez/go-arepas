package post

type PostResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	SubTitle    string `json:"sub_title"`
	Published   bool   `json:"published"`
	Description string `json:"description"`
}
