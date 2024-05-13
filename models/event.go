package models

import (
	"example/rest_api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	Datetime    time.Time
	UserId 		int64
}

// var events = []Event{}

func (e *Event) Save() error{
	query := `
		INSERT INTO events(name, description, location, datetime, user_id)
		VALUES(?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err!=nil {
		return err
	}	
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserId) //Use .Prepare + .Exec for update, insert, delete data
	if err!=nil {
		return err
	}
	
	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) //Use .Query for get data
	if err!=nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserId)
		
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error){
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, datetime = ?
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.Datetime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := `
		DELETE FROM events WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()	

	_, err = stmt.Exec(event.ID)

	return err
}

func (event Event) Register(userId int64) error{
	query := `INSERT INTO registrations(event_id, user_id) VALUES (?,?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()	
	_, err = stmt.Exec(event.ID, userId)

	return err
}

func (event Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()	
	_, err = stmt.Exec(event.ID, userId)

	return err	
}