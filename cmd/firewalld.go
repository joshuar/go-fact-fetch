/*
MIT License

# Copyright (c) 2021 Joshua Rich

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// firewalldCmd represents the firewalld command
var firewalldCmd = &cobra.Command{
	Use:   "firewalld",
	Short: "Extra facts for firewalld (firewall-cmd)",
	Run: func(cmd *cobra.Command, args []string) {
		firewalldAll()
	},
}

func init() {
	rootCmd.AddCommand(firewalldCmd)
}

type firewalld struct {
	DefaultZone string `json:"default_zone"`
}

var firewalldOutput firewalld

func firewalldAll() {
	getDefaultZone()

	jsonOut, err := json.Marshal(firewalldOutput)
	if err != nil {
		log.Debugf("Failed to parse to JSON:", err)
		return
	}
	os.Stdout.Write(jsonOut)
}

func getDefaultZone() {
	out, err := exec.Command("/usr/bin/firewall-cmd", "--get-default-zone").Output()
	if err != nil {
		log.Debugf("Failed to execute command: %v", err)
		firewalldOutput.DefaultZone = ""
	} else {
		firewalldOutput.DefaultZone = string(bytes.TrimSpace(out))
	}
}
