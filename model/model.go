package model

type Domain struct {
	Url string `json:"url" binding:"required"`
}
