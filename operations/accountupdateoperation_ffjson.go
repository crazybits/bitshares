// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: accountupdateoperation.go

package operations

import (
	"bytes"
	"fmt"
	"github.com/denkhaus/bitshares/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AccountUpdateOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *AccountUpdateOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "account":`)

	{

		obj, err = j.Account.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte(',')
	if j.Active != nil {
		if true {
			buf.WriteString(`"active":`)

			{

				err = j.Active.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	buf.WriteString(`"extensions":`)

	{

		obj, err = j.Extensions.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"fee":`)

	{

		err = j.Fee.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteByte(',')
	if j.NewOptions != nil {
		if true {
			buf.WriteString(`"new_options":`)

			{

				err = j.NewOptions.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if j.Owner != nil {
		if true {
			buf.WriteString(`"owner":`)

			{

				err = j.Owner.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAccountUpdateOperationbase = iota
	ffjtAccountUpdateOperationnosuchkey

	ffjtAccountUpdateOperationAccount

	ffjtAccountUpdateOperationActive

	ffjtAccountUpdateOperationExtensions

	ffjtAccountUpdateOperationFee

	ffjtAccountUpdateOperationNewOptions

	ffjtAccountUpdateOperationOwner
)

var ffjKeyAccountUpdateOperationAccount = []byte("account")

var ffjKeyAccountUpdateOperationActive = []byte("active")

var ffjKeyAccountUpdateOperationExtensions = []byte("extensions")

var ffjKeyAccountUpdateOperationFee = []byte("fee")

var ffjKeyAccountUpdateOperationNewOptions = []byte("new_options")

var ffjKeyAccountUpdateOperationOwner = []byte("owner")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AccountUpdateOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AccountUpdateOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAccountUpdateOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtAccountUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAccountUpdateOperationAccount, kn) {
						currentKey = ffjtAccountUpdateOperationAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountUpdateOperationActive, kn) {
						currentKey = ffjtAccountUpdateOperationActive
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyAccountUpdateOperationExtensions, kn) {
						currentKey = ffjtAccountUpdateOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyAccountUpdateOperationFee, kn) {
						currentKey = ffjtAccountUpdateOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyAccountUpdateOperationNewOptions, kn) {
						currentKey = ffjtAccountUpdateOperationNewOptions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyAccountUpdateOperationOwner, kn) {
						currentKey = ffjtAccountUpdateOperationOwner
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountUpdateOperationOwner, kn) {
					currentKey = ffjtAccountUpdateOperationOwner
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountUpdateOperationNewOptions, kn) {
					currentKey = ffjtAccountUpdateOperationNewOptions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountUpdateOperationFee, kn) {
					currentKey = ffjtAccountUpdateOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountUpdateOperationExtensions, kn) {
					currentKey = ffjtAccountUpdateOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountUpdateOperationActive, kn) {
					currentKey = ffjtAccountUpdateOperationActive
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountUpdateOperationAccount, kn) {
					currentKey = ffjtAccountUpdateOperationAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAccountUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtAccountUpdateOperationAccount:
					goto handle_Account

				case ffjtAccountUpdateOperationActive:
					goto handle_Active

				case ffjtAccountUpdateOperationExtensions:
					goto handle_Extensions

				case ffjtAccountUpdateOperationFee:
					goto handle_Fee

				case ffjtAccountUpdateOperationNewOptions:
					goto handle_NewOptions

				case ffjtAccountUpdateOperationOwner:
					goto handle_Owner

				case ffjtAccountUpdateOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Account:

	/* handler: j.Account type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Account.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Active:

	/* handler: j.Active type=types.Authority kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.Active = nil

		} else {

			if j.Active == nil {
				j.Active = new(types.Authority)
			}

			err = j.Active.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=types.AccountUpdateExtensions kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.Extensions.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.Fee.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NewOptions:

	/* handler: j.NewOptions type=types.AccountOptions kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.NewOptions = nil

		} else {

			if j.NewOptions == nil {
				j.NewOptions = new(types.AccountOptions)
			}

			err = j.NewOptions.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Owner:

	/* handler: j.Owner type=types.Authority kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.Owner = nil

		} else {

			if j.Owner == nil {
				j.Owner = new(types.Authority)
			}

			err = j.Owner.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
