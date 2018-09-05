package main

type MyJsonName struct {
	Account struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"account"`
	Comment struct{} `json:"comment"`
	Error   struct {
		App struct {
			BuildUUID            string   `json:"buildUUID"`
			BundleVersion        string   `json:"bundleVersion"`
			CodeBundleID         string   `json:"codeBundleId"`
			DsymUUIDs            []string `json:"dsymUUIDs"`
			Duration             int      `json:"duration"`
			DurationInForeground int      `json:"durationInForeground"`
			ID                   string   `json:"id"`
			InForeground         bool     `json:"inForeground"`
			ReleaseStage         string   `json:"releaseStage"`
			Type                 string   `json:"type"`
			Version              string   `json:"version"`
			VersionCode          string   `json:"versionCode"`
		} `json:"app"`
		AssignedUserID string `json:"assignedUserId"`
		Context        string `json:"context"`
		CreatedIssue   struct {
			ID     string `json:"id"`
			Number int    `json:"number"`
			Type   string `json:"type"`
			URL    string `json:"url"`
		} `json:"createdIssue"`
		Device struct {
			BatteryLevel   float64 `json:"batteryLevel"`
			BrowserName    string  `json:"browserName"`
			BrowserVersion string  `json:"browserVersion"`
			Charging       bool    `json:"charging"`
			FreeDisk       int     `json:"freeDisk"`
			FreeMemory     int     `json:"freeMemory"`
			Hostname       string  `json:"hostname"`
			ID             string  `json:"id"`
			Jailbroken     bool    `json:"jailbroken"`
			Locale         string  `json:"locale"`
			Manufacturer   string  `json:"manufacturer"`
			Model          string  `json:"model"`
			ModelNumber    string  `json:"modelNumber"`
			Orientation    string  `json:"orientation"`
			OsName         string  `json:"osName"`
			OsVersion      string  `json:"osVersion"`
			Time           string  `json:"time"`
			Timezone       string  `json:"timezone"`
			TotalMemory    int     `json:"totalMemory"`
		} `json:"device"`
		ErrorID        string `json:"errorId"`
		ExceptionClass string `json:"exceptionClass"`
		FirstReceived  string `json:"firstReceived"`
		ID             string `json:"id"`
		Message        string `json:"message"`
		MetaData       struct {
			App      struct{} `json:"app"`
			Device   struct{} `json:"device"`
			SomeData struct {
				Key       string `json:"key"`
				SetOfKeys struct {
					Key  string `json:"key"`
					Key2 string `json:"key2"`
				} `json:"setOfKeys"`
			} `json:"someData"`
			SomeMoreData struct{} `json:"someMoreData"`
			User         struct{} `json:"user"`
		} `json:"metaData"`
		ReceivedAt string `json:"receivedAt"`
		RequestURL string `json:"requestUrl"`
		Severity   string `json:"severity"`
		StackTrace []struct {
			Code struct {
				One231 string `json:"1231"`
				One232 string `json:"1232"`
				One233 string `json:"1233"`
				One234 string `json:"1234"`
				One235 string `json:"1235"`
				One236 string `json:"1236"`
				One237 string `json:"1237"`
			} `json:"code"`
			ColumnNumber int    `json:"columnNumber"`
			File         string `json:"file"`
			InProject    bool   `json:"inProject"`
			LineNumber   int    `json:"lineNumber"`
			Method       string `json:"method"`
		} `json:"stackTrace"`
		Status    string `json:"status"`
		Unhandled bool   `json:"unhandled"`
		URL       string `json:"url"`
		User      struct {
			Email string `json:"email"`
			ID    string `json:"id"`
			Name  string `json:"name"`
		} `json:"user"`
	} `json:"error"`
	Project struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"project"`
	Release struct {
		BundleVersion string `json:"bundleVersion"`
		ID            string `json:"id"`
		Metadata      struct {
			MyKey string `json:"myKey"`
		} `json:"metadata"`
		ReleaseStage  string `json:"releaseStage"`
		ReleaseTime   string `json:"releaseTime"`
		ReleasedBy    string `json:"releasedBy"`
		SourceControl struct {
			DiffURL     string `json:"diffUrl"`
			Provider    string `json:"provider"`
			Revision    string `json:"revision"`
			RevisionURL string `json:"revisionUrl"`
		} `json:"sourceControl"`
		URL         string `json:"url"`
		Version     string `json:"version"`
		VersionCode string `json:"versionCode"`
	} `json:"release"`
	Trigger struct {
		Message    string `json:"message"`
		Rate       int    `json:"rate"`
		SnoozeRule struct {
			RuleValue int    `json:"ruleValue"`
			Type      string `json:"type"`
		} `json:"snoozeRule"`
		StateChange string `json:"stateChange"`
		Type        string `json:"type"`
	} `json:"trigger"`
	User struct {
		Email string `json:"email"`
		ID    string `json:"id"`
		Name  string `json:"name"`
	} `json:"user"`
}
