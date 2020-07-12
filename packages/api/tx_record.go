/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package api

import (
	"net/http"
	"strings"

	"github.com/IBAX-io/go-ibax/packages/model"

	"github.com/gorilla/mux"
)

func getTxRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hashes := params["hashes"]

	var (
		hashList   []string
		resultList []interface{}
	)
	if len(hashes) > 0 {
		hashList = strings.Split(hashes, ",")
	}
