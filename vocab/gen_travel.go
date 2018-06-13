//
package vocab

import (
	"fmt"
	"net/url"
	"time"
)

// TravelType is an interface for accepting types that extend from 'Travel'.
type TravelType interface {
	Serializer
	Deserializer
	ActorLen() (l int)
	IsActorObject(index int) (ok bool)
	GetActorObject(index int) (v ObjectType)
	AppendActorObject(v ObjectType)
	PrependActorObject(v ObjectType)
	RemoveActorObject(index int)
	IsActorLink(index int) (ok bool)
	GetActorLink(index int) (v LinkType)
	AppendActorLink(v LinkType)
	PrependActorLink(v LinkType)
	RemoveActorLink(index int)
	IsActorIRI(index int) (ok bool)
	GetActorIRI(index int) (v *url.URL)
	AppendActorIRI(v *url.URL)
	PrependActorIRI(v *url.URL)
	RemoveActorIRI(index int)
	TargetLen() (l int)
	IsTargetObject(index int) (ok bool)
	GetTargetObject(index int) (v ObjectType)
	AppendTargetObject(v ObjectType)
	PrependTargetObject(v ObjectType)
	RemoveTargetObject(index int)
	IsTargetLink(index int) (ok bool)
	GetTargetLink(index int) (v LinkType)
	AppendTargetLink(v LinkType)
	PrependTargetLink(v LinkType)
	RemoveTargetLink(index int)
	IsTargetIRI(index int) (ok bool)
	GetTargetIRI(index int) (v *url.URL)
	AppendTargetIRI(v *url.URL)
	PrependTargetIRI(v *url.URL)
	RemoveTargetIRI(index int)
	ResultLen() (l int)
	IsResultObject(index int) (ok bool)
	GetResultObject(index int) (v ObjectType)
	AppendResultObject(v ObjectType)
	PrependResultObject(v ObjectType)
	RemoveResultObject(index int)
	IsResultLink(index int) (ok bool)
	GetResultLink(index int) (v LinkType)
	AppendResultLink(v LinkType)
	PrependResultLink(v LinkType)
	RemoveResultLink(index int)
	IsResultIRI(index int) (ok bool)
	GetResultIRI(index int) (v *url.URL)
	AppendResultIRI(v *url.URL)
	PrependResultIRI(v *url.URL)
	RemoveResultIRI(index int)
	OriginLen() (l int)
	IsOriginObject(index int) (ok bool)
	GetOriginObject(index int) (v ObjectType)
	AppendOriginObject(v ObjectType)
	PrependOriginObject(v ObjectType)
	RemoveOriginObject(index int)
	IsOriginLink(index int) (ok bool)
	GetOriginLink(index int) (v LinkType)
	AppendOriginLink(v LinkType)
	PrependOriginLink(v LinkType)
	RemoveOriginLink(index int)
	IsOriginIRI(index int) (ok bool)
	GetOriginIRI(index int) (v *url.URL)
	AppendOriginIRI(v *url.URL)
	PrependOriginIRI(v *url.URL)
	RemoveOriginIRI(index int)
	InstrumentLen() (l int)
	IsInstrumentObject(index int) (ok bool)
	GetInstrumentObject(index int) (v ObjectType)
	AppendInstrumentObject(v ObjectType)
	PrependInstrumentObject(v ObjectType)
	RemoveInstrumentObject(index int)
	IsInstrumentLink(index int) (ok bool)
	GetInstrumentLink(index int) (v LinkType)
	AppendInstrumentLink(v LinkType)
	PrependInstrumentLink(v LinkType)
	RemoveInstrumentLink(index int)
	IsInstrumentIRI(index int) (ok bool)
	GetInstrumentIRI(index int) (v *url.URL)
	AppendInstrumentIRI(v *url.URL)
	PrependInstrumentIRI(v *url.URL)
	RemoveInstrumentIRI(index int)
	IsAltitude() (ok bool)
	GetAltitude() (v float64)
	SetAltitude(v float64)
	IsAltitudeIRI() (ok bool)
	GetAltitudeIRI() (v *url.URL)
	SetAltitudeIRI(v *url.URL)
	AttachmentLen() (l int)
	IsAttachmentObject(index int) (ok bool)
	GetAttachmentObject(index int) (v ObjectType)
	AppendAttachmentObject(v ObjectType)
	PrependAttachmentObject(v ObjectType)
	RemoveAttachmentObject(index int)
	IsAttachmentLink(index int) (ok bool)
	GetAttachmentLink(index int) (v LinkType)
	AppendAttachmentLink(v LinkType)
	PrependAttachmentLink(v LinkType)
	RemoveAttachmentLink(index int)
	IsAttachmentIRI(index int) (ok bool)
	GetAttachmentIRI(index int) (v *url.URL)
	AppendAttachmentIRI(v *url.URL)
	PrependAttachmentIRI(v *url.URL)
	RemoveAttachmentIRI(index int)
	AttributedToLen() (l int)
	IsAttributedToObject(index int) (ok bool)
	GetAttributedToObject(index int) (v ObjectType)
	AppendAttributedToObject(v ObjectType)
	PrependAttributedToObject(v ObjectType)
	RemoveAttributedToObject(index int)
	IsAttributedToLink(index int) (ok bool)
	GetAttributedToLink(index int) (v LinkType)
	AppendAttributedToLink(v LinkType)
	PrependAttributedToLink(v LinkType)
	RemoveAttributedToLink(index int)
	IsAttributedToIRI(index int) (ok bool)
	GetAttributedToIRI(index int) (v *url.URL)
	AppendAttributedToIRI(v *url.URL)
	PrependAttributedToIRI(v *url.URL)
	RemoveAttributedToIRI(index int)
	AudienceLen() (l int)
	IsAudienceObject(index int) (ok bool)
	GetAudienceObject(index int) (v ObjectType)
	AppendAudienceObject(v ObjectType)
	PrependAudienceObject(v ObjectType)
	RemoveAudienceObject(index int)
	IsAudienceLink(index int) (ok bool)
	GetAudienceLink(index int) (v LinkType)
	AppendAudienceLink(v LinkType)
	PrependAudienceLink(v LinkType)
	RemoveAudienceLink(index int)
	IsAudienceIRI(index int) (ok bool)
	GetAudienceIRI(index int) (v *url.URL)
	AppendAudienceIRI(v *url.URL)
	PrependAudienceIRI(v *url.URL)
	RemoveAudienceIRI(index int)
	ContentLen() (l int)
	IsContentString(index int) (ok bool)
	GetContentString(index int) (v string)
	AppendContentString(v string)
	PrependContentString(v string)
	RemoveContentString(index int)
	IsContentLangString(index int) (ok bool)
	GetContentLangString(index int) (v string)
	AppendContentLangString(v string)
	PrependContentLangString(v string)
	RemoveContentLangString(index int)
	IsContentIRI(index int) (ok bool)
	GetContentIRI(index int) (v *url.URL)
	AppendContentIRI(v *url.URL)
	PrependContentIRI(v *url.URL)
	RemoveContentIRI(index int)
	ContentMapLanguages() (l []string)
	GetContentMap(l string) (v string)
	SetContentMap(l string, v string)
	ContextLen() (l int)
	IsContextObject(index int) (ok bool)
	GetContextObject(index int) (v ObjectType)
	AppendContextObject(v ObjectType)
	PrependContextObject(v ObjectType)
	RemoveContextObject(index int)
	IsContextLink(index int) (ok bool)
	GetContextLink(index int) (v LinkType)
	AppendContextLink(v LinkType)
	PrependContextLink(v LinkType)
	RemoveContextLink(index int)
	IsContextIRI(index int) (ok bool)
	GetContextIRI(index int) (v *url.URL)
	AppendContextIRI(v *url.URL)
	PrependContextIRI(v *url.URL)
	RemoveContextIRI(index int)
	NameLen() (l int)
	IsNameString(index int) (ok bool)
	GetNameString(index int) (v string)
	AppendNameString(v string)
	PrependNameString(v string)
	RemoveNameString(index int)
	IsNameLangString(index int) (ok bool)
	GetNameLangString(index int) (v string)
	AppendNameLangString(v string)
	PrependNameLangString(v string)
	RemoveNameLangString(index int)
	IsNameIRI(index int) (ok bool)
	GetNameIRI(index int) (v *url.URL)
	AppendNameIRI(v *url.URL)
	PrependNameIRI(v *url.URL)
	RemoveNameIRI(index int)
	NameMapLanguages() (l []string)
	GetNameMap(l string) (v string)
	SetNameMap(l string, v string)
	IsEndTime() (ok bool)
	GetEndTime() (v time.Time)
	SetEndTime(v time.Time)
	IsEndTimeIRI() (ok bool)
	GetEndTimeIRI() (v *url.URL)
	SetEndTimeIRI(v *url.URL)
	GeneratorLen() (l int)
	IsGeneratorObject(index int) (ok bool)
	GetGeneratorObject(index int) (v ObjectType)
	AppendGeneratorObject(v ObjectType)
	PrependGeneratorObject(v ObjectType)
	RemoveGeneratorObject(index int)
	IsGeneratorLink(index int) (ok bool)
	GetGeneratorLink(index int) (v LinkType)
	AppendGeneratorLink(v LinkType)
	PrependGeneratorLink(v LinkType)
	RemoveGeneratorLink(index int)
	IsGeneratorIRI(index int) (ok bool)
	GetGeneratorIRI(index int) (v *url.URL)
	AppendGeneratorIRI(v *url.URL)
	PrependGeneratorIRI(v *url.URL)
	RemoveGeneratorIRI(index int)
	IconLen() (l int)
	IsIconImage(index int) (ok bool)
	GetIconImage(index int) (v ImageType)
	AppendIconImage(v ImageType)
	PrependIconImage(v ImageType)
	RemoveIconImage(index int)
	IsIconLink(index int) (ok bool)
	GetIconLink(index int) (v LinkType)
	AppendIconLink(v LinkType)
	PrependIconLink(v LinkType)
	RemoveIconLink(index int)
	IsIconIRI(index int) (ok bool)
	GetIconIRI(index int) (v *url.URL)
	AppendIconIRI(v *url.URL)
	PrependIconIRI(v *url.URL)
	RemoveIconIRI(index int)
	HasId() (ok bool)
	GetId() (v *url.URL)
	SetId(v *url.URL)
	HasUnknownId() (ok bool)
	GetUnknownId() (v interface{})
	SetUnknownId(i interface{})
	ImageLen() (l int)
	IsImageImage(index int) (ok bool)
	GetImageImage(index int) (v ImageType)
	AppendImageImage(v ImageType)
	PrependImageImage(v ImageType)
	RemoveImageImage(index int)
	IsImageLink(index int) (ok bool)
	GetImageLink(index int) (v LinkType)
	AppendImageLink(v LinkType)
	PrependImageLink(v LinkType)
	RemoveImageLink(index int)
	IsImageIRI(index int) (ok bool)
	GetImageIRI(index int) (v *url.URL)
	AppendImageIRI(v *url.URL)
	PrependImageIRI(v *url.URL)
	RemoveImageIRI(index int)
	InReplyToLen() (l int)
	IsInReplyToObject(index int) (ok bool)
	GetInReplyToObject(index int) (v ObjectType)
	AppendInReplyToObject(v ObjectType)
	PrependInReplyToObject(v ObjectType)
	RemoveInReplyToObject(index int)
	IsInReplyToLink(index int) (ok bool)
	GetInReplyToLink(index int) (v LinkType)
	AppendInReplyToLink(v LinkType)
	PrependInReplyToLink(v LinkType)
	RemoveInReplyToLink(index int)
	IsInReplyToIRI(index int) (ok bool)
	GetInReplyToIRI(index int) (v *url.URL)
	AppendInReplyToIRI(v *url.URL)
	PrependInReplyToIRI(v *url.URL)
	RemoveInReplyToIRI(index int)
	LocationLen() (l int)
	IsLocationObject(index int) (ok bool)
	GetLocationObject(index int) (v ObjectType)
	AppendLocationObject(v ObjectType)
	PrependLocationObject(v ObjectType)
	RemoveLocationObject(index int)
	IsLocationLink(index int) (ok bool)
	GetLocationLink(index int) (v LinkType)
	AppendLocationLink(v LinkType)
	PrependLocationLink(v LinkType)
	RemoveLocationLink(index int)
	IsLocationIRI(index int) (ok bool)
	GetLocationIRI(index int) (v *url.URL)
	AppendLocationIRI(v *url.URL)
	PrependLocationIRI(v *url.URL)
	RemoveLocationIRI(index int)
	PreviewLen() (l int)
	IsPreviewObject(index int) (ok bool)
	GetPreviewObject(index int) (v ObjectType)
	AppendPreviewObject(v ObjectType)
	PrependPreviewObject(v ObjectType)
	RemovePreviewObject(index int)
	IsPreviewLink(index int) (ok bool)
	GetPreviewLink(index int) (v LinkType)
	AppendPreviewLink(v LinkType)
	PrependPreviewLink(v LinkType)
	RemovePreviewLink(index int)
	IsPreviewIRI(index int) (ok bool)
	GetPreviewIRI(index int) (v *url.URL)
	AppendPreviewIRI(v *url.URL)
	PrependPreviewIRI(v *url.URL)
	RemovePreviewIRI(index int)
	IsPublished() (ok bool)
	GetPublished() (v time.Time)
	SetPublished(v time.Time)
	IsPublishedIRI() (ok bool)
	GetPublishedIRI() (v *url.URL)
	SetPublishedIRI(v *url.URL)
	IsReplies() (ok bool)
	GetReplies() (v CollectionType)
	SetReplies(v CollectionType)
	IsRepliesIRI() (ok bool)
	GetRepliesIRI() (v *url.URL)
	SetRepliesIRI(v *url.URL)
	IsStartTime() (ok bool)
	GetStartTime() (v time.Time)
	SetStartTime(v time.Time)
	IsStartTimeIRI() (ok bool)
	GetStartTimeIRI() (v *url.URL)
	SetStartTimeIRI(v *url.URL)
	SummaryLen() (l int)
	IsSummaryString(index int) (ok bool)
	GetSummaryString(index int) (v string)
	AppendSummaryString(v string)
	PrependSummaryString(v string)
	RemoveSummaryString(index int)
	IsSummaryLangString(index int) (ok bool)
	GetSummaryLangString(index int) (v string)
	AppendSummaryLangString(v string)
	PrependSummaryLangString(v string)
	RemoveSummaryLangString(index int)
	IsSummaryIRI(index int) (ok bool)
	GetSummaryIRI(index int) (v *url.URL)
	AppendSummaryIRI(v *url.URL)
	PrependSummaryIRI(v *url.URL)
	RemoveSummaryIRI(index int)
	SummaryMapLanguages() (l []string)
	GetSummaryMap(l string) (v string)
	SetSummaryMap(l string, v string)
	TagLen() (l int)
	IsTagObject(index int) (ok bool)
	GetTagObject(index int) (v ObjectType)
	AppendTagObject(v ObjectType)
	PrependTagObject(v ObjectType)
	RemoveTagObject(index int)
	IsTagLink(index int) (ok bool)
	GetTagLink(index int) (v LinkType)
	AppendTagLink(v LinkType)
	PrependTagLink(v LinkType)
	RemoveTagLink(index int)
	IsTagIRI(index int) (ok bool)
	GetTagIRI(index int) (v *url.URL)
	AppendTagIRI(v *url.URL)
	PrependTagIRI(v *url.URL)
	RemoveTagIRI(index int)
	TypeLen() (l int)
	GetType(index int) (v interface{})
	AppendType(v interface{})
	PrependType(v interface{})
	RemoveType(index int)
	IsUpdated() (ok bool)
	GetUpdated() (v time.Time)
	SetUpdated(v time.Time)
	IsUpdatedIRI() (ok bool)
	GetUpdatedIRI() (v *url.URL)
	SetUpdatedIRI(v *url.URL)
	UrlLen() (l int)
	IsUrlAnyURI(index int) (ok bool)
	GetUrlAnyURI(index int) (v *url.URL)
	AppendUrlAnyURI(v *url.URL)
	PrependUrlAnyURI(v *url.URL)
	RemoveUrlAnyURI(index int)
	IsUrlLink(index int) (ok bool)
	GetUrlLink(index int) (v LinkType)
	AppendUrlLink(v LinkType)
	PrependUrlLink(v LinkType)
	RemoveUrlLink(index int)
	ToLen() (l int)
	IsToObject(index int) (ok bool)
	GetToObject(index int) (v ObjectType)
	AppendToObject(v ObjectType)
	PrependToObject(v ObjectType)
	RemoveToObject(index int)
	IsToLink(index int) (ok bool)
	GetToLink(index int) (v LinkType)
	AppendToLink(v LinkType)
	PrependToLink(v LinkType)
	RemoveToLink(index int)
	IsToIRI(index int) (ok bool)
	GetToIRI(index int) (v *url.URL)
	AppendToIRI(v *url.URL)
	PrependToIRI(v *url.URL)
	RemoveToIRI(index int)
	BtoLen() (l int)
	IsBtoObject(index int) (ok bool)
	GetBtoObject(index int) (v ObjectType)
	AppendBtoObject(v ObjectType)
	PrependBtoObject(v ObjectType)
	RemoveBtoObject(index int)
	IsBtoLink(index int) (ok bool)
	GetBtoLink(index int) (v LinkType)
	AppendBtoLink(v LinkType)
	PrependBtoLink(v LinkType)
	RemoveBtoLink(index int)
	IsBtoIRI(index int) (ok bool)
	GetBtoIRI(index int) (v *url.URL)
	AppendBtoIRI(v *url.URL)
	PrependBtoIRI(v *url.URL)
	RemoveBtoIRI(index int)
	CcLen() (l int)
	IsCcObject(index int) (ok bool)
	GetCcObject(index int) (v ObjectType)
	AppendCcObject(v ObjectType)
	PrependCcObject(v ObjectType)
	RemoveCcObject(index int)
	IsCcLink(index int) (ok bool)
	GetCcLink(index int) (v LinkType)
	AppendCcLink(v LinkType)
	PrependCcLink(v LinkType)
	RemoveCcLink(index int)
	IsCcIRI(index int) (ok bool)
	GetCcIRI(index int) (v *url.URL)
	AppendCcIRI(v *url.URL)
	PrependCcIRI(v *url.URL)
	RemoveCcIRI(index int)
	BccLen() (l int)
	IsBccObject(index int) (ok bool)
	GetBccObject(index int) (v ObjectType)
	AppendBccObject(v ObjectType)
	PrependBccObject(v ObjectType)
	RemoveBccObject(index int)
	IsBccLink(index int) (ok bool)
	GetBccLink(index int) (v LinkType)
	AppendBccLink(v LinkType)
	PrependBccLink(v LinkType)
	RemoveBccLink(index int)
	IsBccIRI(index int) (ok bool)
	GetBccIRI(index int) (v *url.URL)
	AppendBccIRI(v *url.URL)
	PrependBccIRI(v *url.URL)
	RemoveBccIRI(index int)
	IsMediaType() (ok bool)
	GetMediaType() (v string)
	SetMediaType(v string)
	IsMediaTypeIRI() (ok bool)
	GetMediaTypeIRI() (v *url.URL)
	SetMediaTypeIRI(v *url.URL)
	IsDuration() (ok bool)
	GetDuration() (v time.Duration)
	SetDuration(v time.Duration)
	IsDurationIRI() (ok bool)
	GetDurationIRI() (v *url.URL)
	SetDurationIRI(v *url.URL)
	IsSource() (ok bool)
	GetSource() (v ObjectType)
	SetSource(v ObjectType)
	IsSourceIRI() (ok bool)
	GetSourceIRI() (v *url.URL)
	SetSourceIRI(v *url.URL)
	IsInboxOrderedCollection() (ok bool)
	GetInboxOrderedCollection() (v OrderedCollectionType)
	SetInboxOrderedCollection(v OrderedCollectionType)
	IsInboxAnyURI() (ok bool)
	GetInboxAnyURI() (v *url.URL)
	SetInboxAnyURI(v *url.URL)
	IsOutboxOrderedCollection() (ok bool)
	GetOutboxOrderedCollection() (v OrderedCollectionType)
	SetOutboxOrderedCollection(v OrderedCollectionType)
	IsOutboxAnyURI() (ok bool)
	GetOutboxAnyURI() (v *url.URL)
	SetOutboxAnyURI(v *url.URL)
	IsFollowingCollection() (ok bool)
	GetFollowingCollection() (v CollectionType)
	SetFollowingCollection(v CollectionType)
	IsFollowingOrderedCollection() (ok bool)
	GetFollowingOrderedCollection() (v OrderedCollectionType)
	SetFollowingOrderedCollection(v OrderedCollectionType)
	IsFollowingAnyURI() (ok bool)
	GetFollowingAnyURI() (v *url.URL)
	SetFollowingAnyURI(v *url.URL)
	IsFollowersCollection() (ok bool)
	GetFollowersCollection() (v CollectionType)
	SetFollowersCollection(v CollectionType)
	IsFollowersOrderedCollection() (ok bool)
	GetFollowersOrderedCollection() (v OrderedCollectionType)
	SetFollowersOrderedCollection(v OrderedCollectionType)
	IsFollowersAnyURI() (ok bool)
	GetFollowersAnyURI() (v *url.URL)
	SetFollowersAnyURI(v *url.URL)
	IsLikedCollection() (ok bool)
	GetLikedCollection() (v CollectionType)
	SetLikedCollection(v CollectionType)
	IsLikedOrderedCollection() (ok bool)
	GetLikedOrderedCollection() (v OrderedCollectionType)
	SetLikedOrderedCollection(v OrderedCollectionType)
	IsLikedAnyURI() (ok bool)
	GetLikedAnyURI() (v *url.URL)
	SetLikedAnyURI(v *url.URL)
	IsLikesCollection() (ok bool)
	GetLikesCollection() (v CollectionType)
	SetLikesCollection(v CollectionType)
	IsLikesOrderedCollection() (ok bool)
	GetLikesOrderedCollection() (v OrderedCollectionType)
	SetLikesOrderedCollection(v OrderedCollectionType)
	IsLikesAnyURI() (ok bool)
	GetLikesAnyURI() (v *url.URL)
	SetLikesAnyURI(v *url.URL)
	StreamsLen() (l int)
	GetStreams(index int) (v *url.URL)
	AppendStreams(v *url.URL)
	PrependStreams(v *url.URL)
	RemoveStreams(index int)
	HasUnknownStreams() (ok bool)
	GetUnknownStreams() (v interface{})
	SetUnknownStreams(i interface{})
	IsPreferredUsername() (ok bool)
	GetPreferredUsername() (v string)
	SetPreferredUsername(v string)
	IsPreferredUsernameIRI() (ok bool)
	GetPreferredUsernameIRI() (v *url.URL)
	SetPreferredUsernameIRI(v *url.URL)
	PreferredUsernameMapLanguages() (l []string)
	GetPreferredUsernameMap(l string) (v string)
	SetPreferredUsernameMap(l string, v string)
	IsEndpoints() (ok bool)
	GetEndpoints() (v ObjectType)
	SetEndpoints(v ObjectType)
	IsEndpointsIRI() (ok bool)
	GetEndpointsIRI() (v *url.URL)
	SetEndpointsIRI(v *url.URL)
	HasProxyUrl() (ok bool)
	GetProxyUrl() (v *url.URL)
	SetProxyUrl(v *url.URL)
	HasUnknownProxyUrl() (ok bool)
	GetUnknownProxyUrl() (v interface{})
	SetUnknownProxyUrl(i interface{})
	HasOauthAuthorizationEndpoint() (ok bool)
	GetOauthAuthorizationEndpoint() (v *url.URL)
	SetOauthAuthorizationEndpoint(v *url.URL)
	HasUnknownOauthAuthorizationEndpoint() (ok bool)
	GetUnknownOauthAuthorizationEndpoint() (v interface{})
	SetUnknownOauthAuthorizationEndpoint(i interface{})
	HasOauthTokenEndpoint() (ok bool)
	GetOauthTokenEndpoint() (v *url.URL)
	SetOauthTokenEndpoint(v *url.URL)
	HasUnknownOauthTokenEndpoint() (ok bool)
	GetUnknownOauthTokenEndpoint() (v interface{})
	SetUnknownOauthTokenEndpoint(i interface{})
	HasProvideClientKey() (ok bool)
	GetProvideClientKey() (v *url.URL)
	SetProvideClientKey(v *url.URL)
	HasUnknownProvideClientKey() (ok bool)
	GetUnknownProvideClientKey() (v interface{})
	SetUnknownProvideClientKey(i interface{})
	HasSignClientKey() (ok bool)
	GetSignClientKey() (v *url.URL)
	SetSignClientKey(v *url.URL)
	HasUnknownSignClientKey() (ok bool)
	GetUnknownSignClientKey() (v interface{})
	SetUnknownSignClientKey(i interface{})
	HasSharedInbox() (ok bool)
	GetSharedInbox() (v *url.URL)
	SetSharedInbox(v *url.URL)
	HasUnknownSharedInbox() (ok bool)
	GetUnknownSharedInbox() (v interface{})
	SetUnknownSharedInbox(i interface{})
}

// Indicates that the actor is traveling to target from origin. Travel is an IntransitiveObject whose actor specifies the direct object. If the target or origin are not specified, either can be determined by context.
type Travel struct {
	// An unknown value.
	unknown_ map[string]interface{}
	// The 'actor' value could have multiple types and values
	actor []*actorTravelIntermediateType
	// The 'target' value could have multiple types and values
	target []*targetTravelIntermediateType
	// The 'result' value could have multiple types and values
	result []*resultTravelIntermediateType
	// The 'origin' value could have multiple types and values
	origin []*originTravelIntermediateType
	// The 'instrument' value could have multiple types and values
	instrument []*instrumentTravelIntermediateType
	// The functional 'altitude' value could have multiple types, but only a single value
	altitude *altitudeTravelIntermediateType
	// The 'attachment' value could have multiple types and values
	attachment []*attachmentTravelIntermediateType
	// The 'attributedTo' value could have multiple types and values
	attributedTo []*attributedToTravelIntermediateType
	// The 'audience' value could have multiple types and values
	audience []*audienceTravelIntermediateType
	// The 'content' value could have multiple types and values
	content []*contentTravelIntermediateType
	// The 'contentMap' value holds language-specific values for property 'content'
	contentMap map[string]string
	// The 'context' value could have multiple types and values
	context []*contextTravelIntermediateType
	// The 'name' value could have multiple types and values
	name []*nameTravelIntermediateType
	// The 'nameMap' value holds language-specific values for property 'name'
	nameMap map[string]string
	// The functional 'endTime' value could have multiple types, but only a single value
	endTime *endTimeTravelIntermediateType
	// The 'generator' value could have multiple types and values
	generator []*generatorTravelIntermediateType
	// The 'icon' value could have multiple types and values
	icon []*iconTravelIntermediateType
	// The functional 'id' value holds a single type and a single value
	id *url.URL
	// The 'image' value could have multiple types and values
	image []*imageTravelIntermediateType
	// The 'inReplyTo' value could have multiple types and values
	inReplyTo []*inReplyToTravelIntermediateType
	// The 'location' value could have multiple types and values
	location []*locationTravelIntermediateType
	// The 'preview' value could have multiple types and values
	preview []*previewTravelIntermediateType
	// The functional 'published' value could have multiple types, but only a single value
	published *publishedTravelIntermediateType
	// The functional 'replies' value could have multiple types, but only a single value
	replies *repliesTravelIntermediateType
	// The functional 'startTime' value could have multiple types, but only a single value
	startTime *startTimeTravelIntermediateType
	// The 'summary' value could have multiple types and values
	summary []*summaryTravelIntermediateType
	// The 'summaryMap' value holds language-specific values for property 'summary'
	summaryMap map[string]string
	// The 'tag' value could have multiple types and values
	tag []*tagTravelIntermediateType
	// The 'type' value can hold any type and any number of values
	typeName []interface{}
	// The functional 'updated' value could have multiple types, but only a single value
	updated *updatedTravelIntermediateType
	// The 'url' value could have multiple types and values
	url []*urlTravelIntermediateType
	// The 'to' value could have multiple types and values
	to []*toTravelIntermediateType
	// The 'bto' value could have multiple types and values
	bto []*btoTravelIntermediateType
	// The 'cc' value could have multiple types and values
	cc []*ccTravelIntermediateType
	// The 'bcc' value could have multiple types and values
	bcc []*bccTravelIntermediateType
	// The functional 'mediaType' value could have multiple types, but only a single value
	mediaType *mediaTypeTravelIntermediateType
	// The functional 'duration' value could have multiple types, but only a single value
	duration *durationTravelIntermediateType
	// The functional 'source' value could have multiple types, but only a single value
	source *sourceTravelIntermediateType
	// The functional 'inbox' value could have multiple types, but only a single value
	inbox *inboxTravelIntermediateType
	// The functional 'outbox' value could have multiple types, but only a single value
	outbox *outboxTravelIntermediateType
	// The functional 'following' value could have multiple types, but only a single value
	following *followingTravelIntermediateType
	// The functional 'followers' value could have multiple types, but only a single value
	followers *followersTravelIntermediateType
	// The functional 'liked' value could have multiple types, but only a single value
	liked *likedTravelIntermediateType
	// The functional 'likes' value could have multiple types, but only a single value
	likes *likesTravelIntermediateType
	// The 'streams' value holds a single type and any number of values
	streams []*url.URL
	// The functional 'preferredUsername' value could have multiple types, but only a single value
	preferredUsername *preferredUsernameTravelIntermediateType
	// The 'preferredUsernameMap' value holds language-specific values for property 'preferredUsername'
	preferredUsernameMap map[string]string
	// The functional 'endpoints' value could have multiple types, but only a single value
	endpoints *endpointsTravelIntermediateType
	// The functional 'proxyUrl' value holds a single type and a single value
	proxyUrl *url.URL
	// The functional 'oauthAuthorizationEndpoint' value holds a single type and a single value
	oauthAuthorizationEndpoint *url.URL
	// The functional 'oauthTokenEndpoint' value holds a single type and a single value
	oauthTokenEndpoint *url.URL
	// The functional 'provideClientKey' value holds a single type and a single value
	provideClientKey *url.URL
	// The functional 'signClientKey' value holds a single type and a single value
	signClientKey *url.URL
	// The functional 'sharedInbox' value holds a single type and a single value
	sharedInbox *url.URL
}

// ActorLen determines the number of elements able to be used for the IsActorObject, GetActorObject, and RemoveActorObject functions
func (t *Travel) ActorLen() (l int) {
	return len(t.actor)

}

// IsActorObject determines whether the call to GetActorObject is safe for the specified index
func (t *Travel) IsActorObject(index int) (ok bool) {
	return t.actor[index].Object != nil

}

// GetActorObject returns the value safely if IsActorObject returned true for the specified index
func (t *Travel) GetActorObject(index int) (v ObjectType) {
	return t.actor[index].Object

}

// AppendActorObject adds to the back of actor a ObjectType type
func (t *Travel) AppendActorObject(v ObjectType) {
	t.actor = append(t.actor, &actorTravelIntermediateType{Object: v})

}

// PrependActorObject adds to the front of actor a ObjectType type
func (t *Travel) PrependActorObject(v ObjectType) {
	t.actor = append([]*actorTravelIntermediateType{&actorTravelIntermediateType{Object: v}}, t.actor...)

}

// RemoveActorObject deletes the value from the specified index
func (t *Travel) RemoveActorObject(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// IsActorLink determines whether the call to GetActorLink is safe for the specified index
func (t *Travel) IsActorLink(index int) (ok bool) {
	return t.actor[index].Link != nil

}

// GetActorLink returns the value safely if IsActorLink returned true for the specified index
func (t *Travel) GetActorLink(index int) (v LinkType) {
	return t.actor[index].Link

}

// AppendActorLink adds to the back of actor a LinkType type
func (t *Travel) AppendActorLink(v LinkType) {
	t.actor = append(t.actor, &actorTravelIntermediateType{Link: v})

}

// PrependActorLink adds to the front of actor a LinkType type
func (t *Travel) PrependActorLink(v LinkType) {
	t.actor = append([]*actorTravelIntermediateType{&actorTravelIntermediateType{Link: v}}, t.actor...)

}

// RemoveActorLink deletes the value from the specified index
func (t *Travel) RemoveActorLink(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// IsActorIRI determines whether the call to GetActorIRI is safe for the specified index
func (t *Travel) IsActorIRI(index int) (ok bool) {
	return t.actor[index].IRI != nil

}

// GetActorIRI returns the value safely if IsActorIRI returned true for the specified index
func (t *Travel) GetActorIRI(index int) (v *url.URL) {
	return t.actor[index].IRI

}

// AppendActorIRI adds to the back of actor a *url.URL type
func (t *Travel) AppendActorIRI(v *url.URL) {
	t.actor = append(t.actor, &actorTravelIntermediateType{IRI: v})

}

// PrependActorIRI adds to the front of actor a *url.URL type
func (t *Travel) PrependActorIRI(v *url.URL) {
	t.actor = append([]*actorTravelIntermediateType{&actorTravelIntermediateType{IRI: v}}, t.actor...)

}

// RemoveActorIRI deletes the value from the specified index
func (t *Travel) RemoveActorIRI(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// HasUnknownActor determines whether the call to GetUnknownActor is safe
func (t *Travel) HasUnknownActor() (ok bool) {
	return t.actor != nil && t.actor[0].unknown_ != nil

}

// GetUnknownActor returns the unknown value for actor
func (t *Travel) GetUnknownActor() (v interface{}) {
	return t.actor[0].unknown_

}

// SetUnknownActor sets the unknown value of actor
func (t *Travel) SetUnknownActor(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &actorTravelIntermediateType{}
	tmp.unknown_ = i
	t.actor = append(t.actor, tmp)

}

// TargetLen determines the number of elements able to be used for the IsTargetObject, GetTargetObject, and RemoveTargetObject functions
func (t *Travel) TargetLen() (l int) {
	return len(t.target)

}

// IsTargetObject determines whether the call to GetTargetObject is safe for the specified index
func (t *Travel) IsTargetObject(index int) (ok bool) {
	return t.target[index].Object != nil

}

// GetTargetObject returns the value safely if IsTargetObject returned true for the specified index
func (t *Travel) GetTargetObject(index int) (v ObjectType) {
	return t.target[index].Object

}

// AppendTargetObject adds to the back of target a ObjectType type
func (t *Travel) AppendTargetObject(v ObjectType) {
	t.target = append(t.target, &targetTravelIntermediateType{Object: v})

}

// PrependTargetObject adds to the front of target a ObjectType type
func (t *Travel) PrependTargetObject(v ObjectType) {
	t.target = append([]*targetTravelIntermediateType{&targetTravelIntermediateType{Object: v}}, t.target...)

}

// RemoveTargetObject deletes the value from the specified index
func (t *Travel) RemoveTargetObject(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// IsTargetLink determines whether the call to GetTargetLink is safe for the specified index
func (t *Travel) IsTargetLink(index int) (ok bool) {
	return t.target[index].Link != nil

}

// GetTargetLink returns the value safely if IsTargetLink returned true for the specified index
func (t *Travel) GetTargetLink(index int) (v LinkType) {
	return t.target[index].Link

}

// AppendTargetLink adds to the back of target a LinkType type
func (t *Travel) AppendTargetLink(v LinkType) {
	t.target = append(t.target, &targetTravelIntermediateType{Link: v})

}

// PrependTargetLink adds to the front of target a LinkType type
func (t *Travel) PrependTargetLink(v LinkType) {
	t.target = append([]*targetTravelIntermediateType{&targetTravelIntermediateType{Link: v}}, t.target...)

}

// RemoveTargetLink deletes the value from the specified index
func (t *Travel) RemoveTargetLink(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// IsTargetIRI determines whether the call to GetTargetIRI is safe for the specified index
func (t *Travel) IsTargetIRI(index int) (ok bool) {
	return t.target[index].IRI != nil

}

// GetTargetIRI returns the value safely if IsTargetIRI returned true for the specified index
func (t *Travel) GetTargetIRI(index int) (v *url.URL) {
	return t.target[index].IRI

}

// AppendTargetIRI adds to the back of target a *url.URL type
func (t *Travel) AppendTargetIRI(v *url.URL) {
	t.target = append(t.target, &targetTravelIntermediateType{IRI: v})

}

// PrependTargetIRI adds to the front of target a *url.URL type
func (t *Travel) PrependTargetIRI(v *url.URL) {
	t.target = append([]*targetTravelIntermediateType{&targetTravelIntermediateType{IRI: v}}, t.target...)

}

// RemoveTargetIRI deletes the value from the specified index
func (t *Travel) RemoveTargetIRI(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// HasUnknownTarget determines whether the call to GetUnknownTarget is safe
func (t *Travel) HasUnknownTarget() (ok bool) {
	return t.target != nil && t.target[0].unknown_ != nil

}

// GetUnknownTarget returns the unknown value for target
func (t *Travel) GetUnknownTarget() (v interface{}) {
	return t.target[0].unknown_

}

// SetUnknownTarget sets the unknown value of target
func (t *Travel) SetUnknownTarget(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &targetTravelIntermediateType{}
	tmp.unknown_ = i
	t.target = append(t.target, tmp)

}

// ResultLen determines the number of elements able to be used for the IsResultObject, GetResultObject, and RemoveResultObject functions
func (t *Travel) ResultLen() (l int) {
	return len(t.result)

}

// IsResultObject determines whether the call to GetResultObject is safe for the specified index
func (t *Travel) IsResultObject(index int) (ok bool) {
	return t.result[index].Object != nil

}

// GetResultObject returns the value safely if IsResultObject returned true for the specified index
func (t *Travel) GetResultObject(index int) (v ObjectType) {
	return t.result[index].Object

}

// AppendResultObject adds to the back of result a ObjectType type
func (t *Travel) AppendResultObject(v ObjectType) {
	t.result = append(t.result, &resultTravelIntermediateType{Object: v})

}

// PrependResultObject adds to the front of result a ObjectType type
func (t *Travel) PrependResultObject(v ObjectType) {
	t.result = append([]*resultTravelIntermediateType{&resultTravelIntermediateType{Object: v}}, t.result...)

}

// RemoveResultObject deletes the value from the specified index
func (t *Travel) RemoveResultObject(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// IsResultLink determines whether the call to GetResultLink is safe for the specified index
func (t *Travel) IsResultLink(index int) (ok bool) {
	return t.result[index].Link != nil

}

// GetResultLink returns the value safely if IsResultLink returned true for the specified index
func (t *Travel) GetResultLink(index int) (v LinkType) {
	return t.result[index].Link

}

// AppendResultLink adds to the back of result a LinkType type
func (t *Travel) AppendResultLink(v LinkType) {
	t.result = append(t.result, &resultTravelIntermediateType{Link: v})

}

// PrependResultLink adds to the front of result a LinkType type
func (t *Travel) PrependResultLink(v LinkType) {
	t.result = append([]*resultTravelIntermediateType{&resultTravelIntermediateType{Link: v}}, t.result...)

}

// RemoveResultLink deletes the value from the specified index
func (t *Travel) RemoveResultLink(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// IsResultIRI determines whether the call to GetResultIRI is safe for the specified index
func (t *Travel) IsResultIRI(index int) (ok bool) {
	return t.result[index].IRI != nil

}

// GetResultIRI returns the value safely if IsResultIRI returned true for the specified index
func (t *Travel) GetResultIRI(index int) (v *url.URL) {
	return t.result[index].IRI

}

// AppendResultIRI adds to the back of result a *url.URL type
func (t *Travel) AppendResultIRI(v *url.URL) {
	t.result = append(t.result, &resultTravelIntermediateType{IRI: v})

}

// PrependResultIRI adds to the front of result a *url.URL type
func (t *Travel) PrependResultIRI(v *url.URL) {
	t.result = append([]*resultTravelIntermediateType{&resultTravelIntermediateType{IRI: v}}, t.result...)

}

// RemoveResultIRI deletes the value from the specified index
func (t *Travel) RemoveResultIRI(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// HasUnknownResult determines whether the call to GetUnknownResult is safe
func (t *Travel) HasUnknownResult() (ok bool) {
	return t.result != nil && t.result[0].unknown_ != nil

}

// GetUnknownResult returns the unknown value for result
func (t *Travel) GetUnknownResult() (v interface{}) {
	return t.result[0].unknown_

}

// SetUnknownResult sets the unknown value of result
func (t *Travel) SetUnknownResult(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &resultTravelIntermediateType{}
	tmp.unknown_ = i
	t.result = append(t.result, tmp)

}

// OriginLen determines the number of elements able to be used for the IsOriginObject, GetOriginObject, and RemoveOriginObject functions
func (t *Travel) OriginLen() (l int) {
	return len(t.origin)

}

// IsOriginObject determines whether the call to GetOriginObject is safe for the specified index
func (t *Travel) IsOriginObject(index int) (ok bool) {
	return t.origin[index].Object != nil

}

// GetOriginObject returns the value safely if IsOriginObject returned true for the specified index
func (t *Travel) GetOriginObject(index int) (v ObjectType) {
	return t.origin[index].Object

}

// AppendOriginObject adds to the back of origin a ObjectType type
func (t *Travel) AppendOriginObject(v ObjectType) {
	t.origin = append(t.origin, &originTravelIntermediateType{Object: v})

}

// PrependOriginObject adds to the front of origin a ObjectType type
func (t *Travel) PrependOriginObject(v ObjectType) {
	t.origin = append([]*originTravelIntermediateType{&originTravelIntermediateType{Object: v}}, t.origin...)

}

// RemoveOriginObject deletes the value from the specified index
func (t *Travel) RemoveOriginObject(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// IsOriginLink determines whether the call to GetOriginLink is safe for the specified index
func (t *Travel) IsOriginLink(index int) (ok bool) {
	return t.origin[index].Link != nil

}

// GetOriginLink returns the value safely if IsOriginLink returned true for the specified index
func (t *Travel) GetOriginLink(index int) (v LinkType) {
	return t.origin[index].Link

}

// AppendOriginLink adds to the back of origin a LinkType type
func (t *Travel) AppendOriginLink(v LinkType) {
	t.origin = append(t.origin, &originTravelIntermediateType{Link: v})

}

// PrependOriginLink adds to the front of origin a LinkType type
func (t *Travel) PrependOriginLink(v LinkType) {
	t.origin = append([]*originTravelIntermediateType{&originTravelIntermediateType{Link: v}}, t.origin...)

}

// RemoveOriginLink deletes the value from the specified index
func (t *Travel) RemoveOriginLink(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// IsOriginIRI determines whether the call to GetOriginIRI is safe for the specified index
func (t *Travel) IsOriginIRI(index int) (ok bool) {
	return t.origin[index].IRI != nil

}

// GetOriginIRI returns the value safely if IsOriginIRI returned true for the specified index
func (t *Travel) GetOriginIRI(index int) (v *url.URL) {
	return t.origin[index].IRI

}

// AppendOriginIRI adds to the back of origin a *url.URL type
func (t *Travel) AppendOriginIRI(v *url.URL) {
	t.origin = append(t.origin, &originTravelIntermediateType{IRI: v})

}

// PrependOriginIRI adds to the front of origin a *url.URL type
func (t *Travel) PrependOriginIRI(v *url.URL) {
	t.origin = append([]*originTravelIntermediateType{&originTravelIntermediateType{IRI: v}}, t.origin...)

}

// RemoveOriginIRI deletes the value from the specified index
func (t *Travel) RemoveOriginIRI(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// HasUnknownOrigin determines whether the call to GetUnknownOrigin is safe
func (t *Travel) HasUnknownOrigin() (ok bool) {
	return t.origin != nil && t.origin[0].unknown_ != nil

}

// GetUnknownOrigin returns the unknown value for origin
func (t *Travel) GetUnknownOrigin() (v interface{}) {
	return t.origin[0].unknown_

}

// SetUnknownOrigin sets the unknown value of origin
func (t *Travel) SetUnknownOrigin(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &originTravelIntermediateType{}
	tmp.unknown_ = i
	t.origin = append(t.origin, tmp)

}

// InstrumentLen determines the number of elements able to be used for the IsInstrumentObject, GetInstrumentObject, and RemoveInstrumentObject functions
func (t *Travel) InstrumentLen() (l int) {
	return len(t.instrument)

}

// IsInstrumentObject determines whether the call to GetInstrumentObject is safe for the specified index
func (t *Travel) IsInstrumentObject(index int) (ok bool) {
	return t.instrument[index].Object != nil

}

// GetInstrumentObject returns the value safely if IsInstrumentObject returned true for the specified index
func (t *Travel) GetInstrumentObject(index int) (v ObjectType) {
	return t.instrument[index].Object

}

// AppendInstrumentObject adds to the back of instrument a ObjectType type
func (t *Travel) AppendInstrumentObject(v ObjectType) {
	t.instrument = append(t.instrument, &instrumentTravelIntermediateType{Object: v})

}

// PrependInstrumentObject adds to the front of instrument a ObjectType type
func (t *Travel) PrependInstrumentObject(v ObjectType) {
	t.instrument = append([]*instrumentTravelIntermediateType{&instrumentTravelIntermediateType{Object: v}}, t.instrument...)

}

// RemoveInstrumentObject deletes the value from the specified index
func (t *Travel) RemoveInstrumentObject(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// IsInstrumentLink determines whether the call to GetInstrumentLink is safe for the specified index
func (t *Travel) IsInstrumentLink(index int) (ok bool) {
	return t.instrument[index].Link != nil

}

// GetInstrumentLink returns the value safely if IsInstrumentLink returned true for the specified index
func (t *Travel) GetInstrumentLink(index int) (v LinkType) {
	return t.instrument[index].Link

}

// AppendInstrumentLink adds to the back of instrument a LinkType type
func (t *Travel) AppendInstrumentLink(v LinkType) {
	t.instrument = append(t.instrument, &instrumentTravelIntermediateType{Link: v})

}

// PrependInstrumentLink adds to the front of instrument a LinkType type
func (t *Travel) PrependInstrumentLink(v LinkType) {
	t.instrument = append([]*instrumentTravelIntermediateType{&instrumentTravelIntermediateType{Link: v}}, t.instrument...)

}

// RemoveInstrumentLink deletes the value from the specified index
func (t *Travel) RemoveInstrumentLink(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// IsInstrumentIRI determines whether the call to GetInstrumentIRI is safe for the specified index
func (t *Travel) IsInstrumentIRI(index int) (ok bool) {
	return t.instrument[index].IRI != nil

}

// GetInstrumentIRI returns the value safely if IsInstrumentIRI returned true for the specified index
func (t *Travel) GetInstrumentIRI(index int) (v *url.URL) {
	return t.instrument[index].IRI

}

// AppendInstrumentIRI adds to the back of instrument a *url.URL type
func (t *Travel) AppendInstrumentIRI(v *url.URL) {
	t.instrument = append(t.instrument, &instrumentTravelIntermediateType{IRI: v})

}

// PrependInstrumentIRI adds to the front of instrument a *url.URL type
func (t *Travel) PrependInstrumentIRI(v *url.URL) {
	t.instrument = append([]*instrumentTravelIntermediateType{&instrumentTravelIntermediateType{IRI: v}}, t.instrument...)

}

// RemoveInstrumentIRI deletes the value from the specified index
func (t *Travel) RemoveInstrumentIRI(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// HasUnknownInstrument determines whether the call to GetUnknownInstrument is safe
func (t *Travel) HasUnknownInstrument() (ok bool) {
	return t.instrument != nil && t.instrument[0].unknown_ != nil

}

// GetUnknownInstrument returns the unknown value for instrument
func (t *Travel) GetUnknownInstrument() (v interface{}) {
	return t.instrument[0].unknown_

}

// SetUnknownInstrument sets the unknown value of instrument
func (t *Travel) SetUnknownInstrument(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &instrumentTravelIntermediateType{}
	tmp.unknown_ = i
	t.instrument = append(t.instrument, tmp)

}

// IsAltitude determines whether the call to GetAltitude is safe
func (t *Travel) IsAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.float != nil

}

// GetAltitude returns the value safely if IsAltitude returned true
func (t *Travel) GetAltitude() (v float64) {
	return *t.altitude.float

}

// SetAltitude sets the value of altitude to be of float64 type
func (t *Travel) SetAltitude(v float64) {
	t.altitude = &altitudeTravelIntermediateType{float: &v}

}

// IsAltitudeIRI determines whether the call to GetAltitudeIRI is safe
func (t *Travel) IsAltitudeIRI() (ok bool) {
	return t.altitude != nil && t.altitude.IRI != nil

}

// GetAltitudeIRI returns the value safely if IsAltitudeIRI returned true
func (t *Travel) GetAltitudeIRI() (v *url.URL) {
	return t.altitude.IRI

}

// SetAltitudeIRI sets the value of altitude to be of *url.URL type
func (t *Travel) SetAltitudeIRI(v *url.URL) {
	t.altitude = &altitudeTravelIntermediateType{IRI: v}

}

// HasUnknownAltitude determines whether the call to GetUnknownAltitude is safe
func (t *Travel) HasUnknownAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.unknown_ != nil

}

// GetUnknownAltitude returns the unknown value for altitude
func (t *Travel) GetUnknownAltitude() (v interface{}) {
	return t.altitude.unknown_

}

// SetUnknownAltitude sets the unknown value of altitude
func (t *Travel) SetUnknownAltitude(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &altitudeTravelIntermediateType{}
	tmp.unknown_ = i
	t.altitude = tmp

}

// AttachmentLen determines the number of elements able to be used for the IsAttachmentObject, GetAttachmentObject, and RemoveAttachmentObject functions
func (t *Travel) AttachmentLen() (l int) {
	return len(t.attachment)

}

// IsAttachmentObject determines whether the call to GetAttachmentObject is safe for the specified index
func (t *Travel) IsAttachmentObject(index int) (ok bool) {
	return t.attachment[index].Object != nil

}

// GetAttachmentObject returns the value safely if IsAttachmentObject returned true for the specified index
func (t *Travel) GetAttachmentObject(index int) (v ObjectType) {
	return t.attachment[index].Object

}

// AppendAttachmentObject adds to the back of attachment a ObjectType type
func (t *Travel) AppendAttachmentObject(v ObjectType) {
	t.attachment = append(t.attachment, &attachmentTravelIntermediateType{Object: v})

}

// PrependAttachmentObject adds to the front of attachment a ObjectType type
func (t *Travel) PrependAttachmentObject(v ObjectType) {
	t.attachment = append([]*attachmentTravelIntermediateType{&attachmentTravelIntermediateType{Object: v}}, t.attachment...)

}

// RemoveAttachmentObject deletes the value from the specified index
func (t *Travel) RemoveAttachmentObject(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentLink determines whether the call to GetAttachmentLink is safe for the specified index
func (t *Travel) IsAttachmentLink(index int) (ok bool) {
	return t.attachment[index].Link != nil

}

// GetAttachmentLink returns the value safely if IsAttachmentLink returned true for the specified index
func (t *Travel) GetAttachmentLink(index int) (v LinkType) {
	return t.attachment[index].Link

}

// AppendAttachmentLink adds to the back of attachment a LinkType type
func (t *Travel) AppendAttachmentLink(v LinkType) {
	t.attachment = append(t.attachment, &attachmentTravelIntermediateType{Link: v})

}

// PrependAttachmentLink adds to the front of attachment a LinkType type
func (t *Travel) PrependAttachmentLink(v LinkType) {
	t.attachment = append([]*attachmentTravelIntermediateType{&attachmentTravelIntermediateType{Link: v}}, t.attachment...)

}

// RemoveAttachmentLink deletes the value from the specified index
func (t *Travel) RemoveAttachmentLink(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentIRI determines whether the call to GetAttachmentIRI is safe for the specified index
func (t *Travel) IsAttachmentIRI(index int) (ok bool) {
	return t.attachment[index].IRI != nil

}

// GetAttachmentIRI returns the value safely if IsAttachmentIRI returned true for the specified index
func (t *Travel) GetAttachmentIRI(index int) (v *url.URL) {
	return t.attachment[index].IRI

}

// AppendAttachmentIRI adds to the back of attachment a *url.URL type
func (t *Travel) AppendAttachmentIRI(v *url.URL) {
	t.attachment = append(t.attachment, &attachmentTravelIntermediateType{IRI: v})

}

// PrependAttachmentIRI adds to the front of attachment a *url.URL type
func (t *Travel) PrependAttachmentIRI(v *url.URL) {
	t.attachment = append([]*attachmentTravelIntermediateType{&attachmentTravelIntermediateType{IRI: v}}, t.attachment...)

}

// RemoveAttachmentIRI deletes the value from the specified index
func (t *Travel) RemoveAttachmentIRI(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// HasUnknownAttachment determines whether the call to GetUnknownAttachment is safe
func (t *Travel) HasUnknownAttachment() (ok bool) {
	return t.attachment != nil && t.attachment[0].unknown_ != nil

}

// GetUnknownAttachment returns the unknown value for attachment
func (t *Travel) GetUnknownAttachment() (v interface{}) {
	return t.attachment[0].unknown_

}

// SetUnknownAttachment sets the unknown value of attachment
func (t *Travel) SetUnknownAttachment(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attachmentTravelIntermediateType{}
	tmp.unknown_ = i
	t.attachment = append(t.attachment, tmp)

}

// AttributedToLen determines the number of elements able to be used for the IsAttributedToObject, GetAttributedToObject, and RemoveAttributedToObject functions
func (t *Travel) AttributedToLen() (l int) {
	return len(t.attributedTo)

}

// IsAttributedToObject determines whether the call to GetAttributedToObject is safe for the specified index
func (t *Travel) IsAttributedToObject(index int) (ok bool) {
	return t.attributedTo[index].Object != nil

}

// GetAttributedToObject returns the value safely if IsAttributedToObject returned true for the specified index
func (t *Travel) GetAttributedToObject(index int) (v ObjectType) {
	return t.attributedTo[index].Object

}

// AppendAttributedToObject adds to the back of attributedTo a ObjectType type
func (t *Travel) AppendAttributedToObject(v ObjectType) {
	t.attributedTo = append(t.attributedTo, &attributedToTravelIntermediateType{Object: v})

}

// PrependAttributedToObject adds to the front of attributedTo a ObjectType type
func (t *Travel) PrependAttributedToObject(v ObjectType) {
	t.attributedTo = append([]*attributedToTravelIntermediateType{&attributedToTravelIntermediateType{Object: v}}, t.attributedTo...)

}

// RemoveAttributedToObject deletes the value from the specified index
func (t *Travel) RemoveAttributedToObject(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToLink determines whether the call to GetAttributedToLink is safe for the specified index
func (t *Travel) IsAttributedToLink(index int) (ok bool) {
	return t.attributedTo[index].Link != nil

}

// GetAttributedToLink returns the value safely if IsAttributedToLink returned true for the specified index
func (t *Travel) GetAttributedToLink(index int) (v LinkType) {
	return t.attributedTo[index].Link

}

// AppendAttributedToLink adds to the back of attributedTo a LinkType type
func (t *Travel) AppendAttributedToLink(v LinkType) {
	t.attributedTo = append(t.attributedTo, &attributedToTravelIntermediateType{Link: v})

}

// PrependAttributedToLink adds to the front of attributedTo a LinkType type
func (t *Travel) PrependAttributedToLink(v LinkType) {
	t.attributedTo = append([]*attributedToTravelIntermediateType{&attributedToTravelIntermediateType{Link: v}}, t.attributedTo...)

}

// RemoveAttributedToLink deletes the value from the specified index
func (t *Travel) RemoveAttributedToLink(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToIRI determines whether the call to GetAttributedToIRI is safe for the specified index
func (t *Travel) IsAttributedToIRI(index int) (ok bool) {
	return t.attributedTo[index].IRI != nil

}

// GetAttributedToIRI returns the value safely if IsAttributedToIRI returned true for the specified index
func (t *Travel) GetAttributedToIRI(index int) (v *url.URL) {
	return t.attributedTo[index].IRI

}

// AppendAttributedToIRI adds to the back of attributedTo a *url.URL type
func (t *Travel) AppendAttributedToIRI(v *url.URL) {
	t.attributedTo = append(t.attributedTo, &attributedToTravelIntermediateType{IRI: v})

}

// PrependAttributedToIRI adds to the front of attributedTo a *url.URL type
func (t *Travel) PrependAttributedToIRI(v *url.URL) {
	t.attributedTo = append([]*attributedToTravelIntermediateType{&attributedToTravelIntermediateType{IRI: v}}, t.attributedTo...)

}

// RemoveAttributedToIRI deletes the value from the specified index
func (t *Travel) RemoveAttributedToIRI(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// HasUnknownAttributedTo determines whether the call to GetUnknownAttributedTo is safe
func (t *Travel) HasUnknownAttributedTo() (ok bool) {
	return t.attributedTo != nil && t.attributedTo[0].unknown_ != nil

}

// GetUnknownAttributedTo returns the unknown value for attributedTo
func (t *Travel) GetUnknownAttributedTo() (v interface{}) {
	return t.attributedTo[0].unknown_

}

// SetUnknownAttributedTo sets the unknown value of attributedTo
func (t *Travel) SetUnknownAttributedTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attributedToTravelIntermediateType{}
	tmp.unknown_ = i
	t.attributedTo = append(t.attributedTo, tmp)

}

// AudienceLen determines the number of elements able to be used for the IsAudienceObject, GetAudienceObject, and RemoveAudienceObject functions
func (t *Travel) AudienceLen() (l int) {
	return len(t.audience)

}

// IsAudienceObject determines whether the call to GetAudienceObject is safe for the specified index
func (t *Travel) IsAudienceObject(index int) (ok bool) {
	return t.audience[index].Object != nil

}

// GetAudienceObject returns the value safely if IsAudienceObject returned true for the specified index
func (t *Travel) GetAudienceObject(index int) (v ObjectType) {
	return t.audience[index].Object

}

// AppendAudienceObject adds to the back of audience a ObjectType type
func (t *Travel) AppendAudienceObject(v ObjectType) {
	t.audience = append(t.audience, &audienceTravelIntermediateType{Object: v})

}

// PrependAudienceObject adds to the front of audience a ObjectType type
func (t *Travel) PrependAudienceObject(v ObjectType) {
	t.audience = append([]*audienceTravelIntermediateType{&audienceTravelIntermediateType{Object: v}}, t.audience...)

}

// RemoveAudienceObject deletes the value from the specified index
func (t *Travel) RemoveAudienceObject(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceLink determines whether the call to GetAudienceLink is safe for the specified index
func (t *Travel) IsAudienceLink(index int) (ok bool) {
	return t.audience[index].Link != nil

}

// GetAudienceLink returns the value safely if IsAudienceLink returned true for the specified index
func (t *Travel) GetAudienceLink(index int) (v LinkType) {
	return t.audience[index].Link

}

// AppendAudienceLink adds to the back of audience a LinkType type
func (t *Travel) AppendAudienceLink(v LinkType) {
	t.audience = append(t.audience, &audienceTravelIntermediateType{Link: v})

}

// PrependAudienceLink adds to the front of audience a LinkType type
func (t *Travel) PrependAudienceLink(v LinkType) {
	t.audience = append([]*audienceTravelIntermediateType{&audienceTravelIntermediateType{Link: v}}, t.audience...)

}

// RemoveAudienceLink deletes the value from the specified index
func (t *Travel) RemoveAudienceLink(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceIRI determines whether the call to GetAudienceIRI is safe for the specified index
func (t *Travel) IsAudienceIRI(index int) (ok bool) {
	return t.audience[index].IRI != nil

}

// GetAudienceIRI returns the value safely if IsAudienceIRI returned true for the specified index
func (t *Travel) GetAudienceIRI(index int) (v *url.URL) {
	return t.audience[index].IRI

}

// AppendAudienceIRI adds to the back of audience a *url.URL type
func (t *Travel) AppendAudienceIRI(v *url.URL) {
	t.audience = append(t.audience, &audienceTravelIntermediateType{IRI: v})

}

// PrependAudienceIRI adds to the front of audience a *url.URL type
func (t *Travel) PrependAudienceIRI(v *url.URL) {
	t.audience = append([]*audienceTravelIntermediateType{&audienceTravelIntermediateType{IRI: v}}, t.audience...)

}

// RemoveAudienceIRI deletes the value from the specified index
func (t *Travel) RemoveAudienceIRI(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// HasUnknownAudience determines whether the call to GetUnknownAudience is safe
func (t *Travel) HasUnknownAudience() (ok bool) {
	return t.audience != nil && t.audience[0].unknown_ != nil

}

// GetUnknownAudience returns the unknown value for audience
func (t *Travel) GetUnknownAudience() (v interface{}) {
	return t.audience[0].unknown_

}

// SetUnknownAudience sets the unknown value of audience
func (t *Travel) SetUnknownAudience(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &audienceTravelIntermediateType{}
	tmp.unknown_ = i
	t.audience = append(t.audience, tmp)

}

// ContentLen determines the number of elements able to be used for the IsContentString, GetContentString, and RemoveContentString functions
func (t *Travel) ContentLen() (l int) {
	return len(t.content)

}

// IsContentString determines whether the call to GetContentString is safe for the specified index
func (t *Travel) IsContentString(index int) (ok bool) {
	return t.content[index].stringName != nil

}

// GetContentString returns the value safely if IsContentString returned true for the specified index
func (t *Travel) GetContentString(index int) (v string) {
	return *t.content[index].stringName

}

// AppendContentString adds to the back of content a string type
func (t *Travel) AppendContentString(v string) {
	t.content = append(t.content, &contentTravelIntermediateType{stringName: &v})

}

// PrependContentString adds to the front of content a string type
func (t *Travel) PrependContentString(v string) {
	t.content = append([]*contentTravelIntermediateType{&contentTravelIntermediateType{stringName: &v}}, t.content...)

}

// RemoveContentString deletes the value from the specified index
func (t *Travel) RemoveContentString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentLangString determines whether the call to GetContentLangString is safe for the specified index
func (t *Travel) IsContentLangString(index int) (ok bool) {
	return t.content[index].langString != nil

}

// GetContentLangString returns the value safely if IsContentLangString returned true for the specified index
func (t *Travel) GetContentLangString(index int) (v string) {
	return *t.content[index].langString

}

// AppendContentLangString adds to the back of content a string type
func (t *Travel) AppendContentLangString(v string) {
	t.content = append(t.content, &contentTravelIntermediateType{langString: &v})

}

// PrependContentLangString adds to the front of content a string type
func (t *Travel) PrependContentLangString(v string) {
	t.content = append([]*contentTravelIntermediateType{&contentTravelIntermediateType{langString: &v}}, t.content...)

}

// RemoveContentLangString deletes the value from the specified index
func (t *Travel) RemoveContentLangString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentIRI determines whether the call to GetContentIRI is safe for the specified index
func (t *Travel) IsContentIRI(index int) (ok bool) {
	return t.content[index].IRI != nil

}

// GetContentIRI returns the value safely if IsContentIRI returned true for the specified index
func (t *Travel) GetContentIRI(index int) (v *url.URL) {
	return t.content[index].IRI

}

// AppendContentIRI adds to the back of content a *url.URL type
func (t *Travel) AppendContentIRI(v *url.URL) {
	t.content = append(t.content, &contentTravelIntermediateType{IRI: v})

}

// PrependContentIRI adds to the front of content a *url.URL type
func (t *Travel) PrependContentIRI(v *url.URL) {
	t.content = append([]*contentTravelIntermediateType{&contentTravelIntermediateType{IRI: v}}, t.content...)

}

// RemoveContentIRI deletes the value from the specified index
func (t *Travel) RemoveContentIRI(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// HasUnknownContent determines whether the call to GetUnknownContent is safe
func (t *Travel) HasUnknownContent() (ok bool) {
	return t.content != nil && t.content[0].unknown_ != nil

}

// GetUnknownContent returns the unknown value for content
func (t *Travel) GetUnknownContent() (v interface{}) {
	return t.content[0].unknown_

}

// SetUnknownContent sets the unknown value of content
func (t *Travel) SetUnknownContent(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contentTravelIntermediateType{}
	tmp.unknown_ = i
	t.content = append(t.content, tmp)

}

// ContentMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Travel) ContentMapLanguages() (l []string) {
	if t.contentMap == nil || len(t.contentMap) == 0 {
		return nil
	}
	for k := range t.contentMap {
		l = append(l, k)
	}
	return

}

// GetContentMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Travel) GetContentMap(l string) (v string) {
	if t.contentMap == nil {
		return ""
	}
	ok := false
	v, ok = t.contentMap[l]
	if !ok {
		return ""
	}
	return v

}

// SetContentMap sets the value of the property for the specified language
func (t *Travel) SetContentMap(l string, v string) {
	if t.contentMap == nil {
		t.contentMap = make(map[string]string)
	}
	t.contentMap[l] = v

}

// ContextLen determines the number of elements able to be used for the IsContextObject, GetContextObject, and RemoveContextObject functions
func (t *Travel) ContextLen() (l int) {
	return len(t.context)

}

// IsContextObject determines whether the call to GetContextObject is safe for the specified index
func (t *Travel) IsContextObject(index int) (ok bool) {
	return t.context[index].Object != nil

}

// GetContextObject returns the value safely if IsContextObject returned true for the specified index
func (t *Travel) GetContextObject(index int) (v ObjectType) {
	return t.context[index].Object

}

// AppendContextObject adds to the back of context a ObjectType type
func (t *Travel) AppendContextObject(v ObjectType) {
	t.context = append(t.context, &contextTravelIntermediateType{Object: v})

}

// PrependContextObject adds to the front of context a ObjectType type
func (t *Travel) PrependContextObject(v ObjectType) {
	t.context = append([]*contextTravelIntermediateType{&contextTravelIntermediateType{Object: v}}, t.context...)

}

// RemoveContextObject deletes the value from the specified index
func (t *Travel) RemoveContextObject(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextLink determines whether the call to GetContextLink is safe for the specified index
func (t *Travel) IsContextLink(index int) (ok bool) {
	return t.context[index].Link != nil

}

// GetContextLink returns the value safely if IsContextLink returned true for the specified index
func (t *Travel) GetContextLink(index int) (v LinkType) {
	return t.context[index].Link

}

// AppendContextLink adds to the back of context a LinkType type
func (t *Travel) AppendContextLink(v LinkType) {
	t.context = append(t.context, &contextTravelIntermediateType{Link: v})

}

// PrependContextLink adds to the front of context a LinkType type
func (t *Travel) PrependContextLink(v LinkType) {
	t.context = append([]*contextTravelIntermediateType{&contextTravelIntermediateType{Link: v}}, t.context...)

}

// RemoveContextLink deletes the value from the specified index
func (t *Travel) RemoveContextLink(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextIRI determines whether the call to GetContextIRI is safe for the specified index
func (t *Travel) IsContextIRI(index int) (ok bool) {
	return t.context[index].IRI != nil

}

// GetContextIRI returns the value safely if IsContextIRI returned true for the specified index
func (t *Travel) GetContextIRI(index int) (v *url.URL) {
	return t.context[index].IRI

}

// AppendContextIRI adds to the back of context a *url.URL type
func (t *Travel) AppendContextIRI(v *url.URL) {
	t.context = append(t.context, &contextTravelIntermediateType{IRI: v})

}

// PrependContextIRI adds to the front of context a *url.URL type
func (t *Travel) PrependContextIRI(v *url.URL) {
	t.context = append([]*contextTravelIntermediateType{&contextTravelIntermediateType{IRI: v}}, t.context...)

}

// RemoveContextIRI deletes the value from the specified index
func (t *Travel) RemoveContextIRI(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// HasUnknownContext determines whether the call to GetUnknownContext is safe
func (t *Travel) HasUnknownContext() (ok bool) {
	return t.context != nil && t.context[0].unknown_ != nil

}

// GetUnknownContext returns the unknown value for context
func (t *Travel) GetUnknownContext() (v interface{}) {
	return t.context[0].unknown_

}

// SetUnknownContext sets the unknown value of context
func (t *Travel) SetUnknownContext(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contextTravelIntermediateType{}
	tmp.unknown_ = i
	t.context = append(t.context, tmp)

}

// NameLen determines the number of elements able to be used for the IsNameString, GetNameString, and RemoveNameString functions
func (t *Travel) NameLen() (l int) {
	return len(t.name)

}

// IsNameString determines whether the call to GetNameString is safe for the specified index
func (t *Travel) IsNameString(index int) (ok bool) {
	return t.name[index].stringName != nil

}

// GetNameString returns the value safely if IsNameString returned true for the specified index
func (t *Travel) GetNameString(index int) (v string) {
	return *t.name[index].stringName

}

// AppendNameString adds to the back of name a string type
func (t *Travel) AppendNameString(v string) {
	t.name = append(t.name, &nameTravelIntermediateType{stringName: &v})

}

// PrependNameString adds to the front of name a string type
func (t *Travel) PrependNameString(v string) {
	t.name = append([]*nameTravelIntermediateType{&nameTravelIntermediateType{stringName: &v}}, t.name...)

}

// RemoveNameString deletes the value from the specified index
func (t *Travel) RemoveNameString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameLangString determines whether the call to GetNameLangString is safe for the specified index
func (t *Travel) IsNameLangString(index int) (ok bool) {
	return t.name[index].langString != nil

}

// GetNameLangString returns the value safely if IsNameLangString returned true for the specified index
func (t *Travel) GetNameLangString(index int) (v string) {
	return *t.name[index].langString

}

// AppendNameLangString adds to the back of name a string type
func (t *Travel) AppendNameLangString(v string) {
	t.name = append(t.name, &nameTravelIntermediateType{langString: &v})

}

// PrependNameLangString adds to the front of name a string type
func (t *Travel) PrependNameLangString(v string) {
	t.name = append([]*nameTravelIntermediateType{&nameTravelIntermediateType{langString: &v}}, t.name...)

}

// RemoveNameLangString deletes the value from the specified index
func (t *Travel) RemoveNameLangString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameIRI determines whether the call to GetNameIRI is safe for the specified index
func (t *Travel) IsNameIRI(index int) (ok bool) {
	return t.name[index].IRI != nil

}

// GetNameIRI returns the value safely if IsNameIRI returned true for the specified index
func (t *Travel) GetNameIRI(index int) (v *url.URL) {
	return t.name[index].IRI

}

// AppendNameIRI adds to the back of name a *url.URL type
func (t *Travel) AppendNameIRI(v *url.URL) {
	t.name = append(t.name, &nameTravelIntermediateType{IRI: v})

}

// PrependNameIRI adds to the front of name a *url.URL type
func (t *Travel) PrependNameIRI(v *url.URL) {
	t.name = append([]*nameTravelIntermediateType{&nameTravelIntermediateType{IRI: v}}, t.name...)

}

// RemoveNameIRI deletes the value from the specified index
func (t *Travel) RemoveNameIRI(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// HasUnknownName determines whether the call to GetUnknownName is safe
func (t *Travel) HasUnknownName() (ok bool) {
	return t.name != nil && t.name[0].unknown_ != nil

}

// GetUnknownName returns the unknown value for name
func (t *Travel) GetUnknownName() (v interface{}) {
	return t.name[0].unknown_

}

// SetUnknownName sets the unknown value of name
func (t *Travel) SetUnknownName(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &nameTravelIntermediateType{}
	tmp.unknown_ = i
	t.name = append(t.name, tmp)

}

// NameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Travel) NameMapLanguages() (l []string) {
	if t.nameMap == nil || len(t.nameMap) == 0 {
		return nil
	}
	for k := range t.nameMap {
		l = append(l, k)
	}
	return

}

// GetNameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Travel) GetNameMap(l string) (v string) {
	if t.nameMap == nil {
		return ""
	}
	ok := false
	v, ok = t.nameMap[l]
	if !ok {
		return ""
	}
	return v

}

// SetNameMap sets the value of the property for the specified language
func (t *Travel) SetNameMap(l string, v string) {
	if t.nameMap == nil {
		t.nameMap = make(map[string]string)
	}
	t.nameMap[l] = v

}

// IsEndTime determines whether the call to GetEndTime is safe
func (t *Travel) IsEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.dateTime != nil

}

// GetEndTime returns the value safely if IsEndTime returned true
func (t *Travel) GetEndTime() (v time.Time) {
	return *t.endTime.dateTime

}

// SetEndTime sets the value of endTime to be of time.Time type
func (t *Travel) SetEndTime(v time.Time) {
	t.endTime = &endTimeTravelIntermediateType{dateTime: &v}

}

// IsEndTimeIRI determines whether the call to GetEndTimeIRI is safe
func (t *Travel) IsEndTimeIRI() (ok bool) {
	return t.endTime != nil && t.endTime.IRI != nil

}

// GetEndTimeIRI returns the value safely if IsEndTimeIRI returned true
func (t *Travel) GetEndTimeIRI() (v *url.URL) {
	return t.endTime.IRI

}

// SetEndTimeIRI sets the value of endTime to be of *url.URL type
func (t *Travel) SetEndTimeIRI(v *url.URL) {
	t.endTime = &endTimeTravelIntermediateType{IRI: v}

}

// HasUnknownEndTime determines whether the call to GetUnknownEndTime is safe
func (t *Travel) HasUnknownEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.unknown_ != nil

}

// GetUnknownEndTime returns the unknown value for endTime
func (t *Travel) GetUnknownEndTime() (v interface{}) {
	return t.endTime.unknown_

}

// SetUnknownEndTime sets the unknown value of endTime
func (t *Travel) SetUnknownEndTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endTimeTravelIntermediateType{}
	tmp.unknown_ = i
	t.endTime = tmp

}

// GeneratorLen determines the number of elements able to be used for the IsGeneratorObject, GetGeneratorObject, and RemoveGeneratorObject functions
func (t *Travel) GeneratorLen() (l int) {
	return len(t.generator)

}

// IsGeneratorObject determines whether the call to GetGeneratorObject is safe for the specified index
func (t *Travel) IsGeneratorObject(index int) (ok bool) {
	return t.generator[index].Object != nil

}

// GetGeneratorObject returns the value safely if IsGeneratorObject returned true for the specified index
func (t *Travel) GetGeneratorObject(index int) (v ObjectType) {
	return t.generator[index].Object

}

// AppendGeneratorObject adds to the back of generator a ObjectType type
func (t *Travel) AppendGeneratorObject(v ObjectType) {
	t.generator = append(t.generator, &generatorTravelIntermediateType{Object: v})

}

// PrependGeneratorObject adds to the front of generator a ObjectType type
func (t *Travel) PrependGeneratorObject(v ObjectType) {
	t.generator = append([]*generatorTravelIntermediateType{&generatorTravelIntermediateType{Object: v}}, t.generator...)

}

// RemoveGeneratorObject deletes the value from the specified index
func (t *Travel) RemoveGeneratorObject(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorLink determines whether the call to GetGeneratorLink is safe for the specified index
func (t *Travel) IsGeneratorLink(index int) (ok bool) {
	return t.generator[index].Link != nil

}

// GetGeneratorLink returns the value safely if IsGeneratorLink returned true for the specified index
func (t *Travel) GetGeneratorLink(index int) (v LinkType) {
	return t.generator[index].Link

}

// AppendGeneratorLink adds to the back of generator a LinkType type
func (t *Travel) AppendGeneratorLink(v LinkType) {
	t.generator = append(t.generator, &generatorTravelIntermediateType{Link: v})

}

// PrependGeneratorLink adds to the front of generator a LinkType type
func (t *Travel) PrependGeneratorLink(v LinkType) {
	t.generator = append([]*generatorTravelIntermediateType{&generatorTravelIntermediateType{Link: v}}, t.generator...)

}

// RemoveGeneratorLink deletes the value from the specified index
func (t *Travel) RemoveGeneratorLink(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorIRI determines whether the call to GetGeneratorIRI is safe for the specified index
func (t *Travel) IsGeneratorIRI(index int) (ok bool) {
	return t.generator[index].IRI != nil

}

// GetGeneratorIRI returns the value safely if IsGeneratorIRI returned true for the specified index
func (t *Travel) GetGeneratorIRI(index int) (v *url.URL) {
	return t.generator[index].IRI

}

// AppendGeneratorIRI adds to the back of generator a *url.URL type
func (t *Travel) AppendGeneratorIRI(v *url.URL) {
	t.generator = append(t.generator, &generatorTravelIntermediateType{IRI: v})

}

// PrependGeneratorIRI adds to the front of generator a *url.URL type
func (t *Travel) PrependGeneratorIRI(v *url.URL) {
	t.generator = append([]*generatorTravelIntermediateType{&generatorTravelIntermediateType{IRI: v}}, t.generator...)

}

// RemoveGeneratorIRI deletes the value from the specified index
func (t *Travel) RemoveGeneratorIRI(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// HasUnknownGenerator determines whether the call to GetUnknownGenerator is safe
func (t *Travel) HasUnknownGenerator() (ok bool) {
	return t.generator != nil && t.generator[0].unknown_ != nil

}

// GetUnknownGenerator returns the unknown value for generator
func (t *Travel) GetUnknownGenerator() (v interface{}) {
	return t.generator[0].unknown_

}

// SetUnknownGenerator sets the unknown value of generator
func (t *Travel) SetUnknownGenerator(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &generatorTravelIntermediateType{}
	tmp.unknown_ = i
	t.generator = append(t.generator, tmp)

}

// IconLen determines the number of elements able to be used for the IsIconImage, GetIconImage, and RemoveIconImage functions
func (t *Travel) IconLen() (l int) {
	return len(t.icon)

}

// IsIconImage determines whether the call to GetIconImage is safe for the specified index
func (t *Travel) IsIconImage(index int) (ok bool) {
	return t.icon[index].Image != nil

}

// GetIconImage returns the value safely if IsIconImage returned true for the specified index
func (t *Travel) GetIconImage(index int) (v ImageType) {
	return t.icon[index].Image

}

// AppendIconImage adds to the back of icon a ImageType type
func (t *Travel) AppendIconImage(v ImageType) {
	t.icon = append(t.icon, &iconTravelIntermediateType{Image: v})

}

// PrependIconImage adds to the front of icon a ImageType type
func (t *Travel) PrependIconImage(v ImageType) {
	t.icon = append([]*iconTravelIntermediateType{&iconTravelIntermediateType{Image: v}}, t.icon...)

}

// RemoveIconImage deletes the value from the specified index
func (t *Travel) RemoveIconImage(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconLink determines whether the call to GetIconLink is safe for the specified index
func (t *Travel) IsIconLink(index int) (ok bool) {
	return t.icon[index].Link != nil

}

// GetIconLink returns the value safely if IsIconLink returned true for the specified index
func (t *Travel) GetIconLink(index int) (v LinkType) {
	return t.icon[index].Link

}

// AppendIconLink adds to the back of icon a LinkType type
func (t *Travel) AppendIconLink(v LinkType) {
	t.icon = append(t.icon, &iconTravelIntermediateType{Link: v})

}

// PrependIconLink adds to the front of icon a LinkType type
func (t *Travel) PrependIconLink(v LinkType) {
	t.icon = append([]*iconTravelIntermediateType{&iconTravelIntermediateType{Link: v}}, t.icon...)

}

// RemoveIconLink deletes the value from the specified index
func (t *Travel) RemoveIconLink(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconIRI determines whether the call to GetIconIRI is safe for the specified index
func (t *Travel) IsIconIRI(index int) (ok bool) {
	return t.icon[index].IRI != nil

}

// GetIconIRI returns the value safely if IsIconIRI returned true for the specified index
func (t *Travel) GetIconIRI(index int) (v *url.URL) {
	return t.icon[index].IRI

}

// AppendIconIRI adds to the back of icon a *url.URL type
func (t *Travel) AppendIconIRI(v *url.URL) {
	t.icon = append(t.icon, &iconTravelIntermediateType{IRI: v})

}

// PrependIconIRI adds to the front of icon a *url.URL type
func (t *Travel) PrependIconIRI(v *url.URL) {
	t.icon = append([]*iconTravelIntermediateType{&iconTravelIntermediateType{IRI: v}}, t.icon...)

}

// RemoveIconIRI deletes the value from the specified index
func (t *Travel) RemoveIconIRI(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// HasUnknownIcon determines whether the call to GetUnknownIcon is safe
func (t *Travel) HasUnknownIcon() (ok bool) {
	return t.icon != nil && t.icon[0].unknown_ != nil

}

// GetUnknownIcon returns the unknown value for icon
func (t *Travel) GetUnknownIcon() (v interface{}) {
	return t.icon[0].unknown_

}

// SetUnknownIcon sets the unknown value of icon
func (t *Travel) SetUnknownIcon(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &iconTravelIntermediateType{}
	tmp.unknown_ = i
	t.icon = append(t.icon, tmp)

}

// HasId determines whether the call to GetId is safe
func (t *Travel) HasId() (ok bool) {
	return t.id != nil

}

// GetId returns the value for id
func (t *Travel) GetId() (v *url.URL) {
	return t.id

}

// SetId sets the value of id
func (t *Travel) SetId(v *url.URL) {
	t.id = v

}

// HasUnknownId determines whether the call to GetUnknownId is safe
func (t *Travel) HasUnknownId() (ok bool) {
	return t.unknown_ != nil && t.unknown_["id"] != nil

}

// GetUnknownId returns the unknown value for id
func (t *Travel) GetUnknownId() (v interface{}) {
	return t.unknown_["id"]

}

// SetUnknownId sets the unknown value of id
func (t *Travel) SetUnknownId(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["id"] = i

}

// ImageLen determines the number of elements able to be used for the IsImageImage, GetImageImage, and RemoveImageImage functions
func (t *Travel) ImageLen() (l int) {
	return len(t.image)

}

// IsImageImage determines whether the call to GetImageImage is safe for the specified index
func (t *Travel) IsImageImage(index int) (ok bool) {
	return t.image[index].Image != nil

}

// GetImageImage returns the value safely if IsImageImage returned true for the specified index
func (t *Travel) GetImageImage(index int) (v ImageType) {
	return t.image[index].Image

}

// AppendImageImage adds to the back of image a ImageType type
func (t *Travel) AppendImageImage(v ImageType) {
	t.image = append(t.image, &imageTravelIntermediateType{Image: v})

}

// PrependImageImage adds to the front of image a ImageType type
func (t *Travel) PrependImageImage(v ImageType) {
	t.image = append([]*imageTravelIntermediateType{&imageTravelIntermediateType{Image: v}}, t.image...)

}

// RemoveImageImage deletes the value from the specified index
func (t *Travel) RemoveImageImage(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageLink determines whether the call to GetImageLink is safe for the specified index
func (t *Travel) IsImageLink(index int) (ok bool) {
	return t.image[index].Link != nil

}

// GetImageLink returns the value safely if IsImageLink returned true for the specified index
func (t *Travel) GetImageLink(index int) (v LinkType) {
	return t.image[index].Link

}

// AppendImageLink adds to the back of image a LinkType type
func (t *Travel) AppendImageLink(v LinkType) {
	t.image = append(t.image, &imageTravelIntermediateType{Link: v})

}

// PrependImageLink adds to the front of image a LinkType type
func (t *Travel) PrependImageLink(v LinkType) {
	t.image = append([]*imageTravelIntermediateType{&imageTravelIntermediateType{Link: v}}, t.image...)

}

// RemoveImageLink deletes the value from the specified index
func (t *Travel) RemoveImageLink(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageIRI determines whether the call to GetImageIRI is safe for the specified index
func (t *Travel) IsImageIRI(index int) (ok bool) {
	return t.image[index].IRI != nil

}

// GetImageIRI returns the value safely if IsImageIRI returned true for the specified index
func (t *Travel) GetImageIRI(index int) (v *url.URL) {
	return t.image[index].IRI

}

// AppendImageIRI adds to the back of image a *url.URL type
func (t *Travel) AppendImageIRI(v *url.URL) {
	t.image = append(t.image, &imageTravelIntermediateType{IRI: v})

}

// PrependImageIRI adds to the front of image a *url.URL type
func (t *Travel) PrependImageIRI(v *url.URL) {
	t.image = append([]*imageTravelIntermediateType{&imageTravelIntermediateType{IRI: v}}, t.image...)

}

// RemoveImageIRI deletes the value from the specified index
func (t *Travel) RemoveImageIRI(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// HasUnknownImage determines whether the call to GetUnknownImage is safe
func (t *Travel) HasUnknownImage() (ok bool) {
	return t.image != nil && t.image[0].unknown_ != nil

}

// GetUnknownImage returns the unknown value for image
func (t *Travel) GetUnknownImage() (v interface{}) {
	return t.image[0].unknown_

}

// SetUnknownImage sets the unknown value of image
func (t *Travel) SetUnknownImage(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &imageTravelIntermediateType{}
	tmp.unknown_ = i
	t.image = append(t.image, tmp)

}

// InReplyToLen determines the number of elements able to be used for the IsInReplyToObject, GetInReplyToObject, and RemoveInReplyToObject functions
func (t *Travel) InReplyToLen() (l int) {
	return len(t.inReplyTo)

}

// IsInReplyToObject determines whether the call to GetInReplyToObject is safe for the specified index
func (t *Travel) IsInReplyToObject(index int) (ok bool) {
	return t.inReplyTo[index].Object != nil

}

// GetInReplyToObject returns the value safely if IsInReplyToObject returned true for the specified index
func (t *Travel) GetInReplyToObject(index int) (v ObjectType) {
	return t.inReplyTo[index].Object

}

// AppendInReplyToObject adds to the back of inReplyTo a ObjectType type
func (t *Travel) AppendInReplyToObject(v ObjectType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToTravelIntermediateType{Object: v})

}

// PrependInReplyToObject adds to the front of inReplyTo a ObjectType type
func (t *Travel) PrependInReplyToObject(v ObjectType) {
	t.inReplyTo = append([]*inReplyToTravelIntermediateType{&inReplyToTravelIntermediateType{Object: v}}, t.inReplyTo...)

}

// RemoveInReplyToObject deletes the value from the specified index
func (t *Travel) RemoveInReplyToObject(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToLink determines whether the call to GetInReplyToLink is safe for the specified index
func (t *Travel) IsInReplyToLink(index int) (ok bool) {
	return t.inReplyTo[index].Link != nil

}

// GetInReplyToLink returns the value safely if IsInReplyToLink returned true for the specified index
func (t *Travel) GetInReplyToLink(index int) (v LinkType) {
	return t.inReplyTo[index].Link

}

// AppendInReplyToLink adds to the back of inReplyTo a LinkType type
func (t *Travel) AppendInReplyToLink(v LinkType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToTravelIntermediateType{Link: v})

}

// PrependInReplyToLink adds to the front of inReplyTo a LinkType type
func (t *Travel) PrependInReplyToLink(v LinkType) {
	t.inReplyTo = append([]*inReplyToTravelIntermediateType{&inReplyToTravelIntermediateType{Link: v}}, t.inReplyTo...)

}

// RemoveInReplyToLink deletes the value from the specified index
func (t *Travel) RemoveInReplyToLink(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToIRI determines whether the call to GetInReplyToIRI is safe for the specified index
func (t *Travel) IsInReplyToIRI(index int) (ok bool) {
	return t.inReplyTo[index].IRI != nil

}

// GetInReplyToIRI returns the value safely if IsInReplyToIRI returned true for the specified index
func (t *Travel) GetInReplyToIRI(index int) (v *url.URL) {
	return t.inReplyTo[index].IRI

}

// AppendInReplyToIRI adds to the back of inReplyTo a *url.URL type
func (t *Travel) AppendInReplyToIRI(v *url.URL) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToTravelIntermediateType{IRI: v})

}

// PrependInReplyToIRI adds to the front of inReplyTo a *url.URL type
func (t *Travel) PrependInReplyToIRI(v *url.URL) {
	t.inReplyTo = append([]*inReplyToTravelIntermediateType{&inReplyToTravelIntermediateType{IRI: v}}, t.inReplyTo...)

}

// RemoveInReplyToIRI deletes the value from the specified index
func (t *Travel) RemoveInReplyToIRI(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// HasUnknownInReplyTo determines whether the call to GetUnknownInReplyTo is safe
func (t *Travel) HasUnknownInReplyTo() (ok bool) {
	return t.inReplyTo != nil && t.inReplyTo[0].unknown_ != nil

}

// GetUnknownInReplyTo returns the unknown value for inReplyTo
func (t *Travel) GetUnknownInReplyTo() (v interface{}) {
	return t.inReplyTo[0].unknown_

}

// SetUnknownInReplyTo sets the unknown value of inReplyTo
func (t *Travel) SetUnknownInReplyTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inReplyToTravelIntermediateType{}
	tmp.unknown_ = i
	t.inReplyTo = append(t.inReplyTo, tmp)

}

// LocationLen determines the number of elements able to be used for the IsLocationObject, GetLocationObject, and RemoveLocationObject functions
func (t *Travel) LocationLen() (l int) {
	return len(t.location)

}

// IsLocationObject determines whether the call to GetLocationObject is safe for the specified index
func (t *Travel) IsLocationObject(index int) (ok bool) {
	return t.location[index].Object != nil

}

// GetLocationObject returns the value safely if IsLocationObject returned true for the specified index
func (t *Travel) GetLocationObject(index int) (v ObjectType) {
	return t.location[index].Object

}

// AppendLocationObject adds to the back of location a ObjectType type
func (t *Travel) AppendLocationObject(v ObjectType) {
	t.location = append(t.location, &locationTravelIntermediateType{Object: v})

}

// PrependLocationObject adds to the front of location a ObjectType type
func (t *Travel) PrependLocationObject(v ObjectType) {
	t.location = append([]*locationTravelIntermediateType{&locationTravelIntermediateType{Object: v}}, t.location...)

}

// RemoveLocationObject deletes the value from the specified index
func (t *Travel) RemoveLocationObject(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationLink determines whether the call to GetLocationLink is safe for the specified index
func (t *Travel) IsLocationLink(index int) (ok bool) {
	return t.location[index].Link != nil

}

// GetLocationLink returns the value safely if IsLocationLink returned true for the specified index
func (t *Travel) GetLocationLink(index int) (v LinkType) {
	return t.location[index].Link

}

// AppendLocationLink adds to the back of location a LinkType type
func (t *Travel) AppendLocationLink(v LinkType) {
	t.location = append(t.location, &locationTravelIntermediateType{Link: v})

}

// PrependLocationLink adds to the front of location a LinkType type
func (t *Travel) PrependLocationLink(v LinkType) {
	t.location = append([]*locationTravelIntermediateType{&locationTravelIntermediateType{Link: v}}, t.location...)

}

// RemoveLocationLink deletes the value from the specified index
func (t *Travel) RemoveLocationLink(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationIRI determines whether the call to GetLocationIRI is safe for the specified index
func (t *Travel) IsLocationIRI(index int) (ok bool) {
	return t.location[index].IRI != nil

}

// GetLocationIRI returns the value safely if IsLocationIRI returned true for the specified index
func (t *Travel) GetLocationIRI(index int) (v *url.URL) {
	return t.location[index].IRI

}

// AppendLocationIRI adds to the back of location a *url.URL type
func (t *Travel) AppendLocationIRI(v *url.URL) {
	t.location = append(t.location, &locationTravelIntermediateType{IRI: v})

}

// PrependLocationIRI adds to the front of location a *url.URL type
func (t *Travel) PrependLocationIRI(v *url.URL) {
	t.location = append([]*locationTravelIntermediateType{&locationTravelIntermediateType{IRI: v}}, t.location...)

}

// RemoveLocationIRI deletes the value from the specified index
func (t *Travel) RemoveLocationIRI(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// HasUnknownLocation determines whether the call to GetUnknownLocation is safe
func (t *Travel) HasUnknownLocation() (ok bool) {
	return t.location != nil && t.location[0].unknown_ != nil

}

// GetUnknownLocation returns the unknown value for location
func (t *Travel) GetUnknownLocation() (v interface{}) {
	return t.location[0].unknown_

}

// SetUnknownLocation sets the unknown value of location
func (t *Travel) SetUnknownLocation(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &locationTravelIntermediateType{}
	tmp.unknown_ = i
	t.location = append(t.location, tmp)

}

// PreviewLen determines the number of elements able to be used for the IsPreviewObject, GetPreviewObject, and RemovePreviewObject functions
func (t *Travel) PreviewLen() (l int) {
	return len(t.preview)

}

// IsPreviewObject determines whether the call to GetPreviewObject is safe for the specified index
func (t *Travel) IsPreviewObject(index int) (ok bool) {
	return t.preview[index].Object != nil

}

// GetPreviewObject returns the value safely if IsPreviewObject returned true for the specified index
func (t *Travel) GetPreviewObject(index int) (v ObjectType) {
	return t.preview[index].Object

}

// AppendPreviewObject adds to the back of preview a ObjectType type
func (t *Travel) AppendPreviewObject(v ObjectType) {
	t.preview = append(t.preview, &previewTravelIntermediateType{Object: v})

}

// PrependPreviewObject adds to the front of preview a ObjectType type
func (t *Travel) PrependPreviewObject(v ObjectType) {
	t.preview = append([]*previewTravelIntermediateType{&previewTravelIntermediateType{Object: v}}, t.preview...)

}

// RemovePreviewObject deletes the value from the specified index
func (t *Travel) RemovePreviewObject(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewLink determines whether the call to GetPreviewLink is safe for the specified index
func (t *Travel) IsPreviewLink(index int) (ok bool) {
	return t.preview[index].Link != nil

}

// GetPreviewLink returns the value safely if IsPreviewLink returned true for the specified index
func (t *Travel) GetPreviewLink(index int) (v LinkType) {
	return t.preview[index].Link

}

// AppendPreviewLink adds to the back of preview a LinkType type
func (t *Travel) AppendPreviewLink(v LinkType) {
	t.preview = append(t.preview, &previewTravelIntermediateType{Link: v})

}

// PrependPreviewLink adds to the front of preview a LinkType type
func (t *Travel) PrependPreviewLink(v LinkType) {
	t.preview = append([]*previewTravelIntermediateType{&previewTravelIntermediateType{Link: v}}, t.preview...)

}

// RemovePreviewLink deletes the value from the specified index
func (t *Travel) RemovePreviewLink(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewIRI determines whether the call to GetPreviewIRI is safe for the specified index
func (t *Travel) IsPreviewIRI(index int) (ok bool) {
	return t.preview[index].IRI != nil

}

// GetPreviewIRI returns the value safely if IsPreviewIRI returned true for the specified index
func (t *Travel) GetPreviewIRI(index int) (v *url.URL) {
	return t.preview[index].IRI

}

// AppendPreviewIRI adds to the back of preview a *url.URL type
func (t *Travel) AppendPreviewIRI(v *url.URL) {
	t.preview = append(t.preview, &previewTravelIntermediateType{IRI: v})

}

// PrependPreviewIRI adds to the front of preview a *url.URL type
func (t *Travel) PrependPreviewIRI(v *url.URL) {
	t.preview = append([]*previewTravelIntermediateType{&previewTravelIntermediateType{IRI: v}}, t.preview...)

}

// RemovePreviewIRI deletes the value from the specified index
func (t *Travel) RemovePreviewIRI(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// HasUnknownPreview determines whether the call to GetUnknownPreview is safe
func (t *Travel) HasUnknownPreview() (ok bool) {
	return t.preview != nil && t.preview[0].unknown_ != nil

}

// GetUnknownPreview returns the unknown value for preview
func (t *Travel) GetUnknownPreview() (v interface{}) {
	return t.preview[0].unknown_

}

// SetUnknownPreview sets the unknown value of preview
func (t *Travel) SetUnknownPreview(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &previewTravelIntermediateType{}
	tmp.unknown_ = i
	t.preview = append(t.preview, tmp)

}

// IsPublished determines whether the call to GetPublished is safe
func (t *Travel) IsPublished() (ok bool) {
	return t.published != nil && t.published.dateTime != nil

}

// GetPublished returns the value safely if IsPublished returned true
func (t *Travel) GetPublished() (v time.Time) {
	return *t.published.dateTime

}

// SetPublished sets the value of published to be of time.Time type
func (t *Travel) SetPublished(v time.Time) {
	t.published = &publishedTravelIntermediateType{dateTime: &v}

}

// IsPublishedIRI determines whether the call to GetPublishedIRI is safe
func (t *Travel) IsPublishedIRI() (ok bool) {
	return t.published != nil && t.published.IRI != nil

}

// GetPublishedIRI returns the value safely if IsPublishedIRI returned true
func (t *Travel) GetPublishedIRI() (v *url.URL) {
	return t.published.IRI

}

// SetPublishedIRI sets the value of published to be of *url.URL type
func (t *Travel) SetPublishedIRI(v *url.URL) {
	t.published = &publishedTravelIntermediateType{IRI: v}

}

// HasUnknownPublished determines whether the call to GetUnknownPublished is safe
func (t *Travel) HasUnknownPublished() (ok bool) {
	return t.published != nil && t.published.unknown_ != nil

}

// GetUnknownPublished returns the unknown value for published
func (t *Travel) GetUnknownPublished() (v interface{}) {
	return t.published.unknown_

}

// SetUnknownPublished sets the unknown value of published
func (t *Travel) SetUnknownPublished(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &publishedTravelIntermediateType{}
	tmp.unknown_ = i
	t.published = tmp

}

// IsReplies determines whether the call to GetReplies is safe
func (t *Travel) IsReplies() (ok bool) {
	return t.replies != nil && t.replies.Collection != nil

}

// GetReplies returns the value safely if IsReplies returned true
func (t *Travel) GetReplies() (v CollectionType) {
	return t.replies.Collection

}

// SetReplies sets the value of replies to be of CollectionType type
func (t *Travel) SetReplies(v CollectionType) {
	t.replies = &repliesTravelIntermediateType{Collection: v}

}

// IsRepliesIRI determines whether the call to GetRepliesIRI is safe
func (t *Travel) IsRepliesIRI() (ok bool) {
	return t.replies != nil && t.replies.IRI != nil

}

// GetRepliesIRI returns the value safely if IsRepliesIRI returned true
func (t *Travel) GetRepliesIRI() (v *url.URL) {
	return t.replies.IRI

}

// SetRepliesIRI sets the value of replies to be of *url.URL type
func (t *Travel) SetRepliesIRI(v *url.URL) {
	t.replies = &repliesTravelIntermediateType{IRI: v}

}

// HasUnknownReplies determines whether the call to GetUnknownReplies is safe
func (t *Travel) HasUnknownReplies() (ok bool) {
	return t.replies != nil && t.replies.unknown_ != nil

}

// GetUnknownReplies returns the unknown value for replies
func (t *Travel) GetUnknownReplies() (v interface{}) {
	return t.replies.unknown_

}

// SetUnknownReplies sets the unknown value of replies
func (t *Travel) SetUnknownReplies(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &repliesTravelIntermediateType{}
	tmp.unknown_ = i
	t.replies = tmp

}

// IsStartTime determines whether the call to GetStartTime is safe
func (t *Travel) IsStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.dateTime != nil

}

// GetStartTime returns the value safely if IsStartTime returned true
func (t *Travel) GetStartTime() (v time.Time) {
	return *t.startTime.dateTime

}

// SetStartTime sets the value of startTime to be of time.Time type
func (t *Travel) SetStartTime(v time.Time) {
	t.startTime = &startTimeTravelIntermediateType{dateTime: &v}

}

// IsStartTimeIRI determines whether the call to GetStartTimeIRI is safe
func (t *Travel) IsStartTimeIRI() (ok bool) {
	return t.startTime != nil && t.startTime.IRI != nil

}

// GetStartTimeIRI returns the value safely if IsStartTimeIRI returned true
func (t *Travel) GetStartTimeIRI() (v *url.URL) {
	return t.startTime.IRI

}

// SetStartTimeIRI sets the value of startTime to be of *url.URL type
func (t *Travel) SetStartTimeIRI(v *url.URL) {
	t.startTime = &startTimeTravelIntermediateType{IRI: v}

}

// HasUnknownStartTime determines whether the call to GetUnknownStartTime is safe
func (t *Travel) HasUnknownStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.unknown_ != nil

}

// GetUnknownStartTime returns the unknown value for startTime
func (t *Travel) GetUnknownStartTime() (v interface{}) {
	return t.startTime.unknown_

}

// SetUnknownStartTime sets the unknown value of startTime
func (t *Travel) SetUnknownStartTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &startTimeTravelIntermediateType{}
	tmp.unknown_ = i
	t.startTime = tmp

}

// SummaryLen determines the number of elements able to be used for the IsSummaryString, GetSummaryString, and RemoveSummaryString functions
func (t *Travel) SummaryLen() (l int) {
	return len(t.summary)

}

// IsSummaryString determines whether the call to GetSummaryString is safe for the specified index
func (t *Travel) IsSummaryString(index int) (ok bool) {
	return t.summary[index].stringName != nil

}

// GetSummaryString returns the value safely if IsSummaryString returned true for the specified index
func (t *Travel) GetSummaryString(index int) (v string) {
	return *t.summary[index].stringName

}

// AppendSummaryString adds to the back of summary a string type
func (t *Travel) AppendSummaryString(v string) {
	t.summary = append(t.summary, &summaryTravelIntermediateType{stringName: &v})

}

// PrependSummaryString adds to the front of summary a string type
func (t *Travel) PrependSummaryString(v string) {
	t.summary = append([]*summaryTravelIntermediateType{&summaryTravelIntermediateType{stringName: &v}}, t.summary...)

}

// RemoveSummaryString deletes the value from the specified index
func (t *Travel) RemoveSummaryString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryLangString determines whether the call to GetSummaryLangString is safe for the specified index
func (t *Travel) IsSummaryLangString(index int) (ok bool) {
	return t.summary[index].langString != nil

}

// GetSummaryLangString returns the value safely if IsSummaryLangString returned true for the specified index
func (t *Travel) GetSummaryLangString(index int) (v string) {
	return *t.summary[index].langString

}

// AppendSummaryLangString adds to the back of summary a string type
func (t *Travel) AppendSummaryLangString(v string) {
	t.summary = append(t.summary, &summaryTravelIntermediateType{langString: &v})

}

// PrependSummaryLangString adds to the front of summary a string type
func (t *Travel) PrependSummaryLangString(v string) {
	t.summary = append([]*summaryTravelIntermediateType{&summaryTravelIntermediateType{langString: &v}}, t.summary...)

}

// RemoveSummaryLangString deletes the value from the specified index
func (t *Travel) RemoveSummaryLangString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryIRI determines whether the call to GetSummaryIRI is safe for the specified index
func (t *Travel) IsSummaryIRI(index int) (ok bool) {
	return t.summary[index].IRI != nil

}

// GetSummaryIRI returns the value safely if IsSummaryIRI returned true for the specified index
func (t *Travel) GetSummaryIRI(index int) (v *url.URL) {
	return t.summary[index].IRI

}

// AppendSummaryIRI adds to the back of summary a *url.URL type
func (t *Travel) AppendSummaryIRI(v *url.URL) {
	t.summary = append(t.summary, &summaryTravelIntermediateType{IRI: v})

}

// PrependSummaryIRI adds to the front of summary a *url.URL type
func (t *Travel) PrependSummaryIRI(v *url.URL) {
	t.summary = append([]*summaryTravelIntermediateType{&summaryTravelIntermediateType{IRI: v}}, t.summary...)

}

// RemoveSummaryIRI deletes the value from the specified index
func (t *Travel) RemoveSummaryIRI(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// HasUnknownSummary determines whether the call to GetUnknownSummary is safe
func (t *Travel) HasUnknownSummary() (ok bool) {
	return t.summary != nil && t.summary[0].unknown_ != nil

}

// GetUnknownSummary returns the unknown value for summary
func (t *Travel) GetUnknownSummary() (v interface{}) {
	return t.summary[0].unknown_

}

// SetUnknownSummary sets the unknown value of summary
func (t *Travel) SetUnknownSummary(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &summaryTravelIntermediateType{}
	tmp.unknown_ = i
	t.summary = append(t.summary, tmp)

}

// SummaryMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Travel) SummaryMapLanguages() (l []string) {
	if t.summaryMap == nil || len(t.summaryMap) == 0 {
		return nil
	}
	for k := range t.summaryMap {
		l = append(l, k)
	}
	return

}

// GetSummaryMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Travel) GetSummaryMap(l string) (v string) {
	if t.summaryMap == nil {
		return ""
	}
	ok := false
	v, ok = t.summaryMap[l]
	if !ok {
		return ""
	}
	return v

}

// SetSummaryMap sets the value of the property for the specified language
func (t *Travel) SetSummaryMap(l string, v string) {
	if t.summaryMap == nil {
		t.summaryMap = make(map[string]string)
	}
	t.summaryMap[l] = v

}

// TagLen determines the number of elements able to be used for the IsTagObject, GetTagObject, and RemoveTagObject functions
func (t *Travel) TagLen() (l int) {
	return len(t.tag)

}

// IsTagObject determines whether the call to GetTagObject is safe for the specified index
func (t *Travel) IsTagObject(index int) (ok bool) {
	return t.tag[index].Object != nil

}

// GetTagObject returns the value safely if IsTagObject returned true for the specified index
func (t *Travel) GetTagObject(index int) (v ObjectType) {
	return t.tag[index].Object

}

// AppendTagObject adds to the back of tag a ObjectType type
func (t *Travel) AppendTagObject(v ObjectType) {
	t.tag = append(t.tag, &tagTravelIntermediateType{Object: v})

}

// PrependTagObject adds to the front of tag a ObjectType type
func (t *Travel) PrependTagObject(v ObjectType) {
	t.tag = append([]*tagTravelIntermediateType{&tagTravelIntermediateType{Object: v}}, t.tag...)

}

// RemoveTagObject deletes the value from the specified index
func (t *Travel) RemoveTagObject(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagLink determines whether the call to GetTagLink is safe for the specified index
func (t *Travel) IsTagLink(index int) (ok bool) {
	return t.tag[index].Link != nil

}

// GetTagLink returns the value safely if IsTagLink returned true for the specified index
func (t *Travel) GetTagLink(index int) (v LinkType) {
	return t.tag[index].Link

}

// AppendTagLink adds to the back of tag a LinkType type
func (t *Travel) AppendTagLink(v LinkType) {
	t.tag = append(t.tag, &tagTravelIntermediateType{Link: v})

}

// PrependTagLink adds to the front of tag a LinkType type
func (t *Travel) PrependTagLink(v LinkType) {
	t.tag = append([]*tagTravelIntermediateType{&tagTravelIntermediateType{Link: v}}, t.tag...)

}

// RemoveTagLink deletes the value from the specified index
func (t *Travel) RemoveTagLink(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagIRI determines whether the call to GetTagIRI is safe for the specified index
func (t *Travel) IsTagIRI(index int) (ok bool) {
	return t.tag[index].IRI != nil

}

// GetTagIRI returns the value safely if IsTagIRI returned true for the specified index
func (t *Travel) GetTagIRI(index int) (v *url.URL) {
	return t.tag[index].IRI

}

// AppendTagIRI adds to the back of tag a *url.URL type
func (t *Travel) AppendTagIRI(v *url.URL) {
	t.tag = append(t.tag, &tagTravelIntermediateType{IRI: v})

}

// PrependTagIRI adds to the front of tag a *url.URL type
func (t *Travel) PrependTagIRI(v *url.URL) {
	t.tag = append([]*tagTravelIntermediateType{&tagTravelIntermediateType{IRI: v}}, t.tag...)

}

// RemoveTagIRI deletes the value from the specified index
func (t *Travel) RemoveTagIRI(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// HasUnknownTag determines whether the call to GetUnknownTag is safe
func (t *Travel) HasUnknownTag() (ok bool) {
	return t.tag != nil && t.tag[0].unknown_ != nil

}

// GetUnknownTag returns the unknown value for tag
func (t *Travel) GetUnknownTag() (v interface{}) {
	return t.tag[0].unknown_

}

// SetUnknownTag sets the unknown value of tag
func (t *Travel) SetUnknownTag(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &tagTravelIntermediateType{}
	tmp.unknown_ = i
	t.tag = append(t.tag, tmp)

}

// TypeLen determines the number of elements able to be used for the GetType and RemoveType functions
func (t *Travel) TypeLen() (l int) {
	return len(t.typeName)

}

// GetType returns the value for the specified index
func (t *Travel) GetType(index int) (v interface{}) {
	return t.typeName[index]

}

// AppendType adds a value to the back of type
func (t *Travel) AppendType(v interface{}) {
	t.typeName = append(t.typeName, v)

}

// PrependType adds a value to the front of type
func (t *Travel) PrependType(v interface{}) {
	t.typeName = append([]interface{}{v}, t.typeName...)

}

// RemoveType deletes the value from the specified index
func (t *Travel) RemoveType(index int) {
	copy(t.typeName[index:], t.typeName[index+1:])
	t.typeName[len(t.typeName)-1] = nil
	t.typeName = t.typeName[:len(t.typeName)-1]

}

// IsUpdated determines whether the call to GetUpdated is safe
func (t *Travel) IsUpdated() (ok bool) {
	return t.updated != nil && t.updated.dateTime != nil

}

// GetUpdated returns the value safely if IsUpdated returned true
func (t *Travel) GetUpdated() (v time.Time) {
	return *t.updated.dateTime

}

// SetUpdated sets the value of updated to be of time.Time type
func (t *Travel) SetUpdated(v time.Time) {
	t.updated = &updatedTravelIntermediateType{dateTime: &v}

}

// IsUpdatedIRI determines whether the call to GetUpdatedIRI is safe
func (t *Travel) IsUpdatedIRI() (ok bool) {
	return t.updated != nil && t.updated.IRI != nil

}

// GetUpdatedIRI returns the value safely if IsUpdatedIRI returned true
func (t *Travel) GetUpdatedIRI() (v *url.URL) {
	return t.updated.IRI

}

// SetUpdatedIRI sets the value of updated to be of *url.URL type
func (t *Travel) SetUpdatedIRI(v *url.URL) {
	t.updated = &updatedTravelIntermediateType{IRI: v}

}

// HasUnknownUpdated determines whether the call to GetUnknownUpdated is safe
func (t *Travel) HasUnknownUpdated() (ok bool) {
	return t.updated != nil && t.updated.unknown_ != nil

}

// GetUnknownUpdated returns the unknown value for updated
func (t *Travel) GetUnknownUpdated() (v interface{}) {
	return t.updated.unknown_

}

// SetUnknownUpdated sets the unknown value of updated
func (t *Travel) SetUnknownUpdated(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &updatedTravelIntermediateType{}
	tmp.unknown_ = i
	t.updated = tmp

}

// UrlLen determines the number of elements able to be used for the IsUrlAnyURI, GetUrlAnyURI, and RemoveUrlAnyURI functions
func (t *Travel) UrlLen() (l int) {
	return len(t.url)

}

// IsUrlAnyURI determines whether the call to GetUrlAnyURI is safe for the specified index
func (t *Travel) IsUrlAnyURI(index int) (ok bool) {
	return t.url[index].anyURI != nil

}

// GetUrlAnyURI returns the value safely if IsUrlAnyURI returned true for the specified index
func (t *Travel) GetUrlAnyURI(index int) (v *url.URL) {
	return t.url[index].anyURI

}

// AppendUrlAnyURI adds to the back of url a *url.URL type
func (t *Travel) AppendUrlAnyURI(v *url.URL) {
	t.url = append(t.url, &urlTravelIntermediateType{anyURI: v})

}

// PrependUrlAnyURI adds to the front of url a *url.URL type
func (t *Travel) PrependUrlAnyURI(v *url.URL) {
	t.url = append([]*urlTravelIntermediateType{&urlTravelIntermediateType{anyURI: v}}, t.url...)

}

// RemoveUrlAnyURI deletes the value from the specified index
func (t *Travel) RemoveUrlAnyURI(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// IsUrlLink determines whether the call to GetUrlLink is safe for the specified index
func (t *Travel) IsUrlLink(index int) (ok bool) {
	return t.url[index].Link != nil

}

// GetUrlLink returns the value safely if IsUrlLink returned true for the specified index
func (t *Travel) GetUrlLink(index int) (v LinkType) {
	return t.url[index].Link

}

// AppendUrlLink adds to the back of url a LinkType type
func (t *Travel) AppendUrlLink(v LinkType) {
	t.url = append(t.url, &urlTravelIntermediateType{Link: v})

}

// PrependUrlLink adds to the front of url a LinkType type
func (t *Travel) PrependUrlLink(v LinkType) {
	t.url = append([]*urlTravelIntermediateType{&urlTravelIntermediateType{Link: v}}, t.url...)

}

// RemoveUrlLink deletes the value from the specified index
func (t *Travel) RemoveUrlLink(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// HasUnknownUrl determines whether the call to GetUnknownUrl is safe
func (t *Travel) HasUnknownUrl() (ok bool) {
	return t.url != nil && t.url[0].unknown_ != nil

}

// GetUnknownUrl returns the unknown value for url
func (t *Travel) GetUnknownUrl() (v interface{}) {
	return t.url[0].unknown_

}

// SetUnknownUrl sets the unknown value of url
func (t *Travel) SetUnknownUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &urlTravelIntermediateType{}
	tmp.unknown_ = i
	t.url = append(t.url, tmp)

}

// ToLen determines the number of elements able to be used for the IsToObject, GetToObject, and RemoveToObject functions
func (t *Travel) ToLen() (l int) {
	return len(t.to)

}

// IsToObject determines whether the call to GetToObject is safe for the specified index
func (t *Travel) IsToObject(index int) (ok bool) {
	return t.to[index].Object != nil

}

// GetToObject returns the value safely if IsToObject returned true for the specified index
func (t *Travel) GetToObject(index int) (v ObjectType) {
	return t.to[index].Object

}

// AppendToObject adds to the back of to a ObjectType type
func (t *Travel) AppendToObject(v ObjectType) {
	t.to = append(t.to, &toTravelIntermediateType{Object: v})

}

// PrependToObject adds to the front of to a ObjectType type
func (t *Travel) PrependToObject(v ObjectType) {
	t.to = append([]*toTravelIntermediateType{&toTravelIntermediateType{Object: v}}, t.to...)

}

// RemoveToObject deletes the value from the specified index
func (t *Travel) RemoveToObject(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToLink determines whether the call to GetToLink is safe for the specified index
func (t *Travel) IsToLink(index int) (ok bool) {
	return t.to[index].Link != nil

}

// GetToLink returns the value safely if IsToLink returned true for the specified index
func (t *Travel) GetToLink(index int) (v LinkType) {
	return t.to[index].Link

}

// AppendToLink adds to the back of to a LinkType type
func (t *Travel) AppendToLink(v LinkType) {
	t.to = append(t.to, &toTravelIntermediateType{Link: v})

}

// PrependToLink adds to the front of to a LinkType type
func (t *Travel) PrependToLink(v LinkType) {
	t.to = append([]*toTravelIntermediateType{&toTravelIntermediateType{Link: v}}, t.to...)

}

// RemoveToLink deletes the value from the specified index
func (t *Travel) RemoveToLink(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToIRI determines whether the call to GetToIRI is safe for the specified index
func (t *Travel) IsToIRI(index int) (ok bool) {
	return t.to[index].IRI != nil

}

// GetToIRI returns the value safely if IsToIRI returned true for the specified index
func (t *Travel) GetToIRI(index int) (v *url.URL) {
	return t.to[index].IRI

}

// AppendToIRI adds to the back of to a *url.URL type
func (t *Travel) AppendToIRI(v *url.URL) {
	t.to = append(t.to, &toTravelIntermediateType{IRI: v})

}

// PrependToIRI adds to the front of to a *url.URL type
func (t *Travel) PrependToIRI(v *url.URL) {
	t.to = append([]*toTravelIntermediateType{&toTravelIntermediateType{IRI: v}}, t.to...)

}

// RemoveToIRI deletes the value from the specified index
func (t *Travel) RemoveToIRI(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// HasUnknownTo determines whether the call to GetUnknownTo is safe
func (t *Travel) HasUnknownTo() (ok bool) {
	return t.to != nil && t.to[0].unknown_ != nil

}

// GetUnknownTo returns the unknown value for to
func (t *Travel) GetUnknownTo() (v interface{}) {
	return t.to[0].unknown_

}

// SetUnknownTo sets the unknown value of to
func (t *Travel) SetUnknownTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &toTravelIntermediateType{}
	tmp.unknown_ = i
	t.to = append(t.to, tmp)

}

// BtoLen determines the number of elements able to be used for the IsBtoObject, GetBtoObject, and RemoveBtoObject functions
func (t *Travel) BtoLen() (l int) {
	return len(t.bto)

}

// IsBtoObject determines whether the call to GetBtoObject is safe for the specified index
func (t *Travel) IsBtoObject(index int) (ok bool) {
	return t.bto[index].Object != nil

}

// GetBtoObject returns the value safely if IsBtoObject returned true for the specified index
func (t *Travel) GetBtoObject(index int) (v ObjectType) {
	return t.bto[index].Object

}

// AppendBtoObject adds to the back of bto a ObjectType type
func (t *Travel) AppendBtoObject(v ObjectType) {
	t.bto = append(t.bto, &btoTravelIntermediateType{Object: v})

}

// PrependBtoObject adds to the front of bto a ObjectType type
func (t *Travel) PrependBtoObject(v ObjectType) {
	t.bto = append([]*btoTravelIntermediateType{&btoTravelIntermediateType{Object: v}}, t.bto...)

}

// RemoveBtoObject deletes the value from the specified index
func (t *Travel) RemoveBtoObject(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoLink determines whether the call to GetBtoLink is safe for the specified index
func (t *Travel) IsBtoLink(index int) (ok bool) {
	return t.bto[index].Link != nil

}

// GetBtoLink returns the value safely if IsBtoLink returned true for the specified index
func (t *Travel) GetBtoLink(index int) (v LinkType) {
	return t.bto[index].Link

}

// AppendBtoLink adds to the back of bto a LinkType type
func (t *Travel) AppendBtoLink(v LinkType) {
	t.bto = append(t.bto, &btoTravelIntermediateType{Link: v})

}

// PrependBtoLink adds to the front of bto a LinkType type
func (t *Travel) PrependBtoLink(v LinkType) {
	t.bto = append([]*btoTravelIntermediateType{&btoTravelIntermediateType{Link: v}}, t.bto...)

}

// RemoveBtoLink deletes the value from the specified index
func (t *Travel) RemoveBtoLink(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoIRI determines whether the call to GetBtoIRI is safe for the specified index
func (t *Travel) IsBtoIRI(index int) (ok bool) {
	return t.bto[index].IRI != nil

}

// GetBtoIRI returns the value safely if IsBtoIRI returned true for the specified index
func (t *Travel) GetBtoIRI(index int) (v *url.URL) {
	return t.bto[index].IRI

}

// AppendBtoIRI adds to the back of bto a *url.URL type
func (t *Travel) AppendBtoIRI(v *url.URL) {
	t.bto = append(t.bto, &btoTravelIntermediateType{IRI: v})

}

// PrependBtoIRI adds to the front of bto a *url.URL type
func (t *Travel) PrependBtoIRI(v *url.URL) {
	t.bto = append([]*btoTravelIntermediateType{&btoTravelIntermediateType{IRI: v}}, t.bto...)

}

// RemoveBtoIRI deletes the value from the specified index
func (t *Travel) RemoveBtoIRI(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// HasUnknownBto determines whether the call to GetUnknownBto is safe
func (t *Travel) HasUnknownBto() (ok bool) {
	return t.bto != nil && t.bto[0].unknown_ != nil

}

// GetUnknownBto returns the unknown value for bto
func (t *Travel) GetUnknownBto() (v interface{}) {
	return t.bto[0].unknown_

}

// SetUnknownBto sets the unknown value of bto
func (t *Travel) SetUnknownBto(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &btoTravelIntermediateType{}
	tmp.unknown_ = i
	t.bto = append(t.bto, tmp)

}

// CcLen determines the number of elements able to be used for the IsCcObject, GetCcObject, and RemoveCcObject functions
func (t *Travel) CcLen() (l int) {
	return len(t.cc)

}

// IsCcObject determines whether the call to GetCcObject is safe for the specified index
func (t *Travel) IsCcObject(index int) (ok bool) {
	return t.cc[index].Object != nil

}

// GetCcObject returns the value safely if IsCcObject returned true for the specified index
func (t *Travel) GetCcObject(index int) (v ObjectType) {
	return t.cc[index].Object

}

// AppendCcObject adds to the back of cc a ObjectType type
func (t *Travel) AppendCcObject(v ObjectType) {
	t.cc = append(t.cc, &ccTravelIntermediateType{Object: v})

}

// PrependCcObject adds to the front of cc a ObjectType type
func (t *Travel) PrependCcObject(v ObjectType) {
	t.cc = append([]*ccTravelIntermediateType{&ccTravelIntermediateType{Object: v}}, t.cc...)

}

// RemoveCcObject deletes the value from the specified index
func (t *Travel) RemoveCcObject(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcLink determines whether the call to GetCcLink is safe for the specified index
func (t *Travel) IsCcLink(index int) (ok bool) {
	return t.cc[index].Link != nil

}

// GetCcLink returns the value safely if IsCcLink returned true for the specified index
func (t *Travel) GetCcLink(index int) (v LinkType) {
	return t.cc[index].Link

}

// AppendCcLink adds to the back of cc a LinkType type
func (t *Travel) AppendCcLink(v LinkType) {
	t.cc = append(t.cc, &ccTravelIntermediateType{Link: v})

}

// PrependCcLink adds to the front of cc a LinkType type
func (t *Travel) PrependCcLink(v LinkType) {
	t.cc = append([]*ccTravelIntermediateType{&ccTravelIntermediateType{Link: v}}, t.cc...)

}

// RemoveCcLink deletes the value from the specified index
func (t *Travel) RemoveCcLink(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcIRI determines whether the call to GetCcIRI is safe for the specified index
func (t *Travel) IsCcIRI(index int) (ok bool) {
	return t.cc[index].IRI != nil

}

// GetCcIRI returns the value safely if IsCcIRI returned true for the specified index
func (t *Travel) GetCcIRI(index int) (v *url.URL) {
	return t.cc[index].IRI

}

// AppendCcIRI adds to the back of cc a *url.URL type
func (t *Travel) AppendCcIRI(v *url.URL) {
	t.cc = append(t.cc, &ccTravelIntermediateType{IRI: v})

}

// PrependCcIRI adds to the front of cc a *url.URL type
func (t *Travel) PrependCcIRI(v *url.URL) {
	t.cc = append([]*ccTravelIntermediateType{&ccTravelIntermediateType{IRI: v}}, t.cc...)

}

// RemoveCcIRI deletes the value from the specified index
func (t *Travel) RemoveCcIRI(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// HasUnknownCc determines whether the call to GetUnknownCc is safe
func (t *Travel) HasUnknownCc() (ok bool) {
	return t.cc != nil && t.cc[0].unknown_ != nil

}

// GetUnknownCc returns the unknown value for cc
func (t *Travel) GetUnknownCc() (v interface{}) {
	return t.cc[0].unknown_

}

// SetUnknownCc sets the unknown value of cc
func (t *Travel) SetUnknownCc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &ccTravelIntermediateType{}
	tmp.unknown_ = i
	t.cc = append(t.cc, tmp)

}

// BccLen determines the number of elements able to be used for the IsBccObject, GetBccObject, and RemoveBccObject functions
func (t *Travel) BccLen() (l int) {
	return len(t.bcc)

}

// IsBccObject determines whether the call to GetBccObject is safe for the specified index
func (t *Travel) IsBccObject(index int) (ok bool) {
	return t.bcc[index].Object != nil

}

// GetBccObject returns the value safely if IsBccObject returned true for the specified index
func (t *Travel) GetBccObject(index int) (v ObjectType) {
	return t.bcc[index].Object

}

// AppendBccObject adds to the back of bcc a ObjectType type
func (t *Travel) AppendBccObject(v ObjectType) {
	t.bcc = append(t.bcc, &bccTravelIntermediateType{Object: v})

}

// PrependBccObject adds to the front of bcc a ObjectType type
func (t *Travel) PrependBccObject(v ObjectType) {
	t.bcc = append([]*bccTravelIntermediateType{&bccTravelIntermediateType{Object: v}}, t.bcc...)

}

// RemoveBccObject deletes the value from the specified index
func (t *Travel) RemoveBccObject(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccLink determines whether the call to GetBccLink is safe for the specified index
func (t *Travel) IsBccLink(index int) (ok bool) {
	return t.bcc[index].Link != nil

}

// GetBccLink returns the value safely if IsBccLink returned true for the specified index
func (t *Travel) GetBccLink(index int) (v LinkType) {
	return t.bcc[index].Link

}

// AppendBccLink adds to the back of bcc a LinkType type
func (t *Travel) AppendBccLink(v LinkType) {
	t.bcc = append(t.bcc, &bccTravelIntermediateType{Link: v})

}

// PrependBccLink adds to the front of bcc a LinkType type
func (t *Travel) PrependBccLink(v LinkType) {
	t.bcc = append([]*bccTravelIntermediateType{&bccTravelIntermediateType{Link: v}}, t.bcc...)

}

// RemoveBccLink deletes the value from the specified index
func (t *Travel) RemoveBccLink(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccIRI determines whether the call to GetBccIRI is safe for the specified index
func (t *Travel) IsBccIRI(index int) (ok bool) {
	return t.bcc[index].IRI != nil

}

// GetBccIRI returns the value safely if IsBccIRI returned true for the specified index
func (t *Travel) GetBccIRI(index int) (v *url.URL) {
	return t.bcc[index].IRI

}

// AppendBccIRI adds to the back of bcc a *url.URL type
func (t *Travel) AppendBccIRI(v *url.URL) {
	t.bcc = append(t.bcc, &bccTravelIntermediateType{IRI: v})

}

// PrependBccIRI adds to the front of bcc a *url.URL type
func (t *Travel) PrependBccIRI(v *url.URL) {
	t.bcc = append([]*bccTravelIntermediateType{&bccTravelIntermediateType{IRI: v}}, t.bcc...)

}

// RemoveBccIRI deletes the value from the specified index
func (t *Travel) RemoveBccIRI(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// HasUnknownBcc determines whether the call to GetUnknownBcc is safe
func (t *Travel) HasUnknownBcc() (ok bool) {
	return t.bcc != nil && t.bcc[0].unknown_ != nil

}

// GetUnknownBcc returns the unknown value for bcc
func (t *Travel) GetUnknownBcc() (v interface{}) {
	return t.bcc[0].unknown_

}

// SetUnknownBcc sets the unknown value of bcc
func (t *Travel) SetUnknownBcc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &bccTravelIntermediateType{}
	tmp.unknown_ = i
	t.bcc = append(t.bcc, tmp)

}

// IsMediaType determines whether the call to GetMediaType is safe
func (t *Travel) IsMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.mimeMediaTypeValue != nil

}

// GetMediaType returns the value safely if IsMediaType returned true
func (t *Travel) GetMediaType() (v string) {
	return *t.mediaType.mimeMediaTypeValue

}

// SetMediaType sets the value of mediaType to be of string type
func (t *Travel) SetMediaType(v string) {
	t.mediaType = &mediaTypeTravelIntermediateType{mimeMediaTypeValue: &v}

}

// IsMediaTypeIRI determines whether the call to GetMediaTypeIRI is safe
func (t *Travel) IsMediaTypeIRI() (ok bool) {
	return t.mediaType != nil && t.mediaType.IRI != nil

}

// GetMediaTypeIRI returns the value safely if IsMediaTypeIRI returned true
func (t *Travel) GetMediaTypeIRI() (v *url.URL) {
	return t.mediaType.IRI

}

// SetMediaTypeIRI sets the value of mediaType to be of *url.URL type
func (t *Travel) SetMediaTypeIRI(v *url.URL) {
	t.mediaType = &mediaTypeTravelIntermediateType{IRI: v}

}

// HasUnknownMediaType determines whether the call to GetUnknownMediaType is safe
func (t *Travel) HasUnknownMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.unknown_ != nil

}

// GetUnknownMediaType returns the unknown value for mediaType
func (t *Travel) GetUnknownMediaType() (v interface{}) {
	return t.mediaType.unknown_

}

// SetUnknownMediaType sets the unknown value of mediaType
func (t *Travel) SetUnknownMediaType(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &mediaTypeTravelIntermediateType{}
	tmp.unknown_ = i
	t.mediaType = tmp

}

// IsDuration determines whether the call to GetDuration is safe
func (t *Travel) IsDuration() (ok bool) {
	return t.duration != nil && t.duration.duration != nil

}

// GetDuration returns the value safely if IsDuration returned true
func (t *Travel) GetDuration() (v time.Duration) {
	return *t.duration.duration

}

// SetDuration sets the value of duration to be of time.Duration type
func (t *Travel) SetDuration(v time.Duration) {
	t.duration = &durationTravelIntermediateType{duration: &v}

}

// IsDurationIRI determines whether the call to GetDurationIRI is safe
func (t *Travel) IsDurationIRI() (ok bool) {
	return t.duration != nil && t.duration.IRI != nil

}

// GetDurationIRI returns the value safely if IsDurationIRI returned true
func (t *Travel) GetDurationIRI() (v *url.URL) {
	return t.duration.IRI

}

// SetDurationIRI sets the value of duration to be of *url.URL type
func (t *Travel) SetDurationIRI(v *url.URL) {
	t.duration = &durationTravelIntermediateType{IRI: v}

}

// HasUnknownDuration determines whether the call to GetUnknownDuration is safe
func (t *Travel) HasUnknownDuration() (ok bool) {
	return t.duration != nil && t.duration.unknown_ != nil

}

// GetUnknownDuration returns the unknown value for duration
func (t *Travel) GetUnknownDuration() (v interface{}) {
	return t.duration.unknown_

}

// SetUnknownDuration sets the unknown value of duration
func (t *Travel) SetUnknownDuration(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &durationTravelIntermediateType{}
	tmp.unknown_ = i
	t.duration = tmp

}

// IsSource determines whether the call to GetSource is safe
func (t *Travel) IsSource() (ok bool) {
	return t.source != nil && t.source.Object != nil

}

// GetSource returns the value safely if IsSource returned true
func (t *Travel) GetSource() (v ObjectType) {
	return t.source.Object

}

// SetSource sets the value of source to be of ObjectType type
func (t *Travel) SetSource(v ObjectType) {
	t.source = &sourceTravelIntermediateType{Object: v}

}

// IsSourceIRI determines whether the call to GetSourceIRI is safe
func (t *Travel) IsSourceIRI() (ok bool) {
	return t.source != nil && t.source.IRI != nil

}

// GetSourceIRI returns the value safely if IsSourceIRI returned true
func (t *Travel) GetSourceIRI() (v *url.URL) {
	return t.source.IRI

}

// SetSourceIRI sets the value of source to be of *url.URL type
func (t *Travel) SetSourceIRI(v *url.URL) {
	t.source = &sourceTravelIntermediateType{IRI: v}

}

// HasUnknownSource determines whether the call to GetUnknownSource is safe
func (t *Travel) HasUnknownSource() (ok bool) {
	return t.source != nil && t.source.unknown_ != nil

}

// GetUnknownSource returns the unknown value for source
func (t *Travel) GetUnknownSource() (v interface{}) {
	return t.source.unknown_

}

// SetUnknownSource sets the unknown value of source
func (t *Travel) SetUnknownSource(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &sourceTravelIntermediateType{}
	tmp.unknown_ = i
	t.source = tmp

}

// IsInboxOrderedCollection determines whether the call to GetInboxOrderedCollection is safe
func (t *Travel) IsInboxOrderedCollection() (ok bool) {
	return t.inbox != nil && t.inbox.OrderedCollection != nil

}

// GetInboxOrderedCollection returns the value safely if IsInboxOrderedCollection returned true
func (t *Travel) GetInboxOrderedCollection() (v OrderedCollectionType) {
	return t.inbox.OrderedCollection

}

// SetInboxOrderedCollection sets the value of inbox to be of OrderedCollectionType type
func (t *Travel) SetInboxOrderedCollection(v OrderedCollectionType) {
	t.inbox = &inboxTravelIntermediateType{OrderedCollection: v}

}

// IsInboxAnyURI determines whether the call to GetInboxAnyURI is safe
func (t *Travel) IsInboxAnyURI() (ok bool) {
	return t.inbox != nil && t.inbox.anyURI != nil

}

// GetInboxAnyURI returns the value safely if IsInboxAnyURI returned true
func (t *Travel) GetInboxAnyURI() (v *url.URL) {
	return t.inbox.anyURI

}

// SetInboxAnyURI sets the value of inbox to be of *url.URL type
func (t *Travel) SetInboxAnyURI(v *url.URL) {
	t.inbox = &inboxTravelIntermediateType{anyURI: v}

}

// HasUnknownInbox determines whether the call to GetUnknownInbox is safe
func (t *Travel) HasUnknownInbox() (ok bool) {
	return t.inbox != nil && t.inbox.unknown_ != nil

}

// GetUnknownInbox returns the unknown value for inbox
func (t *Travel) GetUnknownInbox() (v interface{}) {
	return t.inbox.unknown_

}

// SetUnknownInbox sets the unknown value of inbox
func (t *Travel) SetUnknownInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inboxTravelIntermediateType{}
	tmp.unknown_ = i
	t.inbox = tmp

}

// IsOutboxOrderedCollection determines whether the call to GetOutboxOrderedCollection is safe
func (t *Travel) IsOutboxOrderedCollection() (ok bool) {
	return t.outbox != nil && t.outbox.OrderedCollection != nil

}

// GetOutboxOrderedCollection returns the value safely if IsOutboxOrderedCollection returned true
func (t *Travel) GetOutboxOrderedCollection() (v OrderedCollectionType) {
	return t.outbox.OrderedCollection

}

// SetOutboxOrderedCollection sets the value of outbox to be of OrderedCollectionType type
func (t *Travel) SetOutboxOrderedCollection(v OrderedCollectionType) {
	t.outbox = &outboxTravelIntermediateType{OrderedCollection: v}

}

// IsOutboxAnyURI determines whether the call to GetOutboxAnyURI is safe
func (t *Travel) IsOutboxAnyURI() (ok bool) {
	return t.outbox != nil && t.outbox.anyURI != nil

}

// GetOutboxAnyURI returns the value safely if IsOutboxAnyURI returned true
func (t *Travel) GetOutboxAnyURI() (v *url.URL) {
	return t.outbox.anyURI

}

// SetOutboxAnyURI sets the value of outbox to be of *url.URL type
func (t *Travel) SetOutboxAnyURI(v *url.URL) {
	t.outbox = &outboxTravelIntermediateType{anyURI: v}

}

// HasUnknownOutbox determines whether the call to GetUnknownOutbox is safe
func (t *Travel) HasUnknownOutbox() (ok bool) {
	return t.outbox != nil && t.outbox.unknown_ != nil

}

// GetUnknownOutbox returns the unknown value for outbox
func (t *Travel) GetUnknownOutbox() (v interface{}) {
	return t.outbox.unknown_

}

// SetUnknownOutbox sets the unknown value of outbox
func (t *Travel) SetUnknownOutbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &outboxTravelIntermediateType{}
	tmp.unknown_ = i
	t.outbox = tmp

}

// IsFollowingCollection determines whether the call to GetFollowingCollection is safe
func (t *Travel) IsFollowingCollection() (ok bool) {
	return t.following != nil && t.following.Collection != nil

}

// GetFollowingCollection returns the value safely if IsFollowingCollection returned true
func (t *Travel) GetFollowingCollection() (v CollectionType) {
	return t.following.Collection

}

// SetFollowingCollection sets the value of following to be of CollectionType type
func (t *Travel) SetFollowingCollection(v CollectionType) {
	t.following = &followingTravelIntermediateType{Collection: v}

}

// IsFollowingOrderedCollection determines whether the call to GetFollowingOrderedCollection is safe
func (t *Travel) IsFollowingOrderedCollection() (ok bool) {
	return t.following != nil && t.following.OrderedCollection != nil

}

// GetFollowingOrderedCollection returns the value safely if IsFollowingOrderedCollection returned true
func (t *Travel) GetFollowingOrderedCollection() (v OrderedCollectionType) {
	return t.following.OrderedCollection

}

// SetFollowingOrderedCollection sets the value of following to be of OrderedCollectionType type
func (t *Travel) SetFollowingOrderedCollection(v OrderedCollectionType) {
	t.following = &followingTravelIntermediateType{OrderedCollection: v}

}

// IsFollowingAnyURI determines whether the call to GetFollowingAnyURI is safe
func (t *Travel) IsFollowingAnyURI() (ok bool) {
	return t.following != nil && t.following.anyURI != nil

}

// GetFollowingAnyURI returns the value safely if IsFollowingAnyURI returned true
func (t *Travel) GetFollowingAnyURI() (v *url.URL) {
	return t.following.anyURI

}

// SetFollowingAnyURI sets the value of following to be of *url.URL type
func (t *Travel) SetFollowingAnyURI(v *url.URL) {
	t.following = &followingTravelIntermediateType{anyURI: v}

}

// HasUnknownFollowing determines whether the call to GetUnknownFollowing is safe
func (t *Travel) HasUnknownFollowing() (ok bool) {
	return t.following != nil && t.following.unknown_ != nil

}

// GetUnknownFollowing returns the unknown value for following
func (t *Travel) GetUnknownFollowing() (v interface{}) {
	return t.following.unknown_

}

// SetUnknownFollowing sets the unknown value of following
func (t *Travel) SetUnknownFollowing(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followingTravelIntermediateType{}
	tmp.unknown_ = i
	t.following = tmp

}

// IsFollowersCollection determines whether the call to GetFollowersCollection is safe
func (t *Travel) IsFollowersCollection() (ok bool) {
	return t.followers != nil && t.followers.Collection != nil

}

// GetFollowersCollection returns the value safely if IsFollowersCollection returned true
func (t *Travel) GetFollowersCollection() (v CollectionType) {
	return t.followers.Collection

}

// SetFollowersCollection sets the value of followers to be of CollectionType type
func (t *Travel) SetFollowersCollection(v CollectionType) {
	t.followers = &followersTravelIntermediateType{Collection: v}

}

// IsFollowersOrderedCollection determines whether the call to GetFollowersOrderedCollection is safe
func (t *Travel) IsFollowersOrderedCollection() (ok bool) {
	return t.followers != nil && t.followers.OrderedCollection != nil

}

// GetFollowersOrderedCollection returns the value safely if IsFollowersOrderedCollection returned true
func (t *Travel) GetFollowersOrderedCollection() (v OrderedCollectionType) {
	return t.followers.OrderedCollection

}

// SetFollowersOrderedCollection sets the value of followers to be of OrderedCollectionType type
func (t *Travel) SetFollowersOrderedCollection(v OrderedCollectionType) {
	t.followers = &followersTravelIntermediateType{OrderedCollection: v}

}

// IsFollowersAnyURI determines whether the call to GetFollowersAnyURI is safe
func (t *Travel) IsFollowersAnyURI() (ok bool) {
	return t.followers != nil && t.followers.anyURI != nil

}

// GetFollowersAnyURI returns the value safely if IsFollowersAnyURI returned true
func (t *Travel) GetFollowersAnyURI() (v *url.URL) {
	return t.followers.anyURI

}

// SetFollowersAnyURI sets the value of followers to be of *url.URL type
func (t *Travel) SetFollowersAnyURI(v *url.URL) {
	t.followers = &followersTravelIntermediateType{anyURI: v}

}

// HasUnknownFollowers determines whether the call to GetUnknownFollowers is safe
func (t *Travel) HasUnknownFollowers() (ok bool) {
	return t.followers != nil && t.followers.unknown_ != nil

}

// GetUnknownFollowers returns the unknown value for followers
func (t *Travel) GetUnknownFollowers() (v interface{}) {
	return t.followers.unknown_

}

// SetUnknownFollowers sets the unknown value of followers
func (t *Travel) SetUnknownFollowers(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followersTravelIntermediateType{}
	tmp.unknown_ = i
	t.followers = tmp

}

// IsLikedCollection determines whether the call to GetLikedCollection is safe
func (t *Travel) IsLikedCollection() (ok bool) {
	return t.liked != nil && t.liked.Collection != nil

}

// GetLikedCollection returns the value safely if IsLikedCollection returned true
func (t *Travel) GetLikedCollection() (v CollectionType) {
	return t.liked.Collection

}

// SetLikedCollection sets the value of liked to be of CollectionType type
func (t *Travel) SetLikedCollection(v CollectionType) {
	t.liked = &likedTravelIntermediateType{Collection: v}

}

// IsLikedOrderedCollection determines whether the call to GetLikedOrderedCollection is safe
func (t *Travel) IsLikedOrderedCollection() (ok bool) {
	return t.liked != nil && t.liked.OrderedCollection != nil

}

// GetLikedOrderedCollection returns the value safely if IsLikedOrderedCollection returned true
func (t *Travel) GetLikedOrderedCollection() (v OrderedCollectionType) {
	return t.liked.OrderedCollection

}

// SetLikedOrderedCollection sets the value of liked to be of OrderedCollectionType type
func (t *Travel) SetLikedOrderedCollection(v OrderedCollectionType) {
	t.liked = &likedTravelIntermediateType{OrderedCollection: v}

}

// IsLikedAnyURI determines whether the call to GetLikedAnyURI is safe
func (t *Travel) IsLikedAnyURI() (ok bool) {
	return t.liked != nil && t.liked.anyURI != nil

}

// GetLikedAnyURI returns the value safely if IsLikedAnyURI returned true
func (t *Travel) GetLikedAnyURI() (v *url.URL) {
	return t.liked.anyURI

}

// SetLikedAnyURI sets the value of liked to be of *url.URL type
func (t *Travel) SetLikedAnyURI(v *url.URL) {
	t.liked = &likedTravelIntermediateType{anyURI: v}

}

// HasUnknownLiked determines whether the call to GetUnknownLiked is safe
func (t *Travel) HasUnknownLiked() (ok bool) {
	return t.liked != nil && t.liked.unknown_ != nil

}

// GetUnknownLiked returns the unknown value for liked
func (t *Travel) GetUnknownLiked() (v interface{}) {
	return t.liked.unknown_

}

// SetUnknownLiked sets the unknown value of liked
func (t *Travel) SetUnknownLiked(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likedTravelIntermediateType{}
	tmp.unknown_ = i
	t.liked = tmp

}

// IsLikesCollection determines whether the call to GetLikesCollection is safe
func (t *Travel) IsLikesCollection() (ok bool) {
	return t.likes != nil && t.likes.Collection != nil

}

// GetLikesCollection returns the value safely if IsLikesCollection returned true
func (t *Travel) GetLikesCollection() (v CollectionType) {
	return t.likes.Collection

}

// SetLikesCollection sets the value of likes to be of CollectionType type
func (t *Travel) SetLikesCollection(v CollectionType) {
	t.likes = &likesTravelIntermediateType{Collection: v}

}

// IsLikesOrderedCollection determines whether the call to GetLikesOrderedCollection is safe
func (t *Travel) IsLikesOrderedCollection() (ok bool) {
	return t.likes != nil && t.likes.OrderedCollection != nil

}

// GetLikesOrderedCollection returns the value safely if IsLikesOrderedCollection returned true
func (t *Travel) GetLikesOrderedCollection() (v OrderedCollectionType) {
	return t.likes.OrderedCollection

}

// SetLikesOrderedCollection sets the value of likes to be of OrderedCollectionType type
func (t *Travel) SetLikesOrderedCollection(v OrderedCollectionType) {
	t.likes = &likesTravelIntermediateType{OrderedCollection: v}

}

// IsLikesAnyURI determines whether the call to GetLikesAnyURI is safe
func (t *Travel) IsLikesAnyURI() (ok bool) {
	return t.likes != nil && t.likes.anyURI != nil

}

// GetLikesAnyURI returns the value safely if IsLikesAnyURI returned true
func (t *Travel) GetLikesAnyURI() (v *url.URL) {
	return t.likes.anyURI

}

// SetLikesAnyURI sets the value of likes to be of *url.URL type
func (t *Travel) SetLikesAnyURI(v *url.URL) {
	t.likes = &likesTravelIntermediateType{anyURI: v}

}

// HasUnknownLikes determines whether the call to GetUnknownLikes is safe
func (t *Travel) HasUnknownLikes() (ok bool) {
	return t.likes != nil && t.likes.unknown_ != nil

}

// GetUnknownLikes returns the unknown value for likes
func (t *Travel) GetUnknownLikes() (v interface{}) {
	return t.likes.unknown_

}

// SetUnknownLikes sets the unknown value of likes
func (t *Travel) SetUnknownLikes(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likesTravelIntermediateType{}
	tmp.unknown_ = i
	t.likes = tmp

}

// StreamsLen determines the number of elements able to be used for the GetStreams and RemoveStreams functions
func (t *Travel) StreamsLen() (l int) {
	return len(t.streams)

}

// GetStreams returns the value for the specified index
func (t *Travel) GetStreams(index int) (v *url.URL) {
	return t.streams[index]

}

// AppendStreams adds a value to the back of streams
func (t *Travel) AppendStreams(v *url.URL) {
	t.streams = append(t.streams, v)

}

// PrependStreams adds a value to the front of streams
func (t *Travel) PrependStreams(v *url.URL) {
	t.streams = append([]*url.URL{v}, t.streams...)

}

// RemoveStreams deletes the value from the specified index
func (t *Travel) RemoveStreams(index int) {
	copy(t.streams[index:], t.streams[index+1:])
	t.streams[len(t.streams)-1] = nil
	t.streams = t.streams[:len(t.streams)-1]

}

// HasUnknownStreams determines whether the call to GetUnknownStreams is safe
func (t *Travel) HasUnknownStreams() (ok bool) {
	return t.unknown_ != nil && t.unknown_["streams"] != nil

}

// GetUnknownStreams returns the unknown value for streams
func (t *Travel) GetUnknownStreams() (v interface{}) {
	return t.unknown_["streams"]

}

// SetUnknownStreams sets the unknown value of streams
func (t *Travel) SetUnknownStreams(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["streams"] = i

}

// IsPreferredUsername determines whether the call to GetPreferredUsername is safe
func (t *Travel) IsPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.stringName != nil

}

// GetPreferredUsername returns the value safely if IsPreferredUsername returned true
func (t *Travel) GetPreferredUsername() (v string) {
	return *t.preferredUsername.stringName

}

// SetPreferredUsername sets the value of preferredUsername to be of string type
func (t *Travel) SetPreferredUsername(v string) {
	t.preferredUsername = &preferredUsernameTravelIntermediateType{stringName: &v}

}

// IsPreferredUsernameIRI determines whether the call to GetPreferredUsernameIRI is safe
func (t *Travel) IsPreferredUsernameIRI() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.IRI != nil

}

// GetPreferredUsernameIRI returns the value safely if IsPreferredUsernameIRI returned true
func (t *Travel) GetPreferredUsernameIRI() (v *url.URL) {
	return t.preferredUsername.IRI

}

// SetPreferredUsernameIRI sets the value of preferredUsername to be of *url.URL type
func (t *Travel) SetPreferredUsernameIRI(v *url.URL) {
	t.preferredUsername = &preferredUsernameTravelIntermediateType{IRI: v}

}

// HasUnknownPreferredUsername determines whether the call to GetUnknownPreferredUsername is safe
func (t *Travel) HasUnknownPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.unknown_ != nil

}

// GetUnknownPreferredUsername returns the unknown value for preferredUsername
func (t *Travel) GetUnknownPreferredUsername() (v interface{}) {
	return t.preferredUsername.unknown_

}

// SetUnknownPreferredUsername sets the unknown value of preferredUsername
func (t *Travel) SetUnknownPreferredUsername(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &preferredUsernameTravelIntermediateType{}
	tmp.unknown_ = i
	t.preferredUsername = tmp

}

// PreferredUsernameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Travel) PreferredUsernameMapLanguages() (l []string) {
	if t.preferredUsernameMap == nil || len(t.preferredUsernameMap) == 0 {
		return nil
	}
	for k := range t.preferredUsernameMap {
		l = append(l, k)
	}
	return

}

// GetPreferredUsernameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Travel) GetPreferredUsernameMap(l string) (v string) {
	if t.preferredUsernameMap == nil {
		return ""
	}
	ok := false
	v, ok = t.preferredUsernameMap[l]
	if !ok {
		return ""
	}
	return v

}

// SetPreferredUsernameMap sets the value of the property for the specified language
func (t *Travel) SetPreferredUsernameMap(l string, v string) {
	if t.preferredUsernameMap == nil {
		t.preferredUsernameMap = make(map[string]string)
	}
	t.preferredUsernameMap[l] = v

}

// IsEndpoints determines whether the call to GetEndpoints is safe
func (t *Travel) IsEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.Object != nil

}

// GetEndpoints returns the value safely if IsEndpoints returned true
func (t *Travel) GetEndpoints() (v ObjectType) {
	return t.endpoints.Object

}

// SetEndpoints sets the value of endpoints to be of ObjectType type
func (t *Travel) SetEndpoints(v ObjectType) {
	t.endpoints = &endpointsTravelIntermediateType{Object: v}

}

// IsEndpointsIRI determines whether the call to GetEndpointsIRI is safe
func (t *Travel) IsEndpointsIRI() (ok bool) {
	return t.endpoints != nil && t.endpoints.IRI != nil

}

// GetEndpointsIRI returns the value safely if IsEndpointsIRI returned true
func (t *Travel) GetEndpointsIRI() (v *url.URL) {
	return t.endpoints.IRI

}

// SetEndpointsIRI sets the value of endpoints to be of *url.URL type
func (t *Travel) SetEndpointsIRI(v *url.URL) {
	t.endpoints = &endpointsTravelIntermediateType{IRI: v}

}

// HasUnknownEndpoints determines whether the call to GetUnknownEndpoints is safe
func (t *Travel) HasUnknownEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.unknown_ != nil

}

// GetUnknownEndpoints returns the unknown value for endpoints
func (t *Travel) GetUnknownEndpoints() (v interface{}) {
	return t.endpoints.unknown_

}

// SetUnknownEndpoints sets the unknown value of endpoints
func (t *Travel) SetUnknownEndpoints(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endpointsTravelIntermediateType{}
	tmp.unknown_ = i
	t.endpoints = tmp

}

// HasProxyUrl determines whether the call to GetProxyUrl is safe
func (t *Travel) HasProxyUrl() (ok bool) {
	return t.proxyUrl != nil

}

// GetProxyUrl returns the value for proxyUrl
func (t *Travel) GetProxyUrl() (v *url.URL) {
	return t.proxyUrl

}

// SetProxyUrl sets the value of proxyUrl
func (t *Travel) SetProxyUrl(v *url.URL) {
	t.proxyUrl = v

}

// HasUnknownProxyUrl determines whether the call to GetUnknownProxyUrl is safe
func (t *Travel) HasUnknownProxyUrl() (ok bool) {
	return t.unknown_ != nil && t.unknown_["proxyUrl"] != nil

}

// GetUnknownProxyUrl returns the unknown value for proxyUrl
func (t *Travel) GetUnknownProxyUrl() (v interface{}) {
	return t.unknown_["proxyUrl"]

}

// SetUnknownProxyUrl sets the unknown value of proxyUrl
func (t *Travel) SetUnknownProxyUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["proxyUrl"] = i

}

// HasOauthAuthorizationEndpoint determines whether the call to GetOauthAuthorizationEndpoint is safe
func (t *Travel) HasOauthAuthorizationEndpoint() (ok bool) {
	return t.oauthAuthorizationEndpoint != nil

}

// GetOauthAuthorizationEndpoint returns the value for oauthAuthorizationEndpoint
func (t *Travel) GetOauthAuthorizationEndpoint() (v *url.URL) {
	return t.oauthAuthorizationEndpoint

}

// SetOauthAuthorizationEndpoint sets the value of oauthAuthorizationEndpoint
func (t *Travel) SetOauthAuthorizationEndpoint(v *url.URL) {
	t.oauthAuthorizationEndpoint = v

}

// HasUnknownOauthAuthorizationEndpoint determines whether the call to GetUnknownOauthAuthorizationEndpoint is safe
func (t *Travel) HasUnknownOauthAuthorizationEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthAuthorizationEndpoint"] != nil

}

// GetUnknownOauthAuthorizationEndpoint returns the unknown value for oauthAuthorizationEndpoint
func (t *Travel) GetUnknownOauthAuthorizationEndpoint() (v interface{}) {
	return t.unknown_["oauthAuthorizationEndpoint"]

}

// SetUnknownOauthAuthorizationEndpoint sets the unknown value of oauthAuthorizationEndpoint
func (t *Travel) SetUnknownOauthAuthorizationEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthAuthorizationEndpoint"] = i

}

// HasOauthTokenEndpoint determines whether the call to GetOauthTokenEndpoint is safe
func (t *Travel) HasOauthTokenEndpoint() (ok bool) {
	return t.oauthTokenEndpoint != nil

}

// GetOauthTokenEndpoint returns the value for oauthTokenEndpoint
func (t *Travel) GetOauthTokenEndpoint() (v *url.URL) {
	return t.oauthTokenEndpoint

}

// SetOauthTokenEndpoint sets the value of oauthTokenEndpoint
func (t *Travel) SetOauthTokenEndpoint(v *url.URL) {
	t.oauthTokenEndpoint = v

}

// HasUnknownOauthTokenEndpoint determines whether the call to GetUnknownOauthTokenEndpoint is safe
func (t *Travel) HasUnknownOauthTokenEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthTokenEndpoint"] != nil

}

// GetUnknownOauthTokenEndpoint returns the unknown value for oauthTokenEndpoint
func (t *Travel) GetUnknownOauthTokenEndpoint() (v interface{}) {
	return t.unknown_["oauthTokenEndpoint"]

}

// SetUnknownOauthTokenEndpoint sets the unknown value of oauthTokenEndpoint
func (t *Travel) SetUnknownOauthTokenEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthTokenEndpoint"] = i

}

// HasProvideClientKey determines whether the call to GetProvideClientKey is safe
func (t *Travel) HasProvideClientKey() (ok bool) {
	return t.provideClientKey != nil

}

// GetProvideClientKey returns the value for provideClientKey
func (t *Travel) GetProvideClientKey() (v *url.URL) {
	return t.provideClientKey

}

// SetProvideClientKey sets the value of provideClientKey
func (t *Travel) SetProvideClientKey(v *url.URL) {
	t.provideClientKey = v

}

// HasUnknownProvideClientKey determines whether the call to GetUnknownProvideClientKey is safe
func (t *Travel) HasUnknownProvideClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["provideClientKey"] != nil

}

// GetUnknownProvideClientKey returns the unknown value for provideClientKey
func (t *Travel) GetUnknownProvideClientKey() (v interface{}) {
	return t.unknown_["provideClientKey"]

}

// SetUnknownProvideClientKey sets the unknown value of provideClientKey
func (t *Travel) SetUnknownProvideClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["provideClientKey"] = i

}

// HasSignClientKey determines whether the call to GetSignClientKey is safe
func (t *Travel) HasSignClientKey() (ok bool) {
	return t.signClientKey != nil

}

// GetSignClientKey returns the value for signClientKey
func (t *Travel) GetSignClientKey() (v *url.URL) {
	return t.signClientKey

}

// SetSignClientKey sets the value of signClientKey
func (t *Travel) SetSignClientKey(v *url.URL) {
	t.signClientKey = v

}

// HasUnknownSignClientKey determines whether the call to GetUnknownSignClientKey is safe
func (t *Travel) HasUnknownSignClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["signClientKey"] != nil

}

// GetUnknownSignClientKey returns the unknown value for signClientKey
func (t *Travel) GetUnknownSignClientKey() (v interface{}) {
	return t.unknown_["signClientKey"]

}

// SetUnknownSignClientKey sets the unknown value of signClientKey
func (t *Travel) SetUnknownSignClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["signClientKey"] = i

}

// HasSharedInbox determines whether the call to GetSharedInbox is safe
func (t *Travel) HasSharedInbox() (ok bool) {
	return t.sharedInbox != nil

}

// GetSharedInbox returns the value for sharedInbox
func (t *Travel) GetSharedInbox() (v *url.URL) {
	return t.sharedInbox

}

// SetSharedInbox sets the value of sharedInbox
func (t *Travel) SetSharedInbox(v *url.URL) {
	t.sharedInbox = v

}

// HasUnknownSharedInbox determines whether the call to GetUnknownSharedInbox is safe
func (t *Travel) HasUnknownSharedInbox() (ok bool) {
	return t.unknown_ != nil && t.unknown_["sharedInbox"] != nil

}

// GetUnknownSharedInbox returns the unknown value for sharedInbox
func (t *Travel) GetUnknownSharedInbox() (v interface{}) {
	return t.unknown_["sharedInbox"]

}

// SetUnknownSharedInbox sets the unknown value of sharedInbox
func (t *Travel) SetUnknownSharedInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["sharedInbox"] = i

}

// AddUnknown adds a raw extension to this object with the specified key
func (t *Travel) AddUnknown(k string, i interface{}) (this *Travel) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_[k] = i
	return t

}

// HasUnknown returns true if there is an unknown object with the specified key
func (t *Travel) HasUnknown(k string) (b bool) {
	if t.unknown_ == nil {
		return false
	}
	_, ok := t.unknown_[k]
	return ok

}

// RemoveUnknown removes a raw extension from this object with the specified key
func (t *Travel) RemoveUnknown(k string) (this *Travel) {
	delete(t.unknown_, k)
	return t

}

// Serialize turns this object into a map[string]interface{}. Note that the "type" property will automatically be populated with "Travel" if not manually set by the caller
func (t *Travel) Serialize() (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	for k, v := range t.unknown_ {
		m[k] = unknownValueSerialize(v)
	}
	var typeAlreadySet bool
	for _, k := range t.typeName {
		if ks, ok := k.(string); ok {
			if ks == "Travel" {
				typeAlreadySet = true
				break
			}
		}
	}
	if !typeAlreadySet {
		t.typeName = append(t.typeName, "Travel")
	}
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceActorTravelIntermediateType(t.actor); err == nil && v != nil {
		if len(v) == 1 {
			m["actor"] = v[0]
		} else {
			m["actor"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceTargetTravelIntermediateType(t.target); err == nil && v != nil {
		if len(v) == 1 {
			m["target"] = v[0]
		} else {
			m["target"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceResultTravelIntermediateType(t.result); err == nil && v != nil {
		if len(v) == 1 {
			m["result"] = v[0]
		} else {
			m["result"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceOriginTravelIntermediateType(t.origin); err == nil && v != nil {
		if len(v) == 1 {
			m["origin"] = v[0]
		} else {
			m["origin"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceInstrumentTravelIntermediateType(t.instrument); err == nil && v != nil {
		if len(v) == 1 {
			m["instrument"] = v[0]
		} else {
			m["instrument"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.altitude != nil {
		if v, err := serializeAltitudeTravelIntermediateType(t.altitude); err == nil {
			m["altitude"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAttachmentTravelIntermediateType(t.attachment); err == nil && v != nil {
		if len(v) == 1 {
			m["attachment"] = v[0]
		} else {
			m["attachment"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAttributedToTravelIntermediateType(t.attributedTo); err == nil && v != nil {
		if len(v) == 1 {
			m["attributedTo"] = v[0]
		} else {
			m["attributedTo"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAudienceTravelIntermediateType(t.audience); err == nil && v != nil {
		if len(v) == 1 {
			m["audience"] = v[0]
		} else {
			m["audience"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceContentTravelIntermediateType(t.content); err == nil && v != nil {
		if len(v) == 1 {
			m["content"] = v[0]
		} else {
			m["content"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition

	// Begin generation by generateNaturalLanguageMap
	if t.contentMap != nil && len(t.contentMap) >= 0 {
		m["contentMap"] = t.contentMap
	}
	// End generation by generateNaturalLanguageMap
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceContextTravelIntermediateType(t.context); err == nil && v != nil {
		if len(v) == 1 {
			m["context"] = v[0]
		} else {
			m["context"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceNameTravelIntermediateType(t.name); err == nil && v != nil {
		if len(v) == 1 {
			m["name"] = v[0]
		} else {
			m["name"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition

	// Begin generation by generateNaturalLanguageMap
	if t.nameMap != nil && len(t.nameMap) >= 0 {
		m["nameMap"] = t.nameMap
	}
	// End generation by generateNaturalLanguageMap
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.endTime != nil {
		if v, err := serializeEndTimeTravelIntermediateType(t.endTime); err == nil {
			m["endTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceGeneratorTravelIntermediateType(t.generator); err == nil && v != nil {
		if len(v) == 1 {
			m["generator"] = v[0]
		} else {
			m["generator"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceIconTravelIntermediateType(t.icon); err == nil && v != nil {
		if len(v) == 1 {
			m["icon"] = v[0]
		} else {
			m["icon"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by RangeReference.Serialize for Value
	if t.id != nil {
		idSerializeFunc := func() (interface{}, error) {
			v := t.id
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		idResult, err := idSerializeFunc()
		if err == nil {
			m["id"] = idResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceImageTravelIntermediateType(t.image); err == nil && v != nil {
		if len(v) == 1 {
			m["image"] = v[0]
		} else {
			m["image"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceInReplyToTravelIntermediateType(t.inReplyTo); err == nil && v != nil {
		if len(v) == 1 {
			m["inReplyTo"] = v[0]
		} else {
			m["inReplyTo"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceLocationTravelIntermediateType(t.location); err == nil && v != nil {
		if len(v) == 1 {
			m["location"] = v[0]
		} else {
			m["location"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSlicePreviewTravelIntermediateType(t.preview); err == nil && v != nil {
		if len(v) == 1 {
			m["preview"] = v[0]
		} else {
			m["preview"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.published != nil {
		if v, err := serializePublishedTravelIntermediateType(t.published); err == nil {
			m["published"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.replies != nil {
		if v, err := serializeRepliesTravelIntermediateType(t.replies); err == nil {
			m["replies"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.startTime != nil {
		if v, err := serializeStartTimeTravelIntermediateType(t.startTime); err == nil {
			m["startTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceSummaryTravelIntermediateType(t.summary); err == nil && v != nil {
		if len(v) == 1 {
			m["summary"] = v[0]
		} else {
			m["summary"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition

	// Begin generation by generateNaturalLanguageMap
	if t.summaryMap != nil && len(t.summaryMap) >= 0 {
		m["summaryMap"] = t.summaryMap
	}
	// End generation by generateNaturalLanguageMap
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceTagTravelIntermediateType(t.tag); err == nil && v != nil {
		if len(v) == 1 {
			m["tag"] = v[0]
		} else {
			m["tag"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalAnyDefinition
	if t.typeName != nil {
		if len(t.typeName) == 1 {
			m["type"] = t.typeName[0]
		} else {
			m["type"] = t.typeName
		}
	}
	// End generation by generateNonFunctionalAnyDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.updated != nil {
		if v, err := serializeUpdatedTravelIntermediateType(t.updated); err == nil {
			m["updated"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceUrlTravelIntermediateType(t.url); err == nil && v != nil {
		if len(v) == 1 {
			m["url"] = v[0]
		} else {
			m["url"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceToTravelIntermediateType(t.to); err == nil && v != nil {
		if len(v) == 1 {
			m["to"] = v[0]
		} else {
			m["to"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceBtoTravelIntermediateType(t.bto); err == nil && v != nil {
		if len(v) == 1 {
			m["bto"] = v[0]
		} else {
			m["bto"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceCcTravelIntermediateType(t.cc); err == nil && v != nil {
		if len(v) == 1 {
			m["cc"] = v[0]
		} else {
			m["cc"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceBccTravelIntermediateType(t.bcc); err == nil && v != nil {
		if len(v) == 1 {
			m["bcc"] = v[0]
		} else {
			m["bcc"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.mediaType != nil {
		if v, err := serializeMediaTypeTravelIntermediateType(t.mediaType); err == nil {
			m["mediaType"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.duration != nil {
		if v, err := serializeDurationTravelIntermediateType(t.duration); err == nil {
			m["duration"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.source != nil {
		if v, err := serializeSourceTravelIntermediateType(t.source); err == nil {
			m["source"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.inbox != nil {
		if v, err := serializeInboxTravelIntermediateType(t.inbox); err == nil {
			m["inbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.outbox != nil {
		if v, err := serializeOutboxTravelIntermediateType(t.outbox); err == nil {
			m["outbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.following != nil {
		if v, err := serializeFollowingTravelIntermediateType(t.following); err == nil {
			m["following"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.followers != nil {
		if v, err := serializeFollowersTravelIntermediateType(t.followers); err == nil {
			m["followers"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.liked != nil {
		if v, err := serializeLikedTravelIntermediateType(t.liked); err == nil {
			m["liked"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.likes != nil {
		if v, err := serializeLikesTravelIntermediateType(t.likes); err == nil {
			m["likes"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by RangeReference.Serialize for Value
	var streamsTemp []interface{}
	for _, v := range t.streams {
		tmp := anyURISerialize(v)
		streamsTemp = append(streamsTemp, tmp)
	}
	if streamsTemp != nil {
		if len(streamsTemp) == 1 {
			m["streams"] = streamsTemp[0]
		} else {
			m["streams"] = streamsTemp
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.preferredUsername != nil {
		if v, err := serializePreferredUsernameTravelIntermediateType(t.preferredUsername); err == nil {
			m["preferredUsername"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition

	// Begin generation by generateNaturalLanguageMap
	if t.preferredUsernameMap != nil && len(t.preferredUsernameMap) >= 0 {
		m["preferredUsernameMap"] = t.preferredUsernameMap
	}
	// End generation by generateNaturalLanguageMap
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.endpoints != nil {
		if v, err := serializeEndpointsTravelIntermediateType(t.endpoints); err == nil {
			m["endpoints"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by RangeReference.Serialize for Value
	if t.proxyUrl != nil {
		proxyUrlSerializeFunc := func() (interface{}, error) {
			v := t.proxyUrl
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		proxyUrlResult, err := proxyUrlSerializeFunc()
		if err == nil {
			m["proxyUrl"] = proxyUrlResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by RangeReference.Serialize for Value
	if t.oauthAuthorizationEndpoint != nil {
		oauthAuthorizationEndpointSerializeFunc := func() (interface{}, error) {
			v := t.oauthAuthorizationEndpoint
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		oauthAuthorizationEndpointResult, err := oauthAuthorizationEndpointSerializeFunc()
		if err == nil {
			m["oauthAuthorizationEndpoint"] = oauthAuthorizationEndpointResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by RangeReference.Serialize for Value
	if t.oauthTokenEndpoint != nil {
		oauthTokenEndpointSerializeFunc := func() (interface{}, error) {
			v := t.oauthTokenEndpoint
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		oauthTokenEndpointResult, err := oauthTokenEndpointSerializeFunc()
		if err == nil {
			m["oauthTokenEndpoint"] = oauthTokenEndpointResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by RangeReference.Serialize for Value
	if t.provideClientKey != nil {
		provideClientKeySerializeFunc := func() (interface{}, error) {
			v := t.provideClientKey
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		provideClientKeyResult, err := provideClientKeySerializeFunc()
		if err == nil {
			m["provideClientKey"] = provideClientKeyResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by RangeReference.Serialize for Value
	if t.signClientKey != nil {
		signClientKeySerializeFunc := func() (interface{}, error) {
			v := t.signClientKey
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		signClientKeyResult, err := signClientKeySerializeFunc()
		if err == nil {
			m["signClientKey"] = signClientKeyResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by RangeReference.Serialize for Value
	if t.sharedInbox != nil {
		sharedInboxSerializeFunc := func() (interface{}, error) {
			v := t.sharedInbox
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		sharedInboxResult, err := sharedInboxSerializeFunc()
		if err == nil {
			m["sharedInbox"] = sharedInboxResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	return

}

// Deserialize populates this object from a map[string]interface{}
func (t *Travel) Deserialize(m map[string]interface{}) (err error) {
	for k, v := range m {
		handled := false
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "actor" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeActorTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.actor = []*actorTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.actor, err = deserializeSliceActorTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &actorTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.actor = []*actorTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "target" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeTargetTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.target = []*targetTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.target, err = deserializeSliceTargetTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &targetTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.target = []*targetTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "result" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeResultTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.result = []*resultTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.result, err = deserializeSliceResultTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &resultTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.result = []*resultTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "origin" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeOriginTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.origin = []*originTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.origin, err = deserializeSliceOriginTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &originTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.origin = []*originTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "instrument" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeInstrumentTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.instrument = []*instrumentTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.instrument, err = deserializeSliceInstrumentTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &instrumentTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.instrument = []*instrumentTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "altitude" {
				t.altitude, err = deserializeAltitudeTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "attachment" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAttachmentTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attachment, err = deserializeSliceAttachmentTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attachmentTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "attributedTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAttributedToTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attributedTo, err = deserializeSliceAttributedToTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attributedToTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "audience" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAudienceTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.audience = []*audienceTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.audience, err = deserializeSliceAudienceTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &audienceTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.audience = []*audienceTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "content" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeContentTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.content = []*contentTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.content, err = deserializeSliceContentTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contentTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.content = []*contentTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition

			// Begin generation by generateNaturalLanguageMap
			if k == "contentMap" {
				if vMap, ok := v.(map[string]interface{}); ok {
					val := make(map[string]string)
					for k, iVal := range vMap {
						if sVal, ok := iVal.(string); ok {
							val[k] = sVal
						}
					}
					t.contentMap = val
					handled = true
				}
			}
			// End generation by generateNaturalLanguageMap
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "context" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeContextTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.context = []*contextTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.context, err = deserializeSliceContextTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contextTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.context = []*contextTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "name" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeNameTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.name = []*nameTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.name, err = deserializeSliceNameTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &nameTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.name = []*nameTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition

			// Begin generation by generateNaturalLanguageMap
			if k == "nameMap" {
				if vMap, ok := v.(map[string]interface{}); ok {
					val := make(map[string]string)
					for k, iVal := range vMap {
						if sVal, ok := iVal.(string); ok {
							val[k] = sVal
						}
					}
					t.nameMap = val
					handled = true
				}
			}
			// End generation by generateNaturalLanguageMap
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "endTime" {
				t.endTime, err = deserializeEndTimeTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "generator" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeGeneratorTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.generator = []*generatorTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.generator, err = deserializeSliceGeneratorTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &generatorTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.generator = []*generatorTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "icon" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeIconTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.icon = []*iconTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.icon, err = deserializeSliceIconTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &iconTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.icon = []*iconTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "id" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.id = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "image" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeImageTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.image = []*imageTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.image, err = deserializeSliceImageTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &imageTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.image = []*imageTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "inReplyTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeInReplyToTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.inReplyTo, err = deserializeSliceInReplyToTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &inReplyToTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "location" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeLocationTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.location = []*locationTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.location, err = deserializeSliceLocationTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &locationTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.location = []*locationTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "preview" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializePreviewTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.preview = []*previewTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.preview, err = deserializeSlicePreviewTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &previewTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.preview = []*previewTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "published" {
				t.published, err = deserializePublishedTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "replies" {
				t.replies, err = deserializeRepliesTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "startTime" {
				t.startTime, err = deserializeStartTimeTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "summary" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeSummaryTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.summary = []*summaryTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.summary, err = deserializeSliceSummaryTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &summaryTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.summary = []*summaryTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition

			// Begin generation by generateNaturalLanguageMap
			if k == "summaryMap" {
				if vMap, ok := v.(map[string]interface{}); ok {
					val := make(map[string]string)
					for k, iVal := range vMap {
						if sVal, ok := iVal.(string); ok {
							val[k] = sVal
						}
					}
					t.summaryMap = val
					handled = true
				}
			}
			// End generation by generateNaturalLanguageMap
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "tag" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeTagTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.tag = []*tagTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.tag, err = deserializeSliceTagTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &tagTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.tag = []*tagTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalAnyDefinition
			if k == "type" {
				if tmpTypeSlice, ok := v.([]interface{}); ok {
					t.typeName = tmpTypeSlice
					handled = true
				} else {
					t.typeName = []interface{}{v}
					handled = true
				}
			}
			// End generation by generateNonFunctionalAnyDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "updated" {
				t.updated, err = deserializeUpdatedTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "url" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeUrlTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.url = []*urlTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.url, err = deserializeSliceUrlTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &urlTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.url = []*urlTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "to" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeToTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.to = []*toTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.to, err = deserializeSliceToTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &toTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.to = []*toTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bto" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBtoTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bto = []*btoTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bto, err = deserializeSliceBtoTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &btoTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bto = []*btoTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "cc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeCcTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.cc = []*ccTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.cc, err = deserializeSliceCcTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &ccTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.cc = []*ccTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bcc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBccTravelIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bcc = []*bccTravelIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bcc, err = deserializeSliceBccTravelIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &bccTravelIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bcc = []*bccTravelIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "mediaType" {
				t.mediaType, err = deserializeMediaTypeTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "duration" {
				t.duration, err = deserializeDurationTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "source" {
				t.source, err = deserializeSourceTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "inbox" {
				t.inbox, err = deserializeInboxTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "outbox" {
				t.outbox, err = deserializeOutboxTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "following" {
				t.following, err = deserializeFollowingTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "followers" {
				t.followers, err = deserializeFollowersTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "liked" {
				t.liked, err = deserializeLikedTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "likes" {
				t.likes, err = deserializeLikesTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "streams" {
				if tmpSlice, ok := v.([]interface{}); ok {
					for _, tmpElem := range tmpSlice {
						if v, ok := tmpElem.(interface{}); ok {
							tmp, err := anyURIDeserialize(v)
							if err != nil {
								return err
							}
							t.streams = append(t.streams, tmp)
							handled = true
						}
					}
				} else if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.streams = append(t.streams, tmp)
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "preferredUsername" {
				t.preferredUsername, err = deserializePreferredUsernameTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition

			// Begin generation by generateNaturalLanguageMap
			if k == "preferredUsernameMap" {
				if vMap, ok := v.(map[string]interface{}); ok {
					val := make(map[string]string)
					for k, iVal := range vMap {
						if sVal, ok := iVal.(string); ok {
							val[k] = sVal
						}
					}
					t.preferredUsernameMap = val
					handled = true
				}
			}
			// End generation by generateNaturalLanguageMap
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "endpoints" {
				t.endpoints, err = deserializeEndpointsTravelIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "proxyUrl" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.proxyUrl = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "oauthAuthorizationEndpoint" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.oauthAuthorizationEndpoint = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "oauthTokenEndpoint" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.oauthTokenEndpoint = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "provideClientKey" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.provideClientKey = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "signClientKey" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.signClientKey = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "sharedInbox" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.sharedInbox = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled && k != "@context" {
			if t.unknown_ == nil {
				t.unknown_ = make(map[string]interface{})
			}
			t.unknown_[k] = unknownValueDeserialize(v)
		}
	}
	return

}

// actorTravelIntermediateType will only have one of its values set at most
type actorTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for actor property
	Object ObjectType
	// Stores possible LinkType type for actor property
	Link LinkType
	// Stores possible *url.URL type for actor property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *actorTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *actorTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// targetTravelIntermediateType will only have one of its values set at most
type targetTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for target property
	Object ObjectType
	// Stores possible LinkType type for target property
	Link LinkType
	// Stores possible *url.URL type for target property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *targetTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *targetTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// resultTravelIntermediateType will only have one of its values set at most
type resultTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for result property
	Object ObjectType
	// Stores possible LinkType type for result property
	Link LinkType
	// Stores possible *url.URL type for result property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *resultTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *resultTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// originTravelIntermediateType will only have one of its values set at most
type originTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for origin property
	Object ObjectType
	// Stores possible LinkType type for origin property
	Link LinkType
	// Stores possible *url.URL type for origin property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *originTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *originTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// instrumentTravelIntermediateType will only have one of its values set at most
type instrumentTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for instrument property
	Object ObjectType
	// Stores possible LinkType type for instrument property
	Link LinkType
	// Stores possible *url.URL type for instrument property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *instrumentTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *instrumentTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// altitudeTravelIntermediateType will only have one of its values set at most
type altitudeTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *float64 type for altitude property
	float *float64
	// Stores possible *url.URL type for altitude property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *altitudeTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.float, err = floatDeserialize(i)
			if err != nil {
				t.float = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *altitudeTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.float != nil {
		i = floatSerialize(*t.float)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// attachmentTravelIntermediateType will only have one of its values set at most
type attachmentTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for attachment property
	Object ObjectType
	// Stores possible LinkType type for attachment property
	Link LinkType
	// Stores possible *url.URL type for attachment property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *attachmentTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *attachmentTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// attributedToTravelIntermediateType will only have one of its values set at most
type attributedToTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for attributedTo property
	Object ObjectType
	// Stores possible LinkType type for attributedTo property
	Link LinkType
	// Stores possible *url.URL type for attributedTo property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *attributedToTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *attributedToTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// audienceTravelIntermediateType will only have one of its values set at most
type audienceTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for audience property
	Object ObjectType
	// Stores possible LinkType type for audience property
	Link LinkType
	// Stores possible *url.URL type for audience property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *audienceTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *audienceTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// contentTravelIntermediateType will only have one of its values set at most
type contentTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for content property
	stringName *string
	// Stores possible *string type for content property
	langString *string
	// Stores possible *url.URL type for content property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *contentTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.langString, err = langStringDeserialize(i)
			if err != nil {
				t.langString = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *contentTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.langString != nil {
		i = langStringSerialize(*t.langString)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// contextTravelIntermediateType will only have one of its values set at most
type contextTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for context property
	Object ObjectType
	// Stores possible LinkType type for context property
	Link LinkType
	// Stores possible *url.URL type for context property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *contextTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *contextTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// nameTravelIntermediateType will only have one of its values set at most
type nameTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for name property
	stringName *string
	// Stores possible *string type for name property
	langString *string
	// Stores possible *url.URL type for name property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *nameTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.langString, err = langStringDeserialize(i)
			if err != nil {
				t.langString = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *nameTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.langString != nil {
		i = langStringSerialize(*t.langString)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// endTimeTravelIntermediateType will only have one of its values set at most
type endTimeTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for endTime property
	dateTime *time.Time
	// Stores possible *url.URL type for endTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endTimeTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *endTimeTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// generatorTravelIntermediateType will only have one of its values set at most
type generatorTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for generator property
	Object ObjectType
	// Stores possible LinkType type for generator property
	Link LinkType
	// Stores possible *url.URL type for generator property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *generatorTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *generatorTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// iconTravelIntermediateType will only have one of its values set at most
type iconTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ImageType type for icon property
	Image ImageType
	// Stores possible LinkType type for icon property
	Link LinkType
	// Stores possible *url.URL type for icon property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *iconTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Image, ok = resolveObject(kind).(ImageType); t.Image != nil && ok {
						err = t.Image.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *iconTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Image != nil {
		i, err = t.Image.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// imageTravelIntermediateType will only have one of its values set at most
type imageTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ImageType type for image property
	Image ImageType
	// Stores possible LinkType type for image property
	Link LinkType
	// Stores possible *url.URL type for image property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *imageTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Image, ok = resolveObject(kind).(ImageType); t.Image != nil && ok {
						err = t.Image.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *imageTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Image != nil {
		i, err = t.Image.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// inReplyToTravelIntermediateType will only have one of its values set at most
type inReplyToTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for inReplyTo property
	Object ObjectType
	// Stores possible LinkType type for inReplyTo property
	Link LinkType
	// Stores possible *url.URL type for inReplyTo property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *inReplyToTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *inReplyToTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// locationTravelIntermediateType will only have one of its values set at most
type locationTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for location property
	Object ObjectType
	// Stores possible LinkType type for location property
	Link LinkType
	// Stores possible *url.URL type for location property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *locationTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *locationTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// previewTravelIntermediateType will only have one of its values set at most
type previewTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for preview property
	Object ObjectType
	// Stores possible LinkType type for preview property
	Link LinkType
	// Stores possible *url.URL type for preview property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *previewTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *previewTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// publishedTravelIntermediateType will only have one of its values set at most
type publishedTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for published property
	dateTime *time.Time
	// Stores possible *url.URL type for published property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *publishedTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *publishedTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// repliesTravelIntermediateType will only have one of its values set at most
type repliesTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for replies property
	Collection CollectionType
	// Stores possible *url.URL type for replies property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *repliesTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *repliesTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// startTimeTravelIntermediateType will only have one of its values set at most
type startTimeTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for startTime property
	dateTime *time.Time
	// Stores possible *url.URL type for startTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *startTimeTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *startTimeTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// summaryTravelIntermediateType will only have one of its values set at most
type summaryTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for summary property
	stringName *string
	// Stores possible *string type for summary property
	langString *string
	// Stores possible *url.URL type for summary property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *summaryTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.langString, err = langStringDeserialize(i)
			if err != nil {
				t.langString = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *summaryTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.langString != nil {
		i = langStringSerialize(*t.langString)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// tagTravelIntermediateType will only have one of its values set at most
type tagTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for tag property
	Object ObjectType
	// Stores possible LinkType type for tag property
	Link LinkType
	// Stores possible *url.URL type for tag property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *tagTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *tagTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// updatedTravelIntermediateType will only have one of its values set at most
type updatedTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for updated property
	dateTime *time.Time
	// Stores possible *url.URL type for updated property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *updatedTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *updatedTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// urlTravelIntermediateType will only have one of its values set at most
type urlTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *url.URL type for url property
	anyURI *url.URL
	// Stores possible LinkType type for url property
	Link LinkType
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *urlTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *urlTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// toTravelIntermediateType will only have one of its values set at most
type toTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for to property
	Object ObjectType
	// Stores possible LinkType type for to property
	Link LinkType
	// Stores possible *url.URL type for to property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *toTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *toTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// btoTravelIntermediateType will only have one of its values set at most
type btoTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for bto property
	Object ObjectType
	// Stores possible LinkType type for bto property
	Link LinkType
	// Stores possible *url.URL type for bto property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *btoTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *btoTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// ccTravelIntermediateType will only have one of its values set at most
type ccTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for cc property
	Object ObjectType
	// Stores possible LinkType type for cc property
	Link LinkType
	// Stores possible *url.URL type for cc property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *ccTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *ccTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// bccTravelIntermediateType will only have one of its values set at most
type bccTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for bcc property
	Object ObjectType
	// Stores possible LinkType type for bcc property
	Link LinkType
	// Stores possible *url.URL type for bcc property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *bccTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *bccTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// mediaTypeTravelIntermediateType will only have one of its values set at most
type mediaTypeTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for mediaType property
	mimeMediaTypeValue *string
	// Stores possible *url.URL type for mediaType property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *mediaTypeTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.mimeMediaTypeValue, err = mimeMediaTypeValueDeserialize(i)
			if err != nil {
				t.mimeMediaTypeValue = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *mediaTypeTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.mimeMediaTypeValue != nil {
		i = mimeMediaTypeValueSerialize(*t.mimeMediaTypeValue)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// durationTravelIntermediateType will only have one of its values set at most
type durationTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Duration type for duration property
	duration *time.Duration
	// Stores possible *url.URL type for duration property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *durationTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.duration, err = durationDeserialize(i)
			if err != nil {
				t.duration = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *durationTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.duration != nil {
		i = durationSerialize(*t.duration)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// sourceTravelIntermediateType will only have one of its values set at most
type sourceTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for source property
	Object ObjectType
	// Stores possible *url.URL type for source property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *sourceTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *sourceTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// inboxTravelIntermediateType will only have one of its values set at most
type inboxTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for inbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for inbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *inboxTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *inboxTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// outboxTravelIntermediateType will only have one of its values set at most
type outboxTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for outbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for outbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *outboxTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *outboxTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// followingTravelIntermediateType will only have one of its values set at most
type followingTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for following property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for following property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for following property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *followingTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *followingTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// followersTravelIntermediateType will only have one of its values set at most
type followersTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for followers property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for followers property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for followers property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *followersTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *followersTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// likedTravelIntermediateType will only have one of its values set at most
type likedTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for liked property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for liked property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for liked property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *likedTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *likedTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// likesTravelIntermediateType will only have one of its values set at most
type likesTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for likes property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for likes property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for likes property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *likesTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *likesTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// preferredUsernameTravelIntermediateType will only have one of its values set at most
type preferredUsernameTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for preferredUsername property
	stringName *string
	// Stores possible *url.URL type for preferredUsername property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *preferredUsernameTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *preferredUsernameTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// endpointsTravelIntermediateType will only have one of its values set at most
type endpointsTravelIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for endpoints property
	Object ObjectType
	// Stores possible *url.URL type for endpoints property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endpointsTravelIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *endpointsTravelIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// deserializeactorTravelIntermediateType will accept a map to create a actorTravelIntermediateType
func deserializeActorTravelIntermediateType(in interface{}) (t *actorTravelIntermediateType, err error) {
	tmp := &actorTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice actorTravelIntermediateType will accept a slice to create a slice of actorTravelIntermediateType
func deserializeSliceActorTravelIntermediateType(in []interface{}) (t []*actorTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &actorTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeactorTravelIntermediateType will accept a actorTravelIntermediateType to create a map
func serializeActorTravelIntermediateType(t *actorTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceactorTravelIntermediateType will accept a slice of actorTravelIntermediateType to create a slice result
func serializeSliceActorTravelIntermediateType(s []*actorTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetargetTravelIntermediateType will accept a map to create a targetTravelIntermediateType
func deserializeTargetTravelIntermediateType(in interface{}) (t *targetTravelIntermediateType, err error) {
	tmp := &targetTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice targetTravelIntermediateType will accept a slice to create a slice of targetTravelIntermediateType
func deserializeSliceTargetTravelIntermediateType(in []interface{}) (t []*targetTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &targetTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetargetTravelIntermediateType will accept a targetTravelIntermediateType to create a map
func serializeTargetTravelIntermediateType(t *targetTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetargetTravelIntermediateType will accept a slice of targetTravelIntermediateType to create a slice result
func serializeSliceTargetTravelIntermediateType(s []*targetTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeresultTravelIntermediateType will accept a map to create a resultTravelIntermediateType
func deserializeResultTravelIntermediateType(in interface{}) (t *resultTravelIntermediateType, err error) {
	tmp := &resultTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice resultTravelIntermediateType will accept a slice to create a slice of resultTravelIntermediateType
func deserializeSliceResultTravelIntermediateType(in []interface{}) (t []*resultTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &resultTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeresultTravelIntermediateType will accept a resultTravelIntermediateType to create a map
func serializeResultTravelIntermediateType(t *resultTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceresultTravelIntermediateType will accept a slice of resultTravelIntermediateType to create a slice result
func serializeSliceResultTravelIntermediateType(s []*resultTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeoriginTravelIntermediateType will accept a map to create a originTravelIntermediateType
func deserializeOriginTravelIntermediateType(in interface{}) (t *originTravelIntermediateType, err error) {
	tmp := &originTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice originTravelIntermediateType will accept a slice to create a slice of originTravelIntermediateType
func deserializeSliceOriginTravelIntermediateType(in []interface{}) (t []*originTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &originTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeoriginTravelIntermediateType will accept a originTravelIntermediateType to create a map
func serializeOriginTravelIntermediateType(t *originTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceoriginTravelIntermediateType will accept a slice of originTravelIntermediateType to create a slice result
func serializeSliceOriginTravelIntermediateType(s []*originTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinstrumentTravelIntermediateType will accept a map to create a instrumentTravelIntermediateType
func deserializeInstrumentTravelIntermediateType(in interface{}) (t *instrumentTravelIntermediateType, err error) {
	tmp := &instrumentTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice instrumentTravelIntermediateType will accept a slice to create a slice of instrumentTravelIntermediateType
func deserializeSliceInstrumentTravelIntermediateType(in []interface{}) (t []*instrumentTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &instrumentTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinstrumentTravelIntermediateType will accept a instrumentTravelIntermediateType to create a map
func serializeInstrumentTravelIntermediateType(t *instrumentTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinstrumentTravelIntermediateType will accept a slice of instrumentTravelIntermediateType to create a slice result
func serializeSliceInstrumentTravelIntermediateType(s []*instrumentTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializealtitudeTravelIntermediateType will accept a map to create a altitudeTravelIntermediateType
func deserializeAltitudeTravelIntermediateType(in interface{}) (t *altitudeTravelIntermediateType, err error) {
	tmp := &altitudeTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice altitudeTravelIntermediateType will accept a slice to create a slice of altitudeTravelIntermediateType
func deserializeSliceAltitudeTravelIntermediateType(in []interface{}) (t []*altitudeTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &altitudeTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializealtitudeTravelIntermediateType will accept a altitudeTravelIntermediateType to create a map
func serializeAltitudeTravelIntermediateType(t *altitudeTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicealtitudeTravelIntermediateType will accept a slice of altitudeTravelIntermediateType to create a slice result
func serializeSliceAltitudeTravelIntermediateType(s []*altitudeTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattachmentTravelIntermediateType will accept a map to create a attachmentTravelIntermediateType
func deserializeAttachmentTravelIntermediateType(in interface{}) (t *attachmentTravelIntermediateType, err error) {
	tmp := &attachmentTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attachmentTravelIntermediateType will accept a slice to create a slice of attachmentTravelIntermediateType
func deserializeSliceAttachmentTravelIntermediateType(in []interface{}) (t []*attachmentTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &attachmentTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattachmentTravelIntermediateType will accept a attachmentTravelIntermediateType to create a map
func serializeAttachmentTravelIntermediateType(t *attachmentTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattachmentTravelIntermediateType will accept a slice of attachmentTravelIntermediateType to create a slice result
func serializeSliceAttachmentTravelIntermediateType(s []*attachmentTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattributedToTravelIntermediateType will accept a map to create a attributedToTravelIntermediateType
func deserializeAttributedToTravelIntermediateType(in interface{}) (t *attributedToTravelIntermediateType, err error) {
	tmp := &attributedToTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attributedToTravelIntermediateType will accept a slice to create a slice of attributedToTravelIntermediateType
func deserializeSliceAttributedToTravelIntermediateType(in []interface{}) (t []*attributedToTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &attributedToTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattributedToTravelIntermediateType will accept a attributedToTravelIntermediateType to create a map
func serializeAttributedToTravelIntermediateType(t *attributedToTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattributedToTravelIntermediateType will accept a slice of attributedToTravelIntermediateType to create a slice result
func serializeSliceAttributedToTravelIntermediateType(s []*attributedToTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeaudienceTravelIntermediateType will accept a map to create a audienceTravelIntermediateType
func deserializeAudienceTravelIntermediateType(in interface{}) (t *audienceTravelIntermediateType, err error) {
	tmp := &audienceTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice audienceTravelIntermediateType will accept a slice to create a slice of audienceTravelIntermediateType
func deserializeSliceAudienceTravelIntermediateType(in []interface{}) (t []*audienceTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &audienceTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeaudienceTravelIntermediateType will accept a audienceTravelIntermediateType to create a map
func serializeAudienceTravelIntermediateType(t *audienceTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceaudienceTravelIntermediateType will accept a slice of audienceTravelIntermediateType to create a slice result
func serializeSliceAudienceTravelIntermediateType(s []*audienceTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontentTravelIntermediateType will accept a map to create a contentTravelIntermediateType
func deserializeContentTravelIntermediateType(in interface{}) (t *contentTravelIntermediateType, err error) {
	tmp := &contentTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contentTravelIntermediateType will accept a slice to create a slice of contentTravelIntermediateType
func deserializeSliceContentTravelIntermediateType(in []interface{}) (t []*contentTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &contentTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontentTravelIntermediateType will accept a contentTravelIntermediateType to create a map
func serializeContentTravelIntermediateType(t *contentTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontentTravelIntermediateType will accept a slice of contentTravelIntermediateType to create a slice result
func serializeSliceContentTravelIntermediateType(s []*contentTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontextTravelIntermediateType will accept a map to create a contextTravelIntermediateType
func deserializeContextTravelIntermediateType(in interface{}) (t *contextTravelIntermediateType, err error) {
	tmp := &contextTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contextTravelIntermediateType will accept a slice to create a slice of contextTravelIntermediateType
func deserializeSliceContextTravelIntermediateType(in []interface{}) (t []*contextTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &contextTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontextTravelIntermediateType will accept a contextTravelIntermediateType to create a map
func serializeContextTravelIntermediateType(t *contextTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontextTravelIntermediateType will accept a slice of contextTravelIntermediateType to create a slice result
func serializeSliceContextTravelIntermediateType(s []*contextTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializenameTravelIntermediateType will accept a map to create a nameTravelIntermediateType
func deserializeNameTravelIntermediateType(in interface{}) (t *nameTravelIntermediateType, err error) {
	tmp := &nameTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice nameTravelIntermediateType will accept a slice to create a slice of nameTravelIntermediateType
func deserializeSliceNameTravelIntermediateType(in []interface{}) (t []*nameTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &nameTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializenameTravelIntermediateType will accept a nameTravelIntermediateType to create a map
func serializeNameTravelIntermediateType(t *nameTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicenameTravelIntermediateType will accept a slice of nameTravelIntermediateType to create a slice result
func serializeSliceNameTravelIntermediateType(s []*nameTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendTimeTravelIntermediateType will accept a map to create a endTimeTravelIntermediateType
func deserializeEndTimeTravelIntermediateType(in interface{}) (t *endTimeTravelIntermediateType, err error) {
	tmp := &endTimeTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endTimeTravelIntermediateType will accept a slice to create a slice of endTimeTravelIntermediateType
func deserializeSliceEndTimeTravelIntermediateType(in []interface{}) (t []*endTimeTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &endTimeTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendTimeTravelIntermediateType will accept a endTimeTravelIntermediateType to create a map
func serializeEndTimeTravelIntermediateType(t *endTimeTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendTimeTravelIntermediateType will accept a slice of endTimeTravelIntermediateType to create a slice result
func serializeSliceEndTimeTravelIntermediateType(s []*endTimeTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializegeneratorTravelIntermediateType will accept a map to create a generatorTravelIntermediateType
func deserializeGeneratorTravelIntermediateType(in interface{}) (t *generatorTravelIntermediateType, err error) {
	tmp := &generatorTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice generatorTravelIntermediateType will accept a slice to create a slice of generatorTravelIntermediateType
func deserializeSliceGeneratorTravelIntermediateType(in []interface{}) (t []*generatorTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &generatorTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializegeneratorTravelIntermediateType will accept a generatorTravelIntermediateType to create a map
func serializeGeneratorTravelIntermediateType(t *generatorTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicegeneratorTravelIntermediateType will accept a slice of generatorTravelIntermediateType to create a slice result
func serializeSliceGeneratorTravelIntermediateType(s []*generatorTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeiconTravelIntermediateType will accept a map to create a iconTravelIntermediateType
func deserializeIconTravelIntermediateType(in interface{}) (t *iconTravelIntermediateType, err error) {
	tmp := &iconTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice iconTravelIntermediateType will accept a slice to create a slice of iconTravelIntermediateType
func deserializeSliceIconTravelIntermediateType(in []interface{}) (t []*iconTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &iconTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeiconTravelIntermediateType will accept a iconTravelIntermediateType to create a map
func serializeIconTravelIntermediateType(t *iconTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceiconTravelIntermediateType will accept a slice of iconTravelIntermediateType to create a slice result
func serializeSliceIconTravelIntermediateType(s []*iconTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeimageTravelIntermediateType will accept a map to create a imageTravelIntermediateType
func deserializeImageTravelIntermediateType(in interface{}) (t *imageTravelIntermediateType, err error) {
	tmp := &imageTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice imageTravelIntermediateType will accept a slice to create a slice of imageTravelIntermediateType
func deserializeSliceImageTravelIntermediateType(in []interface{}) (t []*imageTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &imageTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeimageTravelIntermediateType will accept a imageTravelIntermediateType to create a map
func serializeImageTravelIntermediateType(t *imageTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceimageTravelIntermediateType will accept a slice of imageTravelIntermediateType to create a slice result
func serializeSliceImageTravelIntermediateType(s []*imageTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinReplyToTravelIntermediateType will accept a map to create a inReplyToTravelIntermediateType
func deserializeInReplyToTravelIntermediateType(in interface{}) (t *inReplyToTravelIntermediateType, err error) {
	tmp := &inReplyToTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inReplyToTravelIntermediateType will accept a slice to create a slice of inReplyToTravelIntermediateType
func deserializeSliceInReplyToTravelIntermediateType(in []interface{}) (t []*inReplyToTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &inReplyToTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinReplyToTravelIntermediateType will accept a inReplyToTravelIntermediateType to create a map
func serializeInReplyToTravelIntermediateType(t *inReplyToTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinReplyToTravelIntermediateType will accept a slice of inReplyToTravelIntermediateType to create a slice result
func serializeSliceInReplyToTravelIntermediateType(s []*inReplyToTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelocationTravelIntermediateType will accept a map to create a locationTravelIntermediateType
func deserializeLocationTravelIntermediateType(in interface{}) (t *locationTravelIntermediateType, err error) {
	tmp := &locationTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice locationTravelIntermediateType will accept a slice to create a slice of locationTravelIntermediateType
func deserializeSliceLocationTravelIntermediateType(in []interface{}) (t []*locationTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &locationTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelocationTravelIntermediateType will accept a locationTravelIntermediateType to create a map
func serializeLocationTravelIntermediateType(t *locationTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelocationTravelIntermediateType will accept a slice of locationTravelIntermediateType to create a slice result
func serializeSliceLocationTravelIntermediateType(s []*locationTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreviewTravelIntermediateType will accept a map to create a previewTravelIntermediateType
func deserializePreviewTravelIntermediateType(in interface{}) (t *previewTravelIntermediateType, err error) {
	tmp := &previewTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice previewTravelIntermediateType will accept a slice to create a slice of previewTravelIntermediateType
func deserializeSlicePreviewTravelIntermediateType(in []interface{}) (t []*previewTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &previewTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreviewTravelIntermediateType will accept a previewTravelIntermediateType to create a map
func serializePreviewTravelIntermediateType(t *previewTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreviewTravelIntermediateType will accept a slice of previewTravelIntermediateType to create a slice result
func serializeSlicePreviewTravelIntermediateType(s []*previewTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepublishedTravelIntermediateType will accept a map to create a publishedTravelIntermediateType
func deserializePublishedTravelIntermediateType(in interface{}) (t *publishedTravelIntermediateType, err error) {
	tmp := &publishedTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice publishedTravelIntermediateType will accept a slice to create a slice of publishedTravelIntermediateType
func deserializeSlicePublishedTravelIntermediateType(in []interface{}) (t []*publishedTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &publishedTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepublishedTravelIntermediateType will accept a publishedTravelIntermediateType to create a map
func serializePublishedTravelIntermediateType(t *publishedTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepublishedTravelIntermediateType will accept a slice of publishedTravelIntermediateType to create a slice result
func serializeSlicePublishedTravelIntermediateType(s []*publishedTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializerepliesTravelIntermediateType will accept a map to create a repliesTravelIntermediateType
func deserializeRepliesTravelIntermediateType(in interface{}) (t *repliesTravelIntermediateType, err error) {
	tmp := &repliesTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice repliesTravelIntermediateType will accept a slice to create a slice of repliesTravelIntermediateType
func deserializeSliceRepliesTravelIntermediateType(in []interface{}) (t []*repliesTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &repliesTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializerepliesTravelIntermediateType will accept a repliesTravelIntermediateType to create a map
func serializeRepliesTravelIntermediateType(t *repliesTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicerepliesTravelIntermediateType will accept a slice of repliesTravelIntermediateType to create a slice result
func serializeSliceRepliesTravelIntermediateType(s []*repliesTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializestartTimeTravelIntermediateType will accept a map to create a startTimeTravelIntermediateType
func deserializeStartTimeTravelIntermediateType(in interface{}) (t *startTimeTravelIntermediateType, err error) {
	tmp := &startTimeTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice startTimeTravelIntermediateType will accept a slice to create a slice of startTimeTravelIntermediateType
func deserializeSliceStartTimeTravelIntermediateType(in []interface{}) (t []*startTimeTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &startTimeTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializestartTimeTravelIntermediateType will accept a startTimeTravelIntermediateType to create a map
func serializeStartTimeTravelIntermediateType(t *startTimeTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicestartTimeTravelIntermediateType will accept a slice of startTimeTravelIntermediateType to create a slice result
func serializeSliceStartTimeTravelIntermediateType(s []*startTimeTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesummaryTravelIntermediateType will accept a map to create a summaryTravelIntermediateType
func deserializeSummaryTravelIntermediateType(in interface{}) (t *summaryTravelIntermediateType, err error) {
	tmp := &summaryTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice summaryTravelIntermediateType will accept a slice to create a slice of summaryTravelIntermediateType
func deserializeSliceSummaryTravelIntermediateType(in []interface{}) (t []*summaryTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &summaryTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesummaryTravelIntermediateType will accept a summaryTravelIntermediateType to create a map
func serializeSummaryTravelIntermediateType(t *summaryTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesummaryTravelIntermediateType will accept a slice of summaryTravelIntermediateType to create a slice result
func serializeSliceSummaryTravelIntermediateType(s []*summaryTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetagTravelIntermediateType will accept a map to create a tagTravelIntermediateType
func deserializeTagTravelIntermediateType(in interface{}) (t *tagTravelIntermediateType, err error) {
	tmp := &tagTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice tagTravelIntermediateType will accept a slice to create a slice of tagTravelIntermediateType
func deserializeSliceTagTravelIntermediateType(in []interface{}) (t []*tagTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &tagTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetagTravelIntermediateType will accept a tagTravelIntermediateType to create a map
func serializeTagTravelIntermediateType(t *tagTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetagTravelIntermediateType will accept a slice of tagTravelIntermediateType to create a slice result
func serializeSliceTagTravelIntermediateType(s []*tagTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeupdatedTravelIntermediateType will accept a map to create a updatedTravelIntermediateType
func deserializeUpdatedTravelIntermediateType(in interface{}) (t *updatedTravelIntermediateType, err error) {
	tmp := &updatedTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice updatedTravelIntermediateType will accept a slice to create a slice of updatedTravelIntermediateType
func deserializeSliceUpdatedTravelIntermediateType(in []interface{}) (t []*updatedTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &updatedTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeupdatedTravelIntermediateType will accept a updatedTravelIntermediateType to create a map
func serializeUpdatedTravelIntermediateType(t *updatedTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceupdatedTravelIntermediateType will accept a slice of updatedTravelIntermediateType to create a slice result
func serializeSliceUpdatedTravelIntermediateType(s []*updatedTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeurlTravelIntermediateType will accept a map to create a urlTravelIntermediateType
func deserializeUrlTravelIntermediateType(in interface{}) (t *urlTravelIntermediateType, err error) {
	tmp := &urlTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice urlTravelIntermediateType will accept a slice to create a slice of urlTravelIntermediateType
func deserializeSliceUrlTravelIntermediateType(in []interface{}) (t []*urlTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &urlTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeurlTravelIntermediateType will accept a urlTravelIntermediateType to create a map
func serializeUrlTravelIntermediateType(t *urlTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceurlTravelIntermediateType will accept a slice of urlTravelIntermediateType to create a slice result
func serializeSliceUrlTravelIntermediateType(s []*urlTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetoTravelIntermediateType will accept a map to create a toTravelIntermediateType
func deserializeToTravelIntermediateType(in interface{}) (t *toTravelIntermediateType, err error) {
	tmp := &toTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice toTravelIntermediateType will accept a slice to create a slice of toTravelIntermediateType
func deserializeSliceToTravelIntermediateType(in []interface{}) (t []*toTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &toTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetoTravelIntermediateType will accept a toTravelIntermediateType to create a map
func serializeToTravelIntermediateType(t *toTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetoTravelIntermediateType will accept a slice of toTravelIntermediateType to create a slice result
func serializeSliceToTravelIntermediateType(s []*toTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebtoTravelIntermediateType will accept a map to create a btoTravelIntermediateType
func deserializeBtoTravelIntermediateType(in interface{}) (t *btoTravelIntermediateType, err error) {
	tmp := &btoTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice btoTravelIntermediateType will accept a slice to create a slice of btoTravelIntermediateType
func deserializeSliceBtoTravelIntermediateType(in []interface{}) (t []*btoTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &btoTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebtoTravelIntermediateType will accept a btoTravelIntermediateType to create a map
func serializeBtoTravelIntermediateType(t *btoTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebtoTravelIntermediateType will accept a slice of btoTravelIntermediateType to create a slice result
func serializeSliceBtoTravelIntermediateType(s []*btoTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeccTravelIntermediateType will accept a map to create a ccTravelIntermediateType
func deserializeCcTravelIntermediateType(in interface{}) (t *ccTravelIntermediateType, err error) {
	tmp := &ccTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice ccTravelIntermediateType will accept a slice to create a slice of ccTravelIntermediateType
func deserializeSliceCcTravelIntermediateType(in []interface{}) (t []*ccTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &ccTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeccTravelIntermediateType will accept a ccTravelIntermediateType to create a map
func serializeCcTravelIntermediateType(t *ccTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceccTravelIntermediateType will accept a slice of ccTravelIntermediateType to create a slice result
func serializeSliceCcTravelIntermediateType(s []*ccTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebccTravelIntermediateType will accept a map to create a bccTravelIntermediateType
func deserializeBccTravelIntermediateType(in interface{}) (t *bccTravelIntermediateType, err error) {
	tmp := &bccTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice bccTravelIntermediateType will accept a slice to create a slice of bccTravelIntermediateType
func deserializeSliceBccTravelIntermediateType(in []interface{}) (t []*bccTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &bccTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebccTravelIntermediateType will accept a bccTravelIntermediateType to create a map
func serializeBccTravelIntermediateType(t *bccTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebccTravelIntermediateType will accept a slice of bccTravelIntermediateType to create a slice result
func serializeSliceBccTravelIntermediateType(s []*bccTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializemediaTypeTravelIntermediateType will accept a map to create a mediaTypeTravelIntermediateType
func deserializeMediaTypeTravelIntermediateType(in interface{}) (t *mediaTypeTravelIntermediateType, err error) {
	tmp := &mediaTypeTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice mediaTypeTravelIntermediateType will accept a slice to create a slice of mediaTypeTravelIntermediateType
func deserializeSliceMediaTypeTravelIntermediateType(in []interface{}) (t []*mediaTypeTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &mediaTypeTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializemediaTypeTravelIntermediateType will accept a mediaTypeTravelIntermediateType to create a map
func serializeMediaTypeTravelIntermediateType(t *mediaTypeTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicemediaTypeTravelIntermediateType will accept a slice of mediaTypeTravelIntermediateType to create a slice result
func serializeSliceMediaTypeTravelIntermediateType(s []*mediaTypeTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializedurationTravelIntermediateType will accept a map to create a durationTravelIntermediateType
func deserializeDurationTravelIntermediateType(in interface{}) (t *durationTravelIntermediateType, err error) {
	tmp := &durationTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice durationTravelIntermediateType will accept a slice to create a slice of durationTravelIntermediateType
func deserializeSliceDurationTravelIntermediateType(in []interface{}) (t []*durationTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &durationTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializedurationTravelIntermediateType will accept a durationTravelIntermediateType to create a map
func serializeDurationTravelIntermediateType(t *durationTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicedurationTravelIntermediateType will accept a slice of durationTravelIntermediateType to create a slice result
func serializeSliceDurationTravelIntermediateType(s []*durationTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesourceTravelIntermediateType will accept a map to create a sourceTravelIntermediateType
func deserializeSourceTravelIntermediateType(in interface{}) (t *sourceTravelIntermediateType, err error) {
	tmp := &sourceTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice sourceTravelIntermediateType will accept a slice to create a slice of sourceTravelIntermediateType
func deserializeSliceSourceTravelIntermediateType(in []interface{}) (t []*sourceTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &sourceTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesourceTravelIntermediateType will accept a sourceTravelIntermediateType to create a map
func serializeSourceTravelIntermediateType(t *sourceTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesourceTravelIntermediateType will accept a slice of sourceTravelIntermediateType to create a slice result
func serializeSliceSourceTravelIntermediateType(s []*sourceTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinboxTravelIntermediateType will accept a map to create a inboxTravelIntermediateType
func deserializeInboxTravelIntermediateType(in interface{}) (t *inboxTravelIntermediateType, err error) {
	tmp := &inboxTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inboxTravelIntermediateType will accept a slice to create a slice of inboxTravelIntermediateType
func deserializeSliceInboxTravelIntermediateType(in []interface{}) (t []*inboxTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &inboxTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinboxTravelIntermediateType will accept a inboxTravelIntermediateType to create a map
func serializeInboxTravelIntermediateType(t *inboxTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinboxTravelIntermediateType will accept a slice of inboxTravelIntermediateType to create a slice result
func serializeSliceInboxTravelIntermediateType(s []*inboxTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeoutboxTravelIntermediateType will accept a map to create a outboxTravelIntermediateType
func deserializeOutboxTravelIntermediateType(in interface{}) (t *outboxTravelIntermediateType, err error) {
	tmp := &outboxTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice outboxTravelIntermediateType will accept a slice to create a slice of outboxTravelIntermediateType
func deserializeSliceOutboxTravelIntermediateType(in []interface{}) (t []*outboxTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &outboxTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeoutboxTravelIntermediateType will accept a outboxTravelIntermediateType to create a map
func serializeOutboxTravelIntermediateType(t *outboxTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceoutboxTravelIntermediateType will accept a slice of outboxTravelIntermediateType to create a slice result
func serializeSliceOutboxTravelIntermediateType(s []*outboxTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowingTravelIntermediateType will accept a map to create a followingTravelIntermediateType
func deserializeFollowingTravelIntermediateType(in interface{}) (t *followingTravelIntermediateType, err error) {
	tmp := &followingTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followingTravelIntermediateType will accept a slice to create a slice of followingTravelIntermediateType
func deserializeSliceFollowingTravelIntermediateType(in []interface{}) (t []*followingTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &followingTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowingTravelIntermediateType will accept a followingTravelIntermediateType to create a map
func serializeFollowingTravelIntermediateType(t *followingTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowingTravelIntermediateType will accept a slice of followingTravelIntermediateType to create a slice result
func serializeSliceFollowingTravelIntermediateType(s []*followingTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowersTravelIntermediateType will accept a map to create a followersTravelIntermediateType
func deserializeFollowersTravelIntermediateType(in interface{}) (t *followersTravelIntermediateType, err error) {
	tmp := &followersTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followersTravelIntermediateType will accept a slice to create a slice of followersTravelIntermediateType
func deserializeSliceFollowersTravelIntermediateType(in []interface{}) (t []*followersTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &followersTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowersTravelIntermediateType will accept a followersTravelIntermediateType to create a map
func serializeFollowersTravelIntermediateType(t *followersTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowersTravelIntermediateType will accept a slice of followersTravelIntermediateType to create a slice result
func serializeSliceFollowersTravelIntermediateType(s []*followersTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikedTravelIntermediateType will accept a map to create a likedTravelIntermediateType
func deserializeLikedTravelIntermediateType(in interface{}) (t *likedTravelIntermediateType, err error) {
	tmp := &likedTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likedTravelIntermediateType will accept a slice to create a slice of likedTravelIntermediateType
func deserializeSliceLikedTravelIntermediateType(in []interface{}) (t []*likedTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &likedTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikedTravelIntermediateType will accept a likedTravelIntermediateType to create a map
func serializeLikedTravelIntermediateType(t *likedTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikedTravelIntermediateType will accept a slice of likedTravelIntermediateType to create a slice result
func serializeSliceLikedTravelIntermediateType(s []*likedTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikesTravelIntermediateType will accept a map to create a likesTravelIntermediateType
func deserializeLikesTravelIntermediateType(in interface{}) (t *likesTravelIntermediateType, err error) {
	tmp := &likesTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likesTravelIntermediateType will accept a slice to create a slice of likesTravelIntermediateType
func deserializeSliceLikesTravelIntermediateType(in []interface{}) (t []*likesTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &likesTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikesTravelIntermediateType will accept a likesTravelIntermediateType to create a map
func serializeLikesTravelIntermediateType(t *likesTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikesTravelIntermediateType will accept a slice of likesTravelIntermediateType to create a slice result
func serializeSliceLikesTravelIntermediateType(s []*likesTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreferredUsernameTravelIntermediateType will accept a map to create a preferredUsernameTravelIntermediateType
func deserializePreferredUsernameTravelIntermediateType(in interface{}) (t *preferredUsernameTravelIntermediateType, err error) {
	tmp := &preferredUsernameTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice preferredUsernameTravelIntermediateType will accept a slice to create a slice of preferredUsernameTravelIntermediateType
func deserializeSlicePreferredUsernameTravelIntermediateType(in []interface{}) (t []*preferredUsernameTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &preferredUsernameTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreferredUsernameTravelIntermediateType will accept a preferredUsernameTravelIntermediateType to create a map
func serializePreferredUsernameTravelIntermediateType(t *preferredUsernameTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreferredUsernameTravelIntermediateType will accept a slice of preferredUsernameTravelIntermediateType to create a slice result
func serializeSlicePreferredUsernameTravelIntermediateType(s []*preferredUsernameTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendpointsTravelIntermediateType will accept a map to create a endpointsTravelIntermediateType
func deserializeEndpointsTravelIntermediateType(in interface{}) (t *endpointsTravelIntermediateType, err error) {
	tmp := &endpointsTravelIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endpointsTravelIntermediateType will accept a slice to create a slice of endpointsTravelIntermediateType
func deserializeSliceEndpointsTravelIntermediateType(in []interface{}) (t []*endpointsTravelIntermediateType, err error) {
	for _, i := range in {
		tmp := &endpointsTravelIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendpointsTravelIntermediateType will accept a endpointsTravelIntermediateType to create a map
func serializeEndpointsTravelIntermediateType(t *endpointsTravelIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendpointsTravelIntermediateType will accept a slice of endpointsTravelIntermediateType to create a slice result
func serializeSliceEndpointsTravelIntermediateType(s []*endpointsTravelIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}