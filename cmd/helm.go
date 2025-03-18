package cmd

import (
	"os"
	helmcmd "helm.sh/helm/v4/pkg/cmd"
)


func init() {
    cmd, err := helmcmd.NewRootCmd(os.Stdout, os.Args[1:])
    if err != nil {
        os.Exit(1)
    }
    RootCmd.AddCommand(cmd)
}
