package redis

import (
	"github.com/go-redis/redis"
	"time"
)

// CreatePost 创建帖子
func CreatePost(postID int64) error {
	pipline := rdb.TxPipeline()
	//帖子时间
	pipline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//帖子分数
	pipline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipline.Exec()
	return err
}
