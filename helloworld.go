package main

import "fmt"

func main() {
	fmt.Println("hello world !")
	//fmt.Println("ini satu", 1)
	//fmt.Println("ini benar", true)
	//fmt.Println("jumlah string : ", len("test"))
	//fmt.Println("karakter ke 2 : ", "test"[2])

	/*
							var name string
							name = "ricky lai"
							fmt.Println(name)

							futurName := "ricky Martin"
							fmt.Println(futurName)
							fmt.Println(futurName[0])
							fmt.Println(string(futurName[0]))

							//var nilaiInteger int32 = 33
							//var nilaiInteger2 int32 = 33
							var total = 0
							total = 10 + 1
							fmt.Println(total)

							//array
							var arraynama [2]string
							arraynama[0] = "ricky"
							arraynama[1] = "lai"
							fmt.Println(arraynama[0])

							//array angka
							var arrayvalue = [3]int{1, 2, 3}
							fmt.Println(arrayvalue[2])


						//map
						person := map[string]string{
							"name":    "ricky",
							"address": "cikupa",
						}

						println(person["name"])


					person := "ricky1"

					if person == "ricky" {
						println(person)
					}


				value := 0

				for value < 10 {
					value++
					println("value : ", value)
				}

			arrayNama := []string{"ricky", "lai"}

			for i := 0; i < len(arrayNama); i++ {
				println(arrayNama[i])
			}



		sayhello()
	*/

	sTampung := returnvalue("ricky lai")
	println(sTampung)

}

func sayhello() {
	println("say Helloooo !!!")
}

func returnvalue(sName string) string {
	sReturnname := "Bye " + sName
	return sReturnname
}
