/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package tx

// Header is contain header data
type Header struct {
	ID          int
	Time        int64
	EcosystemID int64
	KeyID       int64
	NetworkID   int64
	PublicKey   []byte
	//
	//Add sub node processing
	PrivateFor []string
