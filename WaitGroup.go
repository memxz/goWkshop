package main
import ("fmt"
		"sync"
		"strconv"
)
func main() {
	wg := sync.WaitGroup{}
	// two goroutines to wait on
	wg.Add(2)
	go func1(&wg)
	go func2(&wg)
	// wait for completion
	wg.Wait()
}

func func1(wg *sync.WaitGroup) {
	// invoked at end of function
	defer wg.Done()
	for i := 1; i <= 50; i++ {
	fmt.Println("go1: " + strconv.Itoa(i))
}
}
func func2(wg *sync.WaitGroup) {
	// invoked at end of function
	defer wg.Done()
	for i := 1; i <= 50; i++ {
	fmt.Println("go2: " + strconv.Itoa(i))
	}
}