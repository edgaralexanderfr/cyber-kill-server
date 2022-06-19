package game

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/net"

var InputEvents []net.InputEvent = []net.InputEvent{
	{
		Event: "MoveUp",
		Type:  net.INPUT_EVENT_TYPE_DIGITAL,
	},
	{
		Event: "MoveRight",
		Type:  net.INPUT_EVENT_TYPE_DIGITAL,
	},
	{
		Event: "MoveDown",
		Type:  net.INPUT_EVENT_TYPE_DIGITAL,
	},
	{
		Event: "MoveLeft",
		Type:  net.INPUT_EVENT_TYPE_DIGITAL,
	},
	{
		Event: "Shoot",
		Type:  net.INPUT_EVENT_TYPE_DIGITAL,
	},
	{
		Event: "SetAim",
		Type:  net.INPUT_EVENT_TYPE_ANALOGIC,
	},
}
