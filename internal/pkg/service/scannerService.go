package service

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

type Connection struct {
	Host      string
	StartPort int
	EndPort   int
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetConnectionDetails() Connection {
	var host string
	var startPort int
	var endPort int

	fmt.Print("Введите адрес хоста: ")
	if _, err := fmt.Fscan(os.Stdin, &host); err != nil {
		log.Fatalf("ошибка при получении адреса хоста: ", err.Error())
		os.Exit(0)
	}

	fmt.Print("Сканируем порты от: ")
	if _, err := fmt.Fscan(os.Stdin, &startPort); err != nil {
		log.Fatalf("ошибка при получении порта: ", err.Error())
		os.Exit(0)
	}

	fmt.Print("Сканируем порты до: ")
	if _, err := fmt.Fscan(os.Stdin, &endPort); err != nil {
		log.Fatalf("ошибка при получении порта: ", err.Error())
		os.Exit(0)
	}
	return Connection{
		Host:      host,
		StartPort: startPort,
		EndPort:   endPort,
	}
}

func (s *Service) ScanPorts(details Connection) {
	var wg sync.WaitGroup

	for details.StartPort <= details.EndPort {
		wg.Add(1)
		go func(host string, port int) {
			isOpen := s.checkPort(host, port)
			if isOpen {
				fmt.Printf("Открытый порт: %s \n", strconv.Itoa(port))
			}
			wg.Done()
		}(details.Host, details.StartPort)
		details.StartPort++
	}
	wg.Wait()
}

func (s *Service) checkPort(host string, port int) bool {

	network := "tcp"
	timeout := 5 * time.Second

	con, err := net.DialTimeout(network, net.JoinHostPort(host, strconv.Itoa(port)), timeout)

	if err != nil {
		return false
	} else {
		if con != nil {
			defer con.Close()
			return true
		}
		return false
	}
}
