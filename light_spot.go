package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type SpotLight struct {
	spot1 *SpotLightMesh
	spot2 *SpotLightMesh
	spot3 *SpotLightMesh
	rot   float32
}

func init() {
	TestMap["light.spot"] = &SpotLight{}
}

func (t *SpotLight) Initialize(ctx *Context) {

	// Adds axis helper
	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)

	// Sets camera position
	ctx.Camera.GetCamera().SetPosition(0, 6, 10)

	// Creates base plane
	geom1 := geometry.NewPlane(6, 6, 16, 16)
	mat1 := material.NewPhong(math32.NewColor(1, 1, 1))
	mat1.SetSide(material.SideDouble)
	plane1 := graphic.NewMesh(geom1, mat1)
	plane1.SetRotationX(math32.Pi / 2)
	ctx.Scene.Add(plane1)

	// Creates left plane
	geom2 := geometry.NewPlane(6, 6, 16, 16)
	mat2 := material.NewPhong(math32.NewColor(1, 1, 1))
	mat2.SetSide(material.SideFront)
	plane2 := graphic.NewMesh(geom2, mat2)
	plane2.SetRotationY(math32.Pi / 2)
	plane2.SetPosition(-3, 3, 0)
	ctx.Scene.Add(plane2)

	// Creates right plane
	geom3 := geometry.NewPlane(6, 6, 16, 16)
	mat3 := material.NewPhong(math32.NewColor(1, 1, 1))
	mat3.SetSide(material.SideFront)
	mat3.SetSpecularColor(&math32.Color{1, 1, 1})
	plane3 := graphic.NewMesh(geom3, mat3)
	plane3.SetRotationY(-math32.Pi / 2)
	plane3.SetPosition(3, 3, 0)
	ctx.Scene.Add(plane3)

	// Creates red spot light
	t.spot1 = NewSpotLightMesh(&math32.Color{1, 0, 0})
	t.spot1.Mesh.SetPosition(-1, 3, 1)
	t.spot1.Light.SetDirection(&math32.Vector3{0, -1, 0})
	ctx.Scene.Add(t.spot1.Mesh)

	// Creates green spot light
	t.spot2 = NewSpotLightMesh(&math32.Color{0, 1, 0})
	t.spot2.Mesh.SetPosition(1, 3, -1)
	t.spot2.Light.SetDirection(&math32.Vector3{0, -1, 0})
	ctx.Scene.Add(t.spot2.Mesh)

	// Creates blue spot light
	t.spot3 = NewSpotLightMesh(&math32.Color{0, 0, 1})
	t.spot3.Mesh.SetPosition(0, 3, 0)
	t.spot3.Light.SetDirection(&math32.Vector3{0, -1, 0})
	ctx.Scene.Add(t.spot3.Mesh)

	// Subscribe to key events
	//	ctx.Gl.Subscribe(gls.OnKeyDown, t.onKey)
	//
	//	// Add controls
	//	if ctx.Control == nil {
	//		return
	//	}
	//	g := ctx.Control.AddGroup("Show lights")
	//	cb1 := g.AddCheckBox("Red").SetValue(t.spot1.Mesh.Visible())
	//	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
	//		t.spot1.Mesh.SetVisible(!t.spot1.Mesh.Visible())
	//	})
	//	cb2 := g.AddCheckBox("Green").SetValue(t.spot2.Mesh.Visible())
	//	cb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
	//		t.spot2.Mesh.SetVisible(!t.spot2.Mesh.Visible())
	//	})
	//	cb3 := g.AddCheckBox("Blue").SetValue(t.spot3.Mesh.Visible())
	//	cb3.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
	//		t.spot3.Mesh.SetVisible(!t.spot3.Mesh.Visible())
	//	})
}

func (t *SpotLight) Render(ctx *Context) {

	t.rot += float32(ctx.TimeDelta.Seconds())
	t.spot1.SetRotationZ(t.rot)
	t.spot2.SetRotationZ(-t.rot)
	t.spot3.Mesh.SetPosition(0, 3+1.5*math32.Sin(t.rot), 0)
}

//func (t *SpotLight) onKey(evname string, ev interface{}) {
//
//	kev := ev.(*gls.KeyEvent)
//	if kev.Action == gls.ActionRelease {
//		return
//	}
//	switch kev.Keycode {
//	case gls.KeyR:
//		t.spot1.Mesh.SetVisible(!t.spot1.Mesh.Visible())
//	case gls.KeyG:
//		t.spot2.Mesh.SetVisible(!t.spot2.Mesh.Visible())
//	case gls.KeyB:
//		t.spot3.Mesh.SetVisible(!t.spot3.Mesh.Visible())
//	}
//}

type SpotLightMesh struct {
	Mesh  *graphic.Mesh
	Light *light.Spot
}

func NewSpotLightMesh(color *math32.Color) *SpotLightMesh {

	l := new(SpotLightMesh)
	geom := geometry.NewCylinder(0, 0.3, 0.5, 16, 16, 0, 2*math32.Pi, true, true)
	mat1 := material.NewStandard(color)
	mat2 := material.NewStandard(color)
	mat2.SetEmissiveColor(color)
	l.Mesh = graphic.NewMesh(geom, nil)
	l.Mesh.AddGroupMaterial(mat1, 0)
	l.Mesh.AddGroupMaterial(mat2, 1)
	l.Light = light.NewSpot(color, 2.0)
	l.Light.SetLinearDecay(0.5)
	l.Light.SetQuadraticDecay(10)
	l.Mesh.Add(l.Light)
	return l
}

func (l *SpotLightMesh) AddScene(ctx *Context) {

	ctx.Scene.Add(l.Mesh)
}

func (l *SpotLightMesh) Position(x, y, z float32) {

	l.Mesh.SetPosition(x, y, z)
}

func (l *SpotLightMesh) SetRotationZ(rot float32) {

	var quat math32.Quaternion
	l.Mesh.SetRotationZ(rot)
	l.Mesh.WorldQuaternion(&quat)
	direction := math32.Vector3{0, -1, 0}
	direction.ApplyQuaternion(&quat)
	l.Light.SetDirection(&direction)
}
