// +build experimental

package cli

func init() {
	experimentalCommands := []Command{
		{"checkpoint", "Checkpoint one or more running containers"},
		{"restore", "Restore one or more checkpointed containers"},
		{"migrate", "Migrate one or more running containers"},
	}

	for _, cmd := range experimentalCommands {
		DockerCommands[cmd.Name] = cmd
	}

	dockerCommands = append(dockerCommands, experimentalCommands...)
}
