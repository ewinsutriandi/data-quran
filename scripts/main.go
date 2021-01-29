package main

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var httpClient = http.Client{Timeout: 5 * time.Minute}

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-scripts",
		Short: "Download Quran data from various sources",
		Args:  cobra.MaximumNArgs(1),
	}

	rootCmd.AddCommand(tanzilTransCmd())

	err := rootCmd.Execute()
	if err != nil {
		logrus.Fatalln(err)
	}
}
