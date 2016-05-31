package client

import "github.com/docker/engine-api/types"

// ContainerMigrate migrates a container from the given container with the given name to another node
func (cli *Client) ContainerMigrate(containerID string, filters types.MigrateFiltersConfig) error {
	resp, err := cli.post("/containers/"+containerID+"/migrate", nil, filters, nil)
	ensureReaderClosed(resp)
	return err
}
