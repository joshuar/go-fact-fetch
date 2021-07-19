/*
MIT License

Copyright (c) 2021 Joshua Rich

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

// tunedAdmCmd represents the tunedAdm command
var tunedCmd = &cobra.Command{
	Use:   "tuned",
	Short: "Extract facts from tuned-adm command",
	Run: func(cmd *cobra.Command, args []string) {
		tunedAll()
	},
}

func init() {
	rootCmd.AddCommand(tunedCmd)
}

type tuned struct {
	ActiveProfile string `json:"active_profile"`
}

var tunedOutput tuned

func tunedAll() {
	out, err := exec.Command("/usr/sbin/tuned-adm", "active").Output()
	if err != nil {
		log.Fatal(err)
	}
	f := func(c rune) bool {
		return c == ':' || unicode.IsControl(c)
	}
	fields := bytes.FieldsFunc(out, f)
	tunedOutput.ActiveProfile = string(bytes.TrimSpace(fields[1]))
	jsonOut, err := json.Marshal(tunedOutput)
	if err != nil {
		log.Errorf("error:", err)
	}
	os.Stdout.Write(jsonOut)
}
