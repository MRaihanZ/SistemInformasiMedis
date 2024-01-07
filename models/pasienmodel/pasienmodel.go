package pasienmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Pasien {
	rows, err := config.DB.Query(`SELECT * FROM pasien`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var pasiens []entities.Pasien

	for rows.Next() {
		var pasien entities.Pasien
		if err := rows.Scan(&pasien.Id, &pasien.Name, &pasien.CreatedAt, &pasien.UpdatedAt); err != nil {
			panic(err)
		}

		pasiens = append(pasiens, pasien)
	}

	return pasiens
}