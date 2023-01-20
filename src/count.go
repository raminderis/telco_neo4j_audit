package main

import (
	"context"
	"fmt"

	"github.com/benpate/convert"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func performCountAudit(session neo4j.SessionWithContext, ctx context.Context) {
	defer wg.Done()
	//Node count audit
	expectedNodeEntitiesCount := expectedNodeEntitiesCount()
	currentNodeEntitiesCount := currentNodeEntitiesCount(session, ctx)
	if expectedNodeEntitiesCount != currentNodeEntitiesCount {
		fmt.Printf("ERROR : Node Count of %v Doesnt Match the expected count of %v! \n", currentNodeEntitiesCount, expectedNodeEntitiesCount)
		errorCount++
	} else {
		fmt.Printf("PASS : Node Count of %v Matches the expected count of %v! \n", currentNodeEntitiesCount, expectedNodeEntitiesCount)
		passCount++
	}

	//location count audit
	expectedLocationEntitiesCount := expectedLocationEntitiesCount()
	currentLocationEntitiesCount := currentLocationEntitiesCount(session, ctx)
	if expectedLocationEntitiesCount != currentLocationEntitiesCount {
		fmt.Printf("ERROR : Location Count of %v Doesnt Match the expected count of %v! \n", currentLocationEntitiesCount, expectedLocationEntitiesCount)
		errorCount++
	} else {
		fmt.Printf("PASS : Location Count of %v Matches the expected count of %v! \n", currentLocationEntitiesCount, expectedLocationEntitiesCount)
		passCount++
	}

	//testagent count audit
	expectedTestAgentEntitiesCount := expectedTestAgentEntitiesCount()
	currentTestAgentEntitiesCount := currentTestAgentEntitiesCount(session, ctx)
	if expectedTestAgentEntitiesCount != currentTestAgentEntitiesCount {
		fmt.Printf("ERROR : TestAgent Count of %v Doesnt Match the expected count of %v! \n", currentTestAgentEntitiesCount, expectedTestAgentEntitiesCount)
		errorCount++
	} else {
		fmt.Printf("PASS : TestAgent Count of %v Matches the expected count of %v! \n", currentTestAgentEntitiesCount, expectedTestAgentEntitiesCount)
		passCount++
	}

	//testagenthost count audit
	expectedTestAgentHostEntitiesCount := expectedTestAgentHostEntitiesCount()
	currentTestAgentHostEntitiesCount := currentTestAgentHostEntitiesCount(session, ctx)
	if expectedTestAgentHostEntitiesCount != currentTestAgentHostEntitiesCount {
		fmt.Printf("ERROR : TestAgentHost Count of %v Doesnt Match the expected count of %v! \n", currentTestAgentHostEntitiesCount, expectedTestAgentHostEntitiesCount)
		errorCount++
	} else {
		fmt.Printf("PASS : TestAgentHost Count of %v Matches the expected count of %v! \n", currentTestAgentHostEntitiesCount, expectedTestAgentHostEntitiesCount)
		passCount++
	}
}

func expectedNodeEntitiesCount() int {
	return 50
}

func expectedLocationEntitiesCount() int {
	return 18
}

func expectedTestAgentEntitiesCount() int {
	return 20
}

func expectedTestAgentHostEntitiesCount() int {
	return 14
}

func currentNodeEntitiesCount(session neo4j.SessionWithContext, ctx context.Context) int {
	nodeCount := 0
	result, err := session.Run(ctx, "MATCH (n:NODE) RETURN COUNT(n) AS count", nil)
	if err != nil {
		fmt.Println(err)
	}
	record, err := result.Single(ctx)
	//fmt.Println("key check: ", record.Keys)
	//fmt.Println("value check: ", record.Values)
	/*****************************START*******************************/
	/* you can either do this or directly return values like what i am doing*/
	/*
		key := record.Keys[0]
		value, flag := record.Get(key)
		if !flag {
			fmt.Println("no values received return 0")
			nodeCount = 0
		} else {
			nodeCount = convert.Int(value)
		}
		fmt.Println("node count is: ", nodeCount)
		return nodeCount
	*/
	/*****************************END*******************************/
	nodeCount = convert.Int(record.Values[0])
	return nodeCount
}

func currentLocationEntitiesCount(session neo4j.SessionWithContext, ctx context.Context) int {
	locationCount := 0
	result, err := session.Run(ctx, "MATCH (n:LOCATION) RETURN COUNT(n) AS count", nil)
	if err != nil {
		fmt.Println(err)
	}
	record, err := result.Single(ctx)
	locationCount = convert.Int(record.Values[0])
	return locationCount
}

func currentTestAgentEntitiesCount(session neo4j.SessionWithContext, ctx context.Context) int {
	testagentCount := 0
	result, err := session.Run(ctx, "MATCH (n:TESTAGENT) RETURN COUNT(n) AS count", nil)
	if err != nil {
		fmt.Println(err)
	}
	record, err := result.Single(ctx)
	testagentCount = convert.Int(record.Values[0])
	return testagentCount
}

func currentTestAgentHostEntitiesCount(session neo4j.SessionWithContext, ctx context.Context) int {
	testagenthostCount := 0
	result, err := session.Run(ctx, "MATCH (n:TESTAGENTHOST) RETURN COUNT(n) AS count", nil)
	if err != nil {
		fmt.Println(err)
	}
	record, err := result.Single(ctx)
	testagenthostCount = convert.Int(record.Values[0])
	return testagenthostCount
}
