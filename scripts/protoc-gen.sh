docker run --rm -v $(pwd):/defs namely/protoc-all \
    -d api/protobuf-spec \
    -i scripts \
    -i vendor \
    --with-docs \
    --with-gateway \
    -o internal/guestcovider_out \
    -l go


mv ./internal/guestcovider_out/guestcovider-services.swagger.json ./api/swagger-spec/swagger.json
rm -R ./internal/guestcoviderpb
mv ./internal/guestcovider_out/internal/guestcoviderpb ./internal
rm ./internal/guestcoviderpb/*.gw.go
rm -R ./internal/guestcovider_out