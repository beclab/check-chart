package main

import (
	"check-chart/cmd"
	"check-chart/validate"
	"fmt"
	"github.com/spf13/cobra"
)

func validChart(dir string) (err error) {
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	err = validate.CheckChartFolder(dir)
	if err != nil {
		return
	}

	err = validate.CheckAppCfg(dir)
	if err != nil {
		return
	}

	err = validate.CheckServiceAccountRole(dir)

	return
}

func main() {

	rootCmd := &cobra.Command{
		Use:   "check-chart",
		Short: "check-chart",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			chart, err := cmd.Flags().GetString("chart")
			if err != nil {
				fmt.Println(err)
				return
			}
			err = validChart(chart)
			if err != nil {
				fmt.Println("check err:", err)
				return
			}
		},
	}

	cfgCmd := &cobra.Command{
		Use:   "cfg",
		Short: "Check if TerminusManifest.yaml is valid",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {

			chart, err := cmd.Flags().GetString("chart")
			if err != nil {
				fmt.Println(err)
				return
			}
			err = validate.AppCfg(chart, true)
			if err != nil {
				fmt.Println("TerminusManifest.yaml is invalid: ", err)
				return
			}
		},
	}

	cfgCmd.Flags().StringP("chart", "c", "", "path of the chart dir")

	rootCmd.Flags().StringP("chart", "c", "", "path of the chart dir")

	rootCmd.AddCommand(cmd.NewVersionCmd())
	rootCmd.AddCommand(cfgCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
