// Code generated by ent, DO NOT EDIT.

package student

import (
	"mocku/backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldID, id))
}

// IdentityCard applies equality check predicate on the "identity_card" field. It's identical to IdentityCardEQ.
func IdentityCard(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldIdentityCard, v))
}

// BirthDate applies equality check predicate on the "birth_date" field. It's identical to BirthDateEQ.
func BirthDate(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldBirthDate, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldPhone, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAddress, v))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldNumber, v))
}

// District applies equality check predicate on the "district" field. It's identical to DistrictEQ.
func District(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldDistrict, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldCity, v))
}

// PostalCode applies equality check predicate on the "postal_code" field. It's identical to PostalCodeEQ.
func PostalCode(v int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldPostalCode, v))
}

// CreditUnitsAccumulated applies equality check predicate on the "credit_units_accumulated" field. It's identical to CreditUnitsAccumulatedEQ.
func CreditUnitsAccumulated(v int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldCreditUnitsAccumulated, v))
}

// TotalAverage applies equality check predicate on the "total_average" field. It's identical to TotalAverageEQ.
func TotalAverage(v float64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldTotalAverage, v))
}

// IdentityCardEQ applies the EQ predicate on the "identity_card" field.
func IdentityCardEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldIdentityCard, v))
}

// IdentityCardNEQ applies the NEQ predicate on the "identity_card" field.
func IdentityCardNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldIdentityCard, v))
}

// IdentityCardIn applies the In predicate on the "identity_card" field.
func IdentityCardIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldIdentityCard, vs...))
}

// IdentityCardNotIn applies the NotIn predicate on the "identity_card" field.
func IdentityCardNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldIdentityCard, vs...))
}

// IdentityCardGT applies the GT predicate on the "identity_card" field.
func IdentityCardGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldIdentityCard, v))
}

// IdentityCardGTE applies the GTE predicate on the "identity_card" field.
func IdentityCardGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldIdentityCard, v))
}

// IdentityCardLT applies the LT predicate on the "identity_card" field.
func IdentityCardLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldIdentityCard, v))
}

// IdentityCardLTE applies the LTE predicate on the "identity_card" field.
func IdentityCardLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldIdentityCard, v))
}

// IdentityCardContains applies the Contains predicate on the "identity_card" field.
func IdentityCardContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldIdentityCard, v))
}

// IdentityCardHasPrefix applies the HasPrefix predicate on the "identity_card" field.
func IdentityCardHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldIdentityCard, v))
}

// IdentityCardHasSuffix applies the HasSuffix predicate on the "identity_card" field.
func IdentityCardHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldIdentityCard, v))
}

// IdentityCardEqualFold applies the EqualFold predicate on the "identity_card" field.
func IdentityCardEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldIdentityCard, v))
}

// IdentityCardContainsFold applies the ContainsFold predicate on the "identity_card" field.
func IdentityCardContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldIdentityCard, v))
}

// BirthDateEQ applies the EQ predicate on the "birth_date" field.
func BirthDateEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldBirthDate, v))
}

// BirthDateNEQ applies the NEQ predicate on the "birth_date" field.
func BirthDateNEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldBirthDate, v))
}

// BirthDateIn applies the In predicate on the "birth_date" field.
func BirthDateIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldBirthDate, vs...))
}

// BirthDateNotIn applies the NotIn predicate on the "birth_date" field.
func BirthDateNotIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldBirthDate, vs...))
}

// BirthDateGT applies the GT predicate on the "birth_date" field.
func BirthDateGT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldBirthDate, v))
}

// BirthDateGTE applies the GTE predicate on the "birth_date" field.
func BirthDateGTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldBirthDate, v))
}

// BirthDateLT applies the LT predicate on the "birth_date" field.
func BirthDateLT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldBirthDate, v))
}

// BirthDateLTE applies the LTE predicate on the "birth_date" field.
func BirthDateLTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldBirthDate, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldPhone, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldAddress, v))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v int) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...int) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...int) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v int) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v int) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v int) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v int) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldNumber, v))
}

// DistrictEQ applies the EQ predicate on the "district" field.
func DistrictEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldDistrict, v))
}

// DistrictNEQ applies the NEQ predicate on the "district" field.
func DistrictNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldDistrict, v))
}

// DistrictIn applies the In predicate on the "district" field.
func DistrictIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldDistrict, vs...))
}

// DistrictNotIn applies the NotIn predicate on the "district" field.
func DistrictNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldDistrict, vs...))
}

// DistrictGT applies the GT predicate on the "district" field.
func DistrictGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldDistrict, v))
}

// DistrictGTE applies the GTE predicate on the "district" field.
func DistrictGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldDistrict, v))
}

// DistrictLT applies the LT predicate on the "district" field.
func DistrictLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldDistrict, v))
}

// DistrictLTE applies the LTE predicate on the "district" field.
func DistrictLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldDistrict, v))
}

// DistrictContains applies the Contains predicate on the "district" field.
func DistrictContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldDistrict, v))
}

// DistrictHasPrefix applies the HasPrefix predicate on the "district" field.
func DistrictHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldDistrict, v))
}

// DistrictHasSuffix applies the HasSuffix predicate on the "district" field.
func DistrictHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldDistrict, v))
}

// DistrictEqualFold applies the EqualFold predicate on the "district" field.
func DistrictEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldDistrict, v))
}

// DistrictContainsFold applies the ContainsFold predicate on the "district" field.
func DistrictContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldDistrict, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldCity, v))
}

// PostalCodeEQ applies the EQ predicate on the "postal_code" field.
func PostalCodeEQ(v int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldPostalCode, v))
}

// PostalCodeNEQ applies the NEQ predicate on the "postal_code" field.
func PostalCodeNEQ(v int) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldPostalCode, v))
}

// PostalCodeIn applies the In predicate on the "postal_code" field.
func PostalCodeIn(vs ...int) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldPostalCode, vs...))
}

// PostalCodeNotIn applies the NotIn predicate on the "postal_code" field.
func PostalCodeNotIn(vs ...int) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldPostalCode, vs...))
}

// PostalCodeGT applies the GT predicate on the "postal_code" field.
func PostalCodeGT(v int) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldPostalCode, v))
}

// PostalCodeGTE applies the GTE predicate on the "postal_code" field.
func PostalCodeGTE(v int) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldPostalCode, v))
}

// PostalCodeLT applies the LT predicate on the "postal_code" field.
func PostalCodeLT(v int) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldPostalCode, v))
}

// PostalCodeLTE applies the LTE predicate on the "postal_code" field.
func PostalCodeLTE(v int) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldPostalCode, v))
}

// CreditUnitsAccumulatedEQ applies the EQ predicate on the "credit_units_accumulated" field.
func CreditUnitsAccumulatedEQ(v int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldCreditUnitsAccumulated, v))
}

// CreditUnitsAccumulatedNEQ applies the NEQ predicate on the "credit_units_accumulated" field.
func CreditUnitsAccumulatedNEQ(v int) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldCreditUnitsAccumulated, v))
}

// CreditUnitsAccumulatedIn applies the In predicate on the "credit_units_accumulated" field.
func CreditUnitsAccumulatedIn(vs ...int) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldCreditUnitsAccumulated, vs...))
}

// CreditUnitsAccumulatedNotIn applies the NotIn predicate on the "credit_units_accumulated" field.
func CreditUnitsAccumulatedNotIn(vs ...int) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldCreditUnitsAccumulated, vs...))
}

// CreditUnitsAccumulatedGT applies the GT predicate on the "credit_units_accumulated" field.
func CreditUnitsAccumulatedGT(v int) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldCreditUnitsAccumulated, v))
}

// CreditUnitsAccumulatedGTE applies the GTE predicate on the "credit_units_accumulated" field.
func CreditUnitsAccumulatedGTE(v int) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldCreditUnitsAccumulated, v))
}

// CreditUnitsAccumulatedLT applies the LT predicate on the "credit_units_accumulated" field.
func CreditUnitsAccumulatedLT(v int) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldCreditUnitsAccumulated, v))
}

// CreditUnitsAccumulatedLTE applies the LTE predicate on the "credit_units_accumulated" field.
func CreditUnitsAccumulatedLTE(v int) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldCreditUnitsAccumulated, v))
}

// TotalAverageEQ applies the EQ predicate on the "total_average" field.
func TotalAverageEQ(v float64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldTotalAverage, v))
}

// TotalAverageNEQ applies the NEQ predicate on the "total_average" field.
func TotalAverageNEQ(v float64) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldTotalAverage, v))
}

// TotalAverageIn applies the In predicate on the "total_average" field.
func TotalAverageIn(vs ...float64) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldTotalAverage, vs...))
}

// TotalAverageNotIn applies the NotIn predicate on the "total_average" field.
func TotalAverageNotIn(vs ...float64) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldTotalAverage, vs...))
}

// TotalAverageGT applies the GT predicate on the "total_average" field.
func TotalAverageGT(v float64) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldTotalAverage, v))
}

// TotalAverageGTE applies the GTE predicate on the "total_average" field.
func TotalAverageGTE(v float64) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldTotalAverage, v))
}

// TotalAverageLT applies the LT predicate on the "total_average" field.
func TotalAverageLT(v float64) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldTotalAverage, v))
}

// TotalAverageLTE applies the LTE predicate on the "total_average" field.
func TotalAverageLTE(v float64) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldTotalAverage, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.Users) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Student) predicate.Student {
	return predicate.Student(sql.NotPredicates(p))
}
