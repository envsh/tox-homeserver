package main

import (
	"fmt"

	"gopkg.in/go-playground/stats.v1"
)

// goroutine+1
func StartStatsServer() { go RunStatsServer() }
func RunStatsServer() {
	config := &stats.ServerConfig{
		Domain: "",
		Port:   3008,
		Debug:  false,
	}

	server, err := stats.NewServer(config)
	if err != nil {
		panic(err)
	}

	for stat := range server.Run() {

		// calculate CPU times
		// totalCPUTimes := stat.CalculateTotalCPUTimes()
		// perCoreCPUTimes := stat.CalculateCPUTimes()

		// Do whatever you want with the data
		// * Save to database
		// * Stream elsewhere
		// * Print to console
		//

		fmt.Println(stat)
	}
}

// goroutine+1
func NewSimpleStatsClient() *stats.ClientStats {
	serverConfig := &stats.ServerConfig{
		Domain: "remoteserver",
		Port:   3008,
		Debug:  false,
	}

	clientConfig := &stats.ClientConfig{
		Domain:           "",
		Port:             3009,
		PollInterval:     1000,
		Debug:            false,
		LogHostInfo:      true,
		LogCPUInfo:       true,
		LogTotalCPUTimes: true,
		LogPerCPUTimes:   true,
		LogMemory:        true,
		LogGoMemory:      true,
	}

	client, err := stats.NewClient(clientConfig, serverConfig)
	if err != nil {
		panic(err)
	}

	go client.Run()
	return client
}
