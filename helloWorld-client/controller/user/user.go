package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	helloWorldSrv "github.com/wa1kman999/helloworld/proto"
	howAreYou "github.com/wa1kman999/howAreYou/proto"
	"go-micro.dev/v4/client"
	"net/http"
)

type Param struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

//type traceIDCtx   struct{}

func Login(c *gin.Context) {
	// 接收参数
	var param Param
	err := c.ShouldBindJSON(&param)
	if err != nil {
		fmt.Printf("参数校验错误:%v", err)
	}
	//客户端
	cli := client.NewClient()
	userSRV := helloWorldSrv.NewHelloWorldService("helloWorld", cli)
	resp, err := userSRV.Login(c.Request.Context(), &helloWorldSrv.UserLoginRequest{
		Account:  param.Account,
		Password: param.Password,
	})

	howAreYou.NewHowAreYouService("")

	if err != nil {
		fmt.Printf("打印err:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":     resp.Code,
			"username": resp.UserName,
			"msg":      resp.Msg,
		})
	}
}
