package main

import (
        "os"
        "flag"
        "fmt"
        "log"
        "github.com/mackerelio/mackerel-client-go"
)

type MackerelAPI struct {
        cli  *mackerel.Client
}

type Allhosts struct {
        id          string
        displayname string
        status      string
        name        string
}

func NewAPI(apiKey string) *MackerelAPI {
        m := mackerel.NewClient(apiKey)
        return &MackerelAPI {
                cli: m,
        }
}

func (m *MackerelAPI) Gethosts(service string, role string) []Allhosts {
        hosts, err := m.cli.FindHosts(&mackerel.FindHostsParam{
                Service: service,
                Roles: []string{role},
        })
        if err != nil{
              log.Fatal(err)
        }

        allhosts := []Allhosts{}
        for _, h := range hosts {
                allhosts = append(allhosts, Allhosts{
                        id:               h.ID,
			displayname:      h.DisplayName,
                        status:           h.Status,
                        name:             h.Name,
                })
        }
        return allhosts
}

func (m *MackerelAPI) UpdateHosts(hostID string, status string) error {
        if status == "working"{
                resp := m.cli.UpdateHostStatus(hostID, "working")
                return resp
        } else {
                resp := m.cli.UpdateHostStatus(hostID, "standby")
                return resp
        }
}


func main() {
        var (
                typeService   string
                typeRole      string
                typeOperation string
        )

        flag.StringVar(&typeOperation, "type", "show", "Specify the operation to Mackerel-Server")
        flag.StringVar(&typeService, "service", "", "The type of monitoring service")
        flag.StringVar(&typeRole, "role", "", "The type of monitoring role")
        flag.Parse()

        apiKey := os.Getenv("MACKEREL_APIKEY")
        mkr := NewAPI(apiKey)

        hosts := mkr.Gethosts(typeService, typeRole)
        switch typeOperation {
        case "show":
                for _, h := range hosts {
                        if h.status != "working"{
                                fmt.Printf("%s, %s, %s, \x1b[31m%s\x1b[0m\n", h.id, h.displayname, h.name, h.status)
                        } else {
                                fmt.Printf("%s, %s, %s, \x1b[32m%s\x1b[0m\n", h.id, h.displayname, h.name, h.status)
                        }
                }
        case "working":
                for _, h := range hosts {
                        mkr.UpdateHosts(h.id, "working")
                }
                after_hosts := mkr.Gethosts(typeService, typeRole)
                for _, h := range after_hosts {
                        fmt.Printf("%s, %s, %s, \x1b[31m%s\x1b[0m\n", h.id, h.displayname, h.name, h.status)
                }
                fmt.Println("Status Updated to working.")
        case "standby":
                for _, h := range hosts {
                        mkr.UpdateHosts(h.id, "standby")
                }
                after_hosts := mkr.Gethosts(typeService, typeRole)
                for _, h := range after_hosts {
                        fmt.Printf("%s, %s, %s, \x1b[31m%s\x1b[0m\n", h.id, h.displayname, h.name, h.status)
                }
                fmt.Println("Status Updated to standby.")
        default:
                fmt.Println("Invalid Commands.")
        }
}