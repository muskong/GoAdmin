package handler

type (
	Id struct {
		Nanoid string `uri:"nanoid", binding:"required"`
	}
)
