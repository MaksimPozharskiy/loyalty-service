package service

import "github.com/MaksimPozharskiy/loyalty-service/internal/repository"

type UserService interface {
}

type UserServiceImp struct {
	storageRepository repository.StorageRepository
}

func NewUserService(storageRepository repository.StorageRepository) UserService {
	return &UserServiceImp{
		storageRepository: storageRepository,
	}
}
