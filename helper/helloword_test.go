package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWord(b *testing.B) {
	for i:=0 ; i<b.N ; i++{
		HelloWord("Rizal")
	}
}

func BenchmarkHelloWordNurizalludin(b *testing.B) {
	for i:=0; i<b.N ; i++{
		HelloWord("Nurizalludin")
	}
}

func BenchmarkHelloWordSub(b *testing.B) {
	b.Run("Rizal", func(b *testing.B){
		for i:=0 ; i<b.N ; i++{
			HelloWord("Rizal")
		}
	})
	b.Run("Nurizalludin",func(b *testing.B) {
		for i:=0; i<b.N ; i++{
			HelloWord("Nurizalludin")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks :=[]struct{
		name string
		request string
	}{
		{
			name:"Rizal",
			request: "Rizal",
		},
		{
			name: "Nurizalludin",
			request: "Nurizalludin",
		},
		{
			name: "Ganteng",
			request: "Ganteng",
		},
	}

	for _,benchmark:=range benchmarks{
		b.Run(benchmark.name,func(b *testing.B) {
			for i:=0 ; i<b.N ; i++{
				HelloWord(benchmark.name)
			}
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Rizal", func(t *testing.T) {
		result := HelloWord("Rizal")
		require.Equal(t, "Hello Rizal", result, "Result must be Hello Rizal")
	})
	t.Run("Nurizalludin", func(t *testing.T) {
		result := HelloWord("Nurizalludin")
		require.Equal(t, "Hello Nurizalludin", result, "Result must be Hello Nurizalludin")

	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("After UNIT TEST")
}

func TestHelloWord(t *testing.T) {
	result := HelloWord("Rizal")
	if result != "Hello Rizal" {
		t.Fail()
	}
	fmt.Println("Test Helloword Done")
}

func TestHelloWordNur(t *testing.T) {
	result := HelloWord("Nurizalludin")
	if result != "Hello Nurizalludin" {
		//menghentikan test dan tidak melanjutkan ke code dibawahnya
		t.Fatal()
	}
}

//with testify

func TestHelloWordAssert(t *testing.T) {
	result := HelloWord("Rizal")
	//asser fail test namun akan mengeksekusi kode yang dibawahnya
	assert.Equal(t, "Hello Rizal", result, "Result must be Hello Rizal")
	fmt.Println("Assert Test is Done")
}

func TestHelloWordRequired(t *testing.T) {
	result := HelloWord("Rizal")
	//require fail test dan tidak melanjutkan ke code dibawahnya
	require.Equal(t, "Hello Rizal", result, "Result must be Hello Rizal")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows")
	}
}

func TestTableHelloWord(t *testing.T) {
	tests:=[]struct{
		name string
		request string
		expected string
	}{
		{
			name:"Rizal",
			request:"Rizal",
			expected:"Hello Rizal",
		},
		{
			name:"Nurizalludin",
			request:"Nurizalludin",
			expected:"Hello Nurizalludin",
		},
		{
			name:"Ebon",
			request:"Ebon",
			expected:"Hello Ebon",
		},
	}

	for _,test:=range tests{
	t.Run(test.name,func(t *testing.T){
		result:=HelloWord(test.request)
		require.Equal(t,test.expected,result)
	})

	}
}
