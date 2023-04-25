package cmd

import (
	"fmt"
	"github.com/jhyehuang/fil-cmd/intern"
	"github.com/jhyehuang/fil-cmd/intern/models"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var RunCommand = cli.Command{
	Name:      "run",
	Aliases:   []string{"run"},
	Usage:     "start fil-data run service",
	UsageText: `fil-data run`,
	//Flags: []cli.Flag{
	//	&cli.StringFlag{
	//		Name:    "pg-uri",
	//		EnvVars: []string{"PG_URI"},
	//		//Required: true,
	//	},
	//},
	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context

		var sm = &intern.ServiceManage{}
		pg_host := os.Getenv("PG_HOST")
		pg_port := os.Getenv("PG_PORT")
		pg_user := os.Getenv("PG_USER")
		pg_password := os.Getenv("PG_PASSWORD")
		pg_db := os.Getenv("PG_DBNAME")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", pg_host, pg_user, pg_password, pg_db, pg_port)
		pd_db, err := initDatabase(dsn)
		if err != nil {
			return err
		}
		// 绑定models
		err = setupDatabase(pd_db)
		if err != nil {
			return err
		}

		sm.DB = pd_db

		return sm.ServeAPI(ctx)

	},
}

func initDatabase(dsn string) (*gorm.DB, error) {
	// https://github.com/jackc/pgx
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}

func setupDatabase(dbsetupDatabase *gorm.DB) error {

	if _, err := models.SetupDatabase(dbsetupDatabase); err != nil {
		return err
	}
	return nil
}
