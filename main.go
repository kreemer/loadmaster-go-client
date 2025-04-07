package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/kreemer/loadmaster-go-client/api"
	"github.com/urfave/cli/v3"
)

func main() {
	api_key := os.Getenv("LOADMASTER_API_KEY")
	ip := os.Getenv("LOADMASTER_IP")

	slog.SetLogLoggerLevel(slog.LevelError)
	var count int
	client := api.NewClientWithApiKey("https://"+ip, api_key)
	app := &cli.Command{
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Show verbose output",
				Config: cli.BoolConfig{
					Count: &count,
				},
				Action: func(ctx context.Context, cmd *cli.Command, _ bool) error {
					if count == 1 {
						slog.SetLogLoggerLevel(slog.LevelWarn)
					} else if count == 2 {
						slog.SetLogLoggerLevel(slog.LevelInfo)
					} else if count == 3 {
						slog.SetLogLoggerLevel(slog.LevelDebug)
					}

					slog.Info("Verbose mode enabled", "Level", count)

					client.SetDebugLevel(uint8(count))

					return nil
				},
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "virtual-service",
				Aliases: []string{"vs"},
				Usage:   "Manage virtual services",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list virtual service",
						Action: func(c context.Context, cmd *cli.Command) error {
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
						Action: func(c context.Context, cmd *cli.Command) error {
							vs_identifier := cmd.Args().First()
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
							&cli.StringFlag{Name: "port", Aliases: []string{"p"}, Required: true},
							&cli.StringFlag{Name: "protocol", Aliases: []string{"t"}},
							&cli.StringFlag{Name: "data", Aliases: []string{"d"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							bytes := []byte(cmd.String("data"))
							params := api.VirtualServiceParameters{}
							json.Unmarshal(bytes, &params)

							response, err := client.AddVirtualService(cmd.String("address"), cmd.String("port"), cmd.String("protocol"), params)

							if err != nil {
								return err
							}

							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "mod",
						Usage: "Modify a virtual service",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data", Aliases: []string{"d"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							vs_identifier := cmd.Args().First()
							if vs_identifier == "" {
								return fmt.Errorf("missing virtual service identifier")
							}
							id, _ := strconv.Atoi(vs_identifier)

							bytes := []byte(cmd.String("data"))
							params := api.VirtualServiceParameters{}
							json.Unmarshal(bytes, &params)

							response, err := client.ModifyVirtualService(id, params)

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
						Action: func(c context.Context, cmd *cli.Command) error {
							vs_identifier := cmd.Args().First()
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
				Commands: []*cli.Command{
					{
						Name:  "show",
						Usage: "show a sub virtual service",
						Action: func(c context.Context, cmd *cli.Command) error {
							vs_identifier := cmd.Args().First()
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data", Aliases: []string{"d"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							vs_identifier := cmd.Args().First()
							if vs_identifier == "" {
								return fmt.Errorf("missing sub virtual service identifier")
							}
							id, _ := strconv.Atoi(vs_identifier)

							bytes := []byte(cmd.String("data"))
							params := api.VirtualServiceParameters{}
							json.Unmarshal(bytes, &params)

							response, err := client.AddSubVirtualService(id, params)
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
						Action: func(c context.Context, cmd *cli.Command) error {
							vs_identifier := cmd.Args().First()
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
					{
						Name:  "mod",
						Usage: "Modify a sub virtual service",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data", Aliases: []string{"d"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							vs_identifier := cmd.Args().First()
							if vs_identifier == "" {
								return fmt.Errorf("missing sub virtual service identifier")
							}
							id, _ := strconv.Atoi(vs_identifier)

							bytes := []byte(cmd.String("data"))
							params := api.VirtualServiceParameters{}
							json.Unmarshal(bytes, &params)

							response, err := client.ModifySubVirtualService(id, params)

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
				Name:    "real-server",
				Aliases: []string{"rs"},
				Usage:   "Manage realserver",
				Commands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add real server",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data", Aliases: []string{"d"}},
							&cli.StringFlag{Name: "vs", Aliases: []string{"v"}},
							&cli.StringFlag{Name: "address", Aliases: []string{"a"}},
							&cli.StringFlag{Name: "port", Aliases: []string{"p"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {

							bytes := []byte(cmd.String("data"))
							params := api.RealServerParameters{}
							json.Unmarshal(bytes, &params)

							response, err := client.AddRealServer(cmd.String("vs"), cmd.String("address"), cmd.String("port"), params)
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "del",
						Usage: "delete a real server",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "vs", Aliases: []string{"v"}},
							&cli.StringFlag{Name: "rs", Aliases: []string{"r"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							response, err := client.DeleteRealServer(cmd.String("vs"), cmd.String("rs"))
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "mod",
						Usage: "modifies a real server",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "vs", Aliases: []string{"v"}},
							&cli.StringFlag{Name: "rs", Aliases: []string{"r"}},
							&cli.StringFlag{Name: "data", Aliases: []string{"d"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							bytes := []byte(cmd.String("data"))
							params := api.RealServerParameters{}
							json.Unmarshal(bytes, &params)
							response, err := client.ModifyRealServer(cmd.String("vs"), cmd.String("rs"), params)
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
				Name:    "rule",
				Aliases: []string{"r"},
				Usage:   "Manage rules",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list rule",
						Action: func(c context.Context, cmd *cli.Command) error {
							response, err := client.ListRule()
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "add",
						Usage: "add rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data", Aliases: []string{"d"}},
							&cli.StringFlag{Name: "type", Aliases: []string{"t"}},
							&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {

							bytes := []byte(cmd.String("data"))
							params := api.GeneralRule{}
							json.Unmarshal(bytes, &params)

							response, err := client.AddRule(cmd.String("type"), cmd.String("name"), params)
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "del",
						Usage: "delete a rule",
						Action: func(c context.Context, cmd *cli.Command) error {
							name := cmd.Args().First()
							if name == "" {
								return fmt.Errorf("missing rule name")
							}
							response, err := client.DeleteRule(name)
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
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}

func prettyPrint(i any) string {
	s, _ := json.Marshal(i)
	return string(s)
}
