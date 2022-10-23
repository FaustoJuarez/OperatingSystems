package semaforo

type empty struct{}

//type Sem struct {
//	ch chan empty
//}
//
//func InitSemaphore(value int) Sem {
//	var sem Sem
//	sem.ch = make(chan empty, value)
//	return sem
//}
//
//func (sem *Sem) Wait() {
//	sem.ch <- empty{}
//}
//
//func (sem *Sem) Signal() {
//	<-sem.ch
//}

//Opcion 2
type Sem struct {
	value chan int
	queue chan empty
}

func InitSemaphore(value int) Sem {
	var s Sem
	s.value = make(chan int, 1)
	s.queue = make(chan empty)
	s.value <- value
	return s
}
func (s Sem) Wait() {
	v := <-s.value
	v--
	s.value <- v
	if v < 0 {
		<-s.queue
	}
}
func (s Sem) Signal() {
	v := <-s.value
	v++
	s.value <- v
	if v <= 0 {
		s.queue <- empty{}
	}
}
