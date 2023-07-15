package main

import "university-api/cmd"

func main() {
	server := cmd.ApiServer{}

	server.Start()
}
