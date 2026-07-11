// Package engine contains the V1 synchronous calculation stages of the
// rendering pipeline. External data access is supplied through explicit
// synchronous function ports rather than concrete repository dependencies.
//
// Functions in this package do not start goroutines, own worker pools, enforce
// concurrency limits, or interpret context cancellation. Callers in the parent
// usecase package own I/O, cancellation, stage ordering, and any future
// parallel scheduling. Independent documents may be processed concurrently;
// calculations within one document retain their defined order where later
// results depend on earlier geometry or routing decisions.
package engine
