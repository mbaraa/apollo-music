package auth

import (
	"github.com/mbaraa/apollo-music/data"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/models"
	"github.com/mbaraa/apollo-music/utils/jwt"
)

type SessionHelper struct {
	repo    data.CRUDRepo[models.User]
	jwtUtil jwt.Manager[entities.JSON]
}

func NewSessionHelper(
	repo data.CRUDRepo[models.User],
	jwtUtil jwt.Manager[entities.JSON],
) *SessionHelper {
	return &SessionHelper{
		repo:    repo,
		jwtUtil: jwtUtil,
	}
}

func (s *SessionHelper) CheckSession(token string) (entities.JSON, int) {
	claims, err := s.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := s.repo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	return response.Build(errors.None, entities.JSON{
		"fullName": dbUser[0].FullName,
		"email":    dbUser[0].Email,
		"publicId": dbUser[0].PublicId,
		"status":   dbUser[0].Status,
	})
}
