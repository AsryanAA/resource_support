package models

const (
	ErrorUnique   = "запись с таким наименованием уже существует"
	ErrorNotFound = "данная запись не найдена"
)

var DataBaseErrors = map[int]string{
	1: "unique constraint",
	2: "record not found",
}
