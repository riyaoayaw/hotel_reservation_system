package hotel

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

// hotel with 5 rooms
// 3 people booking rooms
// random order
// only one at a time
// systems checks for room availabitiy
// each room has its name
// reservation list to indicate which guest is reserved for rooms
// cron - task which happens periodically - 1 or 2 mins
// cron checks and tells which room is vacant
// can withdraw or cancel reservation - room is vacant

// Room struct represents a room
type Room struct {
	sync.Mutex
	RoomNumber int
	Reserved   bool
	GuestName  string
}

// NewRoom creates a new instance of Room
func NewRoom(number int) *Room {
	return &Room{
		RoomNumber: number,
		Reserved:   false,
		GuestName:  "",
	}
}

// checkReserved checks if the room is reserved
func (r *Room) checkReserved() bool {
	r.Lock()
	defer r.Unlock()

	return r.Reserved
}

// bookRoom reserves the room for a guest
func (r *Room) bookRoom(name string) {
	r.Lock()
	defer r.Unlock()

	r.Reserved = true
	r.GuestName = name
}

// cancelRoom cancels reservation for the room
func (r *Room) cancelRoom() {
	r.Lock()
	defer r.Unlock()

	r.Reserved = false
	r.GuestName = ""
}

// create the rooms of the hotel
var Rooms = []*Room{
	NewRoom(101),
	NewRoom(102),
	NewRoom(103),
	NewRoom(104),
	NewRoom(105),
}

// checkVacantRooms checks for vacant rooms and returns the list of vacant rooms
func checkVacantRooms() []*Room {
	var vacantRooms []*Room
	for _, room := range Rooms {
		if !room.checkReserved() {
			vacantRooms = append(vacantRooms, room)
		}
	}

	return vacantRooms
}

// checkVacantRoomsPeriodically checks the vacant rooms every 1 minute and stop when a signal is passed
func checkVacantRoomsPeriodically(ticker *time.Ticker, done <-chan bool) {
	for {
		select {
		case <-ticker.C:
			checkVacantRooms()
		case <-done:
			return
		}
	}
}

// reserveRooms checks for vacant rooms and reserves a room for the guest
func reserveRooms(name string) error {
	vacantRooms := checkVacantRooms()

	if len(vacantRooms) == 0 {
		return errors.New("no rooms available for reservation")
	}

	room := vacantRooms[0]
	room.bookRoom(name)

	fmt.Printf("\nRoom %d successfully reserved for guest %s", room.RoomNumber, room.GuestName)
	return nil
}

// cancelReservation cancels a room reservation
func cancelReservation(roomNo int) error {

	for _, room := range Rooms {
		if room.RoomNumber == roomNo {
			if !room.checkReserved() {
				return errors.New("Room is not reserved")
			} else {
				room.cancelRoom()
				checkVacantRooms()
				fmt.Printf("\nReservation for room %d has been canceled.\n", room.RoomNumber)
				return nil
			}
		}
	}

	return errors.New("Room number not found")
}

// viewRoomDetails shows details of a specific room
func viewRoomDetails(roomNo int) error {

	for _, room := range Rooms {
		if room.RoomNumber == roomNo {
			if room.checkReserved() {
				fmt.Printf("\nRoom %d is reserved by %s", room.RoomNumber, room.GuestName)
			} else {
				fmt.Printf("\nRoom %d is available.\n", room.RoomNumber)
			}
			return nil
		}
	}

	return errors.New("Room number not found")
}

func StartHotelRun() {

	// Create a ticker to trigger every 1 minute
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	// Channel to signal stopping the Go routine
	done := make(chan bool)

	// Start the Go routine
	go checkVacantRoomsPeriodically(ticker, done)

	reader := bufio.NewReader(os.Stdin)

	ch := "Y"
	for ch == "Y" || ch == "y" {
		var choice int

		fmt.Println("\n\n---------Welcome to HOTEL GOLANG---------")
		fmt.Println("\n1. Reserve a room")
		fmt.Println("2. View Room Details")
		fmt.Println("3. Cancel your reservation")
		fmt.Println("4. Exit")
		fmt.Print("\nEnter your choice: ")

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch choice {

		case 1:
			var name string

			fmt.Print("\nEnter guest name: ")
			name, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("\nError reading input:", err)
			}

			err = reserveRooms(name)
			if err != nil {
				fmt.Println("\nError reserving a room:", err)
			}

		case 2:
			var roomNo int

			fmt.Print("\nEnter room number to view details: ")
			_, err := fmt.Scanln(&roomNo)
			if err != nil {
				fmt.Println("\nError reading input:", err)
			}

			err = viewRoomDetails(roomNo)
			if err != nil {
				fmt.Println("\nError viewing room details:", err)
			}

		case 3:
			var roomNo int

			fmt.Print("\nEnter room number for cancellation: ")
			_, err := fmt.Scanln(&roomNo)
			if err != nil {
				fmt.Println("\nError reading input:", err)
			}

			err = cancelReservation(roomNo)
			if err != nil {
				fmt.Println("\nError cancelling a reservation:", err)
			}

		case 4:
			done <- true
			return

		default:
			fmt.Println("\nInvalid Input")
		}

		fmt.Print("\nGo back to menu? (Y/N): ")
		fmt.Scanln(&ch)
	}
}
