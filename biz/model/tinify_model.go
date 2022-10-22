package model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ShrinkResp struct {
	Input  Input  `json:"input"`
	Output Output `json:"output"`
}

type Input struct {
	Size int    `json:"size"`
	Type string `json:"type"`
}

type Output struct {
	Size   int     `json:"size"`
	Type   string  `json:"type"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Ratio  float64 `json:"ratio"`
	Url    string  `json:"url"`
}

func (s *ShrinkResp) String() string {
	b, err := json.Marshal(*s)
	if err != nil {
		return fmt.Sprintf("%+v", *s)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *s)
	}
	return out.String()
}
