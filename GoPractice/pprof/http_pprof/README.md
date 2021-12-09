### 参考
https://segmentfault.com/a/1190000016412013

https://juejin.cn/post/6844904079525675016

### 过程
- 导入`net/http/pprof`包，并运行web服务
- 执行pprof，会生成文件以及进入命令行`go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60`
- 可以直接在pprof命令中查看：help查看帮助
- 也可以在启用web服务查看生成的文件`go tool pprof -http 127.0.0.1:8080 /Users/admin/pprof/pprof.samples.cpu.001.pb.gz`
