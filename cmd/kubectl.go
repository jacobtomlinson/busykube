package cmd

import (
    kubectlcmd "k8s.io/kubectl/pkg/cmd"
)

func init() {
    RootCmd.AddCommand(kubectlcmd.NewDefaultKubectlCommand())
}
