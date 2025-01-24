package model

type User struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Username  string `db:"username" json:"username"`
	Position  string `db:"position" json:"position"`
	Salary    int    `db:"salary" json:"salary"`
	Password  string `db:"password" json:"-"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
