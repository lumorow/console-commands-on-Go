package Commandline

import (
	"archive/tar"
	"compress/gzip"
	"day_2/myRotate/Flags"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Commandline struct {
	Args  []string
	Path  map[string]int
	Arch  string
	Flags Flags.Flags
}

func New(args []string) *Commandline {
	return &Commandline{
		Args: args,
	}
}

func (cl *Commandline) Zip() error {
	wg := new(sync.WaitGroup)
	for path, _ := range cl.Path {
		wg.Add(1)
		go createArch(wg, path, cl.Arch)
	}
	wg.Wait()
	return nil
}

func createArch(wg *sync.WaitGroup, logPath, archPath string) {
	defer wg.Done()
	//
	mtime := getTimeFile(logPath)
	pathSplit := strings.Split(logPath, "/")
	filename := pathSplit[len(pathSplit)-1]
	filenameWithoutExt := strings.Split(filename, ".")
	if archPath == "" {
		archPath += "data/"
	}
	archPath += filenameWithoutExt[0] + "_" + strconv.Itoa(mtime) + ".tag.gz"
	//archive, err := os.Create(archPath)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer archive.Close()
	//gzipWriter := gzip.NewWriter(archive)
	//file, err := os.Open(logPath)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//written, err := gzipWriter.Create(filename)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if _, err := io.Copy(written, file); err != nil {
	//	log.Fatal(err)
	//}

	// Create output file
	out, err := os.Create(archPath)
	if err != nil {
		log.Fatalln("Error writing archive:", err)
	}
	defer out.Close()
	err = createArchive(logPath, out)
	if err != nil {
		log.Fatalln("Error creating archive:", err)
	}
}

func createArchive(file string, buf io.Writer) error {
	gw := gzip.NewWriter(buf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	err := addToArchive(tw, file)
	if err != nil {
		return err
	}
	return nil
}

func addToArchive(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}
	header.Name = filename
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}
	return nil
}

func getTimeFile(path string) int {
	var mtime int
	file, _ := os.Stat(path)
	modTime := file.ModTime()
	y, mon, d := modTime.Date()
	h, min, sec := modTime.Clock()
	mtime = y*31556926 + int(mon)*2629743 + d*86400 + h*3600 + min*60 + sec
	return mtime
}
