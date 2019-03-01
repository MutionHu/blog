package models

type PraiseLog struct {
	Model
	Key    string `sql:"index"`
	UserID int    `sql:"index"`
	Type   string `sql:"index"`
	Flag   bool
}

func (db *DB) UpdateNote4Praise(n *Note) error {
	return db.db.Model(&Note{}).Where("id = ?", n.ID).UpdateColumn("praise", n.Praise).Error
}

func (db *DB) QueryPraiseLog(key string, user_id int, ttype string) (parselog PraiseLog, err error) {
	return parselog, db.db.Where("`key` = ? and user_id =? and type = ? ", key, user_id, ttype).Take(&parselog).Error
}

func (db *DB) SavePraiseLog(p *PraiseLog) error {
	return db.db.Save(&p).Error
}
