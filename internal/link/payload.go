package link

type CreateLink struct {
	Url string `json:"url" validate:"required,url"`
}

type UpdateLink struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash,omitempty"`
}
