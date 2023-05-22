package main

import (
	"log"
	"math/rand"
	"net"
	"time"
	"fmt"
)

// Сетевой адрес.
//
// Служба будет слушать запросы на всех IP-адресах
// компьютера на порту 12345.
// Например, 127.0.0.1:12345
const addr = "0.0.0.0:12345"

// Протокол сетевой службы.
const proto = "tcp4"

var sayings []string = []string{"Don't communicate by sharing memory, share memory by communicating",
	"Concurrency is not parallelism",
	"Channels orchestrate; mutexes serialize",
	"The bigger the interface, the weaker the abstraction",
	"Make the zero value useful",
	"interface{} says nothing",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite",
	"A little copying is better than a little dependency",
	"Syscall must always be guarded with build tags",
	"Cgo must always be guarded with build tags",
	"Cgo is not Go",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Запуск сетевой службы по протоколу TCP
	// на порту 12345.
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// Подключения обрабатываются в бесконечном цикле.
	// Иначе после обслуживания первого подключения сервер
	//завершит работу.
	for {
		// Принимаем подключение.
		conn, err := listener.Accept()
		
		if err != nil {
			log.Fatal(err)
		}
		// Вызов обработчика подключения.
		go handleConn(conn)
	}
}

// Обработчик. Вызывается для каждого соединения.
func handleConn(conn net.Conn) {
	// Закрытие соединения.
	defer conn.Close()
	fmt.Println("conn run")
	
	for {
		num := rand.Intn(len(sayings) - 1)
		conn.Write([]byte(sayings[num]+"\r\n\r\n"))
		time.Sleep(time.Duration(3000 * time.Millisecond))
	}
	
	fmt.Println("conn close")
}
