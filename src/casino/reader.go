package casino

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"poker/src/model"
)

type Reader struct {
	FilePath string
}

type JsonStruct struct {
	Alice  string
	Bob    string
	Result string
}

func (*Reader) ReadFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	matches := &model.Matches{}
	err = json.Unmarshal(data, matches)
	if err != nil {
		panic(err)
	}

	go func(m *model.Matches) {
		for _, v := range m.MatchSlice {
			fmt.Println(v)
		}
	}(matches)
}
