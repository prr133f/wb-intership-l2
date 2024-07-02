package handlers

import "github.com/spf13/cobra"

func (h *Handler) Wget(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		h.Log.Error("wget: no url")
		return
	}

	if err := h.Domain.Wget(args[0]); err != nil {
		h.Log.Error(err.Error())
		return
	}
}
