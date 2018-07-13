package cse

type SearchRet struct {
	Cursor  Cursor   `json:"cursor"`
	Context Context  `json:"context"`
	Results []Result `json:"results"`
}

type Context struct {
	Title        string        `json:"title"`
	TotalResults string        `json:"total_results"`
	Facets       []interface{} `json:"facets"`
}

type Cursor struct {
	CurrentPageIndex     int64  `json:"currentPageIndex"`
	EstimatedResultCount string `json:"estimatedResultCount"`
	MoreResultsURL       string `json:"moreResultsUrl"`
	ResultCount          string `json:"resultCount"`
	SearchResultTime     string `json:"searchResultTime"`
	Pages                []Page `json:"pages"`
}

type Page struct {
	Label int64  `json:"label"`
	Start string `json:"start"`
}

type Result struct {
	GsearchResultClass  string `json:"GsearchResultClass"`
	CacheURL            string `json:"cacheUrl"`
	ClicktrackURL       string `json:"clicktrackUrl"`
	Content             string `json:"content"`
	ContentNoFormatting string `json:"contentNoFormatting"`
	FormattedURL        string `json:"formattedUrl"`
	Title               string `json:"title"`
	TitleNoFormatting   string `json:"titleNoFormatting"`
	UnescapedURL        string `json:"unescapedUrl"`
	URL                 string `json:"url"`
	VisibleURL          string `json:"visibleUrl"`
}
