module github.com/ademuanthony/achibiti

go 1.13

require (
	github.com/ademuanthony/achibiti/acl v0.0.0-00010101000000-000000000000
	github.com/ademuanthony/achibiti/utils v0.0.0-20191101145116-c955652b8ea9
	github.com/davecgh/go-spew v1.1.1
	github.com/decred/slog v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/friendsofgo/errors v0.9.2
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/cors v1.0.0
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jessevdk/go-flags v1.4.0
	github.com/jrick/logrotate v1.0.0
	github.com/kat-co/vala v0.0.0-20170210184112-42e1d8b61f12
	github.com/lib/pq v1.2.0
	github.com/micro/go-micro v1.14.0
	github.com/raedahgroup/dcrextdata v0.0.0-20191001203600-13e716f9d8c3
	github.com/spf13/viper v1.5.0
	github.com/volatiletech/sqlboiler v3.6.0+incompatible
)

replace (
	github.com/ademuanthony/achibiti/acl => ./acl
	github.com/ademuanthony/achibiti/hr => ./hr
	github.com/ademuanthony/achibiti/utils => ./utils
)
