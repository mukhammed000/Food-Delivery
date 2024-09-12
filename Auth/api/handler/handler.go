package handler

import (
	"auth/api/helper"
	"auth/config"
	"auth/service"
	"fmt"
	"math/rand"
	"time"
)

type Handler struct {
	Auth *service.AuthService
}

func NewHandler(auth *service.AuthService) *Handler {
	return &Handler{
		Auth: auth}
}

func (h *Handler) SendEmail(req string) (string, error) {
	cfg := config.Load()

	rand.Seed(time.Now().UnixNano())

	code := rand.Intn(899999) + 100000

	from := "muhammadjonxudaynazarov226@gmail.com"
	password := cfg.EMAIL_PASSWORD
	err := helper.SendVerificationCode(helper.Params{
		From:     from,
		Password: password,
		To:       req,
		Message:  fmt.Sprintf("Hi, here is your verification code: %d", code),
		Code:     fmt.Sprint(code),
	})
	if err != nil {
		return "", fmt.Errorf("failed to send verification email: %v", err)
	}

	return fmt.Sprint(code), nil
}
