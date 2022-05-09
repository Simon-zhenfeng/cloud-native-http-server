FROM golang:1.17 AS build
WORKDIR /http-server/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o http-server main.go

FROM busybox
COPY --from=build /http-server/http-server /http-server/http-server
EXPOSE 8360
ENV ENV local
WORKDIR /http-server/
ENTRYPOINT ["./http-server"]