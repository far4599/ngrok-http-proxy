package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token := os.Getenv("NGROK_AUTHTOKEN")
	if len(token) == 0 {
		log.Fatalln("Please provide ngrok auth token via $NGROK_AUTHTOKEN. \nTo get one sign up at https://dashboard.ngrok.com/signup")
	}

	err := setupNgrokToken("./config.yml", token)
	if err != nil {
		log.Fatalf("failed to write token to ngrok config file: '%s'", err)
	}

	errGroup, errCtx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		return startGOST(errCtx)
	})

	errGroup.Go(func() error {
		return startNgrok(ctx)
	})

	err = errGroup.Wait()
	if err != nil {
		log.Fatalf("exit with error: %s", err)
	}
}

func startGOST(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "./gost", "-L=:8000")

	return cmd.Run()
}

func startNgrok(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "./ngrok", "tcp", "8000", "--config=./config.yml")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	return cmd.Run()
}

func setupNgrokToken(file, token string) error {
	config := struct {
		Version   int    `yaml:"version"`
		Authtoken string `yaml:"authtoken"`
	}{
		Version:   2,
		Authtoken: token,
	}

	fileContent, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, fileContent, 0600)
	if err != nil {
		return err
	}

	return nil
}
