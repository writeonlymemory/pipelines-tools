// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package query provides a sub-tool for querying running pipelines.
package query

import (
	"context"
	"flag"
	"fmt"

	genomics "google.golang.org/api/genomics/v2alpha1"
)

var (
	flags flag.FlagSet

	filter = flags.String("filter", "", "the query filter")
)

func Invoke(ctx context.Context, service *genomics.Service, project string, arguments []string) error {
	flags.Parse(arguments)

	path := fmt.Sprintf("projects/%s/operations", project)
	call := service.Projects.Operations.List(path).Context(ctx)
	if *filter != "" {
		call = call.Filter(*filter)
	}
	resp, err := call.Do()
	if err != nil {
		return err
	}

	for _, operation := range resp.Operations {
		fmt.Println(operation.Name)
	}
	return nil
}
