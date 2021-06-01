// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: ncbi/datasets/v1alpha1/reports/prokaryote_gene_location.proto

package reports

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProkaryoteGeneLocation_Completeness int32

const (
	ProkaryoteGeneLocation_complete ProkaryoteGeneLocation_Completeness = 0
	ProkaryoteGeneLocation_partial  ProkaryoteGeneLocation_Completeness = 1
)

// Enum value maps for ProkaryoteGeneLocation_Completeness.
var (
	ProkaryoteGeneLocation_Completeness_name = map[int32]string{
		0: "complete",
		1: "partial",
	}
	ProkaryoteGeneLocation_Completeness_value = map[string]int32{
		"complete": 0,
		"partial":  1,
	}
)

func (x ProkaryoteGeneLocation_Completeness) Enum() *ProkaryoteGeneLocation_Completeness {
	p := new(ProkaryoteGeneLocation_Completeness)
	*p = x
	return p
}

func (x ProkaryoteGeneLocation_Completeness) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProkaryoteGeneLocation_Completeness) Descriptor() protoreflect.EnumDescriptor {
	return file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_enumTypes[0].Descriptor()
}

func (ProkaryoteGeneLocation_Completeness) Type() protoreflect.EnumType {
	return &file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_enumTypes[0]
}

func (x ProkaryoteGeneLocation_Completeness) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProkaryoteGeneLocation_Completeness.Descriptor instead.
func (ProkaryoteGeneLocation_Completeness) EnumDescriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDescGZIP(), []int{1, 0}
}

type SeqRangeWithAssembly struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The genomic assembly associated with the sequence location of this protein
	AssemblyAccession string `protobuf:"bytes,1,opt,name=assembly_accession,json=assemblyAccession,proto3" json:"assembly_accession,omitempty"`
	// The genomic sequence location of this protein
	SequenceRange *SeqRangeSet `protobuf:"bytes,2,opt,name=sequence_range,json=sequenceRange,proto3" json:"sequence_range,omitempty"`
}

func (x *SeqRangeWithAssembly) Reset() {
	*x = SeqRangeWithAssembly{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeqRangeWithAssembly) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeqRangeWithAssembly) ProtoMessage() {}

func (x *SeqRangeWithAssembly) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeqRangeWithAssembly.ProtoReflect.Descriptor instead.
func (*SeqRangeWithAssembly) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDescGZIP(), []int{0}
}

func (x *SeqRangeWithAssembly) GetAssemblyAccession() string {
	if x != nil {
		return x.AssemblyAccession
	}
	return ""
}

func (x *SeqRangeWithAssembly) GetSequenceRange() *SeqRangeSet {
	if x != nil {
		return x.SequenceRange
	}
	return nil
}

type ProkaryoteGeneLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The RefSeq WP_ prefixed accession for the protein sequence.
	ProteinAccession string `protobuf:"bytes,1,opt,name=protein_accession,json=proteinAccession,proto3" json:"protein_accession,omitempty"`
	// The RefSeq nucleotide mapping for this protein
	RefseqGenomicLocation *SeqRangeWithAssembly `protobuf:"bytes,2,opt,name=refseq_genomic_location,json=refseqGenomicLocation,proto3" json:"refseq_genomic_location,omitempty"`
	// The equivalent GenBank nucleotide mapping for this protein
	GenbankGenomicLocation *SeqRangeWithAssembly `protobuf:"bytes,3,opt,name=genbank_genomic_location,json=genbankGenomicLocation,proto3" json:"genbank_genomic_location,omitempty"`
	// The species level taxonomy information
	Organism *Organism `protobuf:"bytes,4,opt,name=organism,proto3" json:"organism,omitempty"`
	// Whether the assembly is complete or partial
	Completeness ProkaryoteGeneLocation_Completeness `protobuf:"varint,5,opt,name=completeness,proto3,enum=ncbi.datasets.v1alpha1.reports.ProkaryoteGeneLocation_Completeness" json:"completeness,omitempty"`
	// The name of the chromosome, if there is one.
	ChromosomeName string `protobuf:"bytes,6,opt,name=chromosome_name,json=chromosomeName,proto3" json:"chromosome_name,omitempty"`
}

func (x *ProkaryoteGeneLocation) Reset() {
	*x = ProkaryoteGeneLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProkaryoteGeneLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProkaryoteGeneLocation) ProtoMessage() {}

func (x *ProkaryoteGeneLocation) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProkaryoteGeneLocation.ProtoReflect.Descriptor instead.
func (*ProkaryoteGeneLocation) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDescGZIP(), []int{1}
}

func (x *ProkaryoteGeneLocation) GetProteinAccession() string {
	if x != nil {
		return x.ProteinAccession
	}
	return ""
}

func (x *ProkaryoteGeneLocation) GetRefseqGenomicLocation() *SeqRangeWithAssembly {
	if x != nil {
		return x.RefseqGenomicLocation
	}
	return nil
}

func (x *ProkaryoteGeneLocation) GetGenbankGenomicLocation() *SeqRangeWithAssembly {
	if x != nil {
		return x.GenbankGenomicLocation
	}
	return nil
}

func (x *ProkaryoteGeneLocation) GetOrganism() *Organism {
	if x != nil {
		return x.Organism
	}
	return nil
}

func (x *ProkaryoteGeneLocation) GetCompleteness() ProkaryoteGeneLocation_Completeness {
	if x != nil {
		return x.Completeness
	}
	return ProkaryoteGeneLocation_complete
}

func (x *ProkaryoteGeneLocation) GetChromosomeName() string {
	if x != nil {
		return x.ChromosomeName
	}
	return ""
}

// Internal only message
type ProkaryoteGeneLocationDefline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GeneLocation *ProkaryoteGeneLocation `protobuf:"bytes,1,opt,name=gene_location,json=geneLocation,proto3" json:"gene_location,omitempty"`
	Name         string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gene         string                  `protobuf:"bytes,3,opt,name=gene,proto3" json:"gene,omitempty"`
}

func (x *ProkaryoteGeneLocationDefline) Reset() {
	*x = ProkaryoteGeneLocationDefline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProkaryoteGeneLocationDefline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProkaryoteGeneLocationDefline) ProtoMessage() {}

func (x *ProkaryoteGeneLocationDefline) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProkaryoteGeneLocationDefline.ProtoReflect.Descriptor instead.
func (*ProkaryoteGeneLocationDefline) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDescGZIP(), []int{2}
}

func (x *ProkaryoteGeneLocationDefline) GetGeneLocation() *ProkaryoteGeneLocation {
	if x != nil {
		return x.GeneLocation
	}
	return nil
}

func (x *ProkaryoteGeneLocationDefline) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProkaryoteGeneLocationDefline) GetGene() string {
	if x != nil {
		return x.Gene
	}
	return ""
}

var File_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto protoreflect.FileDescriptor

var file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79, 0x6f, 0x74, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a,
	0x2b, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6e, 0x63,
	0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a, 0x14, 0x53,
	0x65, 0x71, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x41, 0x73, 0x73, 0x65, 0x6d,
	0x62, 0x6c, 0x79, 0x12, 0x6c, 0x0a, 0x12, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x3d, 0xc2, 0xf3, 0x18, 0x39, 0x0a, 0x12, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79, 0x2d,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x41, 0x73, 0x73, 0x65, 0x6d,
	0x62, 0x6c, 0x79, 0x20, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x0f, 0x47,
	0x43, 0x46, 0x5f, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x33, 0x38, 0x35, 0x2e, 0x31, 0x52, 0x11,
	0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x64, 0x0a, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x63, 0x62, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x71, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x42, 0x10, 0xc2, 0xf3, 0x18, 0x0c, 0x0a, 0x0a, 0x73, 0x65,
	0x71, 0x2d, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x52, 0x0d, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0xda, 0x06, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x6b,
	0x61, 0x72, 0x79, 0x6f, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x84, 0x01, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x57,
	0xc2, 0xf3, 0x18, 0x36, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x2d, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e,
	0x20, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x0e, 0x57, 0x50, 0x5f, 0x30,
	0x30, 0x30, 0x34, 0x34, 0x33, 0x36, 0x36, 0x35, 0x2e, 0x31, 0xca, 0xf3, 0x18, 0x19, 0x08, 0x06,
	0x12, 0x15, 0x54, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x20, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0xa6, 0x01, 0x0a, 0x17, 0x72, 0x65,
	0x66, 0x73, 0x65, 0x71, 0x5f, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6e, 0x63,
	0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x71,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x41, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c,
	0x79, 0x42, 0x38, 0xc2, 0xf3, 0x18, 0x34, 0x0a, 0x18, 0x72, 0x65, 0x66, 0x73, 0x65, 0x71, 0x2d,
	0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x2d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2d, 0x12, 0x18, 0x52, 0x65, 0x66, 0x53, 0x65, 0x71, 0x20, 0x47, 0x65, 0x6e, 0x6f, 0x6d, 0x69,
	0x63, 0x20, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x52, 0x15, 0x72, 0x65, 0x66,
	0x73, 0x65, 0x71, 0x47, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0xaa, 0x01, 0x0a, 0x18, 0x67, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x67,
	0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x71, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x41, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79, 0x42, 0x3a, 0xc2, 0xf3, 0x18,
	0x36, 0x0a, 0x19, 0x67, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2d, 0x67, 0x65, 0x6e, 0x6f, 0x6d,
	0x69, 0x63, 0x2d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x12, 0x19, 0x47, 0x65,
	0x6e, 0x42, 0x61, 0x6e, 0x6b, 0x20, 0x47, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x20, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x52, 0x16, 0x67, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b,
	0x47, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x60, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x42, 0x1a, 0xc2, 0xf3, 0x18,
	0x16, 0x0a, 0x09, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x2d, 0x12, 0x09, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x20, 0x52, 0x08, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73,
	0x6d, 0x12, 0x89, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x43, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6b, 0x61, 0x72,
	0x79, 0x6f, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x20, 0xc2,
	0xf3, 0x18, 0x1c, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73,
	0x73, 0x12, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x4a, 0x0a,
	0x0f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x6f, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xc2, 0xf3, 0x18, 0x1d, 0x0a, 0x0f, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x6f, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0a, 0x43,
	0x68, 0x72, 0x6f, 0x6d, 0x6f, 0x73, 0x6f, 0x6d, 0x65, 0x52, 0x0e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x6f, 0x73, 0x6f, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0c, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x10, 0x01, 0x22, 0xd4, 0x01, 0x0a, 0x1d, 0x50, 0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79,
	0x6f, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x5b, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79, 0x6f, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x15, 0xca, 0xf3, 0x18, 0x11, 0x08, 0x08, 0x12, 0x0d, 0x54, 0x68, 0x65, 0x20, 0x67,
	0x65, 0x6e, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x04, 0x67, 0x65, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xca, 0xf3,
	0x18, 0x13, 0x08, 0x05, 0x12, 0x0f, 0x54, 0x68, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x20, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x04, 0x67, 0x65, 0x6e, 0x65, 0x42, 0xf2, 0x0a, 0x5a, 0x1e,
	0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0xf8, 0x01,
	0x01, 0xc2, 0xf3, 0x18, 0xca, 0x0a, 0x0a, 0x1f, 0x50, 0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79, 0x6f,
	0x74, 0x65, 0x20, 0x47, 0x65, 0x6e, 0x65, 0x20, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x50, 0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79,
	0x6f, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x57, 0x3c, 0x70, 0x3e, 0x54, 0x68, 0x65, 0x20, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x20, 0x70, 0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79, 0x6f, 0x74, 0x65, 0x20, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x20, 0x61,
	0x20, 0x70, 0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79, 0x6f, 0x74, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65,
	0x20, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x69, 0x6e, 0x1a, 0x3a, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65,
	0x66, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x74, 0x68, 0x65, 0x64, 0x6f, 0x63, 0x73,
	0x2e, 0x69, 0x6f, 0x2f, 0x22, 0x3e, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x4c, 0x69, 0x6e, 0x65, 0x73,
	0x3c, 0x2f, 0x61, 0x3e, 0x1a, 0x56, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x20, 0x69, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x3a, 0x3c, 0x62, 0x72, 0x3e, 0x3c, 0x62, 0x72,
	0x3e, 0x3c, 0x65, 0x6d, 0x3e, 0x20, 0x6e, 0x63, 0x62, 0x69, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x6c, 0x3c,
	0x2f, 0x65, 0x6d, 0x3e, 0x3c, 0x62, 0x72, 0x3e, 0x3c, 0x62, 0x72, 0x3e, 0x1a, 0x4c, 0x45, 0x61,
	0x63, 0x68, 0x20, 0x6c, 0x69, 0x6e, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70,
	0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79, 0x6f, 0x74, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x20, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x68, 0x69,
	0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x69, 0x63, 0x61, 0x6c, 0x1a, 0x34, 0x3c, 0x61, 0x20, 0x68,
	0x72, 0x65, 0x66, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x65,
	0x6e, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x22, 0x3e, 0x4a, 0x53, 0x4f, 0x4e, 0x3c, 0x2f, 0x61, 0x3e,
	0x1a, 0x5d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x72, 0x65,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x61, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x20, 0x70, 0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79, 0x6f, 0x74, 0x65, 0x20, 0x67, 0x65, 0x6e,
	0x65, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x20, 0x54, 0x68, 0x65, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x20, 0x6f, 0x66,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79, 0x6f, 0x74, 0x65, 0x1a,
	0x5e, 0x67, 0x65, 0x6e, 0x65, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x20, 0x69, 0x73, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64,
	0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x20, 0x62,
	0x65, 0x6c, 0x6f, 0x77, 0x20, 0x77, 0x68, 0x65, 0x72, 0x65, 0x20, 0x65, 0x61, 0x63, 0x68, 0x20,
	0x72, 0x6f, 0x77, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x73, 0x20, 0x61, 0x20,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x69, 0x6e, 0x1a,
	0x3f, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x6f, 0x72, 0x20, 0x61,
	0x20, 0x73, 0x75, 0x62, 0x2d, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2c, 0x20,
	0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e,
	0x1a, 0x4d, 0x54, 0x68, 0x65, 0x20, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x74, 0x20,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x69, 0x73, 0x20, 0x3c, 0x65, 0x6d, 0x3e, 0x50,
	0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79, 0x6f, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3c, 0x2f, 0x65, 0x6d, 0x3e, 0x2e, 0x3c, 0x2f, 0x70, 0x3e, 0x1a,
	0x51, 0x3c, 0x70, 0x3e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x20, 0x61, 0x20,
	0x3c, 0x65, 0x6d, 0x3e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x20,
	0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x3c, 0x2f, 0x65, 0x6d, 0x3e, 0x20, 0x63, 0x61,
	0x6e, 0x20, 0x62, 0x65, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74,
	0x68, 0x65, 0x1a, 0x51, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x64, 0x6f,
	0x63, 0x73, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2d, 0x64, 0x6f, 0x63,
	0x73, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x2f, 0x22, 0x3e, 0x64, 0x61, 0x74, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3c, 0x2f, 0x61,
	0x3e, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2d, 0x6c, 0x69, 0x6e, 0x65, 0x20, 0x74,
	0x6f, 0x6f, 0x6c, 0x27, 0x73, 0x1a, 0x37, 0x3c, 0x6e, 0x6f, 0x62, 0x72, 0x3e, 0x3c, 0x63, 0x6f,
	0x64, 0x65, 0x3e, 0x2d, 0x2d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3c, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x3e, 0x3c, 0x2f, 0x6e, 0x6f, 0x62, 0x72, 0x3e, 0x20, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x20, 0x52, 0x65, 0x66, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x1a, 0x5f,
	0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2d, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x63, 0x6c,
	0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2f, 0x22, 0x3e, 0x64,
	0x61, 0x74, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3c, 0x2f, 0x61, 0x3e, 0x20, 0x43, 0x4c,
	0x49, 0x20, 0x74, 0x6f, 0x6f, 0x6c, 0x20, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x20, 0x74, 0x6f, 0x20, 0x73, 0x65, 0x65, 0x20, 0x68, 0x6f, 0x77, 0x20, 0x79, 0x6f, 0x75, 0x1a,
	0x6c, 0x63, 0x61, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x74, 0x6f,
	0x6f, 0x6c, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x20,
	0x70, 0x72, 0x6f, 0x6b, 0x61, 0x72, 0x79, 0x6f, 0x74, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x20,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20,
	0x4c, 0x69, 0x6e, 0x65, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72,
	0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x2e, 0x3c, 0x2f, 0x70, 0x3e, 0x2a, 0xa0, 0x01,
	0x7b, 0x7b, 0x3c, 0x20, 0x72, 0x65, 0x61, 0x64, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x63, 0x6f, 0x64,
	0x65, 0x2d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x3d, 0x22, 0x6a, 0x73,
	0x6f, 0x6e, 0x22, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x3d, 0x22, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x5f,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6b, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x2d, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x20, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x3d, 0x22, 0x53, 0x54, 0x41, 0x52, 0x54, 0x20, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x20, 0x65, 0x6e, 0x64, 0x3d, 0x22, 0x45, 0x4e, 0x44, 0x20, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x3e, 0x7d, 0x7d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDescOnce sync.Once
	file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDescData = file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDesc
)

func file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDescGZIP() []byte {
	file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDescOnce.Do(func() {
		file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDescData)
	})
	return file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDescData
}

var file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_goTypes = []interface{}{
	(ProkaryoteGeneLocation_Completeness)(0), // 0: ncbi.datasets.v1alpha1.reports.ProkaryoteGeneLocation.Completeness
	(*SeqRangeWithAssembly)(nil),             // 1: ncbi.datasets.v1alpha1.reports.SeqRangeWithAssembly
	(*ProkaryoteGeneLocation)(nil),           // 2: ncbi.datasets.v1alpha1.reports.ProkaryoteGeneLocation
	(*ProkaryoteGeneLocationDefline)(nil),    // 3: ncbi.datasets.v1alpha1.reports.ProkaryoteGeneLocationDefline
	(*SeqRangeSet)(nil),                      // 4: ncbi.datasets.v1alpha1.reports.SeqRangeSet
	(*Organism)(nil),                         // 5: ncbi.datasets.v1alpha1.reports.Organism
}
var file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_depIdxs = []int32{
	4, // 0: ncbi.datasets.v1alpha1.reports.SeqRangeWithAssembly.sequence_range:type_name -> ncbi.datasets.v1alpha1.reports.SeqRangeSet
	1, // 1: ncbi.datasets.v1alpha1.reports.ProkaryoteGeneLocation.refseq_genomic_location:type_name -> ncbi.datasets.v1alpha1.reports.SeqRangeWithAssembly
	1, // 2: ncbi.datasets.v1alpha1.reports.ProkaryoteGeneLocation.genbank_genomic_location:type_name -> ncbi.datasets.v1alpha1.reports.SeqRangeWithAssembly
	5, // 3: ncbi.datasets.v1alpha1.reports.ProkaryoteGeneLocation.organism:type_name -> ncbi.datasets.v1alpha1.reports.Organism
	0, // 4: ncbi.datasets.v1alpha1.reports.ProkaryoteGeneLocation.completeness:type_name -> ncbi.datasets.v1alpha1.reports.ProkaryoteGeneLocation.Completeness
	2, // 5: ncbi.datasets.v1alpha1.reports.ProkaryoteGeneLocationDefline.gene_location:type_name -> ncbi.datasets.v1alpha1.reports.ProkaryoteGeneLocation
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_init() }
func file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_init() {
	if File_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto != nil {
		return
	}
	file_ncbi_datasets_v1alpha1_reports_common_proto_init()
	file_ncbi_datasets_v1alpha1_reports_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeqRangeWithAssembly); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProkaryoteGeneLocation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProkaryoteGeneLocationDefline); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_goTypes,
		DependencyIndexes: file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_depIdxs,
		EnumInfos:         file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_enumTypes,
		MessageInfos:      file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_msgTypes,
	}.Build()
	File_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto = out.File
	file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_rawDesc = nil
	file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_goTypes = nil
	file_ncbi_datasets_v1alpha1_reports_prokaryote_gene_location_proto_depIdxs = nil
}
