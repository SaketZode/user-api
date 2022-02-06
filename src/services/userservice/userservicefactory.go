package userservice

func New() (userservice UserService) {
	return &UserServiceImpl{}
}
