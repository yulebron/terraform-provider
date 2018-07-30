package alicloud

import (
		"github.com/aliyun/alibaba-cloud-sdk-go/services/cbn"
		"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func (client *AliyunClient) DescribeCen(cenId string) (c cbn.Cen, err error) {
	request := cbn.CreateDescribeCensRequest()

	for pageNum := 1; ; pageNum++ {
		request.PageNumber = requests.Integer(pageNum)
		resp, _ := client.cenconn.DescribeCens(request)

		cenList := resp.Cens.Cen
		for cenNum := 0; cenNum <= len(cenList); cenNum++ {
			if cenList[cenNum].CenId == cenId {
				return cenList[cenNum], nil
			}
		}

		if pageNum * 10 >= resp.TotalCount {
			return c, GetNotFoundErrorFromString(GetNotFoundMessage("CEN", cenId))
		}

	}

}
