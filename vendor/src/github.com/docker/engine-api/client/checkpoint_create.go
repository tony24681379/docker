package client

import (
	"github.com/docker/engine-api/types"
)

// CheckpointCreate creates a checkpoint from the given container with the given name
func (cli *Client) CheckpointCreate(containerID string, options types.CriuConfig) error {
	resp, err := cli.post("/checkpoints/"+containerID+"/checkpoint", nil, options, nil)
	ensureReaderClosed(resp)
	return err
}
