package redis



import (
	"github.com/robfig/cron"
	//"time"
	"fmt"
	"time"
)


func redisJob()  {
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		zset :=getSoredSetByRange("ZSET",0,10,true)
		for _, set := range zset {
			fmt.Println(zetScore("ZSET",set))
			if zetScore("ZSET",set) < time.Now().Unix(){
				insertList("LIST",set)
				remZet(set)
				fmt.Println(set)
			}
		}
	})
	c.AddFunc(spec, func() {
		if countList("LIST") > 0{

		}
	})
	c.Start()
}
func init() {
	redisJob()
}