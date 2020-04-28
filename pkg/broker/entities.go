package broker

import "time"

type (
	Verification struct {
		Name string `json:"name"`
		Href string `json:"href"`
	}

	Links struct {
		Verification Verification `json:"self"`
	}

	HalPactVerification struct {
		ProviderName string `json:"providerName"`
		ProviderApplicationVersion string `json:"providerApplicationVersion"`
		Success bool `json:"success"`
		VerificationDate time.Time `json:"verificationDate"`
		Links Links `json:"_links"`
	}
)
