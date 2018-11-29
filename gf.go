package gf

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// GoFiler is a GoFile-er
type GoFiler struct {
	Path string
}

// BestPractices implements best practices by smashing all go files into one
func (g *GoFiler) BestPractices() (string, error) {
	files, err := filepath.Glob(g.Path + "*.go")
	if err != nil {
		return "", err
	}
	imports := []string{"import ("}
	code := make([]string, 0)
	pack := ""
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		file, err := os.Open(f)
		if err != nil {
			return "", err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		inImports := false
		for scanner.Scan() {
			s := scanner.Text()
			if strings.HasPrefix(s, "package") {
				pack = s
			} else if strings.HasPrefix(s, "import") {
				// either it's the start of a block
				// or a single line
				if strings.HasSuffix(s, "(") {
					inImports = true
					continue
				}
				parts := strings.Split(s, "import")
				imports = append(imports, "\t"+strings.TrimSpace(parts[1])) // need to tab it once to look nice
			} else if s == ")" && inImports {
				// closing line of a block
				inImports = false
				continue
			} else if inImports {
				imports = append(imports, s)
			} else {
				// it's none of those things, we're done capturing imports
				code = append(code, s)
			}
		}
	}
	data := []string{pack, ""}
	// We need to dedupe but also preserve order...
	lazyUnique := make(map[string]struct{})
	uniqueImports := make([]string, 0)
	for _, i := range imports {
		if _, ok := lazyUnique[i]; !ok {
			lazyUnique[i] = struct{}{}
			uniqueImports = append(uniqueImports, i)
		}
	}
	data = append(data, uniqueImports...)
	data = append(data, ")") // missing from imports at this point
	data = append(data, code...)
	// we have all the go files in the package
	// we need to extract import statements, smush them together
	// write those at the top
	// then put everything else

	return strings.Join(data, "\n"), nil
}
