package proto

type GetTopicReq struct {
	UserName string `json:"user_name" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required" `
	PageSize int    `json:"page_size" form:"page_size"`
}
