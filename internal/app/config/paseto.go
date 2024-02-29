package config

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/spf13/viper"
)

func NewPaseto(viper *viper.Viper) (func(id string) (string, error), func(token string) (paseto.Token, error)) {
	issuer := viper.GetString("APP_HOST")

	secretKey := paseto.NewV4AsymmetricSecretKey()
	pubKey := secretKey.Public()

	encode := func(id string) (string, error) {
		token := paseto.NewToken()

		token.SetAudience("*")
		token.SetIssuer(issuer)
		token.SetSubject(id)

		token.SetExpiration(time.Now().Add(2 * time.Hour))
		token.SetNotBefore(time.Now())
		token.SetIssuedAt(time.Now())

		signed := token.V4Sign(secretKey, nil)

		return signed, nil
	}

	parser := paseto.NewParser()
	parser.AddRule(paseto.ForAudience("*"))
	parser.AddRule(paseto.IssuedBy(issuer))
	parser.AddRule(paseto.NotExpired())

	decode := func(token string) (paseto.Token, error) {
		parser.AddRule(paseto.ValidAt(time.Now()))

		parsedToken, err := parser.ParseV4Public(pubKey, token, nil)
		if err != nil {
			return paseto.Token{}, err
		}

		return *parsedToken, nil
	}

	return encode, decode
}
