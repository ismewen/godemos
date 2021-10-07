package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)
	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

func AdvanceCountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

type Header struct {
	Key, value string
}

type Status struct {
	Code   int
	Reason string
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	// 写http状态
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s \r\n", st.Code, st.Reason)
	if err != nil {
		return err
	}
	// 写请求头
	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.value)
		if err != nil {
			return err
		}
	}
	// 添加换行符
	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}
	// 拷贝到body
	_, err = io.Copy(w, body)
	return err
}

type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, e.err
}

func AdvanceWriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	_, _ = fmt.Fprintf(ew, "HTTP/1.1 %d %s \r\n", st.Code, st.Reason)
	for _, h := range headers {
		_, _ = fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.value)
	}
	_, _ = fmt.Fprintf(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err
}

func doSomething() error {
	return errors.New("test error")
}
func someStuff() error {
	err := doSomething()
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	return nil
}

