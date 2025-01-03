package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Data model
type Users struct {
	UserFirstName string    `bson:"firstName"`
	UserLastName  string    `bson:"lastName"`
	UserEmail     string    `bson:"email"`
	UserZipCode   uint      `bson:"zipCode,omitempty"`
	UserMobile    uint      `bson:"mobile,omitempty"`
	CreatedAt     time.Time `bson:"createdAt"`
}

// type Tickets struct{

// }

var conferenceName = "GoLang Training Conference"

var firstName string
var lastName string
var email string
var zipcode uint
var mobile uint
var client *mongo.Client
var collection *mongo.Collection
var teamName string

//	Planning for future development

//	Do user save and tickets saved in DB.
// To-Do
// Connect MonogoDB or some other (Done Connection)  // in DB Int values are storing but string values are not storing in DB Resolve this error.
// Do future analysis
// Right now it is handling the booking of ticket instead of this make sure to Make a menu list user can navigate to book ticket and Cancel ticket...and so on
// Decide a client Archietecture for this TicketBooking System.
//Studied about MVC pattern
//Need to explore more features
// Need To Resolve The DB Error , Studying the DB Types in MongoDB 
//Do something about storing the Ticket Details in DB.
//Get some details about new features which will be used in this application.

func bookTickets(userTickets uint, remainingTickets uint) uint {
	fmt.Println("Processing....")
	remainingTickets = remainingTickets - userTickets
	fmt.Println("Booking Completed...")
	return remainingTickets
}
func getUserFromDB(db string, table string, name string) {
	collection := client.Database(db).Collection(table)
	// Find a document
	var result Users
	err := collection.FindOne(context.Background(), name).Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User Found:%v\n", result)
}
func storeUserInDB(db string, table string) {
	collection := client.Database(db).Collection(table)
	newUser := Users{
		UserFirstName: firstName,
		UserLastName:  lastName,
		UserEmail:     email,
		UserZipCode:   zipcode,
		UserMobile:    mobile,
		CreatedAt:     time.Now(),
	}

	fmt.Printf("USER GOING TO DB:%v\n", newUser)

	// Insert into MongoDB
	result, err := collection.InsertOne(context.Background(), newUser)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("User Added with Object_Id:%v\n", result.InsertedID)
	}
}

func getCost(userTickets uint, price float32) float32 {
	//get Ticket Cost for differenet type of Ticket booking
	price = price * float32(userTickets)
	return price
}
func getUserDetails() {
	fmt.Printf("Hi Greeting From %v Team :\n", conferenceName) //Greet Customer
	//Get User Details via function needs to implement
	fmt.Println("Enter your first name : ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email : ")
	fmt.Scan(&email)
	fmt.Println("Enter your mobile : ")
	fmt.Scan(&mobile)
	fmt.Println("Enter your postal-code : ")
	fmt.Scan(&zipcode)
	fmt.Println("Enter your  teamName :")
	fmt.Scan(&teamName)
}
func connectMongoDB(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Failed to connect MongoDB")
		return
	}
	fmt.Println("Connected to MongoDB!")
}

// implement features in features branch
func main() {
	//Dot-Env File Handling
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("NOenv File.")
	} else {
		fmt.Println("Got your .env")
	}
	mongo_uri := os.Getenv("MONGO_DB_URI")
	mongo_db := os.Getenv("DB_HOST")
	userTable := os.Getenv("MONGO_DB_USER_TABLE")
	connectMongoDB(mongo_uri)
	var userTickets uint
	var ticketCost float32
	var seatType string
	var validBooking bool
	var taxAmount float32
	var taxValue = 3.5
	// Conference data and collection of bookings
	var totatEarnings float32
	// Use Slices to handle the user details to store everyone's details
	const totalTickets = 7
	var remainingTickets uint = totalTickets
	var bookingUsers []string //Dynamic array is called as slices
	var counter int = 0
	for remainingTickets > 0 {
		//Dynamic Array slices usage
		getUserDetails()

		bookingUsers = append(bookingUsers, firstName+" "+lastName)
		// Notifying user about about ticket details
		fmt.Printf("\n%v Welcome to ticket booking!\n", bookingUsers[counter])
		fmt.Printf("Total number of tickets %v remaining are %v\n", totalTickets, remainingTickets)
		fmt.Println("Grab your tickets before it solds out")
		//ticketBooking Logic
		fmt.Println("Enter number of ticket to be booked:")
		fmt.Scan(&userTickets) //use if conditionals to handle 0 to outOfRange tickets to book
		var limitofUserTicket uint = 6
		if userTickets > limitofUserTicket || userTickets <= 0 {
			fmt.Println("You entered wrong number of Ticket")
			validBooking = false
			fmt.Printf("The lower limit is 0 and upper limit of ticket booking is %v\n", limitofUserTicket)
			fmt.Println("Booking failed...Try Again")
		} else {
			validBooking = true
		}
		if userTickets <= remainingTickets && validBooking {
			//Try user to select a normal seat or premium seats based on their need.
			fmt.Println("Enter N for Normal Seat  or P for Premium Seat :")
			fmt.Scan(&seatType)
			for {
				if seatType == "N" || seatType == "P" {
					break
				} else {
					fmt.Println("Invalid Seat Selection")
					fmt.Println("Enter N for Normal Seat - No Tax  or P for Premium Seat Includes Tax (Case-Sensitive) :")
					fmt.Scan(&seatType)
				}
			}
			switch seatType {
			case "N":
				fmt.Println("Normal Seat Selected")
				ticketCost = 199
				ticketCost = getCost(userTickets, ticketCost)
				totatEarnings += ticketCost
			case "P":
				fmt.Println("Premium Seat Selected")
				ticketCost = 299
				taxAmount = (ticketCost * float32(taxValue)) / 100
				fmt.Printf("Tax For Selceted Seat : %v\n", taxAmount)
				ticketCost = getCost(userTickets, ticketCost+taxAmount)
				totatEarnings += ticketCost
			default:
				fmt.Println("Invalid Seat Selection")
			}
			remainingTickets = bookTickets(userTickets, remainingTickets)
			fmt.Println("-------------------------------------Ticket Summary-------------------------------------")
			fmt.Printf("\n\tCustomer Name :%v \n \tTickets Booked : %v \n \tAvailable tickets after booking  :%v\n \tTickets sent to e-mail : %v\n \tContact-No : %v\n \tPostal-Code %v\n \tTotalCost : %v\n", bookingUsers[counter], userTickets, remainingTickets, email, mobile, zipcode, ticketCost)
			counter += 1
			storeUserInDB(mongo_db, userTable)
			fmt.Println("Happy Ticket Booking...")
			if remainingTickets == 0 {
				fmt.Println("Conference is full,Try next year...")
				break
			}
		} else if userTickets == totalTickets {
			fmt.Printf("You cannot book all tickets once ...\n")
			fmt.Print("Booking Failed...Try Again\n")

		} else if userTickets > remainingTickets && !validBooking {
			fmt.Printf("We only have %v remaining,you can't book %v tickets...\n", remainingTickets, userTickets)
		}
	}
	name := bookingUsers[0]
	getUserFromDB(mongo_db, userTable, name)
	fmt.Printf("Total Earning's after Bookings :%v\n", totatEarnings)
	fmt.Println("Total User Bookings:", len(bookingUsers))
	fmt.Println("Mikka Nandri...Thank You...Arigato Gosaimasu")
}
