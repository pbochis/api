package logic

import (
	"time"

	"github.com/coduno/api/model"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

func init() {
	RegisterResulter(Average, averageResulter)
}

func averageResulter(ctx context.Context, resultKey *datastore.Key) error {
	var result model.Result
	if err := datastore.Get(ctx, resultKey, &result); err != nil {
		return err
	}

	var challenge model.Challenge
	if err := datastore.Get(ctx, result.Challenge, &challenge); err != nil {
		return err
	}
	tasks := make([]model.Task, len(challenge.Tasks))
	if err := datastore.GetMulti(ctx, challenge.Tasks, tasks); err != nil {
		return err
	}

	average := model.Skills{}

	var nrOfComputations float64

	for i, task := range tasks {
		taskResult, err := Tasker(task.Tasker).Call(ctx, challenge.Tasks[i], resultKey, resultKey.Parent().Parent())
		if err != nil {
			// TODO: ignore error for now. We`ll treat it after we have all the taskers available
			//return err
		} else {
			average = average.Add(taskResult)
			nrOfComputations++
		}
	}

	result.Skills = average.DivBy(nrOfComputations)
	result.Computed = time.Now()

	_, err := result.Put(ctx, resultKey)
	return err
}
