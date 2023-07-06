package service

import (
	"encoding/json"
	"fmt"
)

var printList = []string{
	"repo.name",
	"repo.url",
}

func PrintRepoInfo(source_ RepoInfo) {
	fmt.Printf("%s: %s\n", printList[0], source_.RepoName)
	fmt.Printf("%s: %s\n", printList[1], source_.RepoUrl)

	for _, v := range Metrics {
		datum, ok := source_.Data[v]
		if ok {
			fmt.Printf("%s:", v)
			jsonData, err := json.Marshal(datum)
			if err != nil {
				//fmt.Errorf("trans fail", err)
			}
			fmt.Println(string(jsonData))
		}
	}

}