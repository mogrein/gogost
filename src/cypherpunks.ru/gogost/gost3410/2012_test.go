// GoGOST -- Pure Go GOST cryptographic functions library
// Copyright (C) 2015-2019 Sergey Matveev <stargrave@stargrave.org>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package gost3410

import (
	"bytes"
	"crypto/rand"
	"testing"
	"testing/quick"
)

func TestGCL3Vectors(t *testing.T) {
	p := []byte{
		0x45, 0x31, 0xAC, 0xD1, 0xFE, 0x00, 0x23, 0xC7,
		0x55, 0x0D, 0x26, 0x7B, 0x6B, 0x2F, 0xEE, 0x80,
		0x92, 0x2B, 0x14, 0xB2, 0xFF, 0xB9, 0x0F, 0x04,
		0xD4, 0xEB, 0x7C, 0x09, 0xB5, 0xD2, 0xD1, 0x5D,
		0xF1, 0xD8, 0x52, 0x74, 0x1A, 0xF4, 0x70, 0x4A,
		0x04, 0x58, 0x04, 0x7E, 0x80, 0xE4, 0x54, 0x6D,
		0x35, 0xB8, 0x33, 0x6F, 0xAC, 0x22, 0x4D, 0xD8,
		0x16, 0x64, 0xBB, 0xF5, 0x28, 0xBE, 0x63, 0x73,
	}
	q := []byte{
		0x45, 0x31, 0xAC, 0xD1, 0xFE, 0x00, 0x23, 0xC7,
		0x55, 0x0D, 0x26, 0x7B, 0x6B, 0x2F, 0xEE, 0x80,
		0x92, 0x2B, 0x14, 0xB2, 0xFF, 0xB9, 0x0F, 0x04,
		0xD4, 0xEB, 0x7C, 0x09, 0xB5, 0xD2, 0xD1, 0x5D,
		0xA8, 0x2F, 0x2D, 0x7E, 0xCB, 0x1D, 0xBA, 0xC7,
		0x19, 0x90, 0x5C, 0x5E, 0xEC, 0xC4, 0x23, 0xF1,
		0xD8, 0x6E, 0x25, 0xED, 0xBE, 0x23, 0xC5, 0x95,
		0xD6, 0x44, 0xAA, 0xF1, 0x87, 0xE6, 0xE6, 0xDF,
	}
	a := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07,
	}
	b := []byte{
		0x1C, 0xFF, 0x08, 0x06, 0xA3, 0x11, 0x16, 0xDA,
		0x29, 0xD8, 0xCF, 0xA5, 0x4E, 0x57, 0xEB, 0x74,
		0x8B, 0xC5, 0xF3, 0x77, 0xE4, 0x94, 0x00, 0xFD,
		0xD7, 0x88, 0xB6, 0x49, 0xEC, 0xA1, 0xAC, 0x43,
		0x61, 0x83, 0x40, 0x13, 0xB2, 0xAD, 0x73, 0x22,
		0x48, 0x0A, 0x89, 0xCA, 0x58, 0xE0, 0xCF, 0x74,
		0xBC, 0x9E, 0x54, 0x0C, 0x2A, 0xDD, 0x68, 0x97,
		0xFA, 0xD0, 0xA3, 0x08, 0x4F, 0x30, 0x2A, 0xDC,
	}
	x := []byte{
		0x24, 0xD1, 0x9C, 0xC6, 0x45, 0x72, 0xEE, 0x30,
		0xF3, 0x96, 0xBF, 0x6E, 0xBB, 0xFD, 0x7A, 0x6C,
		0x52, 0x13, 0xB3, 0xB3, 0xD7, 0x05, 0x7C, 0xC8,
		0x25, 0xF9, 0x10, 0x93, 0xA6, 0x8C, 0xD7, 0x62,
		0xFD, 0x60, 0x61, 0x12, 0x62, 0xCD, 0x83, 0x8D,
		0xC6, 0xB6, 0x0A, 0xA7, 0xEE, 0xE8, 0x04, 0xE2,
		0x8B, 0xC8, 0x49, 0x97, 0x7F, 0xAC, 0x33, 0xB4,
		0xB5, 0x30, 0xF1, 0xB1, 0x20, 0x24, 0x8A, 0x9A,
	}
	y := []byte{
		0x2B, 0xB3, 0x12, 0xA4, 0x3B, 0xD2, 0xCE, 0x6E,
		0x0D, 0x02, 0x06, 0x13, 0xC8, 0x57, 0xAC, 0xDD,
		0xCF, 0xBF, 0x06, 0x1E, 0x91, 0xE5, 0xF2, 0xC3,
		0xF3, 0x24, 0x47, 0xC2, 0x59, 0xF3, 0x9B, 0x2C,
		0x83, 0xAB, 0x15, 0x6D, 0x77, 0xF1, 0x49, 0x6B,
		0xF7, 0xEB, 0x33, 0x51, 0xE1, 0xEE, 0x4E, 0x43,
		0xDC, 0x1A, 0x18, 0xB9, 0x1B, 0x24, 0x64, 0x0B,
		0x6D, 0xBB, 0x92, 0xCB, 0x1A, 0xDD, 0x37, 0x1E,
	}
	priv := []byte{
		0xD4, 0x8D, 0xA1, 0x1F, 0x82, 0x67, 0x29, 0xC6,
		0xDF, 0xAA, 0x18, 0xFD, 0x7B, 0x6B, 0x63, 0xA2,
		0x14, 0x27, 0x7E, 0x82, 0xD2, 0xDA, 0x22, 0x33,
		0x56, 0xA0, 0x00, 0x22, 0x3B, 0x12, 0xE8, 0x72,
		0x20, 0x10, 0x8B, 0x50, 0x8E, 0x50, 0xE7, 0x0E,
		0x70, 0x69, 0x46, 0x51, 0xE8, 0xA0, 0x91, 0x30,
		0xC9, 0xD7, 0x56, 0x77, 0xD4, 0x36, 0x09, 0xA4,
		0x1B, 0x24, 0xAE, 0xAD, 0x8A, 0x04, 0xA6, 0x0B,
	}
	pubX := []byte{
		0xE1, 0xEF, 0x30, 0xD5, 0x2C, 0x61, 0x33, 0xDD,
		0xD9, 0x9D, 0x1D, 0x5C, 0x41, 0x45, 0x5C, 0xF7,
		0xDF, 0x4D, 0x8B, 0x4C, 0x92, 0x5B, 0xBC, 0x69,
		0xAF, 0x14, 0x33, 0xD1, 0x56, 0x58, 0x51, 0x5A,
		0xDD, 0x21, 0x46, 0x85, 0x0C, 0x32, 0x5C, 0x5B,
		0x81, 0xC1, 0x33, 0xBE, 0x65, 0x5A, 0xA8, 0xC4,
		0xD4, 0x40, 0xE7, 0xB9, 0x8A, 0x8D, 0x59, 0x48,
		0x7B, 0x0C, 0x76, 0x96, 0xBC, 0xC5, 0x5D, 0x11,
	}
	pubY := []byte{
		0xEC, 0xBE, 0x77, 0x36, 0xA9, 0xEC, 0x35, 0x7F,
		0xF2, 0xFD, 0x39, 0x93, 0x1F, 0x4E, 0x11, 0x4C,
		0xB8, 0xCD, 0xA3, 0x59, 0x27, 0x0A, 0xC7, 0xF0,
		0xE7, 0xFF, 0x43, 0xD9, 0x41, 0x94, 0x19, 0xEA,
		0x61, 0xFD, 0x2A, 0xB7, 0x7F, 0x5D, 0x9F, 0x63,
		0x52, 0x3D, 0x3B, 0x50, 0xA0, 0x4F, 0x63, 0xE2,
		0xA0, 0xCF, 0x51, 0xB7, 0xC1, 0x3A, 0xDC, 0x21,
		0x56, 0x0F, 0x0B, 0xD4, 0x0C, 0xC9, 0xC7, 0x37,
	}
	digest := []byte{
		0x37, 0x54, 0xF3, 0xCF, 0xAC, 0xC9, 0xE0, 0x61,
		0x5C, 0x4F, 0x4A, 0x7C, 0x4D, 0x8D, 0xAB, 0x53,
		0x1B, 0x09, 0xB6, 0xF9, 0xC1, 0x70, 0xC5, 0x33,
		0xA7, 0x1D, 0x14, 0x70, 0x35, 0xB0, 0xC5, 0x91,
		0x71, 0x84, 0xEE, 0x53, 0x65, 0x93, 0xF4, 0x41,
		0x43, 0x39, 0x97, 0x6C, 0x64, 0x7C, 0x5D, 0x5A,
		0x40, 0x7A, 0xDE, 0xDB, 0x1D, 0x56, 0x0C, 0x4F,
		0xC6, 0x77, 0x7D, 0x29, 0x72, 0x07, 0x5B, 0x8C,
	}
	signature := []byte{
		0x10, 0x81, 0xB3, 0x94, 0x69, 0x6F, 0xFE, 0x8E,
		0x65, 0x85, 0xE7, 0xA9, 0x36, 0x2D, 0x26, 0xB6,
		0x32, 0x5F, 0x56, 0x77, 0x8A, 0xAD, 0xBC, 0x08,
		0x1C, 0x0B, 0xFB, 0xE9, 0x33, 0xD5, 0x2F, 0xF5,
		0x82, 0x3C, 0xE2, 0x88, 0xE8, 0xC4, 0xF3, 0x62,
		0x52, 0x60, 0x80, 0xDF, 0x7F, 0x70, 0xCE, 0x40,
		0x6A, 0x6E, 0xEB, 0x1F, 0x56, 0x91, 0x9C, 0xB9,
		0x2A, 0x98, 0x53, 0xBD, 0xE7, 0x3E, 0x5B, 0x4A,
		0x2F, 0x86, 0xFA, 0x60, 0xA0, 0x81, 0x09, 0x1A,
		0x23, 0xDD, 0x79, 0x5E, 0x1E, 0x3C, 0x68, 0x9E,
		0xE5, 0x12, 0xA3, 0xC8, 0x2E, 0xE0, 0xDC, 0xC2,
		0x64, 0x3C, 0x78, 0xEE, 0xA8, 0xFC, 0xAC, 0xD3,
		0x54, 0x92, 0x55, 0x84, 0x86, 0xB2, 0x0F, 0x1C,
		0x9E, 0xC1, 0x97, 0xC9, 0x06, 0x99, 0x85, 0x02,
		0x60, 0xC9, 0x3B, 0xCB, 0xCD, 0x9C, 0x5C, 0x33,
		0x17, 0xE1, 0x93, 0x44, 0xE1, 0x73, 0xAE, 0x36,
	}
	c, err := NewCurve(p, q, a, b, x, y)
	if err != nil {
		t.FailNow()
	}
	prv, err := NewPrivateKey(c, Mode2012, priv)
	if err != nil {
		t.FailNow()
	}
	pub, err := prv.PublicKey()
	if err != nil {
		t.FailNow()
	}
	if bytes.Compare(pub.Raw()[:64], pubX) != 0 {
		t.FailNow()
	}
	if bytes.Compare(pub.Raw()[64:], pubY) != 0 {
		t.FailNow()
	}
	ourSign, err := prv.SignDigest(digest, rand.Reader)
	if err != nil {
		t.FailNow()
	}
	valid, err := pub.VerifyDigest(digest, ourSign)
	if err != nil || !valid {
		t.FailNow()
	}
	valid, err = pub.VerifyDigest(digest, signature)
	if err != nil || !valid {
		t.FailNow()
	}
}

func TestRandom2012(t *testing.T) {
	c := CurveIdtc26gost341012512paramSetA()
	f := func(data [31]byte, digest [64]byte) bool {
		prv, err := NewPrivateKey(
			c,
			Mode2012,
			append([]byte{0xde}, data[:]...),
		)
		if err != nil {
			return false
		}
		pub, err := prv.PublicKey()
		if err != nil {
			return false
		}
		pubRaw := pub.Raw()
		pub, err = NewPublicKey(c, Mode2012, pubRaw)
		if err != nil {
			return false
		}
		sign, err := prv.SignDigest(digest[:], rand.Reader)
		if err != nil {
			return false
		}
		valid, err := pub.VerifyDigest(digest[:], sign)
		if err != nil {
			return false
		}
		return valid
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func BenchmarkSign2012(b *testing.B) {
	c := CurveIdtc26gost341012512paramSetA()
	prv, err := GenPrivateKey(c, Mode2012, rand.Reader)
	if err != nil {
		b.FailNow()
	}
	digest := make([]byte, 64)
	rand.Read(digest)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prv.SignDigest(digest, rand.Reader)
	}
}

func BenchmarkVerify2012(b *testing.B) {
	c := CurveIdtc26gost341012512paramSetA()
	prv, err := GenPrivateKey(c, Mode2012, rand.Reader)
	if err != nil {
		b.FailNow()
	}
	digest := make([]byte, 64)
	rand.Read(digest)
	sign, err := prv.SignDigest(digest, rand.Reader)
	if err != nil {
		b.FailNow()
	}
	pub, err := prv.PublicKey()
	if err != nil {
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pub.VerifyDigest(digest, sign)
	}
}
