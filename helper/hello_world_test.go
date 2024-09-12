package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// wajib diawalin dengan nama Test kemudian methodnya tidak mereturn sebuah value
// dan biasanya nama filenya adalah namaFileYangDitest_test.go
// dan parameternya wajib t *testing.T untuk unit test di golang

/*
kita bisa menggunakan go test untuk menjalankan unit test, tapi gakeliatan function apa aja yang di test
kita bisa pake go test -v untuk tau function apa aja yang di test
kita bisa pake go test -v -run=TestHelloWorldSyifa atau go test -v -run TestHelloWorldSyifa
(tapi ini prefix awal, jadi kalau TestHelloWorld doang maka auto ke test method yang TestHelloWorld dan TestHelloWorldSyifa)
untuk baca di root directory maka pake go test ./... (maka otomatis akan test semua package yg ada test nya)
*/
func TestHelloWorldIlham(t *testing.T) {
	result := HelloWorld("Ilham")
	if result != "Hello Ilham" {
		// maka kode program dibawahnya yaitu line 26 tetap dieksekusi, karena si t.Error manggil t.Fail()
		t.Error("Result must be 'Hello Ilham'")
	}

	fmt.Println("TestHelloWorldIlham Done")
}

func TestHelloWorldSyifa(t *testing.T) {
	result := HelloWorld("Syifa")
	if result != "Hello Syifa" {
		// maka kode program dibawahnya yaitu line 36 tidak akan dieksekusi, karena t.Fatal manggil t.FailNow()
		t.Fatal("Result must be 'Hello Syifa'")
	}

	fmt.Println("TestHelloWorldSyifa Done")
}

/*
	terdapat method fail(), failNow(), Error(), dan Fatal() jika kita ingin membatalkan unit test
*/

func TestHelloWorldWithAssert(t *testing.T) {
	result := HelloWorld("Hamzah")
	assert.Equal(t, "Hello Hamzah", result, "Result must be Hello Hamzah") // ini manggil yang t.Fail()
	fmt.Println("TestHelloWorldWithAssert done")
}

func TestHelloWorldWithRequire(t *testing.T) {
	result := HelloWorld("Silmi")
	require.Equal(t, "Hello Silmi", result, "Result must be Hello Silmi") // ini manggil yang t.FailNow()
	fmt.Println("TestHelloWorldWithRequire done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("gabisa jalan di windows")
	}

	result := HelloWorld("Hamzah")
	assert.Equal(t, "Hello Hamzah", result, "Result must be Hello Hamzah")
}

// fitur main yang nama functionya wajib TestMain
// ini hanya akan dieksekusi sekali di package, bukan pertiap function unit test
// jadi kalau beda package wajib bikin TestMain lagi
func TestMain(m *testing.M) {
	// before unit test
	fmt.Println("Before unit test")

	m.Run()

	fmt.Println("After unit test")
}

// kalau mau test subTestnya doang pake go test -v -run=TestSubTest/NamaSubTest
func TestSubTest(t *testing.T) {
	t.Run("Ilham", func(t *testing.T) { // subTest dengan nama Ilham
		result := HelloWorld("Ilham")
		assert.Equal(t, "Hello Ilham", result, "Result must be Hello Ilham")
	})

	t.Run("Sidiq", func(t *testing.T) { // subTest dengan nama Sidiq
		result := HelloWorld("Sidiq")
		assert.Equal(t, "Hello Sidiq", result, "Result must be Hello Sidiq")
	})
}

// TableTest
func TestTableHelloWorld(t *testing.T) {
	tests := []struct { // ini adalah anonymous struct
		name     string
		request  string
		expected string
	}{
		{
			"Ilham",
			"Ilham",
			"Hello Ilham",
		},
		{
			"Muhammad",
			"Muhammad",
			"Hello Muhammad",
		},
		{
			"Sidiq",
			"Sidiq",
			"Hello Sidiq",
		},
	}

	// ini namanya adalah table test, yaitu kita menyediakan data berupa slice yang berisi parameter dan ekspektasi hasil dari unit test
	// lalu slice tersebut kita iterasi dengan sub test
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) { // sub test
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
