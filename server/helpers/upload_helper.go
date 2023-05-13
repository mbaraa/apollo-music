package helpers

import (
	goerrors "errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/mbaraa/apollo-music/config/env"
	"github.com/mbaraa/apollo-music/data"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/enums"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/models"
	"github.com/mbaraa/apollo-music/utils/jwt"
)

type UploadHelper struct {
	storageRepo data.CRUDRepo[models.Storage]
	userRepo    data.CRUDRepo[models.User]
	jwtUtil     jwt.Manager[entities.JSON]
}

func NewUploadHelper(
	storageRepo data.CRUDRepo[models.Storage],
	userRepo data.CRUDRepo[models.User],
	jwtUtil jwt.Manager[entities.JSON],
) *UploadHelper {
	return &UploadHelper{
		storageRepo: storageRepo,
		userRepo:    userRepo,
		jwtUtil:     jwtUtil,
	}
}

// TODO
// handle duplicate files
func (u *UploadHelper) UploadFile(token string, audioType enums.AudioType, fileHeader *multipart.FileHeader) (entities.JSON, int) {
	claims, err := u.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		return response.Build(errors.InvalidToken, nil)
	}

	if !strings.Contains(fileHeader.Header.Get("Content-Type"), "audio") {
		return response.Build(errors.BadRequest, nil)
	}

	dbUser, err := u.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	dbStorage, err := u.storageRepo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	if dbStorage[0].Size-dbStorage[0].Used < fileHeader.Size/(1024*1024) {
		return response.Build(errors.InsufficientStorage, nil)
	}

	saveDirectory := fmt.Sprintf("%s/%s/%s", env.MusicDirectory(), dbUser[0].PublicId, audioType.String())
	err = os.Mkdir(saveDirectory, 0755)
	if err != nil && !goerrors.Is(err, os.ErrExist) {
		fmt.Println(err)
		return response.Build(errors.InternalServerError, nil)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return response.Build(errors.BadRequest, nil)
	}
	defer file.Close()

	filePath := fmt.Sprintf("%s/%s", saveDirectory, fileHeader.Filename)
	fileToSave, err := os.Create(filePath)

	_, err = io.Copy(fileToSave, file)
	if err != nil {
		fmt.Println(err)
		return response.Build(errors.InternalServerError, nil)
	}

	err = u.storageRepo.Update(&models.Storage{
		Used: dbStorage[0].Used + fileHeader.Size/(1024*1024),
	}, "id = ?", dbStorage[0].Id)
	if err != nil {
		fmt.Println(err)
		return response.Build(errors.InternalServerError, nil)
	}

	return response.Build(errors.None, nil)
}
