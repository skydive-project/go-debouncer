/*
 * Copyright (C) 2018 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package debouncer

import (
	"sync/atomic"
	"time"
)

const (
	reset int32 = iota
	fired
)

// Debouncer defines a debouncer object
type Debouncer struct {
	state    int32
	callback func()
	timer    *time.Ticker
	stop     chan bool
	interval time.Duration
}

func (d *Debouncer) debounce() {
	for {
		select {
		case <-d.timer.C:
			if atomic.CompareAndSwapInt32(&d.state, fired, reset) {
				d.callback()
			}
		case <-d.stop:
			d.timer.Stop()
			return
		}
	}
}

// Call calls the callback of the Debouncer object
func (d *Debouncer) Call() {
	atomic.StoreInt32(&d.state, fired)
}

// Stop stops the debouncer
func (d *Debouncer) Stop() {
	d.stop <- true
}

// Start the debouncer
func (d *Debouncer) Start() {
	d.timer = time.NewTicker(d.interval)
	go d.debounce()
}

// New returns a Debouncer calling the given callback after a
// certain amount of time.
func New(interval time.Duration, callback func()) *Debouncer {
	d := &Debouncer{
		state:    reset,
		callback: callback,
		stop:     make(chan bool),
		interval: interval,
	}

	return d
}
