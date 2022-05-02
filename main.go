package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"

	"google.golang.org/api/gmail/v1"
)

func run(ctx context.Context) error {
	service, err := gmail.NewService(ctx)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	message, err := service.Users.Messages.
		Send(
			"me",
			&gmail.Message{
				Raw: base64.StdEncoding.EncodeToString(bytes),
			},
		).
		Do()
	if err != nil {
		return err
	}

	bytes, err = json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(bytes)

	return err
}

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		_, _ = os.Stdout.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
