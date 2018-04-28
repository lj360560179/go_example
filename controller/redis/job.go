package redis



import (
	"github.com/robfig/cron"
	"time"
)


func redisJob()  {
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		zset :=getSoredSetByRange("ZSET",0,10,true)
		for _, set := range zset {
			if zetScore("ZSET",set) > time.Now().UnixNano(){
				insertList("LIST",set)
				remZet(set)
			}
		}
	})
	c.Start()
}
func init() {
	redisJob()
}