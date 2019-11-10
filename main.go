package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type mailReq struct {
	From    string `query:"from"`
	To      string `query:"to"`
	Subject string `query:"subject"`
	Body    string `query:"body"`
	File    string `query:"file"`
}

func callSendmail(req mailReq) error {
	var line = fmt.Sprintf("From: %s\\nTo: %s\\nSubject: %s\\n\\n%s",
		req.From, req.To, req.Subject, req.Body)
	var script = fmt.Sprintf("echo -e '%s'", line)

	fmt.Println(script)
	if len(req.File) > 0 {
		if _, err := os.Stat(req.File); err == nil {
			script = fmt.Sprintf("%s ; uuencode %s %s", script,
				req.File, path.Base(req.File))
		} else {
			fmt.Println(err.Error())
		}
	}

	fmt.Println(script)
	var c1 = exec.Command("/bin/sh", "-c", script)
	var c2 = exec.Command("sendmail", req.To)
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
	return nil
}

func callSendmailMutt(req mailReq) error {

	var args = []string{"-s", req.Subject}
	if len(req.File) > 0 {
		if _, err := os.Stat(req.File); err == nil {
			args = append(args, []string{"-a", req.File}...)
		} else {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(args)
	args = append(args, []string{"--", req.To}...)
	var c1 = exec.Command("echo", "-e", req.Body)
	var c2 = exec.Command("mutt", args...)
	c2.Env = os.Environ()
	c2.Env = append(c2.Env, fmt.Sprintf("EMAIL='%s'", req.From))
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

	var e = echo.New()
	e.Use(middleware.Recover())

	var apiGroup = e.Group("/api")

	// GET HTTP endpoint to send emails via sendmail
	apiGroup.GET("/sendmail", func(c echo.Context) error {

		var req = mailReq{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest,
				map[string]string{
					"status": "error",
					"desc":   err.Error(),
				},
			)
		}
		callSendmail(req)
		return c.JSON(http.StatusOK,
			map[string]string{
				"status": "ok",
			},
		)
	})

	// GET HTTP endpoint to send emails via mutt
	apiGroup.GET("/sendmutt", func(c echo.Context) error {
		var req = mailReq{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest,
				map[string]string{
					"status": "error",
					"desc":   err.Error(),
				},
			)
		}
		callSendmailMutt(req)
		return c.JSON(http.StatusOK,
			map[string]string{
				"status": "ok",
			},
		)
	})

	if err := e.Start(serverAddr); err != nil {
		panic(err.Error())
	}

}
