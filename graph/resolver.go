package graph

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/codercollo/equimper"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

//go:generate go tool gqlgen generate

type Resolver struct {
	AuthService equimper.AuthService
}

type queryResolver struct {
	*Resolver
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct {
	*Resolver
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func buildBadRequestError(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code": http.StatusBadRequest,
		},
	}
}
