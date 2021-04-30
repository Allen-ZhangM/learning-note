### 来源
```
https://mp.weixin.qq.com/s/bF-2lZPoYbHZrNxJqGOn1g
```

### 运行

```
go run trace.go
Hello World
```

会得到一个trace.out文件，然后我们可以用一个工具打开，来分析这个文件。

```
$ go tool trace trace.out
2020/02/23 10:44:11 Parsing trace...
2020/02/23 10:44:11 Splitting trace...
2020/02/23 10:44:11 Opening browser. Trace viewer is listening on http://127.0.0.1:33479
```

我们可以通过浏览器打开http://127.0.0.1:33479网址，点击view trace 能够看见可视化的调度流程。