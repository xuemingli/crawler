package model

import "encoding/json"

type Profile struct {
	Name       string //姓名（昵称）
	Age        int    //年龄
	Gender     string //性别
	Height     int    //身高
	Weight     int    //体重kg
	Income     string //收入
	Marriage   string //婚姻状况
	Education  string //教育
	Birthplace string //籍贯
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
