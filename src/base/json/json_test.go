package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type (
	Server struct {
		Name string `bson:"name" json:"name"` // 名称
	}
)

// struct->json 转换
func Struct2Json(data interface{}) {
	bytes, _ := json.Marshal(data)
	fmt.Println(string(bytes))
}

func TestStruct2Json(t *testing.T) {
	Struct2Json(Server{})
}
