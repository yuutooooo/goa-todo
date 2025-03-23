// Code generated by goa v3.20.0, DO NOT EDIT.
//
// todo HTTP client encoders and decoders
//
// Command:
// $ goa gen backend/design

package client

import (
	todo "backend/gen/todo"
	todoviews "backend/gen/todo/views"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "todo" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		userID int
	)
	{
		p, ok := v.(*todo.CreatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todo", "create", "*todo.CreatePayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateTodoPath(userID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the todo create
// server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*todo.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("todo", "create", "*todo.CreatePayload", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("todo", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the todo
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeCreateResponse may return the following errors:
//   - "BadRequest" (type *todo.APIErrorResult): http.StatusBadRequest
//   - error: internal error
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "create", err)
			}
			err = ValidateCreateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("todo", "create", err)
			}
			res := NewCreateTodoResultCreated(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "create", err)
			}
			err = ValidateCreateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("todo", "create", err)
			}
			return nil, NewCreateBadRequest(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "todo" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		userID int
	)
	{
		p, ok := v.(*todo.ListPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todo", "list", "*todo.ListPayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListTodoPath(userID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the todo list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*todo.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("todo", "list", "*todo.ListPayload", v)
		}
		values := req.URL.Query()
		values.Add("completed", fmt.Sprintf("%v", p.Completed))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the todo list
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "list", err)
			}
			p := NewListTodoCollectionOK(&body)
			view := "default"
			vres := &todoviews.TodoCollection{Projected: p, View: view}
			if err = todoviews.ValidateTodoCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("todo", "list", err)
			}
			res := todo.NewTodoCollection(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "todo" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		userID int
		todoID int
	)
	{
		p, ok := v.(*todo.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todo", "get", "*todo.GetPayload", v)
		}
		userID = p.UserID
		todoID = p.TodoID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetTodoPath(userID, todoID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetResponse returns a decoder for responses returned by the todo get
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeGetResponse may return the following errors:
//   - "NotFound" (type *todo.APIErrorResult): http.StatusNotFound
//   - error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "get", err)
			}
			err = ValidateGetResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("todo", "get", err)
			}
			res := NewGetTodoResultOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "get", err)
			}
			err = ValidateGetNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("todo", "get", err)
			}
			return nil, NewGetNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "get", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "todo" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		userID int
		todoID int
	)
	{
		p, ok := v.(*todo.UpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todo", "update", "*todo.UpdatePayload", v)
		}
		userID = p.UserID
		todoID = p.TodoID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateTodoPath(userID, todoID)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the todo update
// server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*todo.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("todo", "update", "*todo.UpdatePayload", v)
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("todo", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the todo
// update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeUpdateResponse may return the following errors:
//   - "NotFound" (type *todo.APIErrorResult): http.StatusNotFound
//   - "BadRequest" (type *todo.APIErrorResult): http.StatusBadRequest
//   - error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "update", err)
			}
			err = ValidateUpdateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("todo", "update", err)
			}
			res := NewUpdateTodoResultOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body UpdateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "update", err)
			}
			err = ValidateUpdateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("todo", "update", err)
			}
			return nil, NewUpdateNotFound(&body)
		case http.StatusBadRequest:
			var (
				body UpdateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "update", err)
			}
			err = ValidateUpdateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("todo", "update", err)
			}
			return nil, NewUpdateBadRequest(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "todo" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		userID int
		todoID int
	)
	{
		p, ok := v.(*todo.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todo", "delete", "*todo.DeletePayload", v)
		}
		userID = p.UserID
		todoID = p.TodoID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteTodoPath(userID, todoID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteResponse returns a decoder for responses returned by the todo
// delete endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDeleteResponse may return the following errors:
//   - "NotFound" (type *todo.APIErrorResult): http.StatusNotFound
//   - error: internal error
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusNotFound:
			var (
				body DeleteNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "delete", err)
			}
			err = ValidateDeleteNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("todo", "delete", err)
			}
			return nil, NewDeleteNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalTodoResultResponseBodyToTodoviewsTodoResultView builds a value of
// type *todoviews.TodoResultView from a value of type *TodoResultResponseBody.
func unmarshalTodoResultResponseBodyToTodoviewsTodoResultView(v *TodoResultResponseBody) *todoviews.TodoResultView {
	if v == nil {
		return nil
	}
	res := &todoviews.TodoResultView{
		ID:          v.ID,
		UserID:      v.UserID,
		Title:       v.Title,
		Description: v.Description,
		Completed:   v.Completed,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}

	return res
}
