package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	mapWeb     bool
	baseMapURL = "https://www.nycmesh.net/map"

	mapCmd = &cobra.Command{
		Use:   "map [nodenumber ...]",
		Short: "Show a map of nodes",
		Long:  "Open a map of NYC Mesh, optionally focusing on a set of nodes",
		RunE: func(cmd *cobra.Command, args []string) error {
			var url string
			if len(args) > 0 {
				url = fmt.Sprintf("%s/nodes/%s", baseMapURL, strings.Join(args, "-"))
			} else {
				url = baseMapURL
			}
			opener := "xdg-open"
			log.Printf("opening: %s %s", opener, url)

			return exec.Command("xdg-open", url).Run()
		},
	}
)

func init() {
	rootCmd.AddCommand(mapCmd)
	mapCmd.Flags().BoolVarP(&mapWeb, "web", "w", true, "Open with web browser")
	mapCmd.Flags().StringVarP(&baseMapURL, "url", "u", "https://www.nycmesh.net/map", "URL base for mapping nodes")
}
