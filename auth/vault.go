package auth

import (
	"github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
)

func InitVault(viper *viper.Viper) (*api.Client, error) {
	config := api.DefaultConfig()
	config.Address = viper.GetString("VAULT_URI")

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	client.SetToken(viper.GetString("VAULT_TOKEN"))
	return client, nil
}

// Example function to encrypt (tokenize) data using Vault
func Tokenize(client *api.Client, email string) (string, error) {
	// Encrypt the email using Vault's Transit Secrets Engine
	secret, err := client.Logical().Write("transit/encrypt/my-key", map[string]interface{}{
		"plaintext": email,
	})
	if err != nil {
		return "", err
	}

	token := secret.Data["ciphertext"].(string)
	return token, nil
}

// Example function to decrypt (de-tokenize) data using Vault
func Detokenize(client *api.Client, token string) (string, error) {
	// Decrypt the token using Vault's Transit Secrets Engine
	secret, err := client.Logical().Write("transit/decrypt/my-key", map[string]interface{}{
		"ciphertext": token,
	})
	if err != nil {
		return "", err
	}

	email := secret.Data["plaintext"].(string)
	return email, nil
}
