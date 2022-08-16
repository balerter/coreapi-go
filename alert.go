package coreapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type AlertOptions struct {
	Channels []string          `json:"channels"`
	Quiet    bool              `json:"quiet"`
	Repeat   int               `json:"repeat"`
	Image    string            `json:"image"`
	Fields   map[string]string `json:"fields"`
	Escalate map[int][]string  `json:"escalate"`
}

type Alert struct {
	Name       string    `json:"name"`
	Level      int       `json:"level"`
	LastChange time.Time `json:"last_change"`
	Start      time.Time `json:"start"`
	Count      int       `json:"count"`
}

type ModuleAlert struct {
	rf requestFunc
}

// Success calls alert module with success level.
func (m ModuleAlert) Success(alertName, message string, opts *AlertOptions) (*Alert, bool, error) {
	return m.call("success", alertName, message, opts)
}

// Error calls alert module with error level.
func (m ModuleAlert) Error(alertName, message string, opts *AlertOptions) (*Alert, bool, error) {
	return m.call("error", alertName, message, opts)
}

// Warning calls alert module with warning level.
func (m ModuleAlert) Warning(alertName, message string, opts *AlertOptions) (*Alert, bool, error) {
	return m.call("warn", alertName, message, opts)
}

func (m ModuleAlert) call(method, alertName, message string, opts *AlertOptions) (*Alert, bool, error) {
	u := fmt.Sprintf("alert/%s/%s", method, alertName)

	if opts != nil {
		args := url.Values{}
		if len(opts.Channels) > 0 {
			args.Add("channels", strings.Join(opts.Channels, ","))
		}
		if opts.Quiet {
			args.Add("quiet", "true")
		}
		if opts.Repeat > 0 {
			args.Add("repeat", fmt.Sprintf("%d", opts.Repeat))
		}
		if opts.Image != "" {
			args.Add("image", opts.Image)
		}
		if len(opts.Fields) > 0 {
			var fields []string
			for k, v := range opts.Fields {
				fields = append(fields, fmt.Sprintf("%s:%s", k, v))
			}
			args.Add("fields", strings.Join(fields, ","))
		}
		if len(opts.Escalate) > 0 {
			var escalate []string
			for k, v := range opts.Escalate {
				escalate = append(escalate, fmt.Sprintf("%d:%s", k, strings.Join(v, ",")))
			}
			args.Add("escalate", strings.Join(escalate, ";"))
		}

		if len(args) > 0 {
			u = fmt.Sprintf("%s?%s", u, args.Encode())
		}
	}

	resp, err := m.rf(u, "text/plain", []byte(message))
	if err != nil {
		return nil, false, fmt.Errorf("failed to call %s: %w", u, err)
	}

	rsp := struct {
		Alert           *Alert `json:"alert"`
		LevelWasUpdated bool   `json:"level_was_updated"`
	}{}

	errUnmarshal := json.Unmarshal(resp, &rsp)
	if errUnmarshal != nil {
		return nil, false, fmt.Errorf("failed to unmarshal response: %w", errUnmarshal)
	}

	return rsp.Alert, rsp.LevelWasUpdated, nil
}

func (m ModuleAlert) Get(alertName string) (*Alert, error) {
	u := fmt.Sprintf("alert/get/%s", alertName)

	resp, err := m.rf(u, "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call %s: %w", u, err)
	}

	var rsp *Alert

	errUnmarshal := json.Unmarshal(resp, &rsp)
	if errUnmarshal != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", errUnmarshal)
	}

	return rsp, nil
}
