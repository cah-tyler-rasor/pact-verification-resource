package broker

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func loadFixture(t *testing.T, name string) []byte {
	bytes, err := ioutil.ReadFile("./testdata/" + name)
	if err != nil {
		t.Fatalf("Failed to load fixture %s", name)
	}
	return bytes
}

func TestGetValidationsWithoutTag(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/pacts/provider/PROVIDER/consumer/CONSUMER/verification-results/latest", r.URL.Path)
		assert.Equal(t, r.Method, "GET")
		w.Header().Add("Content-Type", "application/json")
		w.Write(loadFixture(t, "pact_verification.json"))
	}))

	client := NewClient(ts.URL)
	got := client.GetValidation("CONSUMER", "PROVIDER")

	theTime, _ := time.Parse(time.RFC3339, "2020-04-22T19:56:35Z")

	want := HalPactVerification{
		ProviderName: "PROVIDER",
		ProviderApplicationVersion: "PROVIDER_VERSION",
		Success: true,
		VerificationDate: theTime,
		Links: Links {
			Verification: Verification {
				Name: "Verification result RESULT_NUMBER for Pact between CONSUMER (CONSUMER_VERSION) and PROVIDER",
				Href: "BROKER_URL/pacts/provider/PROVIDER/consumer/CONSUMER/pact-version/PACT_VERSION/verification-results/RESULT_NUMBER",
			},
		},
	}

	assert.Equal(t, want, got)
}

func TestGetVersionsWithTag(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/pacts/provider/PROVIDER/consumer/CONSUMER/latest/TAG/verification-results/latest", r.URL.Path)
		assert.Equal(t, r.Method, "GET")
		w.Header().Add("Content-Type", "application/json")
		w.Write(loadFixture(t, "pact_verification.json"))
	}))

	client := NewClient(ts.URL)
	got := client.GetTaggedValidation("CONSUMER", "PROVIDER", "TAG")

	theTime, _ := time.Parse(time.RFC3339, "2020-04-22T19:56:35Z")

	want := HalPactVerification{
		ProviderName: "PROVIDER",
		ProviderApplicationVersion: "PROVIDER_VERSION",
		Success: true,
		VerificationDate: theTime,
		Links: Links {
			Verification: Verification {
				Name: "Verification result RESULT_NUMBER for Pact between CONSUMER (CONSUMER_VERSION) and PROVIDER",
				Href: "BROKER_URL/pacts/provider/PROVIDER/consumer/CONSUMER/pact-version/PACT_VERSION/verification-results/RESULT_NUMBER",
			},
		},
	}

	assert.Equal(t, want, got)
}
