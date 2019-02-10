package main

import "io"

type limitReader struct {
	reminder int64
	r io.Reader
}

func min(a int64, b int64) int64 {
	if a > b {
		return b
	}

	return a
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	m := min(lr.reminder, int64(len(p)))
	lr.reminder -= m
	n, err = lr.r.Read(p[0:m])

	if lr.reminder == 0 {
		err = io.EOF
	}

	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{ r: r, reminder: n }
}
