package download

import (
	"context"
	//"bytes"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type downloader struct {
	url        string
	tmpFiles   []string
	outPath    string
	totalSize  int
	nChunk     int
	size       int
	outputSize int
	byteMap    []byteLocation
	wg         sync.WaitGroup
	err        chan error
}

type byteLocation struct {
	start int
	end   int
	size  int
}

func newDownloader(urlStr string, outDir string, nChunk int) (*downloader, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	_, filename := filepath.Split(u.Path)

	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		return nil, err
	}
	var tmpFiles []string
	for i := 0; i < nChunk; i++ {
		tmpFiles = append(tmpFiles, filepath.Join(outDir, filename+strconv.Itoa(i)))
	}

	downloader := downloader{
		url:      urlStr,
		outPath:  filepath.Join(outDir, filename),
		tmpFiles: tmpFiles,
		nChunk:   nChunk,
		byteMap:  make([]byteLocation, nChunk),
	}
	err = downloader.getFileSize()
	if err != nil {
		return nil, err
	}
	downloader.calcByteMap()
	return &downloader, nil
}

func Download(url string, outDir string, nChunk int) error {
	start := time.Now()

	d, err := newDownloader(url, outDir, nChunk)
	if err != nil {
		return err
	}

	bc := context.Background()
	eg, ctx := errgroup.WithContext(bc)
	//ctx, cancel := context.WithTimeout(ctx, 1* time.Second)
	//defer cancel()

	for i := range d.byteMap {
		i := i
		eg.Go(func() error {
			return d.downloadChunk(i, ctx)
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	if err = d.writeFile(); err != nil {
		d.deleteChunks()
		return err
	}
	if err = d.getOutputFileSize(); err != nil {
		return err
	}
	elapsed := time.Since(start)

	d.displaySummary(elapsed)
	return nil
}

func (d *downloader) downloadChunk(idx int, ctx context.Context) error {
	//ctx, cancel := context.WithCancel(ctx)
	loc := d.byteMap[idx]
	client := &http.Client{}
	req, err := http.NewRequest("GET", d.url, nil)
	if err != nil {
		return err
	}
	headerRange := "bytes=" + strconv.Itoa(loc.start) + "-" + strconv.Itoa(loc.end)
	req.Header.Add("Range", headerRange)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(d.tmpFiles[idx])
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("[%d]...Downloaded. Start: %d, End: %d, Size:%d\n", idx, loc.start, loc.end, loc.size)

	select {
	case <-ctx.Done():
		fmt.Println("Cancelled >> ", idx)
		return nil
	case <-time.After(1 * time.Second):
		return fmt.Errorf("Timeout >> ", idx)
	default:
		return nil
	}
}

func (d *downloader) writeFile() error {
	file, err := os.Create(d.outPath)
	defer file.Close()
	if err != nil {
		return err
	}

	if err = d.mergeChunks(file); err != nil {
		return err
	}
	if err = d.deleteChunks(); err != nil {
		return err
	}
	return nil
}

func (d *downloader) mergeChunks(to *os.File) error {
	for _, p := range d.tmpFiles {
		f, _ := os.Open(p)
		defer f.Close()
		if _, err := io.Copy(to, f); err != nil {
			return err
		}
		//fmt.Println(p, " merged")
	}
	return nil
}

func (d *downloader) deleteChunks() error {
	for _, p := range d.tmpFiles {
		if _, err := os.Stat(p); err == nil {
			err := os.Remove(p)
			if err != nil {
				return err
			}
			//fmt.Println(p, " deleted")
		}
	}
	return nil
}

func (d *downloader) getFileSize() error {
	resp, err := http.Head(d.url)
	if err != nil {
		return err
	}
	//fmt.Println(resp)
	d.totalSize = int(resp.ContentLength)
	d.size = d.totalSize / d.nChunk
	//fmt.Println("total sizes ... ", d.totalSize)
	return nil
}

func (d *downloader) calcByteMap() {
	var startByte, endByte int
	for i := 0; i < d.nChunk; i++ {
		startByte = d.size * i
		if i == d.nChunk-1 {
			endByte = d.totalSize - 1
		} else {
			endByte = d.size*(i+1) - 1
		}
		d.byteMap[i] = byteLocation{startByte, endByte, endByte - startByte}
	}
}

func (d *downloader) getOutputFileSize() error {
	fi, err := os.Stat(d.outPath)
	if err != nil {
		return err
	}
	d.outputSize = int(fi.Size())
	return nil
}

func (d *downloader) displaySummary(elapsed time.Duration) {
	format := "%-30s %-30s\n"
	fmt.Println(strings.Repeat("=", 100))
	fmt.Printf("Download Completed!\n")
	fmt.Printf("[Summary]\n")
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf(format, "URL", d.url)
	fmt.Printf(format, "Output File", d.outPath)
	fmt.Printf(format, "Split Count", strconv.Itoa(d.nChunk))
	fmt.Printf(format, "Remote Size (Bytes)", strconv.Itoa(d.totalSize))
	fmt.Printf(format, "Local Size (Bytes)", strconv.Itoa(d.outputSize))
	fmt.Printf(format, "Elapsed", elapsed)
	fmt.Println(strings.Repeat("=", 100))
}
