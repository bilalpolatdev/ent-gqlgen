// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gqlgen-ent/ent/companycareer"
	"gqlgen-ent/ent/companydetail"
	"gqlgen-ent/ent/companyengineer"
	"gqlgen-ent/ent/companyposition"
	"gqlgen-ent/ent/jobdetail"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompanyEngineerCreate is the builder for creating a CompanyEngineer entity.
type CompanyEngineerCreate struct {
	config
	mutation *CompanyEngineerMutation
	hooks    []Hook
}

// SetName sets the "Name" field.
func (cec *CompanyEngineerCreate) SetName(s string) *CompanyEngineerCreate {
	cec.mutation.SetName(s)
	return cec
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableName(s *string) *CompanyEngineerCreate {
	if s != nil {
		cec.SetName(*s)
	}
	return cec
}

// SetAddress sets the "Address" field.
func (cec *CompanyEngineerCreate) SetAddress(s string) *CompanyEngineerCreate {
	cec.mutation.SetAddress(s)
	return cec
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableAddress(s *string) *CompanyEngineerCreate {
	if s != nil {
		cec.SetAddress(*s)
	}
	return cec
}

// SetEmail sets the "Email" field.
func (cec *CompanyEngineerCreate) SetEmail(s string) *CompanyEngineerCreate {
	cec.mutation.SetEmail(s)
	return cec
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableEmail(s *string) *CompanyEngineerCreate {
	if s != nil {
		cec.SetEmail(*s)
	}
	return cec
}

// SetTcNo sets the "TcNo" field.
func (cec *CompanyEngineerCreate) SetTcNo(i int) *CompanyEngineerCreate {
	cec.mutation.SetTcNo(i)
	return cec
}

// SetNillableTcNo sets the "TcNo" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableTcNo(i *int) *CompanyEngineerCreate {
	if i != nil {
		cec.SetTcNo(*i)
	}
	return cec
}

// SetPhone sets the "Phone" field.
func (cec *CompanyEngineerCreate) SetPhone(s string) *CompanyEngineerCreate {
	cec.mutation.SetPhone(s)
	return cec
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillablePhone(s *string) *CompanyEngineerCreate {
	if s != nil {
		cec.SetPhone(*s)
	}
	return cec
}

// SetRegNo sets the "RegNo" field.
func (cec *CompanyEngineerCreate) SetRegNo(i int) *CompanyEngineerCreate {
	cec.mutation.SetRegNo(i)
	return cec
}

// SetNillableRegNo sets the "RegNo" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableRegNo(i *int) *CompanyEngineerCreate {
	if i != nil {
		cec.SetRegNo(*i)
	}
	return cec
}

// SetCertNo sets the "CertNo" field.
func (cec *CompanyEngineerCreate) SetCertNo(i int) *CompanyEngineerCreate {
	cec.mutation.SetCertNo(i)
	return cec
}

// SetNillableCertNo sets the "CertNo" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableCertNo(i *int) *CompanyEngineerCreate {
	if i != nil {
		cec.SetCertNo(*i)
	}
	return cec
}

// SetNote sets the "Note" field.
func (cec *CompanyEngineerCreate) SetNote(s string) *CompanyEngineerCreate {
	cec.mutation.SetNote(s)
	return cec
}

// SetNillableNote sets the "Note" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableNote(s *string) *CompanyEngineerCreate {
	if s != nil {
		cec.SetNote(*s)
	}
	return cec
}

// SetStatus sets the "Status" field.
func (cec *CompanyEngineerCreate) SetStatus(i int) *CompanyEngineerCreate {
	cec.mutation.SetStatus(i)
	return cec
}

// SetNillableStatus sets the "Status" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableStatus(i *int) *CompanyEngineerCreate {
	if i != nil {
		cec.SetStatus(*i)
	}
	return cec
}

// SetDeleted sets the "Deleted" field.
func (cec *CompanyEngineerCreate) SetDeleted(i int) *CompanyEngineerCreate {
	cec.mutation.SetDeleted(i)
	return cec
}

// SetNillableDeleted sets the "Deleted" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableDeleted(i *int) *CompanyEngineerCreate {
	if i != nil {
		cec.SetDeleted(*i)
	}
	return cec
}

// SetEmployment sets the "Employment" field.
func (cec *CompanyEngineerCreate) SetEmployment(t time.Time) *CompanyEngineerCreate {
	cec.mutation.SetEmployment(t)
	return cec
}

// SetNillableEmployment sets the "Employment" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableEmployment(t *time.Time) *CompanyEngineerCreate {
	if t != nil {
		cec.SetEmployment(*t)
	}
	return cec
}

// SetDismissal sets the "Dismissal" field.
func (cec *CompanyEngineerCreate) SetDismissal(t time.Time) *CompanyEngineerCreate {
	cec.mutation.SetDismissal(t)
	return cec
}

// SetNillableDismissal sets the "Dismissal" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableDismissal(t *time.Time) *CompanyEngineerCreate {
	if t != nil {
		cec.SetDismissal(*t)
	}
	return cec
}

// SetCreatedAt sets the "CreatedAt" field.
func (cec *CompanyEngineerCreate) SetCreatedAt(t time.Time) *CompanyEngineerCreate {
	cec.mutation.SetCreatedAt(t)
	return cec
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableCreatedAt(t *time.Time) *CompanyEngineerCreate {
	if t != nil {
		cec.SetCreatedAt(*t)
	}
	return cec
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (cec *CompanyEngineerCreate) SetUpdatedAt(t time.Time) *CompanyEngineerCreate {
	cec.mutation.SetUpdatedAt(t)
	return cec
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableUpdatedAt(t *time.Time) *CompanyEngineerCreate {
	if t != nil {
		cec.SetUpdatedAt(*t)
	}
	return cec
}

// SetEngineerCareerID sets the "engineerCareer" edge to the CompanyCareer entity by ID.
func (cec *CompanyEngineerCreate) SetEngineerCareerID(id int) *CompanyEngineerCreate {
	cec.mutation.SetEngineerCareerID(id)
	return cec
}

// SetNillableEngineerCareerID sets the "engineerCareer" edge to the CompanyCareer entity by ID if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableEngineerCareerID(id *int) *CompanyEngineerCreate {
	if id != nil {
		cec = cec.SetEngineerCareerID(*id)
	}
	return cec
}

// SetEngineerCareer sets the "engineerCareer" edge to the CompanyCareer entity.
func (cec *CompanyEngineerCreate) SetEngineerCareer(c *CompanyCareer) *CompanyEngineerCreate {
	return cec.SetEngineerCareerID(c.ID)
}

// SetEngineerPositionID sets the "engineerPosition" edge to the CompanyPosition entity by ID.
func (cec *CompanyEngineerCreate) SetEngineerPositionID(id int) *CompanyEngineerCreate {
	cec.mutation.SetEngineerPositionID(id)
	return cec
}

// SetNillableEngineerPositionID sets the "engineerPosition" edge to the CompanyPosition entity by ID if the given value is not nil.
func (cec *CompanyEngineerCreate) SetNillableEngineerPositionID(id *int) *CompanyEngineerCreate {
	if id != nil {
		cec = cec.SetEngineerPositionID(*id)
	}
	return cec
}

// SetEngineerPosition sets the "engineerPosition" edge to the CompanyPosition entity.
func (cec *CompanyEngineerCreate) SetEngineerPosition(c *CompanyPosition) *CompanyEngineerCreate {
	return cec.SetEngineerPositionID(c.ID)
}

// AddCompanyOwnerIDs adds the "companyOwners" edge to the CompanyDetail entity by IDs.
func (cec *CompanyEngineerCreate) AddCompanyOwnerIDs(ids ...int) *CompanyEngineerCreate {
	cec.mutation.AddCompanyOwnerIDs(ids...)
	return cec
}

// AddCompanyOwners adds the "companyOwners" edges to the CompanyDetail entity.
func (cec *CompanyEngineerCreate) AddCompanyOwners(c ...*CompanyDetail) *CompanyEngineerCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cec.AddCompanyOwnerIDs(ids...)
}

// AddInspectorIDs adds the "inspectors" edge to the JobDetail entity by IDs.
func (cec *CompanyEngineerCreate) AddInspectorIDs(ids ...int) *CompanyEngineerCreate {
	cec.mutation.AddInspectorIDs(ids...)
	return cec
}

// AddInspectors adds the "inspectors" edges to the JobDetail entity.
func (cec *CompanyEngineerCreate) AddInspectors(j ...*JobDetail) *CompanyEngineerCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cec.AddInspectorIDs(ids...)
}

// AddArchitectIDs adds the "architects" edge to the JobDetail entity by IDs.
func (cec *CompanyEngineerCreate) AddArchitectIDs(ids ...int) *CompanyEngineerCreate {
	cec.mutation.AddArchitectIDs(ids...)
	return cec
}

// AddArchitects adds the "architects" edges to the JobDetail entity.
func (cec *CompanyEngineerCreate) AddArchitects(j ...*JobDetail) *CompanyEngineerCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cec.AddArchitectIDs(ids...)
}

// AddStaticIDs adds the "statics" edge to the JobDetail entity by IDs.
func (cec *CompanyEngineerCreate) AddStaticIDs(ids ...int) *CompanyEngineerCreate {
	cec.mutation.AddStaticIDs(ids...)
	return cec
}

// AddStatics adds the "statics" edges to the JobDetail entity.
func (cec *CompanyEngineerCreate) AddStatics(j ...*JobDetail) *CompanyEngineerCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cec.AddStaticIDs(ids...)
}

// AddMechanicIDs adds the "mechanics" edge to the JobDetail entity by IDs.
func (cec *CompanyEngineerCreate) AddMechanicIDs(ids ...int) *CompanyEngineerCreate {
	cec.mutation.AddMechanicIDs(ids...)
	return cec
}

// AddMechanics adds the "mechanics" edges to the JobDetail entity.
func (cec *CompanyEngineerCreate) AddMechanics(j ...*JobDetail) *CompanyEngineerCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cec.AddMechanicIDs(ids...)
}

// AddElectricIDs adds the "electrics" edge to the JobDetail entity by IDs.
func (cec *CompanyEngineerCreate) AddElectricIDs(ids ...int) *CompanyEngineerCreate {
	cec.mutation.AddElectricIDs(ids...)
	return cec
}

// AddElectrics adds the "electrics" edges to the JobDetail entity.
func (cec *CompanyEngineerCreate) AddElectrics(j ...*JobDetail) *CompanyEngineerCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cec.AddElectricIDs(ids...)
}

// AddControllerIDs adds the "controllers" edge to the JobDetail entity by IDs.
func (cec *CompanyEngineerCreate) AddControllerIDs(ids ...int) *CompanyEngineerCreate {
	cec.mutation.AddControllerIDs(ids...)
	return cec
}

// AddControllers adds the "controllers" edges to the JobDetail entity.
func (cec *CompanyEngineerCreate) AddControllers(j ...*JobDetail) *CompanyEngineerCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cec.AddControllerIDs(ids...)
}

// AddMechaniccontrollerIDs adds the "mechaniccontrollers" edge to the JobDetail entity by IDs.
func (cec *CompanyEngineerCreate) AddMechaniccontrollerIDs(ids ...int) *CompanyEngineerCreate {
	cec.mutation.AddMechaniccontrollerIDs(ids...)
	return cec
}

// AddMechaniccontrollers adds the "mechaniccontrollers" edges to the JobDetail entity.
func (cec *CompanyEngineerCreate) AddMechaniccontrollers(j ...*JobDetail) *CompanyEngineerCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cec.AddMechaniccontrollerIDs(ids...)
}

// AddElectriccontrollerIDs adds the "electriccontrollers" edge to the JobDetail entity by IDs.
func (cec *CompanyEngineerCreate) AddElectriccontrollerIDs(ids ...int) *CompanyEngineerCreate {
	cec.mutation.AddElectriccontrollerIDs(ids...)
	return cec
}

// AddElectriccontrollers adds the "electriccontrollers" edges to the JobDetail entity.
func (cec *CompanyEngineerCreate) AddElectriccontrollers(j ...*JobDetail) *CompanyEngineerCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cec.AddElectriccontrollerIDs(ids...)
}

// Mutation returns the CompanyEngineerMutation object of the builder.
func (cec *CompanyEngineerCreate) Mutation() *CompanyEngineerMutation {
	return cec.mutation
}

// Save creates the CompanyEngineer in the database.
func (cec *CompanyEngineerCreate) Save(ctx context.Context) (*CompanyEngineer, error) {
	cec.defaults()
	return withHooks(ctx, cec.sqlSave, cec.mutation, cec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cec *CompanyEngineerCreate) SaveX(ctx context.Context) *CompanyEngineer {
	v, err := cec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cec *CompanyEngineerCreate) Exec(ctx context.Context) error {
	_, err := cec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cec *CompanyEngineerCreate) ExecX(ctx context.Context) {
	if err := cec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cec *CompanyEngineerCreate) defaults() {
	if _, ok := cec.mutation.Name(); !ok {
		v := companyengineer.DefaultName
		cec.mutation.SetName(v)
	}
	if _, ok := cec.mutation.Status(); !ok {
		v := companyengineer.DefaultStatus
		cec.mutation.SetStatus(v)
	}
	if _, ok := cec.mutation.Deleted(); !ok {
		v := companyengineer.DefaultDeleted
		cec.mutation.SetDeleted(v)
	}
	if _, ok := cec.mutation.CreatedAt(); !ok {
		v := companyengineer.DefaultCreatedAt()
		cec.mutation.SetCreatedAt(v)
	}
	if _, ok := cec.mutation.UpdatedAt(); !ok {
		v := companyengineer.DefaultUpdatedAt()
		cec.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cec *CompanyEngineerCreate) check() error {
	if _, ok := cec.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "CompanyEngineer.Name"`)}
	}
	if _, ok := cec.mutation.Status(); !ok {
		return &ValidationError{Name: "Status", err: errors.New(`ent: missing required field "CompanyEngineer.Status"`)}
	}
	if _, ok := cec.mutation.Deleted(); !ok {
		return &ValidationError{Name: "Deleted", err: errors.New(`ent: missing required field "CompanyEngineer.Deleted"`)}
	}
	if _, ok := cec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "CompanyEngineer.CreatedAt"`)}
	}
	if _, ok := cec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "CompanyEngineer.UpdatedAt"`)}
	}
	return nil
}

func (cec *CompanyEngineerCreate) sqlSave(ctx context.Context) (*CompanyEngineer, error) {
	if err := cec.check(); err != nil {
		return nil, err
	}
	_node, _spec := cec.createSpec()
	if err := sqlgraph.CreateNode(ctx, cec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cec.mutation.id = &_node.ID
	cec.mutation.done = true
	return _node, nil
}

func (cec *CompanyEngineerCreate) createSpec() (*CompanyEngineer, *sqlgraph.CreateSpec) {
	var (
		_node = &CompanyEngineer{config: cec.config}
		_spec = sqlgraph.NewCreateSpec(companyengineer.Table, sqlgraph.NewFieldSpec(companyengineer.FieldID, field.TypeInt))
	)
	if value, ok := cec.mutation.Name(); ok {
		_spec.SetField(companyengineer.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cec.mutation.Address(); ok {
		_spec.SetField(companyengineer.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := cec.mutation.Email(); ok {
		_spec.SetField(companyengineer.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := cec.mutation.TcNo(); ok {
		_spec.SetField(companyengineer.FieldTcNo, field.TypeInt, value)
		_node.TcNo = value
	}
	if value, ok := cec.mutation.Phone(); ok {
		_spec.SetField(companyengineer.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := cec.mutation.RegNo(); ok {
		_spec.SetField(companyengineer.FieldRegNo, field.TypeInt, value)
		_node.RegNo = value
	}
	if value, ok := cec.mutation.CertNo(); ok {
		_spec.SetField(companyengineer.FieldCertNo, field.TypeInt, value)
		_node.CertNo = value
	}
	if value, ok := cec.mutation.Note(); ok {
		_spec.SetField(companyengineer.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if value, ok := cec.mutation.Status(); ok {
		_spec.SetField(companyengineer.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := cec.mutation.Deleted(); ok {
		_spec.SetField(companyengineer.FieldDeleted, field.TypeInt, value)
		_node.Deleted = value
	}
	if value, ok := cec.mutation.Employment(); ok {
		_spec.SetField(companyengineer.FieldEmployment, field.TypeTime, value)
		_node.Employment = value
	}
	if value, ok := cec.mutation.Dismissal(); ok {
		_spec.SetField(companyengineer.FieldDismissal, field.TypeTime, value)
		_node.Dismissal = value
	}
	if value, ok := cec.mutation.CreatedAt(); ok {
		_spec.SetField(companyengineer.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cec.mutation.UpdatedAt(); ok {
		_spec.SetField(companyengineer.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := cec.mutation.EngineerCareerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companyengineer.EngineerCareerTable,
			Columns: []string{companyengineer.EngineerCareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companycareer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.career_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.EngineerPositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companyengineer.EngineerPositionTable,
			Columns: []string{companyengineer.EngineerPositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyposition.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.position_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.CompanyOwnersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyengineer.CompanyOwnersTable,
			Columns: []string{companyengineer.CompanyOwnersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companydetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.InspectorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyengineer.InspectorsTable,
			Columns: []string{companyengineer.InspectorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.ArchitectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyengineer.ArchitectsTable,
			Columns: []string{companyengineer.ArchitectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.StaticsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyengineer.StaticsTable,
			Columns: []string{companyengineer.StaticsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.MechanicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyengineer.MechanicsTable,
			Columns: []string{companyengineer.MechanicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.ElectricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyengineer.ElectricsTable,
			Columns: []string{companyengineer.ElectricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.ControllersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyengineer.ControllersTable,
			Columns: []string{companyengineer.ControllersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.MechaniccontrollersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyengineer.MechaniccontrollersTable,
			Columns: []string{companyengineer.MechaniccontrollersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.ElectriccontrollersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyengineer.ElectriccontrollersTable,
			Columns: []string{companyengineer.ElectriccontrollersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CompanyEngineerCreateBulk is the builder for creating many CompanyEngineer entities in bulk.
type CompanyEngineerCreateBulk struct {
	config
	err      error
	builders []*CompanyEngineerCreate
}

// Save creates the CompanyEngineer entities in the database.
func (cecb *CompanyEngineerCreateBulk) Save(ctx context.Context) ([]*CompanyEngineer, error) {
	if cecb.err != nil {
		return nil, cecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cecb.builders))
	nodes := make([]*CompanyEngineer, len(cecb.builders))
	mutators := make([]Mutator, len(cecb.builders))
	for i := range cecb.builders {
		func(i int, root context.Context) {
			builder := cecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompanyEngineerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cecb *CompanyEngineerCreateBulk) SaveX(ctx context.Context) []*CompanyEngineer {
	v, err := cecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cecb *CompanyEngineerCreateBulk) Exec(ctx context.Context) error {
	_, err := cecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cecb *CompanyEngineerCreateBulk) ExecX(ctx context.Context) {
	if err := cecb.Exec(ctx); err != nil {
		panic(err)
	}
}
