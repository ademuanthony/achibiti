module github.com/ademuanthony/achibiti

go 1.13

require (
	github.com/ademuanthony/achibiti/accounts v0.0.0-20191101111714-6276f3277d84
	github.com/ademuanthony/achibiti/utils v0.0.0-20191101111714-6276f3277d84
	github.com/decred/slog v1.0.0
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jessevdk/go-flags v1.4.0
	github.com/jrick/logrotate v1.0.0
	github.com/lib/pq v1.2.0
	github.com/raedahgroup/dcrextdata v0.0.0-20191001203600-13e716f9d8c3
	github.com/volatiletech/sqlboiler v3.6.0+incompatible
)

replace (
	github.com/ademuanthony/achibiti/accounts => ./accounts
	github.com/ademuanthony/achibiti/utils => ./utils
)
