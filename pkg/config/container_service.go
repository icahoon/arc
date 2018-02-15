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
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
// ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package config

import "github.com/cisco/arc/pkg/msg"

// The configuration of the container_service object.
type ContainerService struct {
	Provider     *Provider              `json:"provider"`
	Name_        string                 `json:"name"`
	Repositories []*ContainerRepository `json:"repositories"`
	Containers   []*Container           `json:containers"`
}

// Name of the container service.
func (c *ContainerService) Name() string {
	return c.Name_
}

// Print provides a user friendly way to view the configuration of the container service.
func (cs *ContainerService) Print() {
	msg.Info("Container Service Config")
	msg.Detail("%-20s\t%s", "name", cs.Name())
	msg.IndentInc()
	if cs.Provider != nil {
		cs.Provider.Print()
	}
	for _, r := range cs.Repositories {
		r.Print()
	}
	msg.Info("Containers")
	for _, c := range cs.Containers {
		c.Print()
	}
	msg.IndentDec()
}
