// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// CollationCharacterSetApplicabilityTable is the database name for the table.
const CollationCharacterSetApplicabilityTable = "information_schema.collation_character_set_applicability"

// CollationCharacterSetApplicability represents a row from 'information_schema.collation_character_set_applicability'.
type CollationCharacterSetApplicability struct {
	CollationCatalog    SQLIdentifier `yaml:"collation_catalog,omitempty"`     // collation_catalog
	CollationSchema     SQLIdentifier `yaml:"collation_schema,omitempty"`      // collation_schema
	CollationName       SQLIdentifier `yaml:"collation_name,omitempty"`        // collation_name
	CharacterSetCatalog SQLIdentifier `yaml:"character_set_catalog,omitempty"` // character_set_catalog
	CharacterSetSchema  SQLIdentifier `yaml:"character_set_schema,omitempty"`  // character_set_schema
	CharacterSetName    SQLIdentifier `yaml:"character_set_name,omitempty"`    // character_set_name
}

// Constants defining each column in the table.
const (
	CollationCharacterSetApplicabilityCollationCatalogField    = "collation_catalog"
	CollationCharacterSetApplicabilityCollationSchemaField     = "collation_schema"
	CollationCharacterSetApplicabilityCollationNameField       = "collation_name"
	CollationCharacterSetApplicabilityCharacterSetCatalogField = "character_set_catalog"
	CollationCharacterSetApplicabilityCharacterSetSchemaField  = "character_set_schema"
	CollationCharacterSetApplicabilityCharacterSetNameField    = "character_set_name"
)

// WhereClauses for every type in CollationCharacterSetApplicability.
var (
	CollationCharacterSetApplicabilityCollationCatalogWhere    SQLIdentifierField = "collation_catalog"
	CollationCharacterSetApplicabilityCollationSchemaWhere     SQLIdentifierField = "collation_schema"
	CollationCharacterSetApplicabilityCollationNameWhere       SQLIdentifierField = "collation_name"
	CollationCharacterSetApplicabilityCharacterSetCatalogWhere SQLIdentifierField = "character_set_catalog"
	CollationCharacterSetApplicabilityCharacterSetSchemaWhere  SQLIdentifierField = "character_set_schema"
	CollationCharacterSetApplicabilityCharacterSetNameWhere    SQLIdentifierField = "character_set_name"
)

// QueryOneCollationCharacterSetApplicability retrieves a row from 'information_schema.collation_character_set_applicability' as a CollationCharacterSetApplicability.
func QueryOneCollationCharacterSetApplicability(db XODB, where WhereClause, order OrderBy) (*CollationCharacterSetApplicability, error) {
	const origsqlstr = `SELECT ` +
		`collation_catalog, collation_schema, collation_name, character_set_catalog, character_set_schema, character_set_name ` +
		`FROM information_schema.collation_character_set_applicability WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	ccsa := &CollationCharacterSetApplicability{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&ccsa.CollationCatalog, &ccsa.CollationSchema, &ccsa.CollationName, &ccsa.CharacterSetCatalog, &ccsa.CharacterSetSchema, &ccsa.CharacterSetName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ccsa, nil
}

// QueryCollationCharacterSetApplicability retrieves rows from 'information_schema.collation_character_set_applicability' as a slice of CollationCharacterSetApplicability.
func QueryCollationCharacterSetApplicability(db XODB, where WhereClause, order OrderBy) ([]*CollationCharacterSetApplicability, error) {
	const origsqlstr = `SELECT ` +
		`collation_catalog, collation_schema, collation_name, character_set_catalog, character_set_schema, character_set_name ` +
		`FROM information_schema.collation_character_set_applicability WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*CollationCharacterSetApplicability
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		ccsa := CollationCharacterSetApplicability{}

		err = q.Scan(&ccsa.CollationCatalog, &ccsa.CollationSchema, &ccsa.CollationName, &ccsa.CharacterSetCatalog, &ccsa.CharacterSetSchema, &ccsa.CharacterSetName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &ccsa)
	}
	return vals, nil
}

// AllCollationCharacterSetApplicability retrieves all rows from 'information_schema.collation_character_set_applicability' as a slice of CollationCharacterSetApplicability.
func AllCollationCharacterSetApplicability(db XODB, order OrderBy) ([]*CollationCharacterSetApplicability, error) {
	const origsqlstr = `SELECT ` +
		`collation_catalog, collation_schema, collation_name, character_set_catalog, character_set_schema, character_set_name ` +
		`FROM information_schema.collation_character_set_applicability`

	sqlstr := origsqlstr + order.String()

	var vals []*CollationCharacterSetApplicability
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		ccsa := CollationCharacterSetApplicability{}

		err = q.Scan(&ccsa.CollationCatalog, &ccsa.CollationSchema, &ccsa.CollationName, &ccsa.CharacterSetCatalog, &ccsa.CharacterSetSchema, &ccsa.CharacterSetName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &ccsa)
	}
	return vals, nil
}