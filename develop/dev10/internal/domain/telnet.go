package domain

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func (d *Domain) Connect(host, port string, timeout time.Duration) error {
	fmt.Printf("Connecting to %s:%s\n", host, port)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		d.Log.Error(err.Error())
		return err
	}
	defer conn.Close()
	fmt.Println("Connected")

	stop := make(chan struct{})

	go d.writer(conn, stop)

	go d.reader(conn, stop)

	<-stop

	return nil
}

func (d *Domain) writer(conn net.Conn, stop chan struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-stop:
			return
		default:
			var line string
			if !scanner.Scan() {
				close(stop)
				return
			}
			line = scanner.Text()

			fmt.Fprintln(conn, line)
		}
	}
}

func (d *Domain) reader(conn net.Conn, stop chan struct{}) {
	var (
		in = bufio.NewReader(conn)
	)

	for {
		select {
		case <-stop:
			return
		default:
			var buf [1024]byte
			_, err := in.Read(buf[:])
			if err != nil {
				d.Log.Error(err.Error())
				return
			}

			fmt.Println("-> " + string(buf[:]))
		}
	}
}
