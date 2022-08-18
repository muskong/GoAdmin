package handler

type (
	Id struct {
		Nanoid string `json:"nanoid" uri:"nanoid" binding:"required"`
	}
	UserRequest struct {
		UserId int `json:"nanoid" uri:"nanoid" binding:"required"`
	}
)
