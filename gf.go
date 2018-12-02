package gf

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// GoFiler is a GoFile-er
type GoFiler struct {
	Pathkage string
	DryRun   bool
	Organize bool
}

// BestPractices implements best practices by smashing all go files into one
func (g *GoFiler) BestPractices() error {
	files, err := filepath.Glob(g.Pathkage + "*.go")
	if err != nil {
		return err
	}
	imports := []string{"import ("}
	consts := []string{"const ("}
	vars := []string{"var ("}
	code := make([]string, 0)
	pack := ""
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") || strings.HasSuffix(f, "Gofile.go") {
			continue
		}
		file, err := os.Open(f)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		inImports := false
		inConsts := false
		inVars := false
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
			} else if g.Organize && strings.HasPrefix(s, "const") {
				// either it's the start of a block
				// or a single line
				if strings.HasSuffix(s, "(") {
					inConsts = true
					continue
				}
				parts := strings.Split(s, "const")
				consts = append(consts, "\t"+strings.TrimSpace(parts[1])) // need to tab it once to look nice
			} else if g.Organize && strings.HasPrefix(s, "var") {
				// either it's the start of a block
				// or a single line
				if strings.HasSuffix(s, "(") {
					inVars = true
					continue
				}
				parts := strings.Split(s, "var")
				vars = append(vars, "\t"+strings.TrimSpace(parts[1])) // need to tab it once to look nice
			} else if s == ")" && inImports {
				// closing line of a block
				inImports = false
				continue
			} else if s == ")" && inConsts {
				// closing line of a block
				inConsts = false
				consts = append(consts, "")
				continue
			} else if s == ")" && inVars {
				// closing line of a block
				inVars = false
				vars = append(vars, "")
				continue
			} else if inImports {
				imports = append(imports, s)
			} else if inConsts {
				consts = append(consts, s)
			} else if inVars {
				vars = append(vars, s)
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
	if len(uniqueImports) > 1 {
		data = append(data, uniqueImports...)
		data = append(data, ")", "") // missing from imports at this point
	}
	if len(consts) > 1 {
		data = append(data, consts[:len(consts)-1]...)
		data = append(data, ")", "") // missing from imports at this point
	}
	if len(vars) > 1 {
		data = append(data, vars[:len(vars)-1]...)
		data = append(data, ")", "") // missing from imports at this point
	}
	data = append(data, code...)

	// We have everything we need to make our Gofile!
	if err = ioutil.WriteFile(path.Join(g.Pathkage, "Gofile.go"), []byte(strings.Join(data, "\n")), 0644); err != nil {
		return err
	}
	if !g.DryRun {

	}
	return nil
}
