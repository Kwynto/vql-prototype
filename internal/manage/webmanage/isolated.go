package webmanage

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Kwynto/gosession"
	"github.com/Kwynto/vql-prototype/internal/engine/basicsystem/gauth"
	"github.com/Kwynto/vql-prototype/pkg/lib/e"
)

// Isolation of statistical data for web access.

type TIsolatedFS struct {
	fs http.FileSystem
}

func (ifs TIsolatedFS) Open(sPath string) (f http.File, err error) {
	sOperation := "internal -> WebManage -> isolated -> Open"
	defer func() { e.Wrapper(sOperation, err) }()

	f, err = ifs.fs.Open(sPath)
	if err != nil {
		return nil, err
	}

	inStat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if inStat.IsDir() {
		sIndex := filepath.Join(sPath, "index.html")
		if _, err := ifs.fs.Open(sIndex); err != nil {
			errClose := f.Close()
			if errClose != nil {
				return nil, errClose
			}
			return nil, err
		}
	}

	return f, nil
}

// Isolation of authorization.
func IsolatedAuth(w http.ResponseWriter, r *http.Request, slIRules []gauth.TRole) bool {
	sSesID := gosession.Start(&w, r)
	inAuth := sSesID.Get("auth")
	sLogin := fmt.Sprint(inAuth)
	stProfile, err := gauth.GetProfile(sLogin)
	if err != nil {
		return true
	}

	if !stProfile.IsAllowed(slIRules) {
		return true
	}

	if stProfile.Status == gauth.NEW {
		stProfile.Status = gauth.ACTIVE
		gauth.UpdateProfile(sLogin, stProfile)
	}

	return false
}
