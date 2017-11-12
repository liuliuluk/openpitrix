// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package runtime

var pkgRuntimeList []RuntimeInterface

type RuntimeInterface interface {
	Name() string

	Run(app string, args ...string) error
}

func RegisterRuntime(runtime RuntimeInterface) {
	pkgRuntimeList = append(pkgRuntimeList, runtime)
}

func getRuntime(name string) RuntimeInterface {
	for _, rt := range pkgRuntimeList {
		if rt.Name() == name {
			return rt
		}
	}
	return nil
}
