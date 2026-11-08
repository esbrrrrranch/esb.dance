package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func writeFixture(t *testing.T, root, name, content string) {
	t.Helper()
	path := filepath.Join(root, name)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func readFixture(t *testing.T, root, name string) string {
	t.Helper()
	content, err := os.ReadFile(filepath.Join(root, name))
	if err != nil {
		t.Fatal(err)
	}
	return string(content)
}

func fixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFixture(t, root, "templates/layout.html", `{{define "body-class"}}{{end}}{{define "layout"}}<body class="{{template "body-class" .}}"><a href="{{.Path}}"{{if eq .Path "/"}} aria-current="page"{{end}}>Home</a>{{if .Canonical}}<link rel="canonical" href="{{.Canonical}}">{{end}}{{template "content" .}}</body>{{end}}`)
	writeFixture(t, root, "www/index.html", `{{define "content"}}Home{{end}}{{template "layout" .}}`)
	writeFixture(t, root, "www/choreography/index.html", `{{define "body-class"}}has-hero{{end}}{{define "content"}}Dance{{end}}{{template "layout" .}}`)
	writeFixture(t, root, "www/404.html", `{{define "content"}}Not found{{end}}{{template "layout" .}}`)
	writeFixture(t, root, "www/emily-schoen-branch/index.html", `<meta http-equiv="refresh" content="0; url=/">`)
	writeFixture(t, root, "www/assets/styles.css", "body { color: black; }")
	writeFixture(t, root, "www/.nojekyll", "")
	return root
}

func TestBuildPagesAndAssets(t *testing.T) {
	root := fixture(t)
	if err := buildSite(root); err != nil {
		t.Fatal(err)
	}
	home := readFixture(t, root, "dist/index.html")
	if !strings.Contains(home, `aria-current="page"`) || strings.Contains(home, "has-hero") || !strings.Contains(home, `href="https://esb.dance/"`) {
		t.Fatalf("incorrect home metadata or leaked page definition: %s", home)
	}
	hero := readFixture(t, root, "dist/choreography/index.html")
	if !strings.Contains(hero, `class="has-hero"`) || strings.Contains(hero, "aria-current") || !strings.Contains(hero, `href="https://esb.dance/choreography/"`) {
		t.Fatalf("incorrect choreography metadata: %s", hero)
	}
	if strings.Contains(readFixture(t, root, "dist/404.html"), "canonical") {
		t.Fatal("404 page must not have a canonical URL")
	}
	for _, name := range []string{".nojekyll", "assets/styles.css", "emily-schoen-branch/index.html"} {
		if readFixture(t, root, "dist/"+name) != readFixture(t, root, "www/"+name) {
			t.Fatalf("changed static file: %s", name)
		}
	}
}

func TestTemplateFailurePreservesBuild(t *testing.T) {
	root := fixture(t)
	if err := buildSite(root); err != nil {
		t.Fatal(err)
	}
	previous := readFixture(t, root, "dist/index.html")
	writeFixture(t, root, "www/index.html", `{{template "missing" .}}`)
	if err := buildSite(root); err == nil {
		t.Fatal("expected missing template to fail the build")
	}
	if got := readFixture(t, root, "dist/index.html"); got != previous {
		t.Fatal("failed build changed the generated page")
	}
	writeFixture(t, root, "www/index.html", `{{define "content"}}Repaired{{end}}{{template "layout" .}}`)
	if err := buildSite(root); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(readFixture(t, root, "dist/index.html"), "Repaired") {
		t.Fatal("repaired template did not update output")
	}
}

func TestRebuildTracksChangesAndRemovals(t *testing.T) {
	root := fixture(t)
	if err := buildSite(root); err != nil {
		t.Fatal(err)
	}
	page := filepath.Join(root, "dist/index.html")
	oldTime := time.Unix(100, 0)
	if err := os.Chtimes(page, oldTime, oldTime); err != nil {
		t.Fatal(err)
	}
	writeFixture(t, root, "www/assets/styles.css", "body { color: blue; }")
	if err := os.Remove(filepath.Join(root, "www/choreography/index.html")); err != nil {
		t.Fatal(err)
	}
	if err := buildSite(root); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(page)
	if err != nil || !info.ModTime().Equal(oldTime) {
		t.Fatal("unchanged page was rewritten")
	}
	if got := readFixture(t, root, "dist/assets/styles.css"); !strings.Contains(got, "blue") {
		t.Fatal("asset change was not copied")
	}
	if _, err := os.Stat(filepath.Join(root, "dist/choreography/index.html")); !os.IsNotExist(err) {
		t.Fatalf("deleted page remains in output: %v", err)
	}
}
