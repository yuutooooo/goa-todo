package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("todo app", func() {
	Title("Todo App")
	Description("ユーザー登録からタスク管理、メモ追加までのアプリ")
	Server("todo server", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

// エラー結果の型定義
var APIErrorResult = Type("APIErrorResult", func() {
	Description("エラーレスポンス")
	Attribute("name", String, "エラー名")
	Attribute("message", String, "エラーメッセージ")
	Required("name", "message")
	ErrorName("name")
})

var User = Type("User", func() {
	Description("user作成に必要なデータ型")
	Attribute("name", String, "ユーザー名", func() {
		MinLength(3)
	})
	Attribute("email", String, "メールアドレス", func() {
		Format("email")
	})
	Attribute("password", String, "パスワード", func() {
		MinLength(8)
	})
	Required("name", "email", "password")
})

var UserResult = Type("UserResult", func() {
	Description("user作成のレスポンス")
	Attribute("id", Int, "userのid")
	Attribute("name", String, "ユーザー名")
	Attribute("email", String, "メールアドレス")
	Attribute("created_at", String, "作成日時")
	Attribute("updated_at", String, "更新日時")
	Required("id", "name", "email", "created_at", "updated_at")
})

var UserUpdatePayload = Type("UserUpdatePayload", func() {
	Description("ユーザー情報更新のデータ型")
	Attribute("name", String, "ユーザー名", func() {
		MinLength(3)
	})
	Attribute("email", String, "メールアドレス", func() {
		Format("email")
	})
})

var Todo = Type("Todo", func() {
	Description("タスク作成のデータ型")
	Attribute("title", String, "タスクのタイトル", func() {
		MinLength(3)
	})
	Attribute("description", String, "タスクの説明", func() {
		MinLength(3)
	})
	Attribute("completed", Boolean, "タスクが完了しているかどうか", func() {
		Default(false)
	})
	Required("title", "description")
})

var TodoResult = Type("TodoResult", func() {
	Description("タスク作成のレスポンス")
	Attribute("id", Int, "タスクのid")
	Attribute("user_id", Int, "ユーザーID")
	Attribute("title", String, "タスクのタイトル")
	Attribute("description", String, "タスクの説明")
	Attribute("completed", Boolean, "タスクが完了しているかどうか")
	Attribute("created_at", String, "作成日時")
	Attribute("updated_at", String, "更新日時")
	Required("id", "user_id", "title", "description", "completed", "created_at", "updated_at")
})

var Memo = Type("Memo", func() {
	Description("メモのデータ型")
	Attribute("content", String, "メモの内容", func() {
		MinLength(3)
	})
	Required("content")
})

var MemoResult = Type("MemoResult", func() {
	Description("メモのレスポンス")
	Attribute("id", Int, "メモのid")
	Attribute("todo_id", Int, "タスクID")
	Attribute("content", String, "メモの内容")
	Attribute("created_at", String, "作成日時")
	Attribute("updated_at", String, "更新日時")
	Required("id", "todo_id", "content", "created_at", "updated_at")
})

var LoginPayload = Type("LoginPayload", func() {
	Description("ログイン用のデータ型")
	Attribute("email", String, "メールアドレス")
	Attribute("password", String, "パスワード")
	Required("email", "password")
})

var TodoResultCollection = ResultType("application/vnd.todo-collection+json", func() {
	TypeName("TodoCollection")
	Attributes(func() {
		Attribute("items", ArrayOf(TodoResult))
	})
	View("default", func() {
		Attribute("items")
	})
})

var MemoResultCollection = ResultType("application/vnd.memo-collection+json", func() {
	TypeName("MemoCollection")
	Attributes(func() {
		Attribute("items", ArrayOf(MemoResult))
	})
	View("default", func() {
		Attribute("items")
	})
})

var _ = Service("user", func() {
	Description("ユーザー管理サービス")

	// ユーザー作成
	Method("create", func() {
		Description("新規ユーザーを作成します")
		Payload(User)
		Result(UserResult)
		Error("BadRequest", APIErrorResult)
		HTTP(func() {
			POST("/users")
			Response(StatusCreated)
			Response("BadRequest", StatusBadRequest)
		})
	})

	// ログイン
	Method("login", func() {
		Description("ユーザーログイン処理を行います")
		Payload(LoginPayload)
		Result(UserResult)
		Error("Unauthorized", APIErrorResult)
		Error("BadRequest", APIErrorResult)
		HTTP(func() {
			POST("/login")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
			Response("BadRequest", StatusBadRequest)
		})
	})

	// ユーザー取得
	Method("get", func() {
		Description("ユーザー情報を取得します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Required("user_id")
		})
		Result(UserResult)
		Error("NotFound", APIErrorResult)
		HTTP(func() {
			GET("/users/{user_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	// ユーザー更新
	Method("update", func() {
		Description("ユーザー情報を更新します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Attribute("name", String, "ユーザー名", func() {
				MinLength(3)
			})
			Attribute("email", String, "メールアドレス", func() {
				Format("email")
			})
			Required("user_id")
		})
		Result(UserResult)
		Error("NotFound", APIErrorResult)
		Error("BadRequest", APIErrorResult)
		HTTP(func() {
			PUT("/users/{user_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})
	})

	// ユーザー削除
	Method("delete", func() {
		Description("ユーザーを削除します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Required("user_id")
		})
		Error("NotFound", APIErrorResult)
		HTTP(func() {
			DELETE("/users/{user_id}")
			Response(StatusNoContent)
			Response("NotFound", StatusNotFound)
		})
	})
})

var _ = Service("todo", func() {
	Description("タスク管理サービス")

	// タスク作成
	Method("create", func() {
		Description("新規タスクを作成します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Extend(Todo)
			Required("user_id")
		})
		Result(TodoResult)
		Error("BadRequest", APIErrorResult)
		HTTP(func() {
			POST("/users/{user_id}/todos")
			Response(StatusCreated)
			Response("BadRequest", StatusBadRequest)
		})
	})

	// タスク一覧取得
	Method("list", func() {
		Description("タスク一覧を取得します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Attribute("completed", Boolean, "完了済みタスクのみ取得", func() {
				Default(false)
			})
			Required("user_id")
		})
		Result(TodoResultCollection)
		HTTP(func() {
			GET("/users/{user_id}/todos")
			Param("completed")
			Response(StatusOK)
		})
	})

	// タスク詳細取得
	Method("get", func() {
		Description("タスク詳細を取得します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Attribute("todo_id", Int, "タスクID")
			Required("user_id", "todo_id")
		})
		Result(TodoResult)
		Error("NotFound", APIErrorResult)
		HTTP(func() {
			GET("/users/{user_id}/todos/{todo_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	// タスク更新
	Method("update", func() {
		Description("タスクを更新します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Attribute("todo_id", Int, "タスクID")
			Attribute("title", String, "タスクのタイトル", func() {
				MinLength(3)
			})
			Attribute("description", String, "タスクの説明", func() {
				MinLength(3)
			})
			Attribute("completed", Boolean, "タスクが完了しているかどうか")
			Required("user_id", "todo_id")
		})
		Result(TodoResult)
		Error("NotFound", APIErrorResult)
		Error("BadRequest", APIErrorResult)
		HTTP(func() {
			PUT("/users/{user_id}/todos/{todo_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})
	})

	// タスク削除
	Method("delete", func() {
		Description("タスクを削除します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Attribute("todo_id", Int, "タスクID")
			Required("user_id", "todo_id")
		})
		Error("NotFound", APIErrorResult)
		HTTP(func() {
			DELETE("/users/{user_id}/todos/{todo_id}")
			Response(StatusNoContent)
			Response("NotFound", StatusNotFound)
		})
	})
})

var _ = Service("memo", func() {
	Description("メモ管理サービス")

	// メモ作成
	Method("create", func() {
		Description("新規メモを作成します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Attribute("todo_id", Int, "タスクID")
			Extend(Memo)
			Required("user_id", "todo_id")
		})
		Result(MemoResult)
		Error("BadRequest", APIErrorResult)
		Error("NotFound", APIErrorResult)
		HTTP(func() {
			POST("/users/{user_id}/todos/{todo_id}/memos")
			Response(StatusCreated)
			Response("BadRequest", StatusBadRequest)
			Response("NotFound", StatusNotFound)
		})
	})

	// メモ一覧取得
	Method("list", func() {
		Description("タスクに関連するメモ一覧を取得します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Attribute("todo_id", Int, "タスクID")
			Required("user_id", "todo_id")
		})
		Result(MemoResultCollection)
		Error("NotFound", APIErrorResult)
		HTTP(func() {
			GET("/users/{user_id}/todos/{todo_id}/memos")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	// メモ詳細取得
	Method("get", func() {
		Description("メモ詳細を取得します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Attribute("todo_id", Int, "タスクID")
			Attribute("memo_id", Int, "メモID")
			Required("user_id", "todo_id", "memo_id")
		})
		Result(MemoResult)
		Error("NotFound", APIErrorResult)
		HTTP(func() {
			GET("/users/{user_id}/todos/{todo_id}/memos/{memo_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	// メモ更新
	Method("update", func() {
		Description("メモを更新します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Attribute("todo_id", Int, "タスクID")
			Attribute("memo_id", Int, "メモID")
			Attribute("content", String, "メモの内容", func() {
				MinLength(3)
			})
			Required("user_id", "todo_id", "memo_id", "content")
		})
		Result(MemoResult)
		Error("NotFound", APIErrorResult)
		Error("BadRequest", APIErrorResult)
		HTTP(func() {
			PUT("/users/{user_id}/todos/{todo_id}/memos/{memo_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})
	})

	// メモ削除
	Method("delete", func() {
		Description("メモを削除します")
		Payload(func() {
			Attribute("user_id", Int, "ユーザーID")
			Attribute("todo_id", Int, "タスクID")
			Attribute("memo_id", Int, "メモID")
			Required("user_id", "todo_id", "memo_id")
		})
		Error("NotFound", APIErrorResult)
		HTTP(func() {
			DELETE("/users/{user_id}/todos/{todo_id}/memos/{memo_id}")
			Response(StatusNoContent)
			Response("NotFound", StatusNotFound)
		})
	})
})
