package s3dm

import "math"

type Frustum struct {
	Xform
	Planes [6]Plane
	Near   float64
	Far    float64
	Fovy   float64
	Aspect float64
}

func NewFrustum(near float64, far float64, fovy float64, aspect float64) *Frustum {
	frustum := new(Frustum)
	frustum.Xform = XformIdentity
	frustum.Near = near
	frustum.Far = far
	frustum.Fovy = fovy
	frustum.Aspect = aspect
	frustum.Update()
	return frustum
}

func (frustum *Frustum) Update() {
	rot := frustum.Rotation.Matrix()
	lookAt := rot.Mulv(V3{0, 0, -1}).Unit()
	angleY := frustum.Fovy * 0.5
	angleX := angleY * frustum.Aspect
	// Near
	frustum.Planes[0].Origin = frustum.Position.Addf(lookAt.Muls(frustum.Near))
	frustum.Planes[0].Normal = lookAt
	// Far
	frustum.Planes[1].Origin = frustum.Position.Addf(lookAt.Muls(frustum.Far))
	frustum.Planes[1].Normal = lookAt.Muls(-1)
	// Top
	frustum.Planes[2].Origin = frustum.Position
	frustum.Planes[2].Normal = rot.Mulv(V3{0, -math.Cos(angleY), -math.Sin(angleY)})
	// Bottom
	frustum.Planes[3].Origin = frustum.Position
	frustum.Planes[3].Normal = rot.Mulv(V3{0, math.Cos(angleY), -math.Sin(angleY)})
	// Left
	frustum.Planes[4].Origin = frustum.Position
	frustum.Planes[4].Normal = rot.Mulv(V3{math.Cos(angleX), 0, -math.Sin(angleX)})
	// Right
	frustum.Planes[5].Origin = frustum.Position
	frustum.Planes[5].Normal = rot.Mulv(V3{-math.Cos(angleX), 0, -math.Sin(angleX)})
}
