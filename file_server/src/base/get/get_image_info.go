// 响应客户端请求,并返回相应的头像路径
package get

import (
	"net/http"
)

// 处理请求
func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "it works! return a Url", http.StatusAccepted)
	} else {
		http.Error(w, "Not GET request", http.StatusInternalServerError)
		return
	}
}