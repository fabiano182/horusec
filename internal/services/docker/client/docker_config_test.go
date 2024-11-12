// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package client

import (
	"context"
	"io"
	"os"
	"strings"
	"testing"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec/internal/utils/testutil"
)

func TestNewDockerAPI(t *testing.T) {
	t.Run("Should not panic when success connect to docker", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewDockerClient()
		})
	})

	t.Run("Should use docker host from environment variable", func(t *testing.T) {
		dockerHost := "http://localhost:1234"
		os.Setenv("DOCKER_HOST", dockerHost)

		c := NewDockerClient()

		assert.Equal(t, dockerHost, c.DaemonHost())
	})
}

func TestMock(t *testing.T) {
	t.Run("Should return expected data to ContainerCreate", func(t *testing.T) {
		m := testutil.NewDockerClientMock()
		m.On("ContainerCreate").Return(container.CreateResponse{}, nil)
		_, err := m.ContainerCreate(context.TODO(), nil, nil, nil, nil, "")
		assert.NoError(t, err)
	})
	t.Run("Should return expected data to ContainerStart", func(t *testing.T) {
		m := testutil.NewDockerClientMock()
		m.On("ContainerStart").Return(nil)
		err := m.ContainerStart(context.TODO(), "", container.StartOptions{})
		assert.NoError(t, err)
	})
	t.Run("Should return expected data to ContainerWait", func(t *testing.T) {
		m := testutil.NewDockerClientMock()
		m.On("ContainerWait").Return(container.WaitResponse{}, nil)
		_, err := m.ContainerWait(context.TODO(), "", "")
		assert.NoError(t, <-err)
	})
	t.Run("Should return expected data to ContainerLogs", func(t *testing.T) {
		m := testutil.NewDockerClientMock()

		m.On("ContainerLogs").Return(io.NopCloser(strings.NewReader("some text")), nil)
		_, err := m.ContainerLogs(context.TODO(), "", container.LogsOptions{})
		assert.NoError(t, err)
	})
	t.Run("Should return expected data to ContainerRemove", func(t *testing.T) {
		m := testutil.NewDockerClientMock()
		m.On("ContainerRemove").Return(nil)
		err := m.ContainerRemove(context.TODO(), "", container.RemoveOptions{})
		assert.NoError(t, err)
	})
	t.Run("Should return expected data to ImageList", func(t *testing.T) {
		m := testutil.NewDockerClientMock()
		m.On("ImageList").Return([]image.Summary{}, nil)
		_, err := m.ImageList(context.TODO(), image.ListOptions{})
		assert.NoError(t, err)
	})
	t.Run("Should return expected data to ImagePull", func(t *testing.T) {
		m := testutil.NewDockerClientMock()
		m.On("ImagePull").Return(io.NopCloser(strings.NewReader("some text")), nil)
		_, err := m.ImagePull(context.TODO(), "", image.PullOptions{})
		assert.NoError(t, err)
	})
	t.Run("Should return expected data to Ping", func(t *testing.T) {
		m := testutil.NewDockerClientMock()
		m.On("Ping").Return(dockerTypes.Ping{}, nil)
		_, err := m.Ping(context.TODO())
		assert.NoError(t, err)
	})
	t.Run("Should return expected data to Ping", func(t *testing.T) {
		m := testutil.NewDockerClientMock()
		m.On("ContainerList").Return([]dockerTypes.Container{}, nil)
		_, err := m.ContainerList(context.TODO(), container.ListOptions{})
		assert.NoError(t, err)
	})
}
