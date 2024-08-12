package hotel

import (
	"testing"
	"time"
)

// Helper function to reset the rooms to their initial state
func resetRooms() {
	for _, room := range Rooms {
		room.cancelRoom()
	}
}

// Test creation of new room instances
func TestNewRoom(t *testing.T) {
	room := NewRoom(101)
	if room.RoomNumber != 101 {
		t.Errorf("Expected room number 101, got %d", room.RoomNumber)
	}
	if room.Reserved {
		t.Errorf("Expected room not to be reserved")
	}
	if room.GuestName != "" {
		t.Errorf("Expected guest name to be empty")
	}
}

// Test reservation of rooms
func TestReserveRooms(t *testing.T) {
	resetRooms()

	// Reserve a room
	err := reserveRooms("john")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the room is reserved
	if Rooms[0].GuestName == "" {
		t.Errorf("Expected room 101 to be reserved")
	}

	// Check reservations when all rooms are reserved
	err = reserveRooms("doe")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = reserveRooms("alex")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = reserveRooms("cael")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = reserveRooms("ryle")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = reserveRooms("atlas")
	if err == nil {
		t.Errorf("Expected error for all rooms occupied")
	}
}

// Test cancelation of reservation
func TestCancelReservation(t *testing.T) {
	resetRooms()

	// Reserve a room
	err := reserveRooms("john")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Cancel the reservation
	err = cancelReservation(101)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the room is available
	if Rooms[0].checkReserved() {
		t.Errorf("Expected room 101 to be available")
	}

	// Cancel the reservation for non-existing room
	err = cancelReservation(10097)
	if err == nil {
		t.Errorf("Expected room to not exist")
	}

	// Cancel the reservation for non-reserved room
	err = cancelReservation(105)
	if err == nil {
		t.Errorf("Expected room to be not reserved")
	}
}

// Test view room details
func TestViewRoomDetails(t *testing.T) {
	resetRooms()

	// Reserve a room
	err := reserveRooms("john")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check room details
	err = viewRoomDetails(101)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test for a non-existing room
	err = viewRoomDetails(10091)
	if err == nil {
		t.Errorf("Expected error for non-existing room")
	}

	// Test for a non-reserved room
	err = viewRoomDetails(105)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

// Test Check Vacant Rooms
func TestCheckVacantRooms(t *testing.T) {
	resetRooms()

	// Reserve a room
	err := reserveRooms("john")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	vacantRooms := checkVacantRooms()
	if len(vacantRooms) != 4 {
		t.Errorf("Expected 4 vacant rooms, got %d", len(vacantRooms))
	}

	// Cancel the reservation
	err = cancelReservation(101)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check vacant rooms
	vacantRooms = checkVacantRooms()
	if len(vacantRooms) != 5 {
		t.Errorf("Expected 5 vacant rooms, got %d", len(vacantRooms))
	}
}

// Test check vacant rooms periodically
func TestCheckVacantRoomsPeriodically(t *testing.T) {
	resetRooms()

	done := make(chan bool)
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		checkVacantRoomsPeriodically(ticker, done)
	}()

	// Test checking vacant rooms periodically
	time.Sleep(2 * time.Second)
	done <- true
}
