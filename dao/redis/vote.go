package redis

import (
	"github.com/go-redis/redis"
	"time"
)

// 推荐阅读
// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm/

// 本项目使用简化版的投票分数
// 投一票就加432分   86400/200  --> 200张赞成票可以给你的帖子续一天

/* 投票的几种情况：
direction=1时，有两种情况：
	1. 之前没有投过票，现在投赞成票    --> 更新分数和投票记录
	2. 之前投反对票，现在改投赞成票    --> 更新分数和投票记录
direction=0时，有两种情况：
	1. 之前投过赞成票，现在要取消投票  --> 更新分数和投票记录
	2. 之前投过反对票，现在要取消投票  --> 更新分数和投票记录
direction=-1时，有两种情况：
	1. 之前没有投过票，现在投反对票    --> 更新分数和投票记录
	2. 之前投赞成票，现在改投反对票    --> 更新分数和投票记录

投票的限制：
每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
	1. 到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2. 到期之后删除那个 KeyPostVotedZSetPF
*/
const (
	oneWeekInSecond = 7 * 24 * 3600
	scorePerVote    = 432
)

func VoteForPost(userID, postID string, value float64) error {
	// 1.判断投票时间限制
	//从redis中找到帖子的发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSecond {
		return ErrVoteTimeExpire
	}
	// 2和3需要放到一个pipline中去执行
	orginValue := rdb.ZScore(getRedisKey(KeyPostScoreZSet+postID), userID).Val()
	diff := value - orginValue

	//开启事务
	pipeline := rdb.TxPipeline()
	//2.更新帖子的分数
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), scorePerVote*diff, postID)
	//3.记录用户的投票数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Member: userID,
			Score:  value,
		})
	}
	_, err := pipeline.Exec()
	return err
}
