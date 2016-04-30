package push

// PushRequest struct represents a message to be pushed to mobile.
type Request struct {
	// push target
	TargetType            string //all, regid, alias, user_account, topic, multi_topic
	TargetValue           string
	RestrictedPackageName string

	// push configuration
	Title       string
	Description string
	Payload     string

	// IOS specific
	ExtraBadge    int
	ExtraCategory int

	// Android specific
	PassThrough           int
	NotifyType            int
	NotifyID              int
	ExtraSourcdURI        string
	ExtraTicker           string
	ExtraNotifyForeground int
	ExtraNotifyEffect     int
	ExtraIntentURI        string
	ExtraWebURI           string
	ExtraFlowControl      int
	ExtraLayoutName       string
	ExtraLayoutValue      string
	ExtraJobkey           string
	ExtraCallback         string
	ExtraLocale           string
	ExtraLocaleNotIn      string
	ExtraModel            string
	ExtraModelNotIn       string
	ExtraAppVersion       string
	ExtraAppVersionNotIn  string
	ExtraConnpt           string

	// Message control
	TimeToSend   int // milliseconds since 1970.1.1 00:00:00 UTC
	TimeToLive   int // milliseconds to keep message
	StoreOffline bool
	ExpireTime   string
	BatchNumber  string
}
