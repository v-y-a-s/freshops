/*
Copyright © 2020 Vyas Sarangapani vyas@freshworks.io

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
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var aboutDesc = `

====================================================================================
=== FRESHOPS CLI ===================================================================
====================================================================================

-> [ Tech Leaders         ]  Common standards and guidelines across all projects
-> [ Developers           ]  Create new projects that adhere to standards
-> [ Quality Assurance    ]  Find technical spec for projects easily
-> [ Support Engineers    ]  Maintain projects with ease
-> [ Project Managers     ]  Audit Report Generation for Projects 
-> [ Human Resources      ]  Onboard new devs easily 
-> [ Resourcing           ]  Move Devs across projects 
-> [ Leadership & Clients ]  Successful project 

====================================================================================

`

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "What is freshops cli?",
	Long: aboutDesc,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(aboutDesc)
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aboutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aboutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
