package doctormodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Doctor {
	rows, err := config.DB.Query(`SELECT * FROM doctor`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var doctors []entities.Doctor

	for rows.Next() {
		var doctor entities.Doctor
		if err := rows.Scan(&doctor.Id, &doctor.Name, &doctor.CreatedAt, &doctor.UpdatedAt); err != nil {
			panic(err)
		}

		doctors = append(doctors, doctor)
	}

	return doctors
}