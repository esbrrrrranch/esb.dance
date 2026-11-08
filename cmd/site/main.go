package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type pageData struct {
	Canonical string
	Path      string
}

func main() {
	log.SetFlags(0)
	if err := buildSite("."); err != nil {
		log.Fatal(err)
	}
	log.Print("built site in dist")
}

func buildSite(root string) error {
	shared, err := template.New("site").Option("missingkey=error").ParseGlob(filepath.Join(root, "templates", "*.html"))
	if err != nil {
		return fmt.Errorf("parse shared templates: %w", err)
	}
	inputs := make(map[string]string)
	rendered := make(map[string][]byte)
	source := filepath.Join(root, "www")
	err = filepath.WalkDir(source, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("source symlinks are not supported: %s", path)
		}
		rel, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		inputs[rel] = path
		if filepath.Ext(path) != ".html" {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		page, err := shared.Clone()
		if err != nil {
			return err
		}
		if _, err = page.New("page").Parse(string(content)); err != nil {
			return fmt.Errorf("parse %s: %w", rel, err)
		}
		url := "/" + filepath.ToSlash(rel)
		url = strings.TrimSuffix(url, "index.html")
		data := pageData{Canonical: "https://esb.dance" + url, Path: url}
		if rel == "404.html" {
			data.Canonical = ""
		}
		var output bytes.Buffer
		if err := page.ExecuteTemplate(&output, "page", data); err != nil {
			return fmt.Errorf("render %s: %w", rel, err)
		}
		rendered[rel] = output.Bytes()
		return nil
	})
	if err != nil {
		return err
	}
	// Render every page before updating output so template errors preserve the last good build.
	destination := filepath.Join(root, "dist")
	for rel, path := range inputs {
		target := filepath.Join(destination, rel)
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}
		if content, ok := rendered[rel]; ok {
			previous, err := os.ReadFile(target)
			if err == nil && bytes.Equal(previous, content) {
				continue
			}
			if err := os.WriteFile(target, content, 0644); err != nil {
				return err
			}
		} else if err := copyAsset(path, target); err != nil {
			return err
		}
	}
	// Remove generated files whose sources were deleted or renamed.
	return filepath.WalkDir(destination, func(path string, entry fs.DirEntry, err error) error {
		if err != nil || entry.IsDir() {
			return err
		}
		rel, err := filepath.Rel(destination, path)
		if err != nil {
			return err
		}
		if _, exists := inputs[rel]; !exists {
			return os.Remove(path)
		}
		return nil
	})
}

func copyAsset(source, target string) error {
	info, err := os.Stat(source)
	if err != nil {
		return err
	}
	previous, err := os.Stat(target)
	if err == nil && previous.Size() == info.Size() && previous.ModTime().Equal(info.ModTime()) {
		return nil
	}
	input, err := os.Open(source)
	if err != nil {
		return err
	}
	defer input.Close()
	output, err := os.Create(target)
	if err != nil {
		return err
	}
	_, copyErr := io.Copy(output, input)
	closeErr := output.Close()
	if err := errors.Join(copyErr, closeErr); err != nil {
		return err
	}
	return os.Chtimes(target, info.ModTime(), info.ModTime())
}
