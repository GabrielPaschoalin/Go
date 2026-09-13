package models

import (
	"gabriel/api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events []Event = []Event{}

// Save insere um novo evento no banco, associado ao usuário informado em UserID.
func (e *Event) Save() error {
	query := `
	INSERT INTO events 
	(name, description, location, dateTime, user_id)
	VALUES
	(?, ?, ?, ?, ?)
	`
	// Guarda a query na memoria e reusa com uma otimizacao melhor
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	// Executa quando terminar de rodar o metodo
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

// GetAllEvents retorna todos os eventos cadastrados no banco.
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	// Usado quando quer pegar um conjunto de linhas
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	// Converter cada linha retornada em um Event
	for rows.Next() {
		var event Event

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

// GetEventByID busca um único evento pelo seu ID.
func GetEventByID(id int64) (*Event, error) {
	query := `
		SELECT *
		FROM events
		WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, err
}

// Update sobrescreve os dados do evento (identificado por event.ID) no banco.
func (event Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ? , location = ?, dateTime = ?
		WHERE id = ?
	`
	// Preparar a query
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	// Executar a query com os novos valores
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

// DeleteEvent remove o evento (identificado por event.ID) do banco.
func (event Event) DeleteEvent() error {
	query := `
		DELETE FROM events
		where id = ?
	`

	// Preparar a query
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	// Executar a query
	_, err = stmt.Exec(event.ID)

	return err
}

// Register cria uma inscrição do usuário informado no evento.
func (e Event) Register(userId int64) error {

	query := `INSERT INTO registrations
			(event_id, user_id)
			VALUES
			(?, ?)
			`

	// Preparar a query para ser usada em segui3da,
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Executar a query
	_, err = stmt.Exec(e.ID, userId)

	return err
}

// DeleteRegistration remove a inscrição do usuário informado no evento.
func (e Event) DeleteRegistration(userId int64) error {

	query := `DELETE FROM registrations
		WHERE event_id = ? and user_id = ?`

	// Preparar a query para ser usada em segui3da,
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Executar a query
	_, err = stmt.Exec(e.ID, userId)

	return err

}
