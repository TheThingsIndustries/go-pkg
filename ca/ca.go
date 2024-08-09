// Copyright © 2024 The Things Industries B.V.

// Package ca provides The Things Industries Certificate Authorities (CAs).
package ca

import _ "embed"

var (
	// RootCA is PEM encoded The Things Industries Root CA.
	//
	//go:embed tti-root.pem
	RootCA []byte

	// TestCA is the PEM encoded The Things Industries Test CA.
	//
	//go:embed tti-test.pem
	TestCA []byte

	// GatewaysCAs is the PEM encoded intermediate CAs for signing gateway identities.
	//
	//go:embed gateways.pem
	GatewaysCAs []byte
)
