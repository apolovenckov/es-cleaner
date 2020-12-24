package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func getAllIndicesK8s(ElasticURL string, elasticIndicesK8s string, elasticAPIFormat string, elasticIndexList string) Indices {

	values := url.Values{
		"format": []string{elasticAPIFormat},
		"h":      []string{elasticIndexList},
	}

	AllIndicesURL := ElasticURL + "_cat/indices/" + elasticIndicesK8s + "?"
	resp, err := http.Get(AllIndicesURL + values.Encode())
	if err != nil {
		logging("error", "Func getAllIndicesK8s, http.Get: "+fmt.Sprint(err))
	}

	dec := json.NewDecoder(resp.Body)

	defer resp.Body.Close()

	var list Indices
	err = dec.Decode(&list)
	if err != nil {
		logging("error", "Func getAllIndicesK8s, dec.Decode: "+fmt.Sprint(err))
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].CreationDate > list[j].CreationDate
	})

	return list
}

func getTaskStatus(Task string, ElasticURL string) bool {

	ElasticTaskStatusURL := ElasticURL + "_tasks/" + Task

	resp, err := http.Get(ElasticTaskStatusURL)
	if err != nil {
		logging("error", "Func getTaskStatus, http.Get: "+fmt.Sprint(err))
	}

	defer resp.Body.Close()

	if strings.Contains(resp.Status, "200") {
		dec := json.NewDecoder(resp.Body)

		var status TaskStatus
		err = dec.Decode(&status)
		if err != nil {
			logging("error", "Func getTaskStatus, dec.Decode: "+fmt.Sprint(err))
		}

		Status := status.Status

		return Status
	} else if strings.Contains(resp.Status, "404 Not Found") {
		logging("warning", "Task: "+Task+". Does not exist")
		return true
	} else {
		fmt.Println("Response Status : ", resp.Status)
		logging("error", "Task: "+Task+". Undefied http method"+resp.Status)
		return false
	}
}

func deleteComplitedTask(Task string, ElasticURL string) {
	ElasticDeleteTaskURL := ElasticURL + ".tasks/_doc/" + Task

	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", ElasticDeleteTaskURL, nil)
	if err != nil {
		logging("error", "Func deleteComplitedTask, http.NewRequest: "+fmt.Sprint(err))
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		logging("error", "Func deleteComplitedTask, client.Do: "+fmt.Sprint(err))
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	var Response DeleteComplitedTask
	err = dec.Decode(&Response)
	if err != nil {
		logging("error", "Func deleteComplitedTask, dec.Decode: "+fmt.Sprint(err))
	}

	result := Response.Result
	logging("info", "Task: "+Task+" "+result)
}

func deleteRequests(ElasticURLdel string, body []uint8) string {
	resp, err := http.Post(ElasticURLdel, "application/json", bytes.NewReader(body))
	if err != nil {
		logging("error", "Func deleteRequests, http.Post: "+fmt.Sprint(err))
	}

	dec := json.NewDecoder(resp.Body)

	var task Task

	err = dec.Decode(&task)
	if err != nil {
		logging("error", "Func deleteRequests, dec.Decode: "+fmt.Sprint(err))
	}
	taskID := task.Task
	return taskID
}
