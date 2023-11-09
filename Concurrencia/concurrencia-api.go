package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https:/outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		go checkAPI(api)
	}

	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Listo , tomo %v segundos\n", elapsed.Seconds())

}

func checkAPI(api string) {
	if _, err := http.Get(api); err != nil {
		fmt.Printf("error: '%s esta caido!\n", api)
		return
	}
	fmt.Printf("SUCCESS: ¡%s esta en funcionamiento!\n", api)
}
