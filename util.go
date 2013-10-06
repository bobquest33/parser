package main

import (
	"archive/zip"
)

func findZipFile(r *zip.ReadCloser, path string) *zip.File {
	for _, f := range r.File {
		if f.Name == path {
			return f
		}
	}
	return nil
}
