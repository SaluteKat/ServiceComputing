// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"agenda/entity/Meeting"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// deleteMeetingCmd represents the deleteMeeting command
var deleteMeetingCmd = &cobra.Command{
	Use:   "deleteMeeting",
	Short: "--title=meeting",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		bytes,_ := ioutil.ReadFile("data/current.txt")
		curuser := string(bytes)
		meeting,_:=cmd.Flags().GetString("title")
		fmt.Println( "' " + curuser + "' called: deleteMeeting, title: " + meeting)
		Meeting.DeleteMeeting(meeting)
		
	},
}

func init() {
	RootCmd.AddCommand(deleteMeetingCmd)

	// Here you will define your flags and configuration settings.

	deleteMeetingCmd.Flags().StringP("title", "t", "Anonymous", "Help message for meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

