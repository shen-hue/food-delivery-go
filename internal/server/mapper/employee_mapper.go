// Package mapper 对应 Java 版 sky-server 的 com.sky.mapper 包。
//
// Java 版使用 MyBatis：接口方法 + 注解SQL/XML文件；
// Go 版使用 pgx 直接编写SQL，配合结构体 db 标签自动扫描，
// 相当于 MyBatis 的 map-underscore-to-camel-case 效果。
//
// 约定：查询不到数据时返回 (nil, nil)，与 MyBatis 返回 null 的行为一致，
// 由 service 层判断并抛出业务异常。
package mapper

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"sky-take-out/internal/pojo/entity"
)

// EmployeeMapper 员工数据访问接口（对应 Java 版 EmployeeMapper）。
type EmployeeMapper struct {
	pool *pgxpool.Pool
}

// NewEmployeeMapper 构造函数（对应 @Mapper 注解的注入）。
func NewEmployeeMapper(pool *pgxpool.Pool) *EmployeeMapper {
	return &EmployeeMapper{pool: pool}
}

// GetByUsername 根据用户名查询员工。
//
// 对应 Java 版：
//
//	@Select("select * from employee where username = #{username}")
//	Employee getByUsername(String username);
func (m *EmployeeMapper) GetByUsername(ctx context.Context, username string) (*entity.Employee, error) {
	rows, err := m.pool.Query(ctx, "select * from employee where username = $1", username)
	if err != nil {
		return nil, err
	}

	employee, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Employee])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // 查无此记录，对应 MyBatis 返回 null
		}
		return nil, err
	}
	return &employee, nil
}

// TODO: 后续练习中你将在这里补充其它方法，例如（参考 Java 完成版）：
//   - Insert(ctx, employee)                // 新增员工
//   - PageQuery(ctx, dto)                  // 分页查询（SQL中用 limit/offset 实现分页）
//   - GetByID(ctx, id)                     // 根据id查询员工
//   - Update(ctx, employee)                // 编辑员工/启用禁用/修改密码
