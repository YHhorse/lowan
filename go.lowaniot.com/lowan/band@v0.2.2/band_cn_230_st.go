package band

import (
	"encoding/binary"
	"time"

	"github.com/brocaar/lorawan"
)

type cn230STBand struct {
	band
}

func (b *cn230STBand) GetRX1FrequencyForJoinFrequency(joinFrequency int) (int, error) {
	return b.GetRX1FrequencyForUplinkFrequency(joinFrequency)
}

//TODO
func (b *cn230STBand) CheckJoinFrequency(frequency int) bool {
	return true
}

func (b *cn230STBand) Name() string {
	return "CN230ST"
}

func (b *cn230STBand) GetDefaults() Defaults {
	return Defaults{
		RX2Frequency:     234300000,
		RX2DataRate:      2,
		MaxFCntGap:       16384,
		ReceiveDelay1:    time.Second,
		ReceiveDelay2:    time.Second * 2,
		JoinAcceptDelay1: time.Second * 5,
		JoinAcceptDelay2: time.Second * 6,
	}
}

func (b *cn230STBand) GetDownlinkTXPower(freq int) int {
	return 27
}

func (b *cn230STBand) GetPingSlotFrequency(devAddr lorawan.DevAddr, beaconTime time.Duration) (int, error) {
	downlinkChannel := (int(binary.BigEndian.Uint32(devAddr[:])) + int(beaconTime/(128*time.Second))) % 8
	return b.downlinkChannels[downlinkChannel].Frequency, nil
}

func (b *cn230STBand) GetRX1ChannelIndexForUplinkChannelIndex(uplinkChannel int) (int, error) {
	return uplinkChannel, nil
}

func (b *cn230STBand) GetRX1FrequencyForUplinkFrequency(uplinkFrequency int) (int, error) {
	uplinkChan, err := b.GetUplinkChannelIndex(uplinkFrequency, true)
	if err != nil {
		return 0, err
	}

	rx1Chan, err := b.GetRX1ChannelIndexForUplinkChannelIndex(uplinkChan)
	if err != nil {
		return 0, err
	}

	return b.downlinkChannels[rx1Chan].Frequency, nil
}

func (b *cn230STBand) GetRX2Frequency(frequency int, isABP bool) int {
	return b.GetDefaults().RX2Frequency
}

func (b *cn230STBand) GetRX2DataRate(uplinkDR int, drOffset int) int {
	return b.GetDefaults().RX2DataRate
}

func (b *cn230STBand) ImplementsTXParamSetup(protocolVersion string) bool {
	return false
}

func (b *cn230STBand) GetDefaultMaxUplinkEIRP() float32 {
	return 19.15
}

func newCN230STBand(repeaterCompatible bool) (Band, error) {
	b := cn230STBand{
		band: band{
			dataRates: map[int]DataRate{
				0: {Modulation: LoRaModulation, SpreadFactor: 11, Bandwidth: 75, uplink: true, downlink: true},
				1: {Modulation: LoRaModulation, SpreadFactor: 10, Bandwidth: 75, uplink: true, downlink: true},
				2: {Modulation: LoRaModulation, SpreadFactor: 9, Bandwidth: 75, uplink: true, downlink: true},
				3: {Modulation: LoRaModulation, SpreadFactor: 8, Bandwidth: 75, uplink: true, downlink: true},
				4: {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 75, uplink: true, downlink: true},
				5: {Modulation: LoRaModulation, SpreadFactor: 6, Bandwidth: 75, uplink: true, downlink: true},

				6:  {Modulation: LoRaModulation, SpreadFactor: 10, Bandwidth: 50, uplink: true, downlink: true},
				7:  {Modulation: LoRaModulation, SpreadFactor: 9, Bandwidth: 50, uplink: true, downlink: true},
				8:  {Modulation: LoRaModulation, SpreadFactor: 8, Bandwidth: 50, uplink: true, downlink: true},
				9:  {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 50, uplink: true, downlink: true},
				10: {Modulation: LoRaModulation, SpreadFactor: 6, Bandwidth: 50, uplink: true, downlink: true},

				11: {Modulation: LoRaModulation, SpreadFactor: 9, Bandwidth: 25, uplink: true, downlink: true},
				12: {Modulation: LoRaModulation, SpreadFactor: 8, Bandwidth: 25, uplink: true, downlink: true},
				13: {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 25, uplink: true, downlink: true},
				14: {Modulation: LoRaModulation, SpreadFactor: 6, Bandwidth: 25, uplink: true, downlink: true},
				15: {Modulation: LoRaModulation, SpreadFactor: 5, Bandwidth: 25, uplink: true, downlink: true},
			},
			rx1DataRateTable: map[int][]int{
				0: {0, 0, 0, 0, 0},
				1: {1, 0, 0, 0, 0},
				2: {2, 1, 0, 0, 0},
				3: {3, 2, 1, 0, 0},
				4: {4, 3, 2, 1, 0},
				5: {5, 4, 3, 2, 1},

				6:  {6, 6, 6, 6, 6},
				7:  {7, 6, 6, 6, 6},
				8:  {8, 7, 6, 6, 6},
				9:  {9, 8, 7, 6, 6},
				10: {10, 9, 8, 7, 6},

				11: {11, 11, 11, 11, 11},
				12: {12, 11, 11, 11, 11},
				13: {13, 12, 11, 11, 11},
				14: {14, 13, 12, 11, 11},
				15: {15, 14, 13, 12, 11},
			},
			txPowerOffsets: []int{
				0,
				-5,
				-8,
				-10,
				-11,
				-13,
				-15,
				-17,
				-20,
				-22,
				-25,
				0, 0, 0, 0, 0,
			},
			uplinkChannels:   make([]Channel, 40),
			downlinkChannels: make([]Channel, 40),
		},
	}

	if repeaterCompatible {
		b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
			latest: {
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.1, 1.0.2B, 1.1.0A, 1.1.0B
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
				},
			},
		}
	} else {
		b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
			LoRaWAN_1_0_1: {
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.1
					0: {M: 59, N: 51},
					1: {M: 108, N: 100},
					2: {M: 208, N: 200},
					3: {M: 258, N: 250},
					4: {M: 258, N: 250},
					5: {M: 258, N: 250},

					6:  {M: 59, N: 51},
					7:  {M: 158, N: 150},
					8:  {M: 258, N: 250},
					9:  {M: 258, N: 250},
					10: {M: 258, N: 250},

					11: {M: 59, N: 51},
					12: {M: 108, N: 100},
					13: {M: 258, N: 250},
					14: {M: 258, N: 250},
					15: {M: 258, N: 250},
				},
			},
			latest: {
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.2B, 1.1.0A, 1.1.0B
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
				},
			},
		}
	}

	// initialize uplink channels
	for i := 0; i < 40; i++ {
		b.uplinkChannels[i] = Channel{
			Frequency: 227000000 + (i * 25000),
			MinDR:     0,
			MaxDR:     15,
			enabled:   true,
		}
	}

	// initialize downlink channels
	for i := 0; i < 40; i++ {
		b.downlinkChannels[i] = Channel{
			Frequency: 234000000 + (i * 25000),
			MinDR:     0,
			MaxDR:     15,
			enabled:   true,
		}
	}

	return &b, nil
}
