package pkg

type Config struct {
	AZURE_PERSONAL_ACCESS_TOKEN string `env:"AZURE_PERSONAL_ACCESS_TOKEN"`
	AZURE_ORGANIZATION_URL      string `env:"AZURE_ORGANIZATION_URL"`
	AZURE_QUERY_ID              string `env:"AZURE_QUERY_ID"`
	AZURE_PROJECT               string `env:"AZURE_PROJECT"`
	TODOIST_API_KEY             string `env:"TODOIST_API_KEY"`
	TODOIST_PROJECT_ID          string `env:"TODOIST_PROJECT_ID"`
}
