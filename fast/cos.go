/*
Copyright 2013 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fast

import "math"

// Fast Cosine approximation with table lookup and linear interpolation.
func Cos(x float64) float64 {
	f := x * cosFactor
	if x < 0 {
		f *= -1
	}
	t := int(f)
	i := t & (cosLen - 1)
	res := cosTable[i] + cosGrad[i]*(f-float64(t))
	if x < 0 {
		return res * -1
	}
	return res
}

const (
	cosLen    = 512
	cosFactor = cosLen / (2 * math.Pi)
)

var cosTable, cosGrad []float64

func init() {
	cosTable = make([]float64, cosLen)
	cosGrad = make([]float64, cosLen)
	step := 1 / cosFactor
	for i := 0; i < cosLen; i++ {
		cosTable[i] = math.Cos(float64(i) * step)
	}
	for i := 0; i < cosLen; i++ {
		cosGrad[i] = cosTable[(i+1)%cosLen] - cosTable[i]
	}
}
