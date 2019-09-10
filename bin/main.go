package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/tuxmonteiro/docker-machine-driver-cloudstack"
)

func main() {
	plugin.RegisterDriver(cloudstack.NewDriver("", ""))
}
