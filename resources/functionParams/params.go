package functionParams

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out/common/result"
	"sky-take-out/resources/commonParams"
	"strconv"
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

func GetUser(in interface{}) int64 {
	user, ok := in.(float64)
	if !ok {
		return 0
	}
	return int64(user)
}

func ToInt(in interface{}) int {
	i, ok := in.(int)
	if ok {
		return i
	}
	i64, ok := in.(int64)
	if ok {
		return int(i64)
	}
	f64, ok := in.(float64)
	if ok {
		return int(f64)
	}
	f32, ok := in.(float32)
	if ok {
		return int(f32)
	}
	str, ok := in.(string)
	if ok {
		i, err := strconv.Atoi(str)
		if err == nil {
			return i
		}
	}
	return -1
}

func Rollback() {
	commonParams.Tx.Rollback()
	commonParams.Tx = nil
}

func Commit() error {
	err := commonParams.Tx.Commit()
	commonParams.Tx = nil
	return err
}

func ExecSQL(SQL string, args []interface{}) (res sql.Result, err error) {
	if commonParams.Tx == nil {
		res, err = commonParams.Db.Exec(SQL, args...)
	} else {
		res, err = commonParams.Tx.Exec(SQL, args...)
	}
	return res, err
}
