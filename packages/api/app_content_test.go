/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppContent(t *testing.T) {
	assert.NoError(t, keyLogin(1))
	if err != nil {
		t.Error(err)
		return
	}

	if len(ret.Blocks) == 0 {
		t.Error("incorrect blocks count")
	}

	if len(ret.Contracts) == 0 {
		t.Error("incorrect contracts count")
	}

	if len(ret.Pages) == 0 {
		t.Error("incorrent pages count")
	}
}
