/*
 * Bamboo - A Vietnamese Input method editor
 * Copyright (C) Luong Thanh Lam <ltlam93@gmail.com>
 *
 * This software is licensed under the MIT license. For more information,
 * see <https://github.com/BambooEngine/bamboo-core/blob/master/LISENCE>.
 */
package bamboo

type InputMethodDefinition map[string]string

var InputMethodDefinitions = map[string]InputMethodDefinition{
	"Telex": {
		"z": "XoaDauThanh",
		"s": "DauSac",
		"f": "DauHuyen",
		"r": "DauHoi",
		"x": "DauNga",
		"j": "DauNang",
		"a": "A_Â",
		"e": "E_Ê",
		"o": "O_Ô",
		"w": "UOA_ƯƠĂ",
		"d": "D_Đ",
	},
	"VNI": {
		"0": "XoaDauThanh",
		"1": "DauSac",
		"2": "DauHuyen",
		"3": "DauHoi",
		"4": "DauNga",
		"5": "DauNang",
		"6": "AEO_ÂÊÔ",
		"7": "UO_ƯƠ",
		"8": "A_Ă",
		"9": "D_Đ",
	},

	"VIQR": {
		"0":  "XoaDauThanh",
		"'":  "DauSac",
		"`":  "DauHuyen",
		"?":  "DauHoi",
		"~":  "DauNga",
		".":  "DauNang",
		"^":  "AEO_ÂÊÔ",
		"+":  "UO_ƯƠ",
		"*":  "UO_ƯƠ",
		"(":  "A_Ă",
		"\\": "D_Đ",
	},
	"Microsoft layout": {
		"8": "DauSac",
		"5": "DauHuyen",
		"6": "DauHoi",
		"7": "DauNga",
		"9": "DauNang",
		"1": "__ă",
		"!": "_Ă",
		"2": "__â",
		"@": "_Â",
		"3": "__ê",
		"#": "_Ê",
		"4": "__ô",
		"$": "_Ô",
		"0": "__đ",
		")": "_Đ",
		"[": "__ư",
		"{": "_Ư",
		"]": "__ơ",
		"}": "_Ơ",
	},
	"Telex 2": {
		"z": "XoaDauThanh",
		"s": "DauSac",
		"f": "DauHuyen",
		"r": "DauHoi",
		"x": "DauNga",
		"j": "DauNang",
		"a": "A_Â",
		"e": "E_Ê",
		"o": "O_Ô",
		"w": "UOA_ƯƠĂ__Ư",
		"d": "D_Đ",
		"]": "__ư",
		"[": "__ơ",
		"}": "_Ư",
		"{": "_Ơ",
	},
	"Telex + VNI": {
		"z": "XoaDauThanh",
		"s": "DauSac",
		"f": "DauHuyen",
		"r": "DauHoi",
		"x": "DauNga",
		"j": "DauNang",
		"a": "A_Â",
		"e": "E_Ê",
		"o": "O_Ô",
		"w": "UOA_ƯƠĂ",
		"d": "D_Đ",
		"0": "XoaDauThanh",
		"1": "DauSac",
		"2": "DauHuyen",
		"3": "DauHoi",
		"4": "DauNga",
		"5": "DauNang",
		"6": "AEO_ÂÊÔ",
		"7": "UO_ƯƠ",
		"8": "A_Ă",
		"9": "D_Đ",
	},
	"Telex + VNI + VIQR": {
		"z":  "XoaDauThanh",
		"s":  "DauSac",
		"f":  "DauHuyen",
		"r":  "DauHoi",
		"x":  "DauNga",
		"j":  "DauNang",
		"a":  "A_Â",
		"e":  "E_Ê",
		"o":  "O_Ô",
		"w":  "UOA_ƯƠĂ",
		"d":  "D_Đ",
		"0":  "XoaDauThanh",
		"1":  "DauSac",
		"2":  "DauHuyen",
		"3":  "DauHoi",
		"4":  "DauNga",
		"5":  "DauNang",
		"6":  "AEO_ÂÊÔ",
		"7":  "UO_ƯƠ",
		"8":  "A_Ă",
		"9":  "D_Đ",
		"'":  "DauSac",
		"`":  "DauHuyen",
		"?":  "DauHoi",
		"~":  "DauNga",
		".":  "DauNang",
		"^":  "AEO_ÂÊÔ",
		"+":  "UO_ƯƠ",
		"*":  "UO_ƯƠ",
		"(":  "A_Ă",
		"\\": "D_Đ",
	},
	"VNI Bàn phím tiếng Pháp": {
		"&":  "XoaDauThanh",
		"é":  "DauSac",
		"\"": "DauHuyen",
		"'":  "DauHoi",
		"(":  "DauNga",
		"-":  "DauNang",
		"è":  "AEO_ÂÊÔ",
		"_":  "UO_ƯƠ",
		"ç":  "A_Ă",
		"à":  "D_Đ",
	},
	"Telex W": {
		"z": "XoaDauThanh",
		"s": "DauSac",
		"f": "DauHuyen",
		"r": "DauHoi",
		"x": "DauNga",
		"j": "DauNang",
		"a": "A_Â",
		"e": "E_Ê",
		"o": "O_Ô",
		"w": "UOA_ƯƠĂ__Ư",
		"d": "D_Đ",
	},
}

func GetInputMethodDefinitions() map[string]InputMethodDefinition {
	var t = make(map[string]InputMethodDefinition)
	for k, v := range InputMethodDefinitions {
		t[k] = v
	}
	return t
}
