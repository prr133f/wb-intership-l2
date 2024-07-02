package domain

import (
	"io"
	"net/http"
	"os"
)

func (d *Domain) Wget(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		d.Log.Error(err.Error())
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create("out.txt")
	if err != nil {
		d.Log.Error(err.Error())
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		d.Log.Error(err.Error())
		return err
	}

	return nil
}
