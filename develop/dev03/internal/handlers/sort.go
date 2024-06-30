package handlers

import (
	"dev03/internal/domain"
	"os"

	"github.com/spf13/cobra"
)

func (h *Handler) Sort(cmd *cobra.Command, args []string) {
	var flags domain.Flags
	if args == nil {
		h.Log.Error("Empty args")
		return
	}

	inName := args[0]

	outName, err := cmd.Flags().GetString("output")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}

	// Read flags
	flags.K, err = cmd.Flags().GetInt("column")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.N, err = cmd.Flags().GetBool("numeric")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.R, err = cmd.Flags().GetBool("reverse")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.U, err = cmd.Flags().GetBool("unique")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}

	file, err := os.Open(inName)
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	defer file.Close()

	outFile, err := os.OpenFile(outName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	defer outFile.Close()

	if err := h.Domain.Sort(file, outFile, flags); err != nil {
		h.Log.Error(err.Error())
		return
	}
}
