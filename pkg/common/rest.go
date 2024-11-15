package common

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func ReadRequestBody(request *http.Request) []byte {

	// 从原有Request.Body读取
	bodyBytes, _ := ioutil.ReadAll(request.Body)

	// 新建缓冲区并替换原有Request.body
	request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes
}
