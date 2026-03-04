package hl7v28

import "time"

// DON - Donation Segment
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/DON
type DON struct {
	DonationIdentificationNumberDin EI `hl7:"POS=2"`
	DonationType CNE `hl7:"POS=3"`
	PhlebotomyStartDateTime time.Time `hl7:"POS=4"`
	PhlebotomyEndDateTime time.Time `hl7:"POS=5"`
	DonationDuration *int `hl7:"POS=6"`
	DonationDurationUnits CNE `hl7:"POS=7"`
	IntendedProcedureType []CNE `hl7:"POS=8"`
	ActualProcedureType []CNE `hl7:"POS=9"`
	DonorEligibilityFlag string `hl7:"POS=10"`
	DonorEligibilityProcedureType []CNE `hl7:"POS=11"`
	DonorEligibilityDate time.Time `hl7:"POS=12"`
	ProcessInterruption CNE `hl7:"POS=13"`
	ProcessInterruptionReason CNE `hl7:"POS=14"`
	PhlebotomyIssue []CNE `hl7:"POS=15"`
	IntendedRecipientBloodRelative string `hl7:"POS=16"`
	IntendedRecipientName XPN `hl7:"POS=17"`
	IntendedRecipientDob time.Time `hl7:"POS=18"`
	IntendedRecipientFacility XON `hl7:"POS=19"`
	IntendedRecipientProcedureDate time.Time `hl7:"POS=20"`
	IntendedRecipientOrderingProvider XPN `hl7:"POS=21"`
	PhlebotomyStatus CNE `hl7:"POS=22"`
	ArmStick CNE `hl7:"POS=23"`
	BleedStartPhlebotomist XPN `hl7:"POS=24"`
	BleedEndPhlebotomist XPN `hl7:"POS=25"`
	AphaeresisTypeMachine string `hl7:"POS=26"`
	AphaeresisMachineSerialNumber string `hl7:"POS=27"`
	DonorReaction string `hl7:"POS=28"`
	FinalReviewStaffID XPN `hl7:"POS=29"`
	FinalReviewDateTime time.Time `hl7:"POS=30"`
	NumberOfTubesCollected *int `hl7:"POS=31"`
	DonationSampleIdentifier []EI `hl7:"POS=32"`
	DonationAcceptStaff XCN `hl7:"POS=33"`
	DonationMaterialReviewStaff []XCN `hl7:"POS=34"`
}

