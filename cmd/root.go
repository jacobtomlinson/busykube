package cmd

import (
    "github.com/spf13/cobra"
    "os"
)

var RootCmd = &cobra.Command{
    Use:   "busykube",
    Short: "A CLI tool that combines the most popular Kubernetes tools",
    Long:  `This tool provides a unified interface for Kubernetes management by combining kubectl and helm commands.`,
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

func Execute() {
    // Check if os.Args[0] matches any subcommand name
    for _, subCmd := range RootCmd.Commands() {
        if os.Args[0] == subCmd.Name() {
            // Inject the subcommand name into os.Args
            os.Args = append([]string{RootCmd.Use, subCmd.Name()}, os.Args[1:]...)
            // Fall back to the root command, which will now process the subcommand
            break
        }
    }

    // Execute the root command (which now includes the subcommand in os.Args if matched)
    if err := RootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
