package dao

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

func CreatePost(postID, communityID int64) error {
	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd("blog:post:time", redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd("blog:post:score", redis.Z{
		Score:  0,
		Member: postID,
	})
	// 更新：把帖子id加到社区的set
	cKey := "blog:community:" + strconv.Itoa(int(communityID))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
}

func PostVote(userID, postID string, value float64) error {
	// 2. 更新贴子的分数
	// 先查当前用户给当前帖子的投票记录
	ov := rdb.ZScore("blog:post:vote:"+postID, userID).Val()

	// 更新：如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if value == ov {
		return errors.New("不允许重复点赞")
	}

	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline := rdb.TxPipeline()

	pipeline.ZIncrBy("blog:post:score", op*diff, postID)

	// 3. 记录用户为该贴子投票的数据
	if value == 0 {
		pipeline.ZRem("blog:post:vote:"+postID, userID)
	} else {
		pipeline.ZAdd("blog:post:vote:"+postID, redis.Z{
			Score:  1,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
