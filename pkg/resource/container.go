//
// Copyright (c) 2018, Cisco Systems
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice, this
//   list of conditions and the following disclaimer in the documentation and/or
//   other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERIVCES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
// ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package resource

import (
	"github.com/cisco/arc/pkg/config"
	"github.com/cisco/arc/pkg/route"
)

// StaticContainer provides the interface to the static portion of a
// container. This information is provided via config file and is implemented
// by config.Container.
type StaticContainer interface {
	config.Printer

	// Name of the container.
	Name() string
}

// DynamicContainer provides access to the dynamic portion of the container.
type DynamicContainer interface {
	Loader
	Creator
	Destroyer
	Provisioner
	Auditor
	Informer
}

// ProviderContainer provides a resource interface for the provider supplied container.
type ProviderContainer interface {
	DynamicContainer
}

// Container provides the resource interface used for the container
// object implemented in the arc package.
type Container interface {
	route.Router
	StaticContainer
	DynamicContainer
	Helper

	// ContainerService provides access to container's parent.
	ContainerService() ContainerService

	// ProviderContainer allows access to the provider's container object.
	ProviderContainer() ProviderContainer
}
