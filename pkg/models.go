package pkg

type Config struct {
	AzurePersonalAccessToken string `env:"AzureQueryID"`
	AzureOrgURL              string `env:"AzureOrgURL"`
	AzureQueryID             string `env:"AzureQueryID"`
	AzureProjectID           string `env:"TodoIstProjectID"`
	TodoIstAPIKey            string `env:"TodoIstAPIKey"`
	TodoIstProjectID         string `env:"AzureQueryID"`
}
