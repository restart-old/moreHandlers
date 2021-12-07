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
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type CHandler interface {
	player.Handler
	Name() string
}

func New(handlers ...CHandler) *Handler {
	return &Handler{
		handlers: handlers,
	}
}

func (h *Handler) AddHandler(handler CHandler) {
	h.handlers = append(h.handlers, handler)
}
func (h *Handler) RemoveHandler(handler CHandler) {
	for n, i := range h.handlers {
		if i.Name() == handler.Name() {
			h.handlers = append(h.handlers[:n], h.handlers[n+1:]...)
		}
	}
}

type Handler struct {
	handlers []CHandler
}

func (h *Handler) HandleMove(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
	for _, handler := range h.handlers {
		handler.HandleMove(ctx, newPos, newYaw, newPitch)
	}
}
func (h *Handler) HandleTeleport(ctx *event.Context, pos mgl64.Vec3) {
	for _, handler := range h.handlers {
		handler.HandleTeleport(ctx, pos)
	}
}
func (h *Handler) HandleToggleSprint(ctx *event.Context, after bool) {
	for _, handler := range h.handlers {
		handler.HandleToggleSprint(ctx, after)
	}
}
func (h *Handler) HandleToggleSneak(ctx *event.Context, after bool) {
	for _, handler := range h.handlers {
		handler.HandleToggleSneak(ctx, after)
	}
}
func (h *Handler) HandleChat(ctx *event.Context, message *string) {
	for _, handler := range h.handlers {
		handler.HandleChat(ctx, message)
	}
}
func (h *Handler) HandleFoodLoss(ctx *event.Context, from, to int) {
	for _, handler := range h.handlers {
		handler.HandleFoodLoss(ctx, from, to)
	}
}
func (h *Handler) HandleHeal(ctx *event.Context, health *float64, src healing.Source) {
	for _, handler := range h.handlers {
		handler.HandleHeal(ctx, health, src)
	}
}
func (h *Handler) HandleHurt(ctx *event.Context, damage *float64, src damage.Source) {
	for _, handler := range h.handlers {
		handler.HandleHurt(ctx, damage, src)
	}
}
func (h *Handler) HandleDeath(src damage.Source) {
	for _, handler := range h.handlers {
		handler.HandleDeath(src)
	}
}
func (h *Handler) HandleRespawn(pos *mgl64.Vec3) {
	for _, handler := range h.handlers {
		handler.HandleRespawn(pos)
	}
}
func (h *Handler) HandleSkinChange(ctx *event.Context, skin skin.Skin) {
	for _, handler := range h.handlers {
		handler.HandleSkinChange(ctx, skin)
	}
}
func (h *Handler) HandleStartBreak(ctx *event.Context, pos cube.Pos) {
	for _, handler := range h.handlers {
		handler.HandleStartBreak(ctx, pos)
	}
}
func (h *Handler) HandleBlockBreak(ctx *event.Context, pos cube.Pos, drops *[]item.Stack) {
	for _, handler := range h.handlers {
		handler.HandleBlockBreak(ctx, pos, drops)
	}
}
func (h *Handler) HandleBlockPlace(ctx *event.Context, pos cube.Pos, b world.Block) {
	for _, handler := range h.handlers {
		handler.HandleBlockPlace(ctx, pos, b)
	}
}
func (h *Handler) HandleBlockPick(ctx *event.Context, pos cube.Pos, b world.Block) {
	for _, handler := range h.handlers {
		handler.HandleBlockPick(ctx, pos, b)
	}
}
func (h *Handler) HandleItemUse(ctx *event.Context) {
	for _, handler := range h.handlers {
		handler.HandleItemUse(ctx)
	}
}
func (h *Handler) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	for _, handler := range h.handlers {
		handler.HandleItemUseOnBlock(ctx, pos, face, clickPos)
	}
}
func (h *Handler) HandleItemUseOnEntity(ctx *event.Context, e world.Entity) {
	for _, handler := range h.handlers {
		handler.HandleItemUseOnEntity(ctx, e)
	}
}
func (h *Handler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64) {
	for _, handler := range h.handlers {
		handler.HandleAttackEntity(ctx, e, force, height)
	}
}
func (h *Handler) HandlePunchAir(ctx *event.Context) {
	for _, handler := range h.handlers {
		handler.HandlePunchAir(ctx)
	}
}
func (h *Handler) HandleSignEdit(ctx *event.Context, oldText, newText string) {
	for _, handler := range h.handlers {
		handler.HandleSignEdit(ctx, oldText, newText)
	}
}
func (h *Handler) HandleItemDamage(ctx *event.Context, i item.Stack, damage int) {
	for _, handler := range h.handlers {
		handler.HandleItemDamage(ctx, i, damage)
	}
}
func (h *Handler) HandleItemPickup(ctx *event.Context, i item.Stack) {
	for _, handler := range h.handlers {
		handler.HandleItemPickup(ctx, i)
	}
}
func (h *Handler) HandleItemDrop(ctx *event.Context, e *entity.Item) {
	for _, handler := range h.handlers {
		handler.HandleItemDrop(ctx, e)
	}
}
func (h *Handler) HandleTransfer(ctx *event.Context, addr *net.UDPAddr) {
	for _, handler := range h.handlers {
		handler.HandleTransfer(ctx, addr)
	}
}
func (h *Handler) HandleCommandExecution(ctx *event.Context, command cmd.Command, args []string) {
	for _, handler := range h.handlers {
		handler.HandleCommandExecution(ctx, command, args)
	}
}
func (h *Handler) HandleQuit() {
	for _, handler := range h.handlers {
		handler.HandleQuit()
	}
}
