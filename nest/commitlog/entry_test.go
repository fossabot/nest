package commitlog

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEntry(t *testing.T) {
	t.Run("should encode to a valid binary array", func(t *testing.T) {
		e := newEntry(1, []byte("test"))
		require.Equal(t, uint64(24), e.Size())
		buf := make([]byte, e.Size())
		require.NoError(t, encodeEntry(e, buf))
		require.Equal(t,
			[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x18,
				0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1,
				0x74, 0x65, 0x73, 0x74,
				0x74, 0x65, 0x73, 0x74,
			},
			buf)
	})
	t.Run("should write into io.Writer", func(t *testing.T) {
		e := newEntry(1, []byte("test"))
		buf := bytes.NewBuffer(nil)
		n, err := writeEntry(e, buf)
		require.NoError(t, err)
		require.Equal(t, int(e.Size()), n)
		require.Equal(t,
			[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x18,
				0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1,
				0x74, 0x65, 0x73, 0x74,
				0x74, 0x65, 0x73, 0x74,
			},
			buf.Bytes())
	})

	t.Run("should decode from a valid binary array", func(t *testing.T) {
		buf := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x18,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1,
			0x74, 0x65, 0x73, 0x74,
			0x74, 0x65, 0x73, 0x74,
		}
		e, err := decodeEntry(buf)
		require.NoError(t, err)
		require.Equal(t, uint64(24), e.Size())
		require.Equal(t, []byte("test"), e.Payload())
	})
	t.Run("should read from io.Reader", func(t *testing.T) {
		r := bytes.NewReader([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x18,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1,
			0x74, 0x65, 0x73, 0x74,
			0x74, 0x65, 0x73, 0x74,
		})
		buf := make([]byte, r.Len())
		e, err := readEntry(r, buf)
		require.NoError(t, err)
		require.Equal(t, uint64(24), e.Size())
		require.Equal(t, []byte("test"), e.Payload())
	})

}