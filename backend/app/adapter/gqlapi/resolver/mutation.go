package resolver

import (
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/requester"
	"github.com/short-d/short/backend/app/usecase/url"
)

// Mutation represents GraphQL mutation resolver
type Mutation struct {
	logger            logger.Logger
	urlCreator        url.Creator
	requesterVerifier requester.Verifier
	authenticator     authenticator.Authenticator
	changeLog         changelog.ChangeLog
}

// AuthMutationArgs represents possible parameters for AuthMutation endpoint
type AuthMutationArgs struct {
	AuthToken       *string
	CaptchaResponse string
}

// AuthMutation extracts user information from authentication token
func (m Mutation) AuthMutation(args *AuthMutationArgs) (*AuthMutation, error) {
	isHuman, err := m.requesterVerifier.IsHuman(args.CaptchaResponse)

	if err != nil {
		return nil, ErrUnknown{}
	}

	if !isHuman {
		return nil, ErrNotHuman{}
	}

	authMutation := newAuthMutation(args.AuthToken, m.authenticator, m.changeLog, m.urlCreator)
	return &authMutation, nil
}

func newMutation(
	logger logger.Logger,
	changeLog changelog.ChangeLog,
	urlCreator url.Creator,
	requesterVerifier requester.Verifier,
	authenticator authenticator.Authenticator,
) Mutation {
	return Mutation{
		logger:            logger,
		changeLog:         changeLog,
		urlCreator:        urlCreator,
		requesterVerifier: requesterVerifier,
		authenticator:     authenticator,
	}
}
