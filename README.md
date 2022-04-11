# inheritance (Miras) ve Struct Yapıları

 ## Struct 
  - Birden fazla veriyi tutabilmektedirler
  - Yazım tipleri <code>type StructName struct{}</code> şeklindedir
  - İçlerinde map, array, struct gibi farklı veri türleri bulundurabilirler.

### Struct Tanımlama ve Kullanım

 Örnek olarak <code>main func()</code> dışında <code>type Devices struct{}</code> şeklinde tanımlandıktan sonra <code>main func()</code> içerisinde:

<code>device:=Device{}</code> ,

<code>var device Device{}</code> (Burada default değerler verilmiş olur)

Şeklinde tanımlanarak kullanılabilir.

***
## İnheritance

- Miras anlamına gelmektedir.
- İç içe structların kullanımı sonrasında oluşmaktadır. 


                    type person struct {
	                name string
	                age  int
                    }

                    type student struct {
	                person
	                okul string
                    }


 Şeklinde Oluşturulan yapılardır.

person structındaki veriler miras olark student structında kullanılabilir. 

Bunun kullanımı;


                    func main() {

                    ogr1 := student{person{name: "taha", age: 12}, "İlkokul"}

	                fmt.Println(ogr1)
					}
    

şeklindedir.