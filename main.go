package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Basicamente es un objeto
type car struct {
	brand string
	year  int
}

func main() {

	// ? Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1416 // Infiere el tipo

	fmt.Println("Hola Mundo!!!")
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// ? Declaración de variables
	base := 12          // Al agregar el : crea la variable y la inicializa, basicamente asi declaramos variables
	var altura int = 14 // Aca asiganmos un tipo a la variable
	var area int        // Aqui podemos crear variables sin asignar valor

	fmt.Println(base, altura, area)

	// ? Zero values(Cuando las creamos tienen valores iniciales)
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// ? Calcular area de cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * 2
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// ? Suma
	result := x + y
	fmt.Println("Suma:", result)

	// ? Resta
	result = y - x
	fmt.Println("Resta:", result)

	// ? Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	// ? División
	result = y / x
	fmt.Println("División:", result)

	// ? Módulo
	result = y % x
	fmt.Println("Módulo:", result)

	// ? Incremental
	x++
	fmt.Println("Incremental:", x)

	// ? Decremental
	x--
	fmt.Println("Decremental:", result)

	// ? Valores primitivos
	// Numeros enteros
	// int = Depende del OS (32 o 64 bits)
	// int8 = 8bits = -128 a 127
	// int16 = 16bits = -2^15 a 2^15-1
	// int32 = 32bits = -2^31 a 2^31-1
	// int64 = 64bits = -2^63 a 2^63-1

	// Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	// uint = Depende del OS (32 o 64 bits)
	// uint8 = 8bits = 0 a 127
	// uint16 = 16bits = 0 a 2^15-1
	// uint32 = 32bits = 0 a 2^31-1
	// uint64 = 64bits = 0 a 2^63-1

	// Números decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	// Textos y booleanos
	// string = ""
	// bool = true or false

	// Números complejos
	// Complex64 = Real e Imaginario float32
	// Complex128 = Real e Imaginario float64
	// Ejemplo : c:=10 + 8i

	// ? Librería fmt
	helloMessage := "Hello"
	worldMessage := "World"

	// El ln hace salto de linea
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// printf agrega una función dónde podemos especificar impresión de tipos de datos
	nombre := "Platzi"
	cursos := 500

	// %s => string, %d => entero , %v => cualquier tipo de dato, \n => Salto de línea
	fmt.Printf("%s tiene más de %d cursos, variable dinámica: %v \n", nombre, cursos, nombre)
	fmt.Printf("%s tiene más de %d cursos, variable dinámica: %v \n", nombre, cursos, cursos)

	// Sprintf guarda en una variable la impresión
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// ! Podemos saber el tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	// ? Uso de funciones
	printMessage("Estoy usando una función!!!")
	value := returnValue(2)
	fmt.Println(value)

	value1, value2 := doubleReturn(2)
	fmt.Printf("value1 => %d, value2 => %d \n", value1, value2)

	// ? Ciclo for
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// ? For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// ? For forever
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	// ? If
	valor1 := 1
	fmt.Println(valor1 == 1 && valor1 != 1)

	// ? Convertir texto a número
	// En value guarda el valor y en err guarda el error, el método Atoi regresa esos dos
	value, err := strconv.Atoi("53")
	// Si el error es diferente de nil, significa que existe y hubo un error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

	// ? Switch Case => Iterar sobre una misma variable
	// La primera parte indica como se define la varialbe
	// La segunda parte indica el parseo de la variable
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// ? Switch sin condición => Para anidar multiples condiciones
	switchValue := 200
	switch {
	case switchValue > 100:
		fmt.Println("Es mayor a 100")
	case switchValue < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condición")
	}

	// ? Defer
	// Hace que justo antes de morir el código, se ejecute la función
	defer fmt.Println("Hola")
	defer fmt.Println("???")
	fmt.Println("Mundo")

	// ? Array => INMUTALBE
	var array [4]int
	array[0] = 1
	array[1] = 2
	// len => Longitud del array
	// cap => Capacidad máxima del array
	fmt.Println(array, len(array), cap(array))

	// ? Slice => MUTABLE
	// No se indica la cantidad de valores que va a tener
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice, len(slice), cap(slice))

	// ? Métodos en el slice
	// ! NOTA: En [2:4] el primer elemento es inclusivo, es decir aparecerá en la impresión, mientras que
	// ! el segundo es exclusivo, es decir el cuarto elemento no aparecerá
	// ! Por ejemplo [2(inclusivo),4(exclusivo)] por lo que se imprimirá el índice 2 y 3 (el 4 ya no por ser exclusivo)
	// Imprimimos el elemento del índice 0
	fmt.Println((slice[0]))
	// Imprimimos todos los elementos desde el índice 0 hasta el 3 (recordar que no se imprime el índice 3 ya que es exclusivo en [:3])
	fmt.Println((slice[:3]))
	// Imprimimos todos los elementos desde el elemento 2 al 4 (se imprimirá el índice 2 pero no el 4)
	fmt.Println((slice[2:4]))

	// ? Append
	// Le agregamos un 3 al array ó slice
	slice = append(slice, 3)
	fmt.Println(slice)

	// ? Append nueva list
	newSlice := []int{8, 9, 10}
	// ! Los 3 puntos indica que va a tomar cada elemento de newSlice y de forma independiente lo va a agregar
	slice = append(slice, newSlice...)

	// ? Recorrido de Slices con Range
	stringSlice := []string{"hola", "que", "hace"}

	// ! range  se utiliza para iterar los elementos, devuelve tanto el índice como el valor, es decir es una función con dos valores en su retorno
	for indice, valor := range stringSlice {
		fmt.Println(indice, valor)
	}

	// ? Definir si un string es un palíndromo (se lee de la misma forma al derecho y revés)
	isPalindromo("ama")
	isPalindromo("Ama")
	isPalindromo("amas")

	// ? Maps
	// Make en este caso sirve para crear diccionarios pero puede crear más cosas
	// Make va a crear un map, cuyas llaves son en string y valor int
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// ? Recorrer map => No recorre en orden el map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// ? Encontrar valor
	valueJose := m["Jose"]
	fmt.Println(valueJose)
	fmt.Println(m["Jose"])

	// ! Si intentamos acceder a un valor que no existe, nos retornará 0
	fmt.Println(m["Josasdfsadfe"])
	// ! Nos podemos asegurar que exista el valor o no con el segundo retorno
	valuePepito, exists := m["a"]
	fmt.Println(valuePepito, exists)

	// ? Instanciar structs(objetos)
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// ? Otra forma de instanciar
	var otherCar car
	otherCar.brand = "Ferrari"
	// Si no especificamos, tomará el valor por defecto
	fmt.Println(otherCar)
}

func printMessage(message string) {
	fmt.Println(message)
}

func tripleArg(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

// El primer paréntesis es para indicar lo que recibe la función
// Si queremos indicar que se retorne dos parámetros lo ponemos entre paréntesis
func doubleReturn(a int) (return1, return2 int) {
	// Retornamos dos enteros como específicamos
	return a, a * 2
}

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		// Lo convertimos en texto porque llega como código ASCII
		textReverse += string(text[i])
	}

	text = strings.ToLower(text)
	textReverse = strings.ToLower(textReverse)

	if text == textReverse {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es un palíndromo")
	}
}
