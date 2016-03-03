// also need TCPChan server

package tcpchan

import (
  "net"
)

type TCPChan struct {
  In  chan []byte
  Out chan []byte
  Err chan error
  Conn net.Conn
}

func (ch *TCPChan) Dial(addr string) error {

  ch.In = make(chan []byte, 100)
  ch.Out = make(chan []byte, 100)
  ch.Err = make(chan error, 100)
  conn, err := net.Dial("tcp", addr)
  if err != nil {
    return err
  }
  ch.Conn = conn
  go ch.checkOutgoing()
  go ch.checkIncoming()
  return nil
}

func (ch *TCPChan) checkOutgoing() {
  for {
    msg := <-ch.Out
    ch.Conn.Write(msg)
  }
}

func (ch *TCPChan) checkIncoming() {

  buffer := make([]byte, 1000)
  for {
    n, err := ch.Conn.Read(buffer)
    if err != nil {
      ch.Err <- err
      continue
    }
    ch.In <- buffer[:n]
  }
}
