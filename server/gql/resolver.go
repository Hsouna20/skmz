package gql

import (
	"context"
	"github.com/Hsouna20/skmz/tree/main/server/db"
	"github.com/Hsouna20/skmz/tree/main/server/gql/gen"
	"github.com/Hsouna20/skmz/tree/main/server/model"
)

type Resolver struct {
	DB db.DB
}

func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Programmers(ctx context.Context, skill string) ([]*model.Programmer, error) {
	return r.DB.GetProgrammers(skill)
}
