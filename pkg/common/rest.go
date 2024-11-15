package common

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	v1 "lite-frame/apis/v1"
	"log"
	"net/http"
)

func ReadRequestBody(request *http.Request) []byte {

	// 从原有Request.Body读取
	bodyBytes, _ := ioutil.ReadAll(request.Body)

	// 新建缓冲区并替换原有Request.body
	request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes
}

func BuildPageParams(c *gin.Context) v1.Page {
	page := v1.Page{}
	err := c.ShouldBindQuery(&page)
	if err != nil {
		log.Fatalf("build page param failed error is %s", err.Error())
	}
	return page
}
