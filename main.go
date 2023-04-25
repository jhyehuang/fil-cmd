/*
Copyright © 2021 NAME HERE rickiey@qq.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"git.sxxfuture.net/filfi/letsfil/fil-data/cmd"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	//utils.LoadEnv()

	app := &cli.App{}

	app.Name = "fil-data"
	app.Usage = "manager service for filbase cluster"
	app.Version = "alpha"
	app.Authors = []*cli.Author{
		{Name: "Huangzhijie", Email: "huangzhijie@sxxfuture.net"},
	}

	app.EnableBashCompletion = true
	app.Writer = os.Stderr
	app.Commands = []*cli.Command{
		&cmd.RunCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	return

}
