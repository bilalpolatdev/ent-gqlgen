// Code generated by ent, DO NOT EDIT.

package companycareer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the companycareer type in the database.
	Label = "company_career"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCareer holds the string denoting the career field in the database.
	FieldCareer = "career"
	// EdgeEngineerCareers holds the string denoting the engineercareers edge name in mutations.
	EdgeEngineerCareers = "engineerCareers"
	// EdgeCompanyOwnerCareers holds the string denoting the companyownercareers edge name in mutations.
	EdgeCompanyOwnerCareers = "companyOwnerCareers"
	// Table holds the table name of the companycareer in the database.
	Table = "company_careers"
	// EngineerCareersTable is the table that holds the engineerCareers relation/edge.
	EngineerCareersTable = "company_engineers"
	// EngineerCareersInverseTable is the table name for the CompanyEngineer entity.
	// It exists in this package in order to avoid circular dependency with the "companyengineer" package.
	EngineerCareersInverseTable = "company_engineers"
	// EngineerCareersColumn is the table column denoting the engineerCareers relation/edge.
	EngineerCareersColumn = "career_id"
	// CompanyOwnerCareersTable is the table that holds the companyOwnerCareers relation/edge.
	CompanyOwnerCareersTable = "company_owners"
	// CompanyOwnerCareersInverseTable is the table name for the CompanyOwner entity.
	// It exists in this package in order to avoid circular dependency with the "companyowner" package.
	CompanyOwnerCareersInverseTable = "company_owners"
	// CompanyOwnerCareersColumn is the table column denoting the companyOwnerCareers relation/edge.
	CompanyOwnerCareersColumn = "career_id"
)

// Columns holds all SQL columns for companycareer fields.
var Columns = []string{
	FieldID,
	FieldCareer,
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

// OrderOption defines the ordering options for the CompanyCareer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCareer orders the results by the Career field.
func ByCareer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCareer, opts...).ToFunc()
}

// ByEngineerCareersCount orders the results by engineerCareers count.
func ByEngineerCareersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEngineerCareersStep(), opts...)
	}
}

// ByEngineerCareers orders the results by engineerCareers terms.
func ByEngineerCareers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEngineerCareersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCompanyOwnerCareersCount orders the results by companyOwnerCareers count.
func ByCompanyOwnerCareersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCompanyOwnerCareersStep(), opts...)
	}
}

// ByCompanyOwnerCareers orders the results by companyOwnerCareers terms.
func ByCompanyOwnerCareers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyOwnerCareersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEngineerCareersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EngineerCareersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EngineerCareersTable, EngineerCareersColumn),
	)
}
func newCompanyOwnerCareersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyOwnerCareersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CompanyOwnerCareersTable, CompanyOwnerCareersColumn),
	)
}