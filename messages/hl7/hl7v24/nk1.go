package hl7v24

import "time"

// HL7 v2.4 - NK1 - Next of kin / associated parties
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/NK1
type NK1 struct {
	SetID                                    int       `hl7:"1" json:"SetID,omitempty"`
	Name                                     []XPN     `hl7:"2" json:"Name,omitempty"`
	Relationship                             CE        `hl7:"3" json:"Relationship,omitempty"`
	Address                                  []XAD     `hl7:"4" json:"Address,omitempty"`
	PhoneNumber                              []XTN     `hl7:"5" json:"PhoneNumber,omitempty"`
	BusinessPhoneNumber                      []XTN     `hl7:"6" json:"BusinessPhoneNumber,omitempty"`
	ContactRole                              CE        `hl7:"7" json:"ContactRole,omitempty"`
	StartDate                                time.Time `hl7:"8" json:"StartDate,omitempty"`
	EndDate                                  time.Time `hl7:"9" json:"EndDate,omitempty"`
	NextOfKinAssociatedPartiesJobTitle       string    `hl7:"10" json:"NextOfKinAssociatedPartiesJobTitle,omitempty"`
	NextOfKinAssociatedPartiesJobCode        CE        `hl7:"11" json:"NextOfKinAssociatedPartiesJobCode,omitempty"`
	NextOfKinAssociatedPartiesEmployeeNumber CX        `hl7:"12" json:"NextOfKinAssociatedPartiesEmployeeNumber,omitempty"`
	OrganizationName                         []XON     `hl7:"13" json:"OrganizationName,omitempty"`
	MaritalStatus                            CE        `hl7:"14" json:"MaritalStatus,omitempty"`
	AdministrativeSex                        string    `hl7:"15" json:"AdministrativeSex,omitempty"`
	DOB                                      time.Time `hl7:"16" json:"DOB,omitempty"`
	LivingDependency                         []string  `hl7:"17" json:"LivingDependency,omitempty"`
	AmbulatoryStatus                         []string  `hl7:"18" json:"AmbulatoryStatus,omitempty"`
	Citzenship                               []CE      `hl7:"19" json:"Citzenship,omitempty"`
	PrimaryLanguage                          CE        `hl7:"20" json:"PrimaryLanguage,omitempty"`
	LivingArrangement                        string    `hl7:"21" json:"LivingArrangement,omitempty"`
	PublicityCode                            CE        `hl7:"22" json:"PublicityCode,omitempty"`
	ProtectionIndicator                      string    `hl7:"23" json:"ProtectionIndicator,omitempty"`
	StudentIndicator                         string    `hl7:"24" json:"StudentIndicator,omitempty"`
	Religion                                 CE        `hl7:"25" json:"Religion,omitempty"`
	MothersMaidenName                        []XPN     `hl7:"26" json:"MothersMaidenName,omitempty"`
	Nationality                              CE        `hl7:"27" json:"Nationality,omitempty"`
	EthnicGroup                              []CE      `hl7:"28" json:"EthnicGroup,omitempty"`
	ContactReason                            []CE      `hl7:"29" json:"ContactReason,omitempty"`
	ContactPersonsName                       []XPN     `hl7:"30" json:"ContactPersonsName,omitempty"`
	ContactPersonsTelephoneNumber            []XTN     `hl7:"31" json:"ContactPersonsTelephoneNumber,omitempty"`
	NextOfKin                                []CX      `hl7:"32" json:"NextOfKin,omitempty"`
	JobStatus                                string    `hl7:"34" json:"JobStatus,omitempty"`
	Race                                     []CE      `hl7:"35" json:"Race,omitempty"`
	Handicap                                 string    `hl7:"36" json:"Handicap,omitempty"`
	ContactPersonSocialSecurityNumber        string    `hl7:"37" json:"ContactPersonSocialSecurityNumber,omitempty"`
}
