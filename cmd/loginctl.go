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

var loginctlCmd = &cobra.Command{
	Use:   "loginctl",
	Short: "Extract facts from loginctl command",
	Run: func(cmd *cobra.Command, args []string) {
		loginctlAll()
	},
}

func init() {
	rootCmd.AddCommand(loginctlCmd)
}

func loginctlAll() {
	out, err := exec.Command("/usr/bin/loginctl", "show-user", os.Getenv("USER")).Output()
	if err != nil {
		log.Debugf("Failed to execute command: %v\nOutput:\n%v", err, out)
	} else {
		lines := bytes.Split(bytes.TrimSpace(out), []byte("\n"))
		f := func(c rune) bool {
			return c == '=' || unicode.IsControl(c)
		}
		valuesMap := make(map[string]string)
		for l := range lines {
			f := bytes.FieldsFunc(lines[l], f)
			log.Debugf("Key: %s, Value: %s", f[0], f[1])
			valuesMap[string(f[0])] = string(f[1])
		}
		jsonOut, err := json.Marshal(valuesMap)
		if err != nil {
			log.Debugf("Failed to parse to JSON:", err)
			os.Stdout.Write(nil)
		}
		os.Stdout.Write(jsonOut)
	}
}
