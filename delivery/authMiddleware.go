package delivery

import (
	"context"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	uc "main/usecase"
	"net/http"
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
		session, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			log.Println("no session")
			returnErrorJSON(w, e.ErrUnauthorized401)
			return
		}
		usLogin, err := amw.usecase.CheckSession(session.Value)
		if err != nil {
			log.Println("no session2")
			returnErrorJSON(w, e.ErrUnauthorized401)
			return
		}

		user, err := amw.usecase.GetTeacherProfileByLogin(usLogin)
		if user.Name == "" {
			log.Println("no user in db")
			returnErrorJSON(w, e.ErrUnauthorized401)
			return
		}
		if err != nil {
			log.Println(e.StacktraceError(err))
			returnErrorJSON(w, e.ErrServerError500)
			return
		}
		next.ServeHTTP(w, r.WithContext(WithUser(r.Context(), user)))
	})
}
