// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package network_bind

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bindEventData struct {
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
	_        [2]byte
	Id       int32
	Fd       int32
	Addrlen  int32
	SaFamily uint16
	Port     uint16
	V4Addr   struct{ S_addr uint32 }
	V6Addr   struct{ S6Addr [16]uint8 }
	UnixAddr struct{ Path [108]int8 }
	Padding  uint32
	Ret      int32
}

// loadBind returns the embedded CollectionSpec for bind.
func loadBind() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BindBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bind: %w", err)
	}

	return spec, err
}

// loadBindObjects loads bind and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bindObjects
//	*bindPrograms
//	*bindMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBindObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBind()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bindSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bindSpecs struct {
	bindProgramSpecs
	bindMapSpecs
}

// bindSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bindProgramSpecs struct {
	KprobeBindEntry   *ebpf.ProgramSpec `ebpf:"kprobe_bind_entry"`
	KretprobeBindExit *ebpf.ProgramSpec `ebpf:"kretprobe_bind_exit"`
}

// bindMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bindMapSpecs struct {
	Event *ebpf.MapSpec `ebpf:"event"`
}

// bindObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBindObjects or ebpf.CollectionSpec.LoadAndAssign.
type bindObjects struct {
	bindPrograms
	bindMaps
}

func (o *bindObjects) Close() error {
	return _BindClose(
		&o.bindPrograms,
		&o.bindMaps,
	)
}

// bindMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBindObjects or ebpf.CollectionSpec.LoadAndAssign.
type bindMaps struct {
	Event *ebpf.Map `ebpf:"event"`
}

func (m *bindMaps) Close() error {
	return _BindClose(
		m.Event,
	)
}

// bindPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBindObjects or ebpf.CollectionSpec.LoadAndAssign.
type bindPrograms struct {
	KprobeBindEntry   *ebpf.Program `ebpf:"kprobe_bind_entry"`
	KretprobeBindExit *ebpf.Program `ebpf:"kretprobe_bind_exit"`
}

func (p *bindPrograms) Close() error {
	return _BindClose(
		p.KprobeBindEntry,
		p.KretprobeBindExit,
	)
}

func _BindClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bind_bpfel_x86.o
var _BindBytes []byte
