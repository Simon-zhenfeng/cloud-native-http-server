package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"http-server/metrics"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

/**
2. 日志以标准输出方式
*/

const EnvVersion = "VERSION"

func main() {
	metrics.Register()
	flag.Set("v", "4")
	glog.V(2).Info("Starting up http server.")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/images", imageHandler)
	http.HandleFunc("/healthz", healthzHandler)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		glog.V(2).Info("http server starting failed: ", err)
	}
}

func imageHandler(writer http.ResponseWriter, request *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	writer.WriteHeader(http.StatusOK)
	io.WriteString(writer, fmt.Sprintf("<h1>%d</h1>", randInt))
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
