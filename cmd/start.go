package cmd

import (
	"github.com/spf13/cobra"
        "net/http"
	"log"
)

var startCmd = &cobra.Command{
	Use: "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Start golang-seed")
		server := http.NewServeMux()

		server.HandleFunc("/@/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status": "OK"}`))
		})

		return http.ListenAndServe(":8080", server)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
