#部署golang环境
FROM golang:1.9
#下载gin包
RUN go get github.com/gin-gonic/gin \
    && go get gopkg.in/mgo.v2 \
    && github.com/jinzhu/gorm \
    && github.com/robfig/cron \
    && github.com/jinzhu/gorm/dialects/mysql
#开放端口
EXPOSE 8080
RUN mkdir /app
ADD main.exe /app/main.exe
WORKDIR  /app
CMD main