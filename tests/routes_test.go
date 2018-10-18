package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/astaxie/beego"
	"github.com/cdvr1993/deployment-manager/models"
	"github.com/cdvr1993/deployment-manager/routers"
	"github.com/cdvr1993/deployment-manager/tests/mocks/services"
	. "github.com/smartystreets/goconvey/convey"
)

type MethodTester struct {
	Body   string
	Method string
	Path   string
	Result interface{}
}

func TestEndpointsAreWorking(t *testing.T) {
	beego.BeeApp = beego.NewApp()

	user := models.User{
		Id:    1,
		Name:  "Cristian",
		Email: "example@example.com",
	}

	routers.InitRouter(routers.ServiceManager{
		UserService: services.NewUserServiceMock(services.UserServiceMethods{
			AddUser: func(u *models.User) {
				u.Id = user.Id
			},
			GetUser: func(e string) models.User {
				return user
			},
		}),
	})

	tests := []MethodTester{
		MethodTester{
			Method: "POST",
			Path:   "/v1/user",
			Body: fmt.Sprintf(
				`{"name": "%s", "email": "%s"}`, user.Name, user.Email,
			),
			Result: user,
		},
		MethodTester{
			Method: "GET",
			Path:   fmt.Sprintf("/v1/user/%s", user.Email),
			Result: user,
		},
	}

	for _, tt := range tests {
		r, _ := http.NewRequest(tt.Method, tt.Path, strings.NewReader(tt.Body))
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		title := fmt.Sprintf("Subject: Test Endpoint '%s - %s'\n", tt.Method, tt.Path)
		Convey(title, t, func() {
			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
			Convey("The result should be of proper type", func() {
				switch v := tt.Result.(type) {
				case models.User:
					var response map[string]models.User
					json.Unmarshal(w.Body.Bytes(), &response)
					So(response["data"], ShouldNotEqual, v)
				default:
					t.Fatalf("Missing case for: '%s'", v)
				}
			})
		})
	}
}
