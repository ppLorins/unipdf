//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package license helps manage commercial licenses and check if they are valid for the version of unipdf used.
package license ;import _ad "github.com/unidoc/unipdf/v3/internal/license";

// LicenseKey represents a loaded license key.
type LicenseKey =_ad .LicenseKey ;

// SetLicenseKey sets and validates the license key.
func SetLicenseKey (content string ,customerName string )error {
	return nil
	//return _ad .SetLicenseKey (content ,customerName );
}

const (LicenseTierUnlicensed =_ad .LicenseTierUnlicensed ;LicenseTierCommunity =_ad .LicenseTierCommunity ;LicenseTierIndividual =_ad .LicenseTierIndividual ;LicenseTierBusiness =_ad .LicenseTierBusiness ;);

// GetLicenseKey returns the currently loaded license key.
func GetLicenseKey ()*LicenseKey {return _ad .GetLicenseKey ()};

// GetMeteredState checks the currently used metered document usage status,
// documents used and credits available.
func GetMeteredState ()(_ad .MeteredStatus ,error ){return _ad .GetMeteredState ()};

// MakeUnlicensedKey returns a default key.
func MakeUnlicensedKey ()*LicenseKey {return _ad .MakeUnlicensedKey ()};

// SetMeteredKey sets the metered API key required for SaaS operation.
// Document usage is reported periodically for the product to function correctly.
func SetMeteredKey (apiKey string )error {return _ad .SetMeteredKey (apiKey )};