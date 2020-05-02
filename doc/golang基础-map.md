# golang基础--map

1.初始化

```go
// 先声明后初始化
var amap map[string]int
amap=make(map[string]int)
amap["tao"] = 18
// 直接初始化
smap := map[string]int{
    "tao":18,
}
tmap := make(map[string]int)
tmap["tao"] = 18
```



2.判断一个值是否在map中

```go
if num,ok :=smap[2];ok{
        fmt.Println("num is exist:",num)
}
```

3.遍历一个map

```go
for key,value := range sMap{
    fmt.Println(key,value)
}
```





























参考资料：

- [GoLang基础数据类型--->字典（map）详解](https://www.cnblogs.com/yinzhengjie/p/7689996.html)