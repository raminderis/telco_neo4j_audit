package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var errorCount, passCount int
var wg sync.WaitGroup

func main() {
	errorCount = 0
	passCount = 0

	// Building a session towards neo4j instance. ensure URL, username and password are correct.
	// In future this authentication will be secured instead of plain text.
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
		fmt.Println(err)
	}
	defer driver.Close(ctx)
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	wg.Add(1)
	go performCountAudit(session, ctx)
	wg.Wait()
	fmt.Printf("AUDIT COMPLETE. Total of %v Errors found and Total of %v valid entries found \n", errorCount, passCount)
}
