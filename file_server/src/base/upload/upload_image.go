// 获取请求并将文件进行存储
package upload

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"strings"
	"path"
	"regexp"

)

// 获取文件大小的接口
type Size interface {
    Size() int64
}

// 判断后缀是否支持,支持返回true
func judge_suffix(suffix string) bool  {
	// 支持的文件格式
	Supported_Format := []string{".jpeg", ".jpg", ".png"}

	sum := 0
	for _, v := range Supported_Format {
		if strings.ToLower(suffix) == v {
			sum++
		}
	}
	if sum == 1 {
		return true
	} else {
		return false
	}
}

// 处理上传文件
func UploadHandler(w http.ResponseWriter, r *http.Request) {

    if r.Method == "POST" {
		// 判断user_id是不是纯数字,如果不是则返回...
		user_id := r.FormValue("user_id")
		if rlt,_ := regexp.MatchString("[^0-9]", user_id); rlt {
			http.Error(w, "user_id must be all digits.", http.StatusNotFound)
			return
		}

		// 判断Content-Type是否为multipart/form-data,如果不是则返回...
		content_type := strings.Split(r.Header["Content-Type"][0], ";")
		if content_type[0] != "multipart/form-data" {
			http.Error(w, "Content-Type must be multipart/form-data.", http.StatusNotFound)
			return
		}

		// 调用MultipartForm来获取post内容,并判断是否post内容是否为空,如果为空则返回...
		r.ParseMultipartForm(32 << 20)
		if r.MultipartForm == nil {
			http.Error(w, "Cannot recive post data.", http.StatusNotFound)
			return
		}

		// 使用FormFile来获取文件,myfile为key值,如果获取失败则返回...
		file, handler, err := r.FormFile("myfile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		defer file.Close()

		// 判断文件大小
		if sizeInterface, ok := file.(Size); ok {
    		file_size := sizeInterface.Size() / (1024*1024)
			if file_size > 20 {
				http.Error(w, "Does not support file over 20MB", http.StatusNotFound)
				return
			}
		} else {
			fmt.Println("无法获取文件大小")
			return
		}

		// 判断文件后缀是否合法
		suffix := strings.ToLower(path.Ext(handler.Filename))
		if rlt := judge_suffix(suffix); ! rlt {
			http.Error(w, "Unsupported suffix", http.StatusNotFound)
			return
		}
		// 写入
		f, err := os.OpenFile("/Users/guoxu/code/go/http/file_server/file/" + user_id + suffix, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		// 图片处理
		rlt := Resize("/Users/guoxu/code/go/http/file_server/file/", user_id + suffix)
		if rlt {
			http.Error(w, "Image resized ok!", http.StatusAccepted)
		}
	} else  {
		http.Error(w, "Not post request", http.StatusNotFound)
		return
	}
}

