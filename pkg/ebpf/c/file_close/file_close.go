// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 Authors of Tarian & the Organization created Tarian

package file_close

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/ringbuf"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -cflags $BPF_CFLAGS -target $CURR_ARCH -type event_data close close.bpf.c -- -I../../../../headers

// loads the ebpf specs like maps, programs
func getEbpfObject() (*closeObjects, error) {
	var bpfObj closeObjects
	err := loadCloseObjects(&bpfObj, nil)
	if err != nil {
		return nil, err
	}

	return &bpfObj, nil
}

type CloseEventData struct {
	closeEventData
}

type CloseDetector struct {
	ebpfLink      link.Link
	ringbufReader *ringbuf.Reader
}

// NewCloseDetector returns a new instance of CloseDetector
func NewCloseDetector() *CloseDetector {
	return &CloseDetector{}
}

// Start the close detector by attaching ebpf program to
// hook in kernel and opens the map to read the data.
// If it cannot be started an error is returned.
func (o *CloseDetector) Start() error {
	bpfObjs, err := getEbpfObject()
	if err != nil {
		return err
	}

	l, err := link.Kprobe("__x64_sys_close", bpfObjs.KprobeClose, nil)
	if err != nil {
		return err
	}

	o.ebpfLink = l

	rd, err := ringbuf.NewReader(bpfObjs.Event)
	if err != nil {
		return err
	}

	o.ringbufReader = rd
	return nil
}

// closes the EBPF objects
func (o *CloseDetector) Close() error {
	err := o.ebpfLink.Close()
	if err != nil {
		return err
	}

	return o.ringbufReader.Close()
}

// reads the next event from the ringbuffer
func (o *CloseDetector) Read() (CloseEventData, error) {
	var event CloseEventData
	// reads the data from ringbuffer
	record, err := o.ringbufReader.Read()
	if err != nil {
		if errors.Is(err, ringbuf.ErrClosed) {
			return event, err
		}

		return event, err
	}

	// read the raw sample from the record.RawSample
	if err := binary.Read(bytes.NewBuffer(record.RawSample), binary.LittleEndian, &event); err != nil {
		return event, err
	}

	return event, nil
}

// reads data from a ring buffer and returns it as an interface
func (o *CloseDetector) ReadAsInterface() (any, error) {
	return o.Read()
}
