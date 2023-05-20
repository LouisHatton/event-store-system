package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/LouisHatton/insight-wave/internal/api/responses"
	"github.com/LouisHatton/insight-wave/internal/users"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type Middleware struct {
	client *auth.Client
	logger *zap.Logger
}

func New(l *zap.Logger, client *auth.Client) (*Middleware, error) {
	return &Middleware{
		client: client,
		logger: l,
	}, nil
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		providedToken, err := extractBearerToken(r)
		if err != nil {
			m.logger.Info("failed to extract bearer token", zap.Error(err))
			render.Render(w, r, responses.ErrUnauthorised())
			return
		}

		token, err := m.client.VerifyIDToken(ctx, providedToken)
		if err != nil {
			m.logger.Info("token provided is invalid", zap.Error(err))
			render.Render(w, r, responses.ErrUnauthorised())
			return
		}

		userId := token.UID
		logger := m.logger.With(zap.String("userId", userId))

		authUser, err := m.client.GetUser(ctx, userId)
		if err != nil {
			logger.Info("could not get user from id in token", zap.Error(err))
			render.Render(w, r, responses.ErrInternalServerError())
			return
		}

		realUser := users.AuthUserRecordToUser(authUser)
		ctx = users.AddUserToContext(ctx, realUser)

		next.ServeHTTP(w, r.Clone(ctx))
	})
}

func extractBearerToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", fmt.Errorf("no Authorization header provided")
	}

	splitHeader := strings.Split(header, " ")
	if len(splitHeader) < 1 {
		return "", fmt.Errorf("invalid Authorization header")
	}
	return splitHeader[1], nil
}
