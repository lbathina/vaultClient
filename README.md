# vaultClient

Demonstrates the HashiCorp Vault GO API usage.

**Create VaultClient Object**

vaultClient := common.NewVaultClient(gitToken, vaultAddress, vaultOrg)


**Authenticate - Login to vault.**

vaultClient.Authenticate()

**use the vaultClient object to perform various read/write/list operations.**

vaultClient.Read(vaultPath)