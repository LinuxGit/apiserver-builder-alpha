/*
Copyright 2017 The Kubernetes Authors.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Generating code from student_types.go file will generate storage and status REST endpoints for
// Student.

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +k8s:openapi-gen=true
// +resource:path=students,rest=StudentREST
// +subresource:request=StudentComputer,path=computer,kind=StudentComputer,rest=StudentComputerREST
// +subresource:request=StudentBook,path=book,kind=StudentBook,rest=StudentBookREST
// +subresource:request=StudentPencil,path=pencil,kind=StudentPencil,rest=StudentPencilREST
type Student struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StudentSpec   `json:"spec,omitempty"`
	Status StudentStatus `json:"status,omitempty"`
}

// StudentSpec defines the desired state of Student
type StudentSpec struct {
	ID int `json:"id,omitempty"`
}

// StudentStatus defines the observed state of Student
type StudentStatus struct {
	// GPA is the GPA of the student.
	GPA float64 `json:"GPA,omitempty"`
}
