package ipconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Ipconv_V42Long(t *testing.T) {
	t.Parallel()

	cases := []struct {
		uint32
		string
	}{
		{0, "0.0.0.0"},
		{1, "0.0.0.1"},
		{16777216, "1.0.0.0"},
		{185273099, "11.11.11.11"},
		{4294967295, "255.255.255.255"},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.uint32, V42Long(tc.string), "%s => %d", tc.string, tc.uint32)
	}
}

func Benchmark_Ipconv_V42Long(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		V42Long("255.255.255.255")
	}
}

func Test_Ipconv_V42Long_Panic(t *testing.T) {
	t.Parallel()

	ips := []string{"111.111.111.111.", "1.1.1", "a.a.a.a", "256.255.255.255", "1.1.1.1.1"}

	for _, ip := range ips {
		t.Run(ip, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(ipv4Error)
					assert.True(t, ok)
					assert.Contains(t, err.Error(), ip)
				}
			}()
			V42Long(ip)
		})
	}
}

func Test_Ipconv_Long2V4(t *testing.T) {
	t.Parallel()

	cases := []struct {
		uint32
		string
	}{
		{0, "0.0.0.0"},
		{1, "0.0.0.1"},
		{16777216, "1.0.0.0"},
		{185273099, "11.11.11.11"},
		{4294967295, "255.255.255.255"},
	}

	for _, tc := range cases {
		assert.Equalf(t, tc.string, Long2V4(tc.uint32), "%d => %s", tc.uint32, tc.string)
	}
}

func Benchmark_Ipconv_Long2V4(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Long2V4(4294967295)
	}
}
