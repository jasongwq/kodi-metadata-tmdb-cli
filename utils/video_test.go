package utils

import (
	"testing"
)

func TestIsFormat(t *testing.T) {
	unit := map[string]string{
		"720":        "",
		"a720p":      "720p",
		"720P":       "720P",
		"1080p":      "1080p",
		"1080P":      "1080P",
		"2k":         "2k",
		"2K":         "2K",
		"4K":         "4K",
		"720p-CMCT":  "720p",
		"-720p-CMCT": "720p",
	}

	for k, v := range unit {
		actual := IsFormat(k)
		if actual != v {
			t.Errorf("isFormat(%s) = %s; expected %s", k, actual, v)
		}
	}
}

func TestIsSeason(t *testing.T) {
	unit := map[string]string{
		"s01":  "s01",
		"S01":  "S01",
		"s1":   "s1",
		"S1":   "S1",
		"S100": "S100",
		"4K":   "",
		"Fall.in.Love.2021.WEB-DL.4k.H265.10bit.AAC-HDCTV FallinLove ": "",
		"Hawkeye.2021S01.Never.Meet.Your.Heroes.2160p":                 "S01",
	}

	for k, v := range unit {
		actual := IsSeason(k)
		if actual != v {
			t.Errorf("isSeason(%s) = %s; expected %s", k, actual, v)
		}
	}
}

func TestSplit(t *testing.T) {
	unit := map[string][]string{
		"[梦蓝字幕组]Crayonshinchan 蜡笔小新[1105][2021.11.06][AVC][1080P][GB_JP][MP4]V2.mp4": []string{
			"梦蓝字幕组",
			"Crayonshinchan",
			"蜡笔小新",
			"1105",
			"2021",
			"11",
			"06",
			"AVC",
		},
		"The Last Son 2021.mkv": []string{
			"The",
			"Last",
			"Son",
			"2021",
			"mkv",
		},
		"Midway 2019 2160p CAN UHD Blu-ray HEVC DTS-HD MA 5.1-THDBST@HDSky.nfo": []string{
			"Midway",
			"2019",
			"2160p",
			"CAN",
			"UHD",
			"Blu",
			"ray",
			"HEVC",
			"DTS",
			"HD",
			"MA",
			"5",
			"1",
			"THDBST",
			"HDSky",
			"nfo",
		},
	}

	for k, v := range unit {
		actual := Split(k)
		for k2, v2 := range v {
			if actual[k2] != v2 {
				t.Errorf("Split(%s) = %s; expected %s", k, actual[k2], v2)
			}
		}
	}
}
