package dayseven

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	totalDisk   = 70000000
	requirement = 30000000
)

type directory struct {
	files       []file
	directories []*directory
	parent      *directory
	name        string
	fileSize    int
}

func (d *directory) size() int {
	if d.fileSize != 0 { // for caching haha.
		return d.fileSize
	}

	sizeOfFiles, sizeOfDirectories := 0, 0
	for _, f := range d.files {
		sizeOfFiles += f.size
	}

	for _, d := range d.directories {
		sizeOfDirectories += d.size()
	}

	d.fileSize = sizeOfDirectories + sizeOfFiles

	return d.fileSize
}

func (d *directory) cd(dest string) *directory {
	if dest == "/" { // edge case, only happens once at the start
		return d
	}
	if dest == ".." {
		return d.parent
	}
	for _, dir := range d.directories {
		if dir.name == dest {
			return dir
		}
	}
	panic("dir not found")
}

func (d *directory) addDir(dir *directory) {
	dir.parent = d
	d.directories = append(d.directories, dir)
}

func (d *directory) addFile(f file) {
	d.files = append(d.files, f)
}

type file struct {
	name string
	size int
}

func parseInput(input string, pwd *directory) *directory {
	switch {
	case isLS(input):
		return pwd
	case isCD(input):
		return pwd.cd(targetDest(input))
	case isDir(input):
		pwd.addDir(createDirectory(input))
		return pwd
	default:
		pwd.addFile(createFile(input))
		return pwd
	}
}

func isLS(s string) bool { return strings.HasPrefix(s, "$ ls") }

func isCD(s string) bool { return strings.HasPrefix(s, "$ cd") }

func isDir(s string) bool { return strings.HasPrefix(s, "dir") }

func targetDest(s string) string {
	cmd := strings.Split(s, " ")
	return cmd[len(cmd)-1]
}

func createFile(s string) file {
	in := strings.Split(s, " ")
	size, err := strconv.Atoi(in[0])
	if err != nil {
		panic(err)
	}
	return file{name: in[1], size: size}
}

func createDirectory(s string) *directory {
	in := strings.Split(s, " ")
	return &directory{name: in[1]}
}

func solutionA() int {
	f, err := os.Open("./day_seven.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	root := createDirectory("dir /")
	pwd := root

	for scanner.Scan() {
		pwd = parseInput(scanner.Text(), pwd)
	}

	return sumDirs(traverseFS(root, []int{}))
}

func traverseFS(dir *directory, small []int) []int {
	for _, d := range dir.directories {
		small = traverseFS(d, small)
	}
	if dir.size() <= 100000 {
		small = append(small, dir.size())
	}

	return small
}

func getDirsAboveMin(dir *directory, passed []int, min int) []int {
	for _, d := range dir.directories {
		passed = getDirsAboveMin(d, passed, min)
	}
	if dir.size() >= min {
		passed = append(passed, dir.size())
	}

	return passed
}

func sumDirs(size []int) int {

	i := 0
	for _, d := range size {
		i += d
	}
	return i
}

func solutionB() int {
	f, err := os.Open("./day_seven.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	root := createDirectory("dir /")
	pwd := root

	for scanner.Scan() {
		pwd = parseInput(scanner.Text(), pwd)
	}

	totalFree := totalDisk - root.size()
	atLeast := requirement - totalFree

	sorted := sort.IntSlice(getDirsAboveMin(root, []int{}, atLeast))
	sorted.Sort()
	return sorted[0]

}
