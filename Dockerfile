ARG GO_VERSION=1.24
ARG NODE_VERSION=22.14.0

FROM node:${NODE_VERSION}-alpine as frontend-build

WORKDIR /app
COPY ./frontend/package.json ./frontend/package-lock.json ./
RUN npm install
COPY ./frontend .
RUN npm run build

FROM golang:${GO_VERSION} AS build
WORKDIR /src

ENV CGO_ENABLED=1

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

COPY . . 
COPY --from=frontend-build /pkg/embed/build ./pkg/static/build

RUN --mount=type=cache,target=/go/pkg/mod/ \
    go build -ldflags='-s -w -extldflags "-static"' -o /bin/cvrs ./cmd/backend/main.go
    # static linking is necessary because of CGO dependency
    # -s -w removes debug info for smaller bin

FROM alpine:latest AS final

ARG GIT_COMMIT=unspecified
LABEL org.opencontainers.image.version=$GIT_COMMIT
LABEL org.opencontainers.image.source=https://github.com/Pineapple217/cvrs

WORKDIR /app
COPY --from=build /bin/cvrs /app/cvrs

EXPOSE 3000

CMD [ "/app/cvrs", "run"]
