package utils

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type Server struct {
	addr  string
	port  string
	store *InMemoryStore
}

func NewServer(addr, port string) *Server {
	return &Server{
		addr:  addr,
		port:  port,
		store: NewInMemoryStore(),
	}
}

func (server *Server) removeExpiredKeys() {
	t := time.NewTicker(10 * time.Second)
	defer t.Stop()
	for range t.C {
		log.Println("Cleaning up expired keys")
		server.store.delExpired()
	}
}

func (server *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	write := func(msg string) {
		_, _ = conn.Write([]byte(msg))
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		switch cmd := ParseCommand(line); cmd {
		case GET:
			key, ok := ParseGetOrDel(line)
			if !ok {
				write("Invalid get command\n")
				continue
			}
			if val, ok := server.store.get(key); ok {
				write(fmt.Sprintf("%s\n", val))
			} else {
				write(fmt.Sprintf("key %s not found\n", key))
			}
		case PUT:
			key, val, ok := ParsePut(line)
			if !ok {
				write("Invalid put command\n")
				continue
			}
			server.store.put(key, val)
			write("OK\n")
		case DEL:
			key, ok := ParseGetOrDel(line)
			if !ok {
				write("Invalid del command\n")
				continue
			}
			server.store.del(key)
			write("OK\n")
		case LIST:
			all := server.store.list()
			var response strings.Builder
			for k, v := range all {
				response.WriteString(fmt.Sprintf("%s=%s\n", k, v))
			}
			write(response.String())
		case UNKNOWN:
			write("Unknown command\n")
		}
	}
}

func (server *Server) Start() {
	fullAddr := fmt.Sprintf("%s:%s", server.addr, server.port)
	listener, err := net.Listen("tcp", fullAddr)
	if err != nil {
		log.Fatal("[err]: failed to start server listener")
	}
	defer listener.Close()

	log.Printf("Server listening on %s", fullAddr)
	go server.removeExpiredKeys()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("[err] accepting connection:", err)
			continue
		}
		go server.handleConnection(conn)
	}
}
