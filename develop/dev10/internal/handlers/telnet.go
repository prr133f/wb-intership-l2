package handlers

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func (h *Handler) Telnet(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		h.Log.Error("Not enough arguments")
		fmt.Println("Not enough arguments")
		return
	}

	tout, err := cmd.Flags().GetString("timeout")
	if err != nil {
		h.Log.Error(err.Error())
		fmt.Println(err.Error())
		return
	}

	timeout, err := time.ParseDuration(tout)
	if err != nil {
		h.Log.Error(err.Error())
		fmt.Println(err.Error())
		return
	}

	err = h.Domain.Connect(args[0], args[1], timeout)
	if err != nil {
		h.Log.Error(err.Error())
		fmt.Println(err.Error())
		return
	}
}
