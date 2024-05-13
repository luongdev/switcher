package activities

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/luongdev/switcher/workflow/enums"
	"github.com/luongdev/switcher/workflow/types"
	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type HttpActivityInput struct {
	Url     string        `json:"url"`
	Method  string        `json:"method"`
	Headers types.Map     `json:"headers"`
	Timeout time.Duration `json:"timeout"`
	Body    types.Map     `json:"body"`
}

func (i *HttpActivityInput) DefaultAndValidate() (err error) {
	if i.Url == "" {
		err = fmt.Errorf("url is required")
		return
	}
	if i.Method == "" {
		i.Method = http.MethodGet
	}

	if i.Method != http.MethodGet && i.Method != http.MethodPost &&
		i.Method != http.MethodPut && i.Method != http.MethodDelete {
		err = fmt.Errorf("method is invalid. Must be one of GET, POST, PUT, DELETE")
		return
	}

	if i.Headers == nil {
		i.Headers = make(types.Map)
		i.Headers["Content-Type"] = "application/json"
	}

	if i.Body == nil {
		i.Body = make(types.Map)
	}

	return
}

func (i *HttpActivityInput) Request() (*http.Request, error) {
	if err := i.DefaultAndValidate(); err != nil {
		return nil, err
	}

	b, err := i.Body.Bytes()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(i.Method, i.Url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	for k, v := range i.Headers {
		req.Header.Set(k, fmt.Sprintf("%v", v))
	}

	return req, nil
}

type HttpActivityOutput struct {
	StatusCode int `json:"statusCode"`
	Body       types.Map
}

type HttpActivity struct {
	logger *zap.Logger
	input  *HttpActivityInput
}

func (h *HttpActivity) HandlerFunc() types.ActivityFunc {
	return func(ctx context.Context, i *types.ActivityInput) (o *types.ActivityOutput, err error) {
		h.logger = activity.GetLogger(ctx)
		if err = i.Convert(&h.input); err != nil {
			return
		}

		return h.execute()
	}
}

func (h *HttpActivity) execute() (o *types.ActivityOutput, err error) {
	req, err := h.input.Request()
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: h.input.Timeout}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	o = &types.ActivityOutput{Success: false, Metadata: make(map[enums.Field]interface{})}
	ao := &HttpActivityOutput{StatusCode: res.StatusCode}

	err = json.NewDecoder(res.Body).Decode(&ao.Body)
	if err != nil {
		return nil, err
	}

	o.Success = res.StatusCode >= 200 && res.StatusCode < 300
	o.Metadata[enums.FieldOutput] = ao

	h.logger.Info("HttpActivity completed",
		zap.Any("input", h.input),
		zap.Any("output", ao),
		zap.Any("bytesRead", ao.Body))

	defer h.responseCloser(res)

	return
}

func (h *HttpActivity) responseCloser(res *http.Response) {
	err := res.Body.Close()
	if err != nil {
		h.logger.Error("Failed to close response body", zap.Error(err))
	}
}

var _ types.Activity = (*HttpActivity)(nil)
