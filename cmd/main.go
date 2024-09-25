package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v3"
	"github.com/lclpedro/scaffold-golang-fiber/configs"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/repositories"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/services"
	"github.com/lclpedro/scaffold-golang-fiber/internal/scaffold/views"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
	"go.uber.org/fx"
)

func newDatabase() (mysql.Connection, error) {
	databaseConfig := mysql.GetDatabaseConfiguration()
	read := mysql.InitMySQLConnection(databaseConfig[mysql.ReadOperation], mysql.ReadOperation)
	write := mysql.InitMySQLConnection(databaseConfig[mysql.WriteOperation], mysql.WriteOperation)
	return mysql.NewConnection(write, read)
}

func newUnitOfWork(conn mysql.Connection) mysql.UnitOfWorkInterface {
	return mysql.NewUnitOfWork(conn)
}

func newRepositories(uowInstance mysql.UnitOfWorkInterface, connMysql mysql.Connection) {
	repositories.RegistryRepositories(uowInstance, connMysql)
}

func newAllServices(uowInstance mysql.UnitOfWorkInterface) *services.AllServices {
	return services.NewAllServices(uowInstance)
}

func newViews(app *fiber.App, allServices *services.AllServices) {
	views.NewAllHandlerViews(app, allServices)
}

func newServer() *fiber.App {
	return fiber.New()
}

func startServer(app *fiber.App, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go app.Listen(":8080")
			return nil
		},
		OnStop: func(context.Context) error {
			return app.Shutdown()
		}},
	)
}

func main() {
	configs.InitConfigs()

	fx.New(
		fx.Provide(
			newDatabase,
			newUnitOfWork,
			newServer,
			newAllServices,
		),
		fx.Invoke(
			startServer,
			newRepositories,
			newViews,
		),
	).Run()
}
