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
					{
						Name:  "show",
						Usage: "show a real server",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
							&cli.StringFlag{Name: "real-server", Aliases: []string{"rs"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							response, err := client.ShowRealServer(cmd.String("virtual-service"), cmd.String("real-server"))
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
			{
				Name:    "rule-assignment",
				Aliases: []string{"ra"},
				Usage:   "Manage rule assignments",
				Commands: []*cli.Command{
					{
						Name:    "real-server",
						Aliases: []string{"rs"},
						Usage:   "Manage rule assignments for real servers",
						Commands: []*cli.Command{
							{
								Name:  "add",
								Usage: "Add a rule to a real server",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
									&cli.StringFlag{Name: "real-server", Aliases: []string{"rs"}},
									&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
								},
								Action: func(c context.Context, cmd *cli.Command) error {
									virtual_service := cmd.String("virtual-service")
									real_server := cmd.String("real-server")
									name := cmd.String("name")

									if virtual_service == "" || real_server == "" || name == "" {
										return fmt.Errorf("missing virtual service, real server or rule name")
									}
									response, err := client.AddRealServerRule(virtual_service, real_server, name)
									if err != nil {
										return err
									}
									fmt.Println(prettyPrint(response))
									return nil
								},
							},
							{
								Name:  "del",
								Usage: "Delete a rule from a real server",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
									&cli.StringFlag{Name: "real-server", Aliases: []string{"rs"}},
									&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
								},
								Action: func(c context.Context, cmd *cli.Command) error {
									virtual_service := cmd.String("virtual-service")
									real_server := cmd.String("real-server")
									name := cmd.String("name")

									if virtual_service == "" || real_server == "" || name == "" {
										return fmt.Errorf("missing virtual service, real server or rule name")
									}
									response, err := client.DeleteRealServerRule(virtual_service, real_server, name)
									if err != nil {
										return err
									}
									fmt.Println(prettyPrint(response))
									return nil
								},
							},
							{
								Name:  "show",
								Usage: "Show a rule assignment from a real server",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
									&cli.StringFlag{Name: "real-server", Aliases: []string{"rs"}},
									&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
								},
								Action: func(c context.Context, cmd *cli.Command) error {
									virtual_service := cmd.String("virtual-service")
									real_server := cmd.String("real-server")
									name := cmd.String("name")

									if virtual_service == "" || real_server == "" || name == "" {
										return fmt.Errorf("missing virtual service, real server or rule name")
									}
									response, err := client.ShowRealServerRule(virtual_service, real_server, name)
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
						Usage:   "Manage rule assignments for sub virtual services",
						Commands: []*cli.Command{
							{
								Name:  "add",
								Usage: "Add a rule to a sub virtual service",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
									&cli.StringFlag{Name: "sub-virtual-service", Aliases: []string{"subvs"}},
									&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
								},
								Action: func(c context.Context, cmd *cli.Command) error {
									virtual_service := cmd.String("virtual-service")
									sub_virtual_service := cmd.String("sub-virtual-service")
									name := cmd.String("name")

									if virtual_service == "" || sub_virtual_service == "" || name == "" {
										return fmt.Errorf("missing virtual service, sub virtual service or rule name")
									}
									response, err := client.AddSubVirtualServiceRule(virtual_service, sub_virtual_service, name)
									if err != nil {
										return err
									}
									fmt.Println(prettyPrint(response))
									return nil
								},
							},
							{
								Name:  "del",
								Usage: "Delete a rule from a sub virtual service",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
									&cli.StringFlag{Name: "sub-virtual-service", Aliases: []string{"subvs"}},
									&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
								},
								Action: func(c context.Context, cmd *cli.Command) error {
									virtual_service := cmd.String("virtual-service")
									sub_virtual_service := cmd.String("sub-virtual-service")
									name := cmd.String("name")

									if virtual_service == "" || sub_virtual_service == "" || name == "" {
										return fmt.Errorf("missing virtual service, sub virtual service or rule name")
									}
									response, err := client.DeleteSubVirtualServiceRule(virtual_service, sub_virtual_service, name)
									if err != nil {
										return err
									}
									fmt.Println(prettyPrint(response))
									return nil
								},
							},
							{
								Name:  "show",
								Usage: "Show a rule assignment from a sub virtual service",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
									&cli.StringFlag{Name: "sub-virtual-service", Aliases: []string{"subvs"}},
									&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
								},
								Action: func(c context.Context, cmd *cli.Command) error {
									virtual_service := cmd.String("virtual-service")
									sub_virtual_service := cmd.String("sub-virtual-service")
									name := cmd.String("name")

									if virtual_service == "" || sub_virtual_service == "" || name == "" {
										return fmt.Errorf("missing virtual service, sub virtual service or rule name")
									}
									response, err := client.ShowSubVirtualServiceRule(virtual_service, sub_virtual_service, name)
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
						Name:    "virtual-service",
						Aliases: []string{"vs"},
						Usage:   "Manage rule assignments for virtual services",
						Commands: []*cli.Command{
							{
								Name:    "pre-rule",
								Aliases: []string{"pre"},
								Usage:   "Manage pre rule assignments for virtual services",
								Commands: []*cli.Command{
									{
										Name:  "add",
										Usage: "Add a prerule to a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.AddVirtualServicePreRule(virtual_service, name)
											if err != nil {
												return err
											}
											fmt.Println(prettyPrint(response))
											return nil
										},
									},
									{
										Name:  "show",
										Usage: "Show a prerule from a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.ShowVirtualServicePreRule(virtual_service, name)
											if err != nil {
												return err
											}
											fmt.Println(prettyPrint(response))
											return nil
										},
									},
									{
										Name:  "del",
										Usage: "Delete a prerule from a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.DeleteVirtualServicePreRule(virtual_service, name)
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
								Name:    "req-rule",
								Aliases: []string{"rq"},
								Usage:   "Manage request rule assignments for virtual services",
								Commands: []*cli.Command{
									{
										Name:  "add",
										Usage: "Add a reqrule to a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.AddVirtualServiceRequestRule(virtual_service, name)
											if err != nil {
												return err
											}
											fmt.Println(prettyPrint(response))
											return nil
										},
									},
									{
										Name:  "show",
										Usage: "Show a reqrule from a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.ShowVirtualServiceRequestRule(virtual_service, name)
											if err != nil {
												return err
											}
											fmt.Println(prettyPrint(response))
											return nil
										},
									},
									{
										Name:  "del",
										Usage: "Delete a reqrule from a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.DeleteVirtualServiceRequestRule(virtual_service, name)
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
								Name:    "res-rule",
								Aliases: []string{"re"},
								Usage:   "Manage response rule assignments for virtual services",
								Commands: []*cli.Command{
									{
										Name:  "add",
										Usage: "Add a response rule to a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.AddVirtualServiceResponseRule(virtual_service, name)
											if err != nil {
												return err
											}
											fmt.Println(prettyPrint(response))
											return nil
										},
									},
									{
										Name:  "show",
										Usage: "Show a response rule from a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.ShowVirtualServiceResponseRule(virtual_service, name)
											if err != nil {
												return err
											}
											fmt.Println(prettyPrint(response))
											return nil
										},
									},
									{
										Name:  "del",
										Usage: "Delete a response rule from a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.DeleteVirtualServiceResponseRule(virtual_service, name)
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
								Name:    "res-body-rule",
								Aliases: []string{"rb"},
								Usage:   "Manage response body rule assignments for virtual services",
								Commands: []*cli.Command{

									{
										Name:  "add",
										Usage: "Add a response body rule to a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.AddVirtualServiceResponseBodyRule(virtual_service, name)
											if err != nil {
												return err
											}
											fmt.Println(prettyPrint(response))
											return nil
										},
									},
									{
										Name:  "show",
										Usage: "Show a response body rule from a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.ShowVirtualServiceResponseBodyRule(virtual_service, name)
											if err != nil {
												return err
											}
											fmt.Println(prettyPrint(response))
											return nil
										},
									},
									{
										Name:  "del",
										Usage: "Delete a response body rule from a virtual service",
										Flags: []cli.Flag{
											&cli.StringFlag{Name: "virtual-service", Aliases: []string{"vs"}},
											&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
										},
										Action: func(c context.Context, cmd *cli.Command) error {
											virtual_service := cmd.String("virtual-service")
											name := cmd.String("name")

											if virtual_service == "" || name == "" {
												return fmt.Errorf("missing virtual service or rule name")
											}
											response, err := client.DeleteVirtualServiceResponseBodyRule(virtual_service, name)
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
					},
				},
			},
			{
				Name:    "certificate",
				Aliases: []string{"c"},
				Usage:   "Manage certificates",
				Commands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add certificate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data", Aliases: []string{"d"}},
							&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
							&cli.StringFlag{Name: "password", Aliases: []string{"p"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							data := cmd.String("data")
							name := cmd.String("name")

							var password *string
							if cmd.IsSet("password") {
								p := cmd.String("password")
								password = &p
							}

							if data == "" || name == "" {
								return fmt.Errorf("missing certificate data or name")
							}

							response, err := client.AddCertificate(name, password, data)
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "del",
						Usage: "delete certificate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							name := cmd.String("name")
							if name == "" {
								return fmt.Errorf("missing certificate name")
							}

							response, err := client.DeleteCertificate(name)
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "show",
						Usage: "show a certificate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							name := cmd.String("name")
							if name == "" {
								return fmt.Errorf("missing certificate name")
							}

							response, err := client.ShowCertificate(name)
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
				Name:    "intermediate-certificate",
				Aliases: []string{"ic"},
				Usage:   "Manage intermediate certificates",
				Commands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add intermediate certificate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data", Aliases: []string{"d"}},
							&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							data := cmd.String("data")
							name := cmd.String("name")

							if data == "" || name == "" {
								return fmt.Errorf("missing certificate data or name")
							}

							response, err := client.AddIntermediateCertificate(name, data)
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "del",
						Usage: "delete intermediate certificate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							name := cmd.String("name")
							if name == "" {
								return fmt.Errorf("missing certificate name")
							}

							response, err := client.DeleteIntermediateCertificate(name)
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "show",
						Usage: "show a intermediate certificate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							name := cmd.String("name")
							if name == "" {
								return fmt.Errorf("missing certificate name")
							}

							response, err := client.ShowIntermediateCertificate(name)
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
				Name:    "administration",
				Aliases: []string{"a"},
				Usage:   "Administration tasks",
				Commands: []*cli.Command{
					{
						Name:  "backup",
						Usage: "Backup the current configuration",
						Action: func(c context.Context, cmd *cli.Command) error {
							response, err := client.Backup()
							if err != nil {
								return err
							}
							fmt.Println(prettyPrint(response))
							return nil
						},
					},
					{
						Name:  "restore",
						Usage: "Restore configuration",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data", Aliases: []string{"d"}},
							&cli.StringFlag{Name: "type", Aliases: []string{"t"}},
						},
						Action: func(c context.Context, cmd *cli.Command) error {
							data := cmd.String("data")
							backup_type := cmd.String("type")
							if data == "" || backup_type == "" {
								return fmt.Errorf("missing data or type")
							}

							response, err := client.Restore(data, backup_type)
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
