package medicinemodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Medicine {
	rows, err := config.DB.Query(`SELECT * FROM medicine`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var medicines []entities.Medicine

	for rows.Next() {
		var medicine entities.Medicine
		if err := rows.Scan(&medicine.Id, &medicine.Name, &medicine.CreatedAt, &medicine.UpdatedAt); err != nil {
			panic(err)
		}

		medicines = append(medicines, medicine)
	}

	return medicines
}