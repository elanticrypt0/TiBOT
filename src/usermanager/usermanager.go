package usermanager

type UserManager struct {
	Users        []User
	selectedUser *User
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (um *UserManager) AddUser(id int64, username string, is_admin bool) {
	um.Users = append(um.Users, User{
		ID:       id,
		Username: username,
		IsAdmin:  is_admin,
	})
}

func (um *UserManager) GetUser() *User {
	return um.selectedUser
}

func (um *UserManager) IsAllowedUser(userID int64) bool {
	userExists := false

	for _, user := range um.Users {
		if user.ID == userID {
			um.selectedUser = &user
			userExists = true
			break
		}
	}

	return userExists
}
