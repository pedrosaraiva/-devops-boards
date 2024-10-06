package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/caarlos0/env"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/workitemtracking"
	"github.com/pedrosaraiva1/devops-boards/internal/azure"
	"github.com/pedrosaraiva1/devops-boards/internal/todoist"
	"github.com/pedrosaraiva1/devops-boards/pkg"
)

const (
	limit = 30
)

func loadConfig() pkg.Config {
	var cfg pkg.Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
func CreateNewWorkItem(ctx context.Context, azureClient azure.Client, todoistClient todoist.Client, cfg pkg.Config, workItem workitemtracking.WorkItemReference) {
	log.Printf("WorkItem: %v", workItem)
	wk, err := azureClient.GetWorkItem(ctx, cfg.AzureProjectID, workItem.Id)
	if err != nil {
		log.Fatal(err)
	}

	task, err := todoistClient.AddTask(todoist.Task{
		Content:   strconv.Itoa(*wk.Id) + "-" + (*wk.Fields)["System.Title"].(string) + " " + (*wk.Fields)["System.State"].(string),
		ProjectID: cfg.TodoIstProjectID,
	})
	if err != nil {
		log.Fatalf("Error creating task: %s", err)
	}
	log.Printf("Task: %v", task)
}

func ProcessCompletedTasks(ctx context.Context, azureClient azure.Client, todoistClient todoist.Client, cfg pkg.Config) {
	response, err := todoistClient.GetCompletedTasks(ctx, cfg.TodoIstProjectID, limit, 0, "", "", false, false)
	if err != nil {
		log.Fatalf("Error getting completed tasks: %s", err)
	}
	log.Printf("Completed Tasks: %v", response)

	for _, item := range response.Items {
		var id = strings.Split(item.Content, "-")[0]
		var idInt int
		idInt, err = strconv.Atoi(id)
		if err != nil {
			log.Fatalf("Error converting id to int: %s", err)
		}
		log.Printf("Item: %v", strings.Split(item.Content, "-")[0])
		err = azureClient.UpdateWorkItemState(ctx, cfg.AzureProjectID, idInt, "Ready")
		if err != nil {
			log.Fatalf("Error updating work item: %s", err)
		}
	}
}

func main() {
	// parse
	cfg := loadConfig()

	ctx := context.Background()

	client := todoist.NewClient(cfg.TodoIstAPIKey)

	workclient, err := azure.NewWorkItemClient(ctx, cfg.AzureOrgURL, cfg.AzurePersonalAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := workclient.QueryWorkItems(ctx, cfg.AzureProjectID, cfg.AzureQueryID)
	if err != nil {
		log.Fatal(err)
	}

	for _, workItem := range *resp.WorkItems {
		CreateNewWorkItem(ctx, workclient, *client, cfg, workItem)
	}

	ProcessCompletedTasks(ctx, workclient, *client, cfg)
}
