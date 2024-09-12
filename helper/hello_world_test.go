package helper

import (
	"fmt"
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
	result := HelloWorld("Hamzahh")
	assert.Equal(t, "Hello Hamzah", result, "Result must be Hello Hamzah") // ini manggil yang t.Fail()
	fmt.Println("TestHelloWorldWithAssert done")
}

func TestHelloWorldWithRequire(t *testing.T) {
	result := HelloWorld("Silmii")
	require.Equal(t, "Hello Silmi", result, "Result must be Hello Silmi") // ini manggil yang t.FailNow()
	fmt.Println("TestHelloWorldWithRequire done")
}
