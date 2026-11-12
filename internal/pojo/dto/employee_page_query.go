package dto

// EmployeePageQueryDTO 员工分页查询的数据传输对象。
//
// 注意：分页查询通过URL查询参数传入（GET /admin/employee/page?page=1&name=xx），
// 练习时在Controller中用 r.URL.Query().Get(...) 手动解析。
type EmployeePageQueryDTO struct {
	Name     string
	Page     int
	PageSize int
}
