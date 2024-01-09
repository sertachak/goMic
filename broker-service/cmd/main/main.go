package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
)

const webPort = "80"

var wg sync.WaitGroup

type Config struct {
}

type Route struct {
	Config
	routeName string
}

func NewConfig() *Config {
	return &Config{}
}

type doc interface {
	printDoc() string
}

type goDoc struct {
	name string
}

func (docgo *goDoc) printDoc() string {
	return "Printdoc for go;"
}

func printDocRegular(doc doc) {
	println("PrintDoc for doc %s", doc.printDoc())
}

func main() {
	documented := goDoc{
		name: "Go",
	}

	documented.printDoc()
	printDocRegular(&documented)

	log.Printf("Go os:%s \t", runtime.GOOS)
	log.Printf("Go :%s \t", runtime.GOARCH)
	log.Printf("Go :%d \t", runtime.NumCPU())
	log.Printf("Go :%d \t", runtime.NumGoroutine())

	app := Route{}
	log.Printf("/T", app)
	routes := app.routes1()

	wgTest()

	log.Printf("Starting broker service on port %s\n", webPort)

	//define https server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: routes,
	}

	//start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func wgTest() {
	wg.Add(102)
	var c = make(chan int)

	var mx sync.Mutex

	sliceA := []int{1, 2, 3, 4}
	sliceB := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}

	count := 0

	go sliceTraverse(sliceA, mx, count)
	go sliceTraverse(sliceB, mx, count)

	log.Println("Before wait")
	log.Printf("Waiting main thread %d", count)

	sharedCount := 0
	var sharedMutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			sharedMutex.Lock()
			v := sharedCount
			v++
			c <- v
			sharedCount = v
			sharedMutex.Unlock()
			wg.Done()
		}()

		go func() {
			fmt.Println(<-c)
		}()
	}

	wg.Wait()
	fmt.Println("Shared count", sharedCount)
}

func sliceTraverse(sliceA []int, mutex sync.Mutex, count int) {
	mutex.Lock()
	defer wg.Done()
	for index, value := range sliceA {
		log.Printf("index : %d value: %d", index, value)
		count++
	}
	log.Printf("Innter routines %d", count)

	mutex.Unlock()
}
