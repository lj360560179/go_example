package redis



import (
	"github.com/robfig/cron"
	"time"
	"fmt"
)


func redisJob()  {
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		zset :=getSoredSetByRange("ZSET",0,10,true)
		for _, set := range zset {
			fmt.Println(zetScore("ZSET",set))
			if zetScore("ZSET",set) > time.Now().UnixNano(){
				insertList("LIST",set)
				remZet(set)
				fmt.Println(set)
			}
		}
	})
	c.Start()
}
func init() {
	redisJob()
}