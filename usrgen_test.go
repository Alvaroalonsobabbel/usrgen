package usrgen_test

import (
	"testing"

	"github.com/Alvaroalonsobabbel/usrgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) { // nolint:funlen
	t.Parallel()

	t.Run("Character Substitution", func(t *testing.T) {
		t.Parallel()

		charSubTests := []struct {
			first, last, want string
		}{
			{"Test", "Person", "tperson"},
			{"Test", "Person2", "tperson2"},
			{"Test", "Pérson", "tperson"},
			{"Test", "Përson", "tperson"},
			{"Test", "Persôn", "tperson"},

			// German specific transliterations
			{"Test", "Persön", "tpersoen"},
			{"Test", "Persün", "tpersuen"},
			{"Test", "Persän", "tpersaen"},
			{"Test", "Perßon", "tpersson"},

			// Special character transliteration
			{"Test", "Person Space", "tpersonspace"},
			{"Test", "Person-Hypen", "tperson-hypen"},

			// Underscores shall be removed
			{"Test", "Person_Underscore", "tpersonunderscore"},
		}

		for _, test := range charSubTests {
			usrgen := usrgen.New(test.first, test.last, "de")
			username, err := usrgen.Generate()
			require.NoError(t, err)
			assert.Equal(t, test.want, username)
		}
	})

	t.Run("Calling Generate() adds letters from the first name", func(t *testing.T) {
		t.Parallel()

		concurrencyTests := []struct {
			n         int
			want      string
			expectErr bool
		}{
			// Examples for user "Test Person"
			{1, "tperson", false},
			{2, "teperson", false},
			{3, "tesperson", false},
			{4, "testperson", false},
			{5, "", true},
		}

		for _, test := range concurrencyTests {
			var (
				usrgen   = usrgen.New("test", "person", "de")
				username string
				err      error
			)

			for range test.n {
				username, err = usrgen.Generate()
			}

			if test.expectErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, test.want, username)
		}
	})
}
