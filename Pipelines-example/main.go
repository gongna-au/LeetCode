package main

import (
	"fmt"

	"github.com/LeetCode/Pipelines-example/business"
	"github.com/LeetCode/Pipelines-example/pipeLine/aggregate"
	"github.com/LeetCode/Pipelines-example/pipeLine/filter"
	"github.com/LeetCode/Pipelines-example/pipeLine/mapFrom"
	"github.com/LeetCode/Pipelines-example/pipeLine/stream"
)

const maxUserID = 100

func main() {
	done := make(chan interface{})
	defer close(done)

	userIDs := make([]business.UserID, maxUserID)
	for i := 1; i <= maxUserID; i++ {
		userIDs = append(userIDs, business.UserID(i))
	}

	plainStructs := map_from.UPAggToPlainStruct(done,
		aggregate.Profile(done,
			filter.InactiveUsers(done,
				aggregate.User(done,
					stream.UserIDs(done, userIDs...)))),
	)

	for ps := range plainStructs {
		fmt.Printf("[result] plain struct for UserID %v is: -> %v \n", ps.UserID, ps)
	}
}
