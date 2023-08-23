package testhelpers

import "io"

// GcpVM represents a GCP VM.
type GcpVM interface {
	GcpResource
	RunCommandViaSsh(command string) error
	StartVm() error
	StopVm() error
}

// GcpStorageBucket represents a GCP storage bucket.
type GcpStorageBucket interface {
	GcpResource
	DownloadData(id string) (io.ByteReader, error)
	UploadData(id string, writer io.ByteWriter) error
}
