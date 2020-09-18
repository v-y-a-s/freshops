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
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type freshSpec struct {
	ProjectName      string
	ClientName       string
	RepoType         string
	BranchingModel   string
	CommitStyle      string
	Hosting          string
	ReleaseTypes     string
	IaacModel        string
	BuildType        string
	RollbackStrategy string

	Testing string

	AccessiblitySupported     string
	UptimeBotEnabled          string
	SecurityBotEnabled        string
	ContainerScanningEnabled  string
	StaticCodeAnalysisEnabled string
	DataEncryption            string
	HighAvilabilityEnabled    string
	TestCoverageCalculator    string
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create freshops spec document",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("\nCreating new project spec....\n\n")

		newSpec := freshSpec{}

		// TODO Error handling
		// TODO Enums
		newSpec.ProjectName, _ = getBasicInput("project name")
		newSpec.ClientName, _ = getBasicInput("client name")

		repoTypeOptions := []string{"Mono Repo", "Frontend Repo", "Backend Repo"}
		newSpec.RepoType, _ = showBasicPrompt("repo type", repoTypeOptions)

		branchingOptions := []string{"Trunk Based Development", "Gitflow", "Github Flow"}
		newSpec.BranchingModel, _ = showBasicPrompt("branching model", branchingOptions)

		commitStyleOptions := []string{"Squash & Merge", "Rebase & Merge", "Merge Commit"}
		newSpec.CommitStyle, _ = showBasicPrompt("your commit style", commitStyleOptions)

		hostingOptions := []string{"AWS", "Azure", "GCP", "OpenShift", "Digital Ocean", "On Prem"}
		newSpec.Hosting, _ = showBasicPrompt("project hosting", hostingOptions)

		releaseTypesOptions := []string{"Feature Flags", "Canary", "Blue Green", "Downtime"}
		newSpec.ReleaseTypes, _ = showBasicPrompt("release management", releaseTypesOptions)

		iaacOptions := []string{"Terraform", "Pulumi", "ARM Templates", "AWS Cloud Formation", "Serverless or SAM Templates", "Chef", "Puppet", "AWS CDK"}
		newSpec.IaacModel, _ = showBasicPrompt("IAAC tool", iaacOptions)

		buildTypesOptions := []string{"Module Level Build", "Full Build"}
		newSpec.BuildType, _ = showBasicPrompt("PR Build type", buildTypesOptions)

		rollbackOptions := []string{"Git Revert", "Manual Rollback", "Deploy Previous Release"}
		newSpec.RollbackStrategy, _ = showBasicPrompt("rollback strategy", rollbackOptions)

		binaryOptions := []string{"Yeaps!", "Nopes!", "N/A!"}
		newSpec.UptimeBotEnabled, _ = showBasicPrompt("if uptime / health check bot is enabled", binaryOptions)
		newSpec.SecurityBotEnabled, _ = showBasicPrompt("if security bot is enabled", binaryOptions)
		newSpec.ContainerScanningEnabled, _ = showBasicPrompt("if container scanning is enabled", binaryOptions)
		newSpec.StaticCodeAnalysisEnabled, _ = showBasicPrompt("if static code analysis is enabled", binaryOptions)
		newSpec.DataEncryption, _ = showBasicPrompt("if data encryption", binaryOptions)
		newSpec.HighAvilabilityEnabled, _ = showBasicPrompt("if HA is configured", binaryOptions)
		newSpec.TestCoverageCalculator, _ = showBasicPrompt("if test covergage is configured", binaryOptions)

		var jsonExport []byte
		jsonExport, err := json.MarshalIndent(newSpec, "", " ")
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(jsonExport))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getBasicInput(message string) (string, error) {

	validate := func(input string) error {
		projname := len(input)
		if projname == 0 {
			return errors.New("Invalid entry, try again")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("Enter %s", message),
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}
	return result, nil
}

func showBasicPrompt(message string, options []string) (string, error) {

	promptSelect := promptui.Select{
		Label: fmt.Sprintf("Select %s", message),
		Items: options,
	}

	_, selected, err := promptSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return selected, nil
}
