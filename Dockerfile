FROM ubi8/go-toolset as build
COPY . .
RUN pwd && go mod download && \
    go build -o main .

FROM ubi8/ubi-micro
COPY --from=build /opt/app-root/src/main .
EXPOSE 8080
CMD ./main
