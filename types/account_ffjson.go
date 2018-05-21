// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: account.go

package types

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *Account) MarshalJSON() ([]byte, error) {
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
func (j *Account) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"id":`)

	{

		obj, err = j.ID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"name":`)
	fflib.WriteJsonString(buf, string(j.Name))
	buf.WriteString(`,"statistics":`)

	{

		obj, err = j.Statistics.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"membership_expiration_date":`)

	{

		obj, err = j.MembershipExpirationDate.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"network_fee_percentage":`)
	fflib.FormatBits2(buf, uint64(j.NetworkFeePercentage), 10, false)
	buf.WriteString(`,"lifetime_referrer_fee_percentage":`)
	fflib.FormatBits2(buf, uint64(j.LifetimeReferrerFeePercentage), 10, false)
	buf.WriteString(`,"referrer_rewards_percentage":`)
	fflib.FormatBits2(buf, uint64(j.ReferrerRewardsPercentage), 10, false)
	buf.WriteString(`,"top_n_control_flags":`)
	fflib.FormatBits2(buf, uint64(j.TopNControlFlags), 10, false)
	buf.WriteString(`,"whitelisting_accounts":`)
	if j.WhitelistingAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.WhitelistingAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"blacklisting_accounts":`)
	if j.BlacklistingAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.BlacklistingAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"whitelisted_accounts":`)
	if j.WhitelistedAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.WhitelistedAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"blacklisted_accounts":`)
	if j.BlacklistedAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.BlacklistedAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"options":`)

	{

		err = j.Options.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"registrar":`)

	{

		obj, err = j.Registrar.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"referrer":`)

	{

		obj, err = j.Referrer.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"lifetime_referrer":`)

	{

		obj, err = j.LifetimeReferrer.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"cashback_vb":`)

	{

		obj, err = j.CashbackVB.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"owner":`)

	{

		err = j.Owner.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"active":`)

	{

		err = j.Active.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"owner_special_authority":`)

	{

		obj, err = j.OwnerSpecialAuthority.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"active_special_authority":`)

	{

		obj, err = j.ActiveSpecialAuthority.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAccountbase = iota
	ffjtAccountnosuchkey

	ffjtAccountID

	ffjtAccountName

	ffjtAccountStatistics

	ffjtAccountMembershipExpirationDate

	ffjtAccountNetworkFeePercentage

	ffjtAccountLifetimeReferrerFeePercentage

	ffjtAccountReferrerRewardsPercentage

	ffjtAccountTopNControlFlags

	ffjtAccountWhitelistingAccounts

	ffjtAccountBlacklistingAccounts

	ffjtAccountWhitelistedAccounts

	ffjtAccountBlacklistedAccounts

	ffjtAccountOptions

	ffjtAccountRegistrar

	ffjtAccountReferrer

	ffjtAccountLifetimeReferrer

	ffjtAccountCashbackVB

	ffjtAccountOwner

	ffjtAccountActive

	ffjtAccountOwnerSpecialAuthority

	ffjtAccountActiveSpecialAuthority
)

var ffjKeyAccountID = []byte("id")

var ffjKeyAccountName = []byte("name")

var ffjKeyAccountStatistics = []byte("statistics")

var ffjKeyAccountMembershipExpirationDate = []byte("membership_expiration_date")

var ffjKeyAccountNetworkFeePercentage = []byte("network_fee_percentage")

var ffjKeyAccountLifetimeReferrerFeePercentage = []byte("lifetime_referrer_fee_percentage")

var ffjKeyAccountReferrerRewardsPercentage = []byte("referrer_rewards_percentage")

var ffjKeyAccountTopNControlFlags = []byte("top_n_control_flags")

var ffjKeyAccountWhitelistingAccounts = []byte("whitelisting_accounts")

var ffjKeyAccountBlacklistingAccounts = []byte("blacklisting_accounts")

var ffjKeyAccountWhitelistedAccounts = []byte("whitelisted_accounts")

var ffjKeyAccountBlacklistedAccounts = []byte("blacklisted_accounts")

var ffjKeyAccountOptions = []byte("options")

var ffjKeyAccountRegistrar = []byte("registrar")

var ffjKeyAccountReferrer = []byte("referrer")

var ffjKeyAccountLifetimeReferrer = []byte("lifetime_referrer")

var ffjKeyAccountCashbackVB = []byte("cashback_vb")

var ffjKeyAccountOwner = []byte("owner")

var ffjKeyAccountActive = []byte("active")

var ffjKeyAccountOwnerSpecialAuthority = []byte("owner_special_authority")

var ffjKeyAccountActiveSpecialAuthority = []byte("active_special_authority")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Account) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Account) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAccountbase
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
				currentKey = ffjtAccountnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAccountActive, kn) {
						currentKey = ffjtAccountActive
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountActiveSpecialAuthority, kn) {
						currentKey = ffjtAccountActiveSpecialAuthority
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffjKeyAccountBlacklistingAccounts, kn) {
						currentKey = ffjtAccountBlacklistingAccounts
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountBlacklistedAccounts, kn) {
						currentKey = ffjtAccountBlacklistedAccounts
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyAccountCashbackVB, kn) {
						currentKey = ffjtAccountCashbackVB
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyAccountID, kn) {
						currentKey = ffjtAccountID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffjKeyAccountLifetimeReferrerFeePercentage, kn) {
						currentKey = ffjtAccountLifetimeReferrerFeePercentage
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountLifetimeReferrer, kn) {
						currentKey = ffjtAccountLifetimeReferrer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyAccountMembershipExpirationDate, kn) {
						currentKey = ffjtAccountMembershipExpirationDate
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyAccountName, kn) {
						currentKey = ffjtAccountName
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountNetworkFeePercentage, kn) {
						currentKey = ffjtAccountNetworkFeePercentage
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyAccountOptions, kn) {
						currentKey = ffjtAccountOptions
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountOwner, kn) {
						currentKey = ffjtAccountOwner
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountOwnerSpecialAuthority, kn) {
						currentKey = ffjtAccountOwnerSpecialAuthority
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyAccountReferrerRewardsPercentage, kn) {
						currentKey = ffjtAccountReferrerRewardsPercentage
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountRegistrar, kn) {
						currentKey = ffjtAccountRegistrar
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountReferrer, kn) {
						currentKey = ffjtAccountReferrer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyAccountStatistics, kn) {
						currentKey = ffjtAccountStatistics
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyAccountTopNControlFlags, kn) {
						currentKey = ffjtAccountTopNControlFlags
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyAccountWhitelistingAccounts, kn) {
						currentKey = ffjtAccountWhitelistingAccounts
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountWhitelistedAccounts, kn) {
						currentKey = ffjtAccountWhitelistedAccounts
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyAccountActiveSpecialAuthority, kn) {
					currentKey = ffjtAccountActiveSpecialAuthority
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountOwnerSpecialAuthority, kn) {
					currentKey = ffjtAccountOwnerSpecialAuthority
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountActive, kn) {
					currentKey = ffjtAccountActive
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountOwner, kn) {
					currentKey = ffjtAccountOwner
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountCashbackVB, kn) {
					currentKey = ffjtAccountCashbackVB
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAccountLifetimeReferrer, kn) {
					currentKey = ffjtAccountLifetimeReferrer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountReferrer, kn) {
					currentKey = ffjtAccountReferrer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountRegistrar, kn) {
					currentKey = ffjtAccountRegistrar
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountOptions, kn) {
					currentKey = ffjtAccountOptions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountBlacklistedAccounts, kn) {
					currentKey = ffjtAccountBlacklistedAccounts
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountWhitelistedAccounts, kn) {
					currentKey = ffjtAccountWhitelistedAccounts
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountBlacklistingAccounts, kn) {
					currentKey = ffjtAccountBlacklistingAccounts
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountWhitelistingAccounts, kn) {
					currentKey = ffjtAccountWhitelistingAccounts
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountTopNControlFlags, kn) {
					currentKey = ffjtAccountTopNControlFlags
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountReferrerRewardsPercentage, kn) {
					currentKey = ffjtAccountReferrerRewardsPercentage
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAccountLifetimeReferrerFeePercentage, kn) {
					currentKey = ffjtAccountLifetimeReferrerFeePercentage
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountNetworkFeePercentage, kn) {
					currentKey = ffjtAccountNetworkFeePercentage
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountMembershipExpirationDate, kn) {
					currentKey = ffjtAccountMembershipExpirationDate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountStatistics, kn) {
					currentKey = ffjtAccountStatistics
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountName, kn) {
					currentKey = ffjtAccountName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountID, kn) {
					currentKey = ffjtAccountID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAccountnosuchkey
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

				case ffjtAccountID:
					goto handle_ID

				case ffjtAccountName:
					goto handle_Name

				case ffjtAccountStatistics:
					goto handle_Statistics

				case ffjtAccountMembershipExpirationDate:
					goto handle_MembershipExpirationDate

				case ffjtAccountNetworkFeePercentage:
					goto handle_NetworkFeePercentage

				case ffjtAccountLifetimeReferrerFeePercentage:
					goto handle_LifetimeReferrerFeePercentage

				case ffjtAccountReferrerRewardsPercentage:
					goto handle_ReferrerRewardsPercentage

				case ffjtAccountTopNControlFlags:
					goto handle_TopNControlFlags

				case ffjtAccountWhitelistingAccounts:
					goto handle_WhitelistingAccounts

				case ffjtAccountBlacklistingAccounts:
					goto handle_BlacklistingAccounts

				case ffjtAccountWhitelistedAccounts:
					goto handle_WhitelistedAccounts

				case ffjtAccountBlacklistedAccounts:
					goto handle_BlacklistedAccounts

				case ffjtAccountOptions:
					goto handle_Options

				case ffjtAccountRegistrar:
					goto handle_Registrar

				case ffjtAccountReferrer:
					goto handle_Referrer

				case ffjtAccountLifetimeReferrer:
					goto handle_LifetimeReferrer

				case ffjtAccountCashbackVB:
					goto handle_CashbackVB

				case ffjtAccountOwner:
					goto handle_Owner

				case ffjtAccountActive:
					goto handle_Active

				case ffjtAccountOwnerSpecialAuthority:
					goto handle_OwnerSpecialAuthority

				case ffjtAccountActiveSpecialAuthority:
					goto handle_ActiveSpecialAuthority

				case ffjtAccountnosuchkey:
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

handle_ID:

	/* handler: j.ID type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ID.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Name:

	/* handler: j.Name type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Name = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Statistics:

	/* handler: j.Statistics type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Statistics.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MembershipExpirationDate:

	/* handler: j.MembershipExpirationDate type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.MembershipExpirationDate.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NetworkFeePercentage:

	/* handler: j.NetworkFeePercentage type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.NetworkFeePercentage.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LifetimeReferrerFeePercentage:

	/* handler: j.LifetimeReferrerFeePercentage type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.LifetimeReferrerFeePercentage.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ReferrerRewardsPercentage:

	/* handler: j.ReferrerRewardsPercentage type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ReferrerRewardsPercentage.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TopNControlFlags:

	/* handler: j.TopNControlFlags type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.TopNControlFlags.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WhitelistingAccounts:

	/* handler: j.WhitelistingAccounts type=types.GrapheneIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for GrapheneIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.WhitelistingAccounts = nil
		} else {

			j.WhitelistingAccounts = []GrapheneID{}

			wantVal := true

			for {

				var tmpJWhitelistingAccounts GrapheneID

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJWhitelistingAccounts type=types.GrapheneID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJWhitelistingAccounts.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.WhitelistingAccounts = append(j.WhitelistingAccounts, tmpJWhitelistingAccounts)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BlacklistingAccounts:

	/* handler: j.BlacklistingAccounts type=types.GrapheneIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for GrapheneIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BlacklistingAccounts = nil
		} else {

			j.BlacklistingAccounts = []GrapheneID{}

			wantVal := true

			for {

				var tmpJBlacklistingAccounts GrapheneID

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBlacklistingAccounts type=types.GrapheneID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJBlacklistingAccounts.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.BlacklistingAccounts = append(j.BlacklistingAccounts, tmpJBlacklistingAccounts)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WhitelistedAccounts:

	/* handler: j.WhitelistedAccounts type=types.GrapheneIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for GrapheneIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.WhitelistedAccounts = nil
		} else {

			j.WhitelistedAccounts = []GrapheneID{}

			wantVal := true

			for {

				var tmpJWhitelistedAccounts GrapheneID

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJWhitelistedAccounts type=types.GrapheneID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJWhitelistedAccounts.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.WhitelistedAccounts = append(j.WhitelistedAccounts, tmpJWhitelistedAccounts)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BlacklistedAccounts:

	/* handler: j.BlacklistedAccounts type=types.GrapheneIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for GrapheneIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BlacklistedAccounts = nil
		} else {

			j.BlacklistedAccounts = []GrapheneID{}

			wantVal := true

			for {

				var tmpJBlacklistedAccounts GrapheneID

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBlacklistedAccounts type=types.GrapheneID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJBlacklistedAccounts.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.BlacklistedAccounts = append(j.BlacklistedAccounts, tmpJBlacklistedAccounts)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Options:

	/* handler: j.Options type=types.AccountOptions kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.Options.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Registrar:

	/* handler: j.Registrar type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Registrar.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Referrer:

	/* handler: j.Referrer type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Referrer.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LifetimeReferrer:

	/* handler: j.LifetimeReferrer type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.LifetimeReferrer.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CashbackVB:

	/* handler: j.CashbackVB type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.CashbackVB.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
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

		} else {

			err = j.Owner.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
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

		} else {

			err = j.Active.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_OwnerSpecialAuthority:

	/* handler: j.OwnerSpecialAuthority type=types.OwnerSpecialAuthority kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.OwnerSpecialAuthority.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ActiveSpecialAuthority:

	/* handler: j.ActiveSpecialAuthority type=types.ActiveSpecialAuthority kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ActiveSpecialAuthority.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
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
