package do

// MedalInfoTab 徽章表
type MedalInfoTab struct {
	Id          uint64 `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserDid     string `gorm:"column:user_did;type:varchar(64);comment:用户id;NOT NULL" json:"user_did"`
	EventUid    string `gorm:"column:event_uid;type:varchar(64);comment:活动uid;NOT NULL" json:"event_uid"`
	Teacher     string `gorm:"column:teacher;type:varchar(16);NOT NULL" json:"teacher"`
	MedalStatus int8   `gorm:"column:medal_status;type:tinyint(3) unsigned;default:0;comment:徽章状态;NOT NULL" json:"medal_status"`
	ChainHash   string `gorm:"column:chain_hash;type:varchar(128);NOT NULL" json:"chain_hash"`
	ChainStatus int8   `gorm:"column:chain_status;type:tinyint(3) unsigned;default:0;comment:上链状态 0:未上链 1:队列中 2:已上链 3:上链失败;NOT NULL" json:"chain_status"`
	Ctime       int32  `gorm:"column:ctime;type:int(10);default:0;NOT NULL" json:"ctime"`
	Mtime       int32  `gorm:"column:mtime;type:int(10);default:0;NOT NULL" json:"mtime"`
}

func (m *MedalInfoTab) TableName() string {
	return "medal_info_tab"
}
