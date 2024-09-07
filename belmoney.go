package belmoney

type StatusId int

const (
	StatusIdCreated                   StatusId = 1
	StatusIdHold                      StatusId = 2
	StatusIdDone                      StatusId = 3
	StatusIdTransmittedToPayee        StatusId = 4
	StatusIdUpdated                   StatusId = 5
	StatusIdVoid                      StatusId = 6
	StatusIdPaid                      StatusId = 7
	StatusIdCancel                    StatusId = 8
	StatusIdRepair                    StatusId = 9
	StatusIdRepaired                  StatusId = 10
	StatusIdProcessingPayment         StatusId = 11
	StatusIdSentToRepair              StatusId = 12
	StatusIdPaidTransmitted           StatusId = 13
	StatusIdProcessingRepaired        StatusId = 14
	StatusIdRepairedPayer             StatusId = 15
	StatusIdReviewByCustomerService   StatusId = 16
	StatusIdVoidAgent                 StatusId = 17
	StatusIdReviewedByCustomerService StatusId = 18
	StatusIdResponseFromPayer         StatusId = 19
	StatusIdRefunded                  StatusId = 20
	StatusIdRefundedToAgent           StatusId = 21
	StatusIdHoldIncomingPayment       StatusId = 22
	StatusIdClearedToPay              StatusId = 23
	StatusIdOutForDelivery            StatusId = 24
	StatusIdFundsReceived             StatusId = 25
)

type AccountTypeId int

const (
	AccountTypeIdSelect           AccountTypeId = 0
	AccountTypeIdChecking         AccountTypeId = 1
	AccountTypeIdSavings          AccountTypeId = 2
	AccountTypeIdBusinessChecking AccountTypeId = 3
	AccountTypeIdBusinessSavings  AccountTypeId = 4
	AccountTypeIdOther            AccountTypeId = 5
)

type PersonIdType int

const (
	PersonIdTypeNone                      PersonIdType = 0
	PersonIdTypeDriversLicense            PersonIdType = 1
	PersonIdTypeSSN                       PersonIdType = 2
	PersonIdTypeAlienRegistrationCard     PersonIdType = 3
	PersonIdTypeStateIssuedCard           PersonIdType = 4
	PersonIdTypePassport                  PersonIdType = 5
	PersonIdTypeForeignPassport           PersonIdType = 6
	PersonIdTypeOther                     PersonIdType = 7
	PersonIdTypeCPF                       PersonIdType = 8
	PersonIdTypeFederalIDNumber           PersonIdType = 9
	PersonIdTypeCNPJ                      PersonIdType = 10
	PersonIdTypeInsuranceNumber           PersonIdType = 11
	PersonIdTypeIDIssuedCard              PersonIdType = 12
	PersonIdTypeUtilityBill               PersonIdType = 13
	PersonIdTypeNationalIdentityCard      PersonIdType = 14
	PersonIdTypeNationalResidenceCard     PersonIdType = 15
	PersonIdTypeForeignIdentityCard       PersonIdType = 16
	PersonIdTypeEuropeanIdentityCard      PersonIdType = 17
	PersonIdTypeEuropeanResidenceCard     PersonIdType = 18
	PersonIdTypeCertificateOfRegistration PersonIdType = 19
	PersonIdTypeFirearmsLicense           PersonIdType = 20
)

type DocumentTypeId int

const (
	DocumentTypeIdGeneral        DocumentTypeId = 0
	DocumentTypeIdIdentification DocumentTypeId = 1
	DocumentTypeIdProofOfAddress DocumentTypeId = 2
	DocumentTypeIdCompliance     DocumentTypeId = 3
	DocumentTypeIdKYCForm        DocumentTypeId = 4
)

type PaymentTypeId int

const (
	PaymentTypeIdAccountDeposit PaymentTypeId = 1
	PaymentTypeIdOfficePickUp   PaymentTypeId = 2
	PaymentTypeIdBankPickUp     PaymentTypeId = 3
	PaymentTypeIdHomeDelivery   PaymentTypeId = 4
	PaymentTypeIdCardTopUp      PaymentTypeId = 5
	PaymentTypeIdWalletTopUp    PaymentTypeId = 6
	PaymentTypeIdPix            PaymentTypeId = 7
)

type GenderCode string

const (
	GenderCodeMale   GenderCode = "M"
	GenderCodeFemale GenderCode = "F"
)

type RateTypeId int

const (
	RateTypeIdGeneral RateTypeId = 0
)

type TransferReasonId int

const (
	TransferReasonIdRemitToBusiness TransferReasonId = 1
	TransferReasonIdRemitToFamily   TransferReasonId = 2
	TransferReasonIdOther           TransferReasonId = 3
	TransferReasonIdSavingsAccount  TransferReasonId = 5
	TransferReasonIdBills           TransferReasonId = 7
	TransferReasonIdBuyProperty     TransferReasonId = 9
	TransferReasonIdChurchDonation  TransferReasonId = 11 // [DIZIMO]
	TransferReasonIdPayVacations    TransferReasonId = 15
	TransferReasonIdSchoolPayment   TransferReasonId = 17
	TransferReasonIdPayWedding      TransferReasonId = 23
	TransferReasonIdRemitToFriend   TransferReasonId = 25
	TransferReasonIdPayCreditCard   TransferReasonId = 27
	TransferReasonIdPayDebt         TransferReasonId = 29
	TransferReasonIdGift            TransferReasonId = 35
)

type RepairReasonId int

const (
	RepairReasonIdWrongAccountNo           RepairReasonId = 1
	RepairReasonIdWrongAgency              RepairReasonId = 2
	RepairReasonIdWrongBank                RepairReasonId = 3
	RepairReasonIdIncorrectBeneficiaryName RepairReasonId = 4
	RepairReasonIdWrongAddress             RepairReasonId = 5
	RepairReasonIdWrongTelephoneNo         RepairReasonId = 6
	RepairReasonIdMissingID                RepairReasonId = 7
	RepairReasonIdCanceledOrSuspendedID    RepairReasonId = 8
	RepairReasonIdInvalidID                RepairReasonId = 9
	RepairReasonIdOthers                   RepairReasonId = 10 // Inform in Note field
)

type CancelReasonId int

const (
	CancelReasonIdGeneral CancelReasonId = 0
)

type RelationTypeId int

const (
	RelationTypeIdFather             RelationTypeId = 1
	RelationTypeIdMother             RelationTypeId = 2
	RelationTypeIdSiblings           RelationTypeId = 3
	RelationTypeIdSonDaughter        RelationTypeId = 4
	RelationTypeIdHusbandWife        RelationTypeId = 5
	RelationTypeIdCousin             RelationTypeId = 6
	RelationTypeIdUncleAunt          RelationTypeId = 7
	RelationTypeIdGrandparent        RelationTypeId = 8
	RelationTypeIdFriend             RelationTypeId = 9
	RelationTypeIdOther              RelationTypeId = 10
	RelationTypeIdBrotherSisterInLaw RelationTypeId = 11
	RelationTypeIdFatherMotherInLaw  RelationTypeId = 12
	RelationTypeIdNephewNiece        RelationTypeId = 13
	RelationTypeIdOwn                RelationTypeId = 14
	RelationTypeIdExHusbandWife      RelationTypeId = 15
	RelationTypeIdGrandchild         RelationTypeId = 16
	RelationTypeIdSonDaughterInLaw   RelationTypeId = 17
	RelationTypeIdBusinessPartner    RelationTypeId = 18
)

type PEPTypeId int

const (
	PEPTypeIdYes PEPTypeId = 1
	PEPTypeIdNo  PEPTypeId = 2
)

type SenderTypeId int

const (
	SenderTypeIdCorporate  SenderTypeId = 0
	SenderTypeIdIndividual SenderTypeId = 1
)

type ProfessionId int
