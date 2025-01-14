package cmd

import (
	"fmt"

	"github.com/gvcgo/saw/blade"
	"github.com/gvcgo/saw/config"
	"github.com/spf13/cobra"
)

// TODO: colorize based on logGroup prefix (/aws/lambda, /aws/kinesisfirehose, etc...)
var groupsConfig config.Configuration

var groupsCommand = &cobra.Command{
	Use:   "groups",
	Short: "List log groups",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		b := blade.NewBlade(&groupsConfig, &awsConfig, nil)
		logGroups := b.GetLogGroups()
		for _, group := range logGroups {
			fmt.Println(*group.LogGroupName)
		}
	},
}

func init() {
	groupsCommand.Flags().StringVar(&groupsConfig.Prefix, "prefix", "", "log group prefix filter")
}
