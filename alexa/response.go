package alexa

// NewSimpleResponse - basic simple response sent back to Alexa
func NewSimpleResponse(title string, text string) Response {
	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "PlainText",
				Text: text,
			},
			Card: &Payload{
				Type:    "Simple",
				Title:   title,
				Content: text,
			},
			ShouldEndSession: true,
		},
	}
	return r
}

// Response - the response sent back to Alexa
type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Body              ResBody                `json:"response"`
}

// ResBody - the response body sent back to Alexa
type ResBody struct {
	OutputSpeech     *Payload `json:"outputSpeech,omitempty"`
	Card             *Payload `json:"card,omitempty"`
	ShouldEndSession bool     `json:"shouldEndSession"`
}

// Payload - contains the output speech sent back to Alexa
type Payload struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Text    string `json:"text,omitempty"`
	SSML    string `json:"ssml,omitempty"`
	Content string `json:"content,omitempty"`
}

// NewSSMLResponse - wrapper response containing the Speech Synthesis Markup Language
func NewSSMLResponse(title string, text string) Response {
	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "SSML",
				SSML: text,
			},
			ShouldEndSession: true,
		},
	}
	return r
}

// SSML - the Speech Synthesis Markup Language
type SSML struct {
	text  string
	pause string
}

// SSMLBuilder - simple builder
type SSMLBuilder struct {
	SSML []SSML
}

// Say - help function using builder to construct a SSML text object
func (builder *SSMLBuilder) Say(text string) {
	builder.SSML = append(builder.SSML, SSML{text: text})
}

// Pause - help function using builder to construct a SSML pause object
func (builder *SSMLBuilder) Pause(pause string) {
	builder.SSML = append(builder.SSML, SSML{pause: pause})
}

// Build - creates string from the SSML array
func (builder *SSMLBuilder) Build() string {
	var response string
	for index, ssml := range builder.SSML {
		if ssml.text != "" {
			response += ssml.text + " "
		} else if ssml.pause != "" && index != len(builder.SSML)-1 {
			response += "<break time='" + ssml.pause + "ms'/> "
		}
	}
	return "<speak>" + response + "</speak>"
}
