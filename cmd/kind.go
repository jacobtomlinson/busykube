package cmd

import (
	"os"
	"sigs.k8s.io/kind/pkg/cmd"
	"sigs.k8s.io/kind/pkg/cmd/kind"
)


func init() {
	logger, streams := cmd.NewLogger(), cmd.StandardIOStreams()
    c := kind.NewCommand(logger, streams)
	c.SetArgs(os.Args[1:])
    RootCmd.AddCommand(c)
}
