package mem

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	parserUtil "github.com/lucasjo/porgex-go/util"
)

func getCgroupModuleInfo(cPath, cFile string) (uint64, error) {
	contents, err := ioutil.ReadFile(filepath.Join(cPath, cFile))

	if err != nil {
		return 0, err
	}

	return parserUtil.ParseUint(strings.TrimSpace(string(contents)), 10, 64)
}
