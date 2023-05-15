package helpers

import (
	"log"
	"os"

	"github.com/mbaraa/apollo-music/config/env"
	"github.com/mbaraa/apollo-music/data"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/models"
	"github.com/mbaraa/apollo-music/utils/jwt"
)

type StorageHelper struct {
	repo     data.CRUDRepo[models.Storage]
	userRepo data.CRUDRepo[models.User]
	jwtUtil  jwt.Manager[entities.JSON]
}

func NewStorageHelper(
	repo data.CRUDRepo[models.Storage],
	userRepo data.CRUDRepo[models.User],
	jwtUtil jwt.Manager[entities.JSON],
) *StorageHelper {
	return &StorageHelper{
		repo:     repo,
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (s *StorageHelper) CheckUserAndStorage(token, userPublicId, storageType string) bool {
	claims, err := s.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		log.Println(err)
		return false
	}

	if userPublicId != claims.Payload["publicId"] {
		log.Println(err)
		return false
	}

	if !s.checkStorageType(storageType) {
		log.Println(err)
		return false
	}

	dbUser, err := s.userRepo.GetByConds("public_id = ?", userPublicId)
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = s.repo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (s *StorageHelper) GetDetails(token string) (entities.JSON, int) {
	claims, err := s.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := s.userRepo.GetByConds("public_id = ?", claims.Payload["publicId"])
	if err != nil {
		log.Println(err)
		return response.Build(errors.Unauthorized, nil)
	}

	dbStorage, err := s.repo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	size := float64(dbStorage[0].Size) / 1024.0
	used := float64(dbStorage[0].Used) / 1024.0
	return response.Build(errors.None, entities.Storage{
		Size: size,
		Used: used,
		Free: size - used,
	})
}

func (s *StorageHelper) createStorage(size int64, user models.User) error {
	err := os.Mkdir(env.MusicDirectory()+"/"+user.PublicId, 0755)
	if err != nil {
		log.Println(err)
		return err
	}

	err = s.repo.Add(&models.Storage{
		UserId: user.Id,
		Size:   size,
		Used:   0,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *StorageHelper) destroyStorage(user models.User) error {
	err := os.RemoveAll(env.MusicDirectory() + "/" + user.PublicId)
	if err != nil {
		log.Println(err)
		return err
	}

	err = s.repo.Delete("user_id = ?", user.Id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *StorageHelper) checkStorageType(tp string) bool {
	return tp == "music"
}
