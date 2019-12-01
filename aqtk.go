package aqtk

// #include <AquesTalk.h>
// #include <stdlib.h>
// #cgo LDFLAGS: -l AquesTalk10 -lstdc++
import "C"
import (
	"fmt"
	"unsafe"
)

// TmplVoice is a helper type
// for defining default voice
// values.
type TmplVoice *C.AQTK_VOICE

var (
	F1 TmplVoice = &C.gVoice_F1
	F2           = &C.gVoice_F2
	F3           = &C.gVoice_F3
	M1           = &C.gVoice_M1
	M2           = &C.gVoice_M2
	R1           = &C.gVoice_R1
	R2           = &C.gVoice_R2
)

// Voice is a container for AQTK_VOICE.
type Voice C.AQTK_VOICE

// NewVoice creates a Voice from a given
// TmplVoice.
func NewVoice(tmpl TmplVoice) *Voice {
	v := Voice(*tmpl)
	return &v
}

// SetSpeed sets the speed of the voice.
// Valid Range: 50-300, Default: 100
func (v *Voice) SetSpeed(speed int) {
	v.spd = C.int(speed)
}

// Speed returns the speed the voice.
func (v *Voice) Speed() int {
	return int(v.spd)
}

// SetVol sets the volume of the voice.
// Valid Range: 0-300
func (v *Voice) SetVol(vol int) {
	v.vol = C.int(vol)
}

// Vol returns the volume of the voice.
func (v *Voice) Vol() int {
	return int(v.vol)
}

// SetHeight sets the pitch height of the voice.
// Valid Range: 20-200.
func (v *Voice) SetHeight(pit int) {
	v.pit = C.int(pit)
}

// SetHeight returns the pitch height
// of the voice.
func (v *Voice) Height() int {
	return int(v.pit)
}

// SetAccent sets the accent of the voice.
// Valid Range: 0-200
func (v *Voice) SetAccent(acc int) {
	v.acc = C.int(acc)
}

// Accent returns the accent of the voice.
func (v *Voice) Accent() int {
	return int(v.acc)
}

// SetPitch1 sets the first property for pitch.
// This coresponds the the lmd member of
// AQTK_VOICE.
// Valid Range: 0-200, Default: 100
func (v *Voice) SetPitch1(lmd int) {
	v.lmd = C.int(lmd)
}

// Pitch1 returns the first property for pitch.
func (v *Voice) Pitch1() int {
	return int(v.lmd)
}

// SetPitch2 sets the second property for pitch.
// This coresponds to the fsc member of
// AQTK_VOICE.
// Valid Range: 50-200, Default: 100
func (v *Voice) SetPitch2(fsc int) {
	v.fsc = C.int(fsc)
}

// Pitch1 returns the second property for pitch.
func (v *Voice) Pitch2() int {
	return int(v.fsc)
}

// Talk returns bytes representing a wav file
// speaking the give string. Only phonetic
// kana is accepted.
func (v *Voice) Talk(s string) ([]byte, error) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	var olen C.int
	wav := C.AquesTalk_Synthe_Utf8((*C.AQTK_VOICE)(v), cs, &olen)
	if wav == nil {
		return nil, fmt.Errorf("error code %d", olen)
	}

	defer C.AquesTalk_FreeWave(wav)
	out := C.GoBytes(unsafe.Pointer(wav), olen)
	return out, nil
}