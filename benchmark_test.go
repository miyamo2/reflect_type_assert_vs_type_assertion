package reflect_type_assert_vs_type_assertion

import (
	"reflect"
	"testing"
	"time"
)

func BenchmarkTypeAssertion(b *testing.B) {
	rv := reflect.ValueOf(new(time.Time)).Elem()
	b.ReportAllocs()
	for b.Loop() {
		_, ok := rv.Interface().(time.Time)
		if !ok {
			b.Fatal("expected true but got false")
		}
	}
}

func BenchmarkReflectTypeAssert(b *testing.B) {
	rv := reflect.ValueOf(new(time.Time)).Elem()
	b.ReportAllocs()
	for b.Loop() {
		_, ok := reflect.TypeAssert[time.Time](rv)
		if !ok {
			b.Fatal("expected true but got false")
		}
	}
}

func BenchmarkTypeAssertionToPointer(b *testing.B) {
	rv := reflect.ValueOf(new(time.Time))
	b.ReportAllocs()
	for b.Loop() {
		_, ok := rv.Interface().(*time.Time)
		if !ok {
			b.Fatal("expected true but got false")
		}
	}
}

func BenchmarkReflectTypeAssertToPointer(b *testing.B) {
	rv := reflect.ValueOf(new(time.Time))
	b.ReportAllocs()
	for b.Loop() {
		_, ok := reflect.TypeAssert[*time.Time](rv)
		if !ok {
			b.Fatal("expected true but got false")
		}
	}
}

func BenchmarkTypeAssertionToInterface(b *testing.B) {
	rv := reflect.ValueOf(new(time.Time)).Elem()
	b.ReportAllocs()
	for b.Loop() {
		_, ok := rv.Interface().(ITime)
		if !ok {
			b.Fatal("expected true but got false")
		}
	}
}

func BenchmarkReflectTypeAssertToInterface(b *testing.B) {
	rv := reflect.ValueOf(new(time.Time)).Elem()
	b.ReportAllocs()
	for b.Loop() {
		_, ok := reflect.TypeAssert[ITime](rv)
		if !ok {
			b.Fatal("expected true but got false")
		}
	}
}
