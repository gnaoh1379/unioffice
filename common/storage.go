// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package common

import "github.com/gnaoh1379/unioffice/common/tempstorage/diskstore"

func init() {
	diskstore.SetAsStorage() // set disk storage by default
}
