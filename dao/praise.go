package dao

import "Project/model"

// TakePraiseDate 将praise表中的数据取出
func TakePraiseDate() model.Praises {
	rows, _ := dB.Query("select * from ?", "user")
	var g model.Praise
	var G model.Praises
	for rows.Next() {
		err := rows.Scan(&g.Author, &g.Reader, &g.Good)
		if err != nil {
			panic(err)
		}
		G = append(G, g)
	}
	return G
}

// TakeGoodMemberDate 将good_member表中的数据取出
func TakeGoodMemberDate() model.Goods {
	rows, _ := dB.Query("select * from ?", "user")
	var o model.Good
	var O model.Goods
	for rows.Next() {
		err := rows.Scan(&o.Author, &o.Member)
		if err != nil {
			panic(err)
		}
		O = append(O, o)
	}
	return O
}

// BringPraiseDate 将数据放入praise表中
func BringPraiseDate(g model.Praise) {
	_, err := dB.Exec("insert into praise (author,reader,good) value (?,?,?)",
		g.Author, g.Reader, g.Good)
	if err != nil {
		return
	}
}

// BringGoodMemberDate 将数据放入good_member表中
func BringGoodMemberDate(o model.Good) {
	_, err := dB.Exec("insert into `good-member` (author,member) value (?,?)",
		o.Author, o.Member)
	if err != nil {
		return
	}
}
