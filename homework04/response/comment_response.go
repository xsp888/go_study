package response

import "homework04/models"

type CommentResponse struct {
	Comments []models.Comment
	Total    int64
	Page     int
	PageSize int
}
