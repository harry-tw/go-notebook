package dao

type OrderDao struct {
	dbURL string
}

func NewOrderDao(dbURL string) (*OrderDao, error) {
	return &OrderDao{dbURL: dbURL}, nil
}

type UserDao struct {
	dbURL string
}

func NewUserDao(dbURL string) (*UserDao, error) {
	return &UserDao{dbURL: dbURL}, nil
}
