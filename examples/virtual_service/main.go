package main

import (
	"encoding/json"
	"net"
	"net/netip"
	"os"

	"github.com/kreemer/loadmaster-go-client/v2/api"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	api_key := os.Getenv("KEMP_API_KEY")
	ip := os.Getenv("KEMP_IP")

	client := api.NewClientWithApiKey(api_key, "https://"+ip)

	a, _ := client.ListVirtualService()
	b, _ := json.Marshal(a)
	println("[")
	println(string(b))

	vs_ip, _ := netip.AddrFromSlice(net.ParseIP("10.0.0.1"))
	req := api.AddVirtualServiceRequest{
		VirtualServiceIdentifier: &api.VirtualServiceIdentifier{
			VS:         vs_ip,
			Port:       8080,
			VSProtocol: "tcp",
		},
	}

	c, _ := client.AddVirtualService(&req)
	d, _ := json.Marshal(c)
	println(string(d))
	println("]")

}
