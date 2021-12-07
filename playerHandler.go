package moreHandlers

import (
	"net"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/entity/healing"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

func NewPlayerHandler(handlers ...PlayerCHandler) *PlayerHandler {
	return &PlayerHandler{
		handlers: handlers,
	}
}

func (h *PlayerHandler) AddHandler(handler PlayerCHandler) {
	h.handlers = append(h.handlers, handler)
}
func (h *PlayerHandler) RemoveHandler(handler PlayerCHandler) {
	for n, i := range h.handlers {
		if i.Name() == handler.Name() {
			h.handlers = append(h.handlers[:n], h.handlers[n+1:]...)
		}
	}
}

type PlayerHandler struct {
	handlers []PlayerCHandler
}

func (h *PlayerHandler) HandleMove(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
	for _, handler := range h.handlers {
		handler.HandleMove(ctx, newPos, newYaw, newPitch)
	}
}
func (h *PlayerHandler) HandleTeleport(ctx *event.Context, pos mgl64.Vec3) {
	for _, handler := range h.handlers {
		handler.HandleTeleport(ctx, pos)
	}
}
func (h *PlayerHandler) HandleToggleSprint(ctx *event.Context, after bool) {
	for _, handler := range h.handlers {
		handler.HandleToggleSprint(ctx, after)
	}
}
func (h *PlayerHandler) HandleToggleSneak(ctx *event.Context, after bool) {
	for _, handler := range h.handlers {
		handler.HandleToggleSneak(ctx, after)
	}
}
func (h *PlayerHandler) HandleChat(ctx *event.Context, message *string) {
	for _, handler := range h.handlers {
		handler.HandleChat(ctx, message)
	}
}
func (h *PlayerHandler) HandleFoodLoss(ctx *event.Context, from, to int) {
	for _, handler := range h.handlers {
		handler.HandleFoodLoss(ctx, from, to)
	}
}
func (h *PlayerHandler) HandleHeal(ctx *event.Context, health *float64, src healing.Source) {
	for _, handler := range h.handlers {
		handler.HandleHeal(ctx, health, src)
	}
}
func (h *PlayerHandler) HandleHurt(ctx *event.Context, damage *float64, src damage.Source) {
	for _, handler := range h.handlers {
		handler.HandleHurt(ctx, damage, src)
	}
}
func (h *PlayerHandler) HandleDeath(src damage.Source) {
	for _, handler := range h.handlers {
		handler.HandleDeath(src)
	}
}
func (h *PlayerHandler) HandleRespawn(pos *mgl64.Vec3) {
	for _, handler := range h.handlers {
		handler.HandleRespawn(pos)
	}
}
func (h *PlayerHandler) HandleSkinChange(ctx *event.Context, skin skin.Skin) {
	for _, handler := range h.handlers {
		handler.HandleSkinChange(ctx, skin)
	}
}
func (h *PlayerHandler) HandleStartBreak(ctx *event.Context, pos cube.Pos) {
	for _, handler := range h.handlers {
		handler.HandleStartBreak(ctx, pos)
	}
}
func (h *PlayerHandler) HandleBlockBreak(ctx *event.Context, pos cube.Pos, drops *[]item.Stack) {
	for _, handler := range h.handlers {
		handler.HandleBlockBreak(ctx, pos, drops)
	}
}
func (h *PlayerHandler) HandleBlockPlace(ctx *event.Context, pos cube.Pos, b world.Block) {
	for _, handler := range h.handlers {
		handler.HandleBlockPlace(ctx, pos, b)
	}
}
func (h *PlayerHandler) HandleBlockPick(ctx *event.Context, pos cube.Pos, b world.Block) {
	for _, handler := range h.handlers {
		handler.HandleBlockPick(ctx, pos, b)
	}
}
func (h *PlayerHandler) HandleItemUse(ctx *event.Context) {
	for _, handler := range h.handlers {
		handler.HandleItemUse(ctx)
	}
}
func (h *PlayerHandler) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	for _, handler := range h.handlers {
		handler.HandleItemUseOnBlock(ctx, pos, face, clickPos)
	}
}
func (h *PlayerHandler) HandleItemUseOnEntity(ctx *event.Context, e world.Entity) {
	for _, handler := range h.handlers {
		handler.HandleItemUseOnEntity(ctx, e)
	}
}
func (h *PlayerHandler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64) {
	for _, handler := range h.handlers {
		handler.HandleAttackEntity(ctx, e, force, height)
	}
}
func (h *PlayerHandler) HandlePunchAir(ctx *event.Context) {
	for _, handler := range h.handlers {
		handler.HandlePunchAir(ctx)
	}
}
func (h *PlayerHandler) HandleSignEdit(ctx *event.Context, oldText, newText string) {
	for _, handler := range h.handlers {
		handler.HandleSignEdit(ctx, oldText, newText)
	}
}
func (h *PlayerHandler) HandleItemDamage(ctx *event.Context, i item.Stack, damage int) {
	for _, handler := range h.handlers {
		handler.HandleItemDamage(ctx, i, damage)
	}
}
func (h *PlayerHandler) HandleItemPickup(ctx *event.Context, i item.Stack) {
	for _, handler := range h.handlers {
		handler.HandleItemPickup(ctx, i)
	}
}
func (h *PlayerHandler) HandleItemDrop(ctx *event.Context, e *entity.Item) {
	for _, handler := range h.handlers {
		handler.HandleItemDrop(ctx, e)
	}
}
func (h *PlayerHandler) HandleTransfer(ctx *event.Context, addr *net.UDPAddr) {
	for _, handler := range h.handlers {
		handler.HandleTransfer(ctx, addr)
	}
}
func (h *PlayerHandler) HandleCommandExecution(ctx *event.Context, command cmd.Command, args []string) {
	for _, handler := range h.handlers {
		handler.HandleCommandExecution(ctx, command, args)
	}
}
func (h *PlayerHandler) HandleQuit() {
	for _, handler := range h.handlers {
		handler.HandleQuit()
	}
}
