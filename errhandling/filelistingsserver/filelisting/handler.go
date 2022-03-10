package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(w http.ResponseWriter,
	r *http.Request) error {
	if strings.Index(r.URL.Path, prefix) != 0 {
		return userError("path must start with " + prefix)
	}

	path := r.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		// panic(err)
		return err
	}
	w.Write(all)
	return nil
}
