package moreHandlers

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
)

func NewInventoryHandler(handlers ...InventoryCHandler) *InventoryHandler {
	return &InventoryHandler{
		handlers: handlers,
	}
}

func (h *InventoryHandler) AddHandler(handler InventoryCHandler) {
	h.handlers = append(h.handlers, handler)
}
func (h *InventoryHandler) RemoveHandler(handler InventoryCHandler) {
	for n, i := range h.handlers {
		if i.Name() == handler.Name() {
			h.handlers = append(h.handlers[:n], h.handlers[n+1:]...)
		}
	}
}

type InventoryHandler struct {
	handlers []InventoryCHandler
}

func (h *InventoryHandler) HandleTake(ctx *event.Context, slot int, it item.Stack) {
	for _, handler := range h.handlers {
		handler.HandleTake(ctx, slot, it)
	}
}
func (h *InventoryHandler) HandlePlace(ctx *event.Context, slot int, it item.Stack) {
	for _, handler := range h.handlers {
		handler.HandlePlace(ctx, slot, it)
	}
}
func (h *InventoryHandler) HandleDrop(ctx *event.Context, slot int, it item.Stack) {
	for _, handler := range h.handlers {
		handler.HandleDrop(ctx, slot, it)
	}
}
