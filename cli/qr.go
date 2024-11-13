package main

import (
	qrcode "github.com/skip2/go-qrcode"
	"github.com/urfave/cli/v2"
)

var qrCommands = &cli.Command{
	Name:      "qr",
	Usage:     "QR Code Tools",
	Args:      true,
	ArgsUsage: "PAYLOAD",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "recovery-level",
			Aliases: []string{"r"},
		},
	},
	Action: qrAction,
}

var qrAction = func(ctx *cli.Context) error {
	return nil
}

func generateQr(content string) (string, error) {
	q, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return "", err
	}

	return q.ToString(false), nil
}
