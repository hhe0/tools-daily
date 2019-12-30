package main

type UserRepository interface {
	GetUsername(username string)
}

type userRepository struct {
	name string
}

func NewUserRepository() UserRepository {
	return &userRepository{
		name: "hehong",
	}
}

func (r userRepository) GetUsername(username string ) {

}