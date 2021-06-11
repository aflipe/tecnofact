package main

func GetAllComidas() ([]comidas, error) {
	//Declare an array because if there's error, we return it empty
	comida := []comidas{}
	bd, err := getDB()
	if err != nil {
		return comida, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT * FROM tecnofac.comidas")
	if err != nil {
		return comida, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var comidas comidas
		err = rows.Scan(&comidas.Id, &comidas.Descripcion, &comidas.Comidastipoid)
		if err != nil {
			return comida, err
		}
		// and append it to the array
		comida = append(comida, comidas)
	}
	return comida, nil
}
