//
// token.go
//
// Copyright (c) 2019 Markku Rossi
//
// All rights reserved.
//

package auth

import (
	"github.com/markkurossi/go-libs/tlv"
)

const (
	TOKEN_VALUES tlv.Type = iota
	TOKEN_SIGNATURE
	T_TENANT_ID
	T_CLIENT_ID
	T_CREATED
	T_SCOPE
)

const (
	SCOPE_ADMIN tlv.Type = iota
)

const (
	KEY_CLIENT_ID_SECRET    = "client-id-secret"
	KEY_TOKEN_SIGNATURE_KEY = "token-signature-key"
)

const (
	ASSET_AUTH_PUBKEY = "auth-pubkey"
)
