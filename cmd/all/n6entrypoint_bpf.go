// Code generated by bpf2go; DO NOT EDIT.

package ebpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadN6Entrypoint returns the embedded CollectionSpec for N6Entrypoint.
func LoadN6Entrypoint() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_N6EntrypointBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load N6Entrypoint: %w", err)
	}

	return spec, err
}

// LoadN6EntrypointObjects loads N6Entrypoint and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*N6EntrypointObjects
//	*N6EntrypointPrograms
//	*N6EntrypointMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadN6EntrypointObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadN6Entrypoint()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// N6EntrypointSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type N6EntrypointSpecs struct {
	N6EntrypointProgramSpecs
	N6EntrypointMapSpecs
}

// N6EntrypointSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type N6EntrypointProgramSpecs struct {
	UpfN6EntrypointFunc *ebpf.ProgramSpec `ebpf:"upf_n6_entrypoint_func"`
}

// N6EntrypointMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type N6EntrypointMapSpecs struct {
	UpfPipeline *ebpf.MapSpec `ebpf:"upf_pipeline"`
}

// N6EntrypointObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadN6EntrypointObjects or ebpf.CollectionSpec.LoadAndAssign.
type N6EntrypointObjects struct {
	N6EntrypointPrograms
	N6EntrypointMaps
}

func (o *N6EntrypointObjects) Close() error {
	return _N6EntrypointClose(
		&o.N6EntrypointPrograms,
		&o.N6EntrypointMaps,
	)
}

// N6EntrypointMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadN6EntrypointObjects or ebpf.CollectionSpec.LoadAndAssign.
type N6EntrypointMaps struct {
	UpfPipeline *ebpf.Map `ebpf:"upf_pipeline"`
}

func (m *N6EntrypointMaps) Close() error {
	return _N6EntrypointClose(
		m.UpfPipeline,
	)
}

// N6EntrypointPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadN6EntrypointObjects or ebpf.CollectionSpec.LoadAndAssign.
type N6EntrypointPrograms struct {
	UpfN6EntrypointFunc *ebpf.Program `ebpf:"upf_n6_entrypoint_func"`
}

func (p *N6EntrypointPrograms) Close() error {
	return _N6EntrypointClose(
		p.UpfN6EntrypointFunc,
	)
}

func _N6EntrypointClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed n6entrypoint_bpf.o
var _N6EntrypointBytes []byte