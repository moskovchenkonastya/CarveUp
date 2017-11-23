package main

import (
	"encoding/json"
	"fmt"
	"time"
	"os"
	"io/ioutil"
	"math/rand"
)

const fileDepot string = "./data/depot.json"
// struct of Depot Coordinates
type DepotCoordinates struct{
	X float32 `json:"x, float32"`
	Y float32 `json:"y, float32"`
}

const fileCoordinates string = "./data/coordinates.json"
// struct of Coordinates
type Coordinates struct {
	Id         int
	hStart	   int
	minStart   int
	XStart     float32
	YStart     float32
	hEnd	   int
	minEnd     int
	XEnd       float32
	YEnd       float32
}

const fileBuses string = "./data/buses.json"
type Buses struct {
	Id       int
	MinTime  int
	Capacity int 
	Price    int
}

type Order struct {
	Id 			int
	Distance	int
	Duration	int
}

type Time struct {
	hour    int
	min  int 
}
// cчитывание файлов json
func readData(filepath string) []byte {

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	json, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	//fmt.Print(string(json))
	return json;
}

func funcRandom() int {
	random := rand.Float64()
	if random <= 0.7 {return 1}
	if random > 0.7 && random <= 0.9 {return 2}
	return 3
}
const numberOrders = 50
const maxNumberIterations =  10


func main() {
	
	// cчитывание Depot Coordinates
	var dataDC []DepotCoordinates
	data := readData(fileDepot)
	err := json.Unmarshal(data, &dataDC)
	if err != nil {
	  fmt.Println("error: Сan not unmarshal filedepot")
	}
	//fmt.Printf( "%#v\n", dataDC);
	
	// считывание данных с coordinates.json
	var dataC []Coordinates
	data = readData(fileCoordinates)
	err = json.Unmarshal(data, &dataC)
	if err != nil {
		fmt.Println("error: Сan not unmarshal fileсoordinates")
	}
	// fmt.Printf("json string:\n\t%#v\n", dataC[0])

	// считывание данных с buses.json
	var dataB []Buses
	data = readData(fileBuses)
	err = json.Unmarshal(data, &dataB)
	if err != nil {
		fmt.Println("error: Сan not unmarshal filebuses")
	}
	// fmt.Printf("json string:\n\t%#v\n", dataB[0])

	// Инициализация
	// static var
	
	//Dynamic variables
	
	iterationCounter := 0
	routesNumber := -1 //Number of routes after each iteration, Ni
	numberOrdersRoute := -1 //Number of orders in route after each iteration of optimizing algorithm, Lni
	maxNumberRoutes := -1 //max Ni
	
	var passengers [numberOrders]float32
	// Number of passengers using random function with distribution: 1 if <= 0.7; 2 if <= 0.2; 3 else
	for passenger := range passengers {
		passenger = funcRandom()
		fmt.Print(passenger)
	}

	// считаем расстояния и время в заказах
	//	json c 
	//
	// инициализация
	var unprocessedRoutes [numberOrders]int
	var admissibleSet [numberOrders]int
	var departureTimeSuitable [numberOrders]Time
	var arrivalTimeSuitable [numberOrders]Time
	var suitableRoutesTime [numberOrders]time.Time
	busesRentalPrice := 0
	totalRoutePassengers := 0
	stillUnprocessed := true
	minDepartureTime := 0

	for i := 0; i <= maxNumberIterations; i++ {
		//Initialization of variables
		totalRoutePassengers = 0
		routesNumber = -1
		stillUnprocessed = true

		//  обнуление значений
		for k := 0; k < numberOrders; k++ {
			unprocessedRoutes[k] = k
			admissibleSet[k] = 1
			departureTimeSuitable[k].hour = 0
			departureTimeSuitable[k].min = 0
			arrivalTimeSuitable[k].hour = 0
			arrivalTimeSuitable[k].min = 0
			suitableRoutesTime[k] = time.Date(2017, 10, 12, 0, 0, 0, 0, time.UTC)
			busesRentalPrice = 0
		}

		for stillUnprocessed && (routesNumber < (numberOrdersRoute - 1)) {
			routesNumber++;

			for k := 0; k < numberOrders; k++ {
				admissibleSet[k] = 1;
			}
			// хз что тут происходит  тип новую память выделяем 
			// не понятно нужно ли все это ???
			if maxNumberRoutes < routesNumber {
				maxNumberRoutes = routesNumber;
				//+2 because of depot
				/*System.out.println(Max_Number_Of_Routes);
				var optimizedRoutesCoordinates[maxNumberRoutes + 2]DepotCoordinates 
				var optimizedDepartureTime[maxNumberRoutes + 2]float32 // time
				var optimizedRoutePassengers[numberOrders + 2]int
				// var optimizedRoutesNumbers[Max_Number_Of_Routes] = new double[Number_of_Orders + 2]
				*/
			}
			
			//Depot counting
			numberOrdersRoute = 0
			for k := 0; k < numberOrders; k++ {
				if(unprocessedRoutes[k] != -1){
					minDepartureTime = k;
					k = numberOrders;
				}
			}

			//Choose the order with the earliest departure time
			for k := 0; k < numberOrders; k++ {
				//Only for unprocessed routes
				if(unprocessedRoutes[k] != -1){
					if dataC[minDepartureTime].hStart > dataC[k].hStart && dataC[minDepartureTime].minStart > dataC[k].minStart {
						minDepartureTime = k;
					}
				}
			}
	}	
}
}
