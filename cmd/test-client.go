package main

import (
	"fmt"
	"os"
)

func main() {
	client, err := http_client.NewClient("https://api.agify.io")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := client.GetAge("Sig")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}
