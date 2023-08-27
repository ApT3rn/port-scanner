package main

import (
	"fmt"
	"github.com/ApT3rn/port-scanner/internal/pkg/service"
)

func main() {

	s := service.NewService()

	details := s.GetConnectionDetails()

	fmt.Println("Ожидайте, идёт проверка портов!")

	s.ScanPorts(details)

	fmt.Println("Все порты проверены!")
}
