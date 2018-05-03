## go gin redis delay

利用redis的有序队列来模拟延迟队列，时间作为分数存入ZSET中，定时从ZSET中取出分数最低的
的与当前时间进行比较，如果达到任务执行时间则取出放入LIST中然后执行业务逻辑