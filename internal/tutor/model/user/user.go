package user

const (
	LoginKey = "ulogin"
	IDKey    = "uid"
)

type User struct {
	Name    string `db:"name"`
	Surname string `db:"surname"`
	Login   string `db:"login"`
	Email   string `db:"email"`
	ID      uint   `db:"id"`
}
