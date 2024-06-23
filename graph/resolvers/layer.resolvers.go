package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"
	"gqlgen-ent/ent"
	"gqlgen-ent/ent/jobdetail"
	"gqlgen-ent/ent/joblayer"
	"gqlgen-ent/graph/generated"
	"gqlgen-ent/graph/model"
	"gqlgen-ent/middlewares"
	"strconv"
	"strings"
	"time"
)

// MoldDate is the resolver for the MoldDate field.
func (r *jobLayerResolver) MoldDate(ctx context.Context, obj *ent.JobLayer) (*string, error) {
	if obj.MoldDate.IsZero() {
		return nil, nil
	}
	moldDate := obj.MoldDate.Format("2006-01-02")
	return &moldDate, nil
}

// ConcreteDate is the resolver for the ConcreteDate field.
func (r *jobLayerResolver) ConcreteDate(ctx context.Context, obj *ent.JobLayer) (*string, error) {
	if obj.ConcreteDate.IsZero() {
		return nil, nil
	}
	concreteDate := obj.ConcreteDate.Format("2006-01-02")
	return &concreteDate, nil
}

// Job is the resolver for the Job field.
func (r *jobLayerResolver) Job(ctx context.Context, obj *ent.JobLayer) (*ent.JobDetail, error) {
	client := middlewares.GetClientFromContext(ctx)
	return client.JobDetail.Query().Where(jobdetail.HasLayersWith(joblayer.IDEQ(obj.ID))).Only(ctx)
}

// CreateLayer is the resolver for the createLayer field.
func (r *mutationResolver) CreateLayer(ctx context.Context, input model.JobLayerInput) (*ent.JobLayer, error) {
	// panic(fmt.Errorf("not implemented: CreateLayer - createLayer"))
	client := middlewares.GetClientFromContext(ctx)

	var moldDatePtr *time.Time
	if input.MoldDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *input.MoldDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse contract date: %v", err)
		}
		moldDatePtr = &parsedDate
	}

	var concreteDatePtr *time.Time
	if input.ConcreteDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *input.ConcreteDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse start date: %v", err)
		}
		concreteDatePtr = &parsedDate
	}

	uppercaseName := strings.ToUpper(*input.Name)

	layer, err := client.JobLayer.Create().
		SetName(uppercaseName).
		SetMetre(*input.Metre).
		SetNillableMoldDate(moldDatePtr).
		SetNillableConcreteDate(concreteDatePtr).
		SetNillableSamples(input.Samples).
		SetNillableConcreteClass(input.ConcreteClass).
		SetNillableWeekResult(input.WeekResult).
		SetNillableMonthResult(input.MonthResult).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// İş detayını bul
	jobDetail, err := client.JobDetail.Query().Where(jobdetail.YibfNoEQ(*input.YibfNo)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş detayı bulunamadı: %v", err)
	}

	// İş detayına katmanı ekle
	_, err = client.JobDetail.UpdateOneID(jobDetail.ID).
		AddLayers(layer).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("katman iş ayrıntısına eklenemedi: %v", err)
	}

	return layer, nil
}

// UpdateLayer is the resolver for the updateLayer field.
func (r *mutationResolver) UpdateLayer(ctx context.Context, id string, input model.JobLayerInput) (*ent.JobLayer, error) {
	client := middlewares.GetClientFromContext(ctx)

	moldDate, err := time.Parse("2006-01-02", *input.MoldDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mold date: %v", err)
	}
	moldDatePtr := &moldDate

	concreteDate, err := time.Parse("2006-01-02", *input.ConcreteDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract date: %v", err)
	}
	concreteDatePtr := &concreteDate
	layerID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert layer ID: %v", err)
	}

	// Update Layer
	layer, err := client.JobLayer.UpdateOneID(layerID).
		SetNillableName(input.Name).
		SetNillableMetre(input.Metre).
		SetNillableMoldDate(moldDatePtr).
		SetNillableConcreteDate(concreteDatePtr).
		SetNillableSamples(input.Samples).
		SetNillableConcreteClass(input.ConcreteClass).
		SetNillableWeekResult(input.WeekResult).
		SetNillableMonthResult(input.MonthResult).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update layer: %v", err)
	}

	return layer, nil
}

// DeleteLayer is the resolver for the deleteLayer field.
func (r *mutationResolver) DeleteLayer(ctx context.Context, id string) (*ent.JobLayer, error) {
	client := middlewares.GetClientFromContext(ctx)

	// ID'yi string'den int'e dönüştür
	layerID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("geçersiz ID formatı: %v", err)
	}

	// İş katmanını ID ile sorgula
	layer, err := client.JobLayer.Query().Where(joblayer.IDEQ(layerID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("kayıt bulunamadı")
		}
		return nil, err
	}

	// İş katmanını sil
	err = client.JobLayer.DeleteOne(layer).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş katmanı silinemedi: %v", err)
	}

	return layer, nil
}

// Layer is the resolver for the Layer field.
func (r *queryResolver) Layer(ctx context.Context, filter *model.LayerFilterInput) ([]*ent.JobLayer, error) {
	client := middlewares.GetClientFromContext(ctx)
	query := client.JobLayer.Query()

	if filter != nil {
		if filter.ID != nil {
			query = query.Where(joblayer.IDEQ(*filter.ID))
		}
		if filter.YibfNo != nil {
			query = query.Where(joblayer.HasLayerWith(jobdetail.YibfNoEQ(*filter.YibfNo)))
		}
	}

	layers, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get layers: %v", err)
	}
	return layers, nil
}

// JobLayer returns generated.JobLayerResolver implementation.
func (r *Resolver) JobLayer() generated.JobLayerResolver { return &jobLayerResolver{r} }

type jobLayerResolver struct{ *Resolver }
