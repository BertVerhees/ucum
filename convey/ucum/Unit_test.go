package ucum

import (
	"ucum"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func SetUpService() {
	var err error
	definitions := os.Getenv("GOPATH") + "/src/ucum/terminology_data/ucum-essence.xml"
	testservice, err = ucum.GetInstanceOfUcumEssenceService(definitions)
	if err != nil {
		panic(err.Error())
	}
}

func TestUnits(t *testing.T) {
	SetUpService()
	Convey("Test Units on errormsg", t, func() {
		for _, u := range units {
			_, s := testservice.Validate(u)
			So(s, ShouldBeEmpty)
			s, _ = testservice.Analyse(u)
			fmt.Println(s)
		}
		for _, u := range wrongunits {
			_, s := testservice.Validate(u)
			So(s, ShouldNotBeEmpty)
		}
	})
	Convey("Test Units on Parsing", t, func() {
		for _, u := range units {
			term, _ := ucum.NewExpressionParser(testservice.Model).Parse(u)
			So(term, ShouldNotBeNil)
		}
	})
}

var wrongunits = []string{
	"m/",
	"",
	"10+3/ul",
	"{a}rad2{b}",
	"{|}1",
	"iU",
	"molv",
	"[BETH'U]",
	"[iIU]",
	"[iIU]/d",
	"[iIU]/L",
	"[iIU]/mL",
	"g/12h",
	"g/48h",
	"g/4h",
	"g/6h",
	"g/72h",
	"m[iIU]/L",
	"mg/12h",
	"mL/10h",
	"mL/12h",
	"mL/2h",
	"mL/4h",
	"mL/5h",
	"mL/6h",
	"mL/72h",
	"mL/8h",
	"mmol/12h",
	"mmol/5h",
	"mmol/6h",
	"mmol/kg[H20]",
	"U/12h",
	"U/1h",
	"U/2h",
	"u[iIU]/mL",
	"cm[H20]",
	"ug(8.h)",
	"ug(8hr)",
	"g.m/{hb}m2",
}
var testservice *ucum.UcumEssenceService
var units = []string{
	"s.m-1.g-1",
	"s/m/mg",
	"s/4.m",
	"10*3/ul",
	"m",
	"/m",
	"10*-3/ul",
	"10*+3/ul",
	"m",
	"m[H2O]",
	"10*23",
	"rad2",
	"m3.kg-1.s-2",
	"4.[pi].10*-7.N/A2",
	"rad2{a}",
	"rad2{?}",
	"{a}.rad2{b}",
	"1{c}",
	"{e}",
	"%",
	"dB[10.nV]",
	"[cup_us]",
	"[foz_br]",
	"[ft_i]",
	"[in_i]",
	"[yd_i]",
	"[gal_br]",
	"[lb_av]",
	"[oz_av]",
	"[pt_br]",
	"[qt_br]",
	"[sft_i]",
	"[sin_i]",
	"[syd_i]",
	"[tbs_us]",
	"[tsp_us]",
	"1/d",
	"1/min",
	"a",
	"cm",
	"cm2",
	"cm3",
	"d",
	"dg",
	"dl",
	"g",
	"g/d",
	"g/l",
	"h",
	"kg",
	"l",
	"m",
	"mm",
	"m2",
	"meq",
	"mg",
	"mg",
	"mg/d",
	"min",
	"ml",
	"ml/s",
	"mm[Hg]",
	"mm2",
	"mm3",
	"mmol",
	"mmol/l",
	"mo",
	"mol",
	"ms",
	"mU",
	"ng",
	"ng",
	"nl",
	"nl",
	"pg/ml",
	"s",
	"U",
	"U/l",
	"ug",
	"ug/min",
	"ul",
	"umol",
	"umol/l",
	"wk",
	"%",
	"[cup_us]",
	"[foz_br]",
	"[gal_br]",
	"[sft_i]",
	"[sin_i]",
	"[lb_av]",
	"[oz_av]",
	"[pt_br]",
	"[qt_br]",
	"[tbs_us]",
	"[tsp_us]",
	"[syd_i]",
	"cm2",
	"cm3",
	"g",
	"kg",
	"l",
	"m2",
	"meq",
	"mg",
	"ml",
	"mm2",
	"mm3",
	"mmol",
	"mU",
	"ng",
	"nl",
	"U",
	"ug",
	"ul",
	"umol",
	"a",
	"d",
	"h",
	"min",
	"mo",
	"s",
	"wk",
	"[ft_i]",
	"[in_i]",
	"[lb_av]",
	"[oz_av]",
	"[yd_i]",
	"cm",
	"g",
	"kg",
	"m",
	"mm",
	"[mi_us]",
	"[yd_i]",
	"deg",
	"km",
	"m",
	"%",
	"/[HPF]",
	"/[LPF]",
	"/L",
	"/mL",
	"/mmol",
	"[APL'U]",
	"[GPL'U]",
	"[IU]",
	"[IU]/d",
	"[IU]/L",
	"[IU]/mL",
	"[MPL'U]",
	"10*12/L",
	"10*6/L",
	"10*9/L",
	"Cel",
	"cm",
	"cm/s",
	"fL",
	"fmol/L",
	"g",
	"g/d",
	"g/g",
	"g/L",
	"h",
	"km",
	"kU/L",
	"L/L",
	"m[IU]/L",
	"mg",
	"mg/d",
	"mg/g",
	"mg/L",
	"mg/mg",
	"mg/mL",
	"min",
	"mL",
	"mL/d",
	"mL/min",
	"mm",
	"mm/h",
	"mm[Hg]",
	"mmol",
	"mmol/d",
	"mmol/g",
	"mmol/kg",
	"mmol/L",
	"mmol/mmol",
	"mU/L",
	"ng/d",
	"ng/g",
	"ng/L",
	"ng/mL",
	"nmol/d",
	"nmol/g",
	"nmol/h/mL",
	"nmol/L",
	"nmol/mmol",
	"nmol/nmol",
	"pg",
	"pg/mL",
	"pmol/d",
	"pmol/g",
	"pmol/h/mg",
	"pmol/h/mL",
	"pmol/L",
	"pmol/mmol",
	"s",
	"U",
	"U/d",
	"U/g",
	"U/kg",
	"U/L",
	"U/mL",
	"u[IU]/mL",
	"ug",
	"ug/d",
	"ug/g",
	"ug/L",
	"ug/mL",
	"um/s",
	"umol",
	"umol/2.h",
	"umol/d",
	"umol/g",
	"umol/L",
	"umol/mmol",
	"umol/umol",
	"wk",
	"[arb'U]",
	"dyn.s/(cm5.m2)",
	"[iU]/mL",
	"mL/h",
	"[bdsk'U]",
	"dyn.s/cm5",
	"K/W",
	"mm[Hg]",
	"{bsa}",
	"cm[H2O]",
	"kg{body_wt}",
	"mm/h",
	"cal",
	"cm[H2O].s/L",
	"kg/m2",
	"mmol/(8.h.kg)",
	"{cfu}",
	"cm[H2O]/(s.m)",
	"kg/h",
	"mmol/(8.h)",
	"[drp]",
	"dB[SPL]",
	"L/(8.h)",
	"mmol/(kg.h)",
	"[ka'U]",
	"REM",
	"L/h",
	"mmol/h",
	"kcal",
	"g{creat}",
	"[lb_av]",
	"ng/(8.h)",
	"kcal/(8.h)",
	"g{hgb}",
	"ng/(8.h.kg)",
	"kcal/d",
	"g{tit_nit}",
	"ms/s",
	"ng/(kg.h)",
	"kcal/h",
	"g{tot_prot}",
	"Ms",
	"ng/h",
	"[knk'U]",
	"g{wet_tis}",
	"meq/(8.h)",
	"osm",
	"[mclg'U]",
	"g.m/m2{hb}",
	"meq/(8.h.kg)",
	"osm/kg",
	"{od}",
	"g.m/{hb}",
	"meq/(kg.h)",
	"osm/L",
	"pH",
	"g/(8.h)",
	"meq/h",
	"pA",
	"[ppb]",
	"g/(8.kg.h)",
	"mg/(8.h)",
	"Pa",
	"[ppm]",
	"g/(kg.h)",
	"mg/(8.h.kg)",
	"[pptr]",
	"g/h",
	"mg/(kg.h)",
	"S",
	"[ppth]",
	"[in_us]",
	"mg/h",
	"[todd'U]",
	"[in_i'Hg]",
	"m[iU]/mL",
	"ug/(8.h.kg)",
	"/[arb'U]",
	"[iU]",
	"mL/{hb}.m2",
	"ug/(kg.h)",
	"[HPF]",
	"[iU]/d",
	"mL/(8.h)",
	"ug/h",
	"/{tot}",
	"[iU]/h",
	"mL/(8.h.kg)",
	"u[iU]",
	"/[iU]",
	"[iU]/kg",
	"mL/{hb}",
	"10*3{rbc}",
	"[iU]/L",
	"mL/(kg.h)",
	"10.L/(min.m2)",
	"[iU]/min",
	"mL/cm[H2O]",
	"%",
	"bar",
	"g/L",
	"L.s",
	"mg",
	"mmol/(kg.d)",
	"ng/L",
	"ueq",
	"/kg",
	"Bq",
	"g/m2",
	"L/(min.m2)",
	"mg/(kg.d)",
	"mmol/(kg.min)",
	"ng/m2",
	"ug",
	"/L",
	"Cel",
	"g/min",
	"L/d",
	"mg/(kg.min)",
	"mmol/kg",
	"ng/min",
	"ug/(kg.d)",
	"/m3",
	"cm",
	"Gy",
	"L/kg",
	"mg/d",
	"mmol/L",
	"ng/mL",
	"ug/(kg.min)",
	"/min",
	"cm2/s",
	"h",
	"L/min",
	"mg/dL",
	"mmol/m2",
	"ng/s",
	"ug/d",
	"/m3",
	"d",
	"hL",
	"L/s",
	"mg/kg",
	"mmol/min",
	"nkat",
	"ug/dL",
	"/min",
	"dB",
	"J/L",
	"lm",
	"mg/L",
	"mol/(kg.s)",
	"nm",
	"ug/g",
	"/mL",
	"deg",
	"kat",
	"m",
	"mg/m2",
	"mol/kg",
	"nmol/s",
	"ug/kg",
	"1/mL",
	"eq",
	"kat/kg",
	"m/s2",
	"mg/m3",
	"mol/L",
	"ns",
	"ug/L",
	"10*12/L",
	"eV",
	"kat/L",
	"m2",
	"mg/min",
	"mol/m3",
	"Ohm",
	"ug/m2",
	"10*3/L",
	"kg",
	"m2/s",
	"mL",
	"mol/s",
	"Ohm.m",
	"ug/min",
	"10*3/mL",
	"fg",
	"kg.m/s",
	"m3/s",
	"mL/(kg.d)",
	"mosm/L",
	"pg",
	"ukat",
	"10*3/mm3",
	"fL",
	"kg/(s.m2)",
	"mbar",
	"mL/(kg.min)",
	"ms",
	"pg/L",
	"um",
	"10*6/L",
	"fmol",
	"kg/L",
	"mbar.s/L",
	"mL/(min.m2)",
	"mV",
	"pg/mL",
	"umol",
	"10*6/mL",
	"g",
	"kg/m3",
	"meq",
	"mL/d",
	"pkat",
	"umol/d",
	"10*6/mm3",
	"g.m",
	"kg/min",
	"meq/(kg.d)",
	"mL/kg",
	"pm",
	"umol/L",
	"10*9/L",
	"g/(kg.d)",
	"kg/mol",
	"meq/(kg.min)",
	"mL/m2",
	"ng",
	"pmol",
	"umol/min",
	"10*9/mL",
	"g/(kg.min)",
	"kg/s",
	"meq/d",
	"mL/mbar",
	"ng/(kg.d)",
	"ps",
	"us",
	"10*9/mm3",
	"g/d",
	"kPa",
	"meq/kg",
	"mL/min",
	"ng/(kg.min)",
	"pt",
	"uV",
	"10.L/min",
	"g/dL",
	"ks",
	"meq/L",
	"mL/s",
	"ng/d",
	"Sv",
	"V",
	"a/m",
	"g/kg",
	"L",
	"meq/min",
	"mm",
	"ng/kg",
	"t",
}
