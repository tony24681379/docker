// +build experimental

package client

import (
	"fmt"

	Cli "github.com/docker/docker/cli"
	flag "github.com/docker/docker/pkg/mflag"
	"github.com/docker/engine-api/types"
)

// CmdRestore restores the process in a checkpointed container
//
// Usage: docker restore CONTAINER
func (cli *DockerCli) CmdRestore(args ...string) error {
	cmd := Cli.Subcmd("restore", []string{"CONTAINER"}, Cli.DockerCommands["restore"].Description, true)
	flImgDir := cmd.String([]string{"-image-dir"}, "", "directory to restore image files from")
	flWorkDir := cmd.String([]string{"-work-dir"}, "", "directory for restore log")
	flForce := cmd.Bool([]string{"-force"}, false, "bypass checks for current container state")

	cmd.Require(flag.Min, 1)
	cmd.ParseFlags(args, true)

	if cmd.NArg() < 1 {
		cmd.Usage()
		return nil
	}

	criuOpts := types.CriuConfig{
		ImagesDirectory: *flImgDir,
		WorkDirectory:   *flWorkDir,
	}

	name := cmd.Arg(0)
	err := cli.client.ContainerRestore(name, criuOpts, *flForce)
	if err != nil {
		fmt.Fprintf(cli.err, "%s\n", err)
		return fmt.Errorf("Error: failed to restore one or more containers")
	}

	fmt.Fprintf(cli.out, "%s\n", name)
	return nil
}
