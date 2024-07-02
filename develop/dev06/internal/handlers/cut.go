package handlers

import (
	"dev06/internal/domain"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func (h *Handler) Cut(cmd *cobra.Command, args []string) {
	var flags domain.Flags
	var err error

	flags.B, err = cmd.Flags().GetString("bytes")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.C, err = cmd.Flags().GetString("characters")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.D, err = cmd.Flags().GetString("delimiter")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.F, err = cmd.Flags().GetString("fields")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.S, err = cmd.Flags().GetBool("separated")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}

	var in io.Reader
	if len(args) == 0 {
		in = os.Stdin
	} else {
		in, err = os.Open(args[0])
		if err != nil {
			h.Log.Error(err.Error())
			return
		}
	}

	ans, err := h.Domain.Cut(in, flags)
	if err != nil {
		h.Log.Error(err.Error())
		return
	}

	for _, row := range ans {
		fmt.Println(row)
	}
}
