package main

import (
	"advent_of_code/calendar"
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
)

var rootCmd = &cobra.Command{
	Use:   "aoc <day>",
	Short: "Run Advent of Code",
	Long: `Run Advent of Code

  Specify the day to run as an argument
  Optionally, specify part 1 or part 2
  Without an argument, all days and all parts will run`,
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if err = cobra.MaximumNArgs(1)(cmd, args); err != nil {
			fmt.Printf("Must specify only up to one day\n")
			return
		}
		if len(args) == 1 {
			if ok, _ := regexp.MatchString(`^\d{1,2}$`, args[0]); ok != true {
				fmt.Printf("%s is an invalid number for a day\n", args[0])
				return
			}
		}
		return
	},
	Run: aoc,
}

var debug bool
var part int

func init() {
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Set log level to debug")
	rootCmd.Flags().IntVarP(&part, "part", "p", 1, "Part 1 or 2. Default is 1.")
}

func main() {
	rootCmd.Execute()
}

func aoc(cmd *cobra.Command, args []string) {
	day := "0"
	if len(args) == 1 {
		day = args[0]
	}
	calendar.Run(day, part, debug)
}
