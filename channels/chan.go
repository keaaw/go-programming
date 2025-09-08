package main

//import "fmt"
//import "time"

func main() {
	ch := make(chan int)

	for i := 0; ; i++ {
		//fmt.Print(i, " ")

		go func() {
			//time.Sleep(10 * time.Second)
			//fmt.Print("R1\n")
			ch <- 1
		}()

		go func() {
			//fmt.Print("R2\n")
			<-ch
		}()

		select {
		case <-ch:
			//fmt.Print("C1\n")
		case ch <- 2:
			//fmt.Print("C2\n")
		}
	}
}
