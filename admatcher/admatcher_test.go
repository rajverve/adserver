package admatcher_test

import (
	"github.com/rajverve/adserver/admatcher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/rajverve/protobuf"
)

var _ = Describe("Admatcher", func() {
	var (
		ad1 = admatcher.Ad{"ad1"}
		ad2 = admatcher.Ad{"ad2"}

		seg1 = SupplySegment{Audience: "Moms", PlaceName: "Target", PlaceType: "Mall"}
		seg2 = SupplySegment{Audience: "Dads"}
		seg3 = SupplySegment{PlaceType: "Pizza"}
		seg4 = SupplySegment{Audience: "Kids", PlaceType: "Mall"}
	)

	Describe("Adding Ads for Segments", func() {
		Context("to empty matcher", func() {
			It("should return be able to find adds that were added", func() {
				admatcher.AddAd(&seg1, &ad1)
				Expect(admatcher.AdsForSegment(&seg1)).To(Equal([]admatcher.Ad{ad1}))
			})
		})
	})

	Describe("Finding Ads for Segments", func() {
		BeforeEach(func() {
			admatcher.AddAd(&seg1, &ad1)
			admatcher.AddAd(&seg2, &ad1)
			admatcher.AddAd(&seg3, &ad2)
			admatcher.AddAd(&seg4, &ad2)
		})

		It("should match partial matches", func() {
			Expect(admatcher.AdsForSegment(&SupplySegment{"Kids", "Round Table", "Pizza"})).To(Equal([]admatcher.Ad{ad2}))
		})

		It("should not match segments that don't match", func() {
			Expect(admatcher.AdsForSegment(&SupplySegment{Audience:"Kooks"})).To(Equal([]admatcher.Ad{}))
		})
	})
})
