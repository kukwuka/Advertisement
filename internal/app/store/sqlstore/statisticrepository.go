package sqlstore

import (
	"Advertising/internal/app/model"
	"errors"
	"github.com/sirupsen/logrus"
	"reflect"
	"time"
)

type AdvertisementRepository struct {
	store *Store
}

//Create ...
func (r *AdvertisementRepository) Create(a *model.Advertisement) (int, error) {
	var id int
	if err := a.Validate(); err != nil {
		return 0, err
	}
	row := r.store.db.QueryRow(
		"INSERT INTO advertisement (date,name, description, cost, img_url) VALUES ($1,$2,$3,$4,$5) Returning id",
		time.Now(),
		a.Name,
		a.Description,
		a.Cost,
		a.ImgUrl)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

//GetAll
func (r *AdvertisementRepository) GetAll(Page int) ([]*model.Advertisement, error) {

	rows, err := r.store.db.Query(`
		SELECT id, date, name, description, cost, img_url 
		FROM advertisement
		ORDER BY cost, date
		OFFSET $1 LIMIT 10`, Page*10)
	if err != nil {
		return nil, err
	}
	us := make([]*model.Advertisement, 0, 10)
	for rows.Next() {
		u := new(model.Advertisement)
		err := rows.Scan(&u.Id, &u.Date, &u.Name, &u.Description, &u.Cost, &u.ImgUrl)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	logrus.Print(reflect.DeepEqual([]*model.Advertisement{},us))
	logrus.Print("Check")
	return us, nil
}

//GetAllMap
func (r *AdvertisementRepository) GetAllMap(Page int) ([]map[string]interface{}, error) {

	rows, err := r.store.db.Query(`
		SELECT id, name, cost, img_url 
		FROM advertisement
		ORDER BY cost, date
		OFFSET $1 LIMIT 10`, Page*10)
	if err != nil {
		return nil, err
	}

	us := make([]map[string]interface{}, 0, 10)
	for rows.Next() {
		model := new(model.Advertisement)
		err := rows.Scan(
			&model.Id,
			&model.Name,
			&model.Cost,
			&model.ImgUrl)
		if err != nil {
			return nil, err
		}
		var imgUrl string
		for in , _ := range model.ImgUrl{
			imgUrl = model.ImgUrl[in]
		}
		u := map[string]interface{}{
			"id":   model.Id,
			"name": model.Name,
			"cost": model.Cost,
			"img":  imgUrl,
		}
		us = append(us, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	if len(us)==0{
		return nil, errors.New("page out of range")
	}
	logrus.Print("Check")
	return us, nil
}

//GetOne
func (r *AdvertisementRepository) GetOne(Id int) (*model.Advertisement, error) {
	rows := r.store.db.QueryRow(`
		SELECT id, date, name, description, cost, img_url 
		FROM advertisement
		WHERE advertisement.id = $1`, Id)
	u := new(model.Advertisement)
	err := rows.Scan(&u.Id, &u.Date, &u.Name, &u.Description, &u.Cost, &u.ImgUrl)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *AdvertisementRepository) Delete(Id int) error {
	_, err := r.store.db.Exec(`
		DELETE FROM advertisement
		WHERE advertisement.id = $1`, Id)
	if err != nil {
		return err
	}
	return nil
}
