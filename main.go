package main

import (
	"Codigo/semaforo"
	"math/rand"
	"sync"
)

func auto_sur(puente1 *semaforo.Sem, puente2 *semaforo.Sem, wg *sync.WaitGroup) {
	numRandom := rand.Intn(10000)
	puente2.Wait()
	println("-- Sur a norte -- ", "Cruzando puente2 Auto nº", numRandom)
	puente2.Signal()
	puente1.Wait()
	println("-- Sur a norte -- ", "Cruzando puente1 Auto nº", numRandom)
	puente1.Signal()
	wg.Done()
}

func auto_norte(puente1 *semaforo.Sem, puente2 *semaforo.Sem, wg *sync.WaitGroup) {
	numRandom := rand.Intn(10000)
	puente1.Wait()
	println("-- Norte a Sur -- ", "Cruzando puente1 Auto nº", numRandom)
	puente1.Signal()
	puente2.Wait()
	println("-- Norte a Sur -- ", "Cruzando puente2 Auto nº", numRandom)
	puente2.Signal()
	wg.Done()
}

func main() {

	puente1 := semaforo.InitSemaphore(1)
	puente2 := semaforo.InitSemaphore(1)
	trafico := 15
	wg := &sync.WaitGroup{}
	wg.Add(trafico)
	for i := 0; i < trafico; i++ {
		if rand.Intn(101) >= 50 {
			go auto_sur(&puente1, &puente2, wg)
		} else {
			go auto_norte(&puente1, &puente2, wg)
		}
	}
	wg.Wait()
}
