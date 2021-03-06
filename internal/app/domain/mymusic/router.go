package mymusic

import (
	"github.com/typical-go/typical-rest-server/internal/app/domain/mymusic/controller"
	"github.com/typical-go/typical-rest-server/pkg/echokit"
	"go.uber.org/dig"
)

type (
	// Router to server
	Router struct {
		dig.In
		SongCntrl controller.SongCntrl
	}
)

var _ echokit.Router = (*Router)(nil)

// SetRoute to echo server
func (r *Router) SetRoute(e echokit.Server) {
	group := e.Group("/mymusic")
	echokit.SetRoute(group,
		&r.SongCntrl,
	)
}
