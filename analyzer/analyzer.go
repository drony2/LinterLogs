package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks slog and zap log messages",
	Run:  run,
}

var (
	sensitiveWords = []string{
		"password",
		"token",
		"secret",
		"api_key",
		"apikey",
	}
)

func run(pass *analysis.Pass) (interface{}, error) {

	for _, file := range pass.Files {

		ast.Inspect(file, func(n ast.Node) bool {

			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			method := sel.Sel.Name

			if !isLogMethod(method) {
				return true
			}

			if !isSupportedLogger(pass, sel) {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			msg, pos, ok := extractStaticMessage(pass, call.Args[0])
			if !ok {
				return true
			}

			checkLowercase(pass, pos, msg)
			checkEnglish(pass, pos, msg)
			checkSpecial(pass, pos, msg)
			checkSensitive(pass, pos, msg)

			return true
		})
	}

	return nil, nil
}

func extractStaticMessage(pass *analysis.Pass, expr ast.Expr) (string, token.Pos, bool) {

	if pass == nil || pass.TypesInfo == nil || expr == nil {
		return "", token.NoPos, false
	}

	if typ := pass.TypesInfo.TypeOf(expr); typ != nil {
		basic, ok := typ.Underlying().(*types.Basic)
		if !ok {
			return "", token.NoPos, false
		}
		if basic.Kind() != types.String && basic.Kind() != types.UntypedString {
			return "", token.NoPos, false
		}
	}

	var builder strings.Builder
	firstPos := token.NoPos
	found := false

	var visit func(ast.Expr)
	visit = func(e ast.Expr) {
		switch v := e.(type) {
		case *ast.BasicLit:
			if v.Kind != token.STRING {
				return
			}
			s, err := strconv.Unquote(v.Value)
			if err != nil {
				return
			}
			if firstPos == token.NoPos {
				firstPos = v.Pos()
			}
			builder.WriteString(s)
			found = true
		case *ast.BinaryExpr:
			if v.Op != token.ADD {
				return
			}
			visit(v.X)
			visit(v.Y)
		case *ast.ParenExpr:
			visit(v.X)
		}
	}

	visit(expr)
	if !found {
		return "", token.NoPos, false
	}

	return builder.String(), firstPos, true
}

func isLogMethod(name string) bool {

	switch name {
	case "Info", "Error", "Warn", "Debug":
		return true
	}

	return false
}

func isSupportedLogger(pass *analysis.Pass, sel *ast.SelectorExpr) bool {

	if pass == nil || pass.TypesInfo == nil {
		return false
	}

	// Package-qualified call (e.g. slog.Info(...)).
	if ident, ok := sel.X.(*ast.Ident); ok {
		obj := pass.TypesInfo.Uses[ident]
		if pkgName, ok := obj.(*types.PkgName); ok {
			pkg := pkgName.Imported()
			if pkg == nil {
				return false
			}
			return isSupportedLoggerPackage(pkg.Path())
		}
	}

	// Method call on a logger value (e.g. logger.Info(...), zap.L().Info(...)).
	typ := pass.TypesInfo.TypeOf(sel.X)
	return isSupportedLoggerType(typ)
}

func isSupportedLoggerPackage(path string) bool {

	switch path {
	case "log/slog":
		return true
	case "go.uber.org/zap":
		return true
	default:
		return false
	}
}

func isSupportedLoggerType(typ types.Type) bool {

	if typ == nil {
		return false
	}

	for {
		ptr, ok := typ.(*types.Pointer)
		if !ok {
			break
		}
		typ = ptr.Elem()
	}

	named, ok := typ.(*types.Named)
	if !ok {
		return false
	}

	obj := named.Obj()
	if obj == nil || obj.Pkg() == nil {
		return false
	}

	return isSupportedLoggerPackage(obj.Pkg().Path())
}

func checkLowercase(pass *analysis.Pass, pos token.Pos, msg string) {

	msg = strings.TrimSpace(msg)
	if msg == "" {
		return
	}

	var first rune
	for _, r := range msg {
		first = r
		break
	}

	if !unicode.IsLower(first) {
		pass.Reportf(pos, "log message must start with lowercase")
	}
}

func checkEnglish(pass *analysis.Pass, pos token.Pos, msg string) {

	for _, r := range msg {
		if unicode.Is(unicode.Cyrillic, r) {
			pass.Reportf(pos, "log message must be in English")
			return
		}
	}
}

func checkSpecial(pass *analysis.Pass, pos token.Pos, msg string) {

	for _, r := range msg {
		if r == ' ' {
			continue
		}
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			continue
		}
		pass.Reportf(pos, "log message contains special characters")
		return
	}
}

func checkSensitive(pass *analysis.Pass, pos token.Pos, msg string) {

	lower := strings.ToLower(msg)

	for _, word := range sensitiveWords {
		if strings.Contains(lower, word) {
			pass.Reportf(pos, "log message contains sensitive data: %s", word)
		}
	}
}
