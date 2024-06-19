// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gqlgen-ent/ent/companydetail"
	"gqlgen-ent/ent/companyowner"
	"gqlgen-ent/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompanyDetailUpdate is the builder for updating CompanyDetail entities.
type CompanyDetailUpdate struct {
	config
	hooks    []Hook
	mutation *CompanyDetailMutation
}

// Where appends a list predicates to the CompanyDetailUpdate builder.
func (cdu *CompanyDetailUpdate) Where(ps ...predicate.CompanyDetail) *CompanyDetailUpdate {
	cdu.mutation.Where(ps...)
	return cdu
}

// SetCompanyCode sets the "CompanyCode" field.
func (cdu *CompanyDetailUpdate) SetCompanyCode(i int) *CompanyDetailUpdate {
	cdu.mutation.ResetCompanyCode()
	cdu.mutation.SetCompanyCode(i)
	return cdu
}

// SetNillableCompanyCode sets the "CompanyCode" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableCompanyCode(i *int) *CompanyDetailUpdate {
	if i != nil {
		cdu.SetCompanyCode(*i)
	}
	return cdu
}

// AddCompanyCode adds i to the "CompanyCode" field.
func (cdu *CompanyDetailUpdate) AddCompanyCode(i int) *CompanyDetailUpdate {
	cdu.mutation.AddCompanyCode(i)
	return cdu
}

// SetName sets the "Name" field.
func (cdu *CompanyDetailUpdate) SetName(s string) *CompanyDetailUpdate {
	cdu.mutation.SetName(s)
	return cdu
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableName(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetName(*s)
	}
	return cdu
}

// SetAddress sets the "Address" field.
func (cdu *CompanyDetailUpdate) SetAddress(s string) *CompanyDetailUpdate {
	cdu.mutation.SetAddress(s)
	return cdu
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableAddress(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetAddress(*s)
	}
	return cdu
}

// ClearAddress clears the value of the "Address" field.
func (cdu *CompanyDetailUpdate) ClearAddress() *CompanyDetailUpdate {
	cdu.mutation.ClearAddress()
	return cdu
}

// SetCity sets the "City" field.
func (cdu *CompanyDetailUpdate) SetCity(s string) *CompanyDetailUpdate {
	cdu.mutation.SetCity(s)
	return cdu
}

// SetNillableCity sets the "City" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableCity(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetCity(*s)
	}
	return cdu
}

// ClearCity clears the value of the "City" field.
func (cdu *CompanyDetailUpdate) ClearCity() *CompanyDetailUpdate {
	cdu.mutation.ClearCity()
	return cdu
}

// SetState sets the "State" field.
func (cdu *CompanyDetailUpdate) SetState(s string) *CompanyDetailUpdate {
	cdu.mutation.SetState(s)
	return cdu
}

// SetNillableState sets the "State" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableState(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetState(*s)
	}
	return cdu
}

// ClearState clears the value of the "State" field.
func (cdu *CompanyDetailUpdate) ClearState() *CompanyDetailUpdate {
	cdu.mutation.ClearState()
	return cdu
}

// SetPhone sets the "Phone" field.
func (cdu *CompanyDetailUpdate) SetPhone(s string) *CompanyDetailUpdate {
	cdu.mutation.SetPhone(s)
	return cdu
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillablePhone(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetPhone(*s)
	}
	return cdu
}

// ClearPhone clears the value of the "Phone" field.
func (cdu *CompanyDetailUpdate) ClearPhone() *CompanyDetailUpdate {
	cdu.mutation.ClearPhone()
	return cdu
}

// SetFax sets the "Fax" field.
func (cdu *CompanyDetailUpdate) SetFax(s string) *CompanyDetailUpdate {
	cdu.mutation.SetFax(s)
	return cdu
}

// SetNillableFax sets the "Fax" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableFax(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetFax(*s)
	}
	return cdu
}

// ClearFax clears the value of the "Fax" field.
func (cdu *CompanyDetailUpdate) ClearFax() *CompanyDetailUpdate {
	cdu.mutation.ClearFax()
	return cdu
}

// SetMobile sets the "Mobile" field.
func (cdu *CompanyDetailUpdate) SetMobile(s string) *CompanyDetailUpdate {
	cdu.mutation.SetMobile(s)
	return cdu
}

// SetNillableMobile sets the "Mobile" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableMobile(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetMobile(*s)
	}
	return cdu
}

// ClearMobile clears the value of the "Mobile" field.
func (cdu *CompanyDetailUpdate) ClearMobile() *CompanyDetailUpdate {
	cdu.mutation.ClearMobile()
	return cdu
}

// SetEmail sets the "Email" field.
func (cdu *CompanyDetailUpdate) SetEmail(s string) *CompanyDetailUpdate {
	cdu.mutation.SetEmail(s)
	return cdu
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableEmail(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetEmail(*s)
	}
	return cdu
}

// ClearEmail clears the value of the "Email" field.
func (cdu *CompanyDetailUpdate) ClearEmail() *CompanyDetailUpdate {
	cdu.mutation.ClearEmail()
	return cdu
}

// SetWebsite sets the "Website" field.
func (cdu *CompanyDetailUpdate) SetWebsite(s string) *CompanyDetailUpdate {
	cdu.mutation.SetWebsite(s)
	return cdu
}

// SetNillableWebsite sets the "Website" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableWebsite(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetWebsite(*s)
	}
	return cdu
}

// ClearWebsite clears the value of the "Website" field.
func (cdu *CompanyDetailUpdate) ClearWebsite() *CompanyDetailUpdate {
	cdu.mutation.ClearWebsite()
	return cdu
}

// SetTaxAdmin sets the "TaxAdmin" field.
func (cdu *CompanyDetailUpdate) SetTaxAdmin(s string) *CompanyDetailUpdate {
	cdu.mutation.SetTaxAdmin(s)
	return cdu
}

// SetNillableTaxAdmin sets the "TaxAdmin" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableTaxAdmin(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetTaxAdmin(*s)
	}
	return cdu
}

// ClearTaxAdmin clears the value of the "TaxAdmin" field.
func (cdu *CompanyDetailUpdate) ClearTaxAdmin() *CompanyDetailUpdate {
	cdu.mutation.ClearTaxAdmin()
	return cdu
}

// SetTaxNo sets the "TaxNo" field.
func (cdu *CompanyDetailUpdate) SetTaxNo(i int) *CompanyDetailUpdate {
	cdu.mutation.ResetTaxNo()
	cdu.mutation.SetTaxNo(i)
	return cdu
}

// SetNillableTaxNo sets the "TaxNo" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableTaxNo(i *int) *CompanyDetailUpdate {
	if i != nil {
		cdu.SetTaxNo(*i)
	}
	return cdu
}

// AddTaxNo adds i to the "TaxNo" field.
func (cdu *CompanyDetailUpdate) AddTaxNo(i int) *CompanyDetailUpdate {
	cdu.mutation.AddTaxNo(i)
	return cdu
}

// ClearTaxNo clears the value of the "TaxNo" field.
func (cdu *CompanyDetailUpdate) ClearTaxNo() *CompanyDetailUpdate {
	cdu.mutation.ClearTaxNo()
	return cdu
}

// SetCommerce sets the "Commerce" field.
func (cdu *CompanyDetailUpdate) SetCommerce(s string) *CompanyDetailUpdate {
	cdu.mutation.SetCommerce(s)
	return cdu
}

// SetNillableCommerce sets the "Commerce" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableCommerce(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetCommerce(*s)
	}
	return cdu
}

// ClearCommerce clears the value of the "Commerce" field.
func (cdu *CompanyDetailUpdate) ClearCommerce() *CompanyDetailUpdate {
	cdu.mutation.ClearCommerce()
	return cdu
}

// SetCommerceReg sets the "CommerceReg" field.
func (cdu *CompanyDetailUpdate) SetCommerceReg(s string) *CompanyDetailUpdate {
	cdu.mutation.SetCommerceReg(s)
	return cdu
}

// SetNillableCommerceReg sets the "CommerceReg" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableCommerceReg(s *string) *CompanyDetailUpdate {
	if s != nil {
		cdu.SetCommerceReg(*s)
	}
	return cdu
}

// ClearCommerceReg clears the value of the "CommerceReg" field.
func (cdu *CompanyDetailUpdate) ClearCommerceReg() *CompanyDetailUpdate {
	cdu.mutation.ClearCommerceReg()
	return cdu
}

// SetVisaDate sets the "VisaDate" field.
func (cdu *CompanyDetailUpdate) SetVisaDate(t time.Time) *CompanyDetailUpdate {
	cdu.mutation.SetVisaDate(t)
	return cdu
}

// SetNillableVisaDate sets the "VisaDate" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableVisaDate(t *time.Time) *CompanyDetailUpdate {
	if t != nil {
		cdu.SetVisaDate(*t)
	}
	return cdu
}

// ClearVisaDate clears the value of the "VisaDate" field.
func (cdu *CompanyDetailUpdate) ClearVisaDate() *CompanyDetailUpdate {
	cdu.mutation.ClearVisaDate()
	return cdu
}

// SetDeleted sets the "Deleted" field.
func (cdu *CompanyDetailUpdate) SetDeleted(i int) *CompanyDetailUpdate {
	cdu.mutation.ResetDeleted()
	cdu.mutation.SetDeleted(i)
	return cdu
}

// SetNillableDeleted sets the "Deleted" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableDeleted(i *int) *CompanyDetailUpdate {
	if i != nil {
		cdu.SetDeleted(*i)
	}
	return cdu
}

// AddDeleted adds i to the "Deleted" field.
func (cdu *CompanyDetailUpdate) AddDeleted(i int) *CompanyDetailUpdate {
	cdu.mutation.AddDeleted(i)
	return cdu
}

// SetCreatedAt sets the "CreatedAt" field.
func (cdu *CompanyDetailUpdate) SetCreatedAt(t time.Time) *CompanyDetailUpdate {
	cdu.mutation.SetCreatedAt(t)
	return cdu
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableCreatedAt(t *time.Time) *CompanyDetailUpdate {
	if t != nil {
		cdu.SetCreatedAt(*t)
	}
	return cdu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (cdu *CompanyDetailUpdate) SetUpdatedAt(t time.Time) *CompanyDetailUpdate {
	cdu.mutation.SetUpdatedAt(t)
	return cdu
}

// SetCompanyOwnerID sets the "companyOwner" edge to the CompanyOwner entity by ID.
func (cdu *CompanyDetailUpdate) SetCompanyOwnerID(id int) *CompanyDetailUpdate {
	cdu.mutation.SetCompanyOwnerID(id)
	return cdu
}

// SetNillableCompanyOwnerID sets the "companyOwner" edge to the CompanyOwner entity by ID if the given value is not nil.
func (cdu *CompanyDetailUpdate) SetNillableCompanyOwnerID(id *int) *CompanyDetailUpdate {
	if id != nil {
		cdu = cdu.SetCompanyOwnerID(*id)
	}
	return cdu
}

// SetCompanyOwner sets the "companyOwner" edge to the CompanyOwner entity.
func (cdu *CompanyDetailUpdate) SetCompanyOwner(c *CompanyOwner) *CompanyDetailUpdate {
	return cdu.SetCompanyOwnerID(c.ID)
}

// Mutation returns the CompanyDetailMutation object of the builder.
func (cdu *CompanyDetailUpdate) Mutation() *CompanyDetailMutation {
	return cdu.mutation
}

// ClearCompanyOwner clears the "companyOwner" edge to the CompanyOwner entity.
func (cdu *CompanyDetailUpdate) ClearCompanyOwner() *CompanyDetailUpdate {
	cdu.mutation.ClearCompanyOwner()
	return cdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cdu *CompanyDetailUpdate) Save(ctx context.Context) (int, error) {
	cdu.defaults()
	return withHooks(ctx, cdu.sqlSave, cdu.mutation, cdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cdu *CompanyDetailUpdate) SaveX(ctx context.Context) int {
	affected, err := cdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cdu *CompanyDetailUpdate) Exec(ctx context.Context) error {
	_, err := cdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdu *CompanyDetailUpdate) ExecX(ctx context.Context) {
	if err := cdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cdu *CompanyDetailUpdate) defaults() {
	if _, ok := cdu.mutation.UpdatedAt(); !ok {
		v := companydetail.UpdateDefaultUpdatedAt()
		cdu.mutation.SetUpdatedAt(v)
	}
}

func (cdu *CompanyDetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(companydetail.Table, companydetail.Columns, sqlgraph.NewFieldSpec(companydetail.FieldID, field.TypeInt))
	if ps := cdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cdu.mutation.CompanyCode(); ok {
		_spec.SetField(companydetail.FieldCompanyCode, field.TypeInt, value)
	}
	if value, ok := cdu.mutation.AddedCompanyCode(); ok {
		_spec.AddField(companydetail.FieldCompanyCode, field.TypeInt, value)
	}
	if value, ok := cdu.mutation.Name(); ok {
		_spec.SetField(companydetail.FieldName, field.TypeString, value)
	}
	if value, ok := cdu.mutation.Address(); ok {
		_spec.SetField(companydetail.FieldAddress, field.TypeString, value)
	}
	if cdu.mutation.AddressCleared() {
		_spec.ClearField(companydetail.FieldAddress, field.TypeString)
	}
	if value, ok := cdu.mutation.City(); ok {
		_spec.SetField(companydetail.FieldCity, field.TypeString, value)
	}
	if cdu.mutation.CityCleared() {
		_spec.ClearField(companydetail.FieldCity, field.TypeString)
	}
	if value, ok := cdu.mutation.State(); ok {
		_spec.SetField(companydetail.FieldState, field.TypeString, value)
	}
	if cdu.mutation.StateCleared() {
		_spec.ClearField(companydetail.FieldState, field.TypeString)
	}
	if value, ok := cdu.mutation.Phone(); ok {
		_spec.SetField(companydetail.FieldPhone, field.TypeString, value)
	}
	if cdu.mutation.PhoneCleared() {
		_spec.ClearField(companydetail.FieldPhone, field.TypeString)
	}
	if value, ok := cdu.mutation.Fax(); ok {
		_spec.SetField(companydetail.FieldFax, field.TypeString, value)
	}
	if cdu.mutation.FaxCleared() {
		_spec.ClearField(companydetail.FieldFax, field.TypeString)
	}
	if value, ok := cdu.mutation.Mobile(); ok {
		_spec.SetField(companydetail.FieldMobile, field.TypeString, value)
	}
	if cdu.mutation.MobileCleared() {
		_spec.ClearField(companydetail.FieldMobile, field.TypeString)
	}
	if value, ok := cdu.mutation.Email(); ok {
		_spec.SetField(companydetail.FieldEmail, field.TypeString, value)
	}
	if cdu.mutation.EmailCleared() {
		_spec.ClearField(companydetail.FieldEmail, field.TypeString)
	}
	if value, ok := cdu.mutation.Website(); ok {
		_spec.SetField(companydetail.FieldWebsite, field.TypeString, value)
	}
	if cdu.mutation.WebsiteCleared() {
		_spec.ClearField(companydetail.FieldWebsite, field.TypeString)
	}
	if value, ok := cdu.mutation.TaxAdmin(); ok {
		_spec.SetField(companydetail.FieldTaxAdmin, field.TypeString, value)
	}
	if cdu.mutation.TaxAdminCleared() {
		_spec.ClearField(companydetail.FieldTaxAdmin, field.TypeString)
	}
	if value, ok := cdu.mutation.TaxNo(); ok {
		_spec.SetField(companydetail.FieldTaxNo, field.TypeInt, value)
	}
	if value, ok := cdu.mutation.AddedTaxNo(); ok {
		_spec.AddField(companydetail.FieldTaxNo, field.TypeInt, value)
	}
	if cdu.mutation.TaxNoCleared() {
		_spec.ClearField(companydetail.FieldTaxNo, field.TypeInt)
	}
	if value, ok := cdu.mutation.Commerce(); ok {
		_spec.SetField(companydetail.FieldCommerce, field.TypeString, value)
	}
	if cdu.mutation.CommerceCleared() {
		_spec.ClearField(companydetail.FieldCommerce, field.TypeString)
	}
	if value, ok := cdu.mutation.CommerceReg(); ok {
		_spec.SetField(companydetail.FieldCommerceReg, field.TypeString, value)
	}
	if cdu.mutation.CommerceRegCleared() {
		_spec.ClearField(companydetail.FieldCommerceReg, field.TypeString)
	}
	if value, ok := cdu.mutation.VisaDate(); ok {
		_spec.SetField(companydetail.FieldVisaDate, field.TypeTime, value)
	}
	if cdu.mutation.VisaDateCleared() {
		_spec.ClearField(companydetail.FieldVisaDate, field.TypeTime)
	}
	if value, ok := cdu.mutation.Deleted(); ok {
		_spec.SetField(companydetail.FieldDeleted, field.TypeInt, value)
	}
	if value, ok := cdu.mutation.AddedDeleted(); ok {
		_spec.AddField(companydetail.FieldDeleted, field.TypeInt, value)
	}
	if value, ok := cdu.mutation.CreatedAt(); ok {
		_spec.SetField(companydetail.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cdu.mutation.UpdatedAt(); ok {
		_spec.SetField(companydetail.FieldUpdatedAt, field.TypeTime, value)
	}
	if cdu.mutation.CompanyOwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companydetail.CompanyOwnerTable,
			Columns: []string{companydetail.CompanyOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyowner.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cdu.mutation.CompanyOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companydetail.CompanyOwnerTable,
			Columns: []string{companydetail.CompanyOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyowner.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{companydetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cdu.mutation.done = true
	return n, nil
}

// CompanyDetailUpdateOne is the builder for updating a single CompanyDetail entity.
type CompanyDetailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CompanyDetailMutation
}

// SetCompanyCode sets the "CompanyCode" field.
func (cduo *CompanyDetailUpdateOne) SetCompanyCode(i int) *CompanyDetailUpdateOne {
	cduo.mutation.ResetCompanyCode()
	cduo.mutation.SetCompanyCode(i)
	return cduo
}

// SetNillableCompanyCode sets the "CompanyCode" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableCompanyCode(i *int) *CompanyDetailUpdateOne {
	if i != nil {
		cduo.SetCompanyCode(*i)
	}
	return cduo
}

// AddCompanyCode adds i to the "CompanyCode" field.
func (cduo *CompanyDetailUpdateOne) AddCompanyCode(i int) *CompanyDetailUpdateOne {
	cduo.mutation.AddCompanyCode(i)
	return cduo
}

// SetName sets the "Name" field.
func (cduo *CompanyDetailUpdateOne) SetName(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetName(s)
	return cduo
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableName(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetName(*s)
	}
	return cduo
}

// SetAddress sets the "Address" field.
func (cduo *CompanyDetailUpdateOne) SetAddress(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetAddress(s)
	return cduo
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableAddress(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetAddress(*s)
	}
	return cduo
}

// ClearAddress clears the value of the "Address" field.
func (cduo *CompanyDetailUpdateOne) ClearAddress() *CompanyDetailUpdateOne {
	cduo.mutation.ClearAddress()
	return cduo
}

// SetCity sets the "City" field.
func (cduo *CompanyDetailUpdateOne) SetCity(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetCity(s)
	return cduo
}

// SetNillableCity sets the "City" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableCity(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetCity(*s)
	}
	return cduo
}

// ClearCity clears the value of the "City" field.
func (cduo *CompanyDetailUpdateOne) ClearCity() *CompanyDetailUpdateOne {
	cduo.mutation.ClearCity()
	return cduo
}

// SetState sets the "State" field.
func (cduo *CompanyDetailUpdateOne) SetState(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetState(s)
	return cduo
}

// SetNillableState sets the "State" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableState(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetState(*s)
	}
	return cduo
}

// ClearState clears the value of the "State" field.
func (cduo *CompanyDetailUpdateOne) ClearState() *CompanyDetailUpdateOne {
	cduo.mutation.ClearState()
	return cduo
}

// SetPhone sets the "Phone" field.
func (cduo *CompanyDetailUpdateOne) SetPhone(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetPhone(s)
	return cduo
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillablePhone(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetPhone(*s)
	}
	return cduo
}

// ClearPhone clears the value of the "Phone" field.
func (cduo *CompanyDetailUpdateOne) ClearPhone() *CompanyDetailUpdateOne {
	cduo.mutation.ClearPhone()
	return cduo
}

// SetFax sets the "Fax" field.
func (cduo *CompanyDetailUpdateOne) SetFax(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetFax(s)
	return cduo
}

// SetNillableFax sets the "Fax" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableFax(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetFax(*s)
	}
	return cduo
}

// ClearFax clears the value of the "Fax" field.
func (cduo *CompanyDetailUpdateOne) ClearFax() *CompanyDetailUpdateOne {
	cduo.mutation.ClearFax()
	return cduo
}

// SetMobile sets the "Mobile" field.
func (cduo *CompanyDetailUpdateOne) SetMobile(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetMobile(s)
	return cduo
}

// SetNillableMobile sets the "Mobile" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableMobile(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetMobile(*s)
	}
	return cduo
}

// ClearMobile clears the value of the "Mobile" field.
func (cduo *CompanyDetailUpdateOne) ClearMobile() *CompanyDetailUpdateOne {
	cduo.mutation.ClearMobile()
	return cduo
}

// SetEmail sets the "Email" field.
func (cduo *CompanyDetailUpdateOne) SetEmail(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetEmail(s)
	return cduo
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableEmail(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetEmail(*s)
	}
	return cduo
}

// ClearEmail clears the value of the "Email" field.
func (cduo *CompanyDetailUpdateOne) ClearEmail() *CompanyDetailUpdateOne {
	cduo.mutation.ClearEmail()
	return cduo
}

// SetWebsite sets the "Website" field.
func (cduo *CompanyDetailUpdateOne) SetWebsite(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetWebsite(s)
	return cduo
}

// SetNillableWebsite sets the "Website" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableWebsite(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetWebsite(*s)
	}
	return cduo
}

// ClearWebsite clears the value of the "Website" field.
func (cduo *CompanyDetailUpdateOne) ClearWebsite() *CompanyDetailUpdateOne {
	cduo.mutation.ClearWebsite()
	return cduo
}

// SetTaxAdmin sets the "TaxAdmin" field.
func (cduo *CompanyDetailUpdateOne) SetTaxAdmin(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetTaxAdmin(s)
	return cduo
}

// SetNillableTaxAdmin sets the "TaxAdmin" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableTaxAdmin(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetTaxAdmin(*s)
	}
	return cduo
}

// ClearTaxAdmin clears the value of the "TaxAdmin" field.
func (cduo *CompanyDetailUpdateOne) ClearTaxAdmin() *CompanyDetailUpdateOne {
	cduo.mutation.ClearTaxAdmin()
	return cduo
}

// SetTaxNo sets the "TaxNo" field.
func (cduo *CompanyDetailUpdateOne) SetTaxNo(i int) *CompanyDetailUpdateOne {
	cduo.mutation.ResetTaxNo()
	cduo.mutation.SetTaxNo(i)
	return cduo
}

// SetNillableTaxNo sets the "TaxNo" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableTaxNo(i *int) *CompanyDetailUpdateOne {
	if i != nil {
		cduo.SetTaxNo(*i)
	}
	return cduo
}

// AddTaxNo adds i to the "TaxNo" field.
func (cduo *CompanyDetailUpdateOne) AddTaxNo(i int) *CompanyDetailUpdateOne {
	cduo.mutation.AddTaxNo(i)
	return cduo
}

// ClearTaxNo clears the value of the "TaxNo" field.
func (cduo *CompanyDetailUpdateOne) ClearTaxNo() *CompanyDetailUpdateOne {
	cduo.mutation.ClearTaxNo()
	return cduo
}

// SetCommerce sets the "Commerce" field.
func (cduo *CompanyDetailUpdateOne) SetCommerce(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetCommerce(s)
	return cduo
}

// SetNillableCommerce sets the "Commerce" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableCommerce(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetCommerce(*s)
	}
	return cduo
}

// ClearCommerce clears the value of the "Commerce" field.
func (cduo *CompanyDetailUpdateOne) ClearCommerce() *CompanyDetailUpdateOne {
	cduo.mutation.ClearCommerce()
	return cduo
}

// SetCommerceReg sets the "CommerceReg" field.
func (cduo *CompanyDetailUpdateOne) SetCommerceReg(s string) *CompanyDetailUpdateOne {
	cduo.mutation.SetCommerceReg(s)
	return cduo
}

// SetNillableCommerceReg sets the "CommerceReg" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableCommerceReg(s *string) *CompanyDetailUpdateOne {
	if s != nil {
		cduo.SetCommerceReg(*s)
	}
	return cduo
}

// ClearCommerceReg clears the value of the "CommerceReg" field.
func (cduo *CompanyDetailUpdateOne) ClearCommerceReg() *CompanyDetailUpdateOne {
	cduo.mutation.ClearCommerceReg()
	return cduo
}

// SetVisaDate sets the "VisaDate" field.
func (cduo *CompanyDetailUpdateOne) SetVisaDate(t time.Time) *CompanyDetailUpdateOne {
	cduo.mutation.SetVisaDate(t)
	return cduo
}

// SetNillableVisaDate sets the "VisaDate" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableVisaDate(t *time.Time) *CompanyDetailUpdateOne {
	if t != nil {
		cduo.SetVisaDate(*t)
	}
	return cduo
}

// ClearVisaDate clears the value of the "VisaDate" field.
func (cduo *CompanyDetailUpdateOne) ClearVisaDate() *CompanyDetailUpdateOne {
	cduo.mutation.ClearVisaDate()
	return cduo
}

// SetDeleted sets the "Deleted" field.
func (cduo *CompanyDetailUpdateOne) SetDeleted(i int) *CompanyDetailUpdateOne {
	cduo.mutation.ResetDeleted()
	cduo.mutation.SetDeleted(i)
	return cduo
}

// SetNillableDeleted sets the "Deleted" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableDeleted(i *int) *CompanyDetailUpdateOne {
	if i != nil {
		cduo.SetDeleted(*i)
	}
	return cduo
}

// AddDeleted adds i to the "Deleted" field.
func (cduo *CompanyDetailUpdateOne) AddDeleted(i int) *CompanyDetailUpdateOne {
	cduo.mutation.AddDeleted(i)
	return cduo
}

// SetCreatedAt sets the "CreatedAt" field.
func (cduo *CompanyDetailUpdateOne) SetCreatedAt(t time.Time) *CompanyDetailUpdateOne {
	cduo.mutation.SetCreatedAt(t)
	return cduo
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableCreatedAt(t *time.Time) *CompanyDetailUpdateOne {
	if t != nil {
		cduo.SetCreatedAt(*t)
	}
	return cduo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (cduo *CompanyDetailUpdateOne) SetUpdatedAt(t time.Time) *CompanyDetailUpdateOne {
	cduo.mutation.SetUpdatedAt(t)
	return cduo
}

// SetCompanyOwnerID sets the "companyOwner" edge to the CompanyOwner entity by ID.
func (cduo *CompanyDetailUpdateOne) SetCompanyOwnerID(id int) *CompanyDetailUpdateOne {
	cduo.mutation.SetCompanyOwnerID(id)
	return cduo
}

// SetNillableCompanyOwnerID sets the "companyOwner" edge to the CompanyOwner entity by ID if the given value is not nil.
func (cduo *CompanyDetailUpdateOne) SetNillableCompanyOwnerID(id *int) *CompanyDetailUpdateOne {
	if id != nil {
		cduo = cduo.SetCompanyOwnerID(*id)
	}
	return cduo
}

// SetCompanyOwner sets the "companyOwner" edge to the CompanyOwner entity.
func (cduo *CompanyDetailUpdateOne) SetCompanyOwner(c *CompanyOwner) *CompanyDetailUpdateOne {
	return cduo.SetCompanyOwnerID(c.ID)
}

// Mutation returns the CompanyDetailMutation object of the builder.
func (cduo *CompanyDetailUpdateOne) Mutation() *CompanyDetailMutation {
	return cduo.mutation
}

// ClearCompanyOwner clears the "companyOwner" edge to the CompanyOwner entity.
func (cduo *CompanyDetailUpdateOne) ClearCompanyOwner() *CompanyDetailUpdateOne {
	cduo.mutation.ClearCompanyOwner()
	return cduo
}

// Where appends a list predicates to the CompanyDetailUpdate builder.
func (cduo *CompanyDetailUpdateOne) Where(ps ...predicate.CompanyDetail) *CompanyDetailUpdateOne {
	cduo.mutation.Where(ps...)
	return cduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cduo *CompanyDetailUpdateOne) Select(field string, fields ...string) *CompanyDetailUpdateOne {
	cduo.fields = append([]string{field}, fields...)
	return cduo
}

// Save executes the query and returns the updated CompanyDetail entity.
func (cduo *CompanyDetailUpdateOne) Save(ctx context.Context) (*CompanyDetail, error) {
	cduo.defaults()
	return withHooks(ctx, cduo.sqlSave, cduo.mutation, cduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cduo *CompanyDetailUpdateOne) SaveX(ctx context.Context) *CompanyDetail {
	node, err := cduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cduo *CompanyDetailUpdateOne) Exec(ctx context.Context) error {
	_, err := cduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cduo *CompanyDetailUpdateOne) ExecX(ctx context.Context) {
	if err := cduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cduo *CompanyDetailUpdateOne) defaults() {
	if _, ok := cduo.mutation.UpdatedAt(); !ok {
		v := companydetail.UpdateDefaultUpdatedAt()
		cduo.mutation.SetUpdatedAt(v)
	}
}

func (cduo *CompanyDetailUpdateOne) sqlSave(ctx context.Context) (_node *CompanyDetail, err error) {
	_spec := sqlgraph.NewUpdateSpec(companydetail.Table, companydetail.Columns, sqlgraph.NewFieldSpec(companydetail.FieldID, field.TypeInt))
	id, ok := cduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CompanyDetail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, companydetail.FieldID)
		for _, f := range fields {
			if !companydetail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != companydetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cduo.mutation.CompanyCode(); ok {
		_spec.SetField(companydetail.FieldCompanyCode, field.TypeInt, value)
	}
	if value, ok := cduo.mutation.AddedCompanyCode(); ok {
		_spec.AddField(companydetail.FieldCompanyCode, field.TypeInt, value)
	}
	if value, ok := cduo.mutation.Name(); ok {
		_spec.SetField(companydetail.FieldName, field.TypeString, value)
	}
	if value, ok := cduo.mutation.Address(); ok {
		_spec.SetField(companydetail.FieldAddress, field.TypeString, value)
	}
	if cduo.mutation.AddressCleared() {
		_spec.ClearField(companydetail.FieldAddress, field.TypeString)
	}
	if value, ok := cduo.mutation.City(); ok {
		_spec.SetField(companydetail.FieldCity, field.TypeString, value)
	}
	if cduo.mutation.CityCleared() {
		_spec.ClearField(companydetail.FieldCity, field.TypeString)
	}
	if value, ok := cduo.mutation.State(); ok {
		_spec.SetField(companydetail.FieldState, field.TypeString, value)
	}
	if cduo.mutation.StateCleared() {
		_spec.ClearField(companydetail.FieldState, field.TypeString)
	}
	if value, ok := cduo.mutation.Phone(); ok {
		_spec.SetField(companydetail.FieldPhone, field.TypeString, value)
	}
	if cduo.mutation.PhoneCleared() {
		_spec.ClearField(companydetail.FieldPhone, field.TypeString)
	}
	if value, ok := cduo.mutation.Fax(); ok {
		_spec.SetField(companydetail.FieldFax, field.TypeString, value)
	}
	if cduo.mutation.FaxCleared() {
		_spec.ClearField(companydetail.FieldFax, field.TypeString)
	}
	if value, ok := cduo.mutation.Mobile(); ok {
		_spec.SetField(companydetail.FieldMobile, field.TypeString, value)
	}
	if cduo.mutation.MobileCleared() {
		_spec.ClearField(companydetail.FieldMobile, field.TypeString)
	}
	if value, ok := cduo.mutation.Email(); ok {
		_spec.SetField(companydetail.FieldEmail, field.TypeString, value)
	}
	if cduo.mutation.EmailCleared() {
		_spec.ClearField(companydetail.FieldEmail, field.TypeString)
	}
	if value, ok := cduo.mutation.Website(); ok {
		_spec.SetField(companydetail.FieldWebsite, field.TypeString, value)
	}
	if cduo.mutation.WebsiteCleared() {
		_spec.ClearField(companydetail.FieldWebsite, field.TypeString)
	}
	if value, ok := cduo.mutation.TaxAdmin(); ok {
		_spec.SetField(companydetail.FieldTaxAdmin, field.TypeString, value)
	}
	if cduo.mutation.TaxAdminCleared() {
		_spec.ClearField(companydetail.FieldTaxAdmin, field.TypeString)
	}
	if value, ok := cduo.mutation.TaxNo(); ok {
		_spec.SetField(companydetail.FieldTaxNo, field.TypeInt, value)
	}
	if value, ok := cduo.mutation.AddedTaxNo(); ok {
		_spec.AddField(companydetail.FieldTaxNo, field.TypeInt, value)
	}
	if cduo.mutation.TaxNoCleared() {
		_spec.ClearField(companydetail.FieldTaxNo, field.TypeInt)
	}
	if value, ok := cduo.mutation.Commerce(); ok {
		_spec.SetField(companydetail.FieldCommerce, field.TypeString, value)
	}
	if cduo.mutation.CommerceCleared() {
		_spec.ClearField(companydetail.FieldCommerce, field.TypeString)
	}
	if value, ok := cduo.mutation.CommerceReg(); ok {
		_spec.SetField(companydetail.FieldCommerceReg, field.TypeString, value)
	}
	if cduo.mutation.CommerceRegCleared() {
		_spec.ClearField(companydetail.FieldCommerceReg, field.TypeString)
	}
	if value, ok := cduo.mutation.VisaDate(); ok {
		_spec.SetField(companydetail.FieldVisaDate, field.TypeTime, value)
	}
	if cduo.mutation.VisaDateCleared() {
		_spec.ClearField(companydetail.FieldVisaDate, field.TypeTime)
	}
	if value, ok := cduo.mutation.Deleted(); ok {
		_spec.SetField(companydetail.FieldDeleted, field.TypeInt, value)
	}
	if value, ok := cduo.mutation.AddedDeleted(); ok {
		_spec.AddField(companydetail.FieldDeleted, field.TypeInt, value)
	}
	if value, ok := cduo.mutation.CreatedAt(); ok {
		_spec.SetField(companydetail.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cduo.mutation.UpdatedAt(); ok {
		_spec.SetField(companydetail.FieldUpdatedAt, field.TypeTime, value)
	}
	if cduo.mutation.CompanyOwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companydetail.CompanyOwnerTable,
			Columns: []string{companydetail.CompanyOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyowner.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cduo.mutation.CompanyOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companydetail.CompanyOwnerTable,
			Columns: []string{companydetail.CompanyOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyowner.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CompanyDetail{config: cduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{companydetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cduo.mutation.done = true
	return _node, nil
}