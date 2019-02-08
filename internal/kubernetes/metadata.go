package kubernetes

// Metadata is a simple tagging system for resources
type Metadata struct {
	Name      string
	Namespace string
	Labels    map[string]string
}

// NewMetadata returns a default empty metadata
func NewMetadata() Metadata {
	return Metadata{Labels: make(map[string]string)}
}
