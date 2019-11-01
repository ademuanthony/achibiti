package web

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

type key int

const (
	ctxSyncDataType key = iota
	ctxChartType
)

func syncDataType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ctxSyncDataType,
			chi.URLParam(r, "dataType"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// getSyncTypeCtx retrieves the syncDataType data from the request context.
// If not set, the return value is an empty string.
func getSyncDataTypeCtx(r *http.Request) string {
	syncType, ok := r.Context().Value(ctxSyncDataType).(string)
	if !ok {
		log.Trace("sync type not set")
		return ""
	}
	return syncType
}

// chartTypeCtx returns a http.HandlerFunc that embeds the value at the url
// part {charttype} into the request context.
func chartTypeCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ctxChartType,
			chi.URLParam(r, "charttype"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// getChartTypeCtx retrieves the ctxChart data from the request context.
// If not set, the return value is an empty string.
func getChartTypeCtx(r *http.Request) string {
	chartType, ok := r.Context().Value(ctxChartType).(string)
	if !ok {
		log.Trace("chart type not set")
		return ""
	}
	return chartType
}
