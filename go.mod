module github.com/JeyKeyAlex/TourProject

go 1.23.0

toolchain go1.24.4

require (
	github.com/JeyKeyAlex/TourProject-proto v0.0.0-20251221125434-4f8e8a0f0fde
	github.com/go-chi/chi/v5 v5.2.3
	github.com/go-kit/kit v0.13.0
	github.com/google/uuid v1.6.0
	github.com/jackc/pgx/v5 v5.7.6
	github.com/joho/godotenv v1.5.1
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/v9 v9.17.2
	github.com/rs/zerolog v1.34.0
	github.com/sethvargo/go-envconfig v1.3.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
)

replace github.com/JeyKeyAlex/TourProject-proto => ../TestProject-proto
