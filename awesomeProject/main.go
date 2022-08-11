package main

// function dışında sayi:= 4 tanımlamaya izin vermez var lı tanımlama gerekir.
func main() {
	/* 1- Değişkenler
	fmt.Println("Hello")

	var message string
	message = "Hello Go ! "
	fmt.Println(message)
	var message string ="Hello Go ! "
	var message = "Merhaba Go ! " // Bu şekilde tanımlarsak reflection arka tarafta tipini algılayıp ona göre tanımlıyor.
	fmt.Println(message)
	// Tek bir satırda birden fazla satır oluşturma ilk değer 0 olarak atanır
	var a, b, c int

	var a, b, c int = 3, 4, 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//var kullanmadan değişken tanımı için := gerekiyor. = le yaparsak hata oluşturur çalışmaz.
	u := 55
	fmt.Println(u)

	a := "metin"
	b := 'M' // char verisi alır tek karakter sadece
	fmt.Println(a, b) // output : metin 77


	var myFloat32 float32 = 44.
	myComplex := complex(3, 4)
	println(myComplex) //(+3.000000e+000+4.000000e+000i)
	println(myFloat32) //+4.400000e+001
	*/

	/* 2 -tür değişimleri

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, z)

	// String den int e dönüştürme
	// hata durumunda hata nesnesi dönüyor onu kullanmak istemiyorsak _ şeklinde yapıyoruz
	number, _ := strconv.Atoi("10")
	fmt.Println(number)
	fmt.Println("Sonuc : " + strconv.Itoa(number))

	*/

	/* 3- Fonksiyon nedir ve neden kullanılır
	fmt.Println(add(2, 3))
	*/

	/* 4 -Struct Yapılar : Alanlardan oluşan bir veri topluluğudur. type bildirimi struct oluştumaya yarar.
	// Yapı alanlarına nokta ile ulaşılır.

	fmt.Println(Vertex{1, 2})

	v := Vertex{3, 4}
	v.X = 15
	fmt.Println(v.X)

	//Nesne Örneği Oluşturma v1
	x := Human{FirstName: "Oğuzhan"}
	fmt.Println(x.FirstName)

	//v2
	x := new(Human)
	x.FirstName = "Oğuzhan"
	fmt.Println(x.FirstName)

	// Constructor kullanımı
	x := NewHuman()
	x.Age = 28
	fmt.Println(x.Age)

	// Parametreli Constructor Kullanımı
	x := NewHumanWithParams("Oğuzhan", "Bektaş", 28)
	fmt.Println(x)

	// Konsol işlemleri
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text : ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter a number : ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number : ", f)
	}

	// For Döngüsü
		for i := 0; i < 5; i++ {
		fmt.Println("Value : ", i)
	}
	for i := 0; ; i++ { // Sonsuz döngü
		fmt.Println("Value : ", i)
	}
	for i := 0; i < 5; { // Arttırmıyor ve sonsuz döngüye giriyor
		fmt.Println("Value : ", i)
	}
	for i := 0; i < 5; {
		fmt.Println("Value : ", i)
		i++
	}
	// Diziler
	// Basit bir dizi
	myArray1 := [3]int{}
	myArray1[0] = 32
	myArray1[1] = 23
	myArray1[2] = 54
	fmt.Println(myArray1)

	var numbers = [5]int{4, 6, 7, 3, 66}
	fmt.Println(numbers)
	fmt.Println("Number of numbers : ", len(numbers))

	myArray1 := [...]int{4, 3, 5, 6, 7, 5} // 3 nokta ile size kadar alıyor
	fmt.Println(myArray1)

	// Renk Dizisi // Foreach gibi...
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"
	fmt.Println(colors)
	for i, value := range colors {
		fmt.Println("i = ", i, " & value = ", value)
	}
	for i, _ := range colors {
		fmt.Println("i = ", i, " & value = ", colors[i])
	}
	for _, value := range colors {
		fmt.Println("Value = ", value)
	}
	//Slice
	var colors = []string{"Red", "Green", "Blue"}
		fmt.Println(colors)
		colors = append(colors, "Purple")
		fmt.Println(colors)
		colors = append(colors[1:len(colors)])// ilk elemanı silecek
		fmt.Println(colors)

	// Capasite he zaman elaman sayısına eşit veya uzun olmak zorundadır.
	numbers := make([]int, 5, 5)
	numbers[0] = 123
	numbers[1] = 43
	numbers[2] = 87
	numbers[3] = 456
	numbers[4] = 654
	fmt.Println(numbers)
	numbers = append(numbers, 342)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))

	// Map nesnesini oluşturma yöntemleri
	// 1. Key Value Pair
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[43] = "Foo"
	myMap[12] = "Bar"
	fmt.Println(myMap)

	// 2.
	states := make(map[string]string)
	states["IST"] = "İstanbul"
	states["ANK"] = "Ankara"
	states["ANT"] = "Antalya"
	fmt.Println(states)

	// şehir listesinde ANT anahtar adına sahip veriyi elde et
	antalya := states["ANT"]
	fmt.Println("Seçili Şehir : ", antalya)

	// ANK yi sil
	delete(states, "ANT")
	fmt.Println(states)

	states["ERZ"] = "Erzurum"

	for k, v := range states {
		fmt.Printf("%v : %v \n", k, v)
	}

	// states map nesnesinin elemano adedince kapasiteli bir key listesi oluştur

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	// Concurrency - Eş zamanlılık
	// Go routines
	// Kanallar
	// Ortam Değişkenlerini Okumak ve Değiştirmek
	// Hata yönetimine genel bakış
	// time paketinin temel kullanımı
	// Tarih ve Zaman Operasyonları
	// String Birleştirme
	// bytes paketi ile
	// strings paketi ile
	// sifre üretme generate password- Basic Version
	// sifre üretme generate password - Gelişmiş Version
	// Dosya
	// Dosya Sıkıştırma Operasyonları : ZIP & TAR
	// Veri Dosyaları ile Çalışmak : XML , JSON ,CSV , YAML , TOML
	// RESTful API Mimari ve Tasarımı
	// RESTful API Programlama Temelleri
	// Go RESTful API Projesi : Kullanıcı Bilgi Sistemi
	// Go RESTful API Projesi : Kullanıcı Kayıt ve Login
	// Mini Go RESTful API & Web Uygulamaları
	// Go RESTful API Projesi : Ürün Yönetimi
	// Golang ile Veritabanı Programlama
	*/

}

func add(x int, y int) int {
	return x + y
}

type Vertex struct {
	X int
	Y int
}
type Human struct {
	FirstName string
	LastName  string
	Age       int
}

//constructor işlemi bu şekilde yapılır. Go desteklemiyor construcyor ı bu şekilde yapmamız gerekiyor.
func NewHuman() *Human {
	h := new(Human)
	return h
}

// multi constructor metot override yok :-(
func NewHumanWithParams(firstName, lastName string, age int) *Human {
	h := new(Human)
	h.FirstName = firstName
	h.LastName = lastName
	h.Age = age
	return h
}
