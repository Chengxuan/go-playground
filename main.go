// Copyright © 2022 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/kaleido-io/kaleido-sdk-go/kaleido"
)

func main() {
	kc := kaleido.NewClient("https://console-ko.kaleido.io/api/v1", "kaleido saas API key")

	cms := []kaleido.Consortium{}
	if _, err := kc.ListConsortium(&cms); err != nil {
		panic(err)
	}
	cs := kaleido.Consortium{}
	kc.GetConsortium(cms[0].ID, &cs)
	env := []kaleido.Environment{}
	if _, err := kc.ListEnvironments(cms[0].ID, &env); err != nil {
		panic(err)
	}
	ns := []kaleido.Node{}
	if _, err := kc.ListNodes(cms[0].ID, env[0].ID, &ns); err != nil {
		panic(err)
	}

	n := kaleido.Node{}
	kc.GetNode(cms[0].ID, env[0].ID, ns[0].ID, &n)
	ac := []kaleido.AppCreds{}
	kc.ListAppCreds(cms[0].ID, env[0].ID, &ac)
	fmt.Println(ns)
}
