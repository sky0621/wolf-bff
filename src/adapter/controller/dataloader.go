package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/sky0621/wolf-bff/src/adapter/controller/gqlmodel"
)

const loadersKey = "dataLoaders"

type DataLoaders struct {
	ContentLoader *gqlmodel.ContentLoader
}

func DataLoaderMiddleware(resolver *Resolver, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &DataLoaders{
			ContentLoader: gqlmodel.NewContentLoader(gqlmodel.ContentLoaderConfig{
				MaxBatch: 100,
				Wait:     1 * time.Millisecond,
				Fetch: func(keys []string) ([][]gqlmodel.Content, []error) {

					// FIXME:

					return nil, nil
				},
			}),
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *DataLoaders {
	return ctx.Value(loadersKey).(*DataLoaders)
}
