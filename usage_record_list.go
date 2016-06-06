package twiliogo

import (
	"encoding/json"
	"net/url"
)

type UsageRecordList struct {
	Client          Client
	Start           int           `json:"start"`
	Total           int           `json:"total"`
	NumPages        int           `json:"num_pages"`
	Page            int           `json:"page"`
	PageSize        int           `json:"page_size"`
	End             int           `json:"end"`
	Uri             string        `json:"uri"`
	FirstPageUri    string        `json:"first_page_uri"`
	LastPageUri     string        `json:"last_page_uri"`
	NextPageUri     string        `json:"next_page_uri"`
	PreviousPageUri string        `json"previous_page_uri"`
	UsageRecords    []UsageRecord `json:"usage_records"`
}

func GetUsageRecordList(client Client, interval string, optionals ...Optional) (*UsageRecordList, error) {
	var usageRecordList *UsageRecordList

	params := url.Values{}

	for _, optional := range optionals {
		param, value := optional.GetParam()
		params.Set(param, value)
	}

	body, err := client.get(nil, "/Usage/Records/"+interval+".json")

	if err != nil {
		return nil, err
	}

	usageRecordList = new(UsageRecordList)
	usageRecordList.Client = client
	err = json.Unmarshal(body, usageRecordList)

	return usageRecordList, err
}

func (usageRecordList *UsageRecordList) GetUsageRecords() []UsageRecord {
	return usageRecordList.UsageRecords
}

func (currentUsageRecordList *UsageRecordList) HasNextPage() bool {
	return currentUsageRecordList.NextPageUri != ""
}

func (currentUsageRecordList *UsageRecordList) NextPage() (*UsageRecordList, error) {
	if !currentUsageRecordList.HasNextPage() {
		return nil, Error{"No next page"}
	}

	return currentUsageRecordList.getPage(currentUsageRecordList.NextPageUri)
}

func (currentUsageRecordList *UsageRecordList) HasPreviousPage() bool {
	return currentUsageRecordList.PreviousPageUri != ""
}

func (currentUsageRecordList *UsageRecordList) PreviousPage() (*UsageRecordList, error) {
	if !currentUsageRecordList.HasPreviousPage() {
		return nil, Error{"No previous page"}
	}

	return currentUsageRecordList.getPage(currentUsageRecordList.NextPageUri)
}

func (currentUsageRecordList *UsageRecordList) FirstPage() (*UsageRecordList, error) {
	return currentUsageRecordList.getPage(currentUsageRecordList.FirstPageUri)
}

func (currentUsageRecordList *UsageRecordList) LastPage() (*UsageRecordList, error) {
	return currentUsageRecordList.getPage(currentUsageRecordList.LastPageUri)
}

func (currentUsageRecordList *UsageRecordList) getPage(uri string) (*UsageRecordList, error) {
	var usageRecordList *UsageRecordList

	client := currentUsageRecordList.Client

	body, err := client.get(nil, uri)

	if err != nil {
		return usageRecordList, err
	}

	usageRecordList = new(UsageRecordList)
	usageRecordList.Client = client
	err = json.Unmarshal(body, usageRecordList)

	return usageRecordList, err
}
