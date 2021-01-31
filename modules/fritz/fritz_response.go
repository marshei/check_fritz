package fritz

import (
	"encoding/xml"
	"fmt"
	"time"
)

// TR064Response is the overriding data structure for the SOAP responses
type TR064Response interface {
}

// WANConnectionInfoResponse is the data structure for responses from WANPPPConnection
type WANConnectionInfoResponse struct {
	TR064Response
	NewEnable                  string `xml:"Body>GetInfoResponse>NewEnable"`
	NewConnectionStatus        string `xml:"Body>GetInfoResponse>NewConnectionStatus"`
	NewPossibleConnectionTypes string `xml:"Body>GetInfoResponse>NewPossibleConnectionTypes"`
	NewConnectionType          string `xml:"Body>GetInfoResponse>NewConnectionType"`
	NewUptime                  string `xml:"Body>GetInfoResponse>NewUptime"`
	NewUpstreamMaxBitRate      string `xml:"Body>GetInfoResponse>NewUpstreamMaxBitRate"`
	NewDownstreamMaxBitRate    string `xml:"Body>GetInfoResponse>NewDownstreamMaxBitRate"`
	NewLastConnectionError     string `xml:"Body>GetInfoResponse>NewLastConnectionError"`
	NewIdleDisconnectTime      string `xml:"Body>GetInfoResponse>NewIdleDisconnectTime"`
	NewRSIPAvailable           string `xml:"Body>GetInfoResponse>NewRSIPAvailable"`
	NewUserName                string `xml:"Body>GetInfoResponse>NewUserName"`
	NewNATEnabled              string `xml:"Body>GetInfoResponse>NewNATEnabled"`
	NewExternalIPAddress       string `xml:"Body>GetInfoResponse>NewExternalIPAddress"`
	NewDNSServers              string `xml:"Body>GetInfoResponse>NewDNSServers"`
	NewMACAddress              string `xml:"Body>GetInfoResponse>NewMACAddress"`
	NewConnectionTrigger       string `xml:"Body>GetInfoResponse>NewConnectionTrigger"`
	NewLastAuthErrorInfo       string `xml:"Body>GetInfoResponse>NewLastAuthErrorInfo"`
	NewMaxCharsUsername        string `xml:"Body>GetInfoResponse>NewMaxCharsUsername"`
	NewMinCharsUsername        string `xml:"Body>GetInfoResponse>NewMinCharsUsername"`
	NewAllowedCharsUsername    string `xml:"Body>GetInfoResponse>NewAllowedCharsUsername"`
	NewMaxCharsPassword        string `xml:"Body>GetInfoResponse>NewMaxCharsPassword"`
	NewMinCharsPassword        string `xml:"Body>GetInfoResponse>NewMinCharsPassword"`
	NewAllowedCharsPassword    string `xml:"Body>GetInfoResponse>NewAllowedCharsPassword"`
	NewTransportType           string `xml:"Body>GetInfoResponse>NewTransportType"`
	NewRouteProtocolRx         string `xml:"Body>GetInfoResponse>NewRouteProtocolRx"`
	NewPPPoEServiceName        string `xml:"Body>GetInfoResponse>NewPPPoEServiceName"`
	NewRemoteIPAddress         string `xml:"Body>GetInfoResponse>NewRemoteIPAddress"`
	NewPPPoEACName             string `xml:"Body>GetInfoResponse>NewPPPoEACName"`
	NewDNSEnabled              string `xml:"Body>GetInfoResponse>NewDNSEnabled"`
	NewDNSOverrideAllowed      string `xml:"Body>GetInfoResponse>NewDNSOverrideAllowed"`
}

// DeviceInfoResponse is the data structure for responses from DeviceInfo
type DeviceInfoResponse struct {
	TR064Response
	NewManufacturerName string `xml:"Body>GetInfoResponse>NewManufacturerName"`
	NewManufacturerOUI  string `xml:"Body>GetInfoResponse>NewManufacturerOUI"`
	NewModelName        string `xml:"Body>GetInfoResponse>NewModelName"`
	NewDescription      string `xml:"Body>GetInfoResponse>NewDescription"`
	NewProductClass     string `xml:"Body>GetInfoResponse>NewProductClass"`
	NewSerialNumber     string `xml:"Body>GetInfoResponse>NewSerialNumber"`
	NewSoftwareVersion  string `xml:"Body>GetInfoResponse>NewSoftwareVersion"`
	NewHardwareVersion  string `xml:"Body>GetInfoResponse>NewHardwareVersion"`
	NewSpecVersion      string `xml:"Body>GetInfoResponse>NewSpecVersion"`
	NewProvisioningCode string `xml:"Body>GetInfoResponse>NewProvisioningCode"`
	NewUpTime           string `xml:"Body>GetInfoResponse>NewUpTime"`
	NewDeviceLog        string `xml:"Body>GetInfoResponse>NewDeviceLog"`
}

// UserInterfaceInfoResponse is the data structure for responses from UserInterface
type UserInterfaceInfoResponse struct {
	TR064Response
	NewUpgradeAvailable       string `xml:"Body>GetInfoResponse>NewUpgradeAvailable"`
	NewPasswordRequired       string `xml:"Body>GetInfoResponse>NewPasswordRequired"`
	NewPasswordUserSelectable string `xml:"Body>GetInfoResponse>NewPasswordUserSelectable"`
	NewWarrantyDate           string `xml:"Body>GetInfoResponse>NewWarrantyDate"`
	NewXAVMDEVersion          string `xml:"Body>GetInfoResponse>NewX_AVM-DE_Version"`
	NewXAVMDEDownloadURL      string `xml:"Body>GetInfoResponse>NewX_AVM-DE_DownloadURL"`
	NewXAVMDEInfoURL          string `xml:"Body>GetInfoResponse>NewX_AVM-DE_InfoURL"`
	NewXAVMDEUpdateState      string `xml:"Body>GetInfoResponse>NewX_AVM-DE_UpdateState"`
	NewXAVMDELaborVersion     string `xml:"Body>GetInfoResponse>NewX_AVM-DE_LaborVersion"`
}

// WANCommonInterfaceOnlineMonitorResponse is the data structure for responses from WANCommonInterfaceConfig
type WANCommonInterfaceOnlineMonitorResponse struct {
	TR064Response
	NewTotalNumberSyncGroups string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>NewTotalNumberSyncGroups"`
	NewSyncGroupName         string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>NewSyncGroupName"`
	NewSyncGroupMode         string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>NewSyncGroupMode"`
	NewMaxDS                 string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newmax_ds"`
	NewMaxUS                 string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newmax_us"`
	NewDSCurrentBPS          string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newds_current_bps"`
	NewMCCurrentBPS          string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newmc_current_bps"`
	NewUSCurrentBPS          string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newus_current_bps"`
	NewPrioRealtimeBPS       string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_realtime_bps"`
	NewPrioHighBPS           string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_high_bps"`
	NewPrioDefaultBPS        string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_default_bps"`
	NewPrioLowBPS            string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_low_bps"`
}

// SmartDeviceInfoResponse is the data structure for responses from X_AVM-DE_Homeauto
type SmartDeviceInfoResponse struct {
	TR064Response
	NewAIN                    string `xml:"Body>GetGenericDeviceInfosResponse>NewAIN"`
	NewDeviceID               string `xml:"Body>GetGenericDeviceInfosResponse>NewDeviceId"`
	NewFunctionBitMask        string `xml:"Body>GetGenericDeviceInfosResponse>NewFunctionBitMask"`
	NewFirmwareVersion        string `xml:"Body>GetGenericDeviceInfosResponse>NewFirmwareVersion"`
	NewManufacturer           string `xml:"Body>GetGenericDeviceInfosResponse>NewManufacturer"`
	NewProductName            string `xml:"Body>GetGenericDeviceInfosResponse>NewProductName"`
	NewDeviceName             string `xml:"Body>GetGenericDeviceInfosResponse>NewDeviceName"`
	NewPresent                string `xml:"Body>GetGenericDeviceInfosResponse>NewPresent"`
	NewMultimeterIsEnabled    string `xml:"Body>GetGenericDeviceInfosResponse>NewMultimeterIsEnabled"`
	NewMultimeterIsValid      string `xml:"Body>GetGenericDeviceInfosResponse>NewMultimeterIsValid"`
	NewMultimeterPower        string `xml:"Body>GetGenericDeviceInfosResponse>NewMultimeterPower"`
	NewMultimeterEnergy       string `xml:"Body>GetGenericDeviceInfosResponse>NewMultimeterEnergy"`
	NewTemperatureIsEnabled   string `xml:"Body>GetGenericDeviceInfosResponse>NewTemperatureIsEnabled"`
	NewTemperatureIsValid     string `xml:"Body>GetGenericDeviceInfosResponse>NewTemperatureIsValid"`
	NewTemperatureCelsius     string `xml:"Body>GetGenericDeviceInfosResponse>NewTemperatureCelsius"`
	NewTemperatureOffset      string `xml:"Body>GetGenericDeviceInfosResponse>NewTemperatureOffset"`
	NewSwitchIsEnabled        string `xml:"Body>GetGenericDeviceInfosResponse>NewSwitchIsEnabled"`
	NewSwitchIsValid          string `xml:"Body>GetGenericDeviceInfosResponse>NewSwitchIsValid"`
	NewSwitchState            string `xml:"Body>GetGenericDeviceInfosResponse>NewSwitchState"`
	NewSwitchMode             string `xml:"Body>GetGenericDeviceInfosResponse>NewSwitchMode"`
	NewSwitchLock             string `xml:"Body>GetGenericDeviceInfosResponse>NewSwitchLock"`
	NewHkrIsEnabled           string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrIsEnabled"`
	NewHkrIsValid             string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrIsValid"`
	NewHkrIsTemperature       string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrIsTemperature"`
	NewHkrSetVentilStatus     string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrSetVentilStatus"`
	NewHkrSetTemperature      string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrSetTemperature"`
	NewHkrReduceVentilStatus  string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrReduceVentilStatus"`
	NewHkrReduceTemperature   string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrReduceTemperature"`
	NewHkrComfortVentilStatus string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrComfortVentilStatus"`
	NewHkrComfortTemperature  string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrComfortTemperature"`
}

// SmartSpecificDeviceInfoResponse is the data structure for responses from X_AVM-DE_Homeauto
type SmartSpecificDeviceInfoResponse struct {
	TR064Response
	NewAIN                    string `xml:"Body>GetSpecificDeviceInfosResponseResponse>NewAIN"`
	NewDeviceID               string `xml:"Body>GetSpecificDeviceInfosResponseResponse>NewDeviceId"`
	NewFunctionBitMask        string `xml:"Body>GetSpecificDeviceInfosResponseResponse>NewFunctionBitMask"`
	NewFirmwareVersion        string `xml:"Body>GetSpecificDeviceInfosResponseResponse>NewFirmwareVersion"`
	NewManufacturer           string `xml:"Body>GetSpecificDeviceInfosResponse>NewManufacturer"`
	NewProductName            string `xml:"Body>GetSpecificDeviceInfosResponse>NewProductName"`
	NewDeviceName             string `xml:"Body>GetSpecificDeviceInfosResponse>NewDeviceName"`
	NewPresent                string `xml:"Body>GetSpecificDeviceInfosResponse>NewPresent"`
	NewMultimeterIsEnabled    string `xml:"Body>GetSpecificDeviceInfosResponse>NewMultimeterIsEnabled"`
	NewMultimeterIsValid      string `xml:"Body>GetSpecificDeviceInfosResponse>NewMultimeterIsValid"`
	NewMultimeterPower        string `xml:"Body>GetSpecificDeviceInfosResponse>NewMultimeterPower"`
	NewMultimeterEnergy       string `xml:"Body>GetSpecificDeviceInfosResponse>NewMultimeterEnergy"`
	NewTemperatureIsEnabled   string `xml:"Body>GetSpecificDeviceInfosResponse>NewTemperatureIsEnabled"`
	NewTemperatureIsValid     string `xml:"Body>GetSpecificDeviceInfosResponse>NewTemperatureIsValid"`
	NewTemperatureCelsius     string `xml:"Body>GetSpecificDeviceInfosResponse>NewTemperatureCelsius"`
	NewTemperatureOffset      string `xml:"Body>GetSpecificDeviceInfosResponse>NewTemperatureOffset"`
	NewSwitchIsEnabled        string `xml:"Body>GetSpecificDeviceInfosResponse>NewSwitchIsEnabled"`
	NewSwitchIsValid          string `xml:"Body>GetSpecificDeviceInfosResponse>NewSwitchIsValid"`
	NewSwitchState            string `xml:"Body>GetSpecificDeviceInfosResponse>NewSwitchState"`
	NewSwitchMode             string `xml:"Body>GetSpecificDeviceInfosResponse>NewSwitchMode"`
	NewSwitchLock             string `xml:"Body>GetSpecificDeviceInfosResponse>NewSwitchLock"`
	NewHkrIsEnabled           string `xml:"Body>GetSpecificDeviceInfosResponse>NewHkrIsEnabled"`
	NewHkrIsValid             string `xml:"Body>GetSpecificDeviceInfosResponse>NewHkrIsValid"`
	NewHkrIsTemperature       string `xml:"Body>GetSpecificDeviceInfosResponse>NewHkrIsTemperature"`
	NewHkrSetVentilStatus     string `xml:"Body>GetSpecificDeviceInfosResponse>NewHkrSetVentilStatus"`
	NewHkrSetTemperature      string `xml:"Body>GetSpecificDeviceInfosResponse>NewHkrSetTemperature"`
	NewHkrReduceVentilStatus  string `xml:"Body>GetSpecificDeviceInfosResponse>NewHkrReduceVentilStatus"`
	NewHkrReduceTemperature   string `xml:"Body>GetSpecificDeviceInfosResponse>NewHkrReduceTemperature"`
	NewHkrComfortVentilStatus string `xml:"Body>GetSpecificDeviceInfosResponse>NewHkrComfortVentilStatus"`
	NewHkrComfortTemperature  string `xml:"Body>GetSpecificDeviceInfosResponse>NewHkrComfortTemperature"`
}

// WANCommonInterfaceCommonLinkPropertiesResponse is the date structure for responses from GetCommonLinkProperties
type WANCommonInterfaceCommonLinkPropertiesResponse struct {
	TR064Response
	NewWANAccessType              string `xml:"Body>GetCommonLinkPropertiesResponse>GetCommonLinkPropertiesResponse"`
	NewLayer1UpstreamMaxBitRate   string `xml:"Body>GetCommonLinkPropertiesResponse>NewLayer1UpstreamMaxBitRate"`
	NewLayer1DownstreamMaxBitRate string `xml:"Body>GetCommonLinkPropertiesResponse>NewLayer1DownstreamMaxBitRate"`
	NewPhysicalLinkStatus         string `xml:"Body>GetCommonLinkPropertiesResponse>NewPhysicalLinkStatus"`
}

// WANCommonInterfaceAddonInfoResponse is the date structure for responses from GetAddonInfo
type WANCommonInterfaceAddonInfoResponse struct {
	TR064Response
	NewByteSendRate                   string `xml:"Body>GetAddonInfosResponse>NewByteSendRate"`
	NewByteReceiveRate                string `xml:"Body>GetAddonInfosResponse>NewByteReceiveRate"`
	NewPacketSendRate                 string `xml:"Body>GetAddonInfosResponse>NewPacketSendRate"`
	NewPacketReceiveRate              string `xml:"Body>GetAddonInfosResponse>NewPacketReceiveRate"`
	NewTotalBytesSent                 string `xml:"Body>GetAddonInfosResponse>NewTotalBytesSent"`
	NewTotalBytesReceived             string `xml:"Body>GetAddonInfosResponse>NewTotalBytesReceived"`
	NewAutoDisconnectTime             string `xml:"Body>GetAddonInfosResponse>NewAutoDisconnectTime"`
	NewIdleDisconnectTime             string `xml:"Body>GetAddonInfosResponse>NewIdleDisconnectTime"`
	NewDNSServer1                     string `xml:"Body>GetAddonInfosResponse>NewDNSServer1"`
	NewDNSServer2                     string `xml:"Body>GetAddonInfosResponse>NewDNSServer2"`
	NewVoipDNSServer1                 string `xml:"Body>GetAddonInfosResponse>NewVoipDNSServer1"`
	NewVoipDNSServer2                 string `xml:"Body>GetAddonInfosResponse>NewVoipDNSServer2"`
	NewUpnpControlEnabled             string `xml:"Body>GetAddonInfosResponse>NewUpnpControlEnabled"`
	NewRoutedBridgedModeBoth          string `xml:"Body>GetAddonInfosResponse>NewRoutedBridgedModeBoth"`
	NewX_AVM_DE_TotalBytesSent64      string `xml:"Body>GetAddonInfosResponse>NewX_AVM_DE_TotalBytesSent64"`
	NewX_AVM_DE_TotalBytesReceived64  string `xml:"Body>GetAddonInfosResponse>NewX_AVM_DE_TotalBytesReceived64"`
	NewX_AVM_DE_WANAccessType         string `xml:"Body>GetAddonInfosResponse>NewX_AVM_DE_WANAccessType"`
}

// WANDSLInterfaceGetInfoResponse is the date structure for responses from GetInfo
type WANDSLInterfaceGetInfoResponse struct {
	TR064Response
	NewEnable                string `xml:"Body>GetInfoResponse>NewEnable"`
	NewStatus                string `xml:"Body>GetInfoResponse>NewStatus"`
	NewDataPath              string `xml:"Body>GetInfoResponse>NewDataPath"`
	NewUpstreamCurrRate      string `xml:"Body>GetInfoResponse>NewUpstreamCurrRate"`
	NewDownstreamCurrRate    string `xml:"Body>GetInfoResponse>NewDownstreamCurrRate"`
	NewUpstreamMaxRate       string `xml:"Body>GetInfoResponse>NewUpstreamMaxRate"`
	NewDownstreamMaxRate     string `xml:"Body>GetInfoResponse>NewDownstreamMaxRate"`
	NewUpstreamNoiseMargin   string `xml:"Body>GetInfoResponse>NewUpstreamNoiseMargin"`
	NewDownstreamNoiseMargin string `xml:"Body>GetInfoResponse>NewDownstreamNoiseMargin"`
	NewUpstreamAttenuation   string `xml:"Body>GetInfoResponse>NewUpstreamAttenuation"`
	NewDownstreamAttenuation string `xml:"Body>GetInfoResponse>NewDownstreamAttenuation"`
	NewATURVendor            string `xml:"Body>GetInfoResponse>NewATURVendor"`
	NewATURCountry           string `xml:"Body>GetInfoResponse>NewATURCountry"`
	NewUpstreamPower         string `xml:"Body>GetInfoResponse>NewUpstreamPower"`
	NewDownstreamPower       string `xml:"Body>GetInfoResponse>NewDownstreamPower"`
}

// UnmarshalSoapResponse unmarshals the soap response to the data structure
func UnmarshalSoapResponse(resp TR064Response, inputXML [][]byte) error {
	for i := range inputXML {

		err := xml.Unmarshal(inputXML[i], &resp)

		if err != nil {
			return err
		}
	}
	return nil
}

// ProcessSoapResponse handles the SOAP response from channels
func ProcessSoapResponse(resps chan []byte, errs chan error, count int, timeout int) ([][]byte, error) {
	results := make([][]byte, 0)

	for {
		timedout := false

		select {
		case err := <-errs:
			if err != nil {
				return results, err
			}
		case res := <-resps:
			count--
			results = append(results, res)

			if count <= 0 {
				break
			}
		case <-time.After(time.Duration(timeout) * time.Second):
			return nil, fmt.Errorf("Ran into a timeout after %d seconds", timeout)
		}

		if count <= 0 || timedout {
			break
		}
	}

	return results, nil
}
