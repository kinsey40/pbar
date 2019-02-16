/*
 * The MIT License
 *
 * Copyright 2018 kinsey40.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * File:   clock.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 *
 */

package clock

import (
	"fmt"
	"time"
)

type Clock interface {
	Now() time.Time
	Subtract(time.Time) time.Duration
	SetStart(time.Time)
	Start() time.Time
	Seconds(time.Duration) float64
	Remaining(float64) time.Duration
	Format(time.Duration) string
}

// clock implements a real-time clock by simply wrapping the time package functions.
type clock struct {
	StartTime time.Time
}

// New returns an instance of a real-time clock.
func NewClock() Clock {
	c := new(clock)

	return c
}

func (c *clock) Now() time.Time {
	return time.Now()
}

func (c *clock) Subtract(now time.Time) time.Duration {
	return now.Sub(c.StartTime)
}

func (c *clock) SetStart(t time.Time) {
	c.StartTime = t
}

func (c *clock) Start() time.Time {
	return c.StartTime
}

func (c *clock) Seconds(d time.Duration) float64 {
	return d.Seconds()
}

func (c *clock) Remaining(fraction float64) time.Duration {
	return time.Duration(fraction) * time.Second
}

func (c *clock) Format(d time.Duration) string {
	secs := (d % time.Minute) / time.Second
	mins := (d % time.Hour) / time.Minute
	hours := d / time.Hour

	if hours == 0 {
		return fmt.Sprintf("%02dm:%02ds", mins, secs)
	}

	return fmt.Sprintf("%02dh:%02dm:%02ds", hours, mins, secs)
}
