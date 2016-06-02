// +build experimental

package client

import (
	"fmt"

	Cli "github.com/docker/docker/cli"
	flag "github.com/docker/docker/pkg/mflag"
	"github.com/docker/engine-api/types"
)

// CmdCheckpoint checkpoints the process running in a container
//
// Usage: docker checkpoint CONTAINER
func (cli *DockerCli) CmdCheckpoint(args ...string) error {
	cmd := Cli.Subcmd("checkpoint", []string{"CONTAINER"}, Cli.DockerCommands["checkpoint"].Description, true)
	flImgDir := cmd.String([]string{"-image-dir"}, "", "directory for storing checkpoint image files")
	flWorkDir := cmd.String([]string{"-work-dir"}, "", "directory for storing log file")
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

	name := cmd.Arg(0)
	err := cli.client.CheckpointCreate(name, criuOpts)
	if err != nil {
		fmt.Fprintf(cli.err, "%s\n", err)
		return fmt.Errorf("Error: failed to checkpoint container")
	}

	fmt.Fprintf(cli.out, "%s\n", name)
	return nil
}

// CmdCheckpointDelete deletes a container's checkpoint
//
// Usage: docker checkpoint delete <CONTAINER> <CHECKPOINT>
func (cli *DockerCli) CmdCheckpointDelete(args ...string) error {
	cmd := Cli.Subcmd("checkpoint delete", []string{"CONTAINER CHECKPOINT"}, "Delete a container's checkpoint", false)
	flImgDir := cmd.String([]string{"-image-dir"}, "", "directory for storing checkpoint image files")
	cmd.Require(flag.Min, 1)

	cmd.ParseFlags(args, true)
	if cmd.NArg() < 1 {
		cmd.Usage()
		return nil
	}

	criuOpts := types.CriuConfig{
		ImagesDirectory: *flImgDir,
	}
	return cli.client.CheckpointDelete(cmd.Arg(0), criuOpts.ImagesDirectory)
}
