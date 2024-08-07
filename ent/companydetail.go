// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"gqlgen-ent/ent/companydetail"
	"gqlgen-ent/ent/companyengineer"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// CompanyDetail is the model entity for the CompanyDetail schema.
type CompanyDetail struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CompanyCode holds the value of the "CompanyCode" field.
	CompanyCode int `json:"CompanyCode,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Address holds the value of the "Address" field.
	Address string `json:"Address,omitempty"`
	// City holds the value of the "City" field.
	City string `json:"City,omitempty"`
	// State holds the value of the "State" field.
	State string `json:"State,omitempty"`
	// Phone holds the value of the "Phone" field.
	Phone string `json:"Phone,omitempty"`
	// Fax holds the value of the "Fax" field.
	Fax string `json:"Fax,omitempty"`
	// Mobile holds the value of the "Mobile" field.
	Mobile string `json:"Mobile,omitempty"`
	// Email holds the value of the "Email" field.
	Email string `json:"Email,omitempty"`
	// Website holds the value of the "Website" field.
	Website string `json:"Website,omitempty"`
	// TaxAdmin holds the value of the "TaxAdmin" field.
	TaxAdmin string `json:"TaxAdmin,omitempty"`
	// TaxNo holds the value of the "TaxNo" field.
	TaxNo int `json:"TaxNo,omitempty"`
	// Commerce holds the value of the "Commerce" field.
	Commerce string `json:"Commerce,omitempty"`
	// CommerceReg holds the value of the "CommerceReg" field.
	CommerceReg string `json:"CommerceReg,omitempty"`
	// VisaDate holds the value of the "VisaDate" field.
	VisaDate time.Time `json:"VisaDate,omitempty"`
	// Deleted holds the value of the "Deleted" field.
	Deleted int `json:"Deleted,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompanyDetailQuery when eager-loading is set.
	Edges        CompanyDetailEdges `json:"edges"`
	owner_id     *int
	selectValues sql.SelectValues
}

// CompanyDetailEdges holds the relations/edges for other nodes in the graph.
type CompanyDetailEdges struct {
	// CompanyOwner holds the value of the companyOwner edge.
	CompanyOwner *CompanyEngineer `json:"companyOwner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// CompanyOwnerOrErr returns the CompanyOwner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompanyDetailEdges) CompanyOwnerOrErr() (*CompanyEngineer, error) {
	if e.CompanyOwner != nil {
		return e.CompanyOwner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "companyOwner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CompanyDetail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case companydetail.FieldID, companydetail.FieldCompanyCode, companydetail.FieldTaxNo, companydetail.FieldDeleted:
			values[i] = new(sql.NullInt64)
		case companydetail.FieldName, companydetail.FieldAddress, companydetail.FieldCity, companydetail.FieldState, companydetail.FieldPhone, companydetail.FieldFax, companydetail.FieldMobile, companydetail.FieldEmail, companydetail.FieldWebsite, companydetail.FieldTaxAdmin, companydetail.FieldCommerce, companydetail.FieldCommerceReg:
			values[i] = new(sql.NullString)
		case companydetail.FieldVisaDate, companydetail.FieldCreatedAt, companydetail.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case companydetail.ForeignKeys[0]: // owner_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CompanyDetail fields.
func (cd *CompanyDetail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case companydetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cd.ID = int(value.Int64)
		case companydetail.FieldCompanyCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CompanyCode", values[i])
			} else if value.Valid {
				cd.CompanyCode = int(value.Int64)
			}
		case companydetail.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				cd.Name = value.String
			}
		case companydetail.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Address", values[i])
			} else if value.Valid {
				cd.Address = value.String
			}
		case companydetail.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field City", values[i])
			} else if value.Valid {
				cd.City = value.String
			}
		case companydetail.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field State", values[i])
			} else if value.Valid {
				cd.State = value.String
			}
		case companydetail.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Phone", values[i])
			} else if value.Valid {
				cd.Phone = value.String
			}
		case companydetail.FieldFax:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Fax", values[i])
			} else if value.Valid {
				cd.Fax = value.String
			}
		case companydetail.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Mobile", values[i])
			} else if value.Valid {
				cd.Mobile = value.String
			}
		case companydetail.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Email", values[i])
			} else if value.Valid {
				cd.Email = value.String
			}
		case companydetail.FieldWebsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Website", values[i])
			} else if value.Valid {
				cd.Website = value.String
			}
		case companydetail.FieldTaxAdmin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TaxAdmin", values[i])
			} else if value.Valid {
				cd.TaxAdmin = value.String
			}
		case companydetail.FieldTaxNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field TaxNo", values[i])
			} else if value.Valid {
				cd.TaxNo = int(value.Int64)
			}
		case companydetail.FieldCommerce:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Commerce", values[i])
			} else if value.Valid {
				cd.Commerce = value.String
			}
		case companydetail.FieldCommerceReg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CommerceReg", values[i])
			} else if value.Valid {
				cd.CommerceReg = value.String
			}
		case companydetail.FieldVisaDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field VisaDate", values[i])
			} else if value.Valid {
				cd.VisaDate = value.Time
			}
		case companydetail.FieldDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Deleted", values[i])
			} else if value.Valid {
				cd.Deleted = int(value.Int64)
			}
		case companydetail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				cd.CreatedAt = value.Time
			}
		case companydetail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				cd.UpdatedAt = value.Time
			}
		case companydetail.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field owner_id", value)
			} else if value.Valid {
				cd.owner_id = new(int)
				*cd.owner_id = int(value.Int64)
			}
		default:
			cd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CompanyDetail.
// This includes values selected through modifiers, order, etc.
func (cd *CompanyDetail) Value(name string) (ent.Value, error) {
	return cd.selectValues.Get(name)
}

// QueryCompanyOwner queries the "companyOwner" edge of the CompanyDetail entity.
func (cd *CompanyDetail) QueryCompanyOwner() *CompanyEngineerQuery {
	return NewCompanyDetailClient(cd.config).QueryCompanyOwner(cd)
}

// Update returns a builder for updating this CompanyDetail.
// Note that you need to call CompanyDetail.Unwrap() before calling this method if this CompanyDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (cd *CompanyDetail) Update() *CompanyDetailUpdateOne {
	return NewCompanyDetailClient(cd.config).UpdateOne(cd)
}

// Unwrap unwraps the CompanyDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cd *CompanyDetail) Unwrap() *CompanyDetail {
	_tx, ok := cd.config.driver.(*txDriver)
	if !ok {
		panic("ent: CompanyDetail is not a transactional entity")
	}
	cd.config.driver = _tx.drv
	return cd
}

// String implements the fmt.Stringer.
func (cd *CompanyDetail) String() string {
	var builder strings.Builder
	builder.WriteString("CompanyDetail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cd.ID))
	builder.WriteString("CompanyCode=")
	builder.WriteString(fmt.Sprintf("%v", cd.CompanyCode))
	builder.WriteString(", ")
	builder.WriteString("Name=")
	builder.WriteString(cd.Name)
	builder.WriteString(", ")
	builder.WriteString("Address=")
	builder.WriteString(cd.Address)
	builder.WriteString(", ")
	builder.WriteString("City=")
	builder.WriteString(cd.City)
	builder.WriteString(", ")
	builder.WriteString("State=")
	builder.WriteString(cd.State)
	builder.WriteString(", ")
	builder.WriteString("Phone=")
	builder.WriteString(cd.Phone)
	builder.WriteString(", ")
	builder.WriteString("Fax=")
	builder.WriteString(cd.Fax)
	builder.WriteString(", ")
	builder.WriteString("Mobile=")
	builder.WriteString(cd.Mobile)
	builder.WriteString(", ")
	builder.WriteString("Email=")
	builder.WriteString(cd.Email)
	builder.WriteString(", ")
	builder.WriteString("Website=")
	builder.WriteString(cd.Website)
	builder.WriteString(", ")
	builder.WriteString("TaxAdmin=")
	builder.WriteString(cd.TaxAdmin)
	builder.WriteString(", ")
	builder.WriteString("TaxNo=")
	builder.WriteString(fmt.Sprintf("%v", cd.TaxNo))
	builder.WriteString(", ")
	builder.WriteString("Commerce=")
	builder.WriteString(cd.Commerce)
	builder.WriteString(", ")
	builder.WriteString("CommerceReg=")
	builder.WriteString(cd.CommerceReg)
	builder.WriteString(", ")
	builder.WriteString("VisaDate=")
	builder.WriteString(cd.VisaDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("Deleted=")
	builder.WriteString(fmt.Sprintf("%v", cd.Deleted))
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(cd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(cd.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CompanyDetails is a parsable slice of CompanyDetail.
type CompanyDetails []*CompanyDetail
