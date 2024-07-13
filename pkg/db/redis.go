package db

import "github.com/liyong2922/TangSengDaoDaoServerLib/pkg/redis"

func NewRedis(addr string, password string) *redis.Conn {
	return redis.New(addr, password)
}
