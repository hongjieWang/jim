package client

import (
	"fmt"
	"github.com/jim/demo/naming/nacos"
	"io/ioutil"
	"net/http"
)

const HTTP = "http://"

type Client struct {
	Name      string
	Namespace string
	Group     string
}

func (c Client) GetServerUrl() string {
	newNaming, err := nacos.NewNaming("110.40.141.168", 8848, c.Namespace)
	if err != nil {
		panic(err)
	}
	name, _ := newNaming.FindByServerName(c.Name, c.Group)
	address := name.ServiceID()
	port := name.PublicPort()
	sprintf := fmt.Sprintf("%s%s:%d", HTTP, address, port)
	return sprintf
}

func (c Client) SendGet() {
	resp, err := http.Get(c.GetServerUrl())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func ClientRun() {
	c := Client{
		Namespace: "test",
		Name:      "myServer",
		Group:     "test",
	}
	c.SendGet()
}
