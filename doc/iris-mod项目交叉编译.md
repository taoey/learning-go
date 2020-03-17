# iris-mod项目交叉编译



```
SET GOARCH=amd64
SET GOOS=linux
go build  cmd/iris-cil-server.go
```





发现一个坑，需要把主程序的包设置为main，如果为其他不行

```
SET GOARCH=amd64 
SET GOOS=linux 
go build  cmd/smartrtb_nginx_server.go
```



