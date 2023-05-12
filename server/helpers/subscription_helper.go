package helpers

import (
	"time"

	"github.com/mbaraa/apollo-music/config/env"
	"github.com/mbaraa/apollo-music/data"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/enums"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/models"
	"github.com/mbaraa/apollo-music/utils/jwt"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/card"
	"github.com/stripe/stripe-go/v74/customer"
	"github.com/stripe/stripe-go/v74/subscription"
)

type SubscriptionHelper struct {
	repo     data.CRUDRepo[models.Subscription]
	userRepo data.CRUDRepo[models.User]
	jwtUtil  jwt.Manager[entities.JSON]
}

func NewSubscriptionHelper(
	repo data.CRUDRepo[models.Subscription],
	userRepo data.CRUDRepo[models.User],
	jwtUtil jwt.Manager[entities.JSON],
) *SubscriptionHelper {
	stripe.Key = env.StripeSecretKey()

	return &SubscriptionHelper{
		repo:     repo,
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (s *SubscriptionHelper) StartSubscription(token, cardToken string, plan enums.Plan) (entities.JSON, int) {
	claims, err := s.jwtUtil.Decode(token, jwt.CheckoutToken)
	if err != nil {
		claims, err = s.jwtUtil.Decode(token, jwt.SessionToken)
		if err != nil {
			return response.Build(errors.InvalidToken, nil)
		}
	}

	dbUser, err := s.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	customerId, err := s.createCustomer(dbUser[0])
	if err != nil {
		return response.Build(errors.PaymentError, nil)
	}

	_, err = s.createCard(cardToken, customerId)
	if err != nil {
		return response.Build(errors.PaymentError, nil)
	}

	subscriptionId, err := s.createSubscription(customerId, plan)
	if err != nil {
		return response.Build(errors.InsufficientFunds, nil)
	}

	err = s.repo.Add(&models.Subscription{
		UserId:               dbUser[0].Id,
		Size:                 plan.Size(),
		Price:                s.calculatePriceInCents(plan.Size()),
		ValidBefore:          time.Now().UTC().Add(time.Hour * 24 * 30),
		StripeCustomerId:     customerId,
		StripeSubscriptionId: subscriptionId,
	})
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	err = s.userRepo.Update(&models.User{
		Status: enums.ActiveStatus,
	}, "id = ?", dbUser[0].Id)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	return response.Build(errors.None, nil)
}

func (s *SubscriptionHelper) CancelSubscription(token string) (entities.JSON, int) {
	claims, err := s.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := s.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	dbSub, err := s.repo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	params := &stripe.SubscriptionParams{CancelAtPeriodEnd: stripe.Bool(true)}
	_, err = subscription.Update(dbSub[0].StripeSubscriptionId, params)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	err = s.userRepo.Update(&models.User{
		Status: enums.InactiveStatus,
	}, "id = ?", dbUser[0].Id)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	err = s.repo.Delete("id = ?", dbSub[0].Id)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	return response.Build(errors.None, nil)
}

func (s *SubscriptionHelper) CheckSubscription(token string) (entities.JSON, int) {
	claims, err := s.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := s.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	dbSub, err := s.repo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	return response.Build(errors.None, entities.Subscription{
		Size:        float64(dbSub[0].Size / 1024),
		Price:       float64(dbSub[0].Price) / 1000.0,
		ValidBefore: dbSub[0].ValidBefore,
	})
}

func (s *SubscriptionHelper) createCustomer(user models.User) (customerId string, err error) {
	params := &stripe.CustomerParams{
		Name:        stripe.String(user.FullName),
		Email:       stripe.String(user.Email),
		Description: stripe.String(user.PublicId),
	}
	cus, err := customer.New(params)
	if err != nil {
		return "", err
	}

	return cus.ID, nil
}

func (s *SubscriptionHelper) createCard(cardToken, customerId string) (cardId string, err error) {
	params := &stripe.CardParams{
		Customer: stripe.String(customerId),
		Token:    stripe.String(cardToken),
	}
	crd, err := card.New(params)
	if err != nil {
		return "", err
	}

	return crd.ID, nil
}

func (s *SubscriptionHelper) createSubscription(customerId string, plan enums.Plan) (subscriptionId string, err error) {
	planPriceId := plan.PlanStripeProductId()

	if len(planPriceId) == 0 {
		return "", errors.ErrPaymentError.New("")
	}

	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerId),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(planPriceId),
			},
		},
	}
	sub, err := subscription.New(params)
	if err != nil {
		return "", err
	}

	return sub.ID, nil
}

func (s *SubscriptionHelper) calculatePriceInCents(migs int64) int64 {
	gigs := float64(migs) / 1024.0
	cents := 1000.0 * ((3.0/80.0)*gigs + (9.0 / 4.0))
	return int64(cents)
}