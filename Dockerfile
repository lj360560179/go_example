#部署golang环境
FROM golang:1.9
#下载gin包
RUN go get github.com/gin-gonic/gin \
    && go get gopkg.in/mgo.v2/bson \
    && go get github.com/jinzhu/gorm \
    && go get github.com/robfig/cron \
    && go get github.com/garyburd/redigo/redis
#开放端口
EXPOSE 707
ADD main.exe /homt/main.exe
CMD main
