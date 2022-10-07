package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	// "os"

	"github.com/fatih/color"
)

type finalStep struct{}

type pullRequest struct {
	Url  string `json:"url"`
	Head struct {
		Ref string `json:"ref"`
	} `json:"head"`
}

// Execute is called whenever final step must be called.
func (s finalStep) Execute(c *context) bool {

	resp, err := http.Get("https://api.github.com/repos/sensorario/ff/pulls")

    if err != nil {
        if !strings.Contains(err.Error(), "dial tcp") && !strings.Contains(err.Error(), "no such host") {
            c.Logger.Warning(color.GreenString("Probabilmente il computer non e' connesso alla rete"))
        } else {
            c.Logger.Warning(color.GreenString("Unknown connection error"))
            c.Logger.Warning(color.GreenString(err.Error()))
        }

        return false
    } else {
        defer resp.Body.Close()
    }

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var pullRequests []pullRequest
	json.Unmarshal(bodyBytes, &pullRequests)

	c.Logger.Info(color.GreenString("info ... "))

	if c.conf.Features.RemoveRemotelyMerged {
		c.Logger.Info(color.RedString("I will remove remotely merged branches"))

		c.Logger.Info(color.RedString("I will prune local branches without remote ones"))
		pruneEverything := &gitCommand{
			c.Logger,
			[]string{"fetch", "origin", "--prune"},
			"Cant prune local branches",
			c.conf,
		}
		pruneEverything.Execute()

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
                c.Logger.Info(color.GreenString("branch found "+branchName))
				if branchName != c.conf.Branches.Historical.Development && !strings.Contains(branchName, "HEAD") {
					removableBranch := strings.Replace(strings.Trim(line, " "), "remotes/origin/", "", 1)
                    c.Logger.Info(color.GreenString("removable branch found "+removableBranch))
					isDifferentFromAllPR := true
					for _, v := range pullRequests {
						if v.Head.Ref == removableBranch {
							isDifferentFromAllPR = false
						}
					}
					if isDifferentFromAllPR {
						deleteRemoteBranch := &gitCommand{
							c.Logger,
							[]string{"push", "origin", ":" + removableBranch},
							strings.Join([]string{"Cant delete branches:", "git","push", "origin", ":" + removableBranch}, " "),
							c.conf,
						}
						outcome := deleteRemoteBranch.Execute()

                        fmt.Println("Result:")
                        fmt.Println(outcome)
					}
				} else {
                    c.Logger.Info(color.GreenString("branch " + branchName))
                }
			}
		}
	} else {
		c.Logger.Info(color.GreenString("leave remotely merged branches"))
	}

	c.Logger.Info(color.GreenString("info ... "))

	return false
}

// Stepname contains the name of current step.
func (s finalStep) Stepname() string {
	return "final"
}
