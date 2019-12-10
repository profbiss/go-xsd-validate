package xsdvalidate

import (
	"fmt"
	"testing"
)

func TestSchemaValidate(t *testing.T) {
	xsdhandler, err := NewXsdHandlerUrl("https://www.fincen.gov/sites/default/files/schema/base/BSA_XML_2.0.xsd", ParsErrDefault)
	if err != nil {
		fmt.Printf("Error: %s %s\n", t.Name(), err.Error())
		t.Fail()
	}
	defer xsdhandler.Free()

}

const xml = `<?xml version="1.0" encoding="UTF-8"?>
<fc2:EFilingBatchXML TotalAmount="0" PartyCount="6" ActivityCount="1" xsi:schemaLocation="www.fincen.gov/base https://www.fincen.gov/sites/default/files/schema/base/EFL_CTRXBatchSchema.xsd"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xmlns:fc2="www.fincen.gov/base">
    <fc2:FormTypeCode>CTRX</fc2:FormTypeCode>
    <fc2:Activity SeqNum="1">
        <fc2:FilingDateText></fc2:FilingDateText>
        <fc2:ActivityAssociation SeqNum="2">
            <fc2:CorrectsAmendsPriorReportIndicator></fc2:CorrectsAmendsPriorReportIndicator>
            <fc2:FinCENDirectBackFileIndicator></fc2:FinCENDirectBackFileIndicator>
            <fc2:InitialReportIndicator>Y</fc2:InitialReportIndicator>
        </fc2:ActivityAssociation>
        <fc2:Party SeqNum="3">
            <fc2:ActivityPartyTypeCode>35</fc2:ActivityPartyTypeCode>
            <fc2:EFilingCoverageBeginningDateText>20191016</fc2:EFilingCoverageBeginningDateText>
            <fc2:EFilingCoverageEndDateText>20191016</fc2:EFilingCoverageEndDateText>
            <fc2:PartyName SeqNum="4">
                <fc2:PartyNameTypeCode>L</fc2:PartyNameTypeCode>
                <fc2:RawPartyFullName>Transmitter Name</fc2:RawPartyFullName>
            </fc2:PartyName>
            <fc2:Address SeqNum="5">
                <fc2:RawCityText>Transmitter City</fc2:RawCityText>
                <fc2:RawCountryCodeText>TC</fc2:RawCountryCodeText>
                <fc2:RawStateCodeText>TS</fc2:RawStateCodeText>
                <fc2:RawStreetAddress1Text>Transmitter street</fc2:RawStreetAddress1Text>
                <fc2:RawZIPCode>89098</fc2:RawZIPCode>
            </fc2:Address>
            <fc2:PhoneNumber SeqNum="6">
                <fc2:PhoneNumberText>1231233423</fc2:PhoneNumberText>
            </fc2:PhoneNumber>
            <fc2:PartyIdentification SeqNum="7">
                <fc2:PartyIdentificationNumberText>123456789</fc2:PartyIdentificationNumberText>
                <fc2:PartyIdentificationTypeCode>4</fc2:PartyIdentificationTypeCode>
            </fc2:PartyIdentification>
            <fc2:PartyIdentification SeqNum="8">
                <fc2:PartyIdentificationNumberText>987654321</fc2:PartyIdentificationNumberText>
                <fc2:PartyIdentificationTypeCode>28</fc2:PartyIdentificationTypeCode>
            </fc2:PartyIdentification>
        </fc2:Party>
        <fc2:Party SeqNum="9">
            <fc2:ActivityPartyTypeCode>37</fc2:ActivityPartyTypeCode>
            <fc2:EFilingCoverageBeginningDateText>20191016</fc2:EFilingCoverageBeginningDateText>
            <fc2:EFilingCoverageEndDateText>20191016</fc2:EFilingCoverageEndDateText>
            <fc2:PartyName SeqNum="10">
                <fc2:PartyNameTypeCode>L</fc2:PartyNameTypeCode>
                <fc2:RawPartyFullName>Transmitter Contact Name</fc2:RawPartyFullName>
            </fc2:PartyName>
        </fc2:Party>
        <fc2:Party SeqNum="11">
            <fc2:ActivityPartyTypeCode>30</fc2:ActivityPartyTypeCode>
            <fc2:PrimaryRegulatorTypeCode>1</fc2:PrimaryRegulatorTypeCode>
            <fc2:PartyName SeqNum="12">
                <fc2:PartyNameTypeCode>L</fc2:PartyNameTypeCode>
                <fc2:RawPartyFullName>THE COLORADO BANK &amp; TRUST COMPANY OF LA JUNTA</fc2:RawPartyFullName>
            </fc2:PartyName>
            <fc2:PartyName SeqNum="13">
                <fc2:PartyNameTypeCode>DBA</fc2:PartyNameTypeCode>
                <fc2:RawPartyFullName>THE COLORADO BANK &amp; TRUST CO OF LA JUNTA</fc2:RawPartyFullName>
            </fc2:PartyName>
            <fc2:Address SeqNum="14">
                <fc2:RawCityText>LA JUNTA</fc2:RawCityText>
                <fc2:RawCountryCodeText>US</fc2:RawCountryCodeText>
                <fc2:RawStateCodeText>CO</fc2:RawStateCodeText>
                <fc2:RawStreetAddress1Text>301 COLORADO AVE</fc2:RawStreetAddress1Text>
                <fc2:RawZIPCode>810503646</fc2:RawZIPCode>
            </fc2:Address>
            <fc2:PartyIdentification SeqNum="15">
                <fc2:PartyIdentificationNumberText>840174015</fc2:PartyIdentificationNumberText>
                <fc2:PartyIdentificationTypeCode>2</fc2:PartyIdentificationTypeCode>
            </fc2:PartyIdentification>
            <fc2:PartyIdentification SeqNum="16">
                <fc2:PartyIdentificationNumberText>285151</fc2:PartyIdentificationNumberText>
                <fc2:PartyIdentificationTypeCode>14</fc2:PartyIdentificationTypeCode>
            </fc2:PartyIdentification>
        </fc2:Party>
        <fc2:Party SeqNum="17">
            <fc2:ActivityPartyTypeCode>8</fc2:ActivityPartyTypeCode>
            <fc2:PartyName SeqNum="18">
                <fc2:PartyNameTypeCode>L</fc2:PartyNameTypeCode>
                <fc2:RawPartyFullName>LA JUNTA</fc2:RawPartyFullName>
            </fc2:PartyName>
            <fc2:PhoneNumber SeqNum="19">
                <fc2:PhoneNumberText>7193848131</fc2:PhoneNumberText>
            </fc2:PhoneNumber>
        </fc2:Party>
        <fc2:Party SeqNum="20">
            <fc2:ActivityPartyTypeCode>34</fc2:ActivityPartyTypeCode>
            <fc2:IndividualEntityCashInAmountText>67780</fc2:IndividualEntityCashInAmountText>
            <fc2:PrimaryRegulatorTypeCode>1</fc2:PrimaryRegulatorTypeCode>
            <fc2:PartyName SeqNum="21">
                <fc2:PartyNameTypeCode>L</fc2:PartyNameTypeCode>
                <fc2:RawPartyFullName>THE COLORADO BANK &amp; TRUST COMPANY OF LA JUNTA</fc2:RawPartyFullName>
            </fc2:PartyName>
            <fc2:PartyName SeqNum="22">
                <fc2:PartyNameTypeCode>DBA</fc2:PartyNameTypeCode>
                <fc2:RawPartyFullName>THE COLORADO BANK &amp; TRUST CO OF LA JUNTA</fc2:RawPartyFullName>
            </fc2:PartyName>
            <fc2:Address SeqNum="23">
                <fc2:RawCityText>PUEBLO</fc2:RawCityText>
                <fc2:RawCountryCodeText>US</fc2:RawCountryCodeText>
                <fc2:RawStateCodeText>CO</fc2:RawStateCodeText>
                <fc2:RawStreetAddress1Text>23051 E US HIGHWAY 50</fc2:RawStreetAddress1Text>
                <fc2:RawZIPCode>810063005</fc2:RawZIPCode>
            </fc2:Address>
            <fc2:PartyIdentification SeqNum="24">
                <fc2:PartyIdentificationNumberText>840174015</fc2:PartyIdentificationNumberText>
                <fc2:PartyIdentificationTypeCode>2</fc2:PartyIdentificationTypeCode>
            </fc2:PartyIdentification>
            <fc2:PartyIdentification SeqNum="25">
                <fc2:PartyIdentificationNumberText>3368215</fc2:PartyIdentificationNumberText>
                <fc2:PartyIdentificationTypeCode>14</fc2:PartyIdentificationTypeCode>
            </fc2:PartyIdentification>
        </fc2:Party>
        <fc2:Party SeqNum="26">
            <fc2:ActivityPartyTypeCode>23</fc2:ActivityPartyTypeCode>
            <fc2:IndividualEntityCashInAmountText>67780</fc2:IndividualEntityCashInAmountText>
            <fc2:PartyAsEntityOrganizationIndicator>Y</fc2:PartyAsEntityOrganizationIndicator>
            <fc2:PartyName SeqNum="27">
                <fc2:PartyNameTypeCode>L</fc2:PartyNameTypeCode>
                <fc2:RawEntityIndividualLastName>ALPINE HERBAL WELLNESS LLC</fc2:RawEntityIndividualLastName>
            </fc2:PartyName>
            <fc2:PartyName SeqNum="28">
                <fc2:PartyNameTypeCode>DBA</fc2:PartyNameTypeCode>
                <fc2:RawEntityIndividualLastName>FROSTED LEAF</fc2:RawEntityIndividualLastName>
            </fc2:PartyName>
            <fc2:Address SeqNum="29">
                <fc2:RawCityText>DENVER</fc2:RawCityText>
                <fc2:RawCountryCodeText>US</fc2:RawCountryCodeText>
                <fc2:RawStateCodeText>CO</fc2:RawStateCodeText>
                <fc2:RawStreetAddress1Text>445 FEDERAL BLVD</fc2:RawStreetAddress1Text>
                <fc2:RawZIPCode>802044744</fc2:RawZIPCode>
            </fc2:Address>
            <fc2:PhoneNumber SeqNum="30">
                <fc2:PhoneNumberText>3039998805</fc2:PhoneNumberText>
            </fc2:PhoneNumber>
            <fc2:PartyIdentification SeqNum="31">
                <fc2:PartyIdentificationNumberText>271406262</fc2:PartyIdentificationNumberText>
                <fc2:PartyIdentificationTypeCode>2</fc2:PartyIdentificationTypeCode>
            </fc2:PartyIdentification>
            <fc2:PartyIdentification SeqNum="32">
                <fc2:OtherIssuerCountryText>US</fc2:OtherIssuerCountryText>
                <fc2:OtherIssuerStateText>CO</fc2:OtherIssuerStateText>
                <fc2:OtherPartyIdentificationTypeText>ARTICLES OF ORGANIZATION</fc2:OtherPartyIdentificationTypeText>
            </fc2:PartyIdentification>
            <fc2:PartyOccupationBusiness SeqNum="33">
                <fc2:NAICSCode>4242</fc2:NAICSCode>
                <fc2:OccupationBusinessText>MARIJUANA DISPENSARY</fc2:OccupationBusinessText>
            </fc2:PartyOccupationBusiness>
            <fc2:ElectronicAddress SeqNum="34">
                <fc2:ElectronicAddressText>ACCOUNTING@FROSTEDLEAF.COM</fc2:ElectronicAddressText>
            </fc2:ElectronicAddress>
            <fc2:Account SeqNum="35">
                <fc2:AccountNumberText>54526</fc2:AccountNumberText>
                <fc2:PartyAccountAssociation SeqNum="36">
                    <fc2:PartyAccountAssociationTypeCode>8</fc2:PartyAccountAssociationTypeCode>
                </fc2:PartyAccountAssociation>
            </fc2:Account>
        </fc2:Party>
        <fc2:CurrencyTransactionActivity SeqNum="37">
            <fc2:AggregateTransactionIndicator>Y</fc2:AggregateTransactionIndicator>
            <fc2:ArmoredCarServiceIndicator></fc2:ArmoredCarServiceIndicator>
            <fc2:ATMIndicator></fc2:ATMIndicator>
            <fc2:MailDepositShipmentIndicator></fc2:MailDepositShipmentIndicator>
            <fc2:NightDepositIndicator></fc2:NightDepositIndicator>
            <fc2:SharedBranchingIndicator></fc2:SharedBranchingIndicator>
            <fc2:TotalCashInReceiveAmountText>67780</fc2:TotalCashInReceiveAmountText>
            <fc2:TotalCashOutAmountText></fc2:TotalCashOutAmountText>
            <fc2:TransactionDateText>20191016</fc2:TransactionDateText>
            <fc2:CurrencyTransactionActivityDetail SeqNum="38">
                <fc2:CurrencyTransactionActivityDetailTypeCode>55</fc2:CurrencyTransactionActivityDetailTypeCode>
                <fc2:DetailTransactionAmountText>50000</fc2:DetailTransactionAmountText>
                <fc2:OtherCurrencyTransactionActivityDetailText></fc2:OtherCurrencyTransactionActivityDetailText>
                <fc2:OtherForeignCurrencyCountryText></fc2:OtherForeignCurrencyCountryText>
            </fc2:CurrencyTransactionActivityDetail>
            <fc2:CurrencyTransactionActivityDetail SeqNum="39">
                <fc2:CurrencyTransactionActivityDetailTypeCode>997</fc2:CurrencyTransactionActivityDetailTypeCode>
                <fc2:DetailTransactionAmountText>10000</fc2:DetailTransactionAmountText>
                <fc2:OtherCurrencyTransactionActivityDetailText>Some other type</fc2:OtherCurrencyTransactionActivityDetailText>
                <fc2:OtherForeignCurrencyCountryText></fc2:OtherForeignCurrencyCountryText>
            </fc2:CurrencyTransactionActivityDetail>
            <fc2:CurrencyTransactionActivityDetail SeqNum="40">
                <fc2:CurrencyTransactionActivityDetailTypeCode>53</fc2:CurrencyTransactionActivityDetailTypeCode>
                <fc2:DetailTransactionAmountText>3000</fc2:DetailTransactionAmountText>
                <fc2:OtherCurrencyTransactionActivityDetailText></fc2:OtherCurrencyTransactionActivityDetailText>
                <fc2:OtherForeignCurrencyCountryText>CA</fc2:OtherForeignCurrencyCountryText>
            </fc2:CurrencyTransactionActivityDetail>
            <fc2:CurrencyTransactionActivityDetail SeqNum="41">
                <fc2:CurrencyTransactionActivityDetailTypeCode>53</fc2:CurrencyTransactionActivityDetailTypeCode>
                <fc2:DetailTransactionAmountText>4780</fc2:DetailTransactionAmountText>
                <fc2:OtherCurrencyTransactionActivityDetailText></fc2:OtherCurrencyTransactionActivityDetailText>
                <fc2:OtherForeignCurrencyCountryText>MX</fc2:OtherForeignCurrencyCountryText>
            </fc2:CurrencyTransactionActivityDetail>
        </fc2:CurrencyTransactionActivity>
    </fc2:Activity>
</fc2:EFilingBatchXML>`
