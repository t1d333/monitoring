// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package apimodels

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	models "github.com/t1d333/vk_edu_db_project/internal/models"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC1727565DecodeGithubComT1d333VkEduDbProjectInternalPostApiModels(in *jlexer.Lexer, out *GetPostResponse) {
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
		case "post":
			if in.IsNull() {
				in.Skip()
				out.Post = nil
			} else {
				if out.Post == nil {
					out.Post = new(models.Post)
				}
				(*out.Post).UnmarshalEasyJSON(in)
			}
		case "author":
			if in.IsNull() {
				in.Skip()
				out.Author = nil
			} else {
				if out.Author == nil {
					out.Author = new(models.User)
				}
				(*out.Author).UnmarshalEasyJSON(in)
			}
		case "forum":
			if in.IsNull() {
				in.Skip()
				out.Forum = nil
			} else {
				if out.Forum == nil {
					out.Forum = new(models.Forum)
				}
				(*out.Forum).UnmarshalEasyJSON(in)
			}
		case "thread":
			if in.IsNull() {
				in.Skip()
				out.Thread = nil
			} else {
				if out.Thread == nil {
					out.Thread = new(models.Thread)
				}
				(*out.Thread).UnmarshalEasyJSON(in)
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
func easyjsonC1727565EncodeGithubComT1d333VkEduDbProjectInternalPostApiModels(out *jwriter.Writer, in GetPostResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"post\":"
		out.RawString(prefix[1:])
		if in.Post == nil {
			out.RawString("null")
		} else {
			(*in.Post).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		if in.Author == nil {
			out.RawString("null")
		} else {
			(*in.Author).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		if in.Forum == nil {
			out.RawString("null")
		} else {
			(*in.Forum).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		if in.Thread == nil {
			out.RawString("null")
		} else {
			(*in.Thread).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetPostResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1727565EncodeGithubComT1d333VkEduDbProjectInternalPostApiModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetPostResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1727565EncodeGithubComT1d333VkEduDbProjectInternalPostApiModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetPostResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1727565DecodeGithubComT1d333VkEduDbProjectInternalPostApiModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetPostResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1727565DecodeGithubComT1d333VkEduDbProjectInternalPostApiModels(l, v)
}
func easyjsonC1727565DecodeGithubComT1d333VkEduDbProjectInternalPostApiModels1(in *jlexer.Lexer, out *GetPostParams) {
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
		case "Author":
			out.Author = bool(in.Bool())
		case "Forum":
			out.Forum = bool(in.Bool())
		case "Thread":
			out.Thread = bool(in.Bool())
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
func easyjsonC1727565EncodeGithubComT1d333VkEduDbProjectInternalPostApiModels1(out *jwriter.Writer, in GetPostParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Author\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Author))
	}
	{
		const prefix string = ",\"Forum\":"
		out.RawString(prefix)
		out.Bool(bool(in.Forum))
	}
	{
		const prefix string = ",\"Thread\":"
		out.RawString(prefix)
		out.Bool(bool(in.Thread))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetPostParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1727565EncodeGithubComT1d333VkEduDbProjectInternalPostApiModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetPostParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1727565EncodeGithubComT1d333VkEduDbProjectInternalPostApiModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetPostParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1727565DecodeGithubComT1d333VkEduDbProjectInternalPostApiModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetPostParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1727565DecodeGithubComT1d333VkEduDbProjectInternalPostApiModels1(l, v)
}