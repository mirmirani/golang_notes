package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

func aa_imperatives() {
	var module = "#1 imperatives"
	var line_break = "\n------------------------------------\n"
	fmt.Println(line_break, module, line_break)

	fmt.Println("check recursive", 7*6*5*4*3*2*1)
	fmt.Println("Hello I am Go")
	fmt.Println("Let's try multiple arguments", 2+2, "is 4",
		"\ndoes new line work too")
	fmt.Printf("My age is %v", 24+1)
	fmt.Print("\nMars\n")
	fmt.Printf("%-20v £%5v\n", "Apples", 2)
	fmt.Printf("%-20v £%5v\n", "Oranges", 1.5)

	// Mars - how long does it take to travel?

	const lightspeed = 299792
	var distance = 56000000

	fmt.Println(distance/lightspeed, "seconds")
	distance = 401000000
	fmt.Println(distance/lightspeed, "seconds")

}

func ab_imperatives() {
	var module = "calculator"
	var line_break = "\n------------------------------------\n"
	fmt.Println(line_break, module, line_break)

	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	rand.Seed(time.Now().UnixNano())
	min := 56000000
	max := 401000000
	fmt.Println(rand.Intn(max-min+1) + min)
	fmt.Println(max)
	fmt.Println(min)
}

/*func ba_types() {
	var module = "module 2: Types"
	var line_break = "\n------------------------------------\n"
	fmt.Println(line_break, module, line_break)
}*/

func ac_imperatives() {
	var module = "loops and branches"
	var line_break = "\n------------------------------------\n"
	fmt.Println(line_break, module, line_break)

	fmt.Println("You find yourself in a dimly lit cavern.")
	var commanda = "walk outside"
	var exit = strings.Contains(commanda, "outside")
	fmt.Println("You leave the cave:", exit)

	fmt.Println("Today is a Wednesday")
	var agenda = "Arabic studies"
	var completed = strings.Contains(agenda, "Arabic")
	fmt.Println("I studied arabic today:", completed)

	fmt.Println("above 18s only")
	var age = 41
	var limit = age > 18
	fmt.Printf("At age %v, do I pass the limit? %v\n", age, limit)

	var command = "go east"

	if command == "go east" {
		fmt.Println("go to see the pyramids")
	} else if command == "go inside" {
		fmt.Println("enter the pyramids")
	} else {
		fmt.Println("what?")
	}

	fmt.Println("The year is 2100, should you leap?")
	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}

	//The not logical operator (!) flips a Boolean value from false to true or vice versa.

	var haveTorch = true
	var litTorch = false
	if !haveTorch || !litTorch {
		fmt.Println("Nothing to see here.")
	}

	//switch
	fmt.Println("Will you enter the cave or go further east?")
	var cmdaa = "go inside"

	switch cmdaa {
	case "go east":
		fmt.Println("you go further east into the jabal")
	case "enter cave", "go inside":
		fmt.Println("you are in the dimly lit cave")
	case "read sign":
		fmt.Println("the sign reads \"no minors\".")
	default:
		fmt.Println("what?")
	}

	fmt.Println("\n")
	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("you find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough")
		fallthrough
	case room == "underwater":
		fmt.Println("the water is freezing cold")
	}

	fmt.Println("\n")

	var count1 = 10

	for count1 > 0 {
		fmt.Println(count1)
		//time.Sleep(time.Second / 3)
		count1--
	}

	fmt.Println("liftoff!")

	fmt.Println("\n")

	var deg1 = 0

	for {
		fmt.Println(deg1)
		deg1++
		if deg1 >= 36 {
			deg1 = 0
			if rand.Intn(1) == 0 {
				break
			}
		}
	}
	//	fmt.Println("\n")

	var mod4 = "variable scope"
	fmt.Println(line_break, mod4, line_break)
	//	fmt.Println("\n")
	//	fmt.Println("\n")

	var count2 = 0

	for count2 = 10; count2 > 0; count2-- {
		fmt.Println(count2)
	}

	fmt.Println(count2)

	for count2 := 10; count2 > 0; count2-- {
		fmt.Println(count2)
	}

	rand.Seed(time.Now().UnixNano())
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}

}

// different scopes
var era = "AD" //global scope

func scopes() {
	year := 2018                               //within scope
	switch month := rand.Intn(12) + 1; month { // month within scope
	case 2:
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
	//month and day are out of scope

}

// year is out of scope

func scope2() { //removing dublicates (code smell)
	year := 2018
	month := rand.Intn(12) + 1
	days_m := 31

	switch month {
	case 2:
		days_m = 28
	case 4, 6, 9, 11:
		days_m = 30
	}

	day := rand.Intn(days_m) + 1
	fmt.Println(era, year, month, day)
}

func capstone() {
	ch = "capstone 1"
	fmt.Println(lb, ch, lb)

	fmt.Printf("\n%-10v %5v %10v %5v %2v\n", "airline", "days", "Ticket", "", "Price")
	fmt.Println("=======================================")

	for i := 1; i <= 10; i++ {

		var airline string
		dmin := 20
		dmax := 40
		var days = rand.Intn(dmax-dmin) + dmin
		var tType1 = "One-way"
		var tType2 = "Round-trip"
		var tType string
		pmin := 30
		pmax := 100
		var price = rand.Intn(pmax-pmin) + pmin
		airnum := rand.Intn(3) + 1

		switch airnum {
		case 1:
			airline = "spaceX"
		case 2:
			airline = "vGalactic"
		case 3:
			airline = "sAdvent"
		}

		if price < 55 {
			tType = tType1
		} else {
			tType = tType2
		}

		fmt.Printf("%-10v %5v %12v %5v %2v\n", airline, days, tType, "£", price)
	}

}

func realnums() {
	ch = "Real numbers (pg. 60)"
	fmt.Println(lb, ch, lb)

	days := 365.2
	var days2 = 365.24
	var days3 float64 = 365.2425
	var inferflt float64 = 42

	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi32, pi64)
	fmt.Println(days, days2, days3, inferflt)

	fmt.Printf("%v\n", pi64)
	fmt.Printf("%f\n", pi64)
	fmt.Printf("%.3f\n", pi64)
	fmt.Printf("%4.2f\n", pi64)
	fmt.Printf("%05.2f\n", pi64)

	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "º F\n")
	fmt.Print((9.0/5.0*celsius)+32, "º F\n")
	farenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Print(farenheit, "º F\n")

	fmt.Printf("\n\ncomparing floating-point numbers%v", lb)
	savings := 0.1
	savings += 0.2
	fmt.Println(savings == 0.3)
	fmt.Println(math.Abs(savings-0.3) < 0.0001)

	fmt.Println("Changed Git Directory")

}

func types() {
	ch = "Real numbers (pg. 60)"
	fmt.Println(lb, ch, lb)

	//var year_1 int = 2018
	//var month_1 uint = 2
	//year_3 := 2018
	//var year_4 = 2018
	//var year_5 int = 2018

	fmt.Println("There are many integer types \nint integers are neg to pos \n while unit integers are positive \n")

	fmt.Println("Interger types:\nint8 'uint8', 'int16', 'uint16', 'int32', 'unit32', 'int64', 'uint64 ")

	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days)

	fmt.Printf("Type %T for %[1]v\n", "purity")

	var red, green, blue uint8 = 255, 141, 213
	println(red, "\n", green, "\n", blue)

	red++
	fmt.Println(red)

	var number int8 = 127
	number++
	fmt.Println(number)

	green = 3
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)

	blue = 255
	blue++
	fmt.Printf("%08b\n", blue)

	future := time.Unix(12622780800, 0)
	fmt.Println(future)
}

func big_numbers() {
	const lightSpeed = 2999792 //km/s
	const secondsPerDay = 86400
	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away.")
	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")
}

func big_numbers2() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away.")
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println("That is", days, "days of travel at light speed.")

}

func big_nums3() {
	//consts are untyped
	const distance = 24000000000000000000
	const lightSpeed = 299792   //km/s
	const secondsPerDay = 86400 //s/d
	//km := distance //will overflow (err)
	days := distance / lightSpeed / secondsPerDay

	fmt.Println("Andromeda Galaxy is", days, "light days away")
	//const values can be assigned to variables as long as they fit.

	//Untyped constants must be converted to typed variables when passed to functions.

}

func multi_text() {
	peace1 := "peace"
	var peace2 = "Peace"
	var peace string = "peace"

	var blank string
	blank = "so much peace"
	fmt.Println(peace, peace1, peace2, blank)
	fmt.Println("Peace be upon you\nUpon you be peace")
	fmt.Println(`Strings can span multiple lines with the \n escape sequence`)
	fmt.Println(`
Peace be upon you
	upon you be peace
	
	`)

	fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)

}

func runes_n_bytes() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v \n", pi, alpha, omega, bang)
	fmt.Printf("%c %c %c %c \n", pi, alpha, omega, bang)

	g1 := 'A'
	var g2 = 'A'
	var g3 rune = 'A'
	var star byte = '*'

	fmt.Printf("%c %c %c %c \n", g1, g2, g3, star)
}

func more_types() {

	peace := "shalom"
	peace = "salām"
	fmt.Printf("%v \n", peace)

	message := "salām"
	c := message[5]
	fmt.Printf("%c\n", c)
	fmt.Printf(`
strings are immutable so
message[5] = 'd' 
is not possible

`)

}

func caesar_cypher() {

	/*
			c := 'a'
		c = c - 3
		fmt.Printf("%c\n", c)

			if c > 'z' {
				c = c - 26
			}
	*/

	message := `fdph
vdz
frqtxhuhg
	`
	for i := 0; i < 20; i++ {
		c := message[i]
		if c < 'a' {
			c = c + 26
		}
		c = c - 3
		fmt.Printf("%c\n", c)

	}

}

func main() {

	rand.Seed(time.Now().UnixNano())

	aa_imperatives()
	ab_imperatives()
	ac_imperatives()
	scopes()
	scope2()
	capstone()
	// Global variables
	realnums()
	//ba_types()
	types()
	big_numbers()
	big_numbers2()
	big_nums3()
	multi_text()
	runes_n_bytes()
	more_types()
	caesar_cypher()
}

var ch string
var lb string = "\n----------------------------------------\n"
