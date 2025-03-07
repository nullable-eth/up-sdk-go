// Copyright 2021 Upbound Inc
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

package controlplanes

import (
	"time"

	"github.com/google/uuid"
)

// Status is the status of a control plane on Upbound.
type Status string

// A control plane will always be in one of the following phases.
const (
	StatusProvisioning Status = "provisioning"
	StatusUpdating     Status = "updating"
	StatusReady        Status = "ready"
	StatusDeleting     Status = "deleting"
)

// ControlPlane describes a control plane.
type ControlPlane struct {
	ID          uuid.UUID  `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatorID   uint       `json:"creatorId,omitempty"`
	Reserved    bool       `json:"reserved"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	ExpiresAt   time.Time  `json:"expiresAt"`
}

// PermissionGroup describes control plane permissions for the authenticated
// user.
type PermissionGroup string

const (
	// PermissionMember has the ability to read the basic environment of the
	// team.
	PermissionMember PermissionGroup = "member"
	// PermissionOwner has the ability to modify any object in a linked control
	// plane, including deleting the control plane.
	PermissionOwner PermissionGroup = "owner"
	// PermissionNone has no permissions on the control plane.
	PermissionNone PermissionGroup = "none"
)

// ControlPlaneResponse is the HTTP body returned by the Upbound API when
// fetching control planes.
type ControlPlaneResponse struct {
	ControlPlane ControlPlane    `json:"controlPlane"`
	Status       Status          `json:"controlPlanestatus,omitempty"`
	Permission   PermissionGroup `json:"controlPlanePermission,omitempty"`
}

// ControlPlaneListResponse is the HTTP body returned when listing control
// planes.
type ControlPlaneListResponse struct {
	ControlPlanes []ControlPlaneResponse `json:"controlPlanes"`
	Size          int                    `json:"size"`
	Page          int                    `json:"page"`
	Count         int                    `json:"count"`
}

// ControlPlaneCreateParameters are the parameters for creating a control plane.
type ControlPlaneCreateParameters struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
