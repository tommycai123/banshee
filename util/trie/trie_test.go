// Copyright 2016 Eleme Inc. All rights reserved.

package trie

import (
	"github.com/eleme/banshee/util"
	"math/rand"
	"testing"
)

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// randKey returns a random trie key with given number of segments.
func randKey(n int) string {
	if n == 0 {
		n = 1 // force n >= 1
	}
	b := make([]byte, 2*n-1)
	for i := 0; i < n; i++ {
		j := rand.Intn(len(letters))
		b[i] = letters[j]
		if i+1 < n {
			b[i+1] = '.'
		}
	}
	return string(b)
}

func TestPut(t *testing.T) {
	tr := New()
	// Case simple
	tr.Put("a.b.c.d", 4)
	tr.Put("a.b.c.d", 99) // case reset
	tr.Put("a.b.c.d.e", 5)
	tr.Put("a.b.c.d.e.f", 6)
	tr.Put("a.b.c.d.e.f.g", 7)
	tr.Put("a.b.c.d.e.f.g.h", 8)
	util.Must(t, tr.Len() == 5)
	util.Must(t, tr.Get("a.b.c.d").(int) == 99)
	util.Must(t, tr.Get("a.b.c.d.e").(int) == 5)
	util.Must(t, tr.Get("a.b.c.d.e.f").(int) == 6)
	util.Must(t, tr.Get("a.b.c.d.e.f.g").(int) == 7)
	util.Must(t, tr.Get("a.b.c.d.e.f.g.h").(int) == 8)
	// Case larger number.
	n := 1024 * 5
	for i := 0; i < n; i++ {
		key := randKey(rand.Intn(128))
		tr.Put(key, i)
		util.Must(t, tr.Get(key).(int) == i)
	}
}

func TestGet(t *testing.T) {
	tr := New()
	// Case not found.
	util.Must(t, tr.Get("not.exist") == nil)
	// Case simple.
	tr.Put("a.b.c.d", 43)
	tr.Put("b.c.d.a", 34)
	tr.Put("m.n.o.p.q", 52)
	util.Must(t, tr.Get("a.b.c.d").(int) == 43)
	util.Must(t, tr.Get("b.c.d.a").(int) == 34)
	util.Must(t, tr.Get("m.n.o.p.q").(int) == 52)
	util.Must(t, tr.Get("a.b.c") == nil)
	// Case Has.
	util.Must(t, tr.Has("a.b.c.d"))
	util.Must(t, !tr.Has("a.b.c.d.e"))
}

func TestPop(t *testing.T) {
	tr := New()
	// Case not found.
	util.Must(t, tr.Pop("not.exist") == nil)
	util.Must(t, tr.Len() == 0)
	// Case simple.
	tr.Put("a.b.c.d", 4)
	tr.Put("a.b.c.d.e", 5)
	tr.Put("a.b.c.d.e.f", 6)
	util.Must(t, tr.Len() == 3)
	util.Must(t, tr.Pop("a.b.c") == nil)
	util.Must(t, tr.Len() == 3)
	util.Must(t, tr.Pop("a.b.c.d").(int) == 4)
	util.Must(t, tr.Len() == 2)
	util.Must(t, tr.Pop("a.b.c.d.e").(int) == 5)
	util.Must(t, tr.Len() == 1)
	util.Must(t, tr.Pop("a.b.c.d.e.f").(int) == 6)
	util.Must(t, tr.Len() == 0)
}

func TestClear(t *testing.T) {
	tr := New()
	// Case simple.
	tr.Put("a.b.c.d", 4)
	tr.Put("a.b.c.d.e", 5)
	tr.Put("a.b.c.d.e.f", 6)
	util.Must(t, tr.Len() == 3)
	tr.Clear()
	util.Must(t, tr.Len() == 0)
	util.Must(t, !tr.Has("a.b.c.d"))
}

func TestMatch(t *testing.T) {
	tr := New()
	// Case simple.
	tr.Put("a.b.c.d", 4)
	tr.Put("a.b.c.f", 9)
	tr.Put("a.b.c.d.e", 5)
	tr.Put("a.b.c.d.e.f", 6)
	tr.Put("m.n.o.p", 43)
	tr.Put("m.n.o.p.q", 53)
	tr.Put("m.n.o.p.q.r", 63)
	var m map[string]interface{}
	// Case x.*
	m = tr.Match("a.b.*.*")
	util.Must(t, len(m) == 2)
	util.Must(t, m["a.b.c.d"].(int) == 4)
	util.Must(t, m["a.b.c.f"].(int) == 9)
	util.Must(t, tr.NumMatch("a.b.*.*") == len(m))
	// Case x
	m = tr.Match("a.b.c.d")
	util.Must(t, len(m) == 1)
	util.Must(t, m["a.b.c.d"].(int) == 4)
	util.Must(t, tr.NumMatch("a.b.c.d") == len(m))
	// Case ""
	m = tr.Match("")
	util.Must(t, len(m) == 0)
	util.Must(t, tr.NumMatch("") == len(m))
	// Case *.x
	m = tr.Match("*.n.o.p")
	util.Must(t, len(m) == 1)
	util.Must(t, m["m.n.o.p"].(int) == 43)
	util.Must(t, tr.NumMatch("*.n.o.p") == len(m))
	// Case *.*
	m = tr.Match("*.b.c.*")
	util.Must(t, len(m) == 2)
	util.Must(t, m["a.b.c.d"].(int) == 4)
	util.Must(t, m["a.b.c.f"].(int) == 9)
	util.Must(t, tr.NumMatch("*.b.c.*") == len(m))
	// Case *...*
	m = tr.Match("*.*.*.*")
	util.Must(t, len(m) == 3)
	util.Must(t, tr.NumMatch("*.*.*.*") == len(m))
	// Case x.*.x
	m = tr.Match("a.*.*.d")
	util.Must(t, len(m) == 1)
	util.Must(t, tr.NumMatch("a.*.*.d") == len(m))
}

func TestMap(t *testing.T) {
	tr := New()
	// Case empty.
	util.Must(t, len(tr.Map()) == 0)
	// Case simple.
	tr.Put("a.b.c.d", 41)
	tr.Put("a.b.c.d.e", 51)
	tr.Put("a.b.c.d.e.f", 61)
	m := tr.Map()
	util.Must(t, len(m) == 3)
	util.Must(t, m["a.b.c.d"].(int) == 41)
	util.Must(t, m["a.b.c.d.e"].(int) == 51)
	util.Must(t, m["a.b.c.d.e.f"].(int) == 61)
}

func TestMatched(t *testing.T) {
	tr := New()

	p1 := "a.b.c.d.e.f"
	p2 := "a.b.c.d.*.f"
	p3 := "*.b.c.d.*.f"
	p4 := "*.b.c.d.*"
	p5 := "*"
	p6 := ""

	tr.Put(p1, 1)
	tr.Put(p2, 2)
	tr.Put(p3, 3)
	tr.Put(p4, 4)
	tr.Put(p5, 5)
	tr.Put(p6, 6)

	var m map[string]interface{}

	m = tr.Matched("a.b.c.d.e.f")
	util.Must(t, len(m) == 3)
	util.Must(t, m[p1].(int) == 1)
	util.Must(t, m[p2].(int) == 2)
	util.Must(t, m[p3].(int) == 3)

	m = tr.Matched("a.b.c.d.e")
	util.Must(t, len(m) == 1)
	util.Must(t, m[p4].(int) == 4)

	m = tr.Matched("a")
	util.Must(t, len(m) == 1)
	util.Must(t, m[p5].(int) == 5)

	m = tr.Matched("not.matched")
	util.Must(t, len(m) == 0)
}

func BenchmarkPutRandKeys(b *testing.B) {
	tr := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Put(randKey(128), i)
	}
}

func BenchmarkPutPrefixedKeys(b *testing.B) {
	tr := New()
	m := 63
	n := 16
	prefix := randKey(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%n == 0 {
			prefix = randKey(m)
		}
		key := prefix + "." + randKey(63)
		tr.Put(key, i)
	}
}
func BenchmarkPutAndGetRandKeys(b *testing.B) {
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Put(randKey(128), i)
		tr.Get(randKey(128))
	}
}

func BenchmarkPutAndGetPrefixedKeys(b *testing.B) {
	tr := New()
	m := 63
	n := 16
	prefix := randKey(m)
	for i := 0; i < b.N; i++ {
		if i%n == 0 {
			prefix = randKey(m)
		}
		key := prefix + "." + randKey(63)
		tr.Put(key, i)
		tr.Get(key)
	}
}
