// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufferedwriter

import (
	"bufio"
	"io"
)

const (
	defaultBufferSize = 4 << 10 // 4096 KB
)

// BufferedWriter is a writter with buffer, it supports io.Writer, io.WriterAt and io.WriterSeeker.
type BufferedWriter interface {
	io.Writer
	// Flush writes any buffered data to the underlying io.Writer.
	Flush() error
}

// New is shorthand for NewSize(w, 4096), a 4KB Buffer.
func New(w io.Writer) BufferedWriter {
	return NewSize(w, defaultBufferSize)
}

// NewSize creates a buffered writer with a specified size while taking into account the underlying capabilities of the writer w,
// which might implement either [io.WriterAt] or [io.WriteSeeker]. This allows the buffered writer to maintain the ability to write at a specific byte position.
//
// Use-case scenario:
//   - An *os.File may be passed to a function receiving an [io.Writer], while that function may assert the original ability of the value to write at a specific byte
//     position if possible to enable a faster processing path. Nonetheless, directly working with *os.File for frequent writing small byte segments can affect
//     performance due to numerous syscalls. To alleviate this issue, incorporating a buffered writer for the *os.File becomes essential.
//     Unlike bufio.Writer that encapsulates everything as a buffered [io.Writer], this approach preserves the inherent capabilities of [io.WriterAt] and [io.WriteSeeker].
//
// Just like any other buffered writer, the Flush() method should be called after the process is completed to write the unwritten buffered data.
func NewSize(w io.Writer, size int) BufferedWriter {
	if size < defaultBufferSize {
		size = defaultBufferSize
	}

	if wra, ok := w.(io.WriterAt); ok {
		return &WriterAt{
			bw:  bufio.NewWriterSize(w, size),
			wra: wra,
		}
	}

	if wrs, ok := w.(io.WriteSeeker); ok {
		return &WriteSeeker{
			bw:  bufio.NewWriterSize(w, size),
			wrs: wrs,
		}
	}

	return bufio.NewWriterSize(w, size)
}

type WriterAt struct {
	bw  *bufio.Writer
	wra io.WriterAt
}

func (w *WriterAt) Write(p []byte) (n int, err error) {
	return w.bw.Write(p)
}

func (w *WriterAt) WriteAt(p []byte, offset int64) (n int, err error) {
	if err = w.bw.Flush(); err != nil { // flush any buffered data from buffered writer in case we are writing on unflushed data.
		return
	}
	w.bw.Reset(w.wra.(io.Writer)) // reset buffer

	return w.wra.WriteAt(p, offset)
}

func (w *WriterAt) Flush() error { return w.bw.Flush() }

type WriteSeeker struct {
	bw  *bufio.Writer
	wrs io.WriteSeeker
}

func (w *WriteSeeker) Write(p []byte) (n int, err error) {
	return w.bw.Write(p)
}

func (w *WriteSeeker) Seek(offset int64, whence int) (n int64, err error) {
	if err = w.bw.Flush(); err != nil { // flush any buffered data from buffered writer in case we are seeking on unflushed data.
		return
	}
	w.bw.Reset(w.wrs) // reset buffer

	n, err = w.wrs.Seek(offset, whence)
	if err != nil {
		return
	}
	return
}

func (w *WriteSeeker) Flush() error { return w.bw.Flush() }
