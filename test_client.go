package main

import (
    "log"
    "runtime"
    "fmt"

    "github.com/zeromq/goczmq"
    "time"
    "sync"
)

func main() {
    // Create a dealer socket and connect it to the router.
    dealer := goczmq.NewDealerChanneler("tcp://127.0.0.1:5555")
    defer dealer.Destroy()

    log.Println("dealer created and connected")

    wg := &sync.WaitGroup{}

    max := runtime.NumCPU()
    //max = 2
    for i := 0; i < max; i++ {
        go func(wg *sync.WaitGroup) {
            for {
                // Receive the reply.
                reply := <-dealer.RecvChan

                t, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", string(reply[0]))
                log.Printf("dealer received '%s' %s", string(reply[0]), time.Now().Sub(t))

                wg.Done()
            }
        }(wg)
    }

    for i:= 0; i<10; i++ {
        wg.Add(1)
        go func() {
            //tStart := time.Now()
            // Send a 'Hello' message from the dealer to the router.
            // Here we send it as a frame ([]byte), with a FlagNone
            // flag to indicate there are no more frames following.
            dealer.SendChan <- [][]byte{[]byte(fmt.Sprintf("%s", time.Now()))}
            log.Println("dealer sent 'Hello'")
        }()

    }

    wg.Wait()

}
