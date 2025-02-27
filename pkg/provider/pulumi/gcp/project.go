// Copyright Nitric Pty Ltd.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcp

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectArgs struct {
	ProjectId     string
	ProjectNumber string
}

type Project struct {
	pulumi.ResourceState

	Name     string
	Services []*projects.Service
}

var requiredServices = []string{
	// Enable IAM
	"iam.googleapis.com",
	// Enable cloud run
	"run.googleapis.com",
	// Enable pubsub
	"pubsub.googleapis.com",
	// Enable cloud scheduler
	"cloudscheduler.googleapis.com",
	// Enable cloud scheduler
	"storage.googleapis.com",
	// Enable Compute API (Networking/Load Balancing)
	"compute.googleapis.com",
	// Enable Container Registry API
	"containerregistry.googleapis.com",
	// Enable firestore API
	"firestore.googleapis.com",
	// Enable ApiGateway API
	"apigateway.googleapis.com",
	// Enable SecretManager API
	"secretmanager.googleapis.com",
}

func newProject(ctx *pulumi.Context, name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	res := &Project{
		Name:     name,
		Services: []*projects.Service{},
	}

	err := ctx.RegisterComponentResource("nitric:project:GcpProject", name, res, opts...)
	if err != nil {
		return nil, err
	}

	deps := []pulumi.Resource{}

	for _, serv := range requiredServices {
		s, err := projects.NewService(ctx, serv+"-enabled", &projects.ServiceArgs{
			DisableDependentServices: pulumi.Bool(true),
			DisableOnDestroy:         pulumi.Bool(false),
			Project:                  pulumi.String(args.ProjectId),
			Service:                  pulumi.String(serv),
		})
		if err != nil {
			return nil, err
		}

		res.Services = append(res.Services, s)
		deps = append(deps, s)
	}

	// Add ServiceAccount Token Creator Role to the default pubsub gservice account
	// services-{projectNumber}@gcp-sa-pubsub.iam.gserviceaccount.com
	_, err = projects.NewIAMMember(ctx, "pubsub-token-creator", &projects.IAMMemberArgs{
		Role:    pulumi.String("roles/iam.serviceAccountTokenCreator"),
		Member:  pulumi.Sprintf("serviceAccount:service-%s@gcp-sa-pubsub.iam.gserviceaccount.com", args.ProjectNumber),
		Project: pulumi.String(args.ProjectId),
	}, pulumi.Parent(res),
		pulumi.DependsOn(deps)) // Only create this once the google managed service account is available
	if err != nil {
		return nil, err
	}

	return res, ctx.RegisterResourceOutputs(res, pulumi.Map{
		"name": pulumi.StringPtr(res.Name),
		//"services": res.Services,
	})
}
