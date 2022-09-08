// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package shared

import (
	"fmt"
	log "github.com/ChainSafe/log15"
)

func GetBytes(key interface{}) ([]byte, error) {
	buf, ok := key.([]byte)
	if !ok {
		log.Error("failed to get bytes from interface")
		return nil, fmt.Errorf("failed to get bytes from interface", "interface", key)
	}

	return buf, nil
}
