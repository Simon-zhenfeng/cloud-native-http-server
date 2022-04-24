package main

import (
	"flag"
	"github.com/golang/glog"
	"io"
	"net/http"
	"os"
	"strings"
)

/**
2. 日志以标准输出方式
*/

const EnvVersion = "VERSION"

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting up http server.")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthzHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		glog.V(2).Info("http server starting failed: ", err)
	}
}

func healthzHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	io.WriteString(writer, "I'm healthy.\n")
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	glog.V(2).Infof("客户端请求IP: %s", request.RemoteAddr)
	setHeader(writer.Header(), request.Header)
	writer.WriteHeader(http.StatusOK)
	glog.V(2).Infof("服务成功返回Http Status: %d", http.StatusOK)
	io.WriteString(writer, "ok\n")
}

func setHeader(responseHeader http.Header, requestHeader http.Header) {
	responseHeader.Add(EnvVersion, os.Getenv(EnvVersion))
	for name, value := range requestHeader {
		responseHeader.Add(name, strings.Join(value, ", "))
	}
}
