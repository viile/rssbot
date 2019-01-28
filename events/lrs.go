package events

import (
	"errors"
	"github.com/viile/rssbot/utils"
)
// 角色
const (
	YuYanJia = iota
	NvWu
	ShouWei
	Baichi
	SheMengRen
	LieRen
	QiShi
	BaiLangWang
	HeiLangWang
	ELingQiShi
	LangRen
	CunMin
)
// 事件
const (
	Sha = iota
	Shou
	She
	Jiu
	Du
	Yan
)

var RoleTable  = map[int]string{
	YuYanJia : "预言家",
	NvWu : "女巫",
	ShouWei : "守卫",
	Baichi : "白痴",
	SheMengRen : "摄梦人",
	LieRen : "猎人",
	QiShi : "骑士",
	BaiLangWang : "白狼王",
	HeiLangWang : "黑狼王",
	ELingQiShi : "恶灵骑士",
	LangRen : "狼人",
	CunMin : "村民",
}

type Room struct{
	ID string
	Nums int
	Roles map[int]int	// 角色表 1:lr,2:cm
	Users map[int]int	// 用户表 1:12323,2:34234
	Status map[int]int // 用户状态表 1:0,2:1    0 live 1 dead
	Events map[int][]int // 事件表	kill:1,sh:1
	RoomStatus int // 1 created 2 started
	RoomOwner int // 房主
}

func NewRoom(id string,owner int,roles []int) *Room {
	roles = utils.Random(roles,6)
	len := len(roles)
	room := &Room{
		ID:id,
		Nums:len,
		Roles:make(map[int]int,len),
		Users:make(map[int]int,len),
		Status:make(map[int]int,len),
		Events:make(map[int][]int,0),
		RoomStatus:1,
		RoomOwner:owner,
	}

	for k,v := range roles {
		room.Roles[k] = v
		room.Status[k] = 0
		room.Users[k] = 0
	}

	return room
}

func (r *Room) Join(id int,index int) error {
	if index <= 0 {
		for k,v := range r.Users {
			if v == 0 {
				r.Users[k] = id
				return nil
			}
		}
		return errors.New("没有空位置了!")
	} else {
		_,ok := r.Users[index]
		if ok {
			return errors.New("该位置已经有人了!")
		} else {
			r.Users[index] = id
		}
	}

	return nil
}

func (r *Room) Start(id int) error {
	if r.RoomOwner != id {
		return errors.New("您不是房主!")
	}
	if r.RoomStatus != 1 {
		return errors.New("该房间已经开始!")
	}

	r.RoomStatus = 2;
	return nil
}

func (r *Room) AddEvent(id int,ops []int) error {
	r.Events[id] = ops
	return nil
}

func (r *Room) Result(id int) error {
	if r.RoomOwner != id {
		return errors.New("您不是房主!")
	}

	//for _,_ := range r.Events {
	//
	//}

	return nil
}