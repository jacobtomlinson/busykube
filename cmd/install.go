package cmd

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
    Use:   "install",
    Short: "Create symlinks for each subcommand in the current binary's directory",
    Run: func(cmd *cobra.Command, args []string) {
        // Get the path to the current binary
        binaryPath, err := os.Executable()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error getting current binary path: %v\n", err)
            os.Exit(1)
        }

        // Get the directory of the current binary
        binaryDir := filepath.Dir(binaryPath)

        // Iterate over all subcommands
        for _, subCmd := range RootCmd.Commands() {
			// Skip specific subcommands
			if subCmd.Name() == "install" || subCmd.Name() == "help" || subCmd.Name() == "completion" || subCmd.Name() == "versions" {
				continue
			}
            symlinkPath := filepath.Join(binaryDir, subCmd.Name())

            // Check if the symlink already exists
            if _, err := os.Lstat(symlinkPath); err == nil {
                fmt.Printf("Symlink for '%s' already exists, skipping...\n", subCmd.Name())
                continue
            }

            // Create the symlink
            err := os.Symlink(binaryPath, symlinkPath)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Error creating symlink for '%s': %v\n", subCmd.Name(), err)
                continue
            }

            fmt.Printf("Created symlink: %s -> %s\n", symlinkPath, binaryPath)
        }
    },
}

func init() {
    RootCmd.AddCommand(installCmd)
}
