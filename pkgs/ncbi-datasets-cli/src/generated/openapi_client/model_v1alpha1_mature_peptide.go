/*
 * NCBI Datasets API
 *
 * NCBI service to query and download biological sequence data across all domains of life from NCBI databases.
 *
 * API version: v1alpha
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datasets
// V1alpha1MaturePeptide struct for V1alpha1MaturePeptide
type V1alpha1MaturePeptide struct {
	AccessionVersion string `json:"accession_version,omitempty"`
	Length int64 `json:"length,omitempty"`
	Name string `json:"name,omitempty"`
}
