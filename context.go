package httplusplus

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	values         map[string]any
}

func (c *Context) Set(key string, val any) {
	if c.values == nil {
		c.values = make(map[string]any)
	}
	c.values[key] = val
}

func (c *Context) Get(key string) any {
	return c.values[key]
}

func (c *Context) Text(code int, s string) error {
	c.ResponseWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.ResponseWriter.WriteHeader(code)
	_, err := c.ResponseWriter.Write([]byte(s))
	return err
}

func (c *Context) Json(code int, data any) error {
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.ResponseWriter.WriteHeader(code)
	return json.NewEncoder(c.ResponseWriter).Encode(data)
}

func (c *Context) BindJson(v any) error {
	if c.Request.Body == nil {
		return errors.New("Request body is empty")
	}
	defer c.Request.Body.Close()
	if err := json.NewDecoder(c.Request.Body).Decode(v); err != nil {
		return fmt.Errorf("Decode json error: %w", err)
	}
	return nil
}

func (c *Context) SendError(status int, message string) error {
	return SendError(status, message)
}
