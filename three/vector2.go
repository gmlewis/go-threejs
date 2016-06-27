package three

import (
	"math"

	"github.com/pkg/errors"
)

// Vector2 represents a two-dimensional vector.
type Vector2 struct {
	X, Y float64
}

// Width returns the x component of v.
func (v *Vector2) Width() float64 {
	return v.X
}

// Height returns the Y component of v.
func (v *Vector2) Height() float64 {
	return v.Y
}

// Set sets the vector values.
func (v *Vector2) Set(x, y float64) *Vector2 {
	v.X, v.Y = x, y
	return v
}

// SetScalar sets both the X and Y dimensions of the vector.
func (v *Vector2) SetScalar(scalar float64) *Vector2 {
	v.X, v.Y = scalar, scalar
	return v
}

// SetX sets the X component of the vector.
func (v *Vector2) SetX(x float64) *Vector2 {
	v.X = x
	return v
}

// SetY sets the Y component of the vector.
func (v *Vector2) SetY(y float64) *Vector2 {
	v.Y = y
	return v
}

// SetComponent sets the indexed component of the vector.
func (v *Vector2) SetComponent(index int, value float64) error {
	switch index {
	case 0:
		v.X = value
	case 1:
		v.Y = value
	default:
		return errors.Errorf("index is out of range: %v", index)
	}
	return nil
}

// GetComponent returns the indexed component of the vector.
func (v *Vector2) GetComponent(index int) (float64, error) {
	switch index {
	case 0:
		return v.X, nil
	case 1:
		return v.Y, nil
	default:
		return 0, errors.Errorf("index is out of range: %v", index)
	}
}

// Copy copies the src vector to v.
func (v *Vector2) Copy(src *Vector2) *Vector2 {
	v.X, v.Y = src.X, src.Y
	return v
}

// Add adds the src vector to v.
func (v *Vector2) Add(src *Vector2) *Vector2 {
	v.X += src.X
	v.Y += src.Y
	return v
}

// AddScalar adds scalar to the components of v.
func (v *Vector2) AddScalar(scalar float64) *Vector2 {
	v.X += scalar
	v.Y += scalar
	return v
}

// AddVectors adds the two vectors and stores the result in v.
func (v *Vector2) AddVectors(a, b *Vector2) *Vector2 {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	return v
}

// AddScaledVector adds scaled vector src to v.
func (v *Vector2) AddScaledVector(src *Vector2, scale float64) *Vector2 {
	v.X = src.X * scale
	v.Y = src.Y * scale
	return v
}

// Sub subtracts src from v and stores the result in v.
func (v *Vector2) Sub(src *Vector2) *Vector2 {
	v.X -= src.X
	v.Y -= src.Y
	return v
}

// SubScalar subtracts scalar from the components of v.
func (v *Vector2) subScalar(scalar float64) *Vector2 {
	v.X -= scalar
	v.Y -= scalar
	return v
}

// SubVectors subtracts the two vectors and stores the result in v.
func (v *Vector2) SubVectors(a, b *Vector2) *Vector2 {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	return v
}

// Multiply multiplies vector v by the src vector.
func (v *Vector2) Multiply(src *Vector2) *Vector2 {
	v.X *= src.X
	v.Y *= src.Y
	return v
}

// MultiplyScalar multiplies the components of v by scalar.
func (v *Vector2) MultiplyScalar(scalar float64) *Vector2 {
	v.X *= scalar
	v.Y *= scalar
	return v
}

// Divide divides vector v by the src vector.
func (v *Vector2) Divide(src *Vector2) *Vector2 {
	v.X /= src.X
	v.Y /= src.Y
	return v
}

// DivideScalar divides the components of v by scalar.
func (v *Vector2) DivideScalar(scalar float64) *Vector2 {
	v.X /= scalar
	v.Y /= scalar
	return v
}

// Min sets the components of v to the min of v and src.
func (v *Vector2) Min(src *Vector2) *Vector2 {
	v.X = math.Min(v.X, src.X)
	v.Y = math.Min(v.Y, src.Y)
	return v
}

// Max sets the components of v to the max of v and src.
func (v *Vector2) Max(src *Vector2) *Vector2 {
	v.X = math.Max(v.X, src.X)
	v.Y = math.Max(v.Y, src.Y)
	return v
}

// Clamp clamps the components of v between min and max.
//
// Note that the min components must be less than the max components
// or this function will not operate correctly.
func (v *Vector2) Clamp(min, max *Vector2) *Vector2 {
	v.X = math.Max(min.X, math.Min(max.X, v.X))
	v.Y = math.Max(min.Y, math.Min(max.Y, v.Y))
	return v
}

// ClampLength clamps length of v between min and max.
func (v *Vector2) ClampLength(min, max float64) *Vector2 {
	length := v.Length()
	v.MultiplyScalar(math.Max(min, math.Min(max, length)) / length)
	return v
}

// Floor calls floor on the components of v.
func (v *Vector2) Floor() *Vector2 {
	v.X = math.Floor(v.X)
	v.Y = math.Floor(v.Y)
	return v
}

// Ceil calls ceil on the components of v.
func (v *Vector2) Ceil() *Vector2 {
	v.X = math.Ceil(v.X)
	v.Y = math.Ceil(v.Y)
	return v
}

// Round rounds the components of v.
func (v *Vector2) Round() *Vector2 {
	v.X = math.Floor(v.X + 0.5)
	v.Y = math.Floor(v.Y + 0.5)
	return v
}

// RoundToZero rounds the components of v to zero.
func (v *Vector2) RoundToZero() *Vector2 {
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
	return v
}

// Negate negates the components of v.
func (v *Vector2) Negate() *Vector2 {
	v.X = -v.X
	v.Y = -v.Y
	return v
}

// Dot returns the dot product of (v*src)
func (v *Vector2) Dot(src *Vector2) float64 {
	return v.X*src.X + v.Y*src.Y
}

// LengthSq returns the length of v squared.
func (v *Vector2) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Length returns the length of v.
func (v *Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// LengthManhattan returns the manhattan length of v.
func (v *Vector2) LengthManhattan() float64 {
	return math.Abs(v.X) + math.Abs(v.Y)
}

// Normalize normalizes v to unit length.
func (v *Vector2) Normalize() *Vector2 {
	return v.DivideScalar(v.Length())
}

// Angle returns the angle in radians with respect to the positive X-axis.
func (v *Vector2) Angle() float64 {
	angle := math.Atan2(v.Y, v.X)
	if angle < 0 {
		angle += 2 * math.Pi
	}
	return angle
}

// DistanceTo returns the distance from v to src.
func (v *Vector2) DistanceTo(src *Vector2) float64 {
	return math.Sqrt(v.DistanceToSquared(src))
}

// DistanceToSquared returns the squared distance from v to src.
func (v *Vector2) DistanceToSquared(src *Vector2) float64 {
	dx, dy := v.X-src.X, v.Y-src.Y
	return dx*dx + dy*dy
}

// SetLength sets the length of v.
func (v *Vector2) SetLength(length float64) *Vector2 {
	return v.MultiplyScalar(length / v.Length())
}

// Lerp linearly interpolates between v and src by alpha.
//
// alpha = 0 returns v unmodified.
// alpha = 1 returns v set to src.
func (v *Vector2) Lerp(src *Vector2, alpha float64) *Vector2 {
	v.X += (src.X - v.X) * alpha
	v.Y += (src.Y - v.Y) * alpha
	return v
}

// LerpVectors linearly interpolates between v1 and v2 by alpha.
//
// alpha = 0 sets v to v1.
// alpha = 1 sets v to v2.
func (v *Vector2) LerpVectors(v1, v2 *Vector2, alpha float64) *Vector2 {
	return v.SubVectors(v2, v1).MultiplyScalar(alpha).Add(v1)
}

// Equals compares v to src and returns true if equal.
func (v *Vector2) Equals(src *Vector2) bool {
	return v.X == src.X && v.Y == src.Y
}

// FromArray sets v to the offset component of array.
func (v *Vector2) FromArray(array []float64, offset int) *Vector2 {
	v.X = array[offset]
	v.Y = array[offset+1]
	return v
}

// ToArray sets the offset components of array from v.
func (v *Vector2) ToArray(array []float64, offset int) []float64 {
	array[offset] = v.X
	array[offset+1] = v.Y
	return array
}

// TODO: implement FromAttribute

// RotateAround rotates v around center by angle radians.
func (v *Vector2) RotateAround(center *Vector2, angle float64) *Vector2 {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	x := v.X - center.X
	y := v.Y - center.Y
	v.X = x*cos - y*sin + center.X
	v.Y = x*sin + y*cos + center.Y
	return v
}
