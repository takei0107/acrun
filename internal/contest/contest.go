package contest

import "github.com/takei0107/acrun/internal/util"

type contest struct {
	name string
}

func ResolveContestName(s string) (string, error) {
	if s != "" {
		return s, nil
	}

	return util.GetCurrentDirName()
}
