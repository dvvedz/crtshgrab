package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type CrtshResult struct {
	Domain string `json:"name_value"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Command takes an argument of domain")
		os.Exit(1)
	}

	da := os.Args[1]
	fmt.Println(da)
	res, err := http.Get(fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", da))

	if err != nil {
		log.Fatal("Http ERROR:", err)
	}

	fmt.Println(res.StatusCode)
	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d, and body: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	// var output []string

	var result []CrtshResult

	json.Unmarshal(body, &result)

	for _, d := range result {
		// output = append(output, d.Domain)
		fmt.Println(d.Domain)
	}

}
