// Code generated by goa v3.20.0, DO NOT EDIT.
//
// todo server HTTP client CLI support package
//
// Command:
// $ goa gen backend/design

package cli

import (
	memoc "backend/gen/http/memo/client"
	todoc "backend/gen/http/todo/client"
	userc "backend/gen/http/user/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `user (create|login|get|update|delete)
todo (create|list|get|update|delete)
memo (create|list|get|update|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` user create --body '{
      "email": "helen_reichert@wuckertroberts.name",
      "name": "0qc",
      "password": "66a"
   }'` + "\n" +
		os.Args[0] + ` todo create --body '{
      "completed": true,
      "description": "9d8",
      "title": "z9z"
   }' --user-id 3730712348700016980` + "\n" +
		os.Args[0] + ` memo create --body '{
      "content": "64x"
   }' --user-id 2071962975920199465 --todo-id 4194034076411350345` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		userCreateBodyFlag = userCreateFlags.String("body", "REQUIRED", "")

		userLoginFlags    = flag.NewFlagSet("login", flag.ExitOnError)
		userLoginBodyFlag = userLoginFlags.String("body", "REQUIRED", "")

		userGetFlags      = flag.NewFlagSet("get", flag.ExitOnError)
		userGetUserIDFlag = userGetFlags.String("user-id", "REQUIRED", "ユーザーID")

		userUpdateFlags      = flag.NewFlagSet("update", flag.ExitOnError)
		userUpdateBodyFlag   = userUpdateFlags.String("body", "REQUIRED", "")
		userUpdateUserIDFlag = userUpdateFlags.String("user-id", "REQUIRED", "ユーザーID")

		userDeleteFlags      = flag.NewFlagSet("delete", flag.ExitOnError)
		userDeleteUserIDFlag = userDeleteFlags.String("user-id", "REQUIRED", "ユーザーID")

		todoFlags = flag.NewFlagSet("todo", flag.ContinueOnError)

		todoCreateFlags      = flag.NewFlagSet("create", flag.ExitOnError)
		todoCreateBodyFlag   = todoCreateFlags.String("body", "REQUIRED", "")
		todoCreateUserIDFlag = todoCreateFlags.String("user-id", "REQUIRED", "ユーザーID")

		todoListFlags         = flag.NewFlagSet("list", flag.ExitOnError)
		todoListUserIDFlag    = todoListFlags.String("user-id", "REQUIRED", "ユーザーID")
		todoListCompletedFlag = todoListFlags.String("completed", "", "")

		todoGetFlags      = flag.NewFlagSet("get", flag.ExitOnError)
		todoGetUserIDFlag = todoGetFlags.String("user-id", "REQUIRED", "ユーザーID")
		todoGetTodoIDFlag = todoGetFlags.String("todo-id", "REQUIRED", "タスクID")

		todoUpdateFlags      = flag.NewFlagSet("update", flag.ExitOnError)
		todoUpdateBodyFlag   = todoUpdateFlags.String("body", "REQUIRED", "")
		todoUpdateUserIDFlag = todoUpdateFlags.String("user-id", "REQUIRED", "ユーザーID")
		todoUpdateTodoIDFlag = todoUpdateFlags.String("todo-id", "REQUIRED", "タスクID")

		todoDeleteFlags      = flag.NewFlagSet("delete", flag.ExitOnError)
		todoDeleteUserIDFlag = todoDeleteFlags.String("user-id", "REQUIRED", "ユーザーID")
		todoDeleteTodoIDFlag = todoDeleteFlags.String("todo-id", "REQUIRED", "タスクID")

		memoFlags = flag.NewFlagSet("memo", flag.ContinueOnError)

		memoCreateFlags      = flag.NewFlagSet("create", flag.ExitOnError)
		memoCreateBodyFlag   = memoCreateFlags.String("body", "REQUIRED", "")
		memoCreateUserIDFlag = memoCreateFlags.String("user-id", "REQUIRED", "ユーザーID")
		memoCreateTodoIDFlag = memoCreateFlags.String("todo-id", "REQUIRED", "タスクID")

		memoListFlags      = flag.NewFlagSet("list", flag.ExitOnError)
		memoListUserIDFlag = memoListFlags.String("user-id", "REQUIRED", "ユーザーID")
		memoListTodoIDFlag = memoListFlags.String("todo-id", "REQUIRED", "タスクID")

		memoGetFlags      = flag.NewFlagSet("get", flag.ExitOnError)
		memoGetUserIDFlag = memoGetFlags.String("user-id", "REQUIRED", "ユーザーID")
		memoGetTodoIDFlag = memoGetFlags.String("todo-id", "REQUIRED", "タスクID")
		memoGetMemoIDFlag = memoGetFlags.String("memo-id", "REQUIRED", "メモID")

		memoUpdateFlags      = flag.NewFlagSet("update", flag.ExitOnError)
		memoUpdateBodyFlag   = memoUpdateFlags.String("body", "REQUIRED", "")
		memoUpdateUserIDFlag = memoUpdateFlags.String("user-id", "REQUIRED", "ユーザーID")
		memoUpdateTodoIDFlag = memoUpdateFlags.String("todo-id", "REQUIRED", "タスクID")
		memoUpdateMemoIDFlag = memoUpdateFlags.String("memo-id", "REQUIRED", "メモID")

		memoDeleteFlags      = flag.NewFlagSet("delete", flag.ExitOnError)
		memoDeleteUserIDFlag = memoDeleteFlags.String("user-id", "REQUIRED", "ユーザーID")
		memoDeleteTodoIDFlag = memoDeleteFlags.String("todo-id", "REQUIRED", "タスクID")
		memoDeleteMemoIDFlag = memoDeleteFlags.String("memo-id", "REQUIRED", "メモID")
	)
	userFlags.Usage = userUsage
	userCreateFlags.Usage = userCreateUsage
	userLoginFlags.Usage = userLoginUsage
	userGetFlags.Usage = userGetUsage
	userUpdateFlags.Usage = userUpdateUsage
	userDeleteFlags.Usage = userDeleteUsage

	todoFlags.Usage = todoUsage
	todoCreateFlags.Usage = todoCreateUsage
	todoListFlags.Usage = todoListUsage
	todoGetFlags.Usage = todoGetUsage
	todoUpdateFlags.Usage = todoUpdateUsage
	todoDeleteFlags.Usage = todoDeleteUsage

	memoFlags.Usage = memoUsage
	memoCreateFlags.Usage = memoCreateUsage
	memoListFlags.Usage = memoListUsage
	memoGetFlags.Usage = memoGetUsage
	memoUpdateFlags.Usage = memoUpdateUsage
	memoDeleteFlags.Usage = memoDeleteUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "user":
			svcf = userFlags
		case "todo":
			svcf = todoFlags
		case "memo":
			svcf = memoFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "user":
			switch epn {
			case "create":
				epf = userCreateFlags

			case "login":
				epf = userLoginFlags

			case "get":
				epf = userGetFlags

			case "update":
				epf = userUpdateFlags

			case "delete":
				epf = userDeleteFlags

			}

		case "todo":
			switch epn {
			case "create":
				epf = todoCreateFlags

			case "list":
				epf = todoListFlags

			case "get":
				epf = todoGetFlags

			case "update":
				epf = todoUpdateFlags

			case "delete":
				epf = todoDeleteFlags

			}

		case "memo":
			switch epn {
			case "create":
				epf = memoCreateFlags

			case "list":
				epf = memoListFlags

			case "get":
				epf = memoGetFlags

			case "update":
				epf = memoUpdateFlags

			case "delete":
				epf = memoDeleteFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "user":
			c := userc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = userc.BuildCreatePayload(*userCreateBodyFlag)
			case "login":
				endpoint = c.Login()
				data, err = userc.BuildLoginPayload(*userLoginBodyFlag)
			case "get":
				endpoint = c.Get()
				data, err = userc.BuildGetPayload(*userGetUserIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = userc.BuildUpdatePayload(*userUpdateBodyFlag, *userUpdateUserIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = userc.BuildDeletePayload(*userDeleteUserIDFlag)
			}
		case "todo":
			c := todoc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = todoc.BuildCreatePayload(*todoCreateBodyFlag, *todoCreateUserIDFlag)
			case "list":
				endpoint = c.List()
				data, err = todoc.BuildListPayload(*todoListUserIDFlag, *todoListCompletedFlag)
			case "get":
				endpoint = c.Get()
				data, err = todoc.BuildGetPayload(*todoGetUserIDFlag, *todoGetTodoIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = todoc.BuildUpdatePayload(*todoUpdateBodyFlag, *todoUpdateUserIDFlag, *todoUpdateTodoIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = todoc.BuildDeletePayload(*todoDeleteUserIDFlag, *todoDeleteTodoIDFlag)
			}
		case "memo":
			c := memoc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = memoc.BuildCreatePayload(*memoCreateBodyFlag, *memoCreateUserIDFlag, *memoCreateTodoIDFlag)
			case "list":
				endpoint = c.List()
				data, err = memoc.BuildListPayload(*memoListUserIDFlag, *memoListTodoIDFlag)
			case "get":
				endpoint = c.Get()
				data, err = memoc.BuildGetPayload(*memoGetUserIDFlag, *memoGetTodoIDFlag, *memoGetMemoIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = memoc.BuildUpdatePayload(*memoUpdateBodyFlag, *memoUpdateUserIDFlag, *memoUpdateTodoIDFlag, *memoUpdateMemoIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = memoc.BuildDeletePayload(*memoDeleteUserIDFlag, *memoDeleteTodoIDFlag, *memoDeleteMemoIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// userUsage displays the usage of the user command and its subcommands.
func userUsage() {
	fmt.Fprintf(os.Stderr, `ユーザー管理サービス
Usage:
    %[1]s [globalflags] user COMMAND [flags]

COMMAND:
    create: 新規ユーザーを作成します
    login: ユーザーログイン処理を行います
    get: ユーザー情報を取得します
    update: ユーザー情報を更新します
    delete: ユーザーを削除します

Additional help:
    %[1]s user COMMAND --help
`, os.Args[0])
}
func userCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user create -body JSON

新規ユーザーを作成します
    -body JSON: 

Example:
    %[1]s user create --body '{
      "email": "helen_reichert@wuckertroberts.name",
      "name": "0qc",
      "password": "66a"
   }'
`, os.Args[0])
}

func userLoginUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user login -body JSON

ユーザーログイン処理を行います
    -body JSON: 

Example:
    %[1]s user login --body '{
      "email": "Rerum ipsam id eius ea ducimus.",
      "password": "Molestias at."
   }'
`, os.Args[0])
}

func userGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user get -user-id INT

ユーザー情報を取得します
    -user-id INT: ユーザーID

Example:
    %[1]s user get --user-id 4198301971282410797
`, os.Args[0])
}

func userUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user update -body JSON -user-id INT

ユーザー情報を更新します
    -body JSON: 
    -user-id INT: ユーザーID

Example:
    %[1]s user update --body '{
      "email": "damon.hahn@morar.org",
      "name": "knj"
   }' --user-id 2578890648404992561
`, os.Args[0])
}

func userDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user delete -user-id INT

ユーザーを削除します
    -user-id INT: ユーザーID

Example:
    %[1]s user delete --user-id 4222867704420030937
`, os.Args[0])
}

// todoUsage displays the usage of the todo command and its subcommands.
func todoUsage() {
	fmt.Fprintf(os.Stderr, `タスク管理サービス
Usage:
    %[1]s [globalflags] todo COMMAND [flags]

COMMAND:
    create: 新規タスクを作成します
    list: タスク一覧を取得します
    get: タスク詳細を取得します
    update: タスクを更新します
    delete: タスクを削除します

Additional help:
    %[1]s todo COMMAND --help
`, os.Args[0])
}
func todoCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo create -body JSON -user-id INT

新規タスクを作成します
    -body JSON: 
    -user-id INT: ユーザーID

Example:
    %[1]s todo create --body '{
      "completed": true,
      "description": "9d8",
      "title": "z9z"
   }' --user-id 3730712348700016980
`, os.Args[0])
}

func todoListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo list -user-id INT -completed BOOL

タスク一覧を取得します
    -user-id INT: ユーザーID
    -completed BOOL: 

Example:
    %[1]s todo list --user-id 8024336273802408321 --completed false
`, os.Args[0])
}

func todoGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo get -user-id INT -todo-id INT

タスク詳細を取得します
    -user-id INT: ユーザーID
    -todo-id INT: タスクID

Example:
    %[1]s todo get --user-id 2490929946099554739 --todo-id 7591740459006074328
`, os.Args[0])
}

func todoUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo update -body JSON -user-id INT -todo-id INT

タスクを更新します
    -body JSON: 
    -user-id INT: ユーザーID
    -todo-id INT: タスクID

Example:
    %[1]s todo update --body '{
      "completed": true,
      "description": "iua",
      "title": "sfq"
   }' --user-id 7113982788117776079 --todo-id 5231178207423615515
`, os.Args[0])
}

func todoDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo delete -user-id INT -todo-id INT

タスクを削除します
    -user-id INT: ユーザーID
    -todo-id INT: タスクID

Example:
    %[1]s todo delete --user-id 7415249690300749021 --todo-id 3877079342457875127
`, os.Args[0])
}

// memoUsage displays the usage of the memo command and its subcommands.
func memoUsage() {
	fmt.Fprintf(os.Stderr, `メモ管理サービス
Usage:
    %[1]s [globalflags] memo COMMAND [flags]

COMMAND:
    create: 新規メモを作成します
    list: タスクに関連するメモ一覧を取得します
    get: メモ詳細を取得します
    update: メモを更新します
    delete: メモを削除します

Additional help:
    %[1]s memo COMMAND --help
`, os.Args[0])
}
func memoCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] memo create -body JSON -user-id INT -todo-id INT

新規メモを作成します
    -body JSON: 
    -user-id INT: ユーザーID
    -todo-id INT: タスクID

Example:
    %[1]s memo create --body '{
      "content": "64x"
   }' --user-id 2071962975920199465 --todo-id 4194034076411350345
`, os.Args[0])
}

func memoListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] memo list -user-id INT -todo-id INT

タスクに関連するメモ一覧を取得します
    -user-id INT: ユーザーID
    -todo-id INT: タスクID

Example:
    %[1]s memo list --user-id 2927266822603476504 --todo-id 6422719988830496889
`, os.Args[0])
}

func memoGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] memo get -user-id INT -todo-id INT -memo-id INT

メモ詳細を取得します
    -user-id INT: ユーザーID
    -todo-id INT: タスクID
    -memo-id INT: メモID

Example:
    %[1]s memo get --user-id 7993033404756377863 --todo-id 6964843889273683616 --memo-id 7363813684071047719
`, os.Args[0])
}

func memoUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] memo update -body JSON -user-id INT -todo-id INT -memo-id INT

メモを更新します
    -body JSON: 
    -user-id INT: ユーザーID
    -todo-id INT: タスクID
    -memo-id INT: メモID

Example:
    %[1]s memo update --body '{
      "content": "hp6"
   }' --user-id 3272338859504802213 --todo-id 4529182633562951640 --memo-id 9210136981290392496
`, os.Args[0])
}

func memoDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] memo delete -user-id INT -todo-id INT -memo-id INT

メモを削除します
    -user-id INT: ユーザーID
    -todo-id INT: タスクID
    -memo-id INT: メモID

Example:
    %[1]s memo delete --user-id 6343982725536526004 --todo-id 7527679246805584637 --memo-id 7720890222943482465
`, os.Args[0])
}
