package ucum

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type UcumModel struct {
	Version      string
	Revision     string
	RevisionDate time.Time
	Prefixes     []*Prefix
	BaseUnits    []*BaseUnit
	DefinedUnits []*DefinedUnit
}

func NewUcumModel(version, revision string, revisionDate time.Time) *UcumModel {
	r := &UcumModel{}
	r.Version = version
	r.Revision = revision
	r.RevisionDate = revisionDate
	r.Prefixes = make([]*Prefix, 0)
	r.BaseUnits = make([]*BaseUnit, 0)
	r.DefinedUnits = make([]*DefinedUnit, 0)
	return r
}

func (u *UcumModel) GetUnit(code string) Uniter {
	for _, unit := range u.BaseUnits {
		if unit.Code == code {
			return unit
		}
	}
	for _, unit := range u.DefinedUnits {
		if unit.Code == code {
			return unit
		}
	}
	return nil
}

func (u *UcumModel) Search(kind ConceptKind, text string, isRegex bool) []Concepter {
	concepts := make([]Concepter, 0)
	if kind == 0 || kind == PREFIX {
		concepts = append(concepts, u.searchPrefixes(text, isRegex)...)
	}
	if kind == 0 || kind == BASEUNIT || kind == UNIT {
		concepts = append(concepts, u.searchUnits(text, isRegex, kind)...)
	}
	return concepts
}

func (u *UcumModel) searchPrefixes(text string, isRegex bool) []Concepter {
	concepts := make([]Concepter, 0)
	for _, c := range u.Prefixes {
		if u.matchesConcept(c, text, isRegex) {
			concepts = append(concepts, c)
		}
	}
	return concepts
}

func (u *UcumModel) getBaseUnit(code string) *BaseUnit {
	for _, unit := range u.BaseUnits {
		if unit.Code == code {
			return unit
		}
	}
	return nil
}

func (u *UcumModel) searchUnits(text string, isRegex bool, kind ConceptKind) []Concepter {
	concepts := make([]Concepter, 0)
	if kind == BASEUNIT {
		for _, unit := range u.BaseUnits {
			if u.matchesUnit(unit, text, isRegex) {
				concepts = append(concepts, unit)
			}
		}
	}
	if kind == UNIT {
		for _, unit := range u.DefinedUnits {
			if u.matchesUnit(unit, text, isRegex) {
				concepts = append(concepts, unit)
			}
		}
	}
	return concepts
}

func (u *UcumModel) matchesUnit(unit Uniter, text string, isRegex bool) bool {
	return u.matches(unit.GetProperty(), text, isRegex) || u.matchesConcept(unit, text, isRegex)
}

func (u *UcumModel) matches(value, text string, isRegEx bool) bool {
	if isRegEx {
		b, _ := regexp.MatchString( text, value)
		return b
	} else {
		return strings.Contains(strings.ToLower(value), strings.ToLower(text))
	}
}

func (u *UcumModel) matchesConcept(concept Concepter, text string, isRegex bool) bool {
	for _, name := range concept.GetNames() {
		if u.matches(name, text, isRegex) {
			return true
		}
	}
	if u.matches(concept.GetCode(), text, isRegex) {
		return true
	}
	if u.matches(concept.GetCodeUC(), text, isRegex) {
		return true
	}
	if u.matches(concept.GetPrintSymbol(), text, isRegex) {
		return true
	}
	return false
}

// Concept=====================================================
/**
Base of Unit and Prefix.
Top class
Code = String (case sensitive c/s)
CodeUC = String, case insensitive c/i)
Kind = ConceptKind (PREFIX, BASEUNIT or UNIT)
Name = full (official) name of the concept
PrintSymbol
 */
type Concepter interface {
	GetDescription() string
	String() string
	GetCode() string
	GetKind() ConceptKind
	GetNames() []string
	GetCodeUC() string
	GetPrintSymbol() string
}

type Concept struct {
	Code        string
	CodeUC      string
	Kind        ConceptKind
	Names       []string
	PrintSymbol string
}

func NewConcept(kind ConceptKind, code string, codeUC string) (*Concept, error) {
	c := &Concept{
		Kind:   kind,
		Code:   code,
		CodeUC: codeUC,
	}
	return c, nil
}

func (c Concept) GetDescription() string {
	description := strings.ToLower(c.Kind.String()) + " " + c.Code + " ('" + c.Names[0] + "')"
	return description
}

func (c Concept) String() string {
	return c.Code + " = " + c.GetDescription()
}

func (c Concept) GetCode() string {
	return c.Code
}
func (c Concept) GetCodeUC() string {
	return c.CodeUC
}
func (c Concept) GetPrintSymbol() string {
	return c.PrintSymbol
}

func (c Concept) GetKind() ConceptKind {
	return c.Kind
}
func (c Concept) GetNames() []string {
	return c.Names
}

//Unit=====================================================
/**
Parent is Concept
Children are BaseUnit and DefinedUnit
 */
type Uniter interface {
	Concepter
	GetProperty() string
}

type Unit struct {
	Concept
	Property string
}

func NewUnit(kind ConceptKind, code string, codeUC string) (*Unit, error) {
	u := &Unit{}
	u.Kind = kind
	u.Code = code
	u.CodeUC = codeUC
	return u, nil
}

func (u Unit) GetDescription() string {
	return strings.ToLower(u.Kind.String()) + " " + u.Code + " ('" + u.Names[0] + "')" + " (" + u.Property + ")"
}

func (u Unit) String() string {
	return u.Code + " = " + u.GetDescription()
}

func (u Unit) GetProperty() string {
	return u.Property
}

//BaseUnit=====================================================
//Parent is Unit
//DIM is character indicating the Property
type BaseUnit struct {
	Unit
	Dim rune
}

func NewBaseUnit(kind ConceptKind, code string, codeUC string) (*BaseUnit, error) {
	b := &BaseUnit{}
	b.Kind = BASEUNIT
	b.Code = code
	b.CodeUC = codeUC
	return b, nil
}

//DefinedUnit=====================================================
/**
Parent is Unit
- Class "dimless" = dimensionless
- Class SI "si" = the SI units (International System of Units)/
SI units are mole, steradian, hertz, newton, pascal, joule, watt, ampère, volt, farad, ohm, siemens, weber, degree,
tesla, henry, lumen, lux, becquerel, gray, sievert
- Class ISO1000 "iso1000" = other units from ISO 1000
- Class "const" = Natural units, velocity of light, Planck constant, etc.
- Class "cgs" = The units of the older Centimeter-Gram-Second (CGS) system
- Class "cust" = Customary units have once been used all over Europe.
Units were taken from nature: anatomical structures (e.g., arm, foot, finger), botanical objects
- Class "us-lengths" = The older U.S. units according to the definition of the inch in the U.S. Metric Law of 1866 and
the definition of foot and yard that was valid from 1893 until 1959
- Class "us-volumes" = “capacity” measures, which are different for fluid goods (wine) and dry goods (grain)
- Class "brit-volumes" = British Imperial volumes according to the Weights and Measures Act of 1824
- Class "avoirdupois" = The avoirdupois system is used in the U.S. as well as in coutries that use the British Imperial system.
Avoirdupois is the default system of mass units used for all goods that “have weight”
- Class "troy" = The troy system originates in Troyes, a City in the Champagne (France) that hosted a major European fair.
- Class "apoth" = The apothecaries' system of mass units
- Class "typeset" = There are three systems of typesetter's lengths in use today: Françcois-Ambroise Didot (1730-1804), Didot, U.S. type foundries
- Class "heat" = Older units of heat (energy) and temperature
- Class "clinical" = Units used mainly in clinical medicine
- Class "chemical" = Units used mainly in chemical and biochemical laboratories
- Class "levels" = Pseudo-units defined to express logarithms of ratios between two quantities of the same kind
- Class "misc" = Not otherwise classified units
- Class "infotech" = Units used in information technology
 */
type DefinedUnit struct {
	Unit
	Class     string
	IsSpecial bool
	Metric    bool
	Value     *Value
}

func NewDefinedUnit(kind ConceptKind, code string, codeUC string) (*DefinedUnit, error) {
	b := &DefinedUnit{}
	b.Kind = UNIT
	b.Code = code
	b.CodeUC = codeUC
	return b, nil
}

func (d DefinedUnit) GetDescription() string {
	return strings.ToLower(d.Kind.String()) + " " + d.Code + " ('" + d.Names[0] + "')" + " (" + d.Property + ")" + " = " + d.Value.GetDescription()
}

//Prefix=====================================================
/**
Parent is Concept
Value = is the scalar value by which the unit atom is multiplied if combined with the prefix.
 */
type Prefix struct {
	Concept
	Value Decimal
}

func NewPrefix(kind ConceptKind, code string, codeUC string) (*Prefix, error) {
	b := &Prefix{}
	b.Kind = PREFIX
	b.Code = code
	b.CodeUC = codeUC
	return b, nil
}

func (p Prefix) GetDescription() string {
	return strings.ToLower(p.Kind.String()) + " " + p.Code + " ('" + p.Names[0] + "')" + " = " + p.Value.String()
}

//Value=====================================================
type Value struct {
	Text   string
	Unit   string
	UnitUC string
	Value  Decimal
}

func NewValue(unit, unitUC string, value Decimal) (*Value, error) {
	v := &Value{}
	v.Unit = unit
	v.UnitUC = unitUC
	v.Value = value
	return v, nil
}

func (v Value) GetDescription() string {
	if v.Value == Zero {
		return v.Unit
	}
	return v.Value.String()
}

//Canonical=====================================================
/**
unit terms that are commonly used in medicine. Since the space of possible unit terms is infinite in theory and very large in practice,
no attempt has been made on a systematic coverage of possible units. All necessary units can be built from the rules of
The Unified Code for Units of Measure and there is no need of a particular term to be enumerated in order to be valid.

The canonical form itself consists of 3 columns: (4.1) the magnitude value of the unit term in terms of the canonical unit;
(4.2) a canonical unit term; (4.3) if applicable a special conversion function code.

A canonical unit is a unit of measurement agreed upon as default in a certain context.
 */
type Canonical struct {
	Units []*CanonicalUnit
	Value Decimal
}

func (c *Canonical) RemoveFromUnits(i int) {
	c.Units[i] = c.Units[len(c.Units)-1]
	c.Units[len(c.Units)-1] = nil
	c.Units = c.Units[:len(c.Units)-1]
}

type ByCode []*CanonicalUnit

func (a ByCode) Len() int           { return len(a) }
func (a ByCode) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCode) Less(i, j int) bool { return a[i].base.Code < a[j].base.Code }

func (c *Canonical) SortUnits() {
	sort.Sort(ByCode(c.Units))
}

func NewCanonical(value Decimal) (*Canonical, error) {
	v := &Canonical{
		Value: value,
		Units: make([]*CanonicalUnit, 0),
	}
	return v, nil
}

func (c *Canonical) MultiplyValueDecimal(multiplicand Decimal) {
	c.Value = c.Value.Multiply(multiplicand)
}

func (c *Canonical) MultiplyValueInt(multiplicand int) error {
	d, err := NewDecimal(strconv.Itoa(multiplicand))
	if err != nil {
		return err
	}
	c.Value = c.Value.Multiply(d)
	return nil
}

func (c *Canonical) DivideValueDecimal(divisor Decimal) {
	c.Value = c.Value.Divide(divisor)
}

func (c *Canonical) DivideValueInt(divisor int) error {
	d, err := NewDecimal(strconv.Itoa(divisor))
	if err != nil {
		return err
	}
	c.Value = c.Value.Divide(d)
	return nil
}

//CanonicalUnit=====================================================
/**
base a canonical unit term;
 */
type CanonicalUnit struct {
	base     *BaseUnit
	Exponent int
}

func NewCanonicalUnit(base *BaseUnit, exponent int) (*CanonicalUnit, error) {
	v := &CanonicalUnit{}
	v.base = base
	v.Exponent = exponent
	return v, nil
}

func (c *CanonicalUnit) Base() *BaseUnit {
	return c.base
}

//Component=====================================================
type Componenter interface {
}

type Component struct {
}

//Factor=====================================================
/**
Parent is component
Connected with TokenType NUMBER
 */
type Factor struct {
	Component
	Value int
}

func NewFactor(value int) *Factor {
	v := &Factor{
		Value: value,
	}
	return v
}

//Symbol=====================================================
/**
// Unit may be Base Unit or DefinedUnit
// Prefix only if unit is metric
 */
type Symbol struct {
	Component
	Unit     Uniter
	Prefix   *Prefix
	Exponent int
}

func NewSymbol(unit Uniter, prefix *Prefix, exponent int) (*Symbol, error) {
	v := &Symbol{}
	v.Unit = unit
	v.Prefix = prefix
	v.Exponent = exponent
	return v, nil
}

func (s *Symbol) HasPrefix() bool {
	return s.Prefix != nil
}

func (s *Symbol) InvertExponent() {
	s.Exponent = -s.Exponent
}

//Term=====================================================
// op-term where op = /
// component
// component-op-term
/**
Parent is Component
 */
type Term struct {
	Component
	Comp Componenter
	Op   Operator
	Term *Term
}

func NewTerm() (*Term, error) {
	return &Term{}, nil
}

func (t *Term) SetTermCheckOp(term *Term) {
	if term != nil {
		t.Term = term
		t.Op = term.Op
	} else {
		t.Term = nil
		t.Op = 0
	}
}

//Pair=====================================================
type Pair struct {
	Value Decimal
	Code  string
}

func NewPair(value Decimal, code string) *Pair {
	p := &Pair{}
	p.Value = value
	p.Code = code
	return p
}