package hl7v24

import "time"

// HL7 v2.4 - NK1 - Next of kin / associated parties
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/NK1
type NK1 struct {
	SetID                                    *int      `hl7:"POS=2" json:"SetID,omitempty"`
	Name                                     []XPN     `hl7:"POS=3" json:"Name,omitempty"`
	Relationship                             CE        `hl7:"POS=4" json:"Relationship,omitempty"`
	Address                                  []XAD     `hl7:"POS=5" json:"Address,omitempty"`
	PhoneNumber                              []XTN     `hl7:"POS=6" json:"PhoneNumber,omitempty"`
	BusinessPhoneNumber                      []XTN     `hl7:"POS=7" json:"BusinessPhoneNumber,omitempty"`
	ContactRole                              CE        `hl7:"POS=8" json:"ContactRole,omitempty"`
	StartDate                                time.Time `hl7:"POS=9" json:"StartDate,omitempty"`
	EndDate                                  time.Time `hl7:"POS=10" json:"EndDate,omitempty"`
	NextOfKinAssociatedPartiesJobTitle       string    `hl7:"POS=11" json:"NextOfKinAssociatedPartiesJobTitle,omitempty"`
	NextOfKinAssociatedPartiesJobCode        CE        `hl7:"POS=12" json:"NextOfKinAssociatedPartiesJobCode,omitempty"`
	NextOfKinAssociatedPartiesEmployeeNumber CX        `hl7:"POS=13" json:"NextOfKinAssociatedPartiesEmployeeNumber,omitempty"`
	OrganizationName                         []XON     `hl7:"POS=14" json:"OrganizationName,omitempty"`
	MaritalStatus                            CE        `hl7:"POS=15" json:"MaritalStatus,omitempty"`
	AdministrativeSex                        string    `hl7:"POS=16" json:"AdministrativeSex,omitempty"`
	DOB                                      time.Time `hl7:"POS=17" json:"DOB,omitempty"`
	LivingDependency                         []string  `hl7:"POS=18" json:"LivingDependency,omitempty"`
	AmbulatoryStatus                         []string  `hl7:"POS=19" json:"AmbulatoryStatus,omitempty"`
	Citzenship                               []CE      `hl7:"POS=20" json:"Citzenship,omitempty"`
	PrimaryLanguage                          CE        `hl7:"POS=21" json:"PrimaryLanguage,omitempty"`
	LivingArrangement                        string    `hl7:"POS=22" json:"LivingArrangement,omitempty"`
	PublicityCode                            CE        `hl7:"POS=23" json:"PublicityCode,omitempty"`
	ProtectionIndicator                      string    `hl7:"POS=24" json:"ProtectionIndicator,omitempty"`
	StudentIndicator                         string    `hl7:"POS=25" json:"StudentIndicator,omitempty"`
	Religion                                 CE        `hl7:"POS=26" json:"Religion,omitempty"`
	MothersMaidenName                        []XPN     `hl7:"POS=27" json:"MothersMaidenName,omitempty"`
	Nationality                              CE        `hl7:"POS=28" json:"Nationality,omitempty"`
	EthnicGroup                              []CE      `hl7:"POS=29" json:"EthnicGroup,omitempty"`
	ContactReason                            []CE      `hl7:"POS=30" json:"ContactReason,omitempty"`
	ContactPersonsName                       []XPN     `hl7:"POS=31" json:"ContactPersonsName,omitempty"`
	ContactPersonsTelephoneNumber            []XTN     `hl7:"POS=32" json:"ContactPersonsTelephoneNumber,omitempty"`
	NextOfKin                                []CX      `hl7:"POS=33" json:"NextOfKin,omitempty"`
	JobStatus                                string    `hl7:"POS=35" json:"JobStatus,omitempty"`
	Race                                     []CE      `hl7:"POS=36" json:"Race,omitempty"`
	Handicap                                 string    `hl7:"POS=37" json:"Handicap,omitempty"`
	ContactPersonSocialSecurityNumber        string    `hl7:"POS=38" json:"ContactPersonSocialSecurityNumber,omitempty"`
}
