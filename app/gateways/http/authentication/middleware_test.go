package authentication

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication/usecase"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const (
	authenticated = "authenticated"
	idNotPresent  = "id not present"
	notString     = "not string"
	wrongReqID    = "wrong reqId"
	accountID     = "642e0d44-9792-4d6f-9a04-b40186dddbef"
)

func TestHandler_Authenticate(t *testing.T) {
	t.Run("should authenticate and send account id", func(t *testing.T) {
		header := http.Header{
			"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjk0ODI3MzQsImlhdCI6MTYyOTQ4MTgzNCwiaWQiOiJkY2E3NzQ2ZC04YWU1LTQ3Y2UtOGExYi0yOGFhOTFhZjkyNWQiLCJpc3MiOiJKV1QiLCJuYW1lIjoiUGF0cmljaWEiLCJzdWIiOiI2NDJlMGQ0NC05NzkyLTRkNmYtOWEwNC1iNDAxODZkZGRiZWYifQ.xqFGOp_3jatWFPLAxe9WtvRSITV1FgQzPAnePwXA2EE"},
		}
		handler := createMiddlewareFakeHandler(nil)
		nextHandlerFunc := createFakeHandlerFunc()
		request, _ := http.NewRequest(http.MethodGet, "/transfers", nil)
		response := httptest.NewRecorder()
		request.Header = header
		authenticationMiddleware := handler.Authenticate(nextHandlerFunc)
		authenticationMiddleware(response, request)

		got := response.Body.String()
		expected := authenticated

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))

	})

	t.Run("should return 400 and error when authorization header is empty", func(t *testing.T) {
		badHeader := http.Header{
			"Authorization": []string{},
		}
		handler := createMiddlewareFakeHandler(nil)
		nextHandlerFunc := createFakeHandlerFunc()
		request, _ := http.NewRequest(http.MethodGet, "/transfers", nil)
		response := httptest.NewRecorder()
		request.Header = badHeader
		authenticationMiddleware := handler.Authenticate(nextHandlerFunc)
		authenticationMiddleware(response, request)

		got := response.Body.String()
		expected := `{"error":"empty authorization header"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))

	})

	t.Run("should return 400 and error when access token is empty", func(t *testing.T) {
		badHeader := http.Header{
			"Authorization": []string{"Bearer"},
		}
		handler := createMiddlewareFakeHandler(nil)
		nextHandlerFunc := createFakeHandlerFunc()
		request, _ := http.NewRequest(http.MethodGet, "/transfers", nil)
		response := httptest.NewRecorder()
		request.Header = badHeader
		authenticationMiddleware := handler.Authenticate(nextHandlerFunc)
		authenticationMiddleware(response, request)

		got := response.Body.String()
		expected := `{"error":"empty token"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
	})

	t.Run("should return 400 and error when the authentication method is wrong", func(t *testing.T) {
		badHeader := http.Header{
			"Authorization": []string{"auth "},
		}
		handler := createMiddlewareFakeHandler(nil)
		nextHandlerFunc := createFakeHandlerFunc()
		request, _ := http.NewRequest(http.MethodGet, "/transfers", nil)
		response := httptest.NewRecorder()
		request.Header = badHeader
		authenticationMiddleware := handler.Authenticate(nextHandlerFunc)
		authenticationMiddleware(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid auth method"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
	})

	t.Run("should return 403 and error when token is invalid", func(t *testing.T) {
		header := http.Header{
			"Authorization": []string{"Bearer "},
		}
		err := usecase.ErrInvalidToken
		handler := createMiddlewareFakeHandler(err)
		nextHandlerFunc := createFakeHandlerFunc()
		request, _ := http.NewRequest(http.MethodGet, "/transfers", nil)
		request.Header = header
		response := httptest.NewRecorder()
		authenticationMiddleware := handler.Authenticate(nextHandlerFunc)
		authenticationMiddleware(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid token"}`

		assert.Equal(t, http.StatusForbidden, response.Code)
		assert.Equal(t, expected, got)
	})
}

func createFakeHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(authenticated))
		return
	}
}

func createMiddlewareFakeHandler(err error) Handler {
	if err != nil {
		return NewHandler(&authentication.UseCaseMock{
			ValidatesTokenFunc: func(tokenString string) (entities.Token, error) {
				return entities.Token{}, err
			},
		}, nil)
	}
	return NewHandler(&authentication.UseCaseMock{
		ValidatesTokenFunc: func(tokenString string) (entities.Token, error) {
			return entities.Token{
				ID:        "0de9ec06-0ca4-4583-9ddc-585ec65a8c29",
				Name:      "Olive Oyl",
				Subject:   accountID,
				Issuer:    "JWT",
				IssuedAt:  time.Now(),
				ExpiredAt: time.Now().Add(time.Minute * 15),
			}, nil
		},
	}, nil)
}
