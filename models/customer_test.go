package models_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xerardoo/sql-testing-example/models"
	"gorm.io/gorm"
	"regexp"
	"syreclabs.com/go/faker"
	"testing"
)

func TestCustomer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Customer Suite")
}

var _ = Describe("Customers", func() {

	var db *sql.DB
	var mock sqlmock.Sqlmock
	var customer models.Customer
	const rowAffected, idDigits = 1, 2

	BeforeEach(func() {
		var err error

		db, mock, err = sqlmock.New() // mock sql.DB
		Expect(err).ShouldNot(HaveOccurred())

		err = models.InitMockDB(db) // mock gorm.db
		Expect(err).ShouldNot(HaveOccurred())

		customer = models.Customer{
			Model:     gorm.Model{ID: uint(faker.Number().NumberInt32(idDigits))},
			FirstName: faker.Name().FirstName(),
			LastName:  faker.Name().LastName(),
			Phone:     faker.PhoneNumber().String(),
			Address:   faker.Address().String(),
			City:      faker.Address().City(),
			State:     faker.Address().State(),
			ZipCode:   faker.Address().ZipCode(),
		}
	})

	AfterEach(func() {
		defer db.Close()
		err := mock.ExpectationsWereMet() // make sure all expectations were met
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("Get All", func() {

		It("Empty", func() {
			query := "SELECT * FROM `customers`"
			rows := sqlmock.NewRows(nil)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			c, err := models.GetCustomers()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(c).Should(BeEmpty())
		})

		It("Exist", func() {
			query := "SELECT * FROM `customers`"
			newId := uint(faker.Number().NumberInt32(4))
			rows := sqlmock.
				NewRows([]string{"id", "first_name", "last_name", "phone", "address", "city", "state", "zip_code"}).
				AddRow(customer.ID, customer.FirstName, customer.LastName, customer.Phone, customer.Address, customer.City, customer.State, customer.ZipCode).
				AddRow(newId, customer.FirstName, customer.LastName, customer.Phone, customer.Address, customer.City, customer.State, customer.ZipCode)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			c, err := models.GetCustomers()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(c[0].ID).Should(Equal(customer.ID))
			Expect(c[0].FirstName).Should(Equal(customer.FirstName))
			Expect(c[1].ID).Should(Equal(newId))
		})
	})

	Context("Find One", func() {

		It("Empty", func() {
			query := "SELECT * FROM `customers` WHERE `customers`.`id` = ?"
			rows := sqlmock.NewRows(nil)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			_, err := models.FindCustomer(customer.ID)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})

		It("Exist", func() {
			query := "SELECT * FROM `customers` WHERE `customers`.`id` = ?"
			rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "phone", "address", "city", "state", "zip_code"}).
				AddRow(customer.ID, customer.FirstName, customer.LastName, customer.Phone, customer.Address, customer.City, customer.State, customer.ZipCode)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			c, err := models.FindCustomer(customer.ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(c.ID).Should(Equal(customer.ID))
			Expect(c.FirstName).Should(Equal(customer.FirstName))
		})
	})

	Context("Save", func() {
		It("Add", func() {
			query := "INSERT INTO `customers` (`created_at`,`updated_at`,`deleted_at`,`first_name`,`last_name`,`phone`,`address`,`city`,`state`,`zip_code`,`id`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					customer.FirstName,
					customer.LastName,
					customer.Phone,
					customer.Address,
					customer.City,
					customer.State,
					customer.ZipCode,
					customer.ID,
				).
				WillReturnResult(sqlmock.NewResult(int64(customer.ID), rowAffected))
			mock.ExpectCommit()

			err := customer.AddCustomer()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Update", func() {
			query := "UPDATE `customers` SET `created_at`=?,`updated_at`=?,`deleted_at`=?,`first_name`=?,`last_name`=?,`phone`=?,`address`=?,`city`=?,`state`=?,`zip_code`=? WHERE `id` = ?"
			customer.ID = 1

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					customer.FirstName,
					customer.LastName,
					customer.Phone,
					customer.Address,
					customer.City,
					customer.State,
					customer.ZipCode,
					customer.ID,
				).WillReturnResult(sqlmock.NewResult(int64(customer.ID), rowAffected))
			mock.ExpectCommit()

			err := customer.UpdateCustomer()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Delete", func() {
		It("Soft-Delete", func() {
			query := " UPDATE `customers` SET `deleted_at`=? WHERE `customers`.`id` = ? AND `customers`.`deleted_at` IS NULL"

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(
					sqlmock.AnyArg(),
					customer.ID,
				).
				WillReturnResult(sqlmock.NewResult(0, rowAffected))
			mock.ExpectCommit()

			err := customer.DeleteCustomer()
			Expect(err).ShouldNot(HaveOccurred())

		})
	})
})
