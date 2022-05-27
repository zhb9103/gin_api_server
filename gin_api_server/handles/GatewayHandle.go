package handles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"net/http"
	//greeter "src/api_gin_gateway/greeter/proto/greeter"
)

//具体的请求回调函数
func CallGreeter(ctx *gin.Context) {
	// 进行request的绑定，拿到一个结构，然后传递给我们的微服务
	//greeter.Request{} 是我们的编辑的GRPC里面结果提，我们定义的请求对象
	//request := greeter.Request{Name: "sadasdas"}
	//// 然后从我们的gin接收相关的传递过来的参数信息
	//ctx.ShouldBindQuery(&request)

	//定义注册中心的对象
	reg := etcd.NewRegistry(registry.Addrs("192.168.219.130:2379"))
	// 注册到我们的ectd上
	service := micro.NewService(
		micro.Registry(reg),
	)
	fmt.Println(service)
	//创建客户端对象
	//helloService := greeter.NewGreeterService("go.micro.service.greeter", service.Client())

	//context.Background():可以简单理解我们知道这个上下文要去干什么
	//context.TODO():可以简单理解我们不清楚要使用哪个上下文、或者还没有可用的上下文
	// 默认生成的hello服务中自带三个接口: Call() Stream() PingPong(),分别对应参数调用、流传输和心跳
	//resp, err := helloService.Call(context.Background(), &greeter.Request{
	//	Name: "xiao xie",
	//})

	//// 然后像我们的微服务提交请求
	//res, err := serviser.greeterClient.Call(context.TODO(), &request)
	//if err != nil {
	//	log.Print(err.Error())
	//	ctx.JSON(http.StatusInternalServerError, gin.H{
	//		"code":    -2,
	//		"message": "server error",
	//	})
	//	return
	//}
	ctx.JSON(http.StatusOK, gin.H{
		"jieguo": "hello",
	})
}
