package context182

type in struct {
	station string
	t       int
}

type ava struct {
	t int
	n int
}
type UndergroundSystem struct {
	checkIn map[int]in
	ava     map[string]map[string]ava
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		checkIn: make(map[int]in),
		ava:     make(map[string]map[string]ava),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.checkIn[id] = in{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	in := this.checkIn[id]

	var x ava
	if _, ok := this.ava[in.station]; !ok {
		this.ava[in.station] = make(map[string]ava)
	} else {
		x = this.ava[in.station][stationName]
	}
	x.t += t - in.t
	x.n++
	this.ava[in.station][stationName] = x
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	x := this.ava[startStation][endStation]
	return float64(x.t) / float64(x.n)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
