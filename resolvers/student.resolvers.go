package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/restuwahyu13/gin-graphql/graph/model"
)

func (r *mutationResolver) CreateStudent(ctx context.Context, input model.StudentInput) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteStudent(ctx context.Context, id string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateStudent(ctx context.Context, id string, input model.StudentInput) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Students(ctx context.Context) ([]*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Student(ctx context.Context, id string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}
