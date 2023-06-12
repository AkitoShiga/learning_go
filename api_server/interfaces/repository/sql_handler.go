package repository

type SqlHandler interface {
	// 抽象化された結果を返す
	// 可変長のinterfaceはプレースホルダの実体
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
