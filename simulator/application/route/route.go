package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type ParcialRouteosition struct {
	ID       string    `json:"routeId"` //Quando for alocado em json a variável terá esse nome
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) LoadPositions() error {
	if r.ID == "" { //Se o ID estiver vazio retorna erro
		return errors.New("route id not informed")
	}
	f, err := os.Open("destinations/" + r.ID + ".txt") //Caso encontre o ID abre o arquivo referente a id
	if err != nil {                                    //Caso a variável erro não esteja vazia, retorna o erro
		return err
	}
	defer f.Close() //Fecha o arquivo

	scanner := bufio.NewScanner(f) //Scan no arq

	for scanner.Scan() { //Percorre o arq
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64) //Realiza o casting
		if err != nil {                             //Caso a variável erro não esteja vazia, retorna o erro
			return nil
		}
		lon, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}

		r.Positions = append(r.Positions, Position{ //Realiza append dos valores do txt para a struct
			Lat: lat,
			Lon: lon,
		})
	}
	return nil
}

func (r *Route) ExportJsonPositions() ([]string, error) {
	var route ParcialRouteosition
	var result []string
	total := len(r.Positions)

	for k, value := range r.Positions { //For para alocar valores nas structs
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{value.Lat, value.Lon}
		route.Finished = false
		if total-1 == k {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}

	return result, nil
}
