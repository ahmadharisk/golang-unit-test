package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("before test")

	m.Run()

	fmt.Println("after test")
}

func TestSayHelloSuccess(t *testing.T) {
	result := SayHello("peppo")
	if result != "Hello peppo" {
		t.Fail()
	}
	fmt.Println("TestSayHelloSuccess done")
}

func TestSayHelloFailNow(t *testing.T) {
	t.FailNow()

	fmt.Println("TestSayHelloFail done")
}

func TestSayHelloError(t *testing.T) {
	t.Error("TestSayHelloError error")
	fmt.Println("TestSayHelloError done")
}

func TestSayHelloFatal(t *testing.T) {
	t.Fatal("TestSayHelloFatal fatal")
	fmt.Println("TestSayHelloFatal done")
}

func TestSayHelloAsserttion(t *testing.T) {
	result := SayHello("peppo")
	// assert sama seperti Fail()
	assert.Equal(t, "Hello peppo", result, "result must be 'Hello peppo'")
	fmt.Println("TestSayHelloAsserttion done")
}

func TestSayHelloRequire(t *testing.T) {
	// require sama seperti FailNow()
	require.Equal(t, "Hello pepp", SayHello("peppo"))
	fmt.Println("TestSayHelloRequire done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("tidak bisa jalan di mac")
	}

	result := SayHello("peppo")
	assert.Equal(t, "Hello peppo", result, "result must be 'Hello peppo'")
}

func TestSubTest(t *testing.T) {
	t.Run("subtest1", func(t *testing.T) {
		result := SayHello("peppo")
		assert.Equal(t, "Hello peppo", result, "result must be 'Hello peppo'")
	})
	t.Run("subtest2", func(t *testing.T) {
		result := SayHello("peppo")
		assert.Equal(t, "Hello peppo", result, "result must be 'Hello peppo'")
	})
}

func TestTableSayHello(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "peppo",
			request:  "peppo",
			expected: "Hello peppo",
		},
		{
			name:     "silpiana",
			request:  "silpiana",
			expected: "Hello silpiana",
		},
		{
			name:     "monsep",
			request:  "monsep",
			expected: "Hello monsep",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := SayHello(tt.request)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("peppo")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("peppo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("peppo")
		}
	})

	b.Run("silpiana", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("silpiana")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmark := []struct {
		name    string
		request string
	}{
		{
			name:    "peppo",
			request: "peppo",
		},
		{
			name:    "silpiana",
			request: "silpiana",
		},
	}

	for _, tt := range benchmark {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(tt.request)
			}
		})
	}
}
