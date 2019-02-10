package vocab

import "net/url"

// ActivityStreamsObjectPropertyIterator represents a single value for the
// "object" property.
type ActivityStreamsObjectPropertyIterator interface {
	// GetActivityStreamsAccept returns the value of this property. When
	// IsActivityStreamsAccept returns false, GetActivityStreamsAccept
	// will return an arbitrary value.
	GetActivityStreamsAccept() ActivityStreamsAccept
	// GetActivityStreamsActivity returns the value of this property. When
	// IsActivityStreamsActivity returns false, GetActivityStreamsActivity
	// will return an arbitrary value.
	GetActivityStreamsActivity() ActivityStreamsActivity
	// GetActivityStreamsAdd returns the value of this property. When
	// IsActivityStreamsAdd returns false, GetActivityStreamsAdd will
	// return an arbitrary value.
	GetActivityStreamsAdd() ActivityStreamsAdd
	// GetActivityStreamsAnnounce returns the value of this property. When
	// IsActivityStreamsAnnounce returns false, GetActivityStreamsAnnounce
	// will return an arbitrary value.
	GetActivityStreamsAnnounce() ActivityStreamsAnnounce
	// GetActivityStreamsApplication returns the value of this property. When
	// IsActivityStreamsApplication returns false,
	// GetActivityStreamsApplication will return an arbitrary value.
	GetActivityStreamsApplication() ActivityStreamsApplication
	// GetActivityStreamsArrive returns the value of this property. When
	// IsActivityStreamsArrive returns false, GetActivityStreamsArrive
	// will return an arbitrary value.
	GetActivityStreamsArrive() ActivityStreamsArrive
	// GetActivityStreamsArticle returns the value of this property. When
	// IsActivityStreamsArticle returns false, GetActivityStreamsArticle
	// will return an arbitrary value.
	GetActivityStreamsArticle() ActivityStreamsArticle
	// GetActivityStreamsAudio returns the value of this property. When
	// IsActivityStreamsAudio returns false, GetActivityStreamsAudio will
	// return an arbitrary value.
	GetActivityStreamsAudio() ActivityStreamsAudio
	// GetActivityStreamsBlock returns the value of this property. When
	// IsActivityStreamsBlock returns false, GetActivityStreamsBlock will
	// return an arbitrary value.
	GetActivityStreamsBlock() ActivityStreamsBlock
	// GetActivityStreamsCollection returns the value of this property. When
	// IsActivityStreamsCollection returns false,
	// GetActivityStreamsCollection will return an arbitrary value.
	GetActivityStreamsCollection() ActivityStreamsCollection
	// GetActivityStreamsCollectionPage returns the value of this property.
	// When IsActivityStreamsCollectionPage returns false,
	// GetActivityStreamsCollectionPage will return an arbitrary value.
	GetActivityStreamsCollectionPage() ActivityStreamsCollectionPage
	// GetActivityStreamsCreate returns the value of this property. When
	// IsActivityStreamsCreate returns false, GetActivityStreamsCreate
	// will return an arbitrary value.
	GetActivityStreamsCreate() ActivityStreamsCreate
	// GetActivityStreamsDelete returns the value of this property. When
	// IsActivityStreamsDelete returns false, GetActivityStreamsDelete
	// will return an arbitrary value.
	GetActivityStreamsDelete() ActivityStreamsDelete
	// GetActivityStreamsDislike returns the value of this property. When
	// IsActivityStreamsDislike returns false, GetActivityStreamsDislike
	// will return an arbitrary value.
	GetActivityStreamsDislike() ActivityStreamsDislike
	// GetActivityStreamsDocument returns the value of this property. When
	// IsActivityStreamsDocument returns false, GetActivityStreamsDocument
	// will return an arbitrary value.
	GetActivityStreamsDocument() ActivityStreamsDocument
	// GetActivityStreamsEvent returns the value of this property. When
	// IsActivityStreamsEvent returns false, GetActivityStreamsEvent will
	// return an arbitrary value.
	GetActivityStreamsEvent() ActivityStreamsEvent
	// GetActivityStreamsFlag returns the value of this property. When
	// IsActivityStreamsFlag returns false, GetActivityStreamsFlag will
	// return an arbitrary value.
	GetActivityStreamsFlag() ActivityStreamsFlag
	// GetActivityStreamsFollow returns the value of this property. When
	// IsActivityStreamsFollow returns false, GetActivityStreamsFollow
	// will return an arbitrary value.
	GetActivityStreamsFollow() ActivityStreamsFollow
	// GetActivityStreamsGroup returns the value of this property. When
	// IsActivityStreamsGroup returns false, GetActivityStreamsGroup will
	// return an arbitrary value.
	GetActivityStreamsGroup() ActivityStreamsGroup
	// GetActivityStreamsIgnore returns the value of this property. When
	// IsActivityStreamsIgnore returns false, GetActivityStreamsIgnore
	// will return an arbitrary value.
	GetActivityStreamsIgnore() ActivityStreamsIgnore
	// GetActivityStreamsImage returns the value of this property. When
	// IsActivityStreamsImage returns false, GetActivityStreamsImage will
	// return an arbitrary value.
	GetActivityStreamsImage() ActivityStreamsImage
	// GetActivityStreamsIntransitiveActivity returns the value of this
	// property. When IsActivityStreamsIntransitiveActivity returns false,
	// GetActivityStreamsIntransitiveActivity will return an arbitrary
	// value.
	GetActivityStreamsIntransitiveActivity() ActivityStreamsIntransitiveActivity
	// GetActivityStreamsInvite returns the value of this property. When
	// IsActivityStreamsInvite returns false, GetActivityStreamsInvite
	// will return an arbitrary value.
	GetActivityStreamsInvite() ActivityStreamsInvite
	// GetActivityStreamsJoin returns the value of this property. When
	// IsActivityStreamsJoin returns false, GetActivityStreamsJoin will
	// return an arbitrary value.
	GetActivityStreamsJoin() ActivityStreamsJoin
	// GetActivityStreamsLeave returns the value of this property. When
	// IsActivityStreamsLeave returns false, GetActivityStreamsLeave will
	// return an arbitrary value.
	GetActivityStreamsLeave() ActivityStreamsLeave
	// GetActivityStreamsLike returns the value of this property. When
	// IsActivityStreamsLike returns false, GetActivityStreamsLike will
	// return an arbitrary value.
	GetActivityStreamsLike() ActivityStreamsLike
	// GetActivityStreamsLink returns the value of this property. When
	// IsActivityStreamsLink returns false, GetActivityStreamsLink will
	// return an arbitrary value.
	GetActivityStreamsLink() ActivityStreamsLink
	// GetActivityStreamsListen returns the value of this property. When
	// IsActivityStreamsListen returns false, GetActivityStreamsListen
	// will return an arbitrary value.
	GetActivityStreamsListen() ActivityStreamsListen
	// GetActivityStreamsMention returns the value of this property. When
	// IsActivityStreamsMention returns false, GetActivityStreamsMention
	// will return an arbitrary value.
	GetActivityStreamsMention() ActivityStreamsMention
	// GetActivityStreamsMove returns the value of this property. When
	// IsActivityStreamsMove returns false, GetActivityStreamsMove will
	// return an arbitrary value.
	GetActivityStreamsMove() ActivityStreamsMove
	// GetActivityStreamsNote returns the value of this property. When
	// IsActivityStreamsNote returns false, GetActivityStreamsNote will
	// return an arbitrary value.
	GetActivityStreamsNote() ActivityStreamsNote
	// GetActivityStreamsObject returns the value of this property. When
	// IsActivityStreamsObject returns false, GetActivityStreamsObject
	// will return an arbitrary value.
	GetActivityStreamsObject() ActivityStreamsObject
	// GetActivityStreamsOffer returns the value of this property. When
	// IsActivityStreamsOffer returns false, GetActivityStreamsOffer will
	// return an arbitrary value.
	GetActivityStreamsOffer() ActivityStreamsOffer
	// GetActivityStreamsOrderedCollection returns the value of this property.
	// When IsActivityStreamsOrderedCollection returns false,
	// GetActivityStreamsOrderedCollection will return an arbitrary value.
	GetActivityStreamsOrderedCollection() ActivityStreamsOrderedCollection
	// GetActivityStreamsOrderedCollectionPage returns the value of this
	// property. When IsActivityStreamsOrderedCollectionPage returns
	// false, GetActivityStreamsOrderedCollectionPage will return an
	// arbitrary value.
	GetActivityStreamsOrderedCollectionPage() ActivityStreamsOrderedCollectionPage
	// GetActivityStreamsOrganization returns the value of this property. When
	// IsActivityStreamsOrganization returns false,
	// GetActivityStreamsOrganization will return an arbitrary value.
	GetActivityStreamsOrganization() ActivityStreamsOrganization
	// GetActivityStreamsPage returns the value of this property. When
	// IsActivityStreamsPage returns false, GetActivityStreamsPage will
	// return an arbitrary value.
	GetActivityStreamsPage() ActivityStreamsPage
	// GetActivityStreamsPerson returns the value of this property. When
	// IsActivityStreamsPerson returns false, GetActivityStreamsPerson
	// will return an arbitrary value.
	GetActivityStreamsPerson() ActivityStreamsPerson
	// GetActivityStreamsPlace returns the value of this property. When
	// IsActivityStreamsPlace returns false, GetActivityStreamsPlace will
	// return an arbitrary value.
	GetActivityStreamsPlace() ActivityStreamsPlace
	// GetActivityStreamsProfile returns the value of this property. When
	// IsActivityStreamsProfile returns false, GetActivityStreamsProfile
	// will return an arbitrary value.
	GetActivityStreamsProfile() ActivityStreamsProfile
	// GetActivityStreamsQuestion returns the value of this property. When
	// IsActivityStreamsQuestion returns false, GetActivityStreamsQuestion
	// will return an arbitrary value.
	GetActivityStreamsQuestion() ActivityStreamsQuestion
	// GetActivityStreamsRead returns the value of this property. When
	// IsActivityStreamsRead returns false, GetActivityStreamsRead will
	// return an arbitrary value.
	GetActivityStreamsRead() ActivityStreamsRead
	// GetActivityStreamsReject returns the value of this property. When
	// IsActivityStreamsReject returns false, GetActivityStreamsReject
	// will return an arbitrary value.
	GetActivityStreamsReject() ActivityStreamsReject
	// GetActivityStreamsRelationship returns the value of this property. When
	// IsActivityStreamsRelationship returns false,
	// GetActivityStreamsRelationship will return an arbitrary value.
	GetActivityStreamsRelationship() ActivityStreamsRelationship
	// GetActivityStreamsRemove returns the value of this property. When
	// IsActivityStreamsRemove returns false, GetActivityStreamsRemove
	// will return an arbitrary value.
	GetActivityStreamsRemove() ActivityStreamsRemove
	// GetActivityStreamsService returns the value of this property. When
	// IsActivityStreamsService returns false, GetActivityStreamsService
	// will return an arbitrary value.
	GetActivityStreamsService() ActivityStreamsService
	// GetActivityStreamsTentativeAccept returns the value of this property.
	// When IsActivityStreamsTentativeAccept returns false,
	// GetActivityStreamsTentativeAccept will return an arbitrary value.
	GetActivityStreamsTentativeAccept() ActivityStreamsTentativeAccept
	// GetActivityStreamsTentativeReject returns the value of this property.
	// When IsActivityStreamsTentativeReject returns false,
	// GetActivityStreamsTentativeReject will return an arbitrary value.
	GetActivityStreamsTentativeReject() ActivityStreamsTentativeReject
	// GetActivityStreamsTombstone returns the value of this property. When
	// IsActivityStreamsTombstone returns false,
	// GetActivityStreamsTombstone will return an arbitrary value.
	GetActivityStreamsTombstone() ActivityStreamsTombstone
	// GetActivityStreamsTravel returns the value of this property. When
	// IsActivityStreamsTravel returns false, GetActivityStreamsTravel
	// will return an arbitrary value.
	GetActivityStreamsTravel() ActivityStreamsTravel
	// GetActivityStreamsUndo returns the value of this property. When
	// IsActivityStreamsUndo returns false, GetActivityStreamsUndo will
	// return an arbitrary value.
	GetActivityStreamsUndo() ActivityStreamsUndo
	// GetActivityStreamsUpdate returns the value of this property. When
	// IsActivityStreamsUpdate returns false, GetActivityStreamsUpdate
	// will return an arbitrary value.
	GetActivityStreamsUpdate() ActivityStreamsUpdate
	// GetActivityStreamsVideo returns the value of this property. When
	// IsActivityStreamsVideo returns false, GetActivityStreamsVideo will
	// return an arbitrary value.
	GetActivityStreamsVideo() ActivityStreamsVideo
	// GetActivityStreamsView returns the value of this property. When
	// IsActivityStreamsView returns false, GetActivityStreamsView will
	// return an arbitrary value.
	GetActivityStreamsView() ActivityStreamsView
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetType returns the value in this property as a Type. Returns nil if
	// the value is not an ActivityStreams type, such as an IRI or another
	// value.
	GetType() Type
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsActivityStreamsAccept returns true if this property has a type of
	// "Accept". When true, use the GetActivityStreamsAccept and
	// SetActivityStreamsAccept methods to access and set this property.
	IsActivityStreamsAccept() bool
	// IsActivityStreamsActivity returns true if this property has a type of
	// "Activity". When true, use the GetActivityStreamsActivity and
	// SetActivityStreamsActivity methods to access and set this property.
	IsActivityStreamsActivity() bool
	// IsActivityStreamsAdd returns true if this property has a type of "Add".
	// When true, use the GetActivityStreamsAdd and SetActivityStreamsAdd
	// methods to access and set this property.
	IsActivityStreamsAdd() bool
	// IsActivityStreamsAnnounce returns true if this property has a type of
	// "Announce". When true, use the GetActivityStreamsAnnounce and
	// SetActivityStreamsAnnounce methods to access and set this property.
	IsActivityStreamsAnnounce() bool
	// IsActivityStreamsApplication returns true if this property has a type
	// of "Application". When true, use the GetActivityStreamsApplication
	// and SetActivityStreamsApplication methods to access and set this
	// property.
	IsActivityStreamsApplication() bool
	// IsActivityStreamsArrive returns true if this property has a type of
	// "Arrive". When true, use the GetActivityStreamsArrive and
	// SetActivityStreamsArrive methods to access and set this property.
	IsActivityStreamsArrive() bool
	// IsActivityStreamsArticle returns true if this property has a type of
	// "Article". When true, use the GetActivityStreamsArticle and
	// SetActivityStreamsArticle methods to access and set this property.
	IsActivityStreamsArticle() bool
	// IsActivityStreamsAudio returns true if this property has a type of
	// "Audio". When true, use the GetActivityStreamsAudio and
	// SetActivityStreamsAudio methods to access and set this property.
	IsActivityStreamsAudio() bool
	// IsActivityStreamsBlock returns true if this property has a type of
	// "Block". When true, use the GetActivityStreamsBlock and
	// SetActivityStreamsBlock methods to access and set this property.
	IsActivityStreamsBlock() bool
	// IsActivityStreamsCollection returns true if this property has a type of
	// "Collection". When true, use the GetActivityStreamsCollection and
	// SetActivityStreamsCollection methods to access and set this
	// property.
	IsActivityStreamsCollection() bool
	// IsActivityStreamsCollectionPage returns true if this property has a
	// type of "CollectionPage". When true, use the
	// GetActivityStreamsCollectionPage and
	// SetActivityStreamsCollectionPage methods to access and set this
	// property.
	IsActivityStreamsCollectionPage() bool
	// IsActivityStreamsCreate returns true if this property has a type of
	// "Create". When true, use the GetActivityStreamsCreate and
	// SetActivityStreamsCreate methods to access and set this property.
	IsActivityStreamsCreate() bool
	// IsActivityStreamsDelete returns true if this property has a type of
	// "Delete". When true, use the GetActivityStreamsDelete and
	// SetActivityStreamsDelete methods to access and set this property.
	IsActivityStreamsDelete() bool
	// IsActivityStreamsDislike returns true if this property has a type of
	// "Dislike". When true, use the GetActivityStreamsDislike and
	// SetActivityStreamsDislike methods to access and set this property.
	IsActivityStreamsDislike() bool
	// IsActivityStreamsDocument returns true if this property has a type of
	// "Document". When true, use the GetActivityStreamsDocument and
	// SetActivityStreamsDocument methods to access and set this property.
	IsActivityStreamsDocument() bool
	// IsActivityStreamsEvent returns true if this property has a type of
	// "Event". When true, use the GetActivityStreamsEvent and
	// SetActivityStreamsEvent methods to access and set this property.
	IsActivityStreamsEvent() bool
	// IsActivityStreamsFlag returns true if this property has a type of
	// "Flag". When true, use the GetActivityStreamsFlag and
	// SetActivityStreamsFlag methods to access and set this property.
	IsActivityStreamsFlag() bool
	// IsActivityStreamsFollow returns true if this property has a type of
	// "Follow". When true, use the GetActivityStreamsFollow and
	// SetActivityStreamsFollow methods to access and set this property.
	IsActivityStreamsFollow() bool
	// IsActivityStreamsGroup returns true if this property has a type of
	// "Group". When true, use the GetActivityStreamsGroup and
	// SetActivityStreamsGroup methods to access and set this property.
	IsActivityStreamsGroup() bool
	// IsActivityStreamsIgnore returns true if this property has a type of
	// "Ignore". When true, use the GetActivityStreamsIgnore and
	// SetActivityStreamsIgnore methods to access and set this property.
	IsActivityStreamsIgnore() bool
	// IsActivityStreamsImage returns true if this property has a type of
	// "Image". When true, use the GetActivityStreamsImage and
	// SetActivityStreamsImage methods to access and set this property.
	IsActivityStreamsImage() bool
	// IsActivityStreamsIntransitiveActivity returns true if this property has
	// a type of "IntransitiveActivity". When true, use the
	// GetActivityStreamsIntransitiveActivity and
	// SetActivityStreamsIntransitiveActivity methods to access and set
	// this property.
	IsActivityStreamsIntransitiveActivity() bool
	// IsActivityStreamsInvite returns true if this property has a type of
	// "Invite". When true, use the GetActivityStreamsInvite and
	// SetActivityStreamsInvite methods to access and set this property.
	IsActivityStreamsInvite() bool
	// IsActivityStreamsJoin returns true if this property has a type of
	// "Join". When true, use the GetActivityStreamsJoin and
	// SetActivityStreamsJoin methods to access and set this property.
	IsActivityStreamsJoin() bool
	// IsActivityStreamsLeave returns true if this property has a type of
	// "Leave". When true, use the GetActivityStreamsLeave and
	// SetActivityStreamsLeave methods to access and set this property.
	IsActivityStreamsLeave() bool
	// IsActivityStreamsLike returns true if this property has a type of
	// "Like". When true, use the GetActivityStreamsLike and
	// SetActivityStreamsLike methods to access and set this property.
	IsActivityStreamsLike() bool
	// IsActivityStreamsLink returns true if this property has a type of
	// "Link". When true, use the GetActivityStreamsLink and
	// SetActivityStreamsLink methods to access and set this property.
	IsActivityStreamsLink() bool
	// IsActivityStreamsListen returns true if this property has a type of
	// "Listen". When true, use the GetActivityStreamsListen and
	// SetActivityStreamsListen methods to access and set this property.
	IsActivityStreamsListen() bool
	// IsActivityStreamsMention returns true if this property has a type of
	// "Mention". When true, use the GetActivityStreamsMention and
	// SetActivityStreamsMention methods to access and set this property.
	IsActivityStreamsMention() bool
	// IsActivityStreamsMove returns true if this property has a type of
	// "Move". When true, use the GetActivityStreamsMove and
	// SetActivityStreamsMove methods to access and set this property.
	IsActivityStreamsMove() bool
	// IsActivityStreamsNote returns true if this property has a type of
	// "Note". When true, use the GetActivityStreamsNote and
	// SetActivityStreamsNote methods to access and set this property.
	IsActivityStreamsNote() bool
	// IsActivityStreamsObject returns true if this property has a type of
	// "Object". When true, use the GetActivityStreamsObject and
	// SetActivityStreamsObject methods to access and set this property.
	IsActivityStreamsObject() bool
	// IsActivityStreamsOffer returns true if this property has a type of
	// "Offer". When true, use the GetActivityStreamsOffer and
	// SetActivityStreamsOffer methods to access and set this property.
	IsActivityStreamsOffer() bool
	// IsActivityStreamsOrderedCollection returns true if this property has a
	// type of "OrderedCollection". When true, use the
	// GetActivityStreamsOrderedCollection and
	// SetActivityStreamsOrderedCollection methods to access and set this
	// property.
	IsActivityStreamsOrderedCollection() bool
	// IsActivityStreamsOrderedCollectionPage returns true if this property
	// has a type of "OrderedCollectionPage". When true, use the
	// GetActivityStreamsOrderedCollectionPage and
	// SetActivityStreamsOrderedCollectionPage methods to access and set
	// this property.
	IsActivityStreamsOrderedCollectionPage() bool
	// IsActivityStreamsOrganization returns true if this property has a type
	// of "Organization". When true, use the
	// GetActivityStreamsOrganization and SetActivityStreamsOrganization
	// methods to access and set this property.
	IsActivityStreamsOrganization() bool
	// IsActivityStreamsPage returns true if this property has a type of
	// "Page". When true, use the GetActivityStreamsPage and
	// SetActivityStreamsPage methods to access and set this property.
	IsActivityStreamsPage() bool
	// IsActivityStreamsPerson returns true if this property has a type of
	// "Person". When true, use the GetActivityStreamsPerson and
	// SetActivityStreamsPerson methods to access and set this property.
	IsActivityStreamsPerson() bool
	// IsActivityStreamsPlace returns true if this property has a type of
	// "Place". When true, use the GetActivityStreamsPlace and
	// SetActivityStreamsPlace methods to access and set this property.
	IsActivityStreamsPlace() bool
	// IsActivityStreamsProfile returns true if this property has a type of
	// "Profile". When true, use the GetActivityStreamsProfile and
	// SetActivityStreamsProfile methods to access and set this property.
	IsActivityStreamsProfile() bool
	// IsActivityStreamsQuestion returns true if this property has a type of
	// "Question". When true, use the GetActivityStreamsQuestion and
	// SetActivityStreamsQuestion methods to access and set this property.
	IsActivityStreamsQuestion() bool
	// IsActivityStreamsRead returns true if this property has a type of
	// "Read". When true, use the GetActivityStreamsRead and
	// SetActivityStreamsRead methods to access and set this property.
	IsActivityStreamsRead() bool
	// IsActivityStreamsReject returns true if this property has a type of
	// "Reject". When true, use the GetActivityStreamsReject and
	// SetActivityStreamsReject methods to access and set this property.
	IsActivityStreamsReject() bool
	// IsActivityStreamsRelationship returns true if this property has a type
	// of "Relationship". When true, use the
	// GetActivityStreamsRelationship and SetActivityStreamsRelationship
	// methods to access and set this property.
	IsActivityStreamsRelationship() bool
	// IsActivityStreamsRemove returns true if this property has a type of
	// "Remove". When true, use the GetActivityStreamsRemove and
	// SetActivityStreamsRemove methods to access and set this property.
	IsActivityStreamsRemove() bool
	// IsActivityStreamsService returns true if this property has a type of
	// "Service". When true, use the GetActivityStreamsService and
	// SetActivityStreamsService methods to access and set this property.
	IsActivityStreamsService() bool
	// IsActivityStreamsTentativeAccept returns true if this property has a
	// type of "TentativeAccept". When true, use the
	// GetActivityStreamsTentativeAccept and
	// SetActivityStreamsTentativeAccept methods to access and set this
	// property.
	IsActivityStreamsTentativeAccept() bool
	// IsActivityStreamsTentativeReject returns true if this property has a
	// type of "TentativeReject". When true, use the
	// GetActivityStreamsTentativeReject and
	// SetActivityStreamsTentativeReject methods to access and set this
	// property.
	IsActivityStreamsTentativeReject() bool
	// IsActivityStreamsTombstone returns true if this property has a type of
	// "Tombstone". When true, use the GetActivityStreamsTombstone and
	// SetActivityStreamsTombstone methods to access and set this property.
	IsActivityStreamsTombstone() bool
	// IsActivityStreamsTravel returns true if this property has a type of
	// "Travel". When true, use the GetActivityStreamsTravel and
	// SetActivityStreamsTravel methods to access and set this property.
	IsActivityStreamsTravel() bool
	// IsActivityStreamsUndo returns true if this property has a type of
	// "Undo". When true, use the GetActivityStreamsUndo and
	// SetActivityStreamsUndo methods to access and set this property.
	IsActivityStreamsUndo() bool
	// IsActivityStreamsUpdate returns true if this property has a type of
	// "Update". When true, use the GetActivityStreamsUpdate and
	// SetActivityStreamsUpdate methods to access and set this property.
	IsActivityStreamsUpdate() bool
	// IsActivityStreamsVideo returns true if this property has a type of
	// "Video". When true, use the GetActivityStreamsVideo and
	// SetActivityStreamsVideo methods to access and set this property.
	IsActivityStreamsVideo() bool
	// IsActivityStreamsView returns true if this property has a type of
	// "View". When true, use the GetActivityStreamsView and
	// SetActivityStreamsView methods to access and set this property.
	IsActivityStreamsView() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsObjectPropertyIterator) bool
	// Name returns the name of this property: "ActivityStreamsObject".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() ActivityStreamsObjectPropertyIterator
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() ActivityStreamsObjectPropertyIterator
	// SetActivityStreamsAccept sets the value of this property. Calling
	// IsActivityStreamsAccept afterwards returns true.
	SetActivityStreamsAccept(v ActivityStreamsAccept)
	// SetActivityStreamsActivity sets the value of this property. Calling
	// IsActivityStreamsActivity afterwards returns true.
	SetActivityStreamsActivity(v ActivityStreamsActivity)
	// SetActivityStreamsAdd sets the value of this property. Calling
	// IsActivityStreamsAdd afterwards returns true.
	SetActivityStreamsAdd(v ActivityStreamsAdd)
	// SetActivityStreamsAnnounce sets the value of this property. Calling
	// IsActivityStreamsAnnounce afterwards returns true.
	SetActivityStreamsAnnounce(v ActivityStreamsAnnounce)
	// SetActivityStreamsApplication sets the value of this property. Calling
	// IsActivityStreamsApplication afterwards returns true.
	SetActivityStreamsApplication(v ActivityStreamsApplication)
	// SetActivityStreamsArrive sets the value of this property. Calling
	// IsActivityStreamsArrive afterwards returns true.
	SetActivityStreamsArrive(v ActivityStreamsArrive)
	// SetActivityStreamsArticle sets the value of this property. Calling
	// IsActivityStreamsArticle afterwards returns true.
	SetActivityStreamsArticle(v ActivityStreamsArticle)
	// SetActivityStreamsAudio sets the value of this property. Calling
	// IsActivityStreamsAudio afterwards returns true.
	SetActivityStreamsAudio(v ActivityStreamsAudio)
	// SetActivityStreamsBlock sets the value of this property. Calling
	// IsActivityStreamsBlock afterwards returns true.
	SetActivityStreamsBlock(v ActivityStreamsBlock)
	// SetActivityStreamsCollection sets the value of this property. Calling
	// IsActivityStreamsCollection afterwards returns true.
	SetActivityStreamsCollection(v ActivityStreamsCollection)
	// SetActivityStreamsCollectionPage sets the value of this property.
	// Calling IsActivityStreamsCollectionPage afterwards returns true.
	SetActivityStreamsCollectionPage(v ActivityStreamsCollectionPage)
	// SetActivityStreamsCreate sets the value of this property. Calling
	// IsActivityStreamsCreate afterwards returns true.
	SetActivityStreamsCreate(v ActivityStreamsCreate)
	// SetActivityStreamsDelete sets the value of this property. Calling
	// IsActivityStreamsDelete afterwards returns true.
	SetActivityStreamsDelete(v ActivityStreamsDelete)
	// SetActivityStreamsDislike sets the value of this property. Calling
	// IsActivityStreamsDislike afterwards returns true.
	SetActivityStreamsDislike(v ActivityStreamsDislike)
	// SetActivityStreamsDocument sets the value of this property. Calling
	// IsActivityStreamsDocument afterwards returns true.
	SetActivityStreamsDocument(v ActivityStreamsDocument)
	// SetActivityStreamsEvent sets the value of this property. Calling
	// IsActivityStreamsEvent afterwards returns true.
	SetActivityStreamsEvent(v ActivityStreamsEvent)
	// SetActivityStreamsFlag sets the value of this property. Calling
	// IsActivityStreamsFlag afterwards returns true.
	SetActivityStreamsFlag(v ActivityStreamsFlag)
	// SetActivityStreamsFollow sets the value of this property. Calling
	// IsActivityStreamsFollow afterwards returns true.
	SetActivityStreamsFollow(v ActivityStreamsFollow)
	// SetActivityStreamsGroup sets the value of this property. Calling
	// IsActivityStreamsGroup afterwards returns true.
	SetActivityStreamsGroup(v ActivityStreamsGroup)
	// SetActivityStreamsIgnore sets the value of this property. Calling
	// IsActivityStreamsIgnore afterwards returns true.
	SetActivityStreamsIgnore(v ActivityStreamsIgnore)
	// SetActivityStreamsImage sets the value of this property. Calling
	// IsActivityStreamsImage afterwards returns true.
	SetActivityStreamsImage(v ActivityStreamsImage)
	// SetActivityStreamsIntransitiveActivity sets the value of this property.
	// Calling IsActivityStreamsIntransitiveActivity afterwards returns
	// true.
	SetActivityStreamsIntransitiveActivity(v ActivityStreamsIntransitiveActivity)
	// SetActivityStreamsInvite sets the value of this property. Calling
	// IsActivityStreamsInvite afterwards returns true.
	SetActivityStreamsInvite(v ActivityStreamsInvite)
	// SetActivityStreamsJoin sets the value of this property. Calling
	// IsActivityStreamsJoin afterwards returns true.
	SetActivityStreamsJoin(v ActivityStreamsJoin)
	// SetActivityStreamsLeave sets the value of this property. Calling
	// IsActivityStreamsLeave afterwards returns true.
	SetActivityStreamsLeave(v ActivityStreamsLeave)
	// SetActivityStreamsLike sets the value of this property. Calling
	// IsActivityStreamsLike afterwards returns true.
	SetActivityStreamsLike(v ActivityStreamsLike)
	// SetActivityStreamsLink sets the value of this property. Calling
	// IsActivityStreamsLink afterwards returns true.
	SetActivityStreamsLink(v ActivityStreamsLink)
	// SetActivityStreamsListen sets the value of this property. Calling
	// IsActivityStreamsListen afterwards returns true.
	SetActivityStreamsListen(v ActivityStreamsListen)
	// SetActivityStreamsMention sets the value of this property. Calling
	// IsActivityStreamsMention afterwards returns true.
	SetActivityStreamsMention(v ActivityStreamsMention)
	// SetActivityStreamsMove sets the value of this property. Calling
	// IsActivityStreamsMove afterwards returns true.
	SetActivityStreamsMove(v ActivityStreamsMove)
	// SetActivityStreamsNote sets the value of this property. Calling
	// IsActivityStreamsNote afterwards returns true.
	SetActivityStreamsNote(v ActivityStreamsNote)
	// SetActivityStreamsObject sets the value of this property. Calling
	// IsActivityStreamsObject afterwards returns true.
	SetActivityStreamsObject(v ActivityStreamsObject)
	// SetActivityStreamsOffer sets the value of this property. Calling
	// IsActivityStreamsOffer afterwards returns true.
	SetActivityStreamsOffer(v ActivityStreamsOffer)
	// SetActivityStreamsOrderedCollection sets the value of this property.
	// Calling IsActivityStreamsOrderedCollection afterwards returns true.
	SetActivityStreamsOrderedCollection(v ActivityStreamsOrderedCollection)
	// SetActivityStreamsOrderedCollectionPage sets the value of this
	// property. Calling IsActivityStreamsOrderedCollectionPage afterwards
	// returns true.
	SetActivityStreamsOrderedCollectionPage(v ActivityStreamsOrderedCollectionPage)
	// SetActivityStreamsOrganization sets the value of this property. Calling
	// IsActivityStreamsOrganization afterwards returns true.
	SetActivityStreamsOrganization(v ActivityStreamsOrganization)
	// SetActivityStreamsPage sets the value of this property. Calling
	// IsActivityStreamsPage afterwards returns true.
	SetActivityStreamsPage(v ActivityStreamsPage)
	// SetActivityStreamsPerson sets the value of this property. Calling
	// IsActivityStreamsPerson afterwards returns true.
	SetActivityStreamsPerson(v ActivityStreamsPerson)
	// SetActivityStreamsPlace sets the value of this property. Calling
	// IsActivityStreamsPlace afterwards returns true.
	SetActivityStreamsPlace(v ActivityStreamsPlace)
	// SetActivityStreamsProfile sets the value of this property. Calling
	// IsActivityStreamsProfile afterwards returns true.
	SetActivityStreamsProfile(v ActivityStreamsProfile)
	// SetActivityStreamsQuestion sets the value of this property. Calling
	// IsActivityStreamsQuestion afterwards returns true.
	SetActivityStreamsQuestion(v ActivityStreamsQuestion)
	// SetActivityStreamsRead sets the value of this property. Calling
	// IsActivityStreamsRead afterwards returns true.
	SetActivityStreamsRead(v ActivityStreamsRead)
	// SetActivityStreamsReject sets the value of this property. Calling
	// IsActivityStreamsReject afterwards returns true.
	SetActivityStreamsReject(v ActivityStreamsReject)
	// SetActivityStreamsRelationship sets the value of this property. Calling
	// IsActivityStreamsRelationship afterwards returns true.
	SetActivityStreamsRelationship(v ActivityStreamsRelationship)
	// SetActivityStreamsRemove sets the value of this property. Calling
	// IsActivityStreamsRemove afterwards returns true.
	SetActivityStreamsRemove(v ActivityStreamsRemove)
	// SetActivityStreamsService sets the value of this property. Calling
	// IsActivityStreamsService afterwards returns true.
	SetActivityStreamsService(v ActivityStreamsService)
	// SetActivityStreamsTentativeAccept sets the value of this property.
	// Calling IsActivityStreamsTentativeAccept afterwards returns true.
	SetActivityStreamsTentativeAccept(v ActivityStreamsTentativeAccept)
	// SetActivityStreamsTentativeReject sets the value of this property.
	// Calling IsActivityStreamsTentativeReject afterwards returns true.
	SetActivityStreamsTentativeReject(v ActivityStreamsTentativeReject)
	// SetActivityStreamsTombstone sets the value of this property. Calling
	// IsActivityStreamsTombstone afterwards returns true.
	SetActivityStreamsTombstone(v ActivityStreamsTombstone)
	// SetActivityStreamsTravel sets the value of this property. Calling
	// IsActivityStreamsTravel afterwards returns true.
	SetActivityStreamsTravel(v ActivityStreamsTravel)
	// SetActivityStreamsUndo sets the value of this property. Calling
	// IsActivityStreamsUndo afterwards returns true.
	SetActivityStreamsUndo(v ActivityStreamsUndo)
	// SetActivityStreamsUpdate sets the value of this property. Calling
	// IsActivityStreamsUpdate afterwards returns true.
	SetActivityStreamsUpdate(v ActivityStreamsUpdate)
	// SetActivityStreamsVideo sets the value of this property. Calling
	// IsActivityStreamsVideo afterwards returns true.
	SetActivityStreamsVideo(v ActivityStreamsVideo)
	// SetActivityStreamsView sets the value of this property. Calling
	// IsActivityStreamsView afterwards returns true.
	SetActivityStreamsView(v ActivityStreamsView)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
}

// When used within an Activity, describes the direct object of the activity. For
// instance, in the activity "John added a movie to his wishlist", the object
// of the activity is the movie added. When used within a Relationship
// describes the entity to which the subject is related.
//
// Example 97 (https://www.w3.org/TR/activitystreams-vocabulary/#ex98-jsonld):
//   {
//     "actor": "http://sally.example.org",
//     "object": "http://example.org/posts/1",
//     "summary": "Sally liked a post",
//     "type": "Like"
//   }
//
// Example 98 (https://www.w3.org/TR/activitystreams-vocabulary/#ex99-jsonld):
//   {
//     "actor": "http://sally.example.org",
//     "object": {
//       "content": "A simple note",
//       "type": "Note"
//     },
//     "type": "Like"
//   }
//
// Example 99 (https://www.w3.org/TR/activitystreams-vocabulary/#ex100-jsonld):
//   {
//     "actor": "http://sally.example.org",
//     "object": [
//       "http://example.org/posts/1",
//       {
//         "content": "That is a tree.",
//         "summary": "A simple note",
//         "type": "Note"
//       }
//     ],
//     "summary": "Sally liked a note",
//     "type": "Like"
//   }
type ActivityStreamsObjectProperty interface {
	// AppendActivityStreamsAccept appends a Accept value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsAccept(v ActivityStreamsAccept)
	// AppendActivityStreamsActivity appends a Activity value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsActivity(v ActivityStreamsActivity)
	// AppendActivityStreamsAdd appends a Add value to the back of a list of
	// the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsAdd(v ActivityStreamsAdd)
	// AppendActivityStreamsAnnounce appends a Announce value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsAnnounce(v ActivityStreamsAnnounce)
	// AppendActivityStreamsApplication appends a Application value to the
	// back of a list of the property "object". Invalidates iterators that
	// are traversing using Prev.
	AppendActivityStreamsApplication(v ActivityStreamsApplication)
	// AppendActivityStreamsArrive appends a Arrive value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsArrive(v ActivityStreamsArrive)
	// AppendActivityStreamsArticle appends a Article value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsArticle(v ActivityStreamsArticle)
	// AppendActivityStreamsAudio appends a Audio value to the back of a list
	// of the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsAudio(v ActivityStreamsAudio)
	// AppendActivityStreamsBlock appends a Block value to the back of a list
	// of the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsBlock(v ActivityStreamsBlock)
	// AppendActivityStreamsCollection appends a Collection value to the back
	// of a list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsCollection(v ActivityStreamsCollection)
	// AppendActivityStreamsCollectionPage appends a CollectionPage value to
	// the back of a list of the property "object". Invalidates iterators
	// that are traversing using Prev.
	AppendActivityStreamsCollectionPage(v ActivityStreamsCollectionPage)
	// AppendActivityStreamsCreate appends a Create value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsCreate(v ActivityStreamsCreate)
	// AppendActivityStreamsDelete appends a Delete value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsDelete(v ActivityStreamsDelete)
	// AppendActivityStreamsDislike appends a Dislike value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsDislike(v ActivityStreamsDislike)
	// AppendActivityStreamsDocument appends a Document value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsDocument(v ActivityStreamsDocument)
	// AppendActivityStreamsEvent appends a Event value to the back of a list
	// of the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsEvent(v ActivityStreamsEvent)
	// AppendActivityStreamsFlag appends a Flag value to the back of a list of
	// the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsFlag(v ActivityStreamsFlag)
	// AppendActivityStreamsFollow appends a Follow value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsFollow(v ActivityStreamsFollow)
	// AppendActivityStreamsGroup appends a Group value to the back of a list
	// of the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsGroup(v ActivityStreamsGroup)
	// AppendActivityStreamsIgnore appends a Ignore value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsIgnore(v ActivityStreamsIgnore)
	// AppendActivityStreamsImage appends a Image value to the back of a list
	// of the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsImage(v ActivityStreamsImage)
	// AppendActivityStreamsIntransitiveActivity appends a
	// IntransitiveActivity value to the back of a list of the property
	// "object". Invalidates iterators that are traversing using Prev.
	AppendActivityStreamsIntransitiveActivity(v ActivityStreamsIntransitiveActivity)
	// AppendActivityStreamsInvite appends a Invite value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsInvite(v ActivityStreamsInvite)
	// AppendActivityStreamsJoin appends a Join value to the back of a list of
	// the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsJoin(v ActivityStreamsJoin)
	// AppendActivityStreamsLeave appends a Leave value to the back of a list
	// of the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsLeave(v ActivityStreamsLeave)
	// AppendActivityStreamsLike appends a Like value to the back of a list of
	// the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsLike(v ActivityStreamsLike)
	// AppendActivityStreamsLink appends a Link value to the back of a list of
	// the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsLink(v ActivityStreamsLink)
	// AppendActivityStreamsListen appends a Listen value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsListen(v ActivityStreamsListen)
	// AppendActivityStreamsMention appends a Mention value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsMention(v ActivityStreamsMention)
	// AppendActivityStreamsMove appends a Move value to the back of a list of
	// the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsMove(v ActivityStreamsMove)
	// AppendActivityStreamsNote appends a Note value to the back of a list of
	// the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsNote(v ActivityStreamsNote)
	// AppendActivityStreamsObject appends a Object value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsObject(v ActivityStreamsObject)
	// AppendActivityStreamsOffer appends a Offer value to the back of a list
	// of the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsOffer(v ActivityStreamsOffer)
	// AppendActivityStreamsOrderedCollection appends a OrderedCollection
	// value to the back of a list of the property "object". Invalidates
	// iterators that are traversing using Prev.
	AppendActivityStreamsOrderedCollection(v ActivityStreamsOrderedCollection)
	// AppendActivityStreamsOrderedCollectionPage appends a
	// OrderedCollectionPage value to the back of a list of the property
	// "object". Invalidates iterators that are traversing using Prev.
	AppendActivityStreamsOrderedCollectionPage(v ActivityStreamsOrderedCollectionPage)
	// AppendActivityStreamsOrganization appends a Organization value to the
	// back of a list of the property "object". Invalidates iterators that
	// are traversing using Prev.
	AppendActivityStreamsOrganization(v ActivityStreamsOrganization)
	// AppendActivityStreamsPage appends a Page value to the back of a list of
	// the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsPage(v ActivityStreamsPage)
	// AppendActivityStreamsPerson appends a Person value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsPerson(v ActivityStreamsPerson)
	// AppendActivityStreamsPlace appends a Place value to the back of a list
	// of the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsPlace(v ActivityStreamsPlace)
	// AppendActivityStreamsProfile appends a Profile value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsProfile(v ActivityStreamsProfile)
	// AppendActivityStreamsQuestion appends a Question value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsQuestion(v ActivityStreamsQuestion)
	// AppendActivityStreamsRead appends a Read value to the back of a list of
	// the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsRead(v ActivityStreamsRead)
	// AppendActivityStreamsReject appends a Reject value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsReject(v ActivityStreamsReject)
	// AppendActivityStreamsRelationship appends a Relationship value to the
	// back of a list of the property "object". Invalidates iterators that
	// are traversing using Prev.
	AppendActivityStreamsRelationship(v ActivityStreamsRelationship)
	// AppendActivityStreamsRemove appends a Remove value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsRemove(v ActivityStreamsRemove)
	// AppendActivityStreamsService appends a Service value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsService(v ActivityStreamsService)
	// AppendActivityStreamsTentativeAccept appends a TentativeAccept value to
	// the back of a list of the property "object". Invalidates iterators
	// that are traversing using Prev.
	AppendActivityStreamsTentativeAccept(v ActivityStreamsTentativeAccept)
	// AppendActivityStreamsTentativeReject appends a TentativeReject value to
	// the back of a list of the property "object". Invalidates iterators
	// that are traversing using Prev.
	AppendActivityStreamsTentativeReject(v ActivityStreamsTentativeReject)
	// AppendActivityStreamsTombstone appends a Tombstone value to the back of
	// a list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsTombstone(v ActivityStreamsTombstone)
	// AppendActivityStreamsTravel appends a Travel value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsTravel(v ActivityStreamsTravel)
	// AppendActivityStreamsUndo appends a Undo value to the back of a list of
	// the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsUndo(v ActivityStreamsUndo)
	// AppendActivityStreamsUpdate appends a Update value to the back of a
	// list of the property "object". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsUpdate(v ActivityStreamsUpdate)
	// AppendActivityStreamsVideo appends a Video value to the back of a list
	// of the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsVideo(v ActivityStreamsVideo)
	// AppendActivityStreamsView appends a View value to the back of a list of
	// the property "object". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsView(v ActivityStreamsView)
	// AppendIRI appends an IRI value to the back of a list of the property
	// "object"
	AppendIRI(v *url.URL)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) ActivityStreamsObjectPropertyIterator
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() ActivityStreamsObjectPropertyIterator
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() ActivityStreamsObjectPropertyIterator
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API method specifically needed only for alternate
	// implementations for go-fed. Applications should not use this
	// method. Panics if the index is out of bounds.
	KindIndex(idx int) int
	// Len returns the number of values that exist for the "object" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsObjectProperty) bool
	// Name returns the name of this property: "object".
	Name() string
	// PrependActivityStreamsAccept prepends a Accept value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsAccept(v ActivityStreamsAccept)
	// PrependActivityStreamsActivity prepends a Activity value to the front
	// of a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsActivity(v ActivityStreamsActivity)
	// PrependActivityStreamsAdd prepends a Add value to the front of a list
	// of the property "object". Invalidates all iterators.
	PrependActivityStreamsAdd(v ActivityStreamsAdd)
	// PrependActivityStreamsAnnounce prepends a Announce value to the front
	// of a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsAnnounce(v ActivityStreamsAnnounce)
	// PrependActivityStreamsApplication prepends a Application value to the
	// front of a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsApplication(v ActivityStreamsApplication)
	// PrependActivityStreamsArrive prepends a Arrive value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsArrive(v ActivityStreamsArrive)
	// PrependActivityStreamsArticle prepends a Article value to the front of
	// a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsArticle(v ActivityStreamsArticle)
	// PrependActivityStreamsAudio prepends a Audio value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsAudio(v ActivityStreamsAudio)
	// PrependActivityStreamsBlock prepends a Block value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsBlock(v ActivityStreamsBlock)
	// PrependActivityStreamsCollection prepends a Collection value to the
	// front of a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsCollection(v ActivityStreamsCollection)
	// PrependActivityStreamsCollectionPage prepends a CollectionPage value to
	// the front of a list of the property "object". Invalidates all
	// iterators.
	PrependActivityStreamsCollectionPage(v ActivityStreamsCollectionPage)
	// PrependActivityStreamsCreate prepends a Create value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsCreate(v ActivityStreamsCreate)
	// PrependActivityStreamsDelete prepends a Delete value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsDelete(v ActivityStreamsDelete)
	// PrependActivityStreamsDislike prepends a Dislike value to the front of
	// a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsDislike(v ActivityStreamsDislike)
	// PrependActivityStreamsDocument prepends a Document value to the front
	// of a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsDocument(v ActivityStreamsDocument)
	// PrependActivityStreamsEvent prepends a Event value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsEvent(v ActivityStreamsEvent)
	// PrependActivityStreamsFlag prepends a Flag value to the front of a list
	// of the property "object". Invalidates all iterators.
	PrependActivityStreamsFlag(v ActivityStreamsFlag)
	// PrependActivityStreamsFollow prepends a Follow value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsFollow(v ActivityStreamsFollow)
	// PrependActivityStreamsGroup prepends a Group value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsGroup(v ActivityStreamsGroup)
	// PrependActivityStreamsIgnore prepends a Ignore value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsIgnore(v ActivityStreamsIgnore)
	// PrependActivityStreamsImage prepends a Image value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsImage(v ActivityStreamsImage)
	// PrependActivityStreamsIntransitiveActivity prepends a
	// IntransitiveActivity value to the front of a list of the property
	// "object". Invalidates all iterators.
	PrependActivityStreamsIntransitiveActivity(v ActivityStreamsIntransitiveActivity)
	// PrependActivityStreamsInvite prepends a Invite value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsInvite(v ActivityStreamsInvite)
	// PrependActivityStreamsJoin prepends a Join value to the front of a list
	// of the property "object". Invalidates all iterators.
	PrependActivityStreamsJoin(v ActivityStreamsJoin)
	// PrependActivityStreamsLeave prepends a Leave value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsLeave(v ActivityStreamsLeave)
	// PrependActivityStreamsLike prepends a Like value to the front of a list
	// of the property "object". Invalidates all iterators.
	PrependActivityStreamsLike(v ActivityStreamsLike)
	// PrependActivityStreamsLink prepends a Link value to the front of a list
	// of the property "object". Invalidates all iterators.
	PrependActivityStreamsLink(v ActivityStreamsLink)
	// PrependActivityStreamsListen prepends a Listen value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsListen(v ActivityStreamsListen)
	// PrependActivityStreamsMention prepends a Mention value to the front of
	// a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsMention(v ActivityStreamsMention)
	// PrependActivityStreamsMove prepends a Move value to the front of a list
	// of the property "object". Invalidates all iterators.
	PrependActivityStreamsMove(v ActivityStreamsMove)
	// PrependActivityStreamsNote prepends a Note value to the front of a list
	// of the property "object". Invalidates all iterators.
	PrependActivityStreamsNote(v ActivityStreamsNote)
	// PrependActivityStreamsObject prepends a Object value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsObject(v ActivityStreamsObject)
	// PrependActivityStreamsOffer prepends a Offer value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsOffer(v ActivityStreamsOffer)
	// PrependActivityStreamsOrderedCollection prepends a OrderedCollection
	// value to the front of a list of the property "object". Invalidates
	// all iterators.
	PrependActivityStreamsOrderedCollection(v ActivityStreamsOrderedCollection)
	// PrependActivityStreamsOrderedCollectionPage prepends a
	// OrderedCollectionPage value to the front of a list of the property
	// "object". Invalidates all iterators.
	PrependActivityStreamsOrderedCollectionPage(v ActivityStreamsOrderedCollectionPage)
	// PrependActivityStreamsOrganization prepends a Organization value to the
	// front of a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsOrganization(v ActivityStreamsOrganization)
	// PrependActivityStreamsPage prepends a Page value to the front of a list
	// of the property "object". Invalidates all iterators.
	PrependActivityStreamsPage(v ActivityStreamsPage)
	// PrependActivityStreamsPerson prepends a Person value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsPerson(v ActivityStreamsPerson)
	// PrependActivityStreamsPlace prepends a Place value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsPlace(v ActivityStreamsPlace)
	// PrependActivityStreamsProfile prepends a Profile value to the front of
	// a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsProfile(v ActivityStreamsProfile)
	// PrependActivityStreamsQuestion prepends a Question value to the front
	// of a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsQuestion(v ActivityStreamsQuestion)
	// PrependActivityStreamsRead prepends a Read value to the front of a list
	// of the property "object". Invalidates all iterators.
	PrependActivityStreamsRead(v ActivityStreamsRead)
	// PrependActivityStreamsReject prepends a Reject value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsReject(v ActivityStreamsReject)
	// PrependActivityStreamsRelationship prepends a Relationship value to the
	// front of a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsRelationship(v ActivityStreamsRelationship)
	// PrependActivityStreamsRemove prepends a Remove value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsRemove(v ActivityStreamsRemove)
	// PrependActivityStreamsService prepends a Service value to the front of
	// a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsService(v ActivityStreamsService)
	// PrependActivityStreamsTentativeAccept prepends a TentativeAccept value
	// to the front of a list of the property "object". Invalidates all
	// iterators.
	PrependActivityStreamsTentativeAccept(v ActivityStreamsTentativeAccept)
	// PrependActivityStreamsTentativeReject prepends a TentativeReject value
	// to the front of a list of the property "object". Invalidates all
	// iterators.
	PrependActivityStreamsTentativeReject(v ActivityStreamsTentativeReject)
	// PrependActivityStreamsTombstone prepends a Tombstone value to the front
	// of a list of the property "object". Invalidates all iterators.
	PrependActivityStreamsTombstone(v ActivityStreamsTombstone)
	// PrependActivityStreamsTravel prepends a Travel value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsTravel(v ActivityStreamsTravel)
	// PrependActivityStreamsUndo prepends a Undo value to the front of a list
	// of the property "object". Invalidates all iterators.
	PrependActivityStreamsUndo(v ActivityStreamsUndo)
	// PrependActivityStreamsUpdate prepends a Update value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsUpdate(v ActivityStreamsUpdate)
	// PrependActivityStreamsVideo prepends a Video value to the front of a
	// list of the property "object". Invalidates all iterators.
	PrependActivityStreamsVideo(v ActivityStreamsVideo)
	// PrependActivityStreamsView prepends a View value to the front of a list
	// of the property "object". Invalidates all iterators.
	PrependActivityStreamsView(v ActivityStreamsView)
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "object".
	PrependIRI(v *url.URL)
	// Remove deletes an element at the specified index from a list of the
	// property "object", regardless of its type. Panics if the index is
	// out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetActivityStreamsAccept sets a Accept value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsAccept(idx int, v ActivityStreamsAccept)
	// SetActivityStreamsActivity sets a Activity value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsActivity(idx int, v ActivityStreamsActivity)
	// SetActivityStreamsAdd sets a Add value to be at the specified index for
	// the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsAdd(idx int, v ActivityStreamsAdd)
	// SetActivityStreamsAnnounce sets a Announce value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsAnnounce(idx int, v ActivityStreamsAnnounce)
	// SetActivityStreamsApplication sets a Application value to be at the
	// specified index for the property "object". Panics if the index is
	// out of bounds. Invalidates all iterators.
	SetActivityStreamsApplication(idx int, v ActivityStreamsApplication)
	// SetActivityStreamsArrive sets a Arrive value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsArrive(idx int, v ActivityStreamsArrive)
	// SetActivityStreamsArticle sets a Article value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsArticle(idx int, v ActivityStreamsArticle)
	// SetActivityStreamsAudio sets a Audio value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsAudio(idx int, v ActivityStreamsAudio)
	// SetActivityStreamsBlock sets a Block value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsBlock(idx int, v ActivityStreamsBlock)
	// SetActivityStreamsCollection sets a Collection value to be at the
	// specified index for the property "object". Panics if the index is
	// out of bounds. Invalidates all iterators.
	SetActivityStreamsCollection(idx int, v ActivityStreamsCollection)
	// SetActivityStreamsCollectionPage sets a CollectionPage value to be at
	// the specified index for the property "object". Panics if the index
	// is out of bounds. Invalidates all iterators.
	SetActivityStreamsCollectionPage(idx int, v ActivityStreamsCollectionPage)
	// SetActivityStreamsCreate sets a Create value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsCreate(idx int, v ActivityStreamsCreate)
	// SetActivityStreamsDelete sets a Delete value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsDelete(idx int, v ActivityStreamsDelete)
	// SetActivityStreamsDislike sets a Dislike value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsDislike(idx int, v ActivityStreamsDislike)
	// SetActivityStreamsDocument sets a Document value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsDocument(idx int, v ActivityStreamsDocument)
	// SetActivityStreamsEvent sets a Event value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsEvent(idx int, v ActivityStreamsEvent)
	// SetActivityStreamsFlag sets a Flag value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsFlag(idx int, v ActivityStreamsFlag)
	// SetActivityStreamsFollow sets a Follow value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsFollow(idx int, v ActivityStreamsFollow)
	// SetActivityStreamsGroup sets a Group value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsGroup(idx int, v ActivityStreamsGroup)
	// SetActivityStreamsIgnore sets a Ignore value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsIgnore(idx int, v ActivityStreamsIgnore)
	// SetActivityStreamsImage sets a Image value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsImage(idx int, v ActivityStreamsImage)
	// SetActivityStreamsIntransitiveActivity sets a IntransitiveActivity
	// value to be at the specified index for the property "object".
	// Panics if the index is out of bounds. Invalidates all iterators.
	SetActivityStreamsIntransitiveActivity(idx int, v ActivityStreamsIntransitiveActivity)
	// SetActivityStreamsInvite sets a Invite value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsInvite(idx int, v ActivityStreamsInvite)
	// SetActivityStreamsJoin sets a Join value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsJoin(idx int, v ActivityStreamsJoin)
	// SetActivityStreamsLeave sets a Leave value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsLeave(idx int, v ActivityStreamsLeave)
	// SetActivityStreamsLike sets a Like value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsLike(idx int, v ActivityStreamsLike)
	// SetActivityStreamsLink sets a Link value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsLink(idx int, v ActivityStreamsLink)
	// SetActivityStreamsListen sets a Listen value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsListen(idx int, v ActivityStreamsListen)
	// SetActivityStreamsMention sets a Mention value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsMention(idx int, v ActivityStreamsMention)
	// SetActivityStreamsMove sets a Move value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsMove(idx int, v ActivityStreamsMove)
	// SetActivityStreamsNote sets a Note value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsNote(idx int, v ActivityStreamsNote)
	// SetActivityStreamsObject sets a Object value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsObject(idx int, v ActivityStreamsObject)
	// SetActivityStreamsOffer sets a Offer value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsOffer(idx int, v ActivityStreamsOffer)
	// SetActivityStreamsOrderedCollection sets a OrderedCollection value to
	// be at the specified index for the property "object". Panics if the
	// index is out of bounds. Invalidates all iterators.
	SetActivityStreamsOrderedCollection(idx int, v ActivityStreamsOrderedCollection)
	// SetActivityStreamsOrderedCollectionPage sets a OrderedCollectionPage
	// value to be at the specified index for the property "object".
	// Panics if the index is out of bounds. Invalidates all iterators.
	SetActivityStreamsOrderedCollectionPage(idx int, v ActivityStreamsOrderedCollectionPage)
	// SetActivityStreamsOrganization sets a Organization value to be at the
	// specified index for the property "object". Panics if the index is
	// out of bounds. Invalidates all iterators.
	SetActivityStreamsOrganization(idx int, v ActivityStreamsOrganization)
	// SetActivityStreamsPage sets a Page value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsPage(idx int, v ActivityStreamsPage)
	// SetActivityStreamsPerson sets a Person value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsPerson(idx int, v ActivityStreamsPerson)
	// SetActivityStreamsPlace sets a Place value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsPlace(idx int, v ActivityStreamsPlace)
	// SetActivityStreamsProfile sets a Profile value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsProfile(idx int, v ActivityStreamsProfile)
	// SetActivityStreamsQuestion sets a Question value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsQuestion(idx int, v ActivityStreamsQuestion)
	// SetActivityStreamsRead sets a Read value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsRead(idx int, v ActivityStreamsRead)
	// SetActivityStreamsReject sets a Reject value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsReject(idx int, v ActivityStreamsReject)
	// SetActivityStreamsRelationship sets a Relationship value to be at the
	// specified index for the property "object". Panics if the index is
	// out of bounds. Invalidates all iterators.
	SetActivityStreamsRelationship(idx int, v ActivityStreamsRelationship)
	// SetActivityStreamsRemove sets a Remove value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsRemove(idx int, v ActivityStreamsRemove)
	// SetActivityStreamsService sets a Service value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsService(idx int, v ActivityStreamsService)
	// SetActivityStreamsTentativeAccept sets a TentativeAccept value to be at
	// the specified index for the property "object". Panics if the index
	// is out of bounds. Invalidates all iterators.
	SetActivityStreamsTentativeAccept(idx int, v ActivityStreamsTentativeAccept)
	// SetActivityStreamsTentativeReject sets a TentativeReject value to be at
	// the specified index for the property "object". Panics if the index
	// is out of bounds. Invalidates all iterators.
	SetActivityStreamsTentativeReject(idx int, v ActivityStreamsTentativeReject)
	// SetActivityStreamsTombstone sets a Tombstone value to be at the
	// specified index for the property "object". Panics if the index is
	// out of bounds. Invalidates all iterators.
	SetActivityStreamsTombstone(idx int, v ActivityStreamsTombstone)
	// SetActivityStreamsTravel sets a Travel value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsTravel(idx int, v ActivityStreamsTravel)
	// SetActivityStreamsUndo sets a Undo value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsUndo(idx int, v ActivityStreamsUndo)
	// SetActivityStreamsUpdate sets a Update value to be at the specified
	// index for the property "object". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsUpdate(idx int, v ActivityStreamsUpdate)
	// SetActivityStreamsVideo sets a Video value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsVideo(idx int, v ActivityStreamsVideo)
	// SetActivityStreamsView sets a View value to be at the specified index
	// for the property "object". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsView(idx int, v ActivityStreamsView)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "object". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// Swap swaps the location of values at two indices for the "object"
	// property.
	Swap(i, j int)
}