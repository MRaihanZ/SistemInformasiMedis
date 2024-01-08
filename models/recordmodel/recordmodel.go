package recordmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func Getall(isIdSet bool, id int) []entities.Record {
	if isIdSet {
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
	} else {
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
			ORDER BY record.id DESC;
		`)

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

}



func Create(record entities.Record) bool {
	result, err := config.DB.Exec(`
		INSERT INTO record(
			id, id_pasien, id_categori, id_doctor, diagnose, description, id_medicine, created_at, updated_at
		) VALUES ("", ?, ?, ?, ?, ?, ?, ?, ?)`,
		record.Pasien.Id,
		record.Category.Id,
		record.Doctor.Id,
		record.Diagnose,
		record.Description,
		record.Medicine.Id,
		record.CreatedAt,
		record.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Record {
	row := config.DB.QueryRow(`
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
			WHERE record.id = ?
			ORDER BY record.id DESC;
		`, id)

	var record entities.Record

	err := row.Scan(
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
	)

	if err != nil {
		panic(err)
	}

	return record
}

func Update(record entities.Record) bool {
	query, err := config.DB.Exec(`
		UPDATE record SET 
			id_pasien = ?, 
			id_categori = ?,
			id_doctor = ?,
			diagnose = ?,
			description = ?,
			id_medicine = ?,
			updated_at = ?
		WHERE id = ?`,
		record.Pasien.Id,
		record.Category.Id,
		record.Doctor.Id,
		record.Diagnose,
		record.Description,
		record.Medicine.Id,
		record.UpdatedAt,
		record.Id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM record WHERE id = ?", id)
	return err
}