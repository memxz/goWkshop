package main
import ("fmt"
		"sync"
		"strconv"
)
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	chan_ma := make(chan int)
	chan_ab := make(chan int)
	go go1(chan_ma, chan_ab, &wg)
	var _ = <-chan_ma
	go go2(chan_ab, &wg)
	wg.Wait()
}

func go1(ma chan int, ab chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ma <- 1
	for i := 1; i <= 50; i++ {
		fmt.Println("go1: " + strconv.Itoa(i))
		ab <- 1
		var _ = <-ab
	}
}
func go2(ab chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		var _ = <-ab
		fmt.Println("go2: " + strconv.Itoa(i))
		ab <- 1
	}
}