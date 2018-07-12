package sqldb_test

import (
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/bbs/models/test/model_helpers"
	"code.cloudfoundry.org/bbs/test_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Suspect ActualLRPs", func() {
	var (
		actualLRP *models.ActualLRP
		guid      string
		index     int32
	)

	BeforeEach(func() {
		guid = "some-guid"
		index = int32(1)
		actualLRP = model_helpers.NewValidActualLRP(guid, index)
		actualLRP.CrashCount = 0
		actualLRP.CrashReason = ""
		actualLRP.Since = fakeClock.Now().UnixNano()
		actualLRP.ModificationTag = models.ModificationTag{}
		actualLRP.ModificationTag.Increment()
		actualLRP.ModificationTag.Increment()

		_, err := sqlDB.CreateUnclaimedActualLRP(logger, &actualLRP.ActualLRPKey)
		Expect(err).NotTo(HaveOccurred())
		_, _, err = sqlDB.ClaimActualLRP(logger, guid, index, &actualLRP.ActualLRPInstanceKey)
		Expect(err).NotTo(HaveOccurred())
		_, _, err = sqlDB.StartActualLRP(logger, &actualLRP.ActualLRPKey, &actualLRP.ActualLRPInstanceKey, &actualLRP.ActualLRPNetInfo)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("RemoveSuspectActualLRP", func() {
		Context("when there is an suspect actualLRP", func() {
			BeforeEach(func() {
				queryStr := "UPDATE actual_lrps SET presence = ? WHERE process_guid = ? AND instance_index = ? AND presence = ?"
				if test_helpers.UsePostgres() {
					queryStr = test_helpers.ReplaceQuestionMarks(queryStr)
				}
				_, err := db.Exec(queryStr, models.ActualLRP_Suspect, actualLRP.ProcessGuid, actualLRP.Index, models.ActualLRP_Ordinary)
				Expect(err).NotTo(HaveOccurred())
			})

			It("removes the suspect actual LRP", func() {
				err := sqlDB.RemoveSuspectActualLRP(logger, &actualLRP.ActualLRPKey, &actualLRP.ActualLRPInstanceKey)
				Expect(err).ToNot(HaveOccurred())

				_, err = sqlDB.ActualLRPGroupByProcessGuidAndIndex(logger, guid, index)
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(models.ErrResourceNotFound))
			})

			Context("when the actual lrp instance key is not the same", func() {
				BeforeEach(func() {
					actualLRP.CellId = "a different cell"
				})

				It("returns a ErrActualLRPCannotBeRemoved error", func() {
					err := sqlDB.RemoveSuspectActualLRP(logger, &actualLRP.ActualLRPKey, &actualLRP.ActualLRPInstanceKey)
					Expect(err).To(Equal(models.ErrActualLRPCannotBeRemoved))
				})
			})

			Context("when the actualLRP is expired", func() {
				It("does not return an error", func() {
					err := sqlDB.RemoveSuspectActualLRP(logger, &actualLRP.ActualLRPKey, &actualLRP.ActualLRPInstanceKey)
					Expect(err).NotTo(HaveOccurred())

					_, err = sqlDB.ActualLRPGroupByProcessGuidAndIndex(logger, guid, index)
					Expect(err).To(HaveOccurred())
					Expect(err).To(Equal(models.ErrResourceNotFound))
				})
			})
		})

		Context("when the actualLRP does not exist", func() {
			It("does not return an error", func() {
				err := sqlDB.RemoveSuspectActualLRP(logger, &actualLRP.ActualLRPKey, &actualLRP.ActualLRPInstanceKey)
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})