/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// JobTemplateSpec describes the data a Job should have when created from a template
type V2alpha1JobTemplateSpec struct {

	// Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	Spec *V1JobSpec `json:"spec,omitempty"`
}
