package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/astaxie/beego"
	"github.com/cdvr1993/deployment-manager/models"
	"github.com/cdvr1993/deployment-manager/routers"
	svcs "github.com/cdvr1993/deployment-manager/services"
	"github.com/cdvr1993/deployment-manager/tests/mocks/services"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	beego.SetLevel(beego.LevelEmergency)
}

func TestRecoverFromPanicPredefinedError(t *testing.T) {
	beego.BeeApp = beego.NewApp()

	email := "example@example.com"
	err := svcs.ErrorUserNotFound(email)

	routers.InitRouter(routers.ServiceManager{
		UserService: services.NewUserServiceMock(services.UserServiceMethods{
			GetUserByEmail: func(e string) (*models.User, error) {
				return nil, err
			},
		}),
	})

	r, _ := http.NewRequest("GET", "/v1/user/example@example.com", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	title := "Recover from panic on GET - /v1/user/example@example.com\n"
	Convey(title, t, func() {
		Convey(fmt.Sprintf("Status Code Should Be %d", err.Status()), func() {
			So(w.Code, ShouldEqual, err.Status())
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("The response must have a list of errors", func() {
			var response map[string][]string
			json.Unmarshal(w.Body.Bytes(), &response)

			for _, e := range response["errors"] {
				if strings.Contains(e, email) {
					return
				}
			}

			t.Fatalf("The response '%s' doesn't contain the error", response)
		})
	})
}

func TestRecoverFromPanicUnknownError(t *testing.T) {
	beego.BeeApp = beego.NewApp()

	errorMsg := "Any error"
	routers.InitRouter(routers.ServiceManager{
		UserService: services.NewUserServiceMock(services.UserServiceMethods{
			GetUserByEmail: func(e string) (*models.User, error) {
				panic(errors.New(errorMsg))
			},
		}),
	})

	r, _ := http.NewRequest("GET", "/v1/user/example@example.com", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	title := "Recover from panic on GET - /v1/user/example@example.com\n"
	Convey(title, t, func() {
		Convey("Status Code Should Be 500", func() {
			So(w.Code, ShouldEqual, 500)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("The response must have a list of errors", func() {
			var response map[string][]string
			json.Unmarshal(w.Body.Bytes(), &response)
			So(response["errors"], ShouldNotEqual, []string{errorMsg})
		})
	})
}
