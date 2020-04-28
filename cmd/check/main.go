package main

import (
	"encoding/json"
	"github.com/cah-tylerrasor/pact-verification-resource/pkg/broker"
	"github.com/cah-tylerrasor/pact-verification-resource/pkg/concourse"
	"os"
	"sort"
	"strings"
)

func main() {
	var request concourse.CheckRequest
	populateRequest(&request)

	client := broker.NewClient(request.Source.BrokerURL)

	if request.Source.Username != nil && request.Source.Password != nil {
		broker.WithBasicAuth(*request.Source.Username, *request.Source.Password)(client)
	}

	var providerVerifications []concourse.Version
	for _, p := range request.Source.Providers {
		var validation broker.HalPactVerification
		if request.Source.Tag == nil || *request.Source.Tag == "" {
			validation = client.GetValidation(request.Source.Consumer, p)
		} else {
			validation = client.GetTaggedValidation(request.Source.Consumer, p, *request.Source.Tag)
		}

		parsedVersion := strings.SplitAfter(validation.Links.Verification.Href, "pact-version/")[1]
		parsedVersion = strings.Split(parsedVersion, "/verification-results")[0]
		providerVerifications = append(providerVerifications, concourse.Version{
			Provider: p,
			ProviderVersion: validation.ProviderApplicationVersion,
			UpdatedAt: validation.VerificationDate,
			PactVersion: parsedVersion,
		})
	}

	sort.SliceStable(providerVerifications, func(i, j int) bool {
		return providerVerifications[i].UpdatedAt.Before(providerVerifications[j].UpdatedAt)
	})

	if err := json.NewEncoder(os.Stdout).Encode(providerVerifications); err != nil {
		concourse.FailTask("error while encoding response: %s", err)
	}
}

func populateRequest(req *concourse.CheckRequest) {
	if err := json.NewDecoder(os.Stdin).Decode(req); err != nil {
		concourse.FailTask("could not decode request: %s", err)
	}
}
