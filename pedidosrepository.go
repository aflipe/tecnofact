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

func getComidasById(id int64) (comidas, error) {
	var comida comidas
	bd, err := getDB()
	if err != nil {
		return comida, err
	}
	row := bd.QueryRow("SELECT * FROM tecnofac.comidas WHERE id = ?", id)
	err = row.Scan(&comida.Id, &comida.Descripcion, &comida.Comidastipoid)
	if err != nil {
		return comida, err
	}
	// Success!
	return comida, nil

}

func GetAllBebidas() ([]bebidas, error) {
	//Declare an array because if there's error, we return it empty
	bebida := []bebidas{}
	bd, err := getDB()
	if err != nil {
		return bebida, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT * FROM tecnofac.bebidas")
	if err != nil {
		return bebida, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var bebidas bebidas
		err = rows.Scan(&bebidas.Id, &bebidas.Descripcion, &bebidas.Bebidastipoid)
		if err != nil {
			return bebida, err
		}
		// and append it to the array
		bebida = append(bebida, bebidas)
	}
	return bebida, nil
}

func getBebidasById(id int64) (bebidas, error) {
	var bebida bebidas
	bd, err := getDB()
	if err != nil {
		return bebida, err
	}
	row := bd.QueryRow("SELECT * FROM tecnofac.bebidas WHERE id = ?", id)
	err = row.Scan(&bebida.Id, &bebida.Descripcion, &bebida.Bebidastipoid)
	if err != nil {
		return bebida, err
	}
	// Success!
	return bebida, nil

}

func GetAllBebidastipo() ([]bebidastipo, error) {
	//Declare an array because if there's error, we return it empty
	bebida := []bebidastipo{}
	bd, err := getDB()
	if err != nil {
		return bebida, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT * FROM tecnofac.bebidastipo")
	if err != nil {
		return bebida, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var bebidastipo bebidastipo
		err = rows.Scan(&bebidastipo.Id, &bebidastipo.Descripcion)
		if err != nil {
			return bebida, err
		}
		// and append it to the array
		bebida = append(bebida, bebidastipo)
	}
	return bebida, nil
}

func getBebidastipoById(id int64) (bebidastipo, error) {
	var bebida bebidastipo
	bd, err := getDB()
	if err != nil {
		return bebida, err
	}
	row := bd.QueryRow("SELECT * FROM tecnofac.bebidastipo WHERE id = ?", id)
	err = row.Scan(&bebida.Id, &bebida.Descripcion)
	if err != nil {
		return bebida, err
	}
	// Success!
	return bebida, nil

}

func GetAllRestaurantes() ([]restaurante, error) {
	//Declare an array because if there's error, we return it empty
	restaurantes := []restaurante{}
	bd, err := getDB()
	if err != nil {
		return restaurantes, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT * FROM tecnofac.restaurantes")
	if err != nil {
		return restaurantes, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var restaurante restaurante
		err = rows.Scan(&restaurante.Id, &restaurante.Descripcion)
		if err != nil {
			return restaurantes, err
		}
		// and append it to the array
		restaurantes = append(restaurantes, restaurante)
	}
	return restaurantes, nil
}

func getRestaurantesById(id int64) (restaurante, error) {
	var restaurantes restaurante
	bd, err := getDB()
	if err != nil {
		return restaurantes, err
	}
	row := bd.QueryRow("SELECT * FROM tecnofac.restaurantes WHERE id = ?", id)
	err = row.Scan(&restaurantes.Id, &restaurantes.Descripcion)
	if err != nil {
		return restaurantes, err
	}
	// Success!
	return restaurantes, nil

}

func createOrden(orden ordenar) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO tecnofac.orden (comida, bebida, observacion) VALUES (?, ?, ?)", orden.Comida, orden.Bebida, orden.Observacion)
	return err
}
