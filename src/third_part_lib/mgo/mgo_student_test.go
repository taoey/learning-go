package mgo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type Student struct {
	Name string `bson:"name" json:"name"`
	Age  int    `bson:"age" json:"age"`
}

// service 插入一个学生
func TestSerInsertStudent01(t *testing.T) {
	session := MongoSession.Copy()
	defer session.Close()

	dataMap := map[string]interface{}{}
	dataMap["name"] = "tao"
	dataMap["age"] = 19

	session.DB("").C("student").Insert(dataMap)
}

// 简单查询
func TestFind01(t *testing.T) {
	session := MongoSession.Copy()
	defer session.Close()

	filter := bson.M{
		"name": "tao",
	}
	result := []Student{}
	session.DB("").C("student").Find(filter).All(&result)
	fmt.Println(result)
}

//复杂查询01
func TestFind02(t *testing.T) {

}
