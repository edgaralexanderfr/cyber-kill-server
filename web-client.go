package main

import (
	"syscall/js"
)

var (
	global                js.Value
	requestAnimationFrame js.Value
	scene                 js.Value
	camera                js.Value
	renderer              js.Value
	cube                  js.Value
	animateFunc           js.Func
)

func main() {
	global = js.Global()
	requestAnimationFrame = global.Get("requestAnimationFrame")
	window := global.Get("window")
	document := global.Get("document")
	body := document.Get("body")
	THREE := global.Get("THREE")

	innerWidth := window.Get("innerWidth").Int()
	innerHeight := window.Get("innerHeight").Int()

	scene = THREE.Get("Scene").New()
	camera = THREE.Get("PerspectiveCamera").New(75, innerWidth/innerHeight, 0.1, 1000)

	renderer = THREE.Get("WebGLRenderer").New()
	renderer.Call("setSize", innerWidth, innerHeight)
	body.Call("appendChild", renderer.Get("domElement"))

	geometry := THREE.Get("BoxGeometry").New(1, 1, 1)
	material := THREE.Get("MeshBasicMaterial").New(map[string]interface{}{
		"color": 0x00ff00,
	})

	cube = THREE.Get("Mesh").New(geometry, material)
	scene.Call("add", cube)

	camera.Get("position").Set("z", 5)

	animateFunc = js.FuncOf(animate)

	animate(js.Value{}, nil)
	select {}
}

func animate(this js.Value, args []js.Value) interface{} {
	requestAnimationFrame.Invoke(animateFunc)
	renderer.Call("render", scene, camera)

	cube.Get("rotation").Set("x", cube.Get("rotation").Get("x").Float()+0.01)
	cube.Get("rotation").Set("y", cube.Get("rotation").Get("y").Float()+0.01)

	return nil
}
