package download

type downloadManagerType struct {
	UserURL string     `json:"user_url"`
	Domain  string     `json:"domain"`
	RootURL string     `json:"root_url"`
	Links   []linkType `json:"links"`
	RootDir string     `json:"root_dir"`
}

type linkType struct {
	Data        string `json:"data"`
	Attr        string `json:"attr"`
	ValOriginal string `json:"val_original"`
	ValNew      string `json:"val_new"`
	GetURL      string `json:"get_url"`
	WrittenOut  bool   `json:"written_out"`
}

func (l linkType) IsEmpty() bool {
	// return l.Data == "" && l.Attr == "" && l.ValOriginal == "" && l.ValNew == "" && l.GetURL == ""

	return l.ValOriginal == ""
}
