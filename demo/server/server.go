package naming

import (
	"fmt"
	"github.com/jim/demo/naming/nacos"
	"net/http"
)

type Server struct {
	Id        string
	Name      string
	Address   string
	Port      uint64
	Namespace string
	Group     string
	Meta      map[string]string
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello golang http By Nacos!")
}

func StartServer() {
	server := Server{
		Id:        "myServer",
		Name:      "myServer",
		Address:   "127.0.0.1",
		Port:      8999,
		Namespace: "test",
		Group:     "test",
	}
	newNaming, err := nacos.NewNaming("110.40.141.168", 8848, server.Namespace)
	newNaming.Register(server)
	http.HandleFunc("/", index)
	// 启动web服务，监听9090端口
	url := fmt.Sprintf("%s:%d", server.Address, server.Port)
	err = http.ListenAndServe(url, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func (s Server) ServiceID() string {
	return s.Id
}

func (s Server) ServiceName() string {
	return s.Name
}

func (s Server) GetMeta() map[string]string {
	return s.Meta
}

func (s Server) PublicAddress() string {
	return s.Address
}

func (s Server) PublicPort() uint64 {
	return s.Port
}

func (s Server) GetNamespace() string {
	return s.Namespace
}

func (s Server) GroupName() string {
	return s.Group
}
