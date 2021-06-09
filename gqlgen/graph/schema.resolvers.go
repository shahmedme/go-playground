package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/shakilahmmeed/gqlgen-todos/graph/generated"
	"github.com/shakilahmmeed/gqlgen-todos/graph/model"
	"github.com/shakilahmmeed/gqlgen-todos/models"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*models.Link, error) {
	link2 := models.Link{
		Title:   input.Title,
		Address: input.Address,
	}

	r.DB.Create(&link2)
	return &link2, nil
}

func (r *mutationResolver) CreateBike(ctx context.Context, input model.NewBike) (*model.Bike, error) {
	bike := model.Bike{
		ID:    input.ID,
		Model: input.Model,
		Owner: input.Owner,
	}

	return &bike, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*models.Link, error) {
	var links []*models.Link
	r.DB.Preload("Links").Find(&links)

	return links, nil
}

func (r queryResolver) Students(ctx context.Context) ([]*models.Student, error) {
	var students []*models.Student
	studentOne := models.Student{
		ID:         5,
		Name:       "Shakil",
		VarsityId:  9376,
		Email:      "shakil@gmail.com",
		Department: "CSE",
	}

	students = append(students, &studentOne)
	return students, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
