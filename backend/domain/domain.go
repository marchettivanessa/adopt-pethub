package domain

import (
	"adopt-pethub/backend/database"
	"errors"
	"fmt"
)

type Appointment struct {
	ID              int
	VisitorName     string
	AppointmentTime string
	CreatedAt       string
	Status          string
	PetID           int
}

type Pet struct {
	ID              int
	Name            string
	Age             string
	Addres          string
	Email           string
	PersonalDetails string
	PhoneNumber     string
}

type Domain struct{}

func (d Domain) GetConfirmedAppointments(appointmentDate string, db *database.Database) ([]Appointment, error) {
	queryResults := []Appointment{}

	query := `
SELECT * FROM appointments WHERE status = 'CONFIRMED' AND appointment_time <= $1`

	rows, err := db.Connection.Query(query, appointmentDate)
	if err != nil {
		return nil, fmt.Errorf("failed to select all appointments in this time range: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var appointment Appointment
		if err := rows.Scan(&appointment.ID, &appointment.VisitorName, &appointment.AppointmentTime, &appointment.CreatedAt, &appointment.Status, &appointment.PetID); err != nil {
			return nil, fmt.Errorf("failed to scan appointment: %w", err)
		}
		queryResults = append(queryResults, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	if len(queryResults) == 0 {
		return nil, nil
	}

	return queryResults, nil
}

func (d Domain) DeleteAppointmentById(appointmetnID int, db *database.Database) error {
	query := `DELETE * FROM appointments WHERE id = $1`

	_, err := db.Connection.Exec(query, appointmetnID)
	if err != nil {
		return fmt.Errorf("failed to delete the appointment %w", err)
	}
	return nil
}

func (d Domain) ValidateAppointment(appointment Appointment, db *database.Database) (*Appointment, error) {
	patientName := appointment.VisitorName
	if patientName == "" {
		return nil, errors.New("not possible to save an appointment without the patient name")
	}

	appointmentTime := appointment.AppointmentTime
	if appointmentTime == "" {
		return nil, errors.New("not possible to save an appointment without the appointment time")
	}

	createdAt := appointment.CreatedAt
	if createdAt == "" {
		return nil, errors.New("not possible to save an appointment without the created at")
	}

	status := appointment.Status
	patientID := appointment.PetID

	query := `
	INSERT INTO scheduler.appointments (patient_name, appointment_time, created_at, status, patient_id) VALUES ($1, $2< $3, $4, $5)
	`

	row := db.Connection.QueryRow(query, patientName, appointmentTime, createdAt, status, patientID)

	var insertedAppointment Appointment
	if err := row.Scan(&insertedAppointment.VisitorName, &insertedAppointment.AppointmentTime, &insertedAppointment.CreatedAt, &insertedAppointment.Status, &insertedAppointment.PetID); err != nil {
		return nil, errors.New("error inserting appointment into database")
	}

	return &insertedAppointment, nil
}
