package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kreemer/loadmaster-go-client/v2/api"
	"github.com/urfave/cli/v2"
)

func main() {

	api_key := os.Getenv("KEMP_API_KEY")
	ip := os.Getenv("KEMP_IP")

	client := api.NewClientWithApiKey(api_key, "https://"+ip)

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "virtual-service",
				Aliases: []string{"vs"},
				Usage:   "Manage virtual services",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list virtual service",
						Action: func(c *cli.Context) error {
							response, err := client.ListVirtualService()
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "show",
						Usage: "show a virtual service",
						Action: func(c *cli.Context) error {
							vs_identifier := c.Args().First()
							if vs_identifier == "" {
								return fmt.Errorf("missing virtual service identifier")
							}
							id, _ := strconv.Atoi(vs_identifier)
							response, err := client.ShowVirtualService(id)
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "add",
						Usage: "add a virtual service",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address", Aliases: []string{"a"}, Required: true},
							&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Required: true},
							&cli.StringFlag{Name: "protocol", Aliases: []string{"t"}},
						},
						Action: func(c *cli.Context) error {
							response, err := client.AddVirtualService(c.String("address"), c.Int("port"), c.String("protocol"), api.VirtualServiceParameters{})
							if err != nil {
								return err
							}

							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "del",
						Usage: "delete a virtual service",
						Action: func(c *cli.Context) error {
							vs_identifier := c.Args().First()
							if vs_identifier == "" {
								return fmt.Errorf("missing virtual service identifier")
							}
							id, _ := strconv.Atoi(vs_identifier)
							response, err := client.DeleteVirtualService(id)
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
				},
			},
			{
				Name:    "sub-virtual-service",
				Aliases: []string{"subvs"},
				Usage:   "Manage sub virtual services",
				Subcommands: []*cli.Command{
					{
						Name:  "show",
						Usage: "show a sub virtual service",
						Action: func(c *cli.Context) error {
							vs_identifier := c.Args().First()
							if vs_identifier == "" {
								return fmt.Errorf("missing sub virtual service identifier")
							}
							id, _ := strconv.Atoi(vs_identifier)
							response, err := client.ShowSubVirtualService(id)
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "add",
						Usage: "add a sub virtual service",
						Action: func(c *cli.Context) error {
							vs_identifier := c.Args().First()
							if vs_identifier == "" {
								return fmt.Errorf("missing sub virtual service identifier")
							}
							id, _ := strconv.Atoi(vs_identifier)

							response, err := client.AddSubVirtualService(id, api.VirtualServiceParameters{})
							if err != nil {
								return err
							}

							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "del",
						Usage: "delete a sub virtual service",
						Action: func(c *cli.Context) error {
							vs_identifier := c.Args().First()
							if vs_identifier == "" {
								return fmt.Errorf("missing sub virtual service identifier")
							}
							id, _ := strconv.Atoi(vs_identifier)
							response, err := client.DeleteSubVirtualService(id)
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func prettyPrint(i any) string {
	s, _ := json.Marshal(i)
	return string(s)
}
