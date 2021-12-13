package impl

import (
	"context"

	"github.com/ericyaoxr/mcube/exception"
	"github.com/ericyaoxr/mcube/sqlbuilder"

	"github.com/ericyaoxr/cmdb/app/bill"
)

const (
	insertBillSQL = `INSERT INTO bill (
		vendor,month,owner_id,owner_name,product_type,product_code,product_detail,
		pay_mode,pay_mode_detail,order_id,
		instance_id,instance_name,
		public_ip,private_ip,instance_config,region_code,region_name,sale_price,save_cost,real_cost,credit_pay,voucher_pay,cash_pay,storedcard_pay,outstanding_amount
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	queryBillSQL = `SELECT * FROM bill`
)

func (s *service) SaveBill(ctx context.Context, h *bill.Bill) (
	*bill.Bill, error) {
	if err := s.save(ctx, h); err != nil {
		return nil, err
	}
	return h, nil
}

func (s *service) QueryBill(ctx context.Context, req *bill.QueryBillRequest) (
	*bill.BillSet, error) {
	query := sqlbuilder.NewQuery(queryBillSQL)

	// if req.Keywords != "" {
	// 	query.Where("id = ? OR name LIKE ? OR Repo = ? OR Branch = ? OR Module LIKE ? OR Topic LIKE ? OR Job LIKE ?",
	// 		req.Keywords,
	// 		"%"+req.Keywords+"%",
	// 		req.Keywords,
	// 		req.Keywords,
	// 		req.Keywords+"%",
	// 		"%"+req.Keywords+"%",
	// 		req.Keywords,
	// 	)
	// }

	querySQL, args := query.Order("project_id").Desc().Limit(req.Page.Offset, uint(req.Page.PageSize)).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query bill error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	set := bill.NewBillSet()
	for rows.Next() {
		ins := bill.NewDefaultBill()
		app := ins
		err := rows.Scan(
			&app.Vendor, &app.Month, &app.OwnerId, &app.OwnerName, &app.ProductType, &app.ProductCode, &app.ProductDetail,
			&app.PayMode, &app.PayModeDetail, &app.OrderId, &app.InstanceId, &app.InstanceName,
			&app.PublicIp, &app.PrivateIp, &app.InstanceConfig, &app.RegionCode, &app.RegionName, &app.SalePrice, &app.SaveCost, &app.RealCost, &app.CreditPay, &app.VoucherPay, &app.CashPay, &app.StoredcardPay, &app.OutstandingAmount,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query bill error, %s", err.Error())
		}
		set.Add(ins)
	}

	// 获取total SELECT COUNT(*) FROMT t Where ....
	countSQL, args := query.BuildCount()
	countStmt, err := s.db.Prepare(countSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	defer countStmt.Close()
	err = countStmt.QueryRow(args...).Scan(&set.Total)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	return set, nil
}

// func (s *service) DescribeBill(ctx context.Context, req *bill.DescribeBillRequest) (
// 	*bill.Bill, error) {
// 	query := sqlbuilder.NewQuery(queryBillSQL)
// 	querySQL, args := query.Where("id = ?", req.Id).BuildQuery()
// 	s.log.Debugf("sql: %s", querySQL)

// 	queryStmt, err := s.db.Prepare(querySQL)
// 	if err != nil {
// 		return nil, exception.NewInternalServerError("prepare query bill error, %s", err.Error())
// 	}
// 	defer queryStmt.Close()

// 	ins := bill.NewDefaultBill()
// 	app := ins
// 	err = queryStmt.QueryRow(args...).Scan(
// 		&app.Vendor, &app.Month, &app.OwnerId, &app.OwnerName, &app.ProductType, &app.ProductCode, &app.ProductDetail,
// 		&app.PayMode, &app.PayModeDetail, &app.OrderId, &app.InstanceId, &app.InstanceName,
// 		&app.PublicIp, &app.PrivateIp, &app.InstanceConfig, &app.RegionCode, &app.RegionName, &app.SalePrice, &app.SaveCost, &app.RealCost, &app.CreditPay, &app.VoucherPay, &app.CashPay, &app.StoredcardPay, &app.OutstandingAmount,
// 	)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, exception.NewNotFound("%#v not found", req)
// 		}
// 		return nil, exception.NewInternalServerError("describe bill error, %s", err.Error())
// 	}

// 	return ins, nil
// }
