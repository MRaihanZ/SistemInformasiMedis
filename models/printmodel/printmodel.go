package printmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll(id int) []entities.Record {
	rows, err := config.DB.Query(`
			SELECT 
				record.id,
				pasien.id,
				pasien.name,
				categories.id,
				categories.name,
				doctor.id,
				doctor.name,
				record.diagnose,
				record.description,
				medicine.id,
				medicine.name,
				record.created_at,
				record.updated_at 
			FROM record
			JOIN pasien ON record.id_pasien = pasien.id
			JOIN categories ON record.id_categori = categories.id
			JOIN doctor ON record.id_doctor = doctor.id
			JOIN medicine ON record.id_medicine = medicine.id
			WHERE pasien.id = ?
			ORDER BY pasien.id DESC;
		`, id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var records []entities.Record

	for rows.Next() {
		var record entities.Record
		if err := rows.Scan(
			&record.Id,
			&record.Pasien.Id,
			&record.Pasien.Name,
			&record.Category.Id,
			&record.Category.Name,
			&record.Doctor.Id,
			&record.Doctor.Name,
			&record.Diagnose,
			&record.Description,
			&record.Medicine.Id,
			&record.Medicine.Name,
			&record.CreatedAt,
			&record.UpdatedAt,
		); err != nil {
			panic(err)
		}
		records = append(records, record)
	}
	return records
}