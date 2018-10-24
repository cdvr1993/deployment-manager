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

	group := models.Group{
		Id:   1,
		Name: "MyGroup",
	}

	routers.InitRouter(routers.ServiceManager{
		GroupService: services.NewGroupServiceMock(services.GroupServiceMethods{
			CreateGroup: func(g *models.Group) {
				g.Id = group.Id
			},
			GetAllGroups: func() []*models.Group {
				return []*models.Group{&group}
			},
			GetGroupByName: func(string) models.Group {
				return group
			},
		}),
		UserService: services.NewUserServiceMock(services.UserServiceMethods{
			AddUser: func(u *models.User) {
				u.Id = user.Id
			},
			GetAll: func() []*models.User {
				return []*models.User{&user}
			},
			GetUserByEmail: func(e string) models.User {
				return user
			},
		}),
	})

	tests := []MethodTester{
		MethodTester{
			Method: "GET",
			Path:   "/v1/user",
			Result: []*models.User{&user},
		},
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
		MethodTester{
			Method: "DELETE",
			Path:   fmt.Sprintf("/v1/user/%d", user.Id),
		},
		MethodTester{
			Method: "PUT",
			Path:   fmt.Sprintf("/v1/user/%d", user.Id),
		},
		MethodTester{
			Method: "GET",
			Path:   fmt.Sprintf("/v1/group/%s", group.Name),
			Result: group,
		},
		MethodTester{
			Method: "GET",
			Path:   "/v1/group",
			Result: []*models.Group{&group},
		},
		MethodTester{
			Method: "POST",
			Path:   "/v1/group",
			Body:   fmt.Sprintf(`{"name": "%s"}`, group.Name),
			Result: group,
		},
		MethodTester{
			Method: "PUT",
			Path:   fmt.Sprintf("/v1/group/%d", group.Id),
			Body:   fmt.Sprintf(`{"name": "%s"}`, group.Name),
			Result: "updated successfully",
		},
		MethodTester{
			Method: "DELETE",
			Path:   fmt.Sprintf("/v1/group/%d", group.Id),
			Result: "deleted successfully",
		},
		MethodTester{
			Method: "POST",
			Path:   fmt.Sprintf("/v1/group/%d/member", group.Id),
			Body:   fmt.Sprintf(`{"id": "%d"}`, user.Id),
			Result: "added successfully",
		},
		MethodTester{
			Method: "DELETE",
			Path:   fmt.Sprintf("/v1/group/%d/member/%d", group.Id, user.Id),
			Result: "removed successfully",
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

			// Update and delete response usually doesn't matter, just the status
			if tt.Result != nil {
				Convey("The result should be of proper type", func() {
					switch v := tt.Result.(type) {
					case models.User:
						var response map[string]models.User
						json.Unmarshal(w.Body.Bytes(), &response)
						So(response["Data"], ShouldResemble, v)
					case []*models.User:
						var response map[string][]models.User
						json.Unmarshal(w.Body.Bytes(), &response)
						So(response["Data"][0], ShouldResemble, *(v[0]))
					case models.Group:
						var response map[string]models.Group
						json.Unmarshal(w.Body.Bytes(), &response)
						So(response["Data"], ShouldResemble, v)
					case []*models.Group:
						var response map[string][]models.Group
						json.Unmarshal(w.Body.Bytes(), &response)
						So(response["Data"][0], ShouldResemble, *(v[0]))
					case string:
						var response map[string]string
						json.Unmarshal(w.Body.Bytes(), &response)
						So(response["Data"], ShouldContainSubstring, v)
					default:
						t.Fatalf("Missing case for: '%s'", v)
					}
				})
			}
		})
	}
}
