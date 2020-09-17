package model

type Topic struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"min=4,max=20"`
	Url   string `json:"url" binding:"omitempty,topurl"`
}

func NewTopic() *Topic {
	return &Topic{}
}
