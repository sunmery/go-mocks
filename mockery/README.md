## start

IDE:
在`main.go`的接口点击运行`Go Generate mockery --name DB`

命令行:
进入到本目录
```shell
go install github.com/vektra/mockery/v2@v2.40.3
cd go-mocks/mockery
go mod tidy
mockery --name DB
```

## 自定义yaml
https://vektra.github.io/mockery/latest/configuration/#parameter-descriptions

## 参考资料
1. https://vektra.github.io/mockery/latest/
2. https://github.com/vektra/mockery
3. https://www.cnblogs.com/guoapeng/p/17359565.html
