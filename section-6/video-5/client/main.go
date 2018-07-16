package main

import (
	"fmt"

	consul "github.com/hashicorp/consul/api"
)

func main() {
	config := consul.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	c, err := consul.NewClient(config)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	serviceNodes, _, err := c.Catalog().Service("example-server", "", nil)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	for i, v := range serviceNodes {
		fmt.Printf("Node %d\n", i)
		fmt.Printf("ServiceID: %s\n", v.ServiceID)
		fmt.Printf("ServiceName: %s\n", v.ServiceName)
		fmt.Printf("ServiceAddress: %s\n", v.ServiceAddress)
		fmt.Printf("ServicePort: %d\n", v.ServicePort)
	}

}
