/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/javorszky/adventofcode2022/day1"
	"github.com/javorszky/adventofcode2022/day2"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")

		l := zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Str("module", "adventofcode").Int("year", 2022).Logger()
		l.Info().Msg("Welcome to Gabor Javorszky's Advent of Code 2022 solutions!")

		day1.Task1(l)
		day1.Task2(l)
		day2.Task1(l)
		day2.Task2(l)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
