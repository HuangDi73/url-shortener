package link

type CreateLink struct {
	Url string `json:"url" validate:"required,url"`
}

type UpdateLink struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash,omitempty"`
}

type GetAllLinksResponse struct {
	Links []Link `json:"links"`
	Count int64  `json:"count"`
}
