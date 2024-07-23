// Copyright © 2024 The Things Industries B.V.

// Package ca provides The Things Industries Root CA.
package ca

import _ "embed"

var (
	// RootCA is PEM encoded The Things Industries Root CA.
	//
	// go:embed tti.pem
	RootCA []byte

	// GatewaysCAs is the PEM encoded gateway identity signer intermediate CAs.
	//
	//go:embed gateways.pem
	GatewaysCAs []byte
)
