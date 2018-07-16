package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	consul "github.com/hashicorp/consul/api"
)

func main() {
	port := flag.Int("port", 0, "Port to use on this server.")
	flag.Parse()
	if *port == 0 {
		fmt.Println("You must specify a Port for the server to run.")
		os.Exit(0)
	}
	portStr := strconv.Itoa(*port)

	config := consul.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	c, err := consul.NewClient(config)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	reg := &consul.AgentServiceRegistration{
		ID:      "example-server-" + portStr,
		Name:    "example-server",
		Address: "127.0.0.1",
		Port:    *port,
	}

	fmt.Println("Registering Service")
	err = c.Agent().ServiceRegister(reg)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Recieved Request on Port: %d\n", *port)
	})

	fmt.Println("Starting Server on Port: " + portStr)
	log.Fatal(http.ListenAndServe(":"+portStr, nil))
}
