// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64
// +build 386 amd64

package network_listen

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type listenEventData struct{ Args [3]uint64 }

// loadListen returns the embedded CollectionSpec for listen.
func loadListen() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ListenBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load listen: %w", err)
	}

	return spec, err
}

// loadListenObjects loads listen and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*listenObjects
//	*listenPrograms
//	*listenMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadListenObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadListen()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// listenSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type listenSpecs struct {
	listenProgramSpecs
	listenMapSpecs
}

// listenSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type listenProgramSpecs struct {
	KprobeListen *ebpf.ProgramSpec `ebpf:"kprobe_listen"`
}

// listenMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type listenMapSpecs struct {
	Event *ebpf.MapSpec `ebpf:"event"`
}

// listenObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadListenObjects or ebpf.CollectionSpec.LoadAndAssign.
type listenObjects struct {
	listenPrograms
	listenMaps
}

func (o *listenObjects) Close() error {
	return _ListenClose(
		&o.listenPrograms,
		&o.listenMaps,
	)
}

// listenMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadListenObjects or ebpf.CollectionSpec.LoadAndAssign.
type listenMaps struct {
	Event *ebpf.Map `ebpf:"event"`
}

func (m *listenMaps) Close() error {
	return _ListenClose(
		m.Event,
	)
}

// listenPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadListenObjects or ebpf.CollectionSpec.LoadAndAssign.
type listenPrograms struct {
	KprobeListen *ebpf.Program `ebpf:"kprobe_listen"`
}

func (p *listenPrograms) Close() error {
	return _ListenClose(
		p.KprobeListen,
	)
}

func _ListenClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed listen_bpfel_x86.o
var _ListenBytes []byte
