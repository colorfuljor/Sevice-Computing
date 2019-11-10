## 使用Martini框架开发 web 服务程序
完整代码在[cloudgo](./cloudgo)目录下

代码中用到了一些库，需要安装：
```
> go get github.com/spf13/pflag
> go get github.com/go-martini/martini
> go get github.com/unrolled/render
```

而且如果不是把cloudgo放在$GOPATH/github.com/github-user目录下，则需要更改main.go中import service包的路径。
