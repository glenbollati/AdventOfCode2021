package main

import (
	aoc "aoc/utils"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Packet struct {
	version int64
	typeID  int64
	literal int64
	subPacketData
	subPackets []Packet
}

type subPacketData struct {
	lenType  int64
	totalLen int64
	count    int64
}

var (
	outerPacket []byte
	packets     []*Packet
)

func load() {
	//defer aoc.TimeTrack(time.Now(), "Loading")
	var hexData string
	if len(os.Args) > 1 {
		hexData = aoc.GetLines(os.Args[1])[0]
	} else {
		hexData = aoc.ReadStdin()[0]
	}
	var str string
	for i := 0; i < len(hexData)-1; i += 2 {
		data, err := strconv.ParseInt(hexData[i:i+2], 16, 64)
		if err != nil {
			panic(err)
		}
		str += fmt.Sprintf("%0.8b", data)
	}
	outerPacket = []byte(str)
}

func parseLiteral(msg []byte) (int64, []byte) {
	var chunk, literal []byte
	for {
		chunk, msg = msg[:5], msg[5:]
		literal = append(literal, chunk[1:]...)
		if chunk[0] == '0' {
			break
		}
	}
	return toInt(literal), msg
}

func toInt(data []byte) int64 {
	return aoc.BinStrToI64(string(data))
}

func parsePacket(msg []byte) (pck Packet, remaining []byte) {
	pck.version, pck.typeID = toInt(msg[:3]), toInt(msg[3:6])
	msg = msg[6:]
	// Literal
	if pck.typeID == 4 {
		pck.literal, remaining = parseLiteral(msg)
		return
	}
	// Operator
	pck.lenType, msg = toInt(msg[0:1]), msg[1:]

	if pck.lenType == 0 {
		pck.totalLen, msg = toInt(msg[:15]), msg[15:]
	} else if pck.lenType == 1 {
		pck.count, msg = toInt(msg[:11]), msg[11:]
	}
	// Parse sub packets
	subPck, consumed := Packet{}, 0
	if pck.lenType == 0 {
		for {
			startLen := len(msg)
			subPck, msg = parsePacket(msg)
			pck.subPackets = append(pck.subPackets, subPck)
			consumed += startLen - len(msg)
			if consumed == int(pck.totalLen) || len(msg) == 0 {
				return pck, msg
			}
		}
	}
	for i := 0; i < int(pck.count); i++ {
		if len(msg) == 0 {
			return pck, msg
		}
		subPck, msg = parsePacket(msg)
		pck.subPackets = append(pck.subPackets, subPck)
	}
	return pck, msg
}

func sumVersions(pck Packet) (sum int64) {
	sum += pck.version
	for _, p := range pck.subPackets {
		sum += sumVersions(p)
	}
	return
}

func evalExpr(pck Packet) (result int64) {
	pcks := pck.subPackets
	switch pck.typeID {
	case 0:
		if len(pcks) == 1 {
			return evalExpr(pcks[0])
		}
		for _, p := range pcks {
			result += evalExpr(p)
		}
		return
	case 1:
		if len(pcks) == 1 {
			return evalExpr(pcks[0])
		}
		result = 1
		for _, p := range pcks {
			result *= evalExpr(p)
		}
		return
	case 2:
		result = math.MaxInt64
		for _, p := range pcks {
			res := evalExpr(p)
			if res < result {
				result = res
			}
		}
		return
	case 3:
		for _, p := range pcks {
			res := evalExpr(p)
			if res > result {
				result = res
			}
		}
		return
	case 4:
		return pck.literal
	case 5:
		if evalExpr(pcks[0]) > evalExpr(pcks[1]) {
			return 1
		}
		return 0
	case 6:
		if evalExpr(pcks[0]) < evalExpr(pcks[1]) {
			return 1
		}
		return 0
	case 7:
		if evalExpr(pcks[0]) == evalExpr(pcks[1]) {
			return 1
		}
		return 0
	}
	panic("unreachable")
}

func p1() {
	//defer aoc.TimeTrack(time.Now(), "Part one")
	rootPacket, _ := parsePacket(outerPacket)
	fmt.Println(sumVersions(rootPacket))
}

func p2() {
	//defer aoc.TimeTrack(time.Now(), "Part two")
	rootPacket, _ := parsePacket(outerPacket)
	fmt.Println(evalExpr(rootPacket))
}

func main() {
	load()
	p1() // 977
	p2() // 101501020883
}
