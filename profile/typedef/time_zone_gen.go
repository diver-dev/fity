// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type TimeZone byte

const (
	TimeZoneAlmaty                   TimeZone = 0
	TimeZoneBangkok                  TimeZone = 1
	TimeZoneBombay                   TimeZone = 2
	TimeZoneBrasilia                 TimeZone = 3
	TimeZoneCairo                    TimeZone = 4
	TimeZoneCapeVerdeIs              TimeZone = 5
	TimeZoneDarwin                   TimeZone = 6
	TimeZoneEniwetok                 TimeZone = 7
	TimeZoneFiji                     TimeZone = 8
	TimeZoneHongKong                 TimeZone = 9
	TimeZoneIslamabad                TimeZone = 10
	TimeZoneKabul                    TimeZone = 11
	TimeZoneMagadan                  TimeZone = 12
	TimeZoneMidAtlantic              TimeZone = 13
	TimeZoneMoscow                   TimeZone = 14
	TimeZoneMuscat                   TimeZone = 15
	TimeZoneNewfoundland             TimeZone = 16
	TimeZoneSamoa                    TimeZone = 17
	TimeZoneSydney                   TimeZone = 18
	TimeZoneTehran                   TimeZone = 19
	TimeZoneTokyo                    TimeZone = 20
	TimeZoneUsAlaska                 TimeZone = 21
	TimeZoneUsAtlantic               TimeZone = 22
	TimeZoneUsCentral                TimeZone = 23
	TimeZoneUsEastern                TimeZone = 24
	TimeZoneUsHawaii                 TimeZone = 25
	TimeZoneUsMountain               TimeZone = 26
	TimeZoneUsPacific                TimeZone = 27
	TimeZoneOther                    TimeZone = 28
	TimeZoneAuckland                 TimeZone = 29
	TimeZoneKathmandu                TimeZone = 30
	TimeZoneEuropeWesternWet         TimeZone = 31
	TimeZoneEuropeCentralCet         TimeZone = 32
	TimeZoneEuropeEasternEet         TimeZone = 33
	TimeZoneJakarta                  TimeZone = 34
	TimeZonePerth                    TimeZone = 35
	TimeZoneAdelaide                 TimeZone = 36
	TimeZoneBrisbane                 TimeZone = 37
	TimeZoneTasmania                 TimeZone = 38
	TimeZoneIceland                  TimeZone = 39
	TimeZoneAmsterdam                TimeZone = 40
	TimeZoneAthens                   TimeZone = 41
	TimeZoneBarcelona                TimeZone = 42
	TimeZoneBerlin                   TimeZone = 43
	TimeZoneBrussels                 TimeZone = 44
	TimeZoneBudapest                 TimeZone = 45
	TimeZoneCopenhagen               TimeZone = 46
	TimeZoneDublin                   TimeZone = 47
	TimeZoneHelsinki                 TimeZone = 48
	TimeZoneLisbon                   TimeZone = 49
	TimeZoneLondon                   TimeZone = 50
	TimeZoneMadrid                   TimeZone = 51
	TimeZoneMunich                   TimeZone = 52
	TimeZoneOslo                     TimeZone = 53
	TimeZoneParis                    TimeZone = 54
	TimeZonePrague                   TimeZone = 55
	TimeZoneReykjavik                TimeZone = 56
	TimeZoneRome                     TimeZone = 57
	TimeZoneStockholm                TimeZone = 58
	TimeZoneVienna                   TimeZone = 59
	TimeZoneWarsaw                   TimeZone = 60
	TimeZoneZurich                   TimeZone = 61
	TimeZoneQuebec                   TimeZone = 62
	TimeZoneOntario                  TimeZone = 63
	TimeZoneManitoba                 TimeZone = 64
	TimeZoneSaskatchewan             TimeZone = 65
	TimeZoneAlberta                  TimeZone = 66
	TimeZoneBritishColumbia          TimeZone = 67
	TimeZoneBoise                    TimeZone = 68
	TimeZoneBoston                   TimeZone = 69
	TimeZoneChicago                  TimeZone = 70
	TimeZoneDallas                   TimeZone = 71
	TimeZoneDenver                   TimeZone = 72
	TimeZoneKansasCity               TimeZone = 73
	TimeZoneLasVegas                 TimeZone = 74
	TimeZoneLosAngeles               TimeZone = 75
	TimeZoneMiami                    TimeZone = 76
	TimeZoneMinneapolis              TimeZone = 77
	TimeZoneNewYork                  TimeZone = 78
	TimeZoneNewOrleans               TimeZone = 79
	TimeZonePhoenix                  TimeZone = 80
	TimeZoneSantaFe                  TimeZone = 81
	TimeZoneSeattle                  TimeZone = 82
	TimeZoneWashingtonDc             TimeZone = 83
	TimeZoneUsArizona                TimeZone = 84
	TimeZoneChita                    TimeZone = 85
	TimeZoneEkaterinburg             TimeZone = 86
	TimeZoneIrkutsk                  TimeZone = 87
	TimeZoneKaliningrad              TimeZone = 88
	TimeZoneKrasnoyarsk              TimeZone = 89
	TimeZoneNovosibirsk              TimeZone = 90
	TimeZonePetropavlovskKamchatskiy TimeZone = 91
	TimeZoneSamara                   TimeZone = 92
	TimeZoneVladivostok              TimeZone = 93
	TimeZoneMexicoCentral            TimeZone = 94
	TimeZoneMexicoMountain           TimeZone = 95
	TimeZoneMexicoPacific            TimeZone = 96
	TimeZoneCapeTown                 TimeZone = 97
	TimeZoneWinkhoek                 TimeZone = 98
	TimeZoneLagos                    TimeZone = 99
	TimeZoneRiyahd                   TimeZone = 100
	TimeZoneVenezuela                TimeZone = 101
	TimeZoneAustraliaLh              TimeZone = 102
	TimeZoneSantiago                 TimeZone = 103
	TimeZoneManual                   TimeZone = 253
	TimeZoneAutomatic                TimeZone = 254
	TimeZoneInvalid                  TimeZone = 0xFF // INVALID
)

var timezonetostrs = map[TimeZone]string{
	TimeZoneAlmaty:                   "almaty",
	TimeZoneBangkok:                  "bangkok",
	TimeZoneBombay:                   "bombay",
	TimeZoneBrasilia:                 "brasilia",
	TimeZoneCairo:                    "cairo",
	TimeZoneCapeVerdeIs:              "cape_verde_is",
	TimeZoneDarwin:                   "darwin",
	TimeZoneEniwetok:                 "eniwetok",
	TimeZoneFiji:                     "fiji",
	TimeZoneHongKong:                 "hong_kong",
	TimeZoneIslamabad:                "islamabad",
	TimeZoneKabul:                    "kabul",
	TimeZoneMagadan:                  "magadan",
	TimeZoneMidAtlantic:              "mid_atlantic",
	TimeZoneMoscow:                   "moscow",
	TimeZoneMuscat:                   "muscat",
	TimeZoneNewfoundland:             "newfoundland",
	TimeZoneSamoa:                    "samoa",
	TimeZoneSydney:                   "sydney",
	TimeZoneTehran:                   "tehran",
	TimeZoneTokyo:                    "tokyo",
	TimeZoneUsAlaska:                 "us_alaska",
	TimeZoneUsAtlantic:               "us_atlantic",
	TimeZoneUsCentral:                "us_central",
	TimeZoneUsEastern:                "us_eastern",
	TimeZoneUsHawaii:                 "us_hawaii",
	TimeZoneUsMountain:               "us_mountain",
	TimeZoneUsPacific:                "us_pacific",
	TimeZoneOther:                    "other",
	TimeZoneAuckland:                 "auckland",
	TimeZoneKathmandu:                "kathmandu",
	TimeZoneEuropeWesternWet:         "europe_western_wet",
	TimeZoneEuropeCentralCet:         "europe_central_cet",
	TimeZoneEuropeEasternEet:         "europe_eastern_eet",
	TimeZoneJakarta:                  "jakarta",
	TimeZonePerth:                    "perth",
	TimeZoneAdelaide:                 "adelaide",
	TimeZoneBrisbane:                 "brisbane",
	TimeZoneTasmania:                 "tasmania",
	TimeZoneIceland:                  "iceland",
	TimeZoneAmsterdam:                "amsterdam",
	TimeZoneAthens:                   "athens",
	TimeZoneBarcelona:                "barcelona",
	TimeZoneBerlin:                   "berlin",
	TimeZoneBrussels:                 "brussels",
	TimeZoneBudapest:                 "budapest",
	TimeZoneCopenhagen:               "copenhagen",
	TimeZoneDublin:                   "dublin",
	TimeZoneHelsinki:                 "helsinki",
	TimeZoneLisbon:                   "lisbon",
	TimeZoneLondon:                   "london",
	TimeZoneMadrid:                   "madrid",
	TimeZoneMunich:                   "munich",
	TimeZoneOslo:                     "oslo",
	TimeZoneParis:                    "paris",
	TimeZonePrague:                   "prague",
	TimeZoneReykjavik:                "reykjavik",
	TimeZoneRome:                     "rome",
	TimeZoneStockholm:                "stockholm",
	TimeZoneVienna:                   "vienna",
	TimeZoneWarsaw:                   "warsaw",
	TimeZoneZurich:                   "zurich",
	TimeZoneQuebec:                   "quebec",
	TimeZoneOntario:                  "ontario",
	TimeZoneManitoba:                 "manitoba",
	TimeZoneSaskatchewan:             "saskatchewan",
	TimeZoneAlberta:                  "alberta",
	TimeZoneBritishColumbia:          "british_columbia",
	TimeZoneBoise:                    "boise",
	TimeZoneBoston:                   "boston",
	TimeZoneChicago:                  "chicago",
	TimeZoneDallas:                   "dallas",
	TimeZoneDenver:                   "denver",
	TimeZoneKansasCity:               "kansas_city",
	TimeZoneLasVegas:                 "las_vegas",
	TimeZoneLosAngeles:               "los_angeles",
	TimeZoneMiami:                    "miami",
	TimeZoneMinneapolis:              "minneapolis",
	TimeZoneNewYork:                  "new_york",
	TimeZoneNewOrleans:               "new_orleans",
	TimeZonePhoenix:                  "phoenix",
	TimeZoneSantaFe:                  "santa_fe",
	TimeZoneSeattle:                  "seattle",
	TimeZoneWashingtonDc:             "washington_dc",
	TimeZoneUsArizona:                "us_arizona",
	TimeZoneChita:                    "chita",
	TimeZoneEkaterinburg:             "ekaterinburg",
	TimeZoneIrkutsk:                  "irkutsk",
	TimeZoneKaliningrad:              "kaliningrad",
	TimeZoneKrasnoyarsk:              "krasnoyarsk",
	TimeZoneNovosibirsk:              "novosibirsk",
	TimeZonePetropavlovskKamchatskiy: "petropavlovsk_kamchatskiy",
	TimeZoneSamara:                   "samara",
	TimeZoneVladivostok:              "vladivostok",
	TimeZoneMexicoCentral:            "mexico_central",
	TimeZoneMexicoMountain:           "mexico_mountain",
	TimeZoneMexicoPacific:            "mexico_pacific",
	TimeZoneCapeTown:                 "cape_town",
	TimeZoneWinkhoek:                 "winkhoek",
	TimeZoneLagos:                    "lagos",
	TimeZoneRiyahd:                   "riyahd",
	TimeZoneVenezuela:                "venezuela",
	TimeZoneAustraliaLh:              "australia_lh",
	TimeZoneSantiago:                 "santiago",
	TimeZoneManual:                   "manual",
	TimeZoneAutomatic:                "automatic",
	TimeZoneInvalid:                  "invalid",
}

func (t TimeZone) String() string {
	val, ok := timezonetostrs[t]
	if !ok {
		return strconv.Itoa(int(t))
	}
	return val
}

var strtotimezone = func() map[string]TimeZone {
	m := make(map[string]TimeZone)
	for t, str := range timezonetostrs {
		m[str] = TimeZone(t)
	}
	return m
}()

// FromString parse string into TimeZone constant it's represent, return TimeZoneInvalid if not found.
func TimeZoneFromString(s string) TimeZone {
	val, ok := strtotimezone[s]
	if !ok {
		return strtotimezone["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListTimeZone() []TimeZone {
	vs := make([]TimeZone, 0, len(timezonetostrs))
	for i := range timezonetostrs {
		vs = append(vs, TimeZone(i))
	}
	return vs
}
