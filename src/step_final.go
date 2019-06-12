package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

type finalStep struct{}

type pullRequest struct {
	Url  string `json:"url"`
	Head struct {
		Ref string `json:"ref"`
	} `json:"head"`
}

func (s finalStep) Execute(c *context) bool {

	resp, _ := http.Get("https://api.github.com/repos/sensorario/ff/pulls")
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var pullRequests []pullRequest
	json.Unmarshal(bodyBytes, &pullRequests)
	foo, _ := json.Marshal(pullRequests)
	fmt.Println(string(foo))

	if c.conf.Features.RemoveRemotelyMerged && c.currentBranch() == c.conf.Branches.Historical.Development {

		c.Logger.Info(color.RedString("will remove remotely merged branches"))

		gitCheckoutToDev := &gitCommand{
			c.Logger,
			[]string{"branch", "-a", "--merged"},
			"Cant list all local branches ",
			c.conf,
		}
		output := gitCheckoutToDev.Execute()

		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if strings.Contains(line, "remotes/origin") {

				branchName := strings.Replace(strings.Trim(line, " "), "remotes/origin/", "", 1)
				if branchName != c.conf.Branches.Historical.Development && !strings.Contains(branchName, "HEAD") {

					// skip all branches listed in pullRequests collection

					deleteRemoteBranch := &gitCommand{
						c.Logger,
						[]string{"push", "origin", ":" + strings.Replace(strings.Trim(line, " "), "remotes/origin/", "", 1)},
						"Cant list all local branches ",
						c.conf,
					}

					deleteRemoteBranch.Execute()
				}
			}
		}

	} else {
		c.Logger.Info(color.GreenString("leave remotely merged branches"))
	}

	return false
}

func (s finalStep) Stepname() string {
	return "final"
}
