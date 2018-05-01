package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

func main() {
	outputPath := flag.String("out", "", "output tar file")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Specify image names to download.\n")
		os.Exit(1)
	}

	out := os.Stdout
	if len(*outputPath) > 0 {
		f, err := os.Create(*outputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create the destination file: %v", err)
			os.Exit(1)
		}
		defer f.Close()
		out = f
	}

	for _, image := range flag.Args() {
		if err := download(image, os.Stderr, out); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to download image: %v", err)
			os.Exit(1)
		}
	}
}

func download(image string, stdout io.Writer, out io.Writer) error {
	d, err := NewDocker(stdout)
	if err != nil {
		return err
	}
	if err := d.pull(image); err != nil {
		return fmt.Errorf("failed to pull: %v", err)
	}
	if err := d.save(image, out); err != nil {
		return fmt.Errorf("failed to save: %v", err)
	}
	return nil
}

type Docker struct {
	cli    *docker.Client
	stdout io.Writer
}

func NewDocker(stdout io.Writer) (*Docker, error) {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &Docker{cli, stdout}, nil
}

func (d *Docker) pull(image string) error {
	resp, err := d.cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	return d.handleResponse(resp)
}

func (d *Docker) save(image string, out io.Writer) error {
	resp, err := d.cli.ImageSave(context.Background(), []string{image})
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, resp); err != nil {
		return err
	}
	return nil
}

func (d *Docker) handleResponse(resp io.ReadCloser) error {
	r := bufio.NewReader(resp)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		var v map[string]interface{}
		if err := json.Unmarshal(line, &v); err != nil {
			return fmt.Errorf("failed to unmarshal response. (line=%s, err=%v)", string(line), err)
		}
		status, exist := v["status"]
		if !exist {
			return fmt.Errorf("Unexpected response returned from daemon: %v", v)
		}
		fmt.Fprintln(d.stdout, status)
	}
}
