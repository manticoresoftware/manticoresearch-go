package openapi

import (
	"encoding/json"
	"fmt"
)

// splitDocumentIDJSON parses a raw JSON value into numeric id and/or uuid string.
func splitDocumentIDJSON(raw json.RawMessage) (id *uint64, uuid *string, err error) {
	if len(raw) == 0 || string(raw) == "null" {
		return nil, nil, nil
	}
	var asString string
	if err := json.Unmarshal(raw, &asString); err == nil {
		return nil, &asString, nil
	}
	var asUint uint64
	if err := json.Unmarshal(raw, &asUint); err == nil {
		return &asUint, nil, nil
	}
	return nil, nil, fmt.Errorf("document id must be uint64 or string, got %s", string(raw))
}

func wireDocumentID(id *uint64, uuid *string) interface{} {
	if uuid != nil {
		return *uuid
	}
	if id != nil {
		return *id
	}
	return nil
}
