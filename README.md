

`TODO:UNKNOWN` 待理解的标记

### 目录

* `client`  客户端二方包
    * TODO
    * `callopt`  调用的配置选项
    * `genericclient`  泛化调用的客户端
* `internal`  私有的应用程序代码库
    * TODO
* `pkg`  外部应用程序可以使用的库代码
    * `acl`  访问控制列表ACL（Access Control List）
    * TODO
    * `generic` 泛化调用相关的
        * 支持二进制、http、map、json映射泛化调用，对应文件 binarythrift_codec.go httpthrift_codec.go jsonthrift_codec.go mapthrift_codec.go

        * `remote`  .
            * TODO:核心
    * `klog` [link](./pkg/klog)
    ```text
    字节RPC框架Kitex的日志库klog竟然也这么小巧！    https://juejin.cn/post/7104825435795980324
    ```
* `server`  服务端二方包
    * `genericserver`  泛化服务
        * TODO
    * `invoke`  .
        * TODO
    * TODO
* `tool`
    * `cmd`  命令行工具那块的
    * `internal`  私有的应用程序代码库
        * TODO
* `transport`  协议转换那块的

===


```sh
$ find . | grep -v "\.[go|proto|iml|xml|yml|sum]" | grep -v "licenses" | grep -v "LICENSE" | grep -v "images" | grep -v "NOTICE"



.
./transport
./internal
./internal/mocks
./internal/mocks/limiter
./internal/mocks/thrift
./internal/mocks/remote
./internal/configutil
./internal/test
./internal/wpool
./internal/server
./internal/client
./internal/stats
./server
./server/invoke
./server/genericserver
./client
./client/callopt
./client/genericclient
./CREDITS
./tool
./tool/cmd
./tool/cmd/kitex
./tool/internal
./tool/internal/pkg
./tool/internal/pkg/util
./tool/internal/pkg/pluginmode
./tool/internal/pkg/pluginmode/thriftgo
./tool/internal/pkg/pluginmode/protoc
./tool/internal/pkg/tpl
./tool/internal/pkg/generator
./tool/internal/pkg/log
./pkg
./pkg/rpcinfo
./pkg/rpcinfo/remoteinfo
./pkg/loadbalance
./pkg/loadbalance/lbcache
./pkg/proxy
./pkg/retry
./pkg/connpool
./pkg/limit
./pkg/gofunc
./pkg/limiter
./pkg/rpctimeout
./pkg/discovery
./pkg/utils
./pkg/protocol
./pkg/protocol/bprotoc
./pkg/protocol/bthrift
./pkg/streaming
./pkg/serviceinfo
./pkg/transmeta
./pkg/http
./pkg/endpoint
./pkg/kerrors
./pkg/registry
./pkg/generic
./pkg/generic/binary_test
./pkg/generic/map_test
./pkg/generic/map_test/idl
./pkg/generic/map_test/conf
./pkg/generic/json_test
./pkg/generic/json_test/idl
./pkg/generic/json_test/conf
./pkg/generic/thrift
./pkg/generic/descriptor
./pkg/generic/http_test
./pkg/generic/http_test/idl
./pkg/generic/http_test/conf
./pkg/klog
./pkg/exception
./pkg/event
./pkg/consts
./pkg/warmup
./pkg/acl
./pkg/stats
./pkg/circuitbreak
./pkg/remote
./pkg/remote/remotesvr
./pkg/remote/connpool
./pkg/remote/codec
./pkg/remote/codec/thrift
./pkg/remote/codec/perrors
./pkg/remote/codec/protobuf
./pkg/remote/remotecli
./pkg/remote/transmeta
./pkg/remote/bound
./pkg/remote/trans
./pkg/remote/trans/netpoll
./pkg/remote/trans/invoke
./pkg/remote/trans/netpollmux
./pkg/remote/trans/nphttp2
./pkg/remote/trans/nphttp2/grpc
./pkg/remote/trans/nphttp2/grpc/syscall
./pkg/remote/trans/nphttp2/grpc/testutils
./pkg/remote/trans/nphttp2/grpc/testutils/leakcheck
./pkg/remote/trans/nphttp2/status
./pkg/remote/trans/nphttp2/codes
./pkg/remote/trans/nphttp2/metadata
./pkg/remote/trans/detection
./pkg/diagnosis
```

```shell
$ find . | grep "\.go" | grep -v "test" | xargs wc -l
   41769 total

```
