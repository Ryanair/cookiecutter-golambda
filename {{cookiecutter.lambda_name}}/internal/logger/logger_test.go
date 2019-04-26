package logger_test

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"

	"stash.ryanair.com/{{cookiecutter.project_name}}/{{cookiecutter.lambda_name}}/internal/logger"
)

func TestLogger(t *testing.T) {
	info := "info"
	infof := "infof"
	infow := "infow"
	warn := "warn"
	warnf := "warnf"
	warnw := "warnw"
	err := "error"
	errf := "errof"
	errw := "errw"

	out := captureLogs(t, func() {
		logger.Initialize()

		logger.Info(info)
		logger.Infof(infof)
		logger.Infow(infow)
		logger.Warn(warn)
		logger.Warnf(warnf)
		logger.Warnw(warnw)
		logger.Error(err)
		logger.Errorf(errf)
		logger.Errorw(errw)
	})

	for _, txt := range []string{info, infof, infow, warn, warnf, warnw, err, errf, errw} {
		if !strings.Contains(out, txt) {
			t.Errorf("log output %q does not contain expected text: %q", out, txt)
		}
	}
}

func captureLogs(t *testing.T, fn func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	stderr := os.Stderr
	os.Stderr = w
	defer func() {
		os.Stderr = stderr
	}()

	fn()

	w.Close()
	var b bytes.Buffer
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		b.Write(scanner.Bytes())
	}

	return b.String()
}
