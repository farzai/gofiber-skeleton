package google_recaptcha

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/farzai/app/pkg/httpclient"
)

type RecaptchaConfig struct {
	SiteKey   string
	SecretKey string
}

type Recaptcha struct {
	client httpclient.Client
	config RecaptchaConfig
}

func NewWithDefaultClient(config RecaptchaConfig) *Recaptcha {
	return New(httpclient.NewClient(), config)
}

func New(client httpclient.Client, config RecaptchaConfig) *Recaptcha {
	return &Recaptcha{
		client: client,
		config: config,
	}
}

func (r *Recaptcha) VerifyV3(token string) error {
	url := "https://www.google.com/recaptcha/api/siteverify"

	payload := map[string]string{
		"secret":  r.config.SecretKey,
		"response": token,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
