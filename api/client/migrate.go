// +build experimental

package client

import (
	"fmt"

	Cli "github.com/docker/docker/cli"
	"github.com/docker/docker/opts"
	flag "github.com/docker/docker/pkg/mflag"
	runconfigopts "github.com/docker/docker/runconfig/opts"
	"github.com/docker/engine-api/types"
)

// CmdMigrate migrate the process running in a container
//
// Usage: docker migrate CONTAINER
func (cli *DockerCli) CmdMigrate(args ...string) error {
	cmd := Cli.Subcmd("migrate", []string{"CONTAINER"}, Cli.DockerCommands["migrate"].Description, true)
	flEnv := opts.NewListOpts(runconfigopts.ValidateEnv)
	flLabels := opts.NewListOpts(runconfigopts.ValidateEnv)
	flEnvFile := opts.NewListOpts(nil)
	flLabelsFile := opts.NewListOpts(nil)
	cmd.Var(&flLabels, []string{"l", "-label"}, "Set meta data on a container")
	cmd.Var(&flLabelsFile, []string{"-label-file"}, "Read in a line delimited file of labels")
	cmd.Var(&flEnv, []string{"e", "-env"}, "Set environment variables")
	cmd.Var(&flEnvFile, []string{"-env-file"}, "Read in a file of environment variables")

	cmd.Require(flag.Min, 1)

	cmd.ParseFlags(args, true)

	if cmd.NArg() < 1 {
		cmd.Usage()
		return nil
	}

	// collect all the environment variables for the container
	envVariables, err := readKVStrings(flEnvFile.GetAll(), flEnv.GetAll())
	if err != nil {
		return err
	}

	// collect all the labels for the container
	labels, err := readKVStrings(flLabelsFile.GetAll(), flLabels.GetAll())
	if err != nil {
		return err
	}

	filters := types.MigrateFiltersConfig{
		EnvVariables: envVariables,
		Labels:       runconfigopts.ConvertKVStringsToMap(labels),
	}

	var encounteredError error
	for _, name := range cmd.Args() {
		err := cli.client.ContainerMigrate(name, filters)
		if err != nil {
			fmt.Fprintf(cli.err, "%s\n", err)
			encounteredError = fmt.Errorf("Error: failed to migrate one or more containers")
		} else {
			fmt.Fprintf(cli.out, "%s\n", name)
		}
	}
	return encounteredError
}

// reads a file of line terminated key=value pairs and override that with override parameter
func readKVStrings(files []string, override []string) ([]string, error) {
	envVariables := []string{}
	for _, ef := range files {
		parsedVars, err := runconfigopts.ParseEnvFile(ef)
		if err != nil {
			return nil, err
		}
		envVariables = append(envVariables, parsedVars...)
	}
	// parse the '-e' and '--env' after, to allow override
	envVariables = append(envVariables, override...)

	return envVariables, nil
}
