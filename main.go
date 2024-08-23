package main

import (
	"log"
	"net/http"
	"os"
)

func main(){
	// 打开一个文件，用于写入日志
	logFile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}
	defer logFile.Close()

	// 设置日志的输出目标为文件
	log.SetOutput(logFile)

	// 创建一个用于输出到控制台的日志实例
	consoleLogger := log.New(os.Stdout, "", log.LstdFlags)

	// 设置路由处理函数
	http.HandleFunc("/status200", handleStatus200)
	http.HandleFunc("/status500", handleStatus500)

	// 启动HTTP服务器，监听在端口8080
	consoleLogger.Println("do start")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		consoleLogger.Println("start success")
	}
}


func handleStatus200(w http.ResponseWriter, r *http.Request) {
	// 记录请求日志
	log.Printf("Received request for /status200 from %s, param %s", r.RemoteAddr, r.URL.Query().Get("a"))

	// 返回状态码200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.URL.Query().Get("a")))
}

func handleStatus500(w http.ResponseWriter, r *http.Request) {
	// 记录请求日志
	log.Printf("Received request for /status500 from %s", r.RemoteAddr)

	// 返回状态码500
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}
