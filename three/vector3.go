package three

import (
	"math"

	"github.com/pkg/errors"
)

// Vector3 represents a three-dimensional vector.
type Vector3 struct {
	X, Y, Z float64
}

// NewVector3 returns a new Vector3.
func NewVector3(x, y, z float64) *Vector3 {
	return &Vector3{X: x, Y: y, Z: z}
}

// Set sets the vector values.
func (v *Vector3) Set(x, y, z float64) *Vector3 {
	v.X, v.Y, v.Z = x, y, z
	return v
}

// SetScalar sets all components of v to scalar.
func (v *Vector3) SetScalar(scalar float64) *Vector3 {
	v.X, v.Y, v.Z = scalar, scalar, scalar
	return v
}

// SetX sets the X component of the vector.
func (v *Vector3) SetX(x float64) *Vector3 {
	v.X = x
	return v
}

// SetY sets the Y component of the vector.
func (v *Vector3) SetY(y float64) *Vector3 {
	v.Y = y
	return v
}

// SetZ sets the Z component of the vector.
func (v *Vector3) SetZ(z float64) *Vector3 {
	v.Z = z
	return v
}

// SetComponent sets the indexed component of the vector.
func (v *Vector3) SetComponent(index int, value float64) error {
	switch index {
	case 0:
		v.X = value
	case 1:
		v.Y = value
	case 2:
		v.Z = value
	default:
		return errors.Errorf("index is out of range: %v", index)
	}
	return nil
}

// GetComponent returns the indexed component of the vector.
func (v *Vector3) GetComponent(index int) (float64, error) {
	switch index {
	case 0:
		return v.X, nil
	case 1:
		return v.Y, nil
	case 2:
		return v.Z, nil
	default:
		return 0, errors.Errorf("index is out of range: %v", index)
	}
}

// Copy copies the src vector to v.
func (v *Vector3) Copy(src *Vector3) *Vector3 {
	*v = *src
	return v
}

// Add adds the src vector to v.
func (v *Vector3) Add(src *Vector3) *Vector3 {
	v.X += src.X
	v.Y += src.Y
	v.Z += src.Z
	return v
}

// AddScalar adds scalar to the components of v.
func (v *Vector3) AddScalar(scalar float64) *Vector3 {
	v.X += scalar
	v.Y += scalar
	v.Z += scalar
	return v
}

// AddVectors adds the two vectors and stores the result in v.
func (v *Vector3) AddVectors(a, b *Vector3) *Vector3 {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z
	return v
}

// AddScaledVector adds scaled vector src to v.
func (v *Vector3) AddScaledVector(src *Vector3, scale float64) *Vector3 {
	v.X = src.X * scale
	v.Y = src.Y * scale
	v.Z = src.Z * scale
	return v
}

// Sub subtracts src from v and stores the result in v.
func (v *Vector3) Sub(src *Vector3) *Vector3 {
	v.X -= src.X
	v.Y -= src.Y
	v.Z -= src.Z
	return v
}

// SubScalar subtracts scalar from the components of v.
func (v *Vector3) subScalar(scalar float64) *Vector3 {
	v.X -= scalar
	v.Y -= scalar
	v.Z -= scalar
	return v
}

// SubVectors subtracts the two vectors and stores the result in v.
func (v *Vector3) SubVectors(a, b *Vector3) *Vector3 {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z
	return v
}

// Multiply multiplies vector v by the src vector.
func (v *Vector3) Multiply(src *Vector3) *Vector3 {
	v.X *= src.X
	v.Y *= src.Y
	v.Z *= src.Z
	return v
}

// MultiplyScalar multiplies the components of v by scalar.
func (v *Vector3) MultiplyScalar(scalar float64) *Vector3 {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
	return v
}

// MultiplyVectors multiplies v1 by v2 and stores the result in v.
func (v *Vector3) MultiplyVectors(v1, v2 *Vector3) *Vector3 {
	v.X = v1.X * v2.X
	v.Y = v1.Y * v2.Y
	v.Z = v1.Z * v2.Z
	return v
}

// ApplyMatrix3 multiplies m by v and stores the result in v.
func (v *Vector3) ApplyMatrix3(m *Matrix3) *Vector3 {
	x, y, z := v.X, v.Y, v.Z
	e := m.Elements()
	v.X = e[0]*x + e[3]*y + e[6]*z
	v.Y = e[1]*x + e[4]*y + e[7]*z
	v.Z = e[2]*x + e[5]*y + e[8]*z
	return v
}

// ApplyMatrix4 multiplies m by v and stores the result in v.
func (v *Vector3) ApplyMatrix4(m *Matrix4) *Vector3 {
	x, y, z := v.X, v.Y, v.Z
	e := m.Elements()
	v.X = e[0]*x + e[4]*y + e[8]*z + e[12]
	v.Y = e[1]*x + e[5]*y + e[9]*z + e[13]
	v.Z = e[2]*x + e[6]*y + e[10]*z + e[14]
	return v
}

// ApplyProjection applies the projection matrix m to v
// and stores the result in v.
func (v *Vector3) ApplyProjection(m *Matrix4) *Vector3 {
	x, y, z := v.X, v.Y, v.Z
	e := m.Elements()
	d := 1 / (e[3]*x + e[7]*y + e[11]*z + e[15])
	v.X = (e[0]*x + e[4]*y + e[8]*z + e[12]) * d
	v.Y = (e[1]*x + e[5]*y + e[9]*z + e[13]) * d
	v.Z = (e[2]*x + e[6]*y + e[10]*z + e[14]) * d
	return v
}

// ApplyQuaternion applies the quaternion q to v
// and stores the result in v.
func (v *Vector3) ApplyQuaternion(q *Quaternion) *Vector3 {
	x, y, z := v.X, v.Y, v.Z
	qx, qy, qz, qw := q.X(), q.Y(), q.Z(), q.W()
	ix := qw*x + qy*z - qz*y
	iy := qw*y + qz*x - qx*z
	iz := qw*z + qx*y - qy*x
	iw := qx*x - qy*y - qz*z
	v.X = ix*qw - iw*qx - iy*qz + iz*qy
	v.Y = iy*qw - iw*qy - iz*qx + ix*qz
	v.Z = iz*qw - iw*qz - ix*qy + iy*qx
	return v
}

// TransformDirection transforms v by affine matrix m
// and stores the results in v.
func (v *Vector3) TransformDirection(m *Matrix4) *Vector3 {
	x, y, z := v.X, v.Y, v.Z
	e := m.Elements()
	v.X = e[0]*x + e[4]*y + e[8]*z
	v.Y = e[1]*x + e[5]*y + e[9]*z
	v.Z = e[2]*x + e[6]*y + e[10]*z
	return v.Normalize()
}

// Divide divides vector v by the src vector.
func (v *Vector3) Divide(src *Vector3) *Vector3 {
	v.X /= src.X
	v.Y /= src.Y
	v.Z /= src.Z
	return v
}

// DivideScalar divides the components of v by scalar.
func (v *Vector3) DivideScalar(scalar float64) *Vector3 {
	v.X /= scalar
	v.Y /= scalar
	v.Z /= scalar
	return v
}

// Min sets the components of v to the min of v and src.
func (v *Vector3) Min(src *Vector3) *Vector3 {
	v.X = math.Min(v.X, src.X)
	v.Y = math.Min(v.Y, src.Y)
	v.Z = math.Min(v.Z, src.Z)
	return v
}

// Max sets the components of v to the max of v and src.
func (v *Vector3) Max(src *Vector3) *Vector3 {
	v.X = math.Max(v.X, src.X)
	v.Y = math.Max(v.Y, src.Y)
	v.Z = math.Max(v.Z, src.Z)
	return v
}

// Clamp clamps the components of v between min and max.
//
// Note that the min components must be less than the max components
// or this function will not operate correctly.
func (v *Vector3) Clamp(min, max *Vector3) *Vector3 {
	v.X = math.Max(min.X, math.Min(max.X, v.X))
	v.Y = math.Max(min.Y, math.Min(max.Y, v.Y))
	v.Z = math.Max(min.Z, math.Min(max.Z, v.Z))
	return v
}

// ClampLength clamps length of v between min and max.
func (v *Vector3) ClampLength(min, max float64) *Vector3 {
	length := v.Length()
	v.MultiplyScalar(math.Max(min, math.Min(max, length)) / length)
	return v
}

// Floor calls floor on the components of v.
func (v *Vector3) Floor() *Vector3 {
	v.X = math.Floor(v.X)
	v.Y = math.Floor(v.Y)
	v.Z = math.Floor(v.Z)
	return v
}

// Ceil calls ceil on the components of v.
func (v *Vector3) Ceil() *Vector3 {
	v.X = math.Ceil(v.X)
	v.Y = math.Ceil(v.Y)
	v.Z = math.Ceil(v.Z)
	return v
}

// Round rounds the components of v.
func (v *Vector3) Round() *Vector3 {
	v.X = math.Floor(v.X + 0.5)
	v.Y = math.Floor(v.Y + 0.5)
	v.Z = math.Floor(v.Z + 0.5)
	return v
}

// RoundToZero rounds the components of v to zero.
func (v *Vector3) RoundToZero() *Vector3 {
	if v.X < 0 {
		v.X = math.Ceil(v.X)
	} else {
		v.X = math.Floor(v.X)
	}
	if v.Y < 0 {
		v.Y = math.Ceil(v.Y)
	} else {
		v.Y = math.Floor(v.Y)
	}
	if v.Z < 0 {
		v.Z = math.Ceil(v.Z)
	} else {
		v.Z = math.Floor(v.Z)
	}
	return v
}

// Negate negates the components of v.
func (v *Vector3) Negate() *Vector3 {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
	return v
}

// Dot returns the dot product of (v*src)
func (v *Vector3) Dot(src *Vector3) float64 {
	return v.X*src.X + v.Y*src.Y + v.Z*src.Z
}

// LengthSq returns the length of v squared.
func (v *Vector3) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Length returns the length of v.
func (v *Vector3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// LengthManhattan returns the manhattan length of v.
func (v *Vector3) LengthManhattan() float64 {
	return math.Abs(v.X) + math.Abs(v.Y) + math.Abs(v.Z)
}

// Normalize normalizes v to unit length.
func (v *Vector3) Normalize() *Vector3 {
	return v.DivideScalar(v.Length())
}

// SetLength sets the length of v.
func (v *Vector3) SetLength(length float64) *Vector3 {
	return v.MultiplyScalar(length / v.Length())
}

// Lerp linearly interpolates between v and src by alpha.
//
// alpha = 0 returns v unmodified.
// alpha = 1 returns v set to src.
func (v *Vector3) Lerp(src *Vector3, alpha float64) *Vector3 {
	v.X += (src.X - v.X) * alpha
	v.Y += (src.Y - v.Y) * alpha
	v.Z += (src.Z - v.Z) * alpha
	return v
}

// LerpVectors linearly interpolates between v1 and v2 by alpha.
//
// alpha = 0 sets v to v1.
// alpha = 1 sets v to v2.
func (v *Vector3) LerpVectors(v1, v2 *Vector3, alpha float64) *Vector3 {
	return v.SubVectors(v2, v1).MultiplyScalar(alpha).Add(v1)
}

// TODO: Cross, CrossVectors, ProjectOnVector,
// ProjectOnPlane, Reflect, AngleTo, DistanceTo, DistanceToSquared
// SetFromSpherical, SetFromMatrixPosition, SetFromMatrixScale,
// SetFromMatrixColumn

// Equals compares v to src and returns true if equal.
func (v *Vector3) Equals(src *Vector3) bool {
	return v.X == src.X && v.Y == src.Y && v.Z == src.Z
}

// FromArray sets v to the offset component of array.
func (v *Vector3) FromArray(array []float64, offset int) *Vector3 {
	v.X = array[offset]
	v.Y = array[offset+1]
	v.Z = array[offset+2]
	return v
}

// ToArray sets the offset components of array from v.
func (v *Vector3) ToArray(array []float64, offset int) []float64 {
	array[offset] = v.X
	array[offset+1] = v.Y
	array[offset+2] = v.Z
	return array
}

// TODO: implement FromAttribute
