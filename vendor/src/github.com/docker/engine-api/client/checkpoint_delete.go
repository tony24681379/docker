package client

import "net/url"

// CheckpointDelete deletes a checkpoint from the given container with the given name
func (cli *Client) CheckpointDelete(containerID string, imgDir string) error {
	query := url.Values{}
	if imgDir != "" {
		query.Set("imgDir", imgDir)
	}
	resp, err := cli.delete("/checkpoints/"+containerID+"/checkpoint", query, nil)
	ensureReaderClosed(resp)
	return err
}
