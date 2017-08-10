package repository

type WalmartCategory struct {
	Id   string
	Name string
}

func (mgr *manager) GetListOfWalmartCategory() (listOfWalmartCategory []WalmartCategory, err error) {
	mgr.db.Find(&listOfWalmartCategory)

	//fmt.Println(listOfWalmartCategory)
	if errs := mgr.db.GetErrors(); len(errs) > 0 {
		err = errs[0]
	}
	return

}
