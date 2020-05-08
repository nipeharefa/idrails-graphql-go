package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nipeharefa/idrails-graphql-go/graph/generated"
	"github.com/nipeharefa/idrails-graphql-go/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{}
	todo.Text = input.Text
	todo.Done = false

	row := r.db.QueryRow("insert into todo(text, done) values ($1, $2) returning id", todo.Text, todo.Done)

	err := row.Scan(&todo.ID)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos := make([]*model.Todo, 0)
	rows, err := r.db.Query("select * from todo order by id desc")
	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		todo := &model.Todo{}
		_ = rows.Scan(&todo.ID, &todo.Text, &todo.Done)
		todos = append(todos, todo)
	}

	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
