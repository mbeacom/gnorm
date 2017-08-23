// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// ViewColumnUsageTable is the database name for the table.
const ViewColumnUsageTable = "information_schema.view_column_usage"

// ViewColumnUsage represents a row from 'information_schema.view_column_usage'.
type ViewColumnUsage struct {
	ViewCatalog  SQLIdentifier `json:"view_catalog"`  // view_catalog
	ViewSchema   SQLIdentifier `json:"view_schema"`   // view_schema
	ViewName     SQLIdentifier `json:"view_name"`     // view_name
	TableCatalog SQLIdentifier `json:"table_catalog"` // table_catalog
	TableSchema  SQLIdentifier `json:"table_schema"`  // table_schema
	TableName    SQLIdentifier `json:"table_name"`    // table_name
	ColumnName   SQLIdentifier `json:"column_name"`   // column_name
}

// Constants defining each column in the table.
const (
	ViewColumnUsageViewCatalogField  = "view_catalog"
	ViewColumnUsageViewSchemaField   = "view_schema"
	ViewColumnUsageViewNameField     = "view_name"
	ViewColumnUsageTableCatalogField = "table_catalog"
	ViewColumnUsageTableSchemaField  = "table_schema"
	ViewColumnUsageTableNameField    = "table_name"
	ViewColumnUsageColumnNameField   = "column_name"
)

// WhereClauses for every type in ViewColumnUsage.
var (
	ViewColumnUsageViewCatalogWhere  SQLIdentifierField = "view_catalog"
	ViewColumnUsageViewSchemaWhere   SQLIdentifierField = "view_schema"
	ViewColumnUsageViewNameWhere     SQLIdentifierField = "view_name"
	ViewColumnUsageTableCatalogWhere SQLIdentifierField = "table_catalog"
	ViewColumnUsageTableSchemaWhere  SQLIdentifierField = "table_schema"
	ViewColumnUsageTableNameWhere    SQLIdentifierField = "table_name"
	ViewColumnUsageColumnNameWhere   SQLIdentifierField = "column_name"
)

// QueryOneViewColumnUsage retrieves a row from 'information_schema.view_column_usage' as a ViewColumnUsage.
func QueryOneViewColumnUsage(db XODB, where WhereClause, order OrderBy) (*ViewColumnUsage, error) {
	const origsqlstr = `SELECT ` +
		`view_catalog, view_schema, view_name, table_catalog, table_schema, table_name, column_name ` +
		`FROM information_schema.view_column_usage WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	vcu := &ViewColumnUsage{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&vcu.ViewCatalog, &vcu.ViewSchema, &vcu.ViewName, &vcu.TableCatalog, &vcu.TableSchema, &vcu.TableName, &vcu.ColumnName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return vcu, nil
}

// QueryViewColumnUsage retrieves rows from 'information_schema.view_column_usage' as a slice of ViewColumnUsage.
func QueryViewColumnUsage(db XODB, where WhereClause, order OrderBy) ([]*ViewColumnUsage, error) {
	const origsqlstr = `SELECT ` +
		`view_catalog, view_schema, view_name, table_catalog, table_schema, table_name, column_name ` +
		`FROM information_schema.view_column_usage WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*ViewColumnUsage
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		vcu := ViewColumnUsage{}

		err = q.Scan(&vcu.ViewCatalog, &vcu.ViewSchema, &vcu.ViewName, &vcu.TableCatalog, &vcu.TableSchema, &vcu.TableName, &vcu.ColumnName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &vcu)
	}
	return vals, nil
}

// AllViewColumnUsage retrieves all rows from 'information_schema.view_column_usage' as a slice of ViewColumnUsage.
func AllViewColumnUsage(db XODB, order OrderBy) ([]*ViewColumnUsage, error) {
	const origsqlstr = `SELECT ` +
		`view_catalog, view_schema, view_name, table_catalog, table_schema, table_name, column_name ` +
		`FROM information_schema.view_column_usage`

	sqlstr := origsqlstr + order.String()

	var vals []*ViewColumnUsage
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		vcu := ViewColumnUsage{}

		err = q.Scan(&vcu.ViewCatalog, &vcu.ViewSchema, &vcu.ViewName, &vcu.TableCatalog, &vcu.TableSchema, &vcu.TableName, &vcu.ColumnName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &vcu)
	}
	return vals, nil
}