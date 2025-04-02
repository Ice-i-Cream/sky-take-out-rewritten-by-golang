package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/functionParams"
	"strings"
)

type CommonController struct {
}

func (c *CommonController) Upload(ctx *gin.Context) {
	exec := func(ctx *gin.Context) (string, error) {
		f, h, err := ctx.Request.FormFile("file")
		if err != nil {
			return "", nil
		}
		u := uuid.New()
		uuidStr := u.String()
		ext := strings.TrimPrefix(h.Filename[strings.LastIndex(h.Filename, "."):], ".")
		if ext == h.Filename {
			ext = ""
		} else {
			ext = "." + ext
		}
		fileName := uuidStr + ext
		var out *os.File
		out, err = os.Create("file/" + fileName)
		if err != nil {
			return "", err
		}
		_, err = io.Copy(out, f)
		if err != nil {
			return "", err
		}
		host := commonParams.ServerHost
		port := commonParams.ServerPort
		url := fmt.Sprintf("http://%s:%s/%s", host, port, "image/"+fileName)
		log.Println(url)
		err = f.Close()
		if err != nil {
			return "", err
		}
		err = out.Close()
		if err != nil {
			return "", err
		}
		return url, nil
	}
	url, err := exec(ctx)
	functionParams.PostProcess(ctx, err, url)
}
