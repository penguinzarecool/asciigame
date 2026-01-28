package main

import (
    "fmt"
    "log"
    "time"
    "net"
    //"encoding/json"
    "github.com/gdamore/tcell/v2"
    "asciigame/network"
)

//edit with all data that I need
//type MyData struct {
    //X int `json:"x"`
    //Y int `json:"y"`
//}

func main() {

    // ------- Network Logic -------

    fmt.Printf("Host a game (1) or Join a game (2): ")
    var input int
    fmt.Scanf("%d", &input)

    var peerAddr *net.UDPAddr

    netw, _ := network.Start(9999)
    defer netw.Close()

    if input == 2 {

        fmt.Printf( "Enter a ip address: " )
        var oppIp string 
        fmt.Scanf("%s", &oppIp)

        peerAddr = &net.UDPAddr {
            IP: net.ParseIP(oppIp),
            Port: 9999,
        }

        netw.Send(network.Message{Type: "hello"}, peerAddr)

    }

    msg, addr, _ := netw.Receive()
        
    switch msg.Type {
    case "hello":
        log.Println("Got hello from", addr)
        netw.Send(network.Message{Type: "hello_ack"}, addr)
    case "hello_ack":
        log.Println("connected to peer:", addr)
    }

    // ------- Game Logic -------

    screen, err := tcell.NewScreen()
    if err != nil {
        log.Fatal(err)
    }
    if err := screen.Init(); err != nil {
        log.Fatal(err)
    }
    defer screen.Fini()

    game := NewGame(146, 52)

    inputCh := make(chan Input)
    quitCh := make(chan struct{})

    startInputLoop(screen, inputCh, quitCh)

    ticker := time.NewTicker(33 * time.Millisecond)
    defer ticker.Stop()

    for {
        select {
        case <- quitCh:
            return

        case input := <-inputCh:
            game.ApplyInput(input)

        case <-ticker.C: 
            Render(screen, game)
        }
    }
}
