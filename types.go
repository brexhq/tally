// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tally

import (
	"time"

	ubertally "github.com/uber-go/tally"
)

// Scope aliases to uber's Scope
type Scope = ubertally.Scope

// Counter aliases to uber's Counter
type Counter = ubertally.Counter

// Gauge aliases to uber's Gauge
type Gauge = ubertally.Gauge

// Timer aliases to uber's Timer
type Timer = ubertally.Timer

// Histogram aliases to uber's Histogram
type Histogram = ubertally.Histogram

// Stopwatch aliases to uber's Stopwatch
type Stopwatch = ubertally.Stopwatch

// NewStopwatch aliases to uber's NewStopwatch
func NewStopwatch(start time.Time, r StopwatchRecorder) Stopwatch {
	return ubertally.NewStopwatch(start, r)
}

// StopwatchRecorder aliases to uber's StopwatchRecorder
type StopwatchRecorder = ubertally.StopwatchRecorder

// Buckets aliases to uber's Buckets
type Buckets = ubertally.Buckets

// BucketPair aliases to uber's BucketPair
type BucketPair = ubertally.BucketPair

// Capabilities aliases to uber's Capabilities
type Capabilities = ubertally.Capabilities
