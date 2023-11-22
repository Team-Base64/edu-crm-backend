package delivery

import (
	"context"
	"errors"
	"log"
	conf "main/config"
	e "main/domain/errors"
	"main/domain/model"
	uc "main/usecase"
	"net/http"
	"strings"
)

type AuthenticationMiddlewareInterface interface {
	CheckAuthMiddleware(next http.Handler) http.Handler
}

type AuthMiddleware struct {
	usecase uc.UsecaseInterface
}

func NewAuthMiddleware(uc uc.UsecaseInterface) AuthenticationMiddlewareInterface {
	return &AuthMiddleware{
		usecase: uc,
	}
}

type KeyUserdata struct {
	key string
}

func WithUser(ctx context.Context, user *model.TeacherDB) context.Context {
	return context.WithValue(ctx, KeyUserdata{"userdata"}, user)

}

func (amw *AuthMiddleware) CheckAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		if r.RequestURI == conf.PathLogin || strings.Contains(r.RequestURI, conf.PathDocs) {
			next.ServeHTTP(w, r)
		} else {
			if r.Method == http.MethodOptions {
				return
			}
			session, err := r.Cookie("session_id")
			if err == http.ErrNoCookie {
				log.Println(e.StacktraceError(err))
				returnErrorJSON(w, e.ErrUnauthorized401)
				return
			}
			usLogin, err := amw.usecase.CheckSession(session.Value)
			if err != nil {
				log.Println(e.StacktraceError(err))
				returnErrorJSON(w, e.ErrUnauthorized401)
				return
			}

			user, err := amw.usecase.GetTeacherProfileByLogin(usLogin)
			// usLogin := "test1"
			// user, err := amw.usecase.GetTeacherProfileByLogin(usLogin)
			if user.Name == "" {
				log.Println(e.StacktraceError(errors.New("no user: " + usLogin)))
				returnErrorJSON(w, e.ErrUnauthorized401)
				return
			}
			if err != nil {
				log.Println(e.StacktraceError(err))
				returnErrorJSON(w, e.ErrServerError500)
				return
			}
			next.ServeHTTP(w, r.WithContext(WithUser(r.Context(), user)))
		}

	})
}
