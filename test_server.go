package main

import (
    "log"

    "github.com/zeromq/goczmq"
    "runtime"
    "sync"
)

func main() {
    // Create a router socket and bind it to port 5555.
    router := goczmq.NewRouterChanneler("tcp://*:5555")
    defer router.Destroy()

    log.Println("router created and bound")

    wg := &sync.WaitGroup{}

    max := runtime.NumCPU()
    //max = 2

    for i:= 0; i< max; i++ {
        wg.Add(1)
        go func(wg *sync.WaitGroup){
            for {
                // Receve the message. Here we call RecvMessage, which
                // will return the message as a slice of frames ([][]byte).
                // Since this is a router socket that support async
                // request / reply, the first frame of the message will
                // be the routing frame.
                request := <-router.RecvChan
                /*request, err := router.RecvMessage()
                if err != nil {
                    log.Fatal(err)
                }*/

                log.Printf("router received '%v' from '%v'", request[1], request[0])

                router.SendChan <- [][]byte{request[0], request[1]}

                /*
                // Send a reply. First we send the routing frame, which
                // lets the dealer know which client to send the message.
                // The FlagMore flag tells the router there will be more
                // frames in this message.
                err = router.SendFrame(request[0], goczmq.FlagMore)
                if err != nil {
                    log.Fatal(err)
                }

                log.Printf("router sent 'World'")

                // Next send the reply. The FlagNone flag tells the router
                // that this is the last frame of the message.
                err = router.SendFrame(request[1], goczmq.FlagNone)
                if err != nil {
                    log.Fatal(err)
                }*/
            }
            wg.Done()
        }(wg)
    }

    wg.Wait()
}
