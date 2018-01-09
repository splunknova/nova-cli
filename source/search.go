package source

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
)

type NovaSearch struct {
	Auth    string
	NovaURL string
	ErrChan chan error
}

type NovaResults struct {
	NovaEvents []struct {
		Time string `json:"time"`
		Raw  string `json:"event.raw"`
	} `json:"events"`
}

type NovaResultsStats struct {
	NovaEvents []map[string]string `json:"events"`
}

// NewNovaSearch creates a new search obj
func NewNovaSearch(novaURL, auth string) *NovaSearch {
	return &NovaSearch{
		Auth:    auth,
		NovaURL: novaURL,
		ErrChan: make(chan error, 5),
	}
}

// WaitAndLogErrors blocks on the pipeline to complete and logs all errors
func (n *NovaSearch) WaitAndLogErrors() (errorsEncountered bool) {
	for e := range n.ErrChan {
		errorsEncountered = true
		log.Error(e)
	}
	return
}


func (n *NovaSearch) Search(keywords, transforms, report string) (StrMatrix) {
	defer close(n.ErrChan)

	log.Debugf("Searching keywords='%+v'", keywords)
	log.Debugf("Searching transforms='%+v'", transforms)
	log.Debugf("Searching report='%+v'", report)

	keywords = fmt.Sprintf("source=%s* %s", novaCLISourcePrefix, keywords)

	params := map[string]string{
		"keywords":   keywords,
		"transforms": transforms,
		"report":     report,
		"count":      defaultSearchResultsCount,
	}

	data := StrMatrix{}

	results, err := Get(n.NovaURL+eventsURLPath, params, n.Auth)
	if err != nil {
		log.Error(err)
		return data
	}
	log.Debugf("Raw Results: %+v\n\n", string(results))

	if report == "" {
		n1 := NovaResults{}
		json.Unmarshal(results, &n1)
		for _, ne := range n1.NovaEvents {
			data = append(data, []string{ne.Time, ne.Raw})
		}
	} else {
		n1 := NovaResultsStats{}
		json.Unmarshal(results, &n1)
		for k, v := range n1.NovaEvents[0] {
			data = append(data, []string{k, v})
		}
	}
	return data
}

