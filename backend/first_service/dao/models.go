package dao

type Message struct {
	ID      int    `db:"id"`
	Message string `db:"message"`
}
