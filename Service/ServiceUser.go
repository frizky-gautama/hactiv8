package Service

// Berisi entitas data
type User struct {
	Nama string
	Age  string
}

type userService struct {
	db []*User
}

type UserHandler interface {
	CreateUser(u *User) string
	GetAllUser() []*User
}

func HandlerFunc(db []*User) UserHandler {
	return &userService{
		db: db,
	}
}

func (u *userService) CreateUser(user *User) string {
	u.db = append(u.db, user)
	return user.Nama + ", " + user.Age + " Tahun" + " Telah Terdaftar"
}

func (u *userService) GetAllUser() []*User {
	return u.db
}
