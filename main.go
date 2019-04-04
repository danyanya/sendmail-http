package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type mailReq struct {
	From    string `query:"from"`
	To      string `query:"to"`
	Subject string `query:"subject"`
	Body    string `query:"body"`
}

func callSendmail(req mailReq) error {
	line := fmt.Sprintf("From: %s\\nTo: %s\\nSubject: %s\\n\\n%s",
		req.From, req.To, req.Subject, req.Body)
	c1 := exec.Command("echo", "-e", line)
	c2 := exec.Command("sendmail", req.To)
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
	return nil
}

func main() {

	var serverAddr = os.Getenv("SERVER_ADDR")
	if len(serverAddr) == 0 {
		panic("Server address must not be empty")
	}

	e := echo.New()
	e.Use(middleware.Recover())

	api_group := e.Group("/api")

	api_group.GET("/sendmail", func(c echo.Context) error {

		req := mailReq{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"status": "error",
				"desc":   err.Error(),
			})
		}
		callSendmail(req)
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	err := e.Start(serverAddr)
	if err != nil {
		panic(err.Error())
	}

}
