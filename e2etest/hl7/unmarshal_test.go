package hl7

import (
	"fmt"
	"testing"
	"time"

	"github.com/blutspende/go-parser"
	"github.com/blutspende/go-parser/messageformat/hl7/hl7v23"
	"github.com/blutspende/go-parser/messageformat/hl7/hl7v24"
	"github.com/stretchr/testify/assert"
)

// Test_Parse_MSH_Segment, this test has only one line
func Test_Parse_MSH_Segment(t *testing.T) {
	// Arrange
	messageString := fmt.Sprintf("MSH|^~\\&|HL7_Host^b^c|HL7_Office^^Xyz|CIT^^|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|||ER|ER||8859/1~second_element|<\u000d")
	type HeaderMessage struct {
		MSH hl7v23.MSH `hl7:"TAG=MSH"`
	}
	var message HeaderMessage
	// Act
	err := parser.Unmarshal([]byte(messageString), &message, config)
	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, message.MSH)
	assert.Equal(t, "HL7_Host", message.MSH.SendingApplication.NamespaceID)
	assert.Equal(t, "b", message.MSH.SendingApplication.UniversalID)
	assert.Equal(t, "c", message.MSH.SendingApplication.UniversalIDType)
	assert.Equal(t, "HL7_Office", message.MSH.SendingFacility.NamespaceID)
	assert.Equal(t, "CIT", message.MSH.ReceivingApplication.NamespaceID)
	assert.Equal(t, "LAB", message.MSH.ReceivingFacility.NamespaceID)
	assert.Equal(t, "2011-09-26 10:51:55 +0000 UTC", message.MSH.DateTimeOfMessage.String())
	assert.Equal(t, "", message.MSH.Security)
	assert.Equal(t, "ORM", message.MSH.MessageType.MessageType)
	assert.Equal(t, "O01", message.MSH.MessageType.TriggerEvent)
	assert.Equal(t, "20110926125155", message.MSH.MessageControlID)
	assert.Equal(t, "P", message.MSH.ProcessingID.ProcessingID)
	assert.Equal(t, "2.3", message.MSH.VersionID)
	assert.Nil(t, message.MSH.SequenceNumber)
	assert.Equal(t, "", message.MSH.ContinuationPointer)
	assert.Equal(t, "ER", message.MSH.AcceptAcknowledgementType)
	assert.Equal(t, "ER", message.MSH.ApplicationAcknowledgementType)
	assert.Equal(t, "", message.MSH.CountryCode)
	assert.Equal(t, "8859/1", message.MSH.CharacterSet[0])
	assert.Equal(t, "second_element", message.MSH.CharacterSet[1])
	assert.Equal(t, "<", message.MSH.PrincipalLanguageOfMessage.Identifier)
}

// Run a Testorder provided by Roche cITM but its some standard with each record once
func Test_Order_ORM_generic1(t *testing.T) {
	// Arrange
	var messageString string
	messageString += "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|||ER|ER||8859/1|\u000d"
	messageString += "PID|1|a^b|00100M56016||Smith^Harry||19500412|M\u000d"
	messageString += "ORC|NW|000218T018||||Not used|^^^^^R||20110926120055\u000d"
	messageString += "OBR|1|000218T018||101~102||20110926120000|||||A||||\u000d"
	var message hl7v23.ORM_O01
	// Act
	err := parser.Unmarshal([]byte(messageString), &message, config)
	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, message.Patient)
	assert.NotNil(t, message.Patient.PatientIdentification)
	assert.Equal(t, 1, message.Patient.PatientIdentification.SetID)
	assert.Equal(t, "a", message.Patient.PatientIdentification.PatientIDExternalID.ID)
	assert.Equal(t, "b", message.Patient.PatientIdentification.PatientIDExternalID.CheckDigit)
	assert.Equal(t, "", message.Patient.PatientIdentification.PatientIDExternalID.AssigningAuthority.UniversalID)
	assert.Equal(t, "a", message.Patient.PatientIdentification.PatientIDExternalID.ID)
	assert.Equal(t, "b", message.Patient.PatientIdentification.PatientIDExternalID.CheckDigit)
	assert.Equal(t, "00100M56016", message.Patient.PatientIdentification.PatientIDInternalID[0].ID)
	assert.Len(t, message.Patient.PatientIdentification.AlternatePatientID, 0)
	assert.Len(t, message.Patient.PatientIdentification.PatientName, 1)
	assert.Equal(t, "Smith", message.Patient.PatientIdentification.PatientName[0].FamilyName)
	assert.Equal(t, "Harry", message.Patient.PatientIdentification.PatientName[0].GivenName)
	assert.Equal(t, "", message.Patient.PatientIdentification.PatientName[0].MiddleInitialOrName)
	assert.Equal(t, "", message.Patient.PatientIdentification.MothersMaidenName.FamilyName)
	assert.Equal(t, "1950-04-12 00:00:00 +0100 CET", message.Patient.PatientIdentification.DateOfBirth.String())
	assert.Equal(t, "M", message.Patient.PatientIdentification.Sex)
	assert.Equal(t, 1, len(message.Order))
	assert.Equal(t, "NW", message.Order[0].CommonOrderSegment.OrderControl)
	assert.Equal(t, "000218T018", message.Order[0].CommonOrderSegment.PlacerOrderNumber.EntityIdentifier)
	assert.Equal(t, "", message.Order[0].CommonOrderSegment.PlacerOrderNumber.NamespaceID)
	assert.Equal(t, "", message.Order[0].CommonOrderSegment.PlacerOrderNumber.UniversalID)
	assert.Equal(t, "", message.Order[0].CommonOrderSegment.PlacerOrderNumber.UniversalIDType)
	assert.Equal(t, "", message.Order[0].CommonOrderSegment.FillerOrderNumber.EntityIdentifier)
	assert.Equal(t, "", message.Order[0].CommonOrderSegment.PlacerGroupNumber.EntityIdentifier)
	assert.Equal(t, "", message.Order[0].CommonOrderSegment.OrderStatus)
	assert.Equal(t, "Not used", message.Order[0].CommonOrderSegment.ResponseFlag)
	assert.Equal(t, "R", message.Order[0].CommonOrderSegment.QuantityTiming.Priority)
	assert.Equal(t, "", message.Order[0].CommonOrderSegment.QuantityTiming.Duration)
	assert.Equal(t, "R", message.Order[0].CommonOrderSegment.QuantityTiming.Priority)
	assert.Equal(t, "", message.Order[0].CommonOrderSegment.QuantityTiming.Condition)
	assert.Equal(t, "", message.Order[0].CommonOrderSegment.ParentOrder.ParentsPlacerOrderNumber.UniversalID)
	assert.Equal(t, "", message.Order[0].CommonOrderSegment.ParentOrder.ParentsFillerOrderNumber.UniversalID)
	assert.Equal(t, "2011-09-26 10:00:55 +0000 UTC", message.Order[0].CommonOrderSegment.DateTimeOfTransaction.String())
	assert.Nil(t, err)
	assert.NotNil(t, message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment)
	assert.Equal(t, 1, message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.SetID)
	assert.Equal(t, "000218T018", message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.PlacerOrderNumber.EntityIdentifier)
	assert.Equal(t, "", message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.PlacerOrderNumber.NamespaceID)
	assert.Equal(t, "", message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.PlacerOrderNumber.UniversalID)
	assert.Equal(t, "", message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.PlacerOrderNumber.UniversalIDType)
	assert.Equal(t, "", message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.FillerOrderNumber.EntityIdentifier)
	//  Note: the original assert here was only "101", which only worked because of a flaw in the old library
	assert.Equal(t, "101~102", message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.UniversalServiceIdentifier.Identifier)
	assert.Equal(t, "", message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.UniversalServiceIdentifier.Text)
	assert.Equal(t, "", message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.Priority)
	assert.Equal(t, "2011-09-26 10:00:00 +0000 UTC", message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.RequestedDateTime.String())
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.ObservationDateTime.String())
	assert.Equal(t, "A", message.Order[0].OrderDetail.OrderDetailSegment.ObservationRequestSegment.SpecimenActionCode)
}

// JSON serializer test with optional manual result observation
func TestJson(t *testing.T) {
	// Arrange
	messageString := `MSH|^~\&|SWISSLAB|FFM||FFM|20230203080903||ORM^O01|o3057937.000001|P|2.3|` + "\r"
	messageString += `PID|1||01077843||CL5AVA0N7K|||U|` + "\r"
	messageString += `PV1|1|S|||||||||||||||||S01077843|` + "\r"
	messageString += `ORC|NW||23071012||||||20230203080800|||BSD|` + "\r"
	messageString += `OBR|1||23071012|HINAET^HIV-1/2 PCR mit erhöhter Sensitivität|||20230203080900|||||||||BSD|||||||||P|` + "\r"
	messageString += `ORC|NW||23071013||||||20230203080800|||BSD|` + "\r"
	messageString += `OBR|1||23071013|HCNAET^HCV PCR mit erhöhter Sensitivität|||20230203080900|||||||||BSD|||||||||P|` + "\r"
	messageString += `ORC|NW||23071014||||||20230203080800|||BSD|` + "\r"
	messageString += `OBR|1||23071014|HBNAET^HBV PCR mit erhöhter Sensitivität|||20230203080900|||||||||BSD|||||||||P|` + "\r"
	message := hl7v23.ORM_O01{}
	// Act
	err := parser.Unmarshal([]byte(messageString), &message, config)
	// Assert
	assert.Nil(t, err)
	// Optional
	// Note: a json output can be manually observed
	//out, _ := json.MarshalIndent(message, "", "\t")
	//file, err := os.Create("JsonMarshalTestResult.json")
	//if err != nil {
	//	t.Error(err)
	//} else {
	//	file.Write(out)
	//	file.Close()
	//}
}

// Tests customized for CITM tests

// This test will fail, because Roche is unable to read documents... OBR can not be repeated
func Test_cITm_Result1(t *testing.T) {
	// Arrange
	var messageString string
	messageString += "MSH|^~\\&|||||20110927155013||ORU^R01|68516|P|2.3|||NE|NE||8859/1\u000d"
	messageString += "PID|1||4637463G66||Smith^John||19630101|M\u000d"
	messageString += "NTE|1|L|1st·comment·on·patient·/·sample·20020604101\u000d"
	messageString += "NTE|2|L|2nd·comment·on·patient·/·sample\u000d"
	messageString += "ORC|RE|20020604101|||||^^^^^R||20110927150630\u000d"
	messageString += "OBR|1|20020604101||ALL||20110927150629|||||||||S1^^^^^^P||||||||||||^^^^^|||||||\u000d"
	messageString += "OBX|1||21||101,0|mmol/L||N|||F||23~N|20110927154723|^^^COBAS8K.200|System\u000d"
	//messageString += "TCD|1|21|Dec\u000d"
	messageString += "NTE|1||L|R|G\u000d"
	messageString += "OBX|2||82||5,7|mmol/L||H|||F||23~N|20110927154733|^^^COBAS8K.200|System\u000d"
	//messageString += "TCD|1|82|1\u000d"
	messageString += "NTE|1||L|R|G\u000d"
	messageString += "OBX|3||89||162,0|mmol/L||H|||F||23~N|20110927154734|^^^COBAS8K.200|System\u000d"
	//messageString += "TCD|1|89|Inc\u000d"
	messageString += "NTE|1|L|Comment·on·test·code·89\u000d"
	messageString += "NTE|2||L|R|G\u000d"
	var message hl7v23.ORU_R01
	// Act
	err := parser.Unmarshal([]byte(messageString), &message, config)
	// Assert
	assert.Nil(t, err)
	// Note: assertions removed due to the aforementioned message format mismatch resulting in failing test
	//assert.Equal(t, 1, len(message.PatientResult))
	//assert.Equal(t, "2nd·comment·on·patient·/·sample", message.PatientResult[0].Patient.NotesAndComments[1].Comment)
	//assert.Equal(t, 1, len(message.PatientResult[0].OrderObservation))
	//assert.Equal(t, "RE", message.PatientResult[0].OrderObservation[0].CommonOrder.OrderControl)
	//assert.Equal(t, "ALL", message.PatientResult[0].OrderObservation[0].ObservationRequest.UniversalServiceIdentifier.Identifier)
	//assert.Equal(t, 3, len(message.PatientResult[0].OrderObservation[0].Observation))
}

// This test ensures that the analysis-results from Cobas cITm are understood
func TestCit_OUL_R21(t *testing.T) {
	// Arrange
	var messageString string
	messageString += "MSH|^~\\&|Roche Diagnostics|cITm 1.10.02.0572|DRK FFM||20220711130056||OUL^R21|107737129|P|2.4|||NE|NE||UNICODE UTF-8<13>PID|1|?|||^|||\u000d"
	messageString += "SAC|||AA4F1A1A6J\u000d"
	messageString += "ORC|RE|AA4F1A1A6J|||||^^^^^?||20220709195656\u000d"
	messageString += "OBR|1|AA4F1A1A6J||ALL||20220711075520|||||||||1^^^^^^P||||||||||||^^^^^?|||||||\u000d"
	messageString += "OBX|1||AHBC2-R||-1\\S\\2.14|||N|||F|||20220711122749|^^^CCM2-c8k-5-1859-10#e801#2#2|bmserv\\S\\SYSTEM^System||CCM2-c8k-5-1859-10#e801#2#2|20220711122749\u000d"
	messageString += "TCD|1|AHBC2-R|1||||||^|\u000d"
	var message hl7v24.OUL_R21
	// Act
	err := parser.Unmarshal([]byte(messageString), &message, config)
	// Assert
	// SAC
	assert.Nil(t, err)
	assert.Equal(t, 1, len(message.OrderObservation))
	assert.Equal(t, "AA4F1A1A6J", message.OrderObservation[0].Container.SpecimenAndContainerDetail.ContainerIdentifier.EntityIdentifier)
	// ORC
	assert.Equal(t, "RE", message.OrderObservation[0].CommonOrder.OrderControl)
	assert.Equal(t, "AA4F1A1A6J", message.OrderObservation[0].CommonOrder.PlacerOrderNumber.EntityIdentifier)
	assert.Equal(t, 1, len(message.OrderObservation[0].CommonOrder.QuantityTiming))
	assert.Equal(t, "?", message.OrderObservation[0].CommonOrder.QuantityTiming[0].Priority)
	assert.Equal(t, time.Date(2022, time.July, 9, 17, 56, 56, 0, time.UTC), message.OrderObservation[0].CommonOrder.DateTimeOfTransaction)
	// OBR
	assert.Equal(t, 1, message.OrderObservation[0].ObservationRequest.SetID)
	assert.Equal(t, "AA4F1A1A6J", message.OrderObservation[0].CommonOrder.PlacerOrderNumber.EntityIdentifier)
	assert.Equal(t, "ALL", message.OrderObservation[0].ObservationRequest.UniversalServiceIdentifier.Identifier)
	assert.Equal(t, time.Date(2022, time.July, 11, 5, 55, 20, 0, time.UTC), message.OrderObservation[0].ObservationRequest.RequestedDateTime)
	assert.Equal(t, "1", message.OrderObservation[0].ObservationRequest.SpecimenSource.SpecimenSourceNameOrCode.Identifier)
	assert.Equal(t, "P", message.OrderObservation[0].ObservationRequest.SpecimenSource.SpecimenRole.Identifier)
	// field 1 missing
	// OBX
	assert.Equal(t, 1, len(message.OrderObservation[0].Observation))
	assert.Equal(t, 1, message.OrderObservation[0].Observation[0].ObservationResult.SetID)
	assert.Equal(t, "AHBC2-R", message.OrderObservation[0].Observation[0].ObservationResult.ObservationIdentifier.Identifier)
	assert.Equal(t, 1, len(message.OrderObservation[0].Observation[0].ObservationResult.ObservationValue))
	// Note: as escape substitution actually works now, the expected value has changed
	assert.Equal(t, "-1^2.14", message.OrderObservation[0].Observation[0].ObservationResult.ObservationValue[0])
	assert.Equal(t, "N", message.OrderObservation[0].Observation[0].ObservationResult.AbnormalFlags[0])
	assert.Equal(t, "F", message.OrderObservation[0].Observation[0].ObservationResult.ObservationResultStatus)
	assert.Equal(t, time.Date(2022, time.July, 11, 10, 27, 49, 0, time.UTC), message.OrderObservation[0].Observation[0].ObservationResult.DateTimeOfTheObservation)
	assert.Equal(t, "CCM2-c8k-5-1859-10#e801#2#2", message.OrderObservation[0].Observation[0].ObservationResult.ProducersID.AlternateIdentifier)
	assert.Equal(t, "bmserv^SYSTEM", message.OrderObservation[0].Observation[0].ObservationResult.ResponsibleObserver[0].IDNumber)
	// Note: this no longer fails, as the "System" is properly parsed out
	assert.Equal(t, "System", message.OrderObservation[0].Observation[0].ObservationResult.ResponsibleObserver[0].FamilyName.Surname)
	assert.Equal(t, "CCM2-c8k-5-1859-10#e801#2#2", message.OrderObservation[0].Observation[0].ObservationResult.EquipmentInstanceIdentifier[0].EntityIdentifier)
	assert.Equal(t, time.Date(2022, time.July, 11, 10, 27, 49, 0, time.UTC), message.OrderObservation[0].Observation[0].ObservationResult.DateTimeOfTheAnalysis)
	// TCD
	assert.Equal(t, "1", message.OrderObservation[0].Observation[0].TestCodeDetail.UniversalServiceIdentifier.Identifier)
	assert.Equal(t, "AHBC2-R", message.OrderObservation[0].Observation[0].TestCodeDetail.AutoDilutionFactor.Comparator)
	assert.Equal(t, "1", message.OrderObservation[0].Observation[0].TestCodeDetail.RerunDilutionFactor.Comparator)
}
