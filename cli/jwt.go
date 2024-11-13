package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/urfave/cli/v2"
	"strings"
	"time"
)

var jwtCommands = &cli.Command{
	Name:  "jwt",
	Usage: "JWT token encoding",
	Subcommands: []*cli.Command{
		{
			Name:      "encode",
			Usage:     "Encode a JWT token.",
			Action:    jwtEncodeAction,
			Args:      true,
			ArgsUsage: "JSON_PAYLOAD",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "secret",
					Aliases:  []string{"s"},
					Usage:    "Sign the token with the secret key `KEY`",
					Required: true,
				},
				&cli.StringFlag{
					Name:    "algorithm",
					Aliases: []string{"a"},
					Usage:   "The signing `ALGORITHM`, HS256, HS384 or HS512",
					Value:   "HS256",
				},
			},
		},
		{
			Name:      "decode",
			Usage:     "decode a JWT token",
			Action:    jwtDecodeAction,
			Args:      true,
			ArgsUsage: "TOKEN",
		},
	},
}

var (
	errJwtSecretMissing    = fmt.Errorf("secret missing")
	errJwtAlgorithmMissing = fmt.Errorf("algorithm missing")
	errJwtAlgorithmUnknown = func(algorithm string) error {
		return fmt.Errorf("unknown algorithm: %s", algorithm)
	}
	errJwtPayloadMissing = fmt.Errorf("payload missing")
)

var jwtEncodeAction = func(ctx *cli.Context) error {
	secret := ctx.String("secret")
	algorithm := ctx.String("algorithm")
	payload := ctx.Args().Get(0)

	res, err := jwtEncode(secret, algorithm, payload, nil)
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}

func jwtEncode(secret string, algorithm string, payload string, timeFn func() time.Time) (string, error) {
	if secret == "" {
		return "", errJwtSecretMissing
	}

	if algorithm == "" {
		return "", errJwtAlgorithmMissing
	}

	if payload == "" {
		return "", errJwtPayloadMissing
	}

	methods := map[string]jwt.SigningMethod{
		"HS256": jwt.SigningMethodHS256,
		"HS384": jwt.SigningMethodHS384,
		"HS512": jwt.SigningMethodHS512,
	}

	method, ok := methods[strings.ToUpper(algorithm)]
	if !ok {
		return "", errJwtAlgorithmUnknown(algorithm)
	}

	now := time.Now()
	if timeFn != nil {
		now = timeFn()
	}

	token := jwt.NewWithClaims(method, jwt.MapClaims{
		"data": payload,
		"exp":  now.Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return tokenString, nil
}

var jwtDecodeAction = func(ctx *cli.Context) error {
	tokenStr := ctx.Args().Get(0)

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := token.Method(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {

	}

	res, err := jwtDecode(token)
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}

type JWT struct {
	Header    map[string]interface{}
	Payload   map[string]interface{}
	Signature string
}

var (
	errJwtInvalidTokenFormat = fmt.Errorf("invalid token format")
)

// ParseJWT parses and prints the JWT token parts
func parseJwt(token string) (*JWT, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errJwtInvalidTokenFormat
	}

	var header map[string]interface{}
	if err := json.Unmarshal([]byte(headerDecoded), &header); err != nil {
		return nil, fmt.Errorf("error unmarshaling header: %v", err)
	}

	// Decode Payload
	payloadDecoded, err := DecodeBase64(parts[1])
	if err != nil {
		return nil, fmt.Errorf("error decoding payload: %v", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal([]byte(payloadDecoded), &payload); err != nil {
		return nil, fmt.Errorf("error unmarshaling payload: %v", err)
	}

	// Signature is simply the third part, already base64url encoded
	signature := parts[2]

	return &JWT{
		Header:    header,
		Payload:   payload,
		Signature: signature,
	}, nil
}

func jwtDecode(tokenStr string) (jwt.Claims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to decode the JWT token")
	}

	return claims, nil
}
