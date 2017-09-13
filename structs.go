package main

type XMLStops struct {
	Stops []Stop `xml:" stop,omitempty" json:"stop,omitempty" db:" stop,omitempty"`
}

type Stop struct {
	Entferung  float64 `xml:" entferung,attr"  json:",omitempty"`
	Haltepunkt string  `xml:" haltepunkt,attr"  json:",omitempty"`
	Lat        string  `xml:" lat,attr"  json:",omitempty"`
	Lng        string  `xml:" lng,attr"  json:",omitempty"`
	Name       string  `xml:" name,attr"  json:",omitempty"`
}

type SessionResponse struct {
	SessionID string `xml:"sessionID"`
}

type ItdRequest struct {
	CalcTime                   string                     `xml:" calcTime,attr"  json:",omitempty"`
	Client                     string                     `xml:" client,attr"  json:",omitempty"`
	ClientIP                   string                     `xml:" clientIP,attr"  json:",omitempty"`
	Language                   string                     `xml:" language,attr"  json:",omitempty"`
	LengthUnit                 string                     `xml:" lengthUnit,attr"  json:",omitempty"`
	Now                        string                     `xml:" now,attr"  json:",omitempty"`
	NowWD                      string                     `xml:" nowWD,attr"  json:",omitempty"`
	ServerID                   string                     `xml:" serverID,attr"  json:",omitempty"`
	SessionID                  string                     `xml:" sessionID,attr"  json:",omitempty"`
	Version                    string                     `xml:" version,attr"  json:",omitempty"`
	VirtDir                    string                     `xml:" virtDir,attr"  json:",omitempty"`
	ItdDepartureMonitorRequest ItdDepartureMonitorRequest `xml:" itdDepartureMonitorRequest,omitempty" json:"itdDepartureMonitorRequest,omitempty" db:" itdDepartureMonitorRequest,omitempty"`
	ItdVersionInfo             ItdVersionInfo             `xml:" itdVersionInfo,omitempty" json:"itdVersionInfo,omitempty" db:" itdVersionInfo,omitempty"`
}

type ItdVersionInfo struct {
	PtKernel PtKernel `xml:" ptKernel,omitempty" json:"ptKernel,omitempty" db:" ptKernel,omitempty"`
}

type PtKernel struct {
	AppVersion AppVersion `xml:" appVersion,omitempty" json:"appVersion,omitempty" db:" appVersion,omitempty"`
	DataBuild  DataBuild  `xml:" dataBuild,omitempty" json:"dataBuild,omitempty" db:" dataBuild,omitempty"`
	DataFormat DataFormat `xml:" dataFormat,omitempty" json:"dataFormat,omitempty" db:" dataFormat,omitempty"`
}

type AppVersion struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DataFormat struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DataBuild struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ItdDepartureMonitorRequest struct {
	RequestID        string           `xml:" requestID,attr"  json:",omitempty"`
	ItdDMDateTime    ItdDMDateTime    `xml:" itdDMDateTime,omitempty" json:"itdDMDateTime,omitempty" db:" itdDMDateTime,omitempty"`
	ItdDateRange     ItdDateRange     `xml:" itdDateRange,omitempty" json:"itdDateRange,omitempty" db:" itdDateRange,omitempty"`
	ItdDateTime      ItdDateTime      `xml:" itdDateTime,omitempty" json:"itdDateTime,omitempty" db:" itdDateTime,omitempty"`
	ItdDepartureList ItdDepartureList `xml:" itdDepartureList,omitempty" json:"itdDepartureList,omitempty" db:" itdDepartureList,omitempty"`
	ItdOdv           ItdOdv           `xml:" itdOdv,omitempty" json:"itdOdv,omitempty" db:" itdOdv,omitempty"`
	ItdServingLines  ItdServingLines  `xml:" itdServingLines,omitempty" json:"itdServingLines,omitempty" db:" itdServingLines,omitempty"`
	ItdTripOptions   ItdTripOptions   `xml:" itdTripOptions,omitempty" json:"itdTripOptions,omitempty" db:" itdTripOptions,omitempty"`
}

type ItdOdv struct {
	AnyObjFilter string      `xml:" anyObjFilter,attr"  json:",omitempty"`
	BuildNiveau  string      `xml:" buildNiveau,attr"  json:",omitempty"`
	Type         string      `xml:" type,attr"  json:",omitempty"`
	Usage        string      `xml:" usage,attr"  json:",omitempty"`
	ItdOdvName   ItdOdvName  `xml:" itdOdvName,omitempty" json:"itdOdvName,omitempty" db:" itdOdvName,omitempty"`
	ItdOdvPlace  ItdOdvPlace `xml:" itdOdvPlace,omitempty" json:"itdOdvPlace,omitempty" db:" itdOdvPlace,omitempty"`
}

type ItdOdvPlace struct {
	Method        string        `xml:" method,attr"  json:",omitempty"`
	State         string        `xml:" state,attr"  json:",omitempty"`
	OdvPlaceElem  OdvPlaceElem  `xml:" odvPlaceElem,omitempty" json:"odvPlaceElem,omitempty" db:" odvPlaceElem,omitempty"`
	OdvPlaceInput OdvPlaceInput `xml:" odvPlaceInput,omitempty" json:"odvPlaceInput,omitempty" db:" odvPlaceInput,omitempty"`
}

type OdvPlaceElem struct {
	MainPlace string `xml:" mainPlace,attr"  json:",omitempty"`
	Omc       string `xml:" omc,attr"  json:",omitempty"`
	PlaceID   string `xml:" placeID,attr"  json:",omitempty"`
	Span      string `xml:" span,attr"  json:",omitempty"`
	Stateless string `xml:" stateless,attr"  json:",omitempty"`
	Type      string `xml:" type,attr"  json:",omitempty"`
	Value     string `xml:" value,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type OdvPlaceInput struct {
}

type ItdOdvName struct {
	Method       string       `xml:" method,attr"  json:",omitempty"`
	State        string       `xml:" state,attr"  json:",omitempty"`
	OdvNameElem  OdvNameElem  `xml:" odvNameElem,omitempty" json:"odvNameElem,omitempty" db:" odvNameElem,omitempty"`
	OdvNameInput OdvNameInput `xml:" odvNameInput,omitempty" json:"odvNameInput,omitempty" db:" odvNameInput,omitempty"`
}

type OdvNameElem struct {
	IsTransferStop string         `xml:" isTransferStop,attr"  json:",omitempty"`
	MapName        string         `xml:" mapName,attr"  json:",omitempty"`
	MatchQuality   string         `xml:" matchQuality,attr"  json:",omitempty"`
	Selected       string         `xml:" selected,attr"  json:",omitempty"`
	Stateless      string         `xml:" stateless,attr"  json:",omitempty"`
	StopID         string         `xml:" stopID,attr"  json:",omitempty"`
	Value          string         `xml:" value,attr"  json:",omitempty"`
	X              string         `xml:" x,attr"  json:",omitempty"`
	Y              string         `xml:" y,attr"  json:",omitempty"`
	ItdMapItemList ItdMapItemList `xml:" itdMapItemList,omitempty" json:"itdMapItemList,omitempty" db:" itdMapItemList,omitempty"`
	Text           string         `xml:",chardata" json:",omitempty"`
}

type ItdMapItemList struct {
}

type OdvNameInput struct {
	Text int32 `xml:",chardata" json:",omitempty"`
}

type ItdDateTime struct {
	TtpFrom string    `xml:" ttpFrom,attr"  json:",omitempty"`
	TtpTo   string    `xml:" ttpTo,attr"  json:",omitempty"`
	ItdDate []ItdDate `xml:" itdDate,omitempty" json:"itdDate,omitempty" db:" itdDate,omitempty"`
	ItdTime ItdTime   `xml:" itdTime,omitempty" json:"itdTime,omitempty" db:" itdTime,omitempty"`
}

type ItdDate struct {
	Day     string `xml:" day,attr"  json:",omitempty"`
	Month   string `xml:" month,attr"  json:",omitempty"`
	Weekday string `xml:" weekday,attr"  json:",omitempty"`
	Year    string `xml:" year,attr"  json:",omitempty"`
}

type ItdTime struct {
	Ap     string `xml:" ap,attr"  json:",omitempty"`
	Hour   string `xml:" hour,attr"  json:",omitempty"`
	Minute string `xml:" minute,attr"  json:",omitempty"`
}

type ItdDMDateTime struct {
	Deparr string `xml:" deparr,attr"  json:",omitempty"`
}

type ItdDateRange struct {
	ItdDate []ItdDate `xml:" itdDate,omitempty" json:"itdDate,omitempty" db:" itdDate,omitempty"`
}

type ItdTripOptions struct {
	ItdItOptions   ItdItOptions   `xml:" itdItOptions,omitempty" json:"itdItOptions,omitempty" db:" itdItOptions,omitempty"`
	ItdPtOptions   ItdPtOptions   `xml:" itdPtOptions,omitempty" json:"itdPtOptions,omitempty" db:" itdPtOptions,omitempty"`
	ItdUsedOptions ItdUsedOptions `xml:" itdUsedOptions,omitempty" json:"itdUsedOptions,omitempty" db:" itdUsedOptions,omitempty"`
}

type ItdPtOptions struct {
	Active                       string        `xml:" active,attr"  json:",omitempty"`
	ActiveCom                    string        `xml:" activeCom,attr"  json:",omitempty"`
	ActiveImp                    string        `xml:" activeImp,attr"  json:",omitempty"`
	ActiveSec                    string        `xml:" activeSec,attr"  json:",omitempty"`
	Assistance                   string        `xml:" assistance,attr"  json:",omitempty"`
	Bike                         string        `xml:" bike,attr"  json:",omitempty"`
	ChangeSpeed                  string        `xml:" changeSpeed,attr"  json:",omitempty"`
	IllumTransfer                string        `xml:" illumTransfer,attr"  json:",omitempty"`
	Level                        string        `xml:" level,attr"  json:",omitempty"`
	LineRestriction              string        `xml:" lineRestriction,attr"  json:",omitempty"`
	LowPlatformVhcl              string        `xml:" lowPlatformVhcl,attr"  json:",omitempty"`
	MaxChanges                   string        `xml:" maxChanges,attr"  json:",omitempty"`
	MaxTime                      string        `xml:" maxTime,attr"  json:",omitempty"`
	MaxWait                      string        `xml:" maxWait,attr"  json:",omitempty"`
	NeedElevatedPlt              string        `xml:" needElevatedPlt,attr"  json:",omitempty"`
	NoCrowded                    string        `xml:" noCrowded,attr"  json:",omitempty"`
	NoElevators                  string        `xml:" noElevators,attr"  json:",omitempty"`
	NoEscalators                 string        `xml:" noEscalators,attr"  json:",omitempty"`
	NoInsecurePlaces             string        `xml:" noInsecurePlaces,attr"  json:",omitempty"`
	NoLonelyTransfer             string        `xml:" noLonelyTransfer,attr"  json:",omitempty"`
	NoSolidStairs                string        `xml:" noSolidStairs,attr"  json:",omitempty"`
	OvergroundTransfer           string        `xml:" overgroundTransfer,attr"  json:",omitempty"`
	Plane                        string        `xml:" plane,attr"  json:",omitempty"`
	PrivateTransport             string        `xml:" privateTransport,attr"  json:",omitempty"`
	RouteType                    string        `xml:" routeType,attr"  json:",omitempty"`
	SOSAvail                     string        `xml:" SOSAvail,attr"  json:",omitempty"`
	TwoSidedAlt                  string        `xml:" twoSidedAlt,attr"  json:",omitempty"`
	UseProxFootSearch            string        `xml:" useProxFootSearch,attr"  json:",omitempty"`
	UseProxFootSearchDestination string        `xml:" useProxFootSearchDestination,attr"  json:",omitempty"`
	UseProxFootSearchOrigin      string        `xml:" useProxFootSearchOrigin,attr"  json:",omitempty"`
	Wheelchair                   string        `xml:" wheelchair,attr"  json:",omitempty"`
	ExcludedMeans                ExcludedMeans `xml:" excludedMeans,omitempty" json:"excludedMeans,omitempty" db:" excludedMeans,omitempty"`
}

type ExcludedMeans struct {
	MeansElem []MeansElem `xml:" meansElem,omitempty" json:"meansElem,omitempty" db:" meansElem,omitempty"`
}

type MeansElem struct {
	Selected string `xml:" selected,attr"  json:",omitempty"`
	Value    string `xml:" value,attr"  json:",omitempty"`
	Text     string `xml:",chardata" json:",omitempty"`
}

type ItdItOptions struct {
	DepartureTransport DepartureTransport `xml:" departureTransport,omitempty" json:"departureTransport,omitempty" db:" departureTransport,omitempty"`
	ItBicycle          ItBicycle          `xml:" itBicycle,omitempty" json:"itBicycle,omitempty" db:" itBicycle,omitempty"`
	ItPedestrian       ItPedestrian       `xml:" itPedestrian,omitempty" json:"itPedestrian,omitempty" db:" itPedestrian,omitempty"`
	ItRouter           ItRouter           `xml:" itRouter,omitempty" json:"itRouter,omitempty" db:" itRouter,omitempty"`
	MitCar             MitCar             `xml:" mitCar,omitempty" json:"mitCar,omitempty" db:" mitCar,omitempty"`
	MitTaxi            MitTaxi            `xml:" mitTaxi,omitempty" json:"mitTaxi,omitempty" db:" mitTaxi,omitempty"`
}

type ItRouter struct {
	LogASCII string `xml:" logASCII,attr"  json:",omitempty"`
	LogSVG   string `xml:" logSVG,attr"  json:",omitempty"`
}

type ItPedestrian struct {
	ComputationType          string `xml:" computationType,attr"  json:",omitempty"`
	ComputeAlternativeRoutes string `xml:" computeAlternativeRoutes,attr"  json:",omitempty"`
	ComputeMonomodalTrip     string `xml:" computeMonomodalTrip,attr"  json:",omitempty"`
	CostFactor               string `xml:" costFactor,attr"  json:",omitempty"`
	DistanceFactor           string `xml:" distanceFactor,attr"  json:",omitempty"`
	IgnoreRestrictions       string `xml:" ignoreRestrictions,attr"  json:",omitempty"`
	ItIncidentData           string `xml:" itIncidentData,attr"  json:",omitempty"`
	Level                    string `xml:" level,attr"  json:",omitempty"`
	MaxLength                string `xml:" maxLength,attr"  json:",omitempty"`
	MaxTime                  string `xml:" maxTime,attr"  json:",omitempty"`
	MinLength                string `xml:" minLength,attr"  json:",omitempty"`
	MinTime                  string `xml:" minTime,attr"  json:",omitempty"`
	NoBridge                 string `xml:" noBridge,attr"  json:",omitempty"`
	NoFerry                  string `xml:" noFerry,attr"  json:",omitempty"`
	NoTunnel                 string `xml:" noTunnel,attr"  json:",omitempty"`
	PrefBench                string `xml:" prefBench,attr"  json:",omitempty"`
	PrefIllum                string `xml:" prefIllum,attr"  json:",omitempty"`
	PrefToilet               string `xml:" prefToilet,attr"  json:",omitempty"`
	SpeedFactor              string `xml:" speedFactor,attr"  json:",omitempty"`
	TraveltimeFactor         string `xml:" traveltimeFactor,attr"  json:",omitempty"`
	UseElevation             string `xml:" useElevation,attr"  json:",omitempty"`
	UseGeoRefs               string `xml:" useGeoRefs,attr"  json:",omitempty"`
	UseHdcAcc                string `xml:" useHdcAcc,attr"  json:",omitempty"`
}

type ItBicycle struct {
	ComputationType          string `xml:" computationType,attr"  json:",omitempty"`
	ComputeAlternativeRoutes string `xml:" computeAlternativeRoutes,attr"  json:",omitempty"`
	ComputeMonomodalTrip     string `xml:" computeMonomodalTrip,attr"  json:",omitempty"`
	CostFactor               string `xml:" costFactor,attr"  json:",omitempty"`
	CycleSpeed               string `xml:" cycleSpeed,attr"  json:",omitempty"`
	DistanceFactor           string `xml:" distanceFactor,attr"  json:",omitempty"`
	ElevFac                  string `xml:" elevFac,attr"  json:",omitempty"`
	IgnoreRestrictions       string `xml:" ignoreRestrictions,attr"  json:",omitempty"`
	InclEnvLeisureRoutes     string `xml:" inclEnvLeisureRoutes,attr"  json:",omitempty"`
	InclEnvTrafficDensity    string `xml:" inclEnvTrafficDensity,attr"  json:",omitempty"`
	InclSurfQualCarry        string `xml:" inclSurfQualCarry,attr"  json:",omitempty"`
	InclSurfQualForrest      string `xml:" inclSurfQualForrest,attr"  json:",omitempty"`
	InclSurfQualGravel       string `xml:" inclSurfQualGravel,attr"  json:",omitempty"`
	InclSurfQualSingleTrial  string `xml:" inclSurfQualSingleTrial,attr"  json:",omitempty"`
	ItIncidentData           string `xml:" itIncidentData,attr"  json:",omitempty"`
	Level                    string `xml:" level,attr"  json:",omitempty"`
	MaxLength                string `xml:" maxLength,attr"  json:",omitempty"`
	MaxTime                  string `xml:" maxTime,attr"  json:",omitempty"`
	MinLength                string `xml:" minLength,attr"  json:",omitempty"`
	MinTime                  string `xml:" minTime,attr"  json:",omitempty"`
	NoBridge                 string `xml:" noBridge,attr"  json:",omitempty"`
	NoFerry                  string `xml:" noFerry,attr"  json:",omitempty"`
	NoTunnel                 string `xml:" noTunnel,attr"  json:",omitempty"`
	PrefHikePath             string `xml:" prefHikePath,attr"  json:",omitempty"`
	PrefPedestrianSeperated  string `xml:" prefPedestrianSeperated,attr"  json:",omitempty"`
	PreferAsphaltTracks      string `xml:" preferAsphaltTracks,attr"  json:",omitempty"`
	PreferGreenTracks        string `xml:" preferGreenTracks,attr"  json:",omitempty"`
	SpeedFactor              string `xml:" speedFactor,attr"  json:",omitempty"`
	TraveltimeFactor         string `xml:" traveltimeFactor,attr"  json:",omitempty"`
	UseDrawClasses           string `xml:" useDrawClasses,attr"  json:",omitempty"`
	UseElevation             string `xml:" useElevation,attr"  json:",omitempty"`
	UseEnv                   string `xml:" useEnv,attr"  json:",omitempty"`
	UseLayoutDirections      string `xml:" useLayoutDirections,attr"  json:",omitempty"`
	UseLayoutStructural      string `xml:" useLayoutStructural,attr"  json:",omitempty"`
	UseNetClasses            string `xml:" useNetClasses,attr"  json:",omitempty"`
	UsePseudoRouting         string `xml:" usePseudoRouting,attr"  json:",omitempty"`
	UseRecommRoute           string `xml:" useRecommRoute,attr"  json:",omitempty"`
	UseSignedRoute           string `xml:" useSignedRoute,attr"  json:",omitempty"`
	UseSurfQuality           string `xml:" useSurfQuality,attr"  json:",omitempty"`
}

type MitCar struct {
	ComputationType          string `xml:" computationType,attr"  json:",omitempty"`
	ComputeAlternativeRoutes string `xml:" computeAlternativeRoutes,attr"  json:",omitempty"`
	ComputeMonomodalTrip     string `xml:" computeMonomodalTrip,attr"  json:",omitempty"`
	CostFactor               string `xml:" costFactor,attr"  json:",omitempty"`
	DistanceFactor           string `xml:" distanceFactor,attr"  json:",omitempty"`
	IgnoreRestrictions       string `xml:" ignoreRestrictions,attr"  json:",omitempty"`
	ItIncidentData           string `xml:" itIncidentData,attr"  json:",omitempty"`
	Level                    string `xml:" level,attr"  json:",omitempty"`
	MaxLength                string `xml:" maxLength,attr"  json:",omitempty"`
	MaxTime                  string `xml:" maxTime,attr"  json:",omitempty"`
	MinLength                string `xml:" minLength,attr"  json:",omitempty"`
	MinTime                  string `xml:" minTime,attr"  json:",omitempty"`
	MitIncidentData          string `xml:" mitIncidentData,attr"  json:",omitempty"`
	MitOnlineData            string `xml:" mitOnlineData,attr"  json:",omitempty"`
	MitProfileData           string `xml:" mitProfileData,attr"  json:",omitempty"`
	NoBridge                 string `xml:" noBridge,attr"  json:",omitempty"`
	NoFerry                  string `xml:" noFerry,attr"  json:",omitempty"`
	NoHighway                string `xml:" noHighway,attr"  json:",omitempty"`
	NoTollRoad               string `xml:" noTollRoad,attr"  json:",omitempty"`
	NoTunnel                 string `xml:" noTunnel,attr"  json:",omitempty"`
	PrefSerSt                string `xml:" prefSerSt,attr"  json:",omitempty"`
	SpeedFactor              string `xml:" speedFactor,attr"  json:",omitempty"`
	TraveltimeFactor         string `xml:" traveltimeFactor,attr"  json:",omitempty"`
	UseElevation             string `xml:" useElevation,attr"  json:",omitempty"`
}

type MitTaxi struct {
	ComputationType          string `xml:" computationType,attr"  json:",omitempty"`
	ComputeAlternativeRoutes string `xml:" computeAlternativeRoutes,attr"  json:",omitempty"`
	ComputeMonomodalTrip     string `xml:" computeMonomodalTrip,attr"  json:",omitempty"`
	CostFactor               string `xml:" costFactor,attr"  json:",omitempty"`
	DistanceFactor           string `xml:" distanceFactor,attr"  json:",omitempty"`
	IgnoreRestrictions       string `xml:" ignoreRestrictions,attr"  json:",omitempty"`
	ItIncidentData           string `xml:" itIncidentData,attr"  json:",omitempty"`
	Level                    string `xml:" level,attr"  json:",omitempty"`
	MaxLength                string `xml:" maxLength,attr"  json:",omitempty"`
	MaxTime                  string `xml:" maxTime,attr"  json:",omitempty"`
	MinLength                string `xml:" minLength,attr"  json:",omitempty"`
	MinTime                  string `xml:" minTime,attr"  json:",omitempty"`
	MitIncidentData          string `xml:" mitIncidentData,attr"  json:",omitempty"`
	MitOnlineData            string `xml:" mitOnlineData,attr"  json:",omitempty"`
	MitProfileData           string `xml:" mitProfileData,attr"  json:",omitempty"`
	NoBridge                 string `xml:" noBridge,attr"  json:",omitempty"`
	NoFerry                  string `xml:" noFerry,attr"  json:",omitempty"`
	NoHighway                string `xml:" noHighway,attr"  json:",omitempty"`
	NoTollRoad               string `xml:" noTollRoad,attr"  json:",omitempty"`
	NoTunnel                 string `xml:" noTunnel,attr"  json:",omitempty"`
	PrefSerSt                string `xml:" prefSerSt,attr"  json:",omitempty"`
	SpeedFactor              string `xml:" speedFactor,attr"  json:",omitempty"`
	TraveltimeFactor         string `xml:" traveltimeFactor,attr"  json:",omitempty"`
	UseElevation             string `xml:" useElevation,attr"  json:",omitempty"`
}

type DepartureTransport struct {
	IndividualTransport []IndividualTransport `xml:" individualTransport,omitempty" json:"individualTransport,omitempty" db:" individualTransport,omitempty"`
}

type IndividualTransport struct {
	MeansCode string `xml:" meansCode,attr"  json:",omitempty"`
	Selected  string `xml:" selected,attr"  json:",omitempty"`
	Speed     string `xml:" speed,attr"  json:",omitempty"`
	Value     string `xml:" value,attr"  json:",omitempty"`
}

type ItdUsedOptions struct {
	CalcCO2           string `xml:" calcCO2,attr"  json:",omitempty"`
	CalcNumberOfTrips string `xml:" calcNumberOfTrips,attr"  json:",omitempty"`
	DwellTime         string `xml:" dwellTime,attr"  json:",omitempty"`
	ItemType          string `xml:" itemType,attr"  json:",omitempty"`
	NextDepsPerLeg    string `xml:" nextDepsPerLeg,attr"  json:",omitempty"`
	RealOnlyInfo      string `xml:" realOnlyInfo,attr"  json:",omitempty"`
	Realtime          string `xml:" realtime,attr"  json:",omitempty"`
}

type ItdServingLines struct {
	ItdServingLine []ItdServingLine `xml:" itdServingLine,omitempty" json:"itdServingLine,omitempty" db:" itdServingLine,omitempty"`
}

type ItdServingLine struct {
	Code             string           `xml:" code,attr"  json:",omitempty"`
	Compound         string           `xml:" compound,attr"  json:",omitempty"`
	DestID           string           `xml:" destID,attr"  json:",omitempty"`
	Direction        string           `xml:" direction,attr"  json:",omitempty"`
	DirectionFrom    string           `xml:" directionFrom,attr"  json:",omitempty"`
	Index            string           `xml:" index,attr"  json:",omitempty"`
	Key              string           `xml:" key,attr"  json:",omitempty"`
	MotType          string           `xml:" motType,attr"  json:",omitempty"`
	MtSubcode        string           `xml:" mtSubcode,attr"  json:",omitempty"`
	Number           string           `xml:" number,attr"  json:",omitempty"`
	OK               string           `xml:" oK,attr"  json:",omitempty"`
	ProductId        string           `xml:" productId,attr"  json:",omitempty"`
	ROP              string           `xml:" ROP,attr"  json:",omitempty"`
	Realtime         string           `xml:" realtime,attr"  json:",omitempty"`
	STT              string           `xml:" STT,attr"  json:",omitempty"`
	Selected         string           `xml:" selected,attr"  json:",omitempty"`
	SpTr             string           `xml:" spTr,attr"  json:",omitempty"`
	Stateless        string           `xml:" stateless,attr"  json:",omitempty"`
	Symbol           string           `xml:" symbol,attr"  json:",omitempty"`
	TTB              string           `xml:" TTB,attr"  json:",omitempty"`
	TrainName        string           `xml:" trainName,attr"  json:",omitempty"`
	Type             string           `xml:" type,attr"  json:",omitempty"`
	VF               string           `xml:" vF,attr"  json:",omitempty"`
	VT               string           `xml:" vT,attr"  json:",omitempty"`
	Valid            string           `xml:" valid,attr"  json:",omitempty"`
	ItdNoTrain       ItdNoTrain       `xml:" itdNoTrain,omitempty" json:"itdNoTrain,omitempty" db:" itdNoTrain,omitempty"`
	ItdOperator      ItdOperator      `xml:" itdOperator,omitempty" json:"itdOperator,omitempty" db:" itdOperator,omitempty"`
	ItdRouteDescText ItdRouteDescText `xml:" itdRouteDescText,omitempty" json:"itdRouteDescText,omitempty" db:" itdRouteDescText,omitempty"`
	MotDivaParams    MotDivaParams    `xml:" motDivaParams,omitempty" json:"motDivaParams,omitempty" db:" motDivaParams,omitempty"`
}

type ItdNoTrain struct {
	Delay string `xml:" delay,attr"  json:",omitempty"`
	Name  string `xml:" name,attr"  json:",omitempty"`
}

type MotDivaParams struct {
	Direction  string `xml:" direction,attr"  json:",omitempty"`
	Line       string `xml:" line,attr"  json:",omitempty"`
	Network    string `xml:" network,attr"  json:",omitempty"`
	Project    string `xml:" project,attr"  json:",omitempty"`
	Supplement string `xml:" supplement,attr"  json:",omitempty"`
}

type ItdRouteDescText struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ItdOperator struct {
	Code Code `xml:" code,omitempty" json:"code,omitempty" db:" code,omitempty"`
	Name Name `xml:" name,omitempty" json:"name,omitempty" db:" name,omitempty"`
}

type Code struct {
	Text int8 `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ItdDepartureList struct {
	ItdDepartures []ItdDeparture `xml:" itdDeparture,omitempty" json:"itdDeparture,omitempty" db:" itdDeparture,omitempty"`
}

type ItdDeparture struct {
	Area           string         `xml:" area,attr"  json:",omitempty"`
	Countdown      string         `xml:" countdown,attr"  json:",omitempty"`
	Gid            string         `xml:" gid,attr"  json:",omitempty"`
	MapName        string         `xml:" mapName,attr"  json:",omitempty"`
	NameWO         string         `xml:" nameWO,attr"  json:",omitempty"`
	Platform       string         `xml:" platform,attr"  json:",omitempty"`
	PlatformName   string         `xml:" platformName,attr"  json:",omitempty"`
	PointGid       string         `xml:" pointGid,attr"  json:",omitempty"`
	StopID         string         `xml:" stopID,attr"  json:",omitempty"`
	StopName       string         `xml:" stopName,attr"  json:",omitempty"`
	X              string         `xml:" x,attr"  json:",omitempty"`
	Y              string         `xml:" y,attr"  json:",omitempty"`
	ItdDateTime    ItdDateTime    `xml:" itdDateTime,omitempty" json:"itdDateTime,omitempty" db:" itdDateTime,omitempty"`
	ItdRTDateTime  ItdRTDateTime  `xml:" itdRTDateTime,omitempty" json:"itdRTDateTime,omitempty" db:" itdRTDateTime,omitempty"`
	ItdServingLine ItdServingLine `xml:" itdServingLine,omitempty" json:"itdServingLine,omitempty" db:" itdServingLine,omitempty"`
	ItdServingTrip ItdServingTrip `xml:" itdServingTrip,omitempty" json:"itdServingTrip,omitempty" db:" itdServingTrip,omitempty"`
}

type ItdRTDateTime struct {
	ItdDate []ItdDate `xml:" itdDate,omitempty" json:"itdDate,omitempty" db:" itdDate,omitempty"`
	ItdTime ItdTime   `xml:" itdTime,omitempty" json:"itdTime,omitempty" db:" itdTime,omitempty"`
}

type ItdServingTrip struct {
	AVMSTripID string `xml:" AVMSTripID,attr"  json:",omitempty"`
	TripCode   string `xml:" tripCode,attr"  json:",omitempty"`
}
