package main
import (
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/builtin/credential/github"
)

type VaultClient struct {
	client       *api.Client
	token        string
	vaultAddress string
	vaultOrg     string
}

// NewVaultClient - creates a new vault client
func NewVaultClient(token string, vaultAddress string, vaultOrg string) *VaultClient {
	vaultCFG := api.DefaultConfig()
	vaultCFG.Address = vaultAddress
	vClient, err := api.NewClient(vaultCFG)
	if err != nil {
		return nil
	} else {
		return &VaultClient{client: vClient,
			token:        token,
			vaultAddress: vaultAddress,
			vaultOrg:     vaultOrg}
	}

}

// Authenticate - authenticates access to vault
func (vc *VaultClient) Authenticate() error {
	mountInput := map[string]string{
		"mount": vc.vaultOrg,
		"token": vc.token,
	}
	cliHandler := github.CLIHandler{}

	cTok, err := cliHandler.Auth(vc.client, mountInput)
	if err != nil {
		return err
	}
	token, err := cTok.TokenID()
	if err != nil {
		return err
	}
	vc.client.SetToken(token)
	return nil
}

// Read - returns secrets from given vault path
func (vc *VaultClient) Read(secretPath string) (map[string]interface{}, error) {
	secrets, err := vc.client.Logical().Read(secretPath)
	if err != nil {
		return nil, err
	}
	return secrets.Data, nil
}

