package alexa

// built-in intents constants
const (
	HelpIntent   = "AMAZON.HelpIntent"
	CancelIntent = "AMAZON.CancelIntent"
	StopIntent   = "AMAZON.StopIntent"
)

// Request - the incoming Alexa skill request
type Request struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Body    ReqBody `json:"request"`
	Context Context `json:"context"`
}

// Session - the incoming Alexa skill request session
type Session struct {
	New         bool   `json:"new"`
	SessionID   string `json:"sessionId"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`
	Attributes map[string]interface{} `json:"attributes"`
	User       struct {
		UserID      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
	} `json:"user"`
}

// Context - the incoming Alexa skill request context
type Context struct {
	System struct {
		APIAccessToken string `json:"apiAccessToken"`
		Device         struct {
			DeviceID string `json:"deviceId,omitempty"`
		} `json:"device,omitempty"`
		Application struct {
			ApplicationID string `json:"applicationId,omitempty"`
		} `json:"application,omitempty"`
	} `json:"System,omitempty"`
}

// ReqBody - the incoming Alexa skill request request body
type ReqBody struct {
	Type        string `json:"type"`
	RequestID   string `json:"requestId"`
	Timestamp   string `json:"timestamp"`
	Locale      string `json:"locale"`
	Intent      Intent `json:"intent,omitempty"`
	Reason      string `json:"reason,omitempty"`
	DialogState string `json:"dialogState,omitempty"`
}

// Intent - the incoming Alexa skill request intent
type Intent struct {
	Name  string          `json:"name"`
	Slots map[string]Slot `json:"slots"`
}

// Slot - the incoming Alexa skill request slot
type Slot struct {
	Name        string      `json:"name"`
	Value       string      `json:"value"`
	Resolutions Resolutions `json:"resolutions"`
}

// Resolutions - the incoming Alexa skill request slot resolutions
type Resolutions struct {
	ResolutionPerAuthority []struct {
		Values []struct {
			Value struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
}
