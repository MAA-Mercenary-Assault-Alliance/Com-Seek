package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type RecaptchaResponse struct {
	Success            bool     `json:"success"`
	ChallengeTimestamp string   `json:"challenge_ts"`
	Hostname           string   `json:"hostname"`
	ErrorCodes         []string `json:"error-codes"`
}

var ErrRecaptchaFailed = errors.New("reCAPTCHA verification failed")

func VerifyRecaptchaToken(token string) error {
	var recaptchaResponse RecaptchaResponse

	secretKey := os.Getenv("RECAPTCHA_SECRET_KEY")

	res, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
		url.Values{
			"secret":   {secretKey},
			"response": {token},
		})

	if err != nil {
		return fmt.Errorf("failed to verify reCAPTCHA token: %w", err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&recaptchaResponse); err != nil {
		return fmt.Errorf("failed to read response from Google: %w", err)
	}

	if !recaptchaResponse.Success {
		return ErrRecaptchaFailed
	}

	return nil
}
