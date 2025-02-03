package jsondata

import (
	_ "embed"
	"encoding/json"
	"log"
)

type JsonData struct {
	VideoNodes []VideoNode
	// Url        string "static/transforming_meta.json"
}

func NewJsonData() *JsonData {
	jd := &JsonData{}
	jd.VideoNodes = jd.loadData()
	return jd
}

type VideoNode struct {
	Module      string `json:"module"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Wistia_link string `json:"wistia_link"`
}

//go:embed transforming_meta.json
var content []byte

func (jd *JsonData) loadData() []VideoNode {
	// content, err := ioutil.ReadFile(jd.Url)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var VideoNodes []VideoNode
	if err := json.Unmarshal(content, &VideoNodes); err != nil {
		log.Fatal(err)
	}

	return VideoNodes
}
