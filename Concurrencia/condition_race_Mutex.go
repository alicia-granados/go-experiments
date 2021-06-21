package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Número de CPUs", runtime.NumCPU())
	fmt.Println("Número de Gorutinas", runtime.NumGoroutine())

	contador := 0
	const gs = 100 // número de  gorutinas

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // metodo de bloqueo
			v := contador
			v++
			//time.Sleep(time.seconds* 2)
			runtime.Gosched()
			contador = v
			mu.Unlock() // método de desbloqueo
			wg.Done()
		}()
		fmt.Println("Número de Gorutinas", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
