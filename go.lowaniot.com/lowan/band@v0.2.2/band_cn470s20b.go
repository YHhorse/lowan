package band

import (
	"encoding/binary"
	"time"

	"github.com/brocaar/lorawan"
)

// NOTE: implement of RP2-1.0.3
type cn470s20bBand struct {
	band
}

func (b *cn470s20bBand) CheckJoinFrequency(frequency int) bool {
	for _, channel := range b.joinChannels {
		if frequency == channel.Frequency {
			return true
		}
	}

	return false
}

func (b *cn470s20bBand) Name() string {
	return "CN470S20B"
}

func (b *cn470s20bBand) GetDefaults() Defaults {
	return Defaults{
		RX2Frequency:     498300000,
		RX2DataRate:      1,
		MaxFCntGap:       16384,
		ReceiveDelay1:    time.Second,
		ReceiveDelay2:    time.Second * 2,
		JoinAcceptDelay1: time.Second * 5,
		JoinAcceptDelay2: time.Second * 6,
	}
}

func (b *cn470s20bBand) GetDownlinkTXPower(freq int) int {
	return 14
}

func (b *cn470s20bBand) GetPingSlotFrequency(devAddr lorawan.DevAddr, beaconTime time.Duration) (int, error) {
	downlinkChannel := (int(binary.BigEndian.Uint32(devAddr[:])) + int(beaconTime/(128*time.Second))) % 8
	return b.downlinkChannels[downlinkChannel].Frequency, nil
}

func (b *cn470s20bBand) GetRX1ChannelIndexForUplinkChannelIndex(uplinkChannel int) (int, error) {
	return uplinkChannel % 64, nil
}

func (b *cn470s20bBand) GetRX1FrequencyForUplinkFrequency(uplinkFrequency int) (int, error) {
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

func (b *cn470s20bBand) GetRX1FrequencyForJoinFrequency(joinFrequency int) (int, error) {
	return b.GetRX1FrequencyForUplinkFrequency(joinFrequency)
}

func (b *cn470s20bBand) GetRX2Frequency(frequency int, isABP bool) int {
	if isABP {
		return b.GetDefaults().RX2Frequency
	} else {
		upIndex, _ := b.GetUplinkChannelIndex(frequency, true)

		rx2Index := (upIndex/32)*32 + 7
		return b.downlinkChannels[rx2Index].Frequency
	}
}

func (b *cn470s20bBand) GetRX2DataRate(uplinkDR int, drOffset int) int {
	return b.GetDefaults().RX2DataRate
}

func (b *cn470s20bBand) ImplementsTXParamSetup(protocolVersion string) bool {
	return false
}

func (b *cn470s20bBand) GetDefaultMaxUplinkEIRP() float32 {
	return 19.15
}

func newCN470S20BBand(repeaterCompatible bool) (Band, error) {
	b := cn470s20bBand{
		band: band{
			dataRates: map[int]DataRate{
				0: {Modulation: LoRaModulation, SpreadFactor: 12, Bandwidth: 125, uplink: true, downlink: true},
				1: {Modulation: LoRaModulation, SpreadFactor: 11, Bandwidth: 125, uplink: true, downlink: true},
				2: {Modulation: LoRaModulation, SpreadFactor: 10, Bandwidth: 125, uplink: true, downlink: true},
				3: {Modulation: LoRaModulation, SpreadFactor: 9, Bandwidth: 125, uplink: true, downlink: true},
				4: {Modulation: LoRaModulation, SpreadFactor: 8, Bandwidth: 125, uplink: true, downlink: true},
				5: {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 125, uplink: true, downlink: true},
				6: {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 500, uplink: true, downlink: true},
				7: {Modulation: FSKModulation, BitRate: 50000, uplink: true, downlink: true},
			},
			rx1DataRateTable: map[int][]int{
				0: {0, 0, 0, 0, 0, 0},
				1: {1, 1, 1, 1, 1, 1},
				2: {2, 1, 1, 1, 1, 1},
				3: {3, 2, 1, 1, 1, 1},
				4: {4, 3, 2, 1, 1, 1},
				5: {5, 4, 3, 2, 1, 1},
				6: {6, 5, 4, 3, 2, 1},
				7: {7, 6, 5, 4, 3, 2},
			},
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
			uplinkChannels:   make([]Channel, 64),
			downlinkChannels: make([]Channel, 64),
			joinChannels:     make([]Channel, 2),
		},
	}

	if repeaterCompatible {
		b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
			latest: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{
					0: {M: 31, N: 23},
					1: {M: 31, N: 23},
					2: {M: 94, N: 86},
					3: {M: 192, N: 184},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
					6: {M: 230, N: 222},
					7: {M: 230, N: 222},
				},
			},
		}
	} else {
		b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
			latest: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{
					0: {M: 31, N: 23},
					1: {M: 31, N: 23},
					2: {M: 94, N: 86},
					3: {M: 192, N: 184},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
					6: {M: 250, N: 242},
					7: {M: 250, N: 242},
				},
			},
		}
	}

	// initialize uplink channels
	for i := 0; i < 32; i++ {
		b.uplinkChannels[i] = Channel{
			Frequency: 476900000 + (i * 200000),
			MinDR:     0,
			MaxDR:     5,
			enabled:   true,
		}
	}

	for i := 32; i < 64; i++ {
		b.uplinkChannels[i] = Channel{
			Frequency: 496900000 + ((i - 32) * 200000),
			MinDR:     0,
			MaxDR:     5,
			enabled:   true,
		}
	}

	// initialize downlink channels
	for i := 0; i < 32; i++ {
		b.downlinkChannels[i] = Channel{
			Frequency: 476900000 + (i * 200000),
			MinDR:     0,
			MaxDR:     5,
			enabled:   true,
		}
	}
	for i := 32; i < 64; i++ {
		b.downlinkChannels[i] = Channel{
			Frequency: 496900000 + ((i - 32) * 200000),
			MinDR:     0,
			MaxDR:     5,
			enabled:   true,
		}
	}

	// initial join channels
	b.joinChannels[0] = Channel{
		Frequency: 479900000,
		MinDR:     0,
		MaxDR:     5,
		enabled:   true,
	}

	b.joinChannels[1] = Channel{
		Frequency: 499900000,
		MinDR:     0,
		MaxDR:     5,
		enabled:   true,
	}

	return &b, nil
}
