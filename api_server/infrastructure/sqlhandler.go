package infrastructure

// DBの実装を取り扱う、初期化処理を実装
// 外部のパッケージを取り扱っているのでinfrastructure層に定義する
// データアクセス層(リポジトリ)からDBドライバに関わる部分を隠蔽しておく
// repositoryで定義されたインターフェースを実装して、ドライバの実装を隠蔽する,またリポジトリのインターフェースに依存させるようにする
import (
	"api_server/interfaces/repository"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() repository.SqlHandler {
	conn, err := sql.Open("sqlite3", "./sample.sql")

	if err != nil {
		panic(err.Error)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (repository.Result, error) {
	res := SqlResult{} // interfaceで定義されたメソッドを定義していればrepository.Resultとしてみなされる
	result, err := handler.Conn.Exec(statement, args...)

	if err != nil {
		return res, err
	}

	res.Result = result

	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (repository.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}

	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
