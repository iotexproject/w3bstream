// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package project

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject(in *jlexer.Lexer, out *Project) {
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
		case "defaultVersion":
			out.DefaultVersion = string(in.String())
		case "versions":
			if in.IsNull() {
				in.Skip()
				out.Versions = nil
			} else {
				in.Delim('[')
				if out.Versions == nil {
					if !in.IsDelim(']') {
						out.Versions = make([]*Config, 0, 8)
					} else {
						out.Versions = []*Config{}
					}
				} else {
					out.Versions = (out.Versions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Config
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Config)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Versions = append(out.Versions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject(out *jwriter.Writer, in Project) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"defaultVersion\":"
		out.RawString(prefix[1:])
		out.String(string(in.DefaultVersion))
	}
	{
		const prefix string = ",\"versions\":"
		out.RawString(prefix)
		if in.Versions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Versions {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Project) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Project) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Project) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Project) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject(l, v)
}
func easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject1(in *jlexer.Lexer, out *Meta) {
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
		case "ProjectID":
			out.ProjectID = uint64(in.Uint64())
		case "Uri":
			out.Uri = string(in.String())
		case "Hash":
			if in.IsNull() {
				in.Skip()
			} else {
				copy(out.Hash[:], in.Bytes())
			}
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
func easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject1(out *jwriter.Writer, in Meta) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ProjectID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ProjectID))
	}
	{
		const prefix string = ",\"Uri\":"
		out.RawString(prefix)
		out.String(string(in.Uri))
	}
	{
		const prefix string = ",\"Hash\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Hash[:])
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Meta) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Meta) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Meta) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Meta) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject1(l, v)
}
func easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject2(in *jlexer.Lexer, out *Config) {
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
		case "version":
			out.Version = string(in.String())
		case "vmTypeID":
			out.VMTypeID = uint64(in.Uint64())
		case "codeExpParam":
			out.CodeExpParam = string(in.String())
		case "code":
			out.Code = string(in.String())
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
func easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject2(out *jwriter.Writer, in Config) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix[1:])
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"vmTypeID\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.VMTypeID))
	}
	if in.CodeExpParam != "" {
		const prefix string = ",\"codeExpParam\":"
		out.RawString(prefix)
		out.String(string(in.CodeExpParam))
	}
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.String(string(in.Code))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Config) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Config) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Config) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Config) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject2(l, v)
}
func easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject3(in *jlexer.Lexer, out *Attribute) {
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
		case "Paused":
			out.Paused = bool(in.Bool())
		case "RequestedProverAmount":
			out.RequestedProverAmount = uint64(in.Uint64())
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
func easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject3(out *jwriter.Writer, in Attribute) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Paused\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Paused))
	}
	{
		const prefix string = ",\"RequestedProverAmount\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.RequestedProverAmount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Attribute) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Attribute) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b166cadEncodeGithubComIotexprojectW3bstreamProject3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Attribute) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Attribute) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b166cadDecodeGithubComIotexprojectW3bstreamProject3(l, v)
}
