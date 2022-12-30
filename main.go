package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"golang.org/x/sync/errgroup"
)

const proxyPort = 8000

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token := os.Getenv("NGROK_AUTHTOKEN")
	if len(token) == 0 {
		log.Fatalln("Please provide ngrok auth token via $NGROK_AUTHTOKEN. \nTo get one visit https://dashboard.ngrok.com/get-started/your-authtoken")
	}

	errGroup, errCtx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		return startGOST(errCtx)
	})

	errGroup.Go(func() error {
		return startNgrok(ctx, token)
	})

	if err := errGroup.Wait(); err != nil {
		log.Fatalf("exit with error: %s", err)
	}
}

func startGOST(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "gost", fmt.Sprintf("-L=:%d", proxyPort))

	cmd.Stdout = stdout{}

	err := cmd.Run()
	if err != nil {
		return cmd.Err
	}

	return err
}

func startNgrok(ctx context.Context, token string) error {
	cmd := exec.CommandContext(ctx, "ngrok", "tcp", fmt.Sprint(proxyPort), "--log", "stdout", "--authtoken", token)

	cmd.Stdout = stdout{}

	err := cmd.Run()
	if err != nil {
		return cmd.Err
	}

	return err
}

type stdout struct{}

func (stdout) Write(p []byte) (int, error) {
	if bytes.Contains(p, []byte("started tunnel")) {
		url := regexp.MustCompile("url=tcp://(\\S+)").FindAllSubmatch(p, -1)
		if len(url) > 0 {
			fmt.Printf("HTTP proxy listens at '%s'", url[0][1])
		}
	} else if bytes.Contains(p, []byte("error")) {
		fmt.Printf("Error: '%s'", p)
	}

	return len(p), nil
}
