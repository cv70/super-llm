package gio

import (
	"io"

	"k8s.io/klog/v2"
)

func ReadAndClose(rc io.ReadCloser) ([]byte, error) {
	defer func() {
		err := rc.Close()
		if err != nil {
			klog.Error(err)
		}
	}()
	return io.ReadAll(rc)
}
