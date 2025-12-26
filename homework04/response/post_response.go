package response

import "homework04/models"

type PostResponse struct {
	Posts    []models.Post
	Total    int64
	Page     int
	PageSize int
}
