package network

import (
    "encoding/json"
    "net"
)

type Network struct {
    conn *net.UDPConn
}

func Start(port int) (*Network, error) {
    addr := &net.UDPAddr{
        IP:   net.IPv4zero,
        Port: port,
    }

    conn, err := net.ListenUDP("udp4", addr)
    if err != nil {
        return nil, err
    }

    return &Network{conn: conn}, nil
}

func (n *Network) Send(msg Message, addr *net.UDPAddr) error {
    data, _ := json.Marshal(msg)
    _, err := n.conn.WriteToUDP(data, addr)
    return err
}

func (n *Network) Receive() (Message, *net.UDPAddr, error) {
    buf := make([]byte, 1024)

    nBytes, addr, err := n.conn.ReadFromUDP(buf)
    if err != nil {
        return Message{}, nil, err
    }

    var msg Message
    err = json.Unmarshal(buf[:nBytes], &msg)
    return msg, addr, err
}

func (n *Network) Close() error {
    return n.conn.Close()
}
