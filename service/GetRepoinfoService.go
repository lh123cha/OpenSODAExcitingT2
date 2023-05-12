package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type RepoInfo struct {
	Name            string `json:"repo.name"`
	Url             string `json:"repo.url"`
	month           string
	openrank        string `json:"repo.index.xlab.openrank"`
	activity        string `json:"repo.index.xlab.activity"`
	dates_and_times string `json:"repo.metric.chaoss.active dates and times"`
}

/*
*
Get_On_certain_repo
*/
func Get_On_certain_repo(repo string, metric string) ([]byte, []byte) {

	Base_url := "https://oss.x-lab.info/open_digger/github/"
	url := Base_url + repo + "/" + strings.ToLower(metric) + ".json"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	repo_name := strings.Split(repo, "/")[1]
	repo_url := "https://github.com/" + repo
	repo_info := map[string]string{
		"repo.name": repo_name,
		"repo.url":  repo_url,
		metric:      string(body),
	}
	bytes, _ := json.Marshal(repo_info)
	fmt.Println(string(bytes))

	return bytes, body
}

func Get_On_certain_month(repo string, metric string, month string) []byte {

	jsonData, body := Get_On_certain_repo(repo, metric)
	var v1 interface{}
	var v2 interface{}

	json.Unmarshal(jsonData, &v1)
	json.Unmarshal(body, &v2)
	data1 := v1.(map[string]interface{})
	data2 := v2.(map[string]interface{})
	repo_info := map[string]string{}
	for k, v := range data2 {
		if k == month {
			repo_info = map[string]string{
				"repo.name": data1["repo.name"].(string),
				"repo.url":  data1["repo.url"].(string),
				"month":     month,
				metric:      strconv.FormatFloat(v.(float64), 'f', 2, 32),
			}
		}
	}
	bytes, _ := json.Marshal(repo_info)
	fmt.Println(string(bytes))
	return bytes

}