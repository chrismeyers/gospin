package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func spin() {
	i := 0

	for {
		resp, err := http.Get(fmt.Sprintf("http://localhost:8080/step/%d", i))

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		var result = Step{}
		json.NewDecoder(resp.Body).Decode(&result)

		fmt.Printf("%s\r", result.Glyph)

		i++
		if result.Reset {
			i = 0
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go run() // Start API server
	spin()   // SPIN FOREVER!
}
