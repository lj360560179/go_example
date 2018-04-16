#部署golang环境
FROM golang:1.9
#下载gin包
RUN go get github.com/gin-gonic/gin \
    go get gopkg.in/mgo.v2/bson
#开放端口
EXPOSE 7070
CMD