package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd export ...
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is simple CLI task tool",
}
