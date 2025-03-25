package functionParams

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out/common/result"
	"strings"
)

func PostProcess(c *gin.Context, err error, data interface{}) {
	var res interface{}
	if err != nil {
		res = result.Error(err.Error())
		c.Status(http.StatusInternalServerError) // 设置状态码
	} else {
		if data != nil {
			res = result.Success(data)
		} else {
			res = result.SuccessNoData()
		}
	}
	c.JSON(http.StatusOK, res) // 发送 JSON 响应
}

func IsExcludedPath(excludedPrefixes []string, excludedPaths []string, urlPath string) bool {

	for _, prefix := range excludedPrefixes {
		if strings.HasPrefix(urlPath, prefix) {
			return false
		}
	}
	for _, path := range excludedPaths {
		if urlPath == path {
			return false
		}
	}
	return true
}
