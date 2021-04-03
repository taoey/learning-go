package mgo

import (
	"gopkg.in/mgo.v2"
	"log"
)

//自定义日志
func init() {
	InitLog()
	InitMongo()
}

const URL string = "127.0.0.1:27017/tao_test"

var MongoSession *mgo.Session

func InitMongo() {
	MongoSession, _ = mgo.Dial(URL)
	MongoSession.SetMode(mgo.Monotonic, true) //连接模式设置
	MongoSession.SetPoolLimit(2000)           // 设置连接池数量
}

func InitLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
