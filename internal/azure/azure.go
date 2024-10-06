package azure

import (
	"context"

	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/webapi"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/workitemtracking"
)

type Client struct {
	Client workitemtracking.Client
}

func NewWorkItemClient(ctx context.Context, organizationURL, personalAccessToken string) (Client, error) {
	connection := azuredevops.NewPatConnection(organizationURL, personalAccessToken)

	client, err := workitemtracking.NewClient(ctx, connection)
	if err != nil {
		return Client{}, err
	}
	return Client{
		Client: client,
	}, nil
}

func (c *Client) QueryWorkItems(ctx context.Context, project, queryID string) (*workitemtracking.WorkItemQueryResult, error) {
	query := uuid.MustParse(queryID)

	return c.Client.QueryById(ctx, workitemtracking.QueryByIdArgs{
		Project: &project,
		Id:      &query,
	})
}

func (c *Client) GetWorkItem(ctx context.Context, project string, id *int) (*workitemtracking.WorkItem, error) {
	return c.Client.GetWorkItem(ctx, workitemtracking.GetWorkItemArgs{
		Project: &project,
		Id:      id,
	})
}

func (c *Client) UpdateWorkItemState(ctx context.Context, project string, id int, state string) error {
	path := "/fields/System.State"
	_, err := c.Client.UpdateWorkItem(ctx, workitemtracking.UpdateWorkItemArgs{
		Project: &project,
		Id:      &id,
		Document: &[]webapi.JsonPatchOperation{
			{
				Op:    &webapi.OperationValues.Replace,
				Value: state,
				Path:  &path,
			},
		},
	})
	return err
}
