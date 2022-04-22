package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

type UserFace struct {
	Browsers []string `json:"browsers"`
	Email    string   `json:"email"`
	Name     string   `json:"name"`
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	sc := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	sc.Split(bufio.ScanLines)

	seenBrowsers := []string{}
	uniqueBrowsers := 0

	buf := bytes.Buffer{}

	//easyjson:json
	users := make([]UserFace, 0)
	for sc.Scan() {
		user := &UserFace{}

		err = user.UnmarshalJSON(sc.Bytes())
		if err != nil {
			panic(err)
		}
		users = append(users, *user)
	}

	for i, user := range users {

		isAndroid := false
		isMSIE := false

		browsers := user.Browsers

		for _, browserRaw := range browsers {
			browser := browserRaw

			if ok := strings.Contains(browser, "Android"); ok {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		for _, browserRaw := range browsers {
			browser := browserRaw

			if ok := strings.Contains(browser, "MSIE"); ok {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := strings.Replace(user.Email, "@", " [at] ", -1)

		buf.WriteByte('[')
		buf.WriteString(strconv.Itoa(i))
		buf.WriteString("] ")
		buf.WriteString(user.Name)
		buf.WriteString(" <")
		buf.WriteString(email)
		buf.WriteString(">\n")

	}

	fmt.Fprintln(out, "found users:\n"+buf.String())
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}

var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9f2eff5fDecodeHw3BenchStructUser(in *jlexer.Lexer, out *UserFace) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeHw3BenchStructUser(out *jwriter.Writer, in UserFace) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		out.RawString(prefix[1:])
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserFace) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeHw3BenchStructUser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserFace) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeHw3BenchStructUser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserFace) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeHw3BenchStructUser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserFace) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeHw3BenchStructUser(l, v)
}
