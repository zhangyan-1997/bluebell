package redis

const (
	Prefix             = "bluebell:"   //项目key的前缀
	KeyPostTimeZSet    = "post:time"   //zset帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  //zset帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" //zset记录用户及投票类型,参数是post_id
)

//给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
