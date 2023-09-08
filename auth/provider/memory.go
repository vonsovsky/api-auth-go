package provider

import (
	"github.com/wernerdweight/api-auth-go/auth/contract"
	"github.com/wernerdweight/api-auth-go/auth/entity"
)

// MemoryApiClientProvider is the simplest implementation of the ApiClientProviderInterface
type MemoryApiClientProvider struct {
	memory []entity.MemoryApiClient
}

func (p MemoryApiClientProvider) ProvideByIdAndSecret(id string, secret string) (contract.ApiClientInterface, *contract.AuthError) {
	for _, client := range p.memory {
		if client.Id == id && client.Secret == secret {
			return &client, nil
		}
	}

	return nil, contract.NewAuthError(contract.ClientNotFound, nil)
}

func (p MemoryApiClientProvider) ProvideByApiKey(apiKey string) (contract.ApiClientInterface, *contract.AuthError) {
	for _, client := range p.memory {
		if client.ApiKey == apiKey {
			return &client, nil
		}
	}

	return nil, contract.NewAuthError(contract.ClientNotFound, nil)
}

func NewMemoryApiClientProvider(memory []entity.MemoryApiClient) *MemoryApiClientProvider {
	return &MemoryApiClientProvider{
		memory: memory,
	}
}

// MemoryApiUserProvider is the simplest implementation of the ApiUserProviderInterface
type MemoryApiUserProvider struct {
	memory []entity.MemoryApiUser
}

func (p MemoryApiUserProvider) ProvideByLoginAndPassword(login string, password string) (contract.ApiUserInterface, *contract.AuthError) {
	for _, user := range p.memory {
		if user.Login == login && user.Password == password {
			return &user, nil
		}
	}

	return nil, contract.NewAuthError(contract.UserNotFound, nil)
}

func (p MemoryApiUserProvider) ProvideByToken(token string) (contract.ApiUserInterface, *contract.AuthError) {
	for _, user := range p.memory {
		if user.AccessToken == token {
			return &user, nil
		}
	}

	return nil, contract.NewAuthError(contract.UserNotFound, nil)
}

func NewMemoryApiUserProvider(memory []entity.MemoryApiUser) *MemoryApiUserProvider {
	return &MemoryApiUserProvider{
		memory: memory,
	}
}
