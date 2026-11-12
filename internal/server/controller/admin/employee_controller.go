// Package admin 对应 Java 版 sky-server 的 com.sky.controller.admin 包。
package admin

import (
	"log/slog"
	"net/http"

	"sky-take-out/internal/common/constant"
	"sky-take-out/internal/common/jsonutil"
	"sky-take-out/internal/common/properties"
	"sky-take-out/internal/common/result"
	"sky-take-out/internal/common/utils"
	"sky-take-out/internal/pojo/dto"
	"sky-take-out/internal/pojo/vo"
	"sky-take-out/internal/server/service"
)

// EmployeeController 员工管理
// （对应 Java 版 @RestController + @RequestMapping("/admin/employee")）。
type EmployeeController struct {
	employeeService service.EmployeeService
	jwtProperties   *properties.JwtProperties
}

// NewEmployeeController 构造函数（对应 @Autowired 依赖注入）。
func NewEmployeeController(employeeService service.EmployeeService,
	jwtProperties *properties.JwtProperties) *EmployeeController {
	return &EmployeeController{
		employeeService: employeeService,
		jwtProperties:   jwtProperties,
	}
}

// Login 登录（对应 @PostMapping("/login")）。
func (c *EmployeeController) Login(w http.ResponseWriter, r *http.Request) error {
	var employeeLoginDTO dto.EmployeeLoginDTO
	if err := jsonutil.ReadJSON(r, &employeeLoginDTO); err != nil {
		return err
	}
	slog.Info("员工登录", "username", employeeLoginDTO.Username)

	employee, err := c.employeeService.Login(r.Context(), &employeeLoginDTO)
	if err != nil {
		return err
	}

	// 登录成功后，生成jwt令牌
	claims := map[string]any{
		constant.JwtClaimEmpID: employee.ID,
	}
	token, err := utils.CreateJWT(
		c.jwtProperties.AdminSecretKey,
		c.jwtProperties.AdminTTL,
		claims)
	if err != nil {
		return err
	}

	employeeLoginVO := vo.EmployeeLoginVO{
		ID:       employee.ID,
		UserName: employee.Username,
		Name:     employee.Name,
		Token:    token,
	}

	return jsonutil.WriteJSON(w, http.StatusOK, result.Success(employeeLoginVO))
}

// Logout 退出（对应 @PostMapping("/logout")）。
func (c *EmployeeController) Logout(w http.ResponseWriter, r *http.Request) error {
	return jsonutil.WriteJSON(w, http.StatusOK, result.SuccessEmpty())
}

// TODO: 后续练习中你将在这里补充其它接口（参考 Java 完成版 food-delivery-j）：
//   - POST   /admin/employee          新增员工
//   - GET    /admin/employee/page     员工分页查询
//   - POST   /admin/employee/status/{status}  启用禁用员工
//   - GET    /admin/employee/{id}     根据id查询员工
//   - PUT    /admin/employee          编辑员工信息
