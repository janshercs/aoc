package dayseven

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// need to have a file system data structure
// directory can have other directories (array of directories?)
// directory can have a parent (unless root node)
// should root node parent point to itself or throw err?
// directory can have files

func TestDirectorySize(t *testing.T) {
	fileA := file{size: 2}
	fileB := file{size: 3}

	d := directory{
		name:  "test",
		files: []file{fileA, fileB},
	}

	assert.Equal(t, 5, d.size())

	fileC := file{size: 5}
	c := directory{
		directories: []*directory{&d},
		files:       []file{fileC},
	}
	assert.Equal(t, 10, c.size())

}

func TestChangeDirectory(t *testing.T) {
	fileC := file{size: 5}
	d := &directory{name: "d"}
	c := &directory{
		directories: []*directory{d},
		files:       []file{fileC},
	}

	fileA := file{size: 2}
	fileB := file{size: 3}

	d.parent = c
	d.files = append(d.files, fileA, fileB)

	pwd := c
	assert.Equal(t, d, pwd.cd("d"))

	assert.Equal(t, c, d.cd(".."))
}

func TestIsCommand(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			input: "$ cd /",
			want:  true,
		},
		{
			input: "$ ls",
			want:  false,
		},
		{
			input: "dir jmtrrrp",
			want:  false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.want, isCD(tc.input))
		})
	}
}

func TestTargetDest(t *testing.T) {
	assert.Equal(t, "/", targetDest("$ cd /"))
	assert.Equal(t, "..", targetDest("$ cd .."))
	assert.Equal(t, "abc", targetDest("$ cd abc"))
}

func TestCreateFile(t *testing.T) {
	assert.Equal(t, file{"hi", 2}, createFile("2 hi"))
}

func TestCreateDir(t *testing.T) {
	assert.Equal(t, &directory{name: "gzqg"}, createDirectory("dir gzqg"))
}

func TestAddDir(t *testing.T) {
	root := createDirectory("dir /")
	root.addDir(createDirectory("dir first"))
	want := &directory{
		name: "/",
		directories: []*directory{
			{
				name:   "first",
				parent: root,
			},
		},
	}

	assert.Equal(t, want, root)
}

func TestSolutions(t *testing.T) {
	assert.Equal(t, 1086293, solutionA())
	assert.Equal(t, 366028, solutionB())
}
