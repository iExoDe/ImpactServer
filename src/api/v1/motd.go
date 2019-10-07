package v1

import (
	"github.com/ImpactDevelopment/ImpactServer/src/util"
	"github.com/labstack/echo"
	"net/url"
)

const motdURL = "https://impactdevelopment.github.io/Resources/data/motd.txt"

func getMotd(c echo.Context) error {
	target, err := url.Parse(motdURL)
	if err != nil {
		return err //wtf
	}
	util.Proxy(c, target)
	return nil
}
