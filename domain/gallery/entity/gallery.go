package entity

type Gallery struct {
	Id     int64 `db:"id"`
	Source string `db:"source"`
}
