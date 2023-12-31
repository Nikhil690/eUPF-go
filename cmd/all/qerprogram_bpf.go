// Code generated by bpf2go; DO NOT EDIT.

package ebpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadQerProgram returns the embedded CollectionSpec for QerProgram.
func LoadQerProgram() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_QerProgramBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load QerProgram: %w", err)
	}

	return spec, err
}

// LoadQerProgramObjects loads QerProgram and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*QerProgramObjects
//	*QerProgramPrograms
//	*QerProgramMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadQerProgramObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadQerProgram()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// QerProgramSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type QerProgramSpecs struct {
	QerProgramProgramSpecs
	QerProgramMapSpecs
}

// QerProgramSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type QerProgramProgramSpecs struct {
	UpfQerProgramFunc *ebpf.ProgramSpec `ebpf:"upf_qer_program_func"`
}

// QerProgramMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type QerProgramMapSpecs struct {
	UpfPipeline *ebpf.MapSpec `ebpf:"upf_pipeline"`
}

// QerProgramObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadQerProgramObjects or ebpf.CollectionSpec.LoadAndAssign.
type QerProgramObjects struct {
	QerProgramPrograms
	QerProgramMaps
}

func (o *QerProgramObjects) Close() error {
	return _QerProgramClose(
		&o.QerProgramPrograms,
		&o.QerProgramMaps,
	)
}

// QerProgramMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadQerProgramObjects or ebpf.CollectionSpec.LoadAndAssign.
type QerProgramMaps struct {
	UpfPipeline *ebpf.Map `ebpf:"upf_pipeline"`
}

func (m *QerProgramMaps) Close() error {
	return _QerProgramClose(
		m.UpfPipeline,
	)
}

// QerProgramPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadQerProgramObjects or ebpf.CollectionSpec.LoadAndAssign.
type QerProgramPrograms struct {
	UpfQerProgramFunc *ebpf.Program `ebpf:"upf_qer_program_func"`
}

func (p *QerProgramPrograms) Close() error {
	return _QerProgramClose(
		p.UpfQerProgramFunc,
	)
}

func _QerProgramClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed qerprogram_bpf.o
var _QerProgramBytes []byte
