/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package api

import (
	"net/http"
)

// nodeContract is used when calling a cron contract in OBS mode
func nodeContractHandler(w http.ResponseWriter, r *http.Request) {
