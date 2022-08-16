package coreapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// DataItem represents data item
type DataItem struct {
	Timestamp float64
	Value     float64
}

// DataSeries represents data series
type DataSeries struct {
	Color      string
	LineColor  string
	PointColor string
	Data       []DataItem
}

type ModuleChart struct {
	rf requestFunc
}

func (c ModuleChart) Render(title string, series []DataSeries) ([]byte, error) {
	req := struct {
		Title  string       `json:"title"`
		Series []DataSeries `json:"series"`
	}{
		Title:  title,
		Series: series,
	}

	payload, errMarshal := json.Marshal(req)
	if errMarshal != nil {
		return nil, fmt.Errorf("request marshal error, %w", errMarshal)
	}

	resp, err := c.rf("chart/render", "application/json", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to call chart/render: %w", err)
	}

	var s string

	errUnmarshal := json.Unmarshal(resp, &s)
	if errUnmarshal != nil {
		return nil, fmt.Errorf("unmarshal response error, %w", errUnmarshal)
	}

	img, errDecode := base64.StdEncoding.DecodeString(s)
	if errDecode != nil {
		return nil, fmt.Errorf("decode response error, %w", errDecode)
	}

	return img, nil
}
