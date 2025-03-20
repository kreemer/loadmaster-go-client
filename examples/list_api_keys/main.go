package main

import (
	"encoding/json"
	"os"

	"github.com/kreemer/loadmaster-go-client/v2/api"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	api_key := os.Getenv("KEMP_API_KEY")
	ip := os.Getenv("KEMP_IP")

	client := api.NewClientWithApiKey(api_key, "https://"+ip)

	a, _ := client.ListApiKey()

	b, _ := json.Marshal(a)

	println(string(b))

}
