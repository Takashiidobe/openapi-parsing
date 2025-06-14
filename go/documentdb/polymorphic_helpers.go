// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.10.7, generator: @autorest/go@4.0.0-preview.72)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package documentdb

import "encoding/json"

func unmarshalBackupPolicyClassification(rawMsg json.RawMessage) (BackupPolicyClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BackupPolicyClassification
	switch m["type"] {
	case string(BackupPolicyTypeContinuous):
		b = &ContinuousModeBackupPolicy{}
	case string(BackupPolicyTypePeriodic):
		b = &PeriodicModeBackupPolicy{}
	default:
		b = &BackupPolicy{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

