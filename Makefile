GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

TAGS=geojson,geoparquet,mbtiles,pmtiles

# https://github.com/marcboeker/go-duckdb?tab=readme-ov-file#vendoring
# go install github.com/goware/modvendor@latest
modvendor:
	modvendor -copy="**/*.a **/*.h" -v

cli:
	go build -tags $(TAGS) -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/show cmd/show/main.go

