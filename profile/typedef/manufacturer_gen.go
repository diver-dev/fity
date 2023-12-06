// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Manufacturer uint16

const (
	ManufacturerGarmin                 Manufacturer = 1
	ManufacturerGarminFr405Antfs       Manufacturer = 2 // Do not use. Used by FR405 for ANTFS man id.
	ManufacturerZephyr                 Manufacturer = 3
	ManufacturerDayton                 Manufacturer = 4
	ManufacturerIdt                    Manufacturer = 5
	ManufacturerSrm                    Manufacturer = 6
	ManufacturerQuarq                  Manufacturer = 7
	ManufacturerIbike                  Manufacturer = 8
	ManufacturerSaris                  Manufacturer = 9
	ManufacturerSparkHk                Manufacturer = 10
	ManufacturerTanita                 Manufacturer = 11
	ManufacturerEchowell               Manufacturer = 12
	ManufacturerDynastreamOem          Manufacturer = 13
	ManufacturerNautilus               Manufacturer = 14
	ManufacturerDynastream             Manufacturer = 15
	ManufacturerTimex                  Manufacturer = 16
	ManufacturerMetrigear              Manufacturer = 17
	ManufacturerXelic                  Manufacturer = 18
	ManufacturerBeurer                 Manufacturer = 19
	ManufacturerCardiosport            Manufacturer = 20
	ManufacturerAAndD                  Manufacturer = 21
	ManufacturerHmm                    Manufacturer = 22
	ManufacturerSuunto                 Manufacturer = 23
	ManufacturerThitaElektronik        Manufacturer = 24
	ManufacturerGpulse                 Manufacturer = 25
	ManufacturerCleanMobile            Manufacturer = 26
	ManufacturerPedalBrain             Manufacturer = 27
	ManufacturerPeaksware              Manufacturer = 28
	ManufacturerSaxonar                Manufacturer = 29
	ManufacturerLemondFitness          Manufacturer = 30
	ManufacturerDexcom                 Manufacturer = 31
	ManufacturerWahooFitness           Manufacturer = 32
	ManufacturerOctaneFitness          Manufacturer = 33
	ManufacturerArchinoetics           Manufacturer = 34
	ManufacturerTheHurtBox             Manufacturer = 35
	ManufacturerCitizenSystems         Manufacturer = 36
	ManufacturerMagellan               Manufacturer = 37
	ManufacturerOsynce                 Manufacturer = 38
	ManufacturerHolux                  Manufacturer = 39
	ManufacturerConcept2               Manufacturer = 40
	ManufacturerShimano                Manufacturer = 41
	ManufacturerOneGiantLeap           Manufacturer = 42
	ManufacturerAceSensor              Manufacturer = 43
	ManufacturerBrimBrothers           Manufacturer = 44
	ManufacturerXplova                 Manufacturer = 45
	ManufacturerPerceptionDigital      Manufacturer = 46
	ManufacturerBf1systems             Manufacturer = 47
	ManufacturerPioneer                Manufacturer = 48
	ManufacturerSpantec                Manufacturer = 49
	ManufacturerMetalogics             Manufacturer = 50
	Manufacturer4Iiiis                 Manufacturer = 51
	ManufacturerSeikoEpson             Manufacturer = 52
	ManufacturerSeikoEpsonOem          Manufacturer = 53
	ManufacturerIforPowell             Manufacturer = 54
	ManufacturerMaxwellGuider          Manufacturer = 55
	ManufacturerStarTrac               Manufacturer = 56
	ManufacturerBreakaway              Manufacturer = 57
	ManufacturerAlatechTechnologyLtd   Manufacturer = 58
	ManufacturerMioTechnologyEurope    Manufacturer = 59
	ManufacturerRotor                  Manufacturer = 60
	ManufacturerGeonaute               Manufacturer = 61
	ManufacturerIdBike                 Manufacturer = 62
	ManufacturerSpecialized            Manufacturer = 63
	ManufacturerWtek                   Manufacturer = 64
	ManufacturerPhysicalEnterprises    Manufacturer = 65
	ManufacturerNorthPoleEngineering   Manufacturer = 66
	ManufacturerBkool                  Manufacturer = 67
	ManufacturerCateye                 Manufacturer = 68
	ManufacturerStagesCycling          Manufacturer = 69
	ManufacturerSigmasport             Manufacturer = 70
	ManufacturerTomtom                 Manufacturer = 71
	ManufacturerPeripedal              Manufacturer = 72
	ManufacturerWattbike               Manufacturer = 73
	ManufacturerMoxy                   Manufacturer = 76
	ManufacturerCiclosport             Manufacturer = 77
	ManufacturerPowerbahn              Manufacturer = 78
	ManufacturerAcornProjectsAps       Manufacturer = 79
	ManufacturerLifebeam               Manufacturer = 80
	ManufacturerBontrager              Manufacturer = 81
	ManufacturerWellgo                 Manufacturer = 82
	ManufacturerScosche                Manufacturer = 83
	ManufacturerMagura                 Manufacturer = 84
	ManufacturerWoodway                Manufacturer = 85
	ManufacturerElite                  Manufacturer = 86
	ManufacturerNielsenKellerman       Manufacturer = 87
	ManufacturerDkCity                 Manufacturer = 88
	ManufacturerTacx                   Manufacturer = 89
	ManufacturerDirectionTechnology    Manufacturer = 90
	ManufacturerMagtonic               Manufacturer = 91
	Manufacturer1Partcarbon            Manufacturer = 92
	ManufacturerInsideRideTechnologies Manufacturer = 93
	ManufacturerSoundOfMotion          Manufacturer = 94
	ManufacturerStryd                  Manufacturer = 95
	ManufacturerIcg                    Manufacturer = 96 // Indoorcycling Group
	ManufacturerMipulse                Manufacturer = 97
	ManufacturerBsxAthletics           Manufacturer = 98
	ManufacturerLook                   Manufacturer = 99
	ManufacturerCampagnoloSrl          Manufacturer = 100
	ManufacturerBodyBikeSmart          Manufacturer = 101
	ManufacturerPraxisworks            Manufacturer = 102
	ManufacturerLimitsTechnology       Manufacturer = 103 // Limits Technology Ltd.
	ManufacturerTopactionTechnology    Manufacturer = 104 // TopAction Technology Inc.
	ManufacturerCosinuss               Manufacturer = 105
	ManufacturerFitcare                Manufacturer = 106
	ManufacturerMagene                 Manufacturer = 107
	ManufacturerGiantManufacturingCo   Manufacturer = 108
	ManufacturerTigrasport             Manufacturer = 109 // Tigrasport
	ManufacturerSalutron               Manufacturer = 110
	ManufacturerTechnogym              Manufacturer = 111
	ManufacturerBrytonSensors          Manufacturer = 112
	ManufacturerLatitudeLimited        Manufacturer = 113
	ManufacturerSoaringTechnology      Manufacturer = 114
	ManufacturerIgpsport               Manufacturer = 115
	ManufacturerThinkrider             Manufacturer = 116
	ManufacturerGopherSport            Manufacturer = 117
	ManufacturerWaterrower             Manufacturer = 118
	ManufacturerOrangetheory           Manufacturer = 119
	ManufacturerInpeak                 Manufacturer = 120
	ManufacturerKinetic                Manufacturer = 121
	ManufacturerJohnsonHealthTech      Manufacturer = 122
	ManufacturerPolarElectro           Manufacturer = 123
	ManufacturerSeesense               Manufacturer = 124
	ManufacturerNciTechnology          Manufacturer = 125
	ManufacturerIqsquare               Manufacturer = 126
	ManufacturerLeomo                  Manufacturer = 127
	ManufacturerIfitCom                Manufacturer = 128
	ManufacturerCorosByte              Manufacturer = 129
	ManufacturerVersaDesign            Manufacturer = 130
	ManufacturerChileaf                Manufacturer = 131
	ManufacturerCycplus                Manufacturer = 132
	ManufacturerGravaaByte             Manufacturer = 133
	ManufacturerSigeyi                 Manufacturer = 134
	ManufacturerCoospo                 Manufacturer = 135
	ManufacturerGeoid                  Manufacturer = 136
	ManufacturerBosch                  Manufacturer = 137
	ManufacturerKyto                   Manufacturer = 138
	ManufacturerKineticSports          Manufacturer = 139
	ManufacturerDecathlonByte          Manufacturer = 140
	ManufacturerTqSystems              Manufacturer = 141
	ManufacturerTagHeuer               Manufacturer = 142
	ManufacturerKeiserFitness          Manufacturer = 143
	ManufacturerZwiftByte              Manufacturer = 144
	ManufacturerPorscheEp              Manufacturer = 145
	ManufacturerBlackbird              Manufacturer = 146
	ManufacturerMeilanByte             Manufacturer = 147
	ManufacturerEzon                   Manufacturer = 148
	ManufacturerLaisi                  Manufacturer = 149
	ManufacturerMyzone                 Manufacturer = 150
	ManufacturerDevelopment            Manufacturer = 255
	ManufacturerHealthandlife          Manufacturer = 257
	ManufacturerLezyne                 Manufacturer = 258
	ManufacturerScribeLabs             Manufacturer = 259
	ManufacturerZwift                  Manufacturer = 260
	ManufacturerWatteam                Manufacturer = 261
	ManufacturerRecon                  Manufacturer = 262
	ManufacturerFaveroElectronics      Manufacturer = 263
	ManufacturerDynovelo               Manufacturer = 264
	ManufacturerStrava                 Manufacturer = 265
	ManufacturerPrecor                 Manufacturer = 266 // Amer Sports
	ManufacturerBryton                 Manufacturer = 267
	ManufacturerSram                   Manufacturer = 268
	ManufacturerNavman                 Manufacturer = 269 // MiTAC Global Corporation (Mio Technology)
	ManufacturerCobi                   Manufacturer = 270 // COBI GmbH
	ManufacturerSpivi                  Manufacturer = 271
	ManufacturerMioMagellan            Manufacturer = 272
	ManufacturerEvesports              Manufacturer = 273
	ManufacturerSensitivusGauge        Manufacturer = 274
	ManufacturerPodoon                 Manufacturer = 275
	ManufacturerLifeTimeFitness        Manufacturer = 276
	ManufacturerFalcoEMotors           Manufacturer = 277 // Falco eMotors Inc.
	ManufacturerMinoura                Manufacturer = 278
	ManufacturerCycliq                 Manufacturer = 279
	ManufacturerLuxottica              Manufacturer = 280
	ManufacturerTrainerRoad            Manufacturer = 281
	ManufacturerTheSufferfest          Manufacturer = 282
	ManufacturerFullspeedahead         Manufacturer = 283
	ManufacturerVirtualtraining        Manufacturer = 284
	ManufacturerFeedbacksports         Manufacturer = 285
	ManufacturerOmata                  Manufacturer = 286
	ManufacturerVdo                    Manufacturer = 287
	ManufacturerMagneticdays           Manufacturer = 288
	ManufacturerHammerhead             Manufacturer = 289
	ManufacturerKineticByKurt          Manufacturer = 290
	ManufacturerShapelog               Manufacturer = 291
	ManufacturerDabuziduo              Manufacturer = 292
	ManufacturerJetblack               Manufacturer = 293
	ManufacturerCoros                  Manufacturer = 294
	ManufacturerVirtugo                Manufacturer = 295
	ManufacturerVelosense              Manufacturer = 296
	ManufacturerCycligentinc           Manufacturer = 297
	ManufacturerTrailforks             Manufacturer = 298
	ManufacturerMahleEbikemotion       Manufacturer = 299
	ManufacturerNurvv                  Manufacturer = 300
	ManufacturerMicroprogram           Manufacturer = 301
	ManufacturerZone5cloud             Manufacturer = 302
	ManufacturerGreenteg               Manufacturer = 303
	ManufacturerYamahaMotors           Manufacturer = 304
	ManufacturerWhoop                  Manufacturer = 305
	ManufacturerGravaa                 Manufacturer = 306
	ManufacturerOnelap                 Manufacturer = 307
	ManufacturerMonarkExercise         Manufacturer = 308
	ManufacturerForm                   Manufacturer = 309
	ManufacturerDecathlon              Manufacturer = 310
	ManufacturerSyncros                Manufacturer = 311
	ManufacturerHeatup                 Manufacturer = 312
	ManufacturerCannondale             Manufacturer = 313
	ManufacturerTrueFitness            Manufacturer = 314
	ManufacturerRgtCycling             Manufacturer = 315
	ManufacturerVasa                   Manufacturer = 316
	ManufacturerRaceRepublic           Manufacturer = 317
	ManufacturerFazua                  Manufacturer = 318
	ManufacturerOrekaTraining          Manufacturer = 319
	ManufacturerLsec                   Manufacturer = 320 // Lishun Electric & Communication
	ManufacturerLululemonStudio        Manufacturer = 321
	ManufacturerShanyue                Manufacturer = 322
	ManufacturerSpinningMda            Manufacturer = 323
	ManufacturerHilldating             Manufacturer = 324
	ManufacturerAeroSensor             Manufacturer = 325
	ManufacturerNike                   Manufacturer = 326
	ManufacturerMagicshine             Manufacturer = 327
	ManufacturerActigraphcorp          Manufacturer = 5759
	ManufacturerInvalid                Manufacturer = 0xFFFF // INVALID
)

var manufacturertostrs = map[Manufacturer]string{
	ManufacturerGarmin:                 "garmin",
	ManufacturerGarminFr405Antfs:       "garmin_fr405_antfs",
	ManufacturerZephyr:                 "zephyr",
	ManufacturerDayton:                 "dayton",
	ManufacturerIdt:                    "idt",
	ManufacturerSrm:                    "srm",
	ManufacturerQuarq:                  "quarq",
	ManufacturerIbike:                  "ibike",
	ManufacturerSaris:                  "saris",
	ManufacturerSparkHk:                "spark_hk",
	ManufacturerTanita:                 "tanita",
	ManufacturerEchowell:               "echowell",
	ManufacturerDynastreamOem:          "dynastream_oem",
	ManufacturerNautilus:               "nautilus",
	ManufacturerDynastream:             "dynastream",
	ManufacturerTimex:                  "timex",
	ManufacturerMetrigear:              "metrigear",
	ManufacturerXelic:                  "xelic",
	ManufacturerBeurer:                 "beurer",
	ManufacturerCardiosport:            "cardiosport",
	ManufacturerAAndD:                  "a_and_d",
	ManufacturerHmm:                    "hmm",
	ManufacturerSuunto:                 "suunto",
	ManufacturerThitaElektronik:        "thita_elektronik",
	ManufacturerGpulse:                 "gpulse",
	ManufacturerCleanMobile:            "clean_mobile",
	ManufacturerPedalBrain:             "pedal_brain",
	ManufacturerPeaksware:              "peaksware",
	ManufacturerSaxonar:                "saxonar",
	ManufacturerLemondFitness:          "lemond_fitness",
	ManufacturerDexcom:                 "dexcom",
	ManufacturerWahooFitness:           "wahoo_fitness",
	ManufacturerOctaneFitness:          "octane_fitness",
	ManufacturerArchinoetics:           "archinoetics",
	ManufacturerTheHurtBox:             "the_hurt_box",
	ManufacturerCitizenSystems:         "citizen_systems",
	ManufacturerMagellan:               "magellan",
	ManufacturerOsynce:                 "osynce",
	ManufacturerHolux:                  "holux",
	ManufacturerConcept2:               "concept2",
	ManufacturerShimano:                "shimano",
	ManufacturerOneGiantLeap:           "one_giant_leap",
	ManufacturerAceSensor:              "ace_sensor",
	ManufacturerBrimBrothers:           "brim_brothers",
	ManufacturerXplova:                 "xplova",
	ManufacturerPerceptionDigital:      "perception_digital",
	ManufacturerBf1systems:             "bf1systems",
	ManufacturerPioneer:                "pioneer",
	ManufacturerSpantec:                "spantec",
	ManufacturerMetalogics:             "metalogics",
	Manufacturer4Iiiis:                 "4iiiis",
	ManufacturerSeikoEpson:             "seiko_epson",
	ManufacturerSeikoEpsonOem:          "seiko_epson_oem",
	ManufacturerIforPowell:             "ifor_powell",
	ManufacturerMaxwellGuider:          "maxwell_guider",
	ManufacturerStarTrac:               "star_trac",
	ManufacturerBreakaway:              "breakaway",
	ManufacturerAlatechTechnologyLtd:   "alatech_technology_ltd",
	ManufacturerMioTechnologyEurope:    "mio_technology_europe",
	ManufacturerRotor:                  "rotor",
	ManufacturerGeonaute:               "geonaute",
	ManufacturerIdBike:                 "id_bike",
	ManufacturerSpecialized:            "specialized",
	ManufacturerWtek:                   "wtek",
	ManufacturerPhysicalEnterprises:    "physical_enterprises",
	ManufacturerNorthPoleEngineering:   "north_pole_engineering",
	ManufacturerBkool:                  "bkool",
	ManufacturerCateye:                 "cateye",
	ManufacturerStagesCycling:          "stages_cycling",
	ManufacturerSigmasport:             "sigmasport",
	ManufacturerTomtom:                 "tomtom",
	ManufacturerPeripedal:              "peripedal",
	ManufacturerWattbike:               "wattbike",
	ManufacturerMoxy:                   "moxy",
	ManufacturerCiclosport:             "ciclosport",
	ManufacturerPowerbahn:              "powerbahn",
	ManufacturerAcornProjectsAps:       "acorn_projects_aps",
	ManufacturerLifebeam:               "lifebeam",
	ManufacturerBontrager:              "bontrager",
	ManufacturerWellgo:                 "wellgo",
	ManufacturerScosche:                "scosche",
	ManufacturerMagura:                 "magura",
	ManufacturerWoodway:                "woodway",
	ManufacturerElite:                  "elite",
	ManufacturerNielsenKellerman:       "nielsen_kellerman",
	ManufacturerDkCity:                 "dk_city",
	ManufacturerTacx:                   "tacx",
	ManufacturerDirectionTechnology:    "direction_technology",
	ManufacturerMagtonic:               "magtonic",
	Manufacturer1Partcarbon:            "1partcarbon",
	ManufacturerInsideRideTechnologies: "inside_ride_technologies",
	ManufacturerSoundOfMotion:          "sound_of_motion",
	ManufacturerStryd:                  "stryd",
	ManufacturerIcg:                    "icg",
	ManufacturerMipulse:                "MiPulse",
	ManufacturerBsxAthletics:           "bsx_athletics",
	ManufacturerLook:                   "look",
	ManufacturerCampagnoloSrl:          "campagnolo_srl",
	ManufacturerBodyBikeSmart:          "body_bike_smart",
	ManufacturerPraxisworks:            "praxisworks",
	ManufacturerLimitsTechnology:       "limits_technology",
	ManufacturerTopactionTechnology:    "topaction_technology",
	ManufacturerCosinuss:               "cosinuss",
	ManufacturerFitcare:                "fitcare",
	ManufacturerMagene:                 "magene",
	ManufacturerGiantManufacturingCo:   "giant_manufacturing_co",
	ManufacturerTigrasport:             "tigrasport",
	ManufacturerSalutron:               "salutron",
	ManufacturerTechnogym:              "technogym",
	ManufacturerBrytonSensors:          "bryton_sensors",
	ManufacturerLatitudeLimited:        "latitude_limited",
	ManufacturerSoaringTechnology:      "soaring_technology",
	ManufacturerIgpsport:               "igpsport",
	ManufacturerThinkrider:             "thinkrider",
	ManufacturerGopherSport:            "gopher_sport",
	ManufacturerWaterrower:             "waterrower",
	ManufacturerOrangetheory:           "orangetheory",
	ManufacturerInpeak:                 "inpeak",
	ManufacturerKinetic:                "kinetic",
	ManufacturerJohnsonHealthTech:      "johnson_health_tech",
	ManufacturerPolarElectro:           "polar_electro",
	ManufacturerSeesense:               "seesense",
	ManufacturerNciTechnology:          "nci_technology",
	ManufacturerIqsquare:               "iqsquare",
	ManufacturerLeomo:                  "leomo",
	ManufacturerIfitCom:                "ifit_com",
	ManufacturerCorosByte:              "coros_byte",
	ManufacturerVersaDesign:            "versa_design",
	ManufacturerChileaf:                "chileaf",
	ManufacturerCycplus:                "cycplus",
	ManufacturerGravaaByte:             "gravaa_byte",
	ManufacturerSigeyi:                 "sigeyi",
	ManufacturerCoospo:                 "coospo",
	ManufacturerGeoid:                  "geoid",
	ManufacturerBosch:                  "bosch",
	ManufacturerKyto:                   "kyto",
	ManufacturerKineticSports:          "kinetic_sports",
	ManufacturerDecathlonByte:          "decathlon_byte",
	ManufacturerTqSystems:              "tq_systems",
	ManufacturerTagHeuer:               "tag_heuer",
	ManufacturerKeiserFitness:          "keiser_fitness",
	ManufacturerZwiftByte:              "zwift_byte",
	ManufacturerPorscheEp:              "porsche_ep",
	ManufacturerBlackbird:              "blackbird",
	ManufacturerMeilanByte:             "meilan_byte",
	ManufacturerEzon:                   "ezon",
	ManufacturerLaisi:                  "laisi",
	ManufacturerMyzone:                 "myzone",
	ManufacturerDevelopment:            "development",
	ManufacturerHealthandlife:          "healthandlife",
	ManufacturerLezyne:                 "lezyne",
	ManufacturerScribeLabs:             "scribe_labs",
	ManufacturerZwift:                  "zwift",
	ManufacturerWatteam:                "watteam",
	ManufacturerRecon:                  "recon",
	ManufacturerFaveroElectronics:      "favero_electronics",
	ManufacturerDynovelo:               "dynovelo",
	ManufacturerStrava:                 "strava",
	ManufacturerPrecor:                 "precor",
	ManufacturerBryton:                 "bryton",
	ManufacturerSram:                   "sram",
	ManufacturerNavman:                 "navman",
	ManufacturerCobi:                   "cobi",
	ManufacturerSpivi:                  "spivi",
	ManufacturerMioMagellan:            "mio_magellan",
	ManufacturerEvesports:              "evesports",
	ManufacturerSensitivusGauge:        "sensitivus_gauge",
	ManufacturerPodoon:                 "podoon",
	ManufacturerLifeTimeFitness:        "life_time_fitness",
	ManufacturerFalcoEMotors:           "falco_e_motors",
	ManufacturerMinoura:                "minoura",
	ManufacturerCycliq:                 "cycliq",
	ManufacturerLuxottica:              "luxottica",
	ManufacturerTrainerRoad:            "trainer_road",
	ManufacturerTheSufferfest:          "the_sufferfest",
	ManufacturerFullspeedahead:         "fullspeedahead",
	ManufacturerVirtualtraining:        "virtualtraining",
	ManufacturerFeedbacksports:         "feedbacksports",
	ManufacturerOmata:                  "omata",
	ManufacturerVdo:                    "vdo",
	ManufacturerMagneticdays:           "magneticdays",
	ManufacturerHammerhead:             "hammerhead",
	ManufacturerKineticByKurt:          "kinetic_by_kurt",
	ManufacturerShapelog:               "shapelog",
	ManufacturerDabuziduo:              "dabuziduo",
	ManufacturerJetblack:               "jetblack",
	ManufacturerCoros:                  "coros",
	ManufacturerVirtugo:                "virtugo",
	ManufacturerVelosense:              "velosense",
	ManufacturerCycligentinc:           "cycligentinc",
	ManufacturerTrailforks:             "trailforks",
	ManufacturerMahleEbikemotion:       "mahle_ebikemotion",
	ManufacturerNurvv:                  "nurvv",
	ManufacturerMicroprogram:           "microprogram",
	ManufacturerZone5cloud:             "zone5cloud",
	ManufacturerGreenteg:               "greenteg",
	ManufacturerYamahaMotors:           "yamaha_motors",
	ManufacturerWhoop:                  "whoop",
	ManufacturerGravaa:                 "gravaa",
	ManufacturerOnelap:                 "onelap",
	ManufacturerMonarkExercise:         "monark_exercise",
	ManufacturerForm:                   "form",
	ManufacturerDecathlon:              "decathlon",
	ManufacturerSyncros:                "syncros",
	ManufacturerHeatup:                 "heatup",
	ManufacturerCannondale:             "cannondale",
	ManufacturerTrueFitness:            "true_fitness",
	ManufacturerRgtCycling:             "RGT_cycling",
	ManufacturerVasa:                   "vasa",
	ManufacturerRaceRepublic:           "race_republic",
	ManufacturerFazua:                  "fazua",
	ManufacturerOrekaTraining:          "oreka_training",
	ManufacturerLsec:                   "lsec",
	ManufacturerLululemonStudio:        "lululemon_studio",
	ManufacturerShanyue:                "shanyue",
	ManufacturerSpinningMda:            "spinning_mda",
	ManufacturerHilldating:             "hilldating",
	ManufacturerAeroSensor:             "aero_sensor",
	ManufacturerNike:                   "nike",
	ManufacturerMagicshine:             "magicshine",
	ManufacturerActigraphcorp:          "actigraphcorp",
	ManufacturerInvalid:                "invalid",
}

func (m Manufacturer) String() string {
	val, ok := manufacturertostrs[m]
	if !ok {
		return strconv.FormatUint(uint64(m), 10)
	}
	return val
}

var strtomanufacturer = func() map[string]Manufacturer {
	m := make(map[string]Manufacturer)
	for t, str := range manufacturertostrs {
		m[str] = Manufacturer(t)
	}
	return m
}()

// FromString parse string into Manufacturer constant it's represent, return ManufacturerInvalid if not found.
func ManufacturerFromString(s string) Manufacturer {
	val, ok := strtomanufacturer[s]
	if !ok {
		return strtomanufacturer["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListManufacturer() []Manufacturer {
	vs := make([]Manufacturer, 0, len(manufacturertostrs))
	for i := range manufacturertostrs {
		vs = append(vs, Manufacturer(i))
	}
	return vs
}
