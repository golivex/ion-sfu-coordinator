// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalizeevents

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Represents user interaction event information sent using the PutEvents API.
type Event struct {
	_ struct{} `type:"structure"`

	// An ID associated with the event. If an event ID is not provided, Amazon Personalize
	// generates a unique ID for the event. An event ID is not used as an input
	// to the model. Amazon Personalize uses the event ID to distinquish unique
	// events. Any subsequent events after the first with the same event ID are
	// not used in model training.
	EventId *string `locationName:"eventId" min:"1" type:"string"`

	// The type of event. This property corresponds to the EVENT_TYPE field of the
	// Interactions schema.
	//
	// EventType is a required field
	EventType *string `locationName:"eventType" min:"1" type:"string" required:"true"`

	// A string map of event-specific data that you might choose to record. For
	// example, if a user rates a movie on your site, you might send the movie ID
	// and rating, and the number of movie ratings made by the user.
	//
	// Each item in the map consists of a key-value pair. For example,
	//
	// {"itemId": "movie1"}
	//
	// {"itemId": "movie2", "eventValue": "4.5"}
	//
	// {"itemId": "movie3", "eventValue": "3", "numberOfRatings": "12"}
	//
	// The keys use camel case names that match the fields in the Interactions schema.
	// The itemId and eventValue keys correspond to the ITEM_ID and EVENT_VALUE
	// fields. In the above example, the eventType might be 'MovieRating' with eventValue
	// being the rating. The numberOfRatings would match the 'NUMBER_OF_RATINGS'
	// field defined in the Interactions schema.
	//
	// Properties is a required field
	Properties aws.JSONValue `locationName:"properties" type:"jsonvalue" required:"true"`

	// The timestamp on the client side when the event occurred.
	//
	// SentAt is a required field
	SentAt *time.Time `locationName:"sentAt" type:"timestamp" required:"true"`
}

// String returns the string representation
func (s Event) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Event) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Event"}
	if s.EventId != nil && len(*s.EventId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventId", 1))
	}

	if s.EventType == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventType"))
	}
	if s.EventType != nil && len(*s.EventType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventType", 1))
	}

	if s.Properties == nil {
		invalidParams.Add(aws.NewErrParamRequired("Properties"))
	}

	if s.SentAt == nil {
		invalidParams.Add(aws.NewErrParamRequired("SentAt"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Event) MarshalFields(e protocol.FieldEncoder) error {
	if s.EventId != nil {
		v := *s.EventId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "eventId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EventType != nil {
		v := *s.EventType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "eventType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Properties != nil {
		v := s.Properties

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "properties", protocol.JSONValue{V: v, EscapeMode: protocol.QuotedEscape}, metadata)
	}
	if s.SentAt != nil {
		v := *s.SentAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sentAt",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	return nil
}
