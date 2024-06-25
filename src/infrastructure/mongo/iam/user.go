package iam

import "MSS/src/domain/user"

type User struct {
	Phone    string `bson:"phone"`
	Password string `bson:"password"`
}

func toDomainUser(u User) user.User {
	return user.NewUser(u.Phone, u.Password)
}

func toDomainUserList(us []User) []user.User {
	res := make([]user.User, 0)
	for _, u := range us {
		res = append(res, toDomainUser(u))
	}
	return res
}
