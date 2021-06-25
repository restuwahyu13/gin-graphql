package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/restuwahyu13/gin-graphql/graph/model"
)

func (r *mutationResolver) CreateTeacher(ctx context.Context, input model.TeacherInput) (*model.Teacher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTeacher(ctx context.Context, id string) (*model.Teacher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTeacher(ctx context.Context, id string, input model.TeacherInput) (*model.Teacher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Teachers(ctx context.Context) ([]*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Teacher(ctx context.Context, id string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}
