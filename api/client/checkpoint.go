// +build experimental

package client

import (
	"fmt"

	"github.com/docker/engine-api/types"
	Cli "github.com/docker/docker/cli"
	flag "github.com/docker/docker/pkg/mflag"
)

// CmdCheckpoint checkpoints the process running in a container
//
// Usage: docker checkpoint CONTAINER
func (cli *DockerCli) CmdCheckpoint(args ...string) error {
	cmd := Cli.Subcmd("checkpoint", []string{"CONTAINER"}, Cli.DockerCommands["checkpoint"].Description, true)
	flImgDir       := cmd.String([]string{"-image-dir"}, "", "directory for storing checkpoint image files")
	flWorkDir      := cmd.String([]string{"-work-dir"}, "", "directory for storing log file")
	flLeaveRunning := cmd.Bool([]string{"-leave-running"}, false, "leave the container running after checkpoint")

	cmd.Require(flag.Min, 1)
    
    cmd.ParseFlags(args, true)

	if cmd.NArg() < 1 {
		cmd.Usage()
		return nil
	}

	criuOpts := types.CriuConfig{
		ImagesDirectory: *flImgDir,
		WorkDirectory:   *flWorkDir,
		LeaveRunning:    *flLeaveRunning,
	}

	var encounteredError error
	for _, name := range cmd.Args() {
		err := cli.client.ContainerCheckpoint(name, criuOpts)
		if err != nil {
			fmt.Fprintf(cli.err, "%s\n", err)
			encounteredError = fmt.Errorf("Error: failed to checkpoint one or more containers")
		} else {
			fmt.Fprintf(cli.out, "%s\n", name)
		}
	}
	return encounteredError
}