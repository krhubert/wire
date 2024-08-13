// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build wireinject
// +build wireinject

package main

import (
	"bytes"
	"strings"

	"github.com/google/wire"
)

func injectDifferentInterfaceImplementation() Foo {
	panic(
		wire.Build(
			wire.Union(
				provideFooStrings,
				wire.Bind(new(Foo), new(*strings.Reader)),

				provideFooBytes,
				wire.Bind(new(Foo), new(*bytes.Buffer)),
			),
		),
	)
}

func injectDifferentInterfaceProvider() Foo {
	panic(
		wire.Build(
			wire.Union(
				provideFooStrings,
				wire.Bind(new(Foo), new(*strings.Reader)),

				provideFooStrings2,
				wire.Bind(new(Foo), new(*strings.Reader)),
			),
		),
	)
}
