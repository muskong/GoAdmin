package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoCore/respond"
)

type _upload struct {
	fileKey string
}

var UploadHandler = &_upload{
	fileKey: "file",
}

func (u *_upload) Upload(c *gin.Context) {
	header, err := c.FormFile(u.fileKey)
	if err != nil {
		c.SecureJSON(respond.Message("传入参数错误"))
		return
	}
	dst := header.Filename
	// gin 简单做了封装,拷贝了文件流
	if err := c.SaveUploadedFile(header, dst); err != nil {
		c.SecureJSON(respond.Message(err.Error()))
		return
	}

	c.SecureJSON(respond.Data(dst))
}
