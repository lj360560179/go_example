package redis



import (
	"github.com/robfig/cron"
	"time"
	"strings"
)


func redisJob()  {
	c := cron.New()
	spec := "*/5 * * * * ?"

	//获取排序列表中时间最靠前的十个，如果达到执行时间则加入LIST中
	c.AddFunc(spec, func() {
		zset :=getSoredSetByRange("ZSET",0,10,true)
		for _, set := range zset {
			if zetScore("ZSET",set) < time.Now().Unix(){
				insertList("LIST",set)
				remZet(set)
			}
		}
	})

	//去除LIST中的数据执行业务
	c.AddFunc(spec, func() {
		count := countList("LIST")
		if count > 0{
			list:=rangeList("LIST",0,count - 1)
			if list!= nil{
				for _,s := range list{
					//不知道为毛取出来会带有[]，所以要删掉
					//执行业务逻辑
					saveMongo(getString(strings.TrimRight(strings.TrimLeft(s,"["),"]")))
					//删掉删掉~
					remListValue("LIST",s,1)
					//删掉删掉~
					remoceString(s)
				}
			}
		}
	})
	c.Start()
}
func init() {
	redisJob()
}