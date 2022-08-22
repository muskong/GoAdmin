package handler

type (
	Id struct {
		Nanoid string `json:"nanoid" uri:"nanoid" binding:"required"`
	}
	UserRequest struct {
		UserId int `json:"nanoid" uri:"nanoid" binding:"required"`
	}
	ProductRequest struct {
		Id   int `json:"id" uri:"id" binding:"required"`
		Uuid int `json:"uuid" uri:"uuid" binding:"required"`
	}
)
