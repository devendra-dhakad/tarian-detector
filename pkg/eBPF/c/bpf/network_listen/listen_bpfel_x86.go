// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package network_listen

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type listenEventData struct {
	EventContext struct {
		Ts        uint64
		StartTime uint64
		Pid       uint32
		Tgid      uint32
		Ppid      uint32
		Glpid     uint32
		Uid       uint32
		Gid       uint32
		Comm      [16]uint8
		Cwd       [32]uint8
		CgroupId  uint64
		NodeInfo  struct {
			Sysname    [65]uint8
			Nodename   [65]uint8
			Release    [65]uint8
			Version    [65]uint8
			Machine    [65]uint8
			Domainname [65]uint8
		}
		MountInfo struct {
			MountId      int32
			MountNsId    uint32
			MountDevname [256]uint8
		}
	}
	_       [2]byte
	Id      int32
	Fd      int32
	Backlog int32
	Ret     int32
}

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
	KprobeListenEntry   *ebpf.ProgramSpec `ebpf:"kprobe_listen_entry"`
	KretprobeListenExit *ebpf.ProgramSpec `ebpf:"kretprobe_listen_exit"`
}

// listenMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type listenMapSpecs struct {
	ListenEventMap *ebpf.MapSpec `ebpf:"listen_event_map"`
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
	ListenEventMap *ebpf.Map `ebpf:"listen_event_map"`
}

func (m *listenMaps) Close() error {
	return _ListenClose(
		m.ListenEventMap,
	)
}

// listenPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadListenObjects or ebpf.CollectionSpec.LoadAndAssign.
type listenPrograms struct {
	KprobeListenEntry   *ebpf.Program `ebpf:"kprobe_listen_entry"`
	KretprobeListenExit *ebpf.Program `ebpf:"kretprobe_listen_exit"`
}

func (p *listenPrograms) Close() error {
	return _ListenClose(
		p.KprobeListenEntry,
		p.KretprobeListenExit,
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
