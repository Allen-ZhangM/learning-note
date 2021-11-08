### 思路
```
利用gomock生成接口的server和client，然后模拟调用
https://github.com/jetstack/navigator/blob/master/vendor/google.golang.org/grpc/Documentation/gomock-example.md
```

### 生成
```
mockgen -source=/go/src/google.golang.org/grpc/examples/helloworld/helloworld/helloworld_grpc.pb.go -destination=grpc_mock_gen.go -package=main
```