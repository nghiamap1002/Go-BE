package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUser() string {
	return "concobebe"
}

func (ur *UserRepo) GetInfoUser2() string {
	return "concobebe"
}
