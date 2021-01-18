package core

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// An Action to change the state of the game based on user input.
// It acts as a higher level abstraction than inputs, allowing you
// to map multiple input interfaces to standard operations.
type Action interface{}

type actionHandler struct {
	actions []Action
}

// ActionHandler creates a new chainable builder to declare what
// inputs trigger what actions. After attaching Actions to different
// input mechanisms, call Execute to execute the matching actions, if any.
func ActionHandler() *actionHandler {
	return &actionHandler{actions: make([]Action, 0, 1)}
}

// OnInput is used to write any arbitrary input handling code and
// return actions.
func (i *actionHandler) OnInput(h func() []Action) *actionHandler {
	actions := h()
	i.actions = append(i.actions, actions...)
	return i
}

// OnKey causes the matching action to fire when a given keyboard button is pressed
func (i *actionHandler) OnKey(keys map[ebiten.Key]Action) *actionHandler {
	for key, action := range keys {
		// inpututil.IsKeyJustPressed(key)
		if !ebiten.IsKeyPressed(key) {
			continue
		}
		if action != nil {
			i.actions = append(i.actions, action)
		}
	}
	return i
}

// var mouseButtons = []ebiten.MouseButton{ebiten.MouseButtonLeft, ebiten.MouseButtonMiddle, ebiten.MouseButtonRight}
// func (i *actionHandler) OnMouseClick(handler func(btn ebiten.MouseButton, screenPoint image.Point) Action) *actionHandler {
// 	for _, b := range mouseButtons {
// 		if inpututil.IsMouseButtonJustPressed(b) {
// 			action := handler(b, image.Pt(ebiten.CursorPosition()))
// 			if action != nil {
// 				i.actions = append(i.actions, action)
// 			}
// 		}
// 	}
// 	return i
// }

func (i *actionHandler) OnLeftMouseClick(handler func(screenPoint image.Point) Action) *actionHandler {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		action := handler(image.Pt(ebiten.CursorPosition()))
		if action != nil {
			i.actions = append(i.actions, action)
		}
	}
	return i
}

// func (i *actionHandler) OnGamepadButton(id int, buttons map[ebiten.GamepadButton]Action) *actionHandler {
// 	for button, action := range buttons {
// 		if !ebiten.IsGamepadButtonPressed(id, button) {
// 			continue
// 		}
// 		if action != nil {
// 			i.actions = append(i.actions, action)
// 		}
// 	}
// 	return i
// }

// Execute will call the given handler for any actions matched by the builders
func (i *actionHandler) Execute(h func(Action)) {
	for _, a := range i.actions {
		h(a)
	}
	i.actions = i.actions[:0]
}
