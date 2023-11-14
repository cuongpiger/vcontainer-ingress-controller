package controller

const (
	// High enough Queries per second (QPS) to fit all expected use cases. QPS=0 is not set here, because client code is overriding it
	defaultQPS = 1e6

	// High enough Burst to fit all expected use cases. Burst=0 is not set here, because client code is overriding it
	defaultBurst = 1e6
)
