package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// 嗯啊 你说啥？

const (
	// JBProducts 产品
	JBProducts = "IntelliJIdea CLion PhpStorm GoLand PyCharm WebStorm Rider DataGrip RubyMine AppCode"

	EvalDir   = "eval"
	ConfigDir = "conf"
	EmptyDir  = ""
)

var (
	jbProducts = strings.Split(JBProducts, " ")
)

func clean(fDir, subDirName string) error {

	return filepath.WalkDir(fDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path != fDir && d.IsDir() {

			if needClean(d.Name()) {
				if subDirName == EmptyDir {
					err = cleanConfigEval(path)
				} else {
					err = cleanConfigEval(filepath.Join(path, ConfigDir))
				}
				if err != nil {
					return err
				}
			}

			return filepath.SkipDir
		}
		return nil
	})

}

// needClean 是否需要删除
func needClean(dirName string) bool {
	for _, jb := range jbProducts {
		if strings.HasPrefix(dirName, jb) ||
			strings.HasPrefix(dirName, fmt.Sprintf(".%s", jb)) {
			return true
		}
	}
	return false
}

// cleanConfigEval 删除
func cleanConfigEval(path string) error {
	path = filepath.Join(path, EvalDir)

	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	fmt.Println("-", path)

	return nil
}
