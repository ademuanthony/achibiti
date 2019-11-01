module github.com/ademuanthony/achibiti/accounts

go 1.13

require (
	github.com/DATA-DOG/go-sqlmock v1.3.3 // indirect
	github.com/ademuanthony/achibiti/utils v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/friendsofgo/errors v0.9.2
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/jessevdk/go-flags v1.4.0
	github.com/kat-co/vala v0.0.0-20170210184112-42e1d8b61f12
	github.com/lib/pq v1.2.0
	github.com/micro/go-micro v1.14.0
	github.com/spf13/viper v1.4.0
	github.com/volatiletech/inflect v0.0.0-20170731032912-e7201282ae8d // indirect
	github.com/volatiletech/null v8.0.0+incompatible // indirect
	github.com/volatiletech/sqlboiler v3.6.0+incompatible
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
)

replace github.com/ademuanthony/achibiti/utils => ../utils
