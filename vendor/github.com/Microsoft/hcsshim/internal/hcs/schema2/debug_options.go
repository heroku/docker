/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type DebugOptions struct {
	// BugcheckSavedStateFileName is the path for the file in which the guest VM state will be saved when
	// the guest crashes.
	BugcheckSavedStateFileName string `json:"BugcheckSavedStateFileName,omitempty"`
	// BugcheckNoCrashdumpSavedStateFileName is the path of the file in which the guest VM state will be
	// saved when the guest crashes but the guest isn't able to generate the crash dump. This usually
	// happens in early boot failures.
	BugcheckNoCrashdumpSavedStateFileName string `json:"BugcheckNoCrashdumpSavedStateFileName,omitempty"`
	TripleFaultSavedStateFileName         string `json:"TripleFaultSavedStateFileName,omitempty"`
	FirmwareDumpFileName                  string `json:"FirmwareDumpFileName,omitempty"`
}