package intern

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ServiceManage struct {
	DB *gorm.DB
}

func (sm *ServiceManage) ServeAPI(ctx context.Context) error {
	e := echo.New()
	e.GET("/debug", serveProfile)
	e.GET("/GetBlockReward", sm.GetBlockReward)
	return e.Start(":33004")
}
