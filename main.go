package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func logging(loglevel string, msg string) {
	dt := time.Now().Format("2006-01-02T15:04:05Z")
	println("{\"level\":\"" + loglevel + "\",\"msg\":\"" + msg + "\",\"time\":\"" + dt + "\"}")
}

func deleteProcessing(list Indices, ParametersToDelete ParamsToDeleteStruct, ElasticURL string) {

	for _, params := range ParametersToDelete {
		Lifetime := params.Lifetime
		ContainerName := params.ContainerName
		Message := params.Message
		indexList := list[(Lifetime - 1):]

		logging("info", "Start delete for: "+ContainerName+". With message: "+Message)
		logging("info", "Need deleted in indexes: "+fmt.Sprint(indexList))

		for _, index := range indexList {
			logging("info", "Start delete for: "+ContainerName+". With message: "+Message+". Index: "+index.Index)
			if Message == "*" {
				indexName := index.Index
				ElasticURLdel := ElasticURL + indexName + "/_delete_by_query?wait_for_completion=false"

				var bodyForPost DeleteContainerNameLog
				bodyForPost.Query.Term.KubContName = ContainerName

				body, err := json.Marshal(bodyForPost)
				if err != nil {
					logging("error", "Func deleteProcessing, json.Marshal: "+fmt.Sprint(err))
				}

				taskID := deleteRequests(ElasticURLdel, body)
				logging("info", "Task: "+taskID+" started")
				for true {
					if getTaskStatus(taskID, ElasticURL) {
						deleteComplitedTask(taskID, ElasticURL)
						break
					}
					logging("info", "Task: "+taskID+" not ended. Wait...")
					time.Sleep(10 * time.Second)
				}

			} else {
				indexName := index.Index
				ElasticURLdel := ElasticURL + indexName + "/_delete_by_query?wait_for_completion=false"

				var bodyForPost DeleteMessageLog
				bodyForPost.Query.Bool.Filter.Term.ContainerName = ContainerName
				bodyForPost.Query.Bool.Must.Match.MessageLog = Message

				body, err := json.Marshal(bodyForPost)
				if err != nil {
					log.Fatal("Func getTaskToDelete, json.Marshal:", err)
				}

				taskID := deleteRequests(ElasticURLdel, body)
				logging("info", "Task: "+taskID+". started")
				for true {
					if getTaskStatus(taskID, ElasticURL) {
						deleteComplitedTask(taskID, ElasticURL)
						break
					}
					time.Sleep(10 * time.Second)
				}

			}
		}
	}
}

func main() {
	logging("info", "Service started")

	ElasticURL, elasticIndicesK8s, elasticAPIFormat, elasticIndexList, ParamsToDelete := getElasticEnv()

	listIndexes := getAllIndicesK8s(ElasticURL, elasticIndicesK8s, elasticAPIFormat, elasticIndexList)

	deleteProcessing(listIndexes, ParamsToDelete, ElasticURL)

}
