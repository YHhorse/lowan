package band

import (
	"fmt"
	"time"

	"github.com/brocaar/lorawan"
)

const (
	as923FreqOffset2 = -18000
	as923FreqOffset3 = -66000
	as923OffsetStep  = 100
)

type as923Band struct {
	band
	dwellTime    lorawan.DwellTime
	chPlan       int // 1, 2, 3. default 1
	freqOffset   int
	freqOffsetHz int
}

func (b *as923Band) GetRX1FrequencyForJoinFrequency(joinFrequency int) (int, error) {
	return b.GetRX1FrequencyForUplinkFrequency(joinFrequency)
}

//TODO
func (b *as923Band) CheckJoinFrequency(frequency int) bool {
	return true
}

func (b *as923Band) Name() string {
	switch b.chPlan {
	case 2:
		return string(AS923G2)
	case 3:
		return string(AS923G3)
	default:
		return string(AS923)
	}
}

func (b *as923Band) GetDefaults() Defaults {
	return Defaults{
		RX2Frequency:     923200000 + b.freqOffsetHz,
		RX2DataRate:      2,
		MaxFCntGap:       16384,
		ReceiveDelay1:    time.Second,
		ReceiveDelay2:    time.Second * 2,
		JoinAcceptDelay1: time.Second * 5,
		JoinAcceptDelay2: time.Second * 6,
	}
}

func (b *as923Band) GetDownlinkTXPower(freq int) int {
	return 14
}

func (b *as923Band) GetPingSlotFrequency(lorawan.DevAddr, time.Duration) (int, error) {
	return 923400000 + b.freqOffsetHz, nil
}

func (b *as923Band) GetRX1ChannelIndexForUplinkChannelIndex(uplinkChannel int) (int, error) {
	return uplinkChannel, nil
}

func (b *as923Band) GetRX1FrequencyForUplinkFrequency(uplinkFrequency int) (int, error) {
	return uplinkFrequency, nil
}

func (b *as923Band) GetRX1DataRateIndex(uplinkDR, rx1DROffset int) (int, error) {
	if rx1DROffset < 0 || rx1DROffset > 7 {
		return 0, fmt.Errorf("lorawan/band: invalid RX1 data-rate offset: %d", rx1DROffset)
	}

	if uplinkDR < 0 || uplinkDR > 7 {
		return 0, fmt.Errorf("lorawan/band: invalid uplink data-rate: %d", uplinkDR)
	}

	minDR := 0
	if b.dwellTime == lorawan.DwellTime400ms {
		minDR = 2
	}

	effectiveRX1DROffset := []int{0, 1, 2, 3, 4, 5, -1, -2}[rx1DROffset]
	dr := uplinkDR - effectiveRX1DROffset

	if dr < minDR {
		dr = minDR
	}

	if dr > 5 {
		dr = 5
	}

	return dr, nil
}

func (b *as923Band) GetRX2Frequency(frequency int, isABP bool) int {
	return b.GetDefaults().RX2Frequency
}

func (b *as923Band) GetRX2DataRate(uplinkDR int, drOffset int) int {
	return b.GetDefaults().RX2DataRate
}

func (b *as923Band) ImplementsTXParamSetup(protocolVersion string) bool {
	return true
}

func (b *as923Band) GetDefaultMaxUplinkEIRP() float32 {
	return 16
}

func newAS923Band(repeaterCompatible bool, dt lorawan.DwellTime) (Band, error) {
	return newAs923Band(repeaterCompatible, dt, 1)
}

func newAS923Band2(repeaterCompatible bool, dt lorawan.DwellTime) (Band, error) {
	return newAs923Band(repeaterCompatible, dt, 2)
}

func newAS923Band3(repeaterCompatible bool, dt lorawan.DwellTime) (Band, error) {
	return newAs923Band(repeaterCompatible, dt, 3)
}

func calcAs923FreqOffset(plan int) (chPlan, offset, offsetHz int) {
	switch plan {
	case 2:
		return plan, as923FreqOffset2, as923FreqOffset2 * as923OffsetStep
	case 3:
		return plan, as923FreqOffset3, as923FreqOffset3 * as923OffsetStep
	default:
		return 1, 0, 0
	}
}

func newAs923Band(repeaterCompatible bool, dt lorawan.DwellTime, plan int) (Band, error) {
	chPlan, freqOffset, freqOffsetHz := calcAs923FreqOffset(plan)
	b := as923Band{
		dwellTime:    dt,
		chPlan:       chPlan,
		freqOffset:   freqOffset,
		freqOffsetHz: freqOffsetHz,
		band: band{
			supportsExtraChannels: true,
			dataRates: map[int]DataRate{
				0: {Modulation: LoRaModulation, SpreadFactor: 12, Bandwidth: 125, uplink: true, downlink: true},
				1: {Modulation: LoRaModulation, SpreadFactor: 11, Bandwidth: 125, uplink: true, downlink: true},
				2: {Modulation: LoRaModulation, SpreadFactor: 10, Bandwidth: 125, uplink: true, downlink: true},
				3: {Modulation: LoRaModulation, SpreadFactor: 9, Bandwidth: 125, uplink: true, downlink: true},
				4: {Modulation: LoRaModulation, SpreadFactor: 8, Bandwidth: 125, uplink: true, downlink: true},
				5: {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 125, uplink: true, downlink: true},
				6: {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 250, uplink: true, downlink: true},
				7: {Modulation: FSKModulation, BitRate: 50000, uplink: true, downlink: true},
			},
			rx1DataRateTable: map[int][]int{}, // implemented as function
			txPowerOffsets: []int{
				0,
				-2,
				-4,
				-6,
				-8,
				-10,
				-12,
				-14,
			},
			uplinkChannels: []Channel{
				{Frequency: 923200000 + freqOffsetHz, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 923400000 + freqOffsetHz, MinDR: 0, MaxDR: 5, enabled: true},
			},
			downlinkChannels: []Channel{
				{Frequency: 923200000 + freqOffsetHz, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 923400000 + freqOffsetHz, MinDR: 0, MaxDR: 5, enabled: true},
			},
		},
	}

	if dt == lorawan.DwellTime400ms {
		if repeaterCompatible {
			// repeater compatible + dwell time
			b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
				latest: map[string]map[int]MaxPayloadSize{ // LoRaWAN 1.0.2B, 1.1.0A, 1.1.0B
					latest: map[int]MaxPayloadSize{
						0: {M: 0, N: 0},
						1: {M: 0, N: 0},
						2: {M: 19, N: 11},
						3: {M: 61, N: 53},
						4: {M: 133, N: 125},
						5: {M: 250, N: 242},
						6: {M: 250, N: 242},
						7: {M: 250, N: 242},
					},
				},
			}
		} else {
			// not repeater compatible + dwell time
			b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
				latest: map[string]map[int]MaxPayloadSize{
					latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.2B, 1.1.0A, 1.1.0B
						0: {M: 0, N: 0},
						1: {M: 0, N: 0},
						2: {M: 19, N: 11},
						3: {M: 61, N: 53},
						4: {M: 133, N: 125},
						5: {M: 250, N: 242},
						6: {M: 250, N: 242},
						7: {M: 250, N: 242},
					},
				},
			}
		}
	} else {
		if repeaterCompatible {
			// repeater compatible + no dwell time
			b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
				latest: map[string]map[int]MaxPayloadSize{
					latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.2B, 1.1.0A, 1.1.0B
						0: {M: 59, N: 51},
						1: {M: 59, N: 51},
						2: {M: 59, N: 51},
						3: {M: 123, N: 115},
						4: {M: 230, N: 222},
						5: {M: 230, N: 222},
						6: {M: 230, N: 222},
						7: {M: 230, N: 222},
					},
				},
			}
		} else {
			// not repeater compatible + no dwell time
			b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
				latest: map[string]map[int]MaxPayloadSize{
					latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.2B, 1.1.0A, 1.1.0B
						0: {M: 59, N: 51},
						1: {M: 59, N: 51},
						2: {M: 59, N: 51},
						3: {M: 123, N: 115},
						4: {M: 250, N: 242},
						5: {M: 250, N: 242},
						6: {M: 250, N: 242},
						7: {M: 250, N: 242},
					},
				},
			}
		}
	}

	return &b, nil
}
