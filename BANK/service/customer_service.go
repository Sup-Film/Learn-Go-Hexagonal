package service

import "bank/repository"

/*
ตัว CustomerService จะเป็นตัวกลางระหว่าง handler กับ repository แปลว่าตัวมันเองไม่มี data อยู่ในมือ ตัวมันเองทำหน้าที่แค่เป็นตัวจัดการคอยรับและส่งข้อมูลระหว่าง handler กับ repository เท่านั้น
เพราะฉะนั้นตัวมันเองเลยจะต้องมีแหล่งเก็บข้อมูลเพื่อใช้ในการดึง data ขึ้นมา ซึ่งแหล่งเก็บข้อมูลนี้ก็คือ repository นั่นเอง เลยต้องมีการทำ dependency injection โดยการส่ง repository เข้ามาใน constructor function ของ service ด้วย

แปลว่าตัว Service นี้จะไม่รู้จักกับ database โดยตรง มันจะรู้จักแค่ repository ที่เป็น interface เท่านั้น เราจะเปลี่ยน database เป็นอะไรก็ได้โดยไม่กระทบกับ service

และเราจะไม่ได้ทำการ export customerService ออกไปโดยตรง แต่จะให้ส่วนอื่นมาเรียกใช้ผ่านตัว interface CustomerService แทน เพื่อเป็นการซ่อน implementation details ของ service ไว้ โดยจะสร้างผ่านฟังก์ชัน NewCustomerService เท่านั้น
*/
type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}

func (s *customerService) GetCustomers() ([]CustomerResponse, error) {
	return nil, nil
}

func (s *customerService) GetCustomer(id int) (*CustomerResponse, error) {
	return nil, nil
}
