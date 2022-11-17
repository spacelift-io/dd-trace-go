// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016 Datadog, Inc.

//go:build !go1.22

// TODO(knusbaum): This file should be deleted once go1.21 falls out of support
package tracer

import (
	"math"
	"math/rand/v2"
)

func randUint64() uint64 {
	return rand.Uint64()
}

func generateSpanID(startTime int64) uint64 {
	return rand.Uint64() & math.MaxInt64
}

// generateSpanID returns a random uint64 that has been XORd with the startTime.
// This is done to get around the 32-bit random seed limitation that may create collisions if there is a large number
// of go services all generating spans.
func generateSpanID(startTime int64) uint64 {
	return random.Uint64() ^ uint64(startTime)
}

func randUint64() uint64 {
	return random.Uint64()
}
