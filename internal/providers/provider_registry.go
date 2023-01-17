package providers

import (
	"log"
)

type ProviderRegistry struct {
	providers map[string]SMSProvider
}

func NewProviderRegistry(providers ...SMSProvider) *ProviderRegistry {
	registry := ProviderRegistry{
		providers: map[string]SMSProvider{},
	}

	for _, p := range providers {
		registry.providers[p.Name()] = p
	}
	return &registry
}

func (r ProviderRegistry) GetProvider(name string) SMSProvider {
	p, ok := r.providers[name]

	if !ok {
		log.Fatalln("could not find the given provider")
	}

	return p
}
