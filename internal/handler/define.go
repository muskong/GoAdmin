package handler

type (
	Id struct {
		Nanoid string `uri:"nanoid", binding:"required"`
	}
	UserResquet struct {
		UserId int `uri:"nanoid", binding:"required"`
	}
)
