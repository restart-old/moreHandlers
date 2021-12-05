package main

import (
	"time"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/dragonfly-on-steroids/moreHandlers"
)

func main() {
	c := server.DefaultConfig()
	s := server.New(&c, nil)
	s.Start()
	for {
		p, err := s.Accept()
		if err != nil {
			return
		}
		handler2 := &testHandler2{p: p}
		handler1 := &testHandler1{p: p}
		h := moreHandlers.New(handler1, handler2)
		p.Handle(h)
		go func() {
			time.Sleep(5 * time.Second)
			h.RemoveHandler(&testHandler2{})
		}()
	}
}

type testHandler1 struct {
	p *player.Player
	player.NopHandler
}

func (*testHandler1) Name() string {
	return "handler1"
}

func (h *testHandler1) HandlePunchAir(ctx *event.Context) {
	h.p.Message("handler 1")
}

type testHandler2 struct {
	p *player.Player
	player.NopHandler
}

func (*testHandler2) Name() string {
	return "handler2"
}

func (h *testHandler2) HandlePunchAir(ctx *event.Context) {
	h.p.Message("handler 2")
}
