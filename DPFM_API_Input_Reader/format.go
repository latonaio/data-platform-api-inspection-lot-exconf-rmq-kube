package dpfm_api_input_reader

import (
	"data-platform-api-inspection-lot-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *HeaderSDC) ConvertToInspectionLotHeader() *requests.InspectionLotHeader {
	data := sdc.InspectionLotHeader
	return &requests.InspectionLotHeader{
		InspectionLot: data.InspectionLot,
	}
}
