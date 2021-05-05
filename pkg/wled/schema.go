package wled

import (
	"net"
	"net/http"
	"sync"
)

type Protocol byte

const (
	ProtocolDNRGB Protocol = 4

	WARLSPORT = 21324
)

// Led stores information about a single LED.
// Storing X and Y in order to potentially expand to 2d graphics.
type Led struct {
	xLocation int
	yLocation int

	state LedState

	lock sync.Mutex
}

type LedState struct {
	Red   byte
	Green byte
	Blue  byte
}

type Server struct {
	conns      []*net.UDPConn
	ips        []string
	httpClient *http.Client
	lock       sync.Mutex

	leds []Led
}

type APIResponse struct {
	State    *APIState `json:"state,omitempty"`
	Info     *APIInfo  `json:"info,omitempty"`
	Effects  []string  `json:"effects,omitempty"`
	Palettes []string  `json:"palettes,omitempty"`
	Success  *bool     `json:"success,omitempty"`
}

type APIInfo struct {
	Ver  string `json:"ver,omitempty"`
	Vid  int    `json:"vid,omitempty"`
	Leds struct {
		Count  int   `json:"count,omitempty"`
		Rgbw   bool  `json:"rgbw,omitempty"`
		Pin    []int `json:"pin,omitempty"`
		Pwr    int   `json:"pwr,omitempty"`
		Maxpwr int   `json:"maxpwr,omitempty"`
		Maxseg int   `json:"maxseg,omitempty"`
	} `json:"leds,omitempty"`
	Name     string `json:"name,omitempty"`
	Udpport  int    `json:"udpport,omitempty"`
	Live     bool   `json:"live,omitempty"`
	Fxcount  int    `json:"fxcount,omitempty"`
	Palcount int    `json:"palcount,omitempty"`
	Arch     string `json:"arch,omitempty"`
	Core     string `json:"core,omitempty"`
	Freeheap int    `json:"freeheap,omitempty"`
	Uptime   int    `json:"uptime,omitempty"`
	Opt      int    `json:"opt,omitempty"`
	Brand    string `json:"brand,omitempty"`
	Product  string `json:"product,omitempty"`
	Btype    string `json:"btype,omitempty"`
	Mac      string `json:"mac,omitempty"`
}

type APIState struct {
	On         *bool               `json:"on,omitempty"`
	Bri        int                 `json:"bri,omitempty"`
	Transition int                 `json:"transition,omitempty"`
	Ps         int                 `json:"ps,omitempty"`
	Pl         int                 `json:"pl,omitempty"`
	Nl         *APINightLightState `json:"nl,omitempty"`
	Udpn       *UDPState           `json:"udpn,omitempty"`
	Segments   []APISegment        `json:"seg,omitempty"`
}

type APISegment struct {
	Start           int     `json:"start,omitempty"`
	Stop            int     `json:"stop,omitempty"`
	Len             int     `json:"len,omitempty"`
	Col             [][]int `json:"col,omitempty"`
	EffectId        int     `json:"fx,omitempty"`
	EffectSpeed     int     `json:"sx,omitempty"`
	EffectIntensity int     `json:"ix,omitempty"`
	ColourPaletteID int     `json:"pal,omitempty"`
	Sel             bool    `json:"sel,omitempty"`
	Rev             bool    `json:"rev,omitempty"`
	Cln             int     `json:"cln,omitempty"`
}

type APINightLightState struct {
	On   bool `json:"on,omitempty"`
	Dur  int  `json:"dur,omitempty"`
	Fade bool `json:"fade,omitempty"`
	Tbri int  `json:"tbri,omitempty"`
}

type UDPState struct {
	Send bool `json:"send,omitempty"`
	Recv bool `json:"recv,omitempty"`
}
