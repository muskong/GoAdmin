package public

type (
	Jwt struct {
		Token string `json:"token" uri:"token" binding:"required"`
	}
)
