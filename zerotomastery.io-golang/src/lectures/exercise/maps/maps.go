//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
func printServerStatus(servers map[string]int) {
	fmt.Println("\nThere are", len(servers), "servers")
	statusList := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			statusList[Online] += 1
		case Offline:
			statusList[Offline] += 1
		case Maintenance:
			statusList[Maintenance] += 1
		case Retired:
			statusList[Retired] += 1
		default:
			panic("Unhandled server status")
		}
	}

	fmt.Println(statusList[Online], "servers online")
	fmt.Println(statusList[Offline], "servers offline")
	fmt.Println(statusList[Maintenance], "servers in maintenance")
	fmt.Println(statusList[Retired], "servers retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Store the existing slice of servers in a map
	//* Default all of the servers to `Online`
	serverStatusList := make(map[string]int)
	for _, server := range servers {
		serverStatusList[server] = Online
	}

	//* Perform the following status changes and display server info:
	//  - display server info
	printServerStatus(serverStatusList)
	//  - change `darkstar` to `Retired`
	//  - change `aiur` to `Offline`
	serverStatusList["darkstar"] = Retired
	serverStatusList["aiur"] = Offline
	//  - display server info
	printServerStatus(serverStatusList)
	//  - change all servers to `Maintenance`
	for server := range serverStatusList {
		serverStatusList[server] = Maintenance
	}
	//  - display server info
	printServerStatus(serverStatusList)
}
