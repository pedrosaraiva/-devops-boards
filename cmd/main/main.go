package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/caarlos0/env"
	"github.com/pedrosaraiva1/devops-boards/internal/azure"
	"github.com/pedrosaraiva1/devops-boards/internal/todoist"
	"github.com/pedrosaraiva1/devops-boards/pkg"
)

func main() {
	// parse
	var cfg pkg.Config
	err := env.Parse(&cfg)

	ctx := context.Background()

	workclient, err := azure.NewWorkItemClient(ctx, cfg.AZURE_ORGANIZATION_URL, cfg.AZURE_PERSONAL_ACCESS_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := workclient.QueryWorkItems(ctx, cfg.AZURE_PROJECT, cfg.AZURE_QUERY_ID)
	if err != nil {
		log.Fatal(err)
	}

	client := todoist.NewClient(cfg.TODOIST_API_KEY)

	for _, workItem := range *resp.WorkItems {
		log.Printf("WorkItem: %v", workItem)
		wk, err := workclient.GetWorkItem(ctx, cfg.AZURE_PROJECT, workItem.Id)
		if err != nil {
			log.Fatal(err)
		}

		task, err := client.AddTask(todoist.Task{
			Content:   strconv.Itoa(*wk.Id) + "-" + (*wk.Fields)["System.Title"].(string) + " " + (*wk.Fields)["System.State"].(string),
			ProjectID: cfg.TODOIST_PROJECT_ID,
		})
		if err != nil {
			log.Fatalf("Error creating task: %s", err)
		}

		response, err := client.GetCompletedTasks(ctx, cfg.TODOIST_PROJECT_ID, 30, 0, "", "", false, false)
		if err != nil {
			log.Fatalf("Error getting completed tasks: %s", err)
		}
		log.Printf("Completed Tasks: %v", response)
		log.Print(task)
		log.Printf("WorkItem: %v", wk)

		for _, item := range response.Items {
			var id = strings.Split(item.Content, "-")[0]
			idInt, err := strconv.Atoi(id)
			if err != nil {
				log.Fatalf("Error converting id to int: %s", err)
			}
			log.Printf("Item: %v", strings.Split(item.Content, "-")[0])
			err = workclient.UpdateWorkItemState(ctx, cfg.AZURE_PROJECT, idInt, "Ready")
			if err != nil {
				log.Fatalf("Error updating work item: %s", err)
			}
		}
	}
}
