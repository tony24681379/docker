package client

import (
	"github.com/docker/engine-api/types"
)

// ContainerCheckpoint creates a checkpoint from the given container with the given name
func (cli *Client) ContainerCheckpoint(containerID string, options types.CriuConfig) error {
	resp, err := cli.post("/containers/"+containerID+"/checkpoint", nil, options, nil)
	ensureReaderClosed(resp)
	return err
}