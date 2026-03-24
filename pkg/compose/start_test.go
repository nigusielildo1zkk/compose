/*
   Copyright 2020 Docker Compose CLI authors

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

package compose

import (
	"testing"

	"github.com/compose-spec/compose-go/v2/types"
	"gotest.tools/v3/assert"
)

func TestRestartsOnExit0(t *testing.T) {
	tests := []struct {
		name     string
		service  types.ServiceConfig
		expected bool
	}{
		{
			name:     "no restart policy",
			service:  types.ServiceConfig{Name: "init"},
			expected: false,
		},
		{
			name:     "restart no",
			service:  types.ServiceConfig{Name: "init", Restart: types.RestartPolicyNo},
			expected: false,
		},
		{
			name:     "restart on-failure",
			service:  types.ServiceConfig{Name: "init", Restart: types.RestartPolicyOnFailure},
			expected: false,
		},
		{
			name:     "restart always",
			service:  types.ServiceConfig{Name: "web", Restart: types.RestartPolicyAlways},
			expected: true,
		},
		{
			name:     "restart unless-stopped",
			service:  types.ServiceConfig{Name: "web", Restart: types.RestartPolicyUnlessStopped},
			expected: true,
		},
		{
			name: "deploy restart policy condition any",
			service: types.ServiceConfig{
				Name: "web",
				Deploy: &types.DeployConfig{
					RestartPolicy: &types.RestartPolicy{Condition: "any"},
				},
			},
			expected: true,
		},
		{
			name: "deploy restart policy condition on-failure",
			service: types.ServiceConfig{
				Name: "web",
				Deploy: &types.DeployConfig{
					RestartPolicy: &types.RestartPolicy{Condition: "on-failure"},
				},
			},
			expected: false,
		},
		{
			name: "deploy restart policy condition none",
			service: types.ServiceConfig{
				Name: "web",
				Deploy: &types.DeployConfig{
					RestartPolicy: &types.RestartPolicy{Condition: "none"},
				},
			},
			expected: false,
		},
		{
			name: "deploy restart policy condition unless-stopped",
			service: types.ServiceConfig{
				Name: "web",
				Deploy: &types.DeployConfig{
					RestartPolicy: &types.RestartPolicy{Condition: "unless-stopped"},
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			project := &types.Project{
				Services: types.Services{
					tt.service.Name: tt.service,
				},
			}
			assert.Equal(t, restartsOnExit0(project, tt.service.Name), tt.expected)
		})
	}
}
