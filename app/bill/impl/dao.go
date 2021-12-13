package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ericyaoxr/cmdb/app/bill"
)

func (s *service) save(ctx context.Context, ins *bill.Bill) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 开启一个事务
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 执行结果提交或者回滚事务
	// 当使用sql.Tx的操作方式操作数据后，需要我们使用sql.Tx的Commit()方法显式地提交事务，
	// 如果出错，则可以使用sql.Tx中的Rollback()方法回滚事务，保持数据的一致性
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	// 避免SQL注入, 请使用Prepare
	stmt, err = tx.Prepare(insertBillSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()
	app := ins
	_, err = stmt.Exec(
		&app.Vendor, &app.Month, &app.OwnerId, &app.OwnerName, &app.ProductType, &app.ProductCode, &app.ProductDetail,
		&app.PayMode, &app.PayModeDetail, &app.OrderId, &app.InstanceId, &app.InstanceName,
		&app.PublicIp, &app.PrivateIp, &app.InstanceConfig, &app.RegionCode, &app.RegionName, &app.SalePrice, &app.SaveCost, &app.RealCost, &app.CreditPay, &app.VoucherPay, &app.CashPay, &app.StoredcardPay, &app.OutstandingAmount,
	)
	if err != nil {
		return fmt.Errorf("save bill resource info error, %s", err)
	}

	return tx.Commit()
}
