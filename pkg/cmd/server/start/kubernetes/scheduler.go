package kubernetes

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"k8s.io/kubernetes/pkg/util"
	"k8s.io/kubernetes/plugin/cmd/kube-scheduler/app"
)

const schedulerLong = `
Start Kubernetes scheduler

This command launches an instance of the Kubernetes controller-manager (kube-controller-manager).`

// NewSchedulerCommand provides a CLI handler for the 'scheduler' command
func NewSchedulerCommand(name, fullName string, out io.Writer) *cobra.Command {
	s := app.NewSchedulerServer()

	cmd := &cobra.Command{
		Use:   name,
		Short: "Launch Kubernetes scheduler (kube-scheduler)",
		Long:  controllersLong,
		Run: func(c *cobra.Command, args []string) {
			startProfiler()

			util.InitLogs()
			defer util.FlushLogs()

			if err := s.Run(pflag.CommandLine.Args()); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}
	cmd.SetOutput(out)

	flags := cmd.Flags()
	flags.SetNormalizeFunc(util.WordSepNormalizeFunc)
	flags.AddGoFlagSet(flag.CommandLine)
	s.AddFlags(flags)

	return cmd
}
