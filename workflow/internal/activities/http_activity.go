package activities

import (
	"bytes"
	"context"
	"fmt"
	"github.com/luongdev/switcher/workflow/enums"
	"github.com/luongdev/switcher/workflow/types"
	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type HttpActivityInput struct {
	types.ActivityInput

	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    []byte            `json:"body"`
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
		i.Headers = make(map[string]string)
		i.Headers["Content-Type"] = "application/json"
	}
	if i.Body == nil {
		i.Body = []byte{}
	}

	return
}

func (i *HttpActivityInput) Request() (*http.Request, error) {
	if err := i.DefaultAndValidate(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(i.Method, i.Url, bytes.NewBuffer(i.Body))
	if err != nil {
		return nil, err
	}

	for k, v := range i.Headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

type HttpActivityOutput struct {
	StatusCode int `json:"statusCode"`
	Body       []byte
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

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	o = &types.ActivityOutput{Success: false, Metadata: make(map[enums.Field]interface{})}
	ao := &HttpActivityOutput{StatusCode: res.StatusCode, Body: []byte{}}

	ao.Body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	o.Success = res.StatusCode >= 200 && res.StatusCode < 300
	o.Metadata[enums.FieldOutput] = ao

	h.logger.Info("HttpActivity completed",
		zap.Any("input", h.input),
		zap.Any("output", ao),
		zap.Any("bytesRead", string(ao.Body)))

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
