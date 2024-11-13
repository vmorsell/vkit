package main

import (
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli/v2"
)

var base64Commands = &cli.Command{
	Name:    "base64",
	Aliases: []string{"b64"},
	Usage:   "base64 tools",
	Subcommands: []*cli.Command{
		{
			Name:   "encode",
			Usage:  "encode a string to base64",
			Action: base64EncodeAction,
		},
		{
			Name:   "decode",
			Usage:  "decode a base64 encode string",
			Action: base64DecodeAction,
		},
	},
}

// base64EncodeAction handles the encode call by the CLI
var base64EncodeAction = func(ctx *cli.Context) error {
	in := ctx.Args().Get(0)

	res := base64.StdEncoding.EncodeToString([]byte(in))

	fmt.Println(res)
	return nil
}

// base64DecodeAction is the decoding function called by the CLI
var base64DecodeAction = func(ctx *cli.Context) error {
	in := ctx.Args().Get(0)

	bytes, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	return nil
}
