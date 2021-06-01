/*
 * NCBI Datasets API
 *
 * NCBI service to query and download biological sequence data across all domains of life from NCBI databases.
 *
 * API version: v1alpha
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datasets
// GeneDatasetRequestSort struct for GeneDatasetRequestSort
type GeneDatasetRequestSort struct {
	Direction V1alpha1SortDirection `json:"direction,omitempty"`
	Field GeneDatasetRequestSortField `json:"field,omitempty"`
}
