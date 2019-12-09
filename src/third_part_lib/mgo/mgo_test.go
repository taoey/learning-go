package mgo

import (
    "gopkg.in/mgo.v2"
    "log"
    "testing"
)

// 数据库连接测试
func TestConnect(t *testing.T) {
    url  := "127.0.0.1:27017/tao_test"
    MongoSession, err := mgo.Dial(url)
    if err != nil {
        log.Print(err)
    }else{
    }
    MongoSession.SetMode(mgo.Monotonic, true) //连接模式设置
    MongoSession.SetPoolLimit(2000)           // 设置连接池数量
    session := MongoSession.Copy()
    c := session.DB("").C("student")
    data := map[string]string{}
    data["tao"] = "111"
    c.Insert(data)

}
