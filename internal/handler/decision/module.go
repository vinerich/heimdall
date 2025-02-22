// Copyright 2022 Dimitrij Drus <dadrus@gmx.de>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package decision

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"go.uber.org/fx"

	"github.com/dadrus/heimdall/internal/config"
	"github.com/dadrus/heimdall/internal/handler/listener"
)

var Module = fx.Options( // nolint: gochecknoglobals
	fx.Provide(fx.Annotated{Name: "decision", Target: newApp}),
	fx.Invoke(
		newHandler,
		registerHooks,
	),
)

type hooksArgs struct {
	fx.In

	Lifecycle fx.Lifecycle
	Config    *config.Configuration
	Logger    zerolog.Logger
	App       *fiber.App `name:"decision"`
}

func registerHooks(args hooksArgs) {
	ln, err := listener.New(args.App.Config().Network, args.Config.Serve.Decision)
	if err != nil {
		args.Logger.Fatal().Err(err).Msg("Could not create listener for the Decision service")

		return
	}

	args.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					args.Logger.Info().Str("_address", ln.Addr().String()).Msg("Decision service starts listening")

					if err = args.App.Listener(ln); err != nil {
						args.Logger.Fatal().Err(err).Msg("Could not start Decision service")
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				args.Logger.Info().Msg("Tearing down Decision service")

				return args.App.Shutdown()
			},
		},
	)
}
