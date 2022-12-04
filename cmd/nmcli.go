/*
MIT License

# Copyright (c) 2021 Josh Rich

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
	"unicode"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type nmcli struct {
	Connections []connection `json:"connection"`
	Devices     []device     `json:"device"`
}

type connection struct {
	Name   string `json:"name"`
	Uuid   string `json:"uuid"`
	Conn   string `json:"type"`
	Device string `json:"device"`
}

type device struct {
	Device     string `json:"device"`
	Medium     string `json:"type"`
	State      string `json:"state"`
	Connection string `json:"connection"`
}

// NetworkManagerCmd represents the NetworkManager command
var nmcliCmd = &cobra.Command{
	Use:   "nmcli",
	Short: "Extract facts from nmcli command",
	Run: func(cmd *cobra.Command, args []string) {
		nmcliAll()
	},
}

func init() {
	rootCmd.AddCommand(nmcliCmd)
	nmcliCmd.AddCommand(connectionCmd)
	nmcliCmd.AddCommand(deviceCmd)
}

var nmcliOutput nmcli

var connectionCmd = &cobra.Command{
	Use:   "go-fact-fetch nmcli connection",
	Short: "Get overview of connections from NetworkManager",
	Run: func(cmd *cobra.Command, args []string) {
		nmcliConnection()
		nmCliJson()
	},
}

var deviceCmd = &cobra.Command{
	Use:   "go-fact-fetch nmcli device",
	Short: "Get overview of devices from NetworkManager",
	Run: func(cmd *cobra.Command, args []string) {
		nmcliDevice()
		nmCliJson()
	},
}

func nmcliAll() {
	nmcliConnection()
	nmcliDevice()
	nmCliJson()
}

func nmcliConnection() {
	out, err := exec.Command("/usr/bin/nmcli", "--terse", "connection", "show").Output()
	if err != nil {
		log.Debugf("Failed to execute command: %v", err)
		nmcliOutput.Connections = nil
	} else {
		lines := bytes.Split(bytes.TrimSpace(out), []byte("\n"))
		f := func(c rune) bool {
			return c == ':' || unicode.IsControl(c)
		}
		nmcliOutput.Connections = make([]connection, len(lines))
		for l := range lines {
			f := bytes.FieldsFunc(lines[l], f)
			c := connection{
				Name: string(f[0]),
				Uuid: string(f[1]),
				Conn: string(f[2]),
			}
			if len(f) > 3 {
				c.Device = string(f[3])
			}
			nmcliOutput.Connections[l] = c
		}
	}
}

func nmcliDevice() {
	out, err := exec.Command("/usr/bin/nmcli", "--terse", "device").Output()
	if err != nil {
		log.Debugf("Failed to execute command: %v", err)
		nmcliOutput.Devices = nil
	} else {
		lines := bytes.Split(bytes.TrimSpace(out), []byte("\n"))
		f := func(c rune) bool {
			return c == ':' || unicode.IsControl(c)
		}
		nmcliOutput.Devices = make([]device, len(lines))
		for l := range lines {
			f := bytes.FieldsFunc(lines[l], f)
			c := device{
				Device: string(f[0]),
				Medium: string(f[1]),
				State:  string(f[2]),
			}
			if len(f) > 3 {
				c.Connection = string(f[3])
			}
			nmcliOutput.Devices[l] = c
		}
	}
}

func nmCliJson() {
	jsonOut, err := json.Marshal(nmcliOutput)
	if err != nil {
		log.Debugf("Failed to parse to JSON:", err)
		os.Stdout.Write(nil)
	}
	os.Stdout.Write(jsonOut)
}
