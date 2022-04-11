package main

import "fmt"

/*
type Device struct {
	deviceName   string
	deciveType   int
	deviceStatus bool
}

func main() {

	//                   Struct Yapısı

	device := Device{deviceName: "Buzdolabı", deciveType: 2, deviceStatus: true}
	fmt.Println(device)
	deviceList := []Device{{"Rana", 12, true}, {"Taha", 89, true}}

	fmt.Println(deviceList)

}
*/

/*
type Device struct {
	deviceName   string
	deciveType   int
	deviceStatus bool
}
/*buraya device adresi gönderiliyor adresten status deiştiriliyor
func (a *Device) statusChange() bool {
	a.deviceStatus = true
	return true
}
func main() {
	k := Device{deviceName: "Buzdolabı", deciveType: 2, deviceStatus: true}
	c := &k
	fmt.Println(c.statusChange())//k nın adresini statuschange func a gönderiyor.
	fmt.Println(k)

}*/

//             ÖRNEK METHOD
/*
type ucgen struct {
	a, b, c int
	h       int
}

//method tanımlarken böyle
func (u ucgen) ucgenCevre() int {
	cevre := u.a + u.b + u.c
	return cevre
}
func (u ucgen) ucgenHipotenus() int {

	hipo := (u.a * u.a) + (u.b * u.b)
	return int(math.Sqrt(float64(hipo)))

}

func (u ucgen) ucgenAlan() int {
	return (u.a * u.h / 2)
}
func main() {
	b := ucgen{3, 4, 5, 10}
	fmt.Println(b.ucgenCevre())
	fmt.Println(b.ucgenAlan())
	fmt.Println(b.ucgenHipotenus())

}
*/
/*
type Device struct {
	decviceName  string
	deviceType   int
	deviceStatus bool
}

func main() {
	//        array yapısı
	/*
		d := [10]Device{}
		d[0]=Device{decviceName: "CCM"}
		fmt.Println(d)
*/

//map yapısı
/*
	m := make(map[int]Device)
	fmt.Println(m)
	m[0] = Device{decviceName: "Tech"}
	fmt.Println(m)

}
*/

/*
type person struct {
	name string
	age  int
}

type student struct {
	person
	okul string
}

func main() {

	//inherthance Miras olayı
	ogr1 := student{person{name: "taha", age: 12}, "İlkokul"}

	fmt.Println(ogr1)
}
*/

//   ornek

type sirket struct {
	sirketName    string
	sirketType    int
	sirketCalisan []calisan
}
type calisan struct {
	calisanName  string
	calisanBirim int
}

func main() {
	calisan1 := calisan{"ali", 1}
	calisan2 := calisan{"veli", 2}

	calisanlar := []calisan{calisan1, calisan2}
	sirket := sirket{"CMM", 3, calisanlar}
	fmt.Println(sirket)

}
