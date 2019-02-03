// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtp

import "github.com/pkg/errors"

// Common error definitions.
var (
	ErrInvalidVersion      = errors.New("got invalid version")
	ErrInvalidLength       = errors.New("length value is invalid")
	ErrTooShortToDecode    = errors.New("too short to decode as GTP")
	ErrTooShortToSerialize = errors.New("too short to serialize")
)
