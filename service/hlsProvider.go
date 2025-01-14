package service

import (
	"context"
	"fmt"
	"go.uber.org/dig"
	"main/service/api"
	"main/service/internal/config"
	"net/http"
)

func HlsProvider() Service {
	once.Do(func() {
		srv = &hlsProvider{}
	})

	return srv
}

type hlsProvider struct{}

func (srv *hlsProvider) Run() {
	container := dig.New()

	srv.provideConfig(container)

	srv.provideService(container)

	srv.provideController(container)

	srv.invokeApiRoutes(container)

	if err := container.Invoke(srv.run); err != nil {
		panic(err)
	}
}

func (srv *hlsProvider) provideConfig(container *dig.Container) {
	if err := container.Provide(config.NewConfig); err != nil {
		panic(err)
	}
}

func (srv *hlsProvider) provideService(container *dig.Container) {
	if err := container.Provide(func() context.Context {
		return context.TODO()
	}); err != nil {
		panic(err)
	}

	if err := container.Provide(api.NewServer); err != nil {
		panic(err)
	}

	if err := container.Provide(api.NewGinEngine); err != nil {
		panic(err)
	}

	if err := container.Provide(api.NewRouterRoot); err != nil {
		panic(err)
	}
}

func (srv *hlsProvider) provideController(container *dig.Container) {
}

func (srv *hlsProvider) invokeApiRoutes(container *dig.Container) {
	if err := container.Invoke(api.NewBasic); err != nil {
		panic(err)
	}
}

func (srv *hlsProvider) run(server *http.Server) {
	fmt.Printf("HlsProvider starts at %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
