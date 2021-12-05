# moreHandlers
 
moreHandlers is a library which makes possible the use of multiple handlers for the MCBE server software https://github.com/df-mc/dragonfly

You may create a new `player.Handler` using the `New` method:
```go
handler := moreHandlers.New(handler1, handler2)
```

this example assumes that `handler1` and `handler2` are `player.Handler`:

```go
// PlayerHandler is the default handler used for the player.
// It does basic stuff.
type PlayerHandler struct{
     player.NopHandler
}

// Define handler1.
var handler1 = PlayerHandler{}

// CustomItemHandler is the handler which is used to handle when custom items are used.
type CustomItemHandler struct{
     player.NopHandler
}

// Define handler2.
var handler2 = CustomItemHandler{}
```
