package auth

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/spf13/viper"
)

type Paseto struct {
	privateKey paseto.V4AsymmetricSecretKey
	publicKey  paseto.V4AsymmetricPublicKey
	parser     paseto.Parser
}

func NewPaseto() Paseto {
	key := paseto.NewV4AsymmetricSecretKey()

	parser := paseto.NewParser()
	parser.AddRule(paseto.ForAudience("*"))
	parser.AddRule(paseto.IssuedBy(viper.GetString("APP_HOST")))
	parser.AddRule(paseto.NotExpired())

	return Paseto{key, key.Public(), parser}
}

func (p *Paseto) Encode(token paseto.Token) (string, error) {
	signed := token.V4Sign(p.privateKey, nil)

	return signed, nil
}

func (p *Paseto) Decode(signed string) (paseto.Token, error) {
	p.parser.AddRule(paseto.ValidAt(time.Now()))

	parsedToken, err := p.parser.ParseV4Public(p.publicKey, signed, nil)
	if err != nil {
		return paseto.Token{}, err
	}

	return *parsedToken, nil
}
