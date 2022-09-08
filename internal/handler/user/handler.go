package user

type (
	Id struct {
		Uuid string `json:"uuid" uri:"uuid" binding:"required"`
	}
	UserRequest struct {
		UserId int `json:"uuid" uri:"uuid" binding:"required"`
	}
	ProductRequest struct {
		Id int `json:"id" uri:"id" binding:"required"`
		// Uuid int `json:"uuid" uri:"uuid" binding:"required"`
	}
	ProductServiceRequest struct {
		FileName string `json:"fileName" uri:"fileName" binding:"required"`
	}
)
