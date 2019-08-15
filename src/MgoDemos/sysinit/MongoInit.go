package sysinit

import (
	"gopkg.in/mgo.v2"
)

const URL string = "127.0.0.1:27017"
const DB string = "test"

var MongoSession *mgo.Session

func InitMongo() {
	MongoSession, _ = mgo.Dial(URL)
	MongoSession.SetMode(mgo.Monotonic, true) //连接模式设置
	MongoSession.SetPoolLimit(2000)           // 设置连接池数量
}
