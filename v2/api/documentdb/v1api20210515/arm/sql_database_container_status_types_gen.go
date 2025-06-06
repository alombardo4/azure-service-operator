// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type SqlDatabaseContainer_STATUS struct {
	// Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of an Azure Cosmos DB container
	Properties *SqlContainerGetProperties_STATUS `json:"properties,omitempty"`
	Tags       map[string]string                 `json:"tags,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

// The properties of an Azure Cosmos DB container
type SqlContainerGetProperties_STATUS struct {
	// Options: Cosmos DB options resource object
	Options  *OptionsResource_STATUS                    `json:"options,omitempty"`
	Resource *SqlContainerGetProperties_Resource_STATUS `json:"resource,omitempty"`
}

type SqlContainerGetProperties_Resource_STATUS struct {
	// AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	// ConflictResolutionPolicy: The conflict resolution policy for the container.
	ConflictResolutionPolicy *ConflictResolutionPolicy_STATUS `json:"conflictResolutionPolicy,omitempty"`

	// DefaultTtl: Default time to live
	DefaultTtl *int `json:"defaultTtl,omitempty"`

	// Etag: A system generated property representing the resource etag required for optimistic concurrency control.
	Etag *string `json:"_etag,omitempty"`

	// Id: Name of the Cosmos DB SQL container
	Id *string `json:"id,omitempty"`

	// IndexingPolicy: The configuration of the indexing policy. By default, the indexing is automatic for all document paths
	// within the container
	IndexingPolicy *IndexingPolicy_STATUS `json:"indexingPolicy,omitempty"`

	// PartitionKey: The configuration of the partition key to be used for partitioning data into multiple partitions
	PartitionKey *ContainerPartitionKey_STATUS `json:"partitionKey,omitempty"`

	// Rid: A system generated property. A unique identifier.
	Rid *string `json:"_rid,omitempty"`

	// Ts: A system generated property that denotes the last updated timestamp of the resource.
	Ts *float64 `json:"_ts,omitempty"`

	// UniqueKeyPolicy: The unique key policy configuration for specifying uniqueness constraints on documents in the
	// collection in the Azure Cosmos DB service.
	UniqueKeyPolicy *UniqueKeyPolicy_STATUS `json:"uniqueKeyPolicy,omitempty"`
}

// The conflict resolution policy for the container.
type ConflictResolutionPolicy_STATUS struct {
	// ConflictResolutionPath: The conflict resolution path in the case of LastWriterWins mode.
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty"`

	// ConflictResolutionProcedure: The procedure to resolve conflicts in the case of custom mode.
	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty"`

	// Mode: Indicates the conflict resolution mode.
	Mode *ConflictResolutionPolicy_Mode_STATUS `json:"mode,omitempty"`
}

// The configuration of the partition key to be used for partitioning data into multiple partitions
type ContainerPartitionKey_STATUS struct {
	// Kind: Indicates the kind of algorithm used for partitioning. For MultiHash, multiple partition keys (upto three maximum)
	// are supported for container create
	Kind *ContainerPartitionKey_Kind_STATUS `json:"kind,omitempty"`

	// Paths: List of paths using which data within the container can be partitioned
	Paths []string `json:"paths,omitempty"`

	// SystemKey: Indicates if the container is using a system generated partition key
	SystemKey *bool `json:"systemKey,omitempty"`

	// Version: Indicates the version of the partition key definition
	Version *int `json:"version,omitempty"`
}

// Cosmos DB indexing policy
type IndexingPolicy_STATUS struct {
	// Automatic: Indicates if the indexing policy is automatic
	Automatic *bool `json:"automatic,omitempty"`

	// CompositeIndexes: List of composite path list
	CompositeIndexes [][]CompositePath_STATUS `json:"compositeIndexes,omitempty"`

	// ExcludedPaths: List of paths to exclude from indexing
	ExcludedPaths []ExcludedPath_STATUS `json:"excludedPaths,omitempty"`

	// IncludedPaths: List of paths to include in the indexing
	IncludedPaths []IncludedPath_STATUS `json:"includedPaths,omitempty"`

	// IndexingMode: Indicates the indexing mode.
	IndexingMode *IndexingPolicy_IndexingMode_STATUS `json:"indexingMode,omitempty"`

	// SpatialIndexes: List of spatial specifics
	SpatialIndexes []SpatialSpec_STATUS `json:"spatialIndexes,omitempty"`
}

// The unique key policy configuration for specifying uniqueness constraints on documents in the collection in the Azure
// Cosmos DB service.
type UniqueKeyPolicy_STATUS struct {
	// UniqueKeys: List of unique keys on that enforces uniqueness constraint on documents in the collection in the Azure
	// Cosmos DB service.
	UniqueKeys []UniqueKey_STATUS `json:"uniqueKeys,omitempty"`
}

type CompositePath_STATUS struct {
	// Order: Sort order for composite paths.
	Order *CompositePath_Order_STATUS `json:"order,omitempty"`

	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

type ConflictResolutionPolicy_Mode_STATUS string

const (
	ConflictResolutionPolicy_Mode_STATUS_Custom         = ConflictResolutionPolicy_Mode_STATUS("Custom")
	ConflictResolutionPolicy_Mode_STATUS_LastWriterWins = ConflictResolutionPolicy_Mode_STATUS("LastWriterWins")
)

// Mapping from string to ConflictResolutionPolicy_Mode_STATUS
var conflictResolutionPolicy_Mode_STATUS_Values = map[string]ConflictResolutionPolicy_Mode_STATUS{
	"custom":         ConflictResolutionPolicy_Mode_STATUS_Custom,
	"lastwriterwins": ConflictResolutionPolicy_Mode_STATUS_LastWriterWins,
}

type ContainerPartitionKey_Kind_STATUS string

const (
	ContainerPartitionKey_Kind_STATUS_Hash      = ContainerPartitionKey_Kind_STATUS("Hash")
	ContainerPartitionKey_Kind_STATUS_MultiHash = ContainerPartitionKey_Kind_STATUS("MultiHash")
	ContainerPartitionKey_Kind_STATUS_Range     = ContainerPartitionKey_Kind_STATUS("Range")
)

// Mapping from string to ContainerPartitionKey_Kind_STATUS
var containerPartitionKey_Kind_STATUS_Values = map[string]ContainerPartitionKey_Kind_STATUS{
	"hash":      ContainerPartitionKey_Kind_STATUS_Hash,
	"multihash": ContainerPartitionKey_Kind_STATUS_MultiHash,
	"range":     ContainerPartitionKey_Kind_STATUS_Range,
}

type ExcludedPath_STATUS struct {
	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

// The paths that are included in indexing
type IncludedPath_STATUS struct {
	// Indexes: List of indexes for this path
	Indexes []Indexes_STATUS `json:"indexes,omitempty"`

	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

type IndexingPolicy_IndexingMode_STATUS string

const (
	IndexingPolicy_IndexingMode_STATUS_Consistent = IndexingPolicy_IndexingMode_STATUS("consistent")
	IndexingPolicy_IndexingMode_STATUS_Lazy       = IndexingPolicy_IndexingMode_STATUS("lazy")
	IndexingPolicy_IndexingMode_STATUS_None       = IndexingPolicy_IndexingMode_STATUS("none")
)

// Mapping from string to IndexingPolicy_IndexingMode_STATUS
var indexingPolicy_IndexingMode_STATUS_Values = map[string]IndexingPolicy_IndexingMode_STATUS{
	"consistent": IndexingPolicy_IndexingMode_STATUS_Consistent,
	"lazy":       IndexingPolicy_IndexingMode_STATUS_Lazy,
	"none":       IndexingPolicy_IndexingMode_STATUS_None,
}

type SpatialSpec_STATUS struct {
	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`

	// Types: List of path's spatial type
	Types []SpatialType_STATUS `json:"types,omitempty"`
}

// The unique key on that enforces uniqueness constraint on documents in the collection in the Azure Cosmos DB service.
type UniqueKey_STATUS struct {
	// Paths: List of paths must be unique for each document in the Azure Cosmos DB service
	Paths []string `json:"paths,omitempty"`
}

type CompositePath_Order_STATUS string

const (
	CompositePath_Order_STATUS_Ascending  = CompositePath_Order_STATUS("ascending")
	CompositePath_Order_STATUS_Descending = CompositePath_Order_STATUS("descending")
)

// Mapping from string to CompositePath_Order_STATUS
var compositePath_Order_STATUS_Values = map[string]CompositePath_Order_STATUS{
	"ascending":  CompositePath_Order_STATUS_Ascending,
	"descending": CompositePath_Order_STATUS_Descending,
}

// The indexes for the path.
type Indexes_STATUS struct {
	// DataType: The datatype for which the indexing behavior is applied to.
	DataType *Indexes_DataType_STATUS `json:"dataType,omitempty"`

	// Kind: Indicates the type of index.
	Kind *Indexes_Kind_STATUS `json:"kind,omitempty"`

	// Precision: The precision of the index. -1 is maximum precision.
	Precision *int `json:"precision,omitempty"`
}

// Indicates the spatial type of index.
type SpatialType_STATUS string

const (
	SpatialType_STATUS_LineString   = SpatialType_STATUS("LineString")
	SpatialType_STATUS_MultiPolygon = SpatialType_STATUS("MultiPolygon")
	SpatialType_STATUS_Point        = SpatialType_STATUS("Point")
	SpatialType_STATUS_Polygon      = SpatialType_STATUS("Polygon")
)

// Mapping from string to SpatialType_STATUS
var spatialType_STATUS_Values = map[string]SpatialType_STATUS{
	"linestring":   SpatialType_STATUS_LineString,
	"multipolygon": SpatialType_STATUS_MultiPolygon,
	"point":        SpatialType_STATUS_Point,
	"polygon":      SpatialType_STATUS_Polygon,
}

type Indexes_DataType_STATUS string

const (
	Indexes_DataType_STATUS_LineString   = Indexes_DataType_STATUS("LineString")
	Indexes_DataType_STATUS_MultiPolygon = Indexes_DataType_STATUS("MultiPolygon")
	Indexes_DataType_STATUS_Number       = Indexes_DataType_STATUS("Number")
	Indexes_DataType_STATUS_Point        = Indexes_DataType_STATUS("Point")
	Indexes_DataType_STATUS_Polygon      = Indexes_DataType_STATUS("Polygon")
	Indexes_DataType_STATUS_String       = Indexes_DataType_STATUS("String")
)

// Mapping from string to Indexes_DataType_STATUS
var indexes_DataType_STATUS_Values = map[string]Indexes_DataType_STATUS{
	"linestring":   Indexes_DataType_STATUS_LineString,
	"multipolygon": Indexes_DataType_STATUS_MultiPolygon,
	"number":       Indexes_DataType_STATUS_Number,
	"point":        Indexes_DataType_STATUS_Point,
	"polygon":      Indexes_DataType_STATUS_Polygon,
	"string":       Indexes_DataType_STATUS_String,
}

type Indexes_Kind_STATUS string

const (
	Indexes_Kind_STATUS_Hash    = Indexes_Kind_STATUS("Hash")
	Indexes_Kind_STATUS_Range   = Indexes_Kind_STATUS("Range")
	Indexes_Kind_STATUS_Spatial = Indexes_Kind_STATUS("Spatial")
)

// Mapping from string to Indexes_Kind_STATUS
var indexes_Kind_STATUS_Values = map[string]Indexes_Kind_STATUS{
	"hash":    Indexes_Kind_STATUS_Hash,
	"range":   Indexes_Kind_STATUS_Range,
	"spatial": Indexes_Kind_STATUS_Spatial,
}
