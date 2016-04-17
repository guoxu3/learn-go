// 用户头像存储服务
package main

import "net/http"
import	(
	"base/upload"
	"base/get"
)

func main()  {
	http.HandleFunc("/file", upload.UploadHandler)
	http.HandleFunc("/get", get.GetHandler)
	http.ListenAndServe(":8899", nil)
}

