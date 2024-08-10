package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	// selectChannelDoneAhead()
	// fibo_main()
	waitGroupTest()
}

func waitGroupTest() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	wc := new(sync.WaitGroup)
	wc.Add(2)

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		fmt.Println("Goroutine 1 done.")
		wc.Done()
	}()

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		fmt.Println("Goroutine 2 done.")
		wc.Done()
	}()

	wc.Wait()
	fmt.Println("All Goroutines done")
}

func fibo_main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		// want to get data 10 times
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, y+x
		case <-quit:
			fmt.Println("quittt")
			return
		}
	}
}

func selectChannelDoneAhead() {
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 2
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("Ch1 come first with value:", v1)
		fmt.Println("then ch2 with value:", <-ch2)
	case v2 := <-ch2:
		fmt.Println("Ch2 come first with value:", v2)
		fmt.Println("then ch1 with value:", <-ch1)
	}
}

func moreBasicForClose() {
	myChan := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			myChan <- i
		}
		close(myChan)
	}()

	for value := range myChan {
		fmt.Printf("Value: %d\n", value)
	}
}

func basicTest() {
	myChan := make(chan int)
	go receiveOnly(myChan)
	go sendOnly(myChan)
	go receiveAndSend(myChan)
	myChan <- 1
	time.Sleep(1 * time.Second)
}

func receiveAndSend(c chan int) {
	fmt.Printf("Received: %d\n", <-c)
	fmt.Printf("Sending 2...\n")
	c <- 2
}

func receiveOnly(c <-chan int) {
	fmt.Printf("Received: %d\n", <-c)
}

func sendOnly(c chan<- int) {
	c <- 2
}

func getLinkWithChannel() {
	c := make(chan string)
	links := []string{
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://google.com",
	}

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		errString := fmt.Sprint("error from get: ", err)
		c <- link
		panic(errString)
	}

	fmt.Println(link, "is up!")
	// when should we use go routine with channel?
	c <- link

}
