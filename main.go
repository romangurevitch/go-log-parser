/*
Copyright © 2019 Roman Gurevitch romanil85@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"bufio"
	"fmt"
	"github.com/romangurevitch/go-log-parser/analyzer"
	"github.com/romangurevitch/go-log-parser/httplog"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const useHelp = "go-log-parser <log file path>"
const shortHelp = "Parse HTTP log file"
const longHelp = `Parse and analyze HTTP log file, reports: 
  Number of unique IP addresses.
  Top 3 most active IP addresses.
  Top 3 most visited URLs.`

func main() {
	var cmdPrint = &cobra.Command{
		Use:   useHelp,
		Short: shortHelp,
		Long:  longHelp,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := analyseLog(args[0]); err != nil {
				log.Fatalln(err)
			}
		},
	}

	if err := cmdPrint.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func analyseLog(logFilePath string) error {
	file, err := os.Open(logFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	logAnalyzer := analyzer.New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := &httplog.Log{}
		ok, err := entry.Extract(scanner.Bytes())
		if !ok {
			if err != nil {
				return err
			}
			continue
		}

		logAnalyzer.AddEntry(entry)
	}

	fmt.Println(logAnalyzer.Report())
	return nil
}
