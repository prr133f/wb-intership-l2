package handlers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"errors"
)

func (h *Handler) Run() error {
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := in.ReadString('\n')
		if err != nil {
			h.Log.Error(err.Error())
			return err
		}

		args := strings.Fields(input)

		output, err := h.Domain.ServeCommand(args[0], args[1:])
		if err != nil {
			if errors.Is(err, errors.ErrUnsupported) {
				print(strings.NewReader("unsupported command\n"))
				h.Log.Error(err.Error())
				continue
			}
			h.Log.Error(err.Error())
			return err
		}

		if output == nil {
			continue
		}
		print(output)
	}
}

func print(r io.Reader) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
