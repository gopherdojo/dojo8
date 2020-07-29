package pdl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type PDL struct {
	url      string
	proc     uint
	fileSize uint
	split    uint
	dir      string
	filename string
}

type Range struct {
	start uint
	end   uint
}

func New(proc int, url, dir string) (*PDL, error) {
	if url == "" {
		return nil, errors.New("no url specified")
	}

	pdl := &PDL{
		proc: uint(proc),
		url:  url,
		dir:  dir,
	}

	pdl.setSize()
	pdl.setFilename()

	return pdl, nil
}

func (p *PDL) Run() error {
	if err := p.download(); err != nil {
		return err
	}

	if err := p.merge(); err != nil {
		return err
	}

	fmt.Println("finished")
	return nil
}

func (p *PDL) download() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	eg, egctx := errgroup.WithContext(ctx)

	for i := 0; i < int(p.proc); i++ {
		i := i
		eg.Go(func() error {
			return p.partialDownload(egctx, i)
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func (p *PDL) merge() (rerr error) {
	out, err := os.Create(p.filePath())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to Create %s in Merge", p.filePath()))
	}
	defer func() {
		if err := out.Close(); err != nil {
			rerr = err
		}
	}()

	for i := 0; i < int(p.proc); i++ {
		worker := i + 1

		tmpfile, err := os.Open(p.tmpFilePath(worker))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to Open %s in Merge", p.tmpFilePath(worker)))
		}

		_, err = io.Copy(out, tmpfile)

		// Not use defer
		tmpfile.Close()

		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to Copy %s in Merge", p.tmpFilePath(worker)))
		}

		// delete
		if err := os.Remove(p.tmpFilePath(worker)); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to Remove %s in Merge", p.tmpFilePath(worker)))
		}
	}

	return nil
}

func (p *PDL) makeRange(i, proc uint) Range {
	start := p.split * i
	end := start + p.split - 1
	if i == proc-1 {
		end = p.fileSize
	}

	return Range{
		start: start,
		end:   end,
	}
}

func (p *PDL) setSize() error {
	resp, err := http.Head(p.url)
	if err != nil {
		return errors.Wrap(err, "failed to get Head")
	}

	p.fileSize = uint(resp.ContentLength)
	p.split = p.fileSize / p.proc

	return nil
}

func (p *PDL) setFilename() {
	token := strings.Split(p.url, "/")

	var original string
	for i := 1; original == ""; i++ {
		original = token[len(token)-i]
	}

	p.filename = original
}

func (p *PDL) tmpFilename(worker int) string {
	return fmt.Sprintf("%s.%d", p.filename, worker)
}

func (p *PDL) partialDownload(ctx context.Context, index int) error {
	worker := index + 1

	fmt.Printf("start download worker: %d\n", worker)

	// request
	req, err := http.NewRequest("GET", p.url, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create NewRequest for GET")
	}

	r := p.makeRange(uint(index), p.proc)

	// set header
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", r.start, r.end))

	// do
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to access")
	}
	defer resp.Body.Close()

	// write
	if err := p.writeTmpfile(resp.Body, int(worker)); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		fmt.Printf("cancelled worker: %d\n", worker)
		return ctx.Err()
	default:
		fmt.Printf("finish download worker: %d\n", worker)
		return nil
	}
}

func (p *PDL) filePath() string {
	return filepath.Join(p.dir, p.filename)
}

func (p *PDL) tmpFilePath(worker int) string {
	return filepath.Join(p.dir, p.tmpFilename(worker))
}

func (p *PDL) writeTmpfile(body io.Reader, worker int) (rerr error) {
	out, err := os.Create(p.tmpFilePath(worker))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to create file, worker: %d", worker))
	}
	defer func() {
		if err := out.Close(); err != nil {
			rerr = errors.Wrap(err, fmt.Sprintf("failed to close file, worker: %d", worker))
		}
	}()

	if _, err := io.Copy(out, body); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to write file, worker: %d", worker))
	}

	return nil
}
