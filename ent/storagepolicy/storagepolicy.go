// Code generated by ent, DO NOT EDIT.

package storagepolicy

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cloudreve/Cloudreve/v4/inventory/types"
)

const (
	// Label holds the string label denoting the storagepolicy type in the database.
	Label = "storage_policy"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldServer holds the string denoting the server field in the database.
	FieldServer = "server"
	// FieldBucketName holds the string denoting the bucket_name field in the database.
	FieldBucketName = "bucket_name"
	// FieldIsPrivate holds the string denoting the is_private field in the database.
	FieldIsPrivate = "is_private"
	// FieldAccessKey holds the string denoting the access_key field in the database.
	FieldAccessKey = "access_key"
	// FieldSecretKey holds the string denoting the secret_key field in the database.
	FieldSecretKey = "secret_key"
	// FieldMaxSize holds the string denoting the max_size field in the database.
	FieldMaxSize = "max_size"
	// FieldDirNameRule holds the string denoting the dir_name_rule field in the database.
	FieldDirNameRule = "dir_name_rule"
	// FieldFileNameRule holds the string denoting the file_name_rule field in the database.
	FieldFileNameRule = "file_name_rule"
	// FieldSettings holds the string denoting the settings field in the database.
	FieldSettings = "settings"
	// FieldNodeID holds the string denoting the node_id field in the database.
	FieldNodeID = "node_id"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// EdgeEntities holds the string denoting the entities edge name in mutations.
	EdgeEntities = "entities"
	// EdgeNode holds the string denoting the node edge name in mutations.
	EdgeNode = "node"
	// Table holds the table name of the storagepolicy in the database.
	Table = "storage_policies"
	// GroupsTable is the table that holds the groups relation/edge.
	GroupsTable = "groups"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// GroupsColumn is the table column denoting the groups relation/edge.
	GroupsColumn = "storage_policy_id"
	// FilesTable is the table that holds the files relation/edge.
	FilesTable = "files"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "files"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "storage_policy_files"
	// EntitiesTable is the table that holds the entities relation/edge.
	EntitiesTable = "entities"
	// EntitiesInverseTable is the table name for the Entity entity.
	// It exists in this package in order to avoid circular dependency with the "entity" package.
	EntitiesInverseTable = "entities"
	// EntitiesColumn is the table column denoting the entities relation/edge.
	EntitiesColumn = "storage_policy_entities"
	// NodeTable is the table that holds the node relation/edge.
	NodeTable = "storage_policies"
	// NodeInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	NodeInverseTable = "nodes"
	// NodeColumn is the table column denoting the node relation/edge.
	NodeColumn = "node_id"
)

// Columns holds all SQL columns for storagepolicy fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldType,
	FieldServer,
	FieldBucketName,
	FieldIsPrivate,
	FieldAccessKey,
	FieldSecretKey,
	FieldMaxSize,
	FieldDirNameRule,
	FieldFileNameRule,
	FieldSettings,
	FieldNodeID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/cloudreve/Cloudreve/v4/ent/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultSettings holds the default value on creation for the "settings" field.
	DefaultSettings *types.PolicySetting
)

// OrderOption defines the ordering options for the StoragePolicy queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByServer orders the results by the server field.
func ByServer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldServer, opts...).ToFunc()
}

// ByBucketName orders the results by the bucket_name field.
func ByBucketName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBucketName, opts...).ToFunc()
}

// ByIsPrivate orders the results by the is_private field.
func ByIsPrivate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPrivate, opts...).ToFunc()
}

// ByAccessKey orders the results by the access_key field.
func ByAccessKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccessKey, opts...).ToFunc()
}

// BySecretKey orders the results by the secret_key field.
func BySecretKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecretKey, opts...).ToFunc()
}

// ByMaxSize orders the results by the max_size field.
func ByMaxSize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxSize, opts...).ToFunc()
}

// ByDirNameRule orders the results by the dir_name_rule field.
func ByDirNameRule(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDirNameRule, opts...).ToFunc()
}

// ByFileNameRule orders the results by the file_name_rule field.
func ByFileNameRule(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileNameRule, opts...).ToFunc()
}

// ByNodeID orders the results by the node_id field.
func ByNodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNodeID, opts...).ToFunc()
}

// ByGroupsCount orders the results by groups count.
func ByGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupsStep(), opts...)
	}
}

// ByGroups orders the results by groups terms.
func ByGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFilesCount orders the results by files count.
func ByFilesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFilesStep(), opts...)
	}
}

// ByFiles orders the results by files terms.
func ByFiles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFilesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEntitiesCount orders the results by entities count.
func ByEntitiesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEntitiesStep(), opts...)
	}
}

// ByEntities orders the results by entities terms.
func ByEntities(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEntitiesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNodeField orders the results by node field.
func ByNodeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodeStep(), sql.OrderByField(field, opts...))
	}
}
func newGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, GroupsTable, GroupsColumn),
	)
}
func newFilesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FilesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
	)
}
func newEntitiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EntitiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EntitiesTable, EntitiesColumn),
	)
}
func newNodeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
	)
}
