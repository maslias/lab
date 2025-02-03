package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type Message struct {
	from    net.Addr
	payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Message
	peerch     chan net.Conn
	peers      map[net.Conn]bool
}

func main() {
	s := NewServer(":3000")
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	sendFile(1024 * 8)
	// }()
	// go func() {
	// 	for msg := range s.msgch {
	// 		fmt.Printf("msg from: %s size: %d msg: %s \n", msg.from.String(), len(msg.payload), string(msg.payload))
	// 	}
	// }()
	log.Fatal(s.Run())
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message),
		peerch:     make(chan net.Conn),
		peers:      make(map[net.Conn]bool),
	}
}

func (s *Server) Run() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	go s.stateConn()
	go s.acceptConn()

	<-s.quitch

	return nil
}

func (s *Server) stateConn() {
	for {
		select {
		case conn := <-s.peerch:
			s.peers[conn] = true
			fmt.Printf("new peer connected %s - total peers: %d \n", conn.RemoteAddr().String(), len(s.peers))
			go s.readConn(conn)
		case msg := <-s.msgch:
			if strings.TrimSpace(string(msg.payload)) == "killer" {
				s.writeAll("bye bye")
				for peer := range s.peers {
					peer.Close()
				}
				s.quitch <- struct{}{}
			}
			fmt.Printf("msg from: %s size: %d msg: %s \n", msg.from.String(), len(msg.payload), string(msg.payload))
			s.writeAllPeersAcceptConn(msg)
		}
	}
}

func (s *Server) acceptConn() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		s.peerch <- conn
	}
}

func (s *Server) readConn(conn net.Conn) {
	defer conn.Close()
	// buf := new(bytes.Buffer)
	buf := make([]byte, 2048)

	for {
		// var size int64
		// binary.Read(conn, binary.LittleEndian, &size)
		// _, err := io.CopyN(buf, conn, size)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}

		s.msgch <- Message{
			from: conn.RemoteAddr(),
			// payload: buf.Bytes(),
			payload: buf[:n],
		}
	}
}

func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	// binary.Write(conn, binary.LittleEndian, int64(size))
	// n, err := io.CopyN(conn, bytes.NewBuffer(file), int64(size))
	n, err := conn.Write([]byte("hello there from simulator"))
	if err != nil {
		return err
	}
	fmt.Printf("SIMULATE CLIENT - written file: %d \n", n)

	return nil
}

func (s *Server) writeAllPeersAcceptConn(msg Message) {
	for p := range s.peers {
		if p.RemoteAddr() == msg.from {
			continue
		}
		_, err := p.Write([]byte(msg.from.String() + "wrote: " + string(msg.payload) + "\n"))
		if err != nil {
			fmt.Println("write error:", err)
			continue
		}
	}
}

func (s *Server) writeAll(str string) {
	for p := range s.peers {
		_, err := p.Write([]byte(str))
		if err != nil {
			fmt.Println("write error:", err)
			continue
		}
	}
}
