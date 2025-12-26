package request

type CommentRequest struct {
	PostID   uint
	Page     int
	PageSize int
}
