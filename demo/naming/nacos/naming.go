package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func init() {

	// 创建clientConfig的另一种方式
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId("test"), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// 创建serverConfig的另一种方式
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			"110.40.141.168",
			8848,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}
	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "110.40.141.168",
		Port:        8848,
		ServiceName: "jdd-order-center",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		GroupName:   "test", // 默认值DEFAULT_GROUP
	})

	// SelectAllInstance可以返回全部实例列表,包括healthy=false,enable=false,weight<=0
	instances, _ := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: "jdd-order-center",
		GroupName:   "test", // 默认值DEFAULT_GROUP
	})
	for i, instance := range instances {
		fmt.Println(i)
		fmt.Println(instance.Ip)
		fmt.Println(instance.Port)
	}
}

func main() {

}
