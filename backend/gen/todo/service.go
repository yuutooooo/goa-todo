// Code generated by goa v3.20.0, DO NOT EDIT.
//
// todo service
//
// Command:
// $ goa gen backend/design

package todo

import (
	todoviews "backend/gen/todo/views"
	"context"
)

// タスク管理サービス
type Service interface {
	// 新規タスクを作成します
	Create(context.Context, *CreatePayload) (res *TodoResult, err error)
	// タスク一覧を取得します
	List(context.Context, *ListPayload) (res *TodoCollection, err error)
	// タスク詳細を取得します
	Get(context.Context, *GetPayload) (res *TodoResult, err error)
	// タスクを更新します
	Update(context.Context, *UpdatePayload) (res *TodoResult, err error)
	// タスクを削除します
	Delete(context.Context, *DeletePayload) (err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "todo app"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "todo"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"create", "list", "get", "update", "delete"}

// エラーレスポンス
type APIErrorResult struct {
	// エラー名
	Name string
	// エラーメッセージ
	Message string
}

// CreatePayload is the payload type of the todo service create method.
type CreatePayload struct {
	// ユーザーID
	UserID int
	// タスクのタイトル
	Title string
	// タスクの説明
	Description string
	// タスクが完了しているかどうか
	Completed bool
}

// DeletePayload is the payload type of the todo service delete method.
type DeletePayload struct {
	// ユーザーID
	UserID int
	// タスクID
	TodoID int
}

// GetPayload is the payload type of the todo service get method.
type GetPayload struct {
	// ユーザーID
	UserID int
	// タスクID
	TodoID int
}

// ListPayload is the payload type of the todo service list method.
type ListPayload struct {
	// ユーザーID
	UserID int
	// 完了済みタスクのみ取得
	Completed bool
}

// TodoCollection is the result type of the todo service list method.
type TodoCollection struct {
	Items []*TodoResult
}

// TodoResult is the result type of the todo service create method.
type TodoResult struct {
	// タスクのid
	ID int
	// ユーザーID
	UserID int
	// タスクのタイトル
	Title string
	// タスクの説明
	Description string
	// タスクが完了しているかどうか
	Completed bool
	// 作成日時
	CreatedAt string
	// 更新日時
	UpdatedAt string
}

// UpdatePayload is the payload type of the todo service update method.
type UpdatePayload struct {
	// ユーザーID
	UserID int
	// タスクID
	TodoID int
	// タスクのタイトル
	Title *string
	// タスクの説明
	Description *string
	// タスクが完了しているかどうか
	Completed *bool
}

// Error returns an error description.
func (e *APIErrorResult) Error() string {
	return "エラーレスポンス"
}

// ErrorName returns "APIErrorResult".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *APIErrorResult) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "APIErrorResult".
func (e *APIErrorResult) GoaErrorName() string {
	return e.Name
}

// NewTodoCollection initializes result type TodoCollection from viewed result
// type TodoCollection.
func NewTodoCollection(vres *todoviews.TodoCollection) *TodoCollection {
	return newTodoCollection(vres.Projected)
}

// NewViewedTodoCollection initializes viewed result type TodoCollection from
// result type TodoCollection using the given view.
func NewViewedTodoCollection(res *TodoCollection, view string) *todoviews.TodoCollection {
	p := newTodoCollectionView(res)
	return &todoviews.TodoCollection{Projected: p, View: "default"}
}

// newTodoCollection converts projected type TodoCollection to service type
// TodoCollection.
func newTodoCollection(vres *todoviews.TodoCollectionView) *TodoCollection {
	res := &TodoCollection{}
	if vres.Items != nil {
		res.Items = make([]*TodoResult, len(vres.Items))
		for i, val := range vres.Items {
			res.Items[i] = transformTodoviewsTodoResultViewToTodoResult(val)
		}
	}
	return res
}

// newTodoCollectionView projects result type TodoCollection to projected type
// TodoCollectionView using the "default" view.
func newTodoCollectionView(res *TodoCollection) *todoviews.TodoCollectionView {
	vres := &todoviews.TodoCollectionView{}
	if res.Items != nil {
		vres.Items = make([]*todoviews.TodoResultView, len(res.Items))
		for i, val := range res.Items {
			vres.Items[i] = transformTodoResultToTodoviewsTodoResultView(val)
		}
	}
	return vres
}

// transformTodoviewsTodoResultViewToTodoResult builds a value of type
// *TodoResult from a value of type *todoviews.TodoResultView.
func transformTodoviewsTodoResultViewToTodoResult(v *todoviews.TodoResultView) *TodoResult {
	if v == nil {
		return nil
	}
	res := &TodoResult{
		ID:          *v.ID,
		UserID:      *v.UserID,
		Title:       *v.Title,
		Description: *v.Description,
		Completed:   *v.Completed,
		CreatedAt:   *v.CreatedAt,
		UpdatedAt:   *v.UpdatedAt,
	}

	return res
}

// transformTodoResultToTodoviewsTodoResultView builds a value of type
// *todoviews.TodoResultView from a value of type *TodoResult.
func transformTodoResultToTodoviewsTodoResultView(v *TodoResult) *todoviews.TodoResultView {
	if v == nil {
		return nil
	}
	res := &todoviews.TodoResultView{
		ID:          &v.ID,
		UserID:      &v.UserID,
		Title:       &v.Title,
		Description: &v.Description,
		Completed:   &v.Completed,
		CreatedAt:   &v.CreatedAt,
		UpdatedAt:   &v.UpdatedAt,
	}

	return res
}
