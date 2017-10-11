package youtube

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// Search for a Youtube Video
func Search(term string, numResults int) ([]YoutubeVideo, error) {
	var results []YoutubeVideo
	cmd := exec.Command("youtube-dl", fmt.Sprintf("ytsearch%d:%s", numResults, term), "-f", "bestaudio[vcodec=none,abr=160]", "--skip-download", "--no-playlist", "--print-json")

	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	err := cmd.Run()
	if err != nil {
		return results, err
	}

	if outBuffer.Len() == 0 {
		return results, errors.New(errBuffer.String())
	}

	lines := strings.Split(outBuffer.String(), "\n")
	for _, line := range lines {
		var result YoutubeVideo
		err := json.Unmarshal([]byte(line), &result)
		if err == nil {
			results = append(results, result)
		}
	}

	return results, nil
}
