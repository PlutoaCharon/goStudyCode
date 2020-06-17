package search

import (
	"encoding/json"
	"os"
)

const dataFile = "golangExercise\\chapter2\\sample\\data\\data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 当函数返回时 关闭文件
	defer file.Close()

	// 将文件解码到一个切片里
	// 这个切片的每一项都是指向Feed类型的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
