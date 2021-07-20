package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kazu1029/dynamodb-go/pkg/dynamo"
	"github.com/kazu1029/dynamodb-go/switchtoggle/toggle"
)

func main() {
	ctx := context.Background()
	tableName := "ToggleStateTable"
	db, cleanup := dynamo.SetupTable(ctx, tableName, "./switchtoggle/template.yml")
	defer cleanup()

	now := time.Now()
	togl := toggle.NewToggle(db, tableName)
	err := togl.Save(ctx, toggle.Switch{ID: "123", State: true, CreatedAt: now})
	if err != nil {
		panic(err)
	}

	fmt.Printf("========================== Start Saving toggle ===========================\n")
	s, err := togl.Latest(ctx, "123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", s)
	fmt.Printf("========================== End Saving toggle ===========================\n\n")

	fmt.Printf("========================== Start Latest ===========================\n")
	err = togl.Save(ctx, toggle.Switch{ID: "123", State: false, CreatedAt: now.Add(10 * time.Second)})
	if err != nil {
		panic(err)
	}

	s, err = togl.Latest(ctx, "123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", s)
	fmt.Printf("========================== End Latest ===========================\n\n")

	fmt.Printf("========================== Start Dropping out of order switch ===========================\n")
	err = togl.Save(ctx, toggle.Switch{ID: "123", State: false, CreatedAt: now.Add(-10 * time.Second)})
	if err != nil {
		panic(err)
	}
	s, err = togl.Latest(ctx, "123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", s)
	fmt.Printf("========================== End Dropping out of order switch ===========================\n\n")
}
