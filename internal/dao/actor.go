package dao

import "TicketSales/internal/model"

func (d *Dao) ActorList(name string, offset, size int) ([]*model.Actor, error) {
	actor := model.Actor{Name1: name, Name2: name}
	return actor.SearchList(d.engine, offset, size)
}

func (d *Dao) GetActor(id int) (model.Actor, error) {
	return model.Actor{ActorID: id}.Get(d.engine)
}

func (d *Dao) UpdateActor(id int, n1, n2, intro, ava string, b int64) error {
	return model.Actor{ActorID: id, Name1: n1, Name2: n2, Introduction: intro, Avatar: ava, Birthday: b}.Update(d.engine)
}

func (d *Dao) DelActor(id int) error {
	return model.Actor{ActorID: id}.Delete(d.engine)
}

func (d *Dao) CreateActor(n1, n2, intro, ava string, b int64) error {
	return model.Actor{Name1: n1, Name2: n2, Introduction: intro, Avatar: ava, Birthday: b}.Create(d.engine)
}
