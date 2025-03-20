package cmd

import (
    "debug/buildinfo"
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
    Use:   "versions",
    Short: "Print versions of kubectl, kind, and helm",
    Run: func(cmd *cobra.Command, args []string) {
        printVersions()
    },
}

func init() {
    RootCmd.AddCommand(versionsCmd)
}

func printVersions() {
    // Open the current executable
    exe, err := os.Executable()
    if err != nil {
        fmt.Printf("Error getting executable: %v\n", err)
        return
    }

    // Read build information from the executable
    info, err := buildinfo.ReadFile(exe)
    if err != nil {
        fmt.Printf("Error reading build info: %v\n", err)
        return
    }

    // Define the modules to look for
    modules := map[string]string{
        "kubectl": "k8s.io/client-go",
        "kind":    "sigs.k8s.io/kind",
        "helm":    "helm.sh/helm/v3",
    }

    // Find and print the versions
    for name, path := range modules {
        version := findModuleVersion(info, path)
        if version != "" {
            fmt.Printf("%s: %s\n", name, version)
        } else {
            fmt.Printf("%s: unknown\n", name)
        }
    }
}

func findModuleVersion(info *buildinfo.BuildInfo, modulePath string) string {
    for _, dep := range info.Deps {
        if dep.Path == modulePath {
            return dep.Version
        }
    }
    return ""
}
