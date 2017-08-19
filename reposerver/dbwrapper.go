package main

import (
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

type Task struct {
	Description string
}

func Tmp() {

	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := "gorepo"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the kind for the new entity.
	kind := "Task"
	// Sets the name/ID for the new entity.
	name := "sampletask1"
	// Creates a Key instance.
	taskKey := datastore.NameKey(kind, name, nil)

	// Creates a Task instance.
	task := Task{
		Description: "Buy milk",
	}

	// Saves the new entity.
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

	fmt.Printf("Saved %v: %v\n", taskKey, task.Description)

	//	ctx := context.Background()
	//	client, _ := datastore.NewClient(ctx, "my-proj")
	//	taskKey := datastore.NameKey("Task", "sampleTask", nil)
	// [START lookup]
	var task2 Task
	err = client.Get(ctx, taskKey, &task2)
	// [END lookup]
	_ = err // Make sure you check err.

	fmt.Printf("Got %v: %v\n", taskKey, task2.Description)

	client.Close()
}
