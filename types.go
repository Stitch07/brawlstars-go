package brawlstars

type Player struct {
	Tag                   string      `json:"tag"`
	Name                  string      `json:"name"`
	ID                    TagID       `json:"id"`
	NameColorCode         string      `json:"nameColorCode"`
	BrawlersUnlocked      uint8       `json:"brawlersUnlocked"`
	Victories             uint        `json:"victories"`
	SoloShowdownVictories uint        `json:"soloShowdownVictories"`
	DuoShowdownVictories  uint        `json:"duoShowdownVictories"`
	TotalExp              uint        `json:"totalExp"`
	ExpFmt                string      `json:"expFmt"`
	ExpLevel              uint        `json:"expLevel"`
	Trophies              uint        `json:"trophies"`
	HighestTrophies       uint        `json:"highestTrophies"`
	AvatarID              uint        `json:"avatarId"`
	AvatarURL             string      `json:"avatarUrl"`
	BestTimeAsBigBrawler  string      `json:"bestTimeAsBigBrawler"`
	BestRoboRumbleTime    string      `json:"bestRoboRumbleTime"`
	HasSkins              bool        `json:"hasSkins"`
	Club                  PartialClub `json:"club"`
	Brawlers              []Brawler   `json:"brawler"`
}

type TagID struct {
	High uint `json:"high"`
	Low  uint `json:"low"`
}

type PartialClub struct {
	ID               TagID  `json:"id"`
	Tag              string `json:"tag"`
	Name             string `json:"name"`
	Role             string `json:"role"`
	BadgeID          uint   `json:"badgeId"`
	BadgeURL         string `json:"badgeUrl"`
	Members          uint   `json:"members"`
	Trophies         uint   `json:"trophies"`
	RequiredTrophies uint   `json:"requiredTrophies"`
	OnlineMembers    uint   `json:"onlineMembers"`
}

type Brawler struct {
	Name            string `json:"name"`
	HasSkin         bool   `json:"hasSkin"`
	Skin            string `json:"skin"`
	Trophies        uint   `json:"trophies"`
	HighestTrophies uint   `json:"highestTrophies"`
	Power           uint8  `json:"power"`
	Rank            uint8  `json:"rank"`
}

type PlayerSearch struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type Club struct {
	Tag              string `json:"tag"`
	ID               TagID  `json:"id"`
	Name             string `json:"name"`
	Region           string `json:"region"`
	BadgeID          uint   `json:"badgeId"`
	BadgeURL         string `json:"badgeUrl"`
	Status           string `json:"status"`
	MemberCount      uint   `json:"membersCount"`
	OnlineMembers    uint   `json:"onlineMembers"`
	Trophies         uint   `json:"trophies"`
	RequiredTrophies uint   `json:"requiredTrophies"`
	Description      string `json:"description"`
}

type ClubMember struct {
	Tag  string `json:"tag"`
	ID   TagID  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	//ExpLevel string `json:"expLevel"`
	Trophies                 uint   `json:"trophies"`
	OnlineLessThanOneHourAgo bool   `json:"onlineLessThanOneHourAgo"`
	AvatarID                 uint   `json:"avatarId"`
	NameColorCode            string `json:"nameColorCode"`
	AvatarURL                string `json:"avatarUrl"`
}

type ClubSearch struct {
	Tag              string `json:"tag"`
	ID               TagID  `json:"id"`
	Name             string `json:"name"`
	MemberCount      uint   `json:"membersCount"`
	Trophies         uint   `json:"trophies"`
	RequiredTrophies uint   `json:"requiredTrophies"`
	Type             string `json:"type"`
}

type Event struct {
	Slot               uint8  `json:"slot"`
	SlotName           string `json:"slotName"`
	StartTimeInSeconds uint   `json:"startTimeInSeconds"`
	StartTime          string `json:"startTime"`
	EndTimeInSeconds   uint   `json:"endTimeInSeconds"`
	EndTime            string `json:"endTime"`
	FreeKeys           uint8  `json:"freeKeys"`
	MapID              uint   `json:"mapId"`
	MapName            string `json:"mapName"`
	MapImageURL        string `json:"mapImageURL"`
	GameMode           string `json:"gameMode"`
	HasModifier        bool   `json:"hasModifier"`
	ModifierID         uint   `json:"modifierId"`
	ModifierName       string `json:"modifierName"`
}

type Events struct {
	Current  []*Event `json:"current"`
	Upcoming []*Event `json:"upcoming"`
}

type TopClub struct {
	Tag         string `json:"tag"`
	ID          TagID  `json:"id"`
	Name        string `json:"name"`
	MemberCount uint   `json:"membersCount"`
	Trophies    uint   `json:"trophies"`
	BadgeID     uint   `json:"badgeId"`
	BadgeURL    string `json:"badgeUrl"`
	Position    uint   `json:"position"`
}

type TopPlayer struct {
	Tag           string `json:"tag"`
	ID            TagID  `json:"id"`
	Name          string `json:"name"`
	NameColorCode string `json:"nameColorCode"`
	AvatarID      uint   `json:"avatarId"`
	AvatarURL     string `json:"avatarUrl"`
	Position      uint   `json:"position"`
	Trophies      uint   `json:"trophies"`
	ClubName      string `json:"clubName"`
	ExpLevel      uint   `json:"expLevel"`
}
