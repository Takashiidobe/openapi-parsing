package specfinder

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Find(apiVersion, target string) ([]string, error) {
	versionSegment := filepath.Join("stable", apiVersion)
	var matches []string
	err := filepath.WalkDir("../azure-rest-api-specs/specification", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		if !strings.Contains(path, versionSegment) || strings.Contains(path, "/preview/") {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), target) {
				matches = append(matches, path)
				break
			}
		}
		return nil
	})
	return matches, err
}
