package repository

type UserRepository {	
	Save(user entity.User) (entity.User, error)

}