package main
import ("fmt"
		"time"
		"strconv"
)
func main() {
	// start first thread
	go func1()
	// start second thread
	go func2()
	// let both threads have a chance
	// to run and finish
	time.Sleep(5 * time.Second)
}

func func1() {
	for i := 1; i <= 50; i++ {
	fmt.Println("go1: " + strconv.Itoa(i))
	}
}
func func2() {
	for i := 1; i <= 50; i++ {
	fmt.Println("go2: " + strconv.Itoa(i))
	}
}