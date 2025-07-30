package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Lunaryx-org/Lunapackage/model"
)

func SendPhoneNumber(AnonKey, url, PhoneNumber string) error {
	request := model.SupabasePhoneLogin{
		StructPhoneNumber: PhoneNumber,
	}

	reqPayload, err := json.Marshal(request)
	if err != nil {

	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqPayload))
	if err != nil {
		fmt.Println("", err)
		return err
	}

	req.Header.Set("apikey", AnonKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(":", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

func VerifyPhoneOTP(AnonKey, url, PhoneNumber, OTP string) error {

	return nil
}
