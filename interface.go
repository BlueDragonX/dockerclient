package dockerclient

import (
	"io"
)

type Callback func(*Event, ...interface{})

type Client interface {
	Info() (*Info, error)
	ListContainers(all bool) ([]Container, error)
	InspectContainer(id string) (*ContainerInfo, error)
	CreateContainer(config *ContainerConfig, name string) (string, error)
	ContainerLogs(id string, stdout bool, stderr bool) (io.ReadCloser, error)
	StartContainer(id string, config *HostConfig) error
	StopContainer(id string, timeout int) error
	RestartContainer(id string, timeout int) error
	KillContainer(id string) error
	StartMonitorEvents(cb Callback, args ...interface{})
	StopAllMonitorEvents()
	Version() (*Version, error)
	PullImage(name, tag string) error
	RemoveContainer(id string, force bool) error
	ListImages() ([]*Image, error)
	RemoveImage(name string) error
}
