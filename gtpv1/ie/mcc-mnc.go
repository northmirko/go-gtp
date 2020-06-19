// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"

	"github.com/wmnsk/go-gtp/utils"
)

// MCC returns MCC value if type matches.
func (i *IE) MCC() (string, error) {
	switch i.Type {
	case RouteingAreaIdentity:
		if len(i.Payload) < 2 {
			return "", io.ErrUnexpectedEOF
		}
		return utils.SwappedBytesToStr(i.Payload[0:2], false), nil
	case UserLocationInformation:
		if len(i.Payload) < 3 {
			return "", io.ErrUnexpectedEOF
		}
		return utils.SwappedBytesToStr(i.Payload[1:3], false), nil
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}

// MustMCC returns MCC in string if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustMCC() string {
	v, _ := i.MCC()
	return v
}

// MNC returns MNC value if type matches.
func (i *IE) MNC() (string, error) {
	switch i.Type {
	case RouteingAreaIdentity:
		if len(i.Payload) < 2 {
			return "", io.ErrUnexpectedEOF
		}
		return utils.SwappedBytesToStr(i.Payload[1:2], true), nil
	case UserLocationInformation:
		if len(i.Payload) < 3 {
			return "", io.ErrUnexpectedEOF
		}
		return utils.SwappedBytesToStr(i.Payload[2:3], true), nil
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}

// MustMNC returns MNC in string if type matches.
// This should only be used if it is assured to have the value.
func (i *IE) MustMNC() string {
	v, _ := i.MNC()
	return v
}