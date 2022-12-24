package session

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/ryoshindo/webauthn/backend/model"
	"github.com/ryoshindo/webauthn/backend/repository"
)

type contextKey struct {
	name string
}

var (
	accountCtxKey        = &contextKey{"account"}
	sessionCtxKey        = &contextKey{"session"}
	sessionHandlerCtxKey = &contextKey{"session_handler"}
)

func Middleware(accountRepo repository.AccountRepository, sessionRepo repository.SessionRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			srw := &sessionResponseWriter{w, nil}
			sessionHandler := &sessionHandler{accountRepo, sessionRepo, srw}

			ctx := context.WithValue(r.Context(), sessionHandlerCtxKey, sessionHandler)

			c, err := r.Cookie(getCookieKey())
			if err != nil || c == nil {
				next.ServeHTTP(srw, r.WithContext(ctx))
				return
			}

			sess, err := sessionRepo.FindByToken(ctx, c.Value)
			if err != nil {
				next.ServeHTTP(srw, r.WithContext(ctx))
				return
			}

			account, err := accountRepo.FindByID(ctx, sess.AccountID)
			fmt.Println(sess.AccountID, account, err)
			if err != nil {
				next.ServeHTTP(srw, r.WithContext(ctx))
				return
			}

			ctx = context.WithValue(ctx, sessionCtxKey, sess)
			ctx = context.WithValue(ctx, accountCtxKey, account)

			next.ServeHTTP(srw, r.WithContext(ctx))
		})
	}
}

type sessionHandler struct {
	accountRepo repository.AccountRepository
	sessionRepo repository.SessionRepository

	responseWriter *sessionResponseWriter
}

type sessionResponseWriter struct {
	http.ResponseWriter

	session *model.Session
}

func getCookieKey() string {
	if value, ok := os.LookupEnv("SESSION_TOKEN_COOKIE_KEY"); ok {
		return value
	}

	return "webauthn_app_session_token"
}

func (w *sessionResponseWriter) Write(b []byte) (int, error) {
	if w.session != nil {
		http.SetCookie(w, &http.Cookie{
			Name:     getCookieKey(),
			Value:    w.session.Token,
			HttpOnly: true,
		})
	}

	return w.ResponseWriter.Write(b)
}

func AccountFromSession(ctx context.Context) *model.Account {
	account, _ := ctx.Value(accountCtxKey).(*model.Account)
	return account
}

func CreateSession(ctx context.Context, account *model.Account) error {
	handler, _ := ctx.Value(sessionHandlerCtxKey).(*sessionHandler)
	if handler == nil {
		return nil
	}

	return handler.CreateSession(ctx, account)
}

func DeleteSession(ctx context.Context) error {
	handler, _ := ctx.Value(sessionHandlerCtxKey).(*sessionHandler)
	if handler == nil {
		return nil
	}

	session, _ := ctx.Value(sessionCtxKey).(*model.Session)
	if session == nil {
		return nil
	}

	return handler.DeleteSession(ctx, session)
}

func (sh *sessionHandler) CreateSession(ctx context.Context, account *model.Account) error {
	session, err := model.NewSession(account.ID)
	if err != nil {
		return err
	}

	if err := sh.sessionRepo.Create(ctx, session); err != nil {
		return err
	}

	sh.responseWriter.session = session

	return nil
}

func (sh *sessionHandler) DeleteSession(ctx context.Context, session *model.Session) error {
	return sh.sessionRepo.Delete(ctx, session)
}
