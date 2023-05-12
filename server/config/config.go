package config

import (
	"github.com/mbaraa/apollo-music/data"
	"github.com/mbaraa/apollo-music/data/db"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/models"
	"github.com/mbaraa/apollo-music/utils/jwt"
)

var (
	dbConnection     = db.GetDBConnector()
	userRepo         data.CRUDRepo[models.User]
	verificationRepo data.CRUDRepo[models.Verification]
	subscriptionRepo data.CRUDRepo[models.Subscription]
)

var (
	jwtUtil jwt.Manager[entities.JSON]
)

func init() {
	db.InitTables()
	userRepo = db.NewBaseDB[models.User](dbConnection)
	verificationRepo = db.NewBaseDB[models.Verification](dbConnection)
	subscriptionRepo = db.NewBaseDB[models.Subscription](dbConnection)
	jwtUtil = jwt.NewJWTImpl()
}

func UserRepo() data.CRUDRepo[models.User] {
	return userRepo
}

func VerificationRepo() data.CRUDRepo[models.Verification] {
	return verificationRepo
}

func SubscriptionRepo() data.CRUDRepo[models.Subscription] {
	return subscriptionRepo
}

func JWTUtil() jwt.Manager[entities.JSON] {
	return jwtUtil
}
