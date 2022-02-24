// Copyright © 2021 NAME HERE <EMAIL ADDRESS>
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

/*
import (
	"Hybrid_Cloud/hybridctl/util"
	"fmt"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

var describeAddonInput eks.DescribeAddonInput

// describeAddonCmd represents the describeAddon command
var describeAddonCmd = &cobra.Command{
	Use:   "describe-addon",
	Short: "A brief description of your command",
	Long: `
	- describe-addon
		hybridctl describe-addon <clusterName> <addonName>

	- platform
		- eks (elastic kubernetes service)`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		if len(args) == 0 || len(args) == 1 {
			fmt.Println("Run 'hybridctl join --help' to view all commands")
		} else if args[1] == "" {
			fmt.Println("Run 'hybridctl join --help' to view all commands")
		} else {
			describeAddonInput.ClusterName = &args[0]
			describeAddonInput.AddonName = &args[1]
			describeAddon(describeAddonInput)
		}
	},
}

func describeAddon(describeAddonInput eks.DescribeAddonInput) {
	httpPostUrl := "http://localhost:8080/describeAddon"
	var output eks.DescribeAddonOutput
	// var message util.Addon
	util.GetJson(httpPostUrl, describeAddonInput, &output)
	fmt.Println(output)
	// bytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Println(bytes)
	// }
	// defer response.Body.Close()
	// json.Unmarshal(bytes, &output)
	// fmt.Println(string(bytes))
	// if output.Addon == nil {
	// 	fmt.Println("A")
	// 	json.Unmarshal(bytes, &message)
	// 	fmt.Println(message.Message_)
	// } else {
	// 	fmt.Println(output)
	// }
}

func init() {
	EksCmd.AddCommand(describeAddonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//describeAddonCmd.Flags().StringVarP(&cobrautil.Option_context, "context", "c", "", "input a option")
}
*/
