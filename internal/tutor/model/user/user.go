package user

const (
	LoginKey = "ulogin"
	IDKey    = "uid"
)

type User struct {
	Name    string
	Surname string
	Login   string
	ID      uint
}
