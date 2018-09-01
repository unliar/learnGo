package account

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	ASV "github.com/unliar/proto/account"
	"strconv"
)

var (
	AccountSVService ASV.AccountSVService
)

func init() {
	AccountSVService = ASV.NewAccountSVService("unliar-account", client.DefaultClient)
}

func GetUserBase(c *gin.Context) {
	var err error
	uid := c.Param("uid")
	UID, err := strconv.ParseInt(uid, 10, 64)
	resp, err := AccountSVService.GetUserBase(context.TODO(), &ASV.UserId{
		UID: UID,
	})
	fmt.Println(resp)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, resp)
}

func GetHealthStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  200,
		"message": "api server ok",
	})
}