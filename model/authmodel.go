package model

type SupabasePhoneLogin struct {
	StructPhoneNumber string `json:"phone_number"`
}

type SupabasePhoneVerify struct {
	StructPhoneNumber string `json:"phone_number"`
	StrcutOTPcode     string `json:"token"`
}
