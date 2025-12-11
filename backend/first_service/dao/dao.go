package dao

type DAO interface {
	CreateMessage(message string) error
}
