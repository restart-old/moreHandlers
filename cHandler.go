package moreHandlers

import (
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
)

type PlayerCHandler interface {
	player.Handler
	Name() string
}
type InventoryCHandler interface {
	inventory.Handler
	Name() string
}
