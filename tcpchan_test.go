package tcpchan

import (
  "testing"
  "net"

  "github.com/filwisher/tcptest"
)

var messages = []string{
  "hello",
  "本語本本語語",
  "\t\n\aHello",
}

func TestStartServer(t *testing.T) {
  buffer := make([]byte, 1024)
  tcptest.NewServer(":8080", func (conn net.Conn) {
    for {
      n, err := conn.Read(buffer)
      if err != nil {
        panic("Test server failed")
      }
      conn.Write(buffer[:n])
    }
  })
}

func TestConn(t *testing.T) {

  var c TCPChan
  err := c.Dial(":8080")
  if err != nil {
    t.Errorf("Failed with err %s", err.Error())
  }

  for _, msg := range(messages) { 
    c.Out <- []byte(msg)
    rcv := <-c.In

    if msg != string(rcv) {
      t.Errorf("Sent %s, got back %s", msg, string(rcv))
    }
  }
}
