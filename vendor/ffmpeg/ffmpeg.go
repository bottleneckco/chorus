package ffmpeg

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// Segment a file into segments of 5s each. Returns list of files' absolute paths.
func Segment(filename string) (SegmentResult, error) {
	timeNow := time.Now().Nanosecond()
	targetDir := path.Join(os.TempDir(), strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename)))

	result := SegmentResult{
		ContainerDir: targetDir,
	}

	err := os.MkdirAll(targetDir, 0700)
	if err != nil {
		return result, err
	}

	cmd := exec.Command("ffmpeg", "-i", filename, "-segment_time", "5", "-b:a", "320k", "-f", "segment", fmt.Sprintf("%dsegment%%03d.mp3", timeNow))
	cmd.Dir = targetDir

	err = cmd.Run()
	if err != nil {
		return result, err
	}

	dirContents, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return result, err
	}

	for _, segmentFile := range dirContents {
		if !segmentFile.IsDir() {
			result.SegmentFileNames = append(result.SegmentFileNames, segmentFile.Name())
		}
	}

	return result, nil
}
