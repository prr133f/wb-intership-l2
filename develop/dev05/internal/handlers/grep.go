package handlers

import (
	"dev05/internal/domain"
	"os"

	"github.com/spf13/cobra"
)

func (h *Handlers) Grep(cmd *cobra.Command, args []string) {
	var flags domain.Flags
	var err error
	if len(args) < 2 {
		h.Log.Error("Not enough arguments")
		return
	}

	pattern := args[0]
	filenames := args[1:]

	flags.After, err = cmd.Flags().GetInt("after")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.Before, err = cmd.Flags().GetInt("before")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.Context, err = cmd.Flags().GetInt("context")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.Count, err = cmd.Flags().GetBool("count")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.IgnoreCase, err = cmd.Flags().GetBool("ignore-case")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.Invert, err = cmd.Flags().GetBool("invert")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.Fixed, err = cmd.Flags().GetBool("fixed")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}
	flags.LineNum, err = cmd.Flags().GetBool("line-num")
	if err != nil {
		h.Log.Error(err.Error())
		return
	}

	files := make([]*os.File, 0, len(filenames))
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			h.Log.Error(err.Error())
			return
		}
		defer file.Close()

		files = append(files, file)
	}

	if err := h.Domain.Grep(files, pattern, flags); err != nil {
		h.Log.Error(err.Error())
		return
	}
}
