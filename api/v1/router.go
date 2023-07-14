package v1

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leantech/school-system-api/api/middleware"
	"github.com/leantech/school-system-api/api/swagger"
	"github.com/leantech/school-system-api/api/v1/course"
	"github.com/leantech/school-system-api/api/v1/health"
	"github.com/leantech/school-system-api/api/v1/student"
	"github.com/leantech/school-system-api/api/v1/user"
	"github.com/leantech/school-system-api/config"
	"github.com/leantech/school-system-api/db/mysql"
	"github.com/leantech/school-system-api/model"
	"golang.org/x/crypto/bcrypt"
)

type Option struct {
	DB *sql.DB
}

func Register(g *echo.Group, opts Option) {

	env := config.GetEnv()
	doc := env.Doc
	if doc.Enabled {
		swagger.Register(swagger.Options{
			Title:       doc.Title,
			Description: doc.Description,
			Version:     doc.Version,
			BasePath:    env.Server.BasePath,
			Group:       g.Group("/swagger"),
		})
	}

	g.GET("/health", health.Handle)

	userRoute(g, opts)
	studentRoute(g, opts)
	courseRoute(g, opts)
}

func userRoute(g *echo.Group, opts Option) {
	env := config.GetEnv()

	userRepo := mysql.NewUserRepository(opts.DB)
	userHandler := user.NewHandler(userRepo, bcrypt.GenerateFromPassword, bcrypt.CompareHashAndPassword, env.Authorization.Secret)
	userCreate := middleware.NewController(userHandler.Create, http.StatusCreated, new(model.User))
	userLogin := middleware.NewController(userHandler.Login, http.StatusCreated, new(model.LoginRequest))

	userGroup := g.Group("/user")
	userGroup.POST("", userCreate.Handle, middleware.CheckRole("admin"))
	userGroup.POST("/login", userLogin.Handle)
}

func studentRoute(g *echo.Group, opts Option) {
	studentRepo := mysql.NewStudentRepository(opts.DB)
	studentHandler := student.NewHandler(studentRepo)
	studentCreate := middleware.NewController(studentHandler.Create, http.StatusCreated, new(model.Student))
	studentUpdate := middleware.NewController(studentHandler.Update, http.StatusOK, new(model.UpdateStudentRequest))
	studentDelete := middleware.NewController(studentHandler.Delete, http.StatusOK, new(model.DeleteStudentRequest))

	studentGroup := g.Group("/student")
	studentGroup.POST("", studentCreate.Handle)
	studentGroup.PUT("/:id", studentUpdate.Handle)
	studentGroup.DELETE("/:id", studentDelete.Handle)
}

func courseRoute(g *echo.Group, opts Option) {
	courseRepo := mysql.NewCourseRepository(opts.DB, mysql.NewStudentRepository(opts.DB))

	courseHandler := course.NewHandler(courseRepo)
	courseCreate := middleware.NewController(courseHandler.Create, http.StatusCreated, new(model.Course))
	courseEnroll := middleware.NewController(courseHandler.EnrollStudent, http.StatusCreated, new(model.EnrollStudentRequest))

	courseGroup := g.Group("/course")
	courseGroup.POST("", courseCreate.Handle, middleware.CheckRole("admin"))
	courseGroup.POST("/enroll", courseEnroll.Handle, middleware.CheckRole("admin"))
}
