package main

import (
	"fmt"
	"golang-filestore/handler"
	"net/http"
)

func main() {
	// 静态资源处理
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	// 静态接口路由设置
	http.HandleFunc("/file/upload", handler.HTTPIntercepter(handler.UploadHandler))
	http.HandleFunc("/file/upload/suc", handler.HTTPIntercepter(handler.UploadSucHandler))
	http.HandleFunc("/file/meta", handler.HTTPIntercepter(handler.GetFileMetaHandler))
	http.HandleFunc("/file/query", handler.HTTPIntercepter(handler.FileQueryHandler))
	http.HandleFunc("/file/download", handler.HTTPIntercepter(handler.DownloadHandler))
	http.HandleFunc("/file/download/range", handler.HTTPIntercepter(handler.RangeDownloadHandler))
	http.HandleFunc("/file/update", handler.HTTPIntercepter(handler.FileMetaUpdateHandler))
	http.HandleFunc("/file/delete", handler.HTTPIntercepter(handler.FileDeleteHandler))

	// 秒传接口
	http.HandleFunc("/file/fastupload", handler.HTTPIntercepter(handler.TryFastUploadHandler))

	// 分块上传几口
	http.HandleFunc("/file/mpupload/init", handler.HTTPIntercepter(handler.InitialMultipartUploadHandler))
	http.HandleFunc("/file/mpupload/uppart", handler.HTTPIntercepter(handler.UploadPartHandler))
	http.HandleFunc("/file/mpupload/complete", handler.HTTPIntercepter(handler.CompleteUploadHandler))

	// 用户相关接口
	http.HandleFunc("/user/signup", handler.SignupHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/info", handler.HTTPIntercepter(handler.UserInfoHandler))

	// 监听端口
	fmt.Println("上传服务正在启动，监听端口：8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
