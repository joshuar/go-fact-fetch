// Copyright (c) 2023 Joshua Rich <joshua.rich@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type systemctl struct {
	Version  string   `json:"version"`
	Features []string `json:"features"`
}

// zswapAdmCmd represents the zswapAdm command
var systemCtlCmd = &cobra.Command{
	Use:   "systemctl",
	Short: "Extract facts from systemctl command",
	Run: func(cmd *cobra.Command, args []string) {
		systemCtlAll()
	},
}

func init() {
	rootCmd.AddCommand(systemCtlCmd)
}

func systemCtlAll() {
	systemCtlVersion()
}

func systemCtlVersion() {
	out, err := exec.Command("/bin/systemctl", "--version").Output()
	if err != nil {
		log.Debug().Err(err).Msg("Failed to execute command.")
	} else {
		lines := bytes.Split(bytes.TrimSpace(out), []byte("\n"))
		versionRe := regexp.MustCompile(`^systemd\s(\d+)`)
		version := versionRe.FindStringSubmatch(string(lines[0]))
		if version == nil {
			log.Debug().Msg("No version match.")
		}
		features := strings.Split(string(lines[1]), " ")

		s := &systemctl{
			Version:  version[1],
			Features: features,
		}

		jsonOut, err := json.Marshal(s)
		if err != nil {
			log.Debug().Err(err).Msg("Failed to marshal JSON.")
			os.Exit(-1)
		}
		os.Stdout.Write(jsonOut)
	}
}
