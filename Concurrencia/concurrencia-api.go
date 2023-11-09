package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	//canalse declara
	//canal := make(chan int)
	//envia datos al canal
	//canal <- 15
	//RECIBE DATOS DEL CANAL
	//valor := <-canal
	//fmt.Println(valor)

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https:/outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	//fmt.Println(ch)

	//time.Sleep(1 * time.Second)

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Listo , tomo %v segundos\n,", elapsed.Seconds())

}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		//fmt.Printf("error: '%s esta caido!\n", api)
		ch <- fmt.Sprintf("error: '%s esta caido!\n", api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: ¡%s esta en funcionamiento!\n", api)
}
