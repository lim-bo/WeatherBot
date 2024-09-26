// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package entity

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

func easyjson163c17a9DecodeWeatherbotEntity(in *jlexer.Lexer, out *WeatherCast) {
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
		case "coord":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Coord = make(map[string]float64)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 float64
					v1 = float64(in.Float64())
					(out.Coord)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "main":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Main = make(map[string]float64)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 float64
					v2 = float64(in.Float64())
					(out.Main)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "wind":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Wind = make(map[string]float64)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 float64
					v3 = float64(in.Float64())
					(out.Wind)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "cod":
			out.ResponseCode = int(in.Int())
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
func easyjson163c17a9EncodeWeatherbotEntity(out *jwriter.Writer, in WeatherCast) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"coord\":"
		out.RawString(prefix[1:])
		if in.Coord == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.Coord {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				out.Float64(float64(v4Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"main\":"
		out.RawString(prefix)
		if in.Main == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.Main {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				out.Float64(float64(v5Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"wind\":"
		out.RawString(prefix)
		if in.Wind == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v6First := true
			for v6Name, v6Value := range in.Wind {
				if v6First {
					v6First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v6Name))
				out.RawByte(':')
				out.Float64(float64(v6Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"cod\":"
		out.RawString(prefix)
		out.Int(int(in.ResponseCode))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WeatherCast) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeWeatherbotEntity(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WeatherCast) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeWeatherbotEntity(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WeatherCast) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeWeatherbotEntity(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WeatherCast) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeWeatherbotEntity(l, v)
}
func easyjson163c17a9DecodeWeatherbotEntity1(in *jlexer.Lexer, out *ForecastUnit) {
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
		case "dt":
			out.Dt = int64(in.Int64())
		case "main":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Main = make(map[string]float64)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 float64
					v7 = float64(in.Float64())
					(out.Main)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
		case "wind":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Wind = make(map[string]float64)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v8 float64
					v8 = float64(in.Float64())
					(out.Wind)[key] = v8
					in.WantComma()
				}
				in.Delim('}')
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
func easyjson163c17a9EncodeWeatherbotEntity1(out *jwriter.Writer, in ForecastUnit) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"dt\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Dt))
	}
	{
		const prefix string = ",\"main\":"
		out.RawString(prefix)
		if in.Main == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v9First := true
			for v9Name, v9Value := range in.Main {
				if v9First {
					v9First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v9Name))
				out.RawByte(':')
				out.Float64(float64(v9Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"wind\":"
		out.RawString(prefix)
		if in.Wind == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v10First := true
			for v10Name, v10Value := range in.Wind {
				if v10First {
					v10First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v10Name))
				out.RawByte(':')
				out.Float64(float64(v10Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ForecastUnit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeWeatherbotEntity1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ForecastUnit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeWeatherbotEntity1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ForecastUnit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeWeatherbotEntity1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ForecastUnit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeWeatherbotEntity1(l, v)
}
func easyjson163c17a9DecodeWeatherbotEntity2(in *jlexer.Lexer, out *Forecast) {
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
		case "cod":
			out.ResponseCode = string(in.String())
		case "list":
			if in.IsNull() {
				in.Skip()
				out.List = nil
			} else {
				in.Delim('[')
				if out.List == nil {
					if !in.IsDelim(']') {
						out.List = make([]ForecastUnit, 0, 2)
					} else {
						out.List = []ForecastUnit{}
					}
				} else {
					out.List = (out.List)[:0]
				}
				for !in.IsDelim(']') {
					var v11 ForecastUnit
					(v11).UnmarshalEasyJSON(in)
					out.List = append(out.List, v11)
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
func easyjson163c17a9EncodeWeatherbotEntity2(out *jwriter.Writer, in Forecast) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cod\":"
		out.RawString(prefix[1:])
		out.String(string(in.ResponseCode))
	}
	{
		const prefix string = ",\"list\":"
		out.RawString(prefix)
		if in.List == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.List {
				if v12 > 0 {
					out.RawByte(',')
				}
				(v13).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Forecast) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeWeatherbotEntity2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Forecast) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeWeatherbotEntity2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Forecast) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeWeatherbotEntity2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Forecast) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeWeatherbotEntity2(l, v)
}
