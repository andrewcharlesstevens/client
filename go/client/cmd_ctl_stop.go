// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

// +build !darwin

package client

import (
	"github.com/keybase/cli"
	"github.com/keybase/client/go/install"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
)

func NewCmdCtlStop(cl *libcmdline.CommandLine, g *libkb.GlobalContext) cli.Command {
	return cli.Command{
		Name:  "stop",
		Usage: "Stop Keybase",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "shutdown",
				Usage: "Only shutdown the background service",
			},
		},
		Action: func(c *cli.Context) {
			cl.ChooseCommand(newCmdCtlStop(g), "stop", c)
			cl.SetForkCmd(libcmdline.NoFork)
			cl.SetNoStandalone()
		},
	}
}

func newCmdCtlStop(g *libkb.GlobalContext) *CmdCtlStop {
	return &CmdCtlStop{
		Contextified: libkb.NewContextified(g),
	}
}

type CmdCtlStop struct {
	libkb.Contextified
	shutdown bool
}

func (s *CmdCtlStop) ParseArgv(ctx *cli.Context) error {
	s.shutdown = ctx.Bool("shutdown")
	return nil
}

func (s *CmdCtlStop) Run() (err error) {
	mctx := libkb.NewMetaContextTODO(s.G())
	if !s.shutdown {
		install.StopAllButService(mctx, keybase1.ExitCode_OK)
	}
	cli, err := GetCtlClient(s.G())
	if err != nil {
		return err
	}
	return cli.StopService(mctx.Ctx(), keybase1.StopServiceArg{ExitCode: keybase1.ExitCode_OK})
}

func (s *CmdCtlStop) GetUsage() libkb.Usage {
	return libkb.Usage{}
}
