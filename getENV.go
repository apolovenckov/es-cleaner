package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getElasticEnv() (string, string, string, string, ParamsToDeleteStruct) {
	elasticServer := "192.168.102.253"
	elasticServerEnv, elasticServerHas := os.LookupEnv("ELASTIC_HOST")
	if elasticServerHas {
		elasticServer = elasticServerEnv
	}

	elasticPort := "9200"
	elasticPortEnv, elasticPortHas := os.LookupEnv("ELASTIC_PORT")
	if elasticPortHas {
		elasticPort = elasticPortEnv
	}

	elasticIndicesK8s := "k8s-dev-*"
	elasticIndicesK8sEnv, elasticIndicesK8sHas := os.LookupEnv("ELASTIC_INDICES_K8S")
	if elasticIndicesK8sHas {
		elasticIndicesK8s = elasticIndicesK8sEnv
	}

	elasticAPIFormat := "JSON"
	elasticAPIFormatEnv, elasticAPIFormatHas := os.LookupEnv("ELASTIC_API_FORMAT")
	if elasticAPIFormatHas {
		elasticAPIFormat = elasticAPIFormatEnv
	}

	elasticIndexList := "index,creation.date"
	elasticIndexListEnv, elasticIndexListHas := os.LookupEnv("ELASTIC_INDEX_LIST")
	if elasticIndexListHas {
		elasticIndexList = elasticIndexListEnv
	}

	ElasticURL := "http://" + elasticServer + ":" + elasticPort + "/"
	ElasticURLEnv, ElasticURLHas := os.LookupEnv("ELASTIC_URL")
	if ElasticURLHas {
		ElasticURL = ElasticURLEnv
	}

	ParamsToDelete := "[{\"ContainerName\":\"feed\",\"Lifetime\":2,\"Message\":\"*\"}]"
	ParamsToDeleteEnv, ParamsToDeleteHas := os.LookupEnv("PARAMS_TO_DELETE")
	if ParamsToDeleteHas {
		ParamsToDelete = ParamsToDeleteEnv
	}

	return ElasticURL, elasticIndicesK8s, elasticAPIFormat, elasticIndexList, parsParamsToDelete(ParamsToDelete)
}

func parsParamsToDelete(ParamsToDelete string) ParamsToDeleteStruct {
	rawParamsToDelete := json.RawMessage(ParamsToDelete)
	bytes, err := rawParamsToDelete.MarshalJSON()
	if err != nil {
		logging("error", "Func parsParamsToDelete, MarshalJSON: "+fmt.Sprint(err))
	}

	var ParametersToDelete ParamsToDeleteStruct
	err = json.Unmarshal(bytes, &ParametersToDelete)
	if err != nil {
		logging("error", "Func parsParamsToDelete, json.Unmarshal: "+fmt.Sprint(err))
	}

	return ParametersToDelete
}
