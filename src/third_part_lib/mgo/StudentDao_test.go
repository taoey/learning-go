package mgo

import (
    "fmt"
    "github.com/Taoey/Lets-Go/src/third_part_lib/mgo/sysinit"
    "gopkg.in/mgo.v2/bson"
    "log"
    "math/rand"
    "reflect"
    "strconv"
    "strings"
    "testing"
    "time"
)

//自定义日志
func init() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestInsertOne(t *testing.T) {
    persionGV := Person{"gv", 20}
    query := FindOne(bson.M{
        "name": "gv",
    })
    if reflect.DeepEqual(query, Person{}) {
        InsertOne(persionGV)
    }
    fmt.Println(query)
}

func TestFind(t *testing.T) {

    SQL := bson.M{
        "name": "tao",
    }
    list := Find(SQL)
    log.Println(list)
}

func TestInsertDemo(t *testing.T) {

    start := time.Now().UnixNano() / 1e6
    time.Sleep(time.Duration(2) * time.Second)

    sysinit.InitMongo()
    session := sysinit.MongoSession.Copy()
    defer session.Close()

    collection := session.DB("").C("student")
    for i := 0; i < 100000; i++ {
        SQL := bson.M{
            "uid":  i,
            "name": strings.Join([]string{"S", strconv.Itoa(i)}, ""),
            "age":  rand.Intn(80),
        }
        collection.Insert(SQL)
    }

    end := time.Now().UnixNano() / 1e6
    fmt.Println(end - start)
}

func TestString(t *testing.T) {
    join := strings.Join([]string{"S", strconv.Itoa(1)}, "")
    fmt.Println(join)
}

func TestInsertGradeDemo(t *testing.T) {

    start := time.Now().UnixNano() / 1e6
    time.Sleep(time.Duration(2) * time.Second)

    sysinit.InitMongo()
    session := sysinit.MongoSession.Copy()
    defer session.Close()

    collection := session.DB("").C("grade")
    for i := 0; i < 100000; i++ {
        SQL := bson.M{
            "sid":     i,
            "math":    rand.Intn(100),
            "english": rand.Intn(100),
            "history": rand.Intn(100),
            "bid":     1,
        }
        collection.Insert(SQL)
    }

    end := time.Now().UnixNano() / 1e6
    fmt.Println(end - start)
}

func TestTime(t *testing.T) {
    start := time.Now().UnixNano() / 1e6
    time.Sleep(time.Second * 2)

    end := time.Now().UnixNano() / 1e6
    fmt.Println(end - start)
}
