package main

import (
	"fmt"
	"strconv"

	"github.com/mcktr/check_fritz/modules/fritz"
	"github.com/mcktr/check_fritz/modules/perfdata"
	"github.com/mcktr/check_fritz/modules/thresholds"
)

// CheckDownstreamMax checks the maximum downstream that is available on this internet connection
func CheckTotals(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	var soapReq fritz.SoapData

  // Get total bytes sent
  soapReq = fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/igdupnp/control/WANCommonIFC1",
              "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1", "GetAddonInfos")

	go fritz.DoSoapRequest(&soapReq, resps, errs, aI.Debug)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1, *aI.Timeout)

	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

  soapResp := fritz.WANCommonInterfaceAddonInfoResponse{}
  err = fritz.UnmarshalSoapResponse(&soapResp, res)

  if err != nil {
    panic(err)
  }

  total_bytes_received, err := strconv.ParseFloat(soapResp.NewX_AVM_DE_TotalBytesReceived64, 64)

  if err != nil {
    panic(err)
  }

  total_bytes_sent, err := strconv.ParseFloat(soapResp.NewX_AVM_DE_TotalBytesSent64, 64)

  if err != nil {
    panic(err)
  }

	received_perfData := perfdata.CreatePerformanceData("total_received", ByteCountInGB(int64(total_bytes_received)), "GB")
	sent_perfData := perfdata.CreatePerformanceData("total_sent", ByteCountInGB(int64(total_bytes_sent)), "GB")

	GlobalReturnCode = exitOk

	if thresholds.IsSet(aI.Warning) {
		received_perfData.SetWarning(*aI.Warning)

		if thresholds.CheckLower(*aI.Warning, total_bytes_received) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		received_perfData.SetCritical(*aI.Critical)

		if thresholds.CheckLower(*aI.Critical, total_bytes_received) {
			GlobalReturnCode = exitCritical
		}
	}

	output := " - Total Received / Sent: " + ByteCountSI(int64(total_bytes_received)) + " / " +
	  ByteCountSI(int64(total_bytes_sent)) + " " + received_perfData.GetPerformanceDataAsString() + " " +
	  sent_perfData.GetPerformanceDataAsStringWithoutSeparator()

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK" + output + "\n")
	case exitWarning:
		fmt.Print("WARNING" + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL" + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate totals\n")
	}
}

func ByteCountSI(b int64) string {
    const unit = 1000
    if b < unit {
        return fmt.Sprintf("%d B", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %cB",
        float64(b)/float64(div), "kMGTPE"[exp])
}

func ByteCountInGB(b int64) float64 {
    return float64(b) / (1000 * 1000 * 1000);
}
