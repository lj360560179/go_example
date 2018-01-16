#部署golang环境
FROM golang:1.9
#下载gin包
RUN go get github.com/gin-gonic/gin
#开放端口
EXPOSE 7070