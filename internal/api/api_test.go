package api_test

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/ozonva/ova-food-api/internal/repo"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/jmoiron/sqlx"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ozonva/ova-food-api/internal/api"
	"github.com/ozonva/ova-food-api/internal/food"
	desc "github.com/ozonva/ova-food-api/pkg/ova-food-api"
)

var _ = Describe("Api", func() {
	var (
		ctx       context.Context
		coffee    = food.Food{Id: 0, UserId: 0, Type: food.Drinks, Name: "Coffee", PortionSize: 60}
		pizza     = food.Food{Id: 1, UserId: 0, Type: food.Foods, Name: "Pizza", PortionSize: 300}
		err       error
		db        *sql.DB
		sqlxDB    *sqlx.DB
		mock      sqlmock.Sqlmock
		repoTest  repo.Repo
		apiTest   desc.OvaFoodApiServer
		descrResp *desc.DescribeFoodV1Response
		listResp  *desc.ListFoodsV1Response
	)

	BeforeEach(func() {
		ctx = context.Background()
		db, mock, err = sqlmock.New()
		sqlxDB = sqlx.NewDb(db, "sqlmock")
		repoTest = repo.NewRepo(sqlxDB)
	})
	JustBeforeEach(func() {
		apiTest = api.NewFoodAPI(repoTest)
	})
	AfterEach(func() {
		mock.ExpectClose()
		db.Close()
	})

	Context("add coffee", func() {
		BeforeEach(func() {
		})
		It("add coffee", func() {
			req := &desc.CreateFoodV1Request{
				Food: &desc.CreationFood{
					UserId:      coffee.UserId,
					FoodT:       desc.FoodType(coffee.Type),
					Name:        coffee.Name,
					PortionSize: coffee.PortionSize,
				},
			}
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO food_info (user_id,type,name,portion_size) VALUES ($1,$2,$3,$4)")).
				WithArgs(coffee.UserId, coffee.Type, coffee.Name, coffee.PortionSize).
				WillReturnResult(sqlmock.NewResult(0, 1))
			func() {
				_, err = apiTest.CreateFoodV1(ctx, req)
			}()
			gomega.Expect(err).Should(gomega.BeNil())
		})
		It("add  - internal error", func() {
			req := &desc.CreateFoodV1Request{
				Food: &desc.CreationFood{
					UserId:      coffee.UserId,
					FoodT:       desc.FoodType(coffee.Type),
					Name:        coffee.Name,
					PortionSize: coffee.PortionSize,
				},
			}
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO food_info (user_id,type,name,portion_size) VALUES ($1,$2,$3,$4)")).
				WithArgs(coffee.UserId, coffee.Type, coffee.Name, coffee.PortionSize).
				WillReturnError(sqlmock.ErrCancelled)
			func() {
				_, err = apiTest.CreateFoodV1(ctx, req)
			}()
			gomega.Expect(err).ShouldNot(gomega.BeNil())
		})

	})
	Context("describe food", func() {
		BeforeEach(func() {
		})
		It("describe food", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, user_id, type, name, portion_size FROM food_info WHERE id = $1")).
				WithArgs(coffee.Id).
				WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "type", "name", "portion_size"}).
					AddRow(coffee.Id, coffee.UserId, coffee.Type, coffee.Name, coffee.PortionSize))
			func() {
				descrResp, err = apiTest.DescribeFoodV1(ctx, &desc.DescribeFoodV1Request{
					FoodId: coffee.Id,
				})
			}()
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(descrResp).ShouldNot(gomega.BeNil())
		})
		It("describe food with wrong id", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, user_id, type, name, portion_size FROM food_info WHERE id = $1")).
				WithArgs(8).
				WillReturnError(repo.HaveNotElementErr)
			func() {
				descrResp, err = apiTest.DescribeFoodV1(ctx, &desc.DescribeFoodV1Request{
					FoodId: 8,
				})
			}()
			gomega.Expect(err).ShouldNot(gomega.BeNil())
			gomega.Expect(descrResp).Should(gomega.BeNil())
		})
		It("describe food with internal error", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, user_id, type, name, portion_size FROM food_info WHERE id = $1")).
				WithArgs(coffee.Id).
				WillReturnError(sqlmock.ErrCancelled)
			func() {
				descrResp, err = apiTest.DescribeFoodV1(ctx, &desc.DescribeFoodV1Request{
					FoodId: coffee.Id,
				})
			}()
			gomega.Expect(err).ShouldNot(gomega.BeNil())
			gomega.Expect(descrResp).Should(gomega.BeNil())
		})
	})
	Context("list foods", func() {
		BeforeEach(func() {
		})
		It("list foods", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, user_id, type, name, portion_size FROM food_info WHERE id = $1")).
				WithArgs(coffee.Id).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "user_id", "type", "name", "portion_size"}).
						AddRow(coffee.Id, coffee.UserId, coffee.Type, coffee.Name, coffee.PortionSize))
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, user_id, type, name, portion_size FROM food_info WHERE id = $1")).
				WithArgs(pizza.Id).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "user_id", "type", "name", "portion_size"}).
						AddRow(pizza.Id, pizza.UserId, pizza.Type, pizza.Name, pizza.PortionSize))
			func() {
				listResp, err = apiTest.ListFoodsV1(ctx, &desc.ListFoodsV1Request{
					Ids: []uint64{coffee.Id, pizza.Id},
				})
			}()
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(listResp).ShouldNot(gomega.BeNil())
		})
		It("list foods with wrong id", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, user_id, type, name, portion_size FROM food_info WHERE id = $1")).
				WithArgs(coffee.Id).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "user_id", "type", "name", "portion_size"}).
						AddRow(coffee.Id, coffee.UserId, coffee.Type, coffee.Name, coffee.PortionSize))
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, user_id, type, name, portion_size FROM food_info WHERE id = $1")).
				WithArgs(8).
				WillReturnError(repo.HaveNotElementErr)
			func() {
				listResp, err = apiTest.ListFoodsV1(ctx, &desc.ListFoodsV1Request{
					Ids: []uint64{coffee.Id, 8},
				})
			}()
			gomega.Expect(err).ShouldNot(gomega.BeNil())
			gomega.Expect(listResp).Should(gomega.BeNil())
		})
	})
	Context("remove food", func() {
		BeforeEach(func() {
		})
		It("remove food", func() {
			mock.ExpectExec(regexp.QuoteMeta("DELETE FROM food_info WHERE id = $1")).
				WithArgs(coffee.Id).WillReturnResult(sqlmock.NewResult(0, 1))
			func() {
				_, err = apiTest.RemoveFoodV1(ctx, &desc.RemoveFoodV1Request{
					FoodId: coffee.Id,
				})
			}()
			gomega.Expect(err).Should(gomega.BeNil())
		})

		It("internal error", func() {
			mock.ExpectExec(regexp.QuoteMeta("DELETE FROM food_info WHERE id = $1")).
				WithArgs(coffee.Id).WillReturnError(sqlmock.ErrCancelled)
			func() {
				_, err = apiTest.RemoveFoodV1(ctx, &desc.RemoveFoodV1Request{
					FoodId: coffee.Id,
				})
			}()
			gomega.Expect(err).ShouldNot(gomega.BeNil())
		})
	})
})
