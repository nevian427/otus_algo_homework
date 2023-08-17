package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolveLuckyTicket(t *testing.T) {
	testfiles, err := filepath.Glob("testfiles/test.*.in")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range testfiles {
		file := file[:strings.LastIndex(file, ".")]
		in, _ := os.ReadFile(file + ".in")
		if err != nil {
			t.Logf("error read input test file %s: %s\n", file, err)
			continue
		}
		input, _ := strconv.ParseUint(string(in), 10, 64)

		out, _ := os.ReadFile(file + ".out")
		if err != nil {
			t.Logf("error read output test file %s: %s\n", file, err)
			continue
		}
		output, _ := strconv.ParseUint(string(out), 10, 64)

		t.Run(filepath.Base(file)+" table calc", func(t *testing.T) {
			t.Parallel()

			res := solveLuckyTicketTable(input)
			require.Equal(t, output, res)
		})

		t.Run(filepath.Base(file)+" recursive calc", func(t *testing.T) {
			if input > 3 {
				t.Skip("too slow - skip")
			}
			count = 0
			solveLuckyTicketRecursive(input, 0, 0)
			require.Equal(t, output, count)
		})

		t.Run(filepath.Base(file)+" sum calc", func(t *testing.T) {
			t.Parallel()

			if input > 6 {
				t.Skip("too slow - skip")
			}
			res := solveLuckyTicketSum(input)
			require.Equal(t, output, res)
		})
	}

}
