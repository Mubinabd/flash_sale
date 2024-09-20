package handlers

import (
	"context"
	pb "flashSale_gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// @Summary       Create a Review
// @Description   Submit a review for a product
// @Tags          Review
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         CreateReviewReq body pb.CreateReviewReq true "Create Review Request"
// @Success       200  {object} pb.Void "Review created successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/reviews [post]
func (h *Handler) CreateReview(c *gin.Context) {
	req := &pb.CreateReviewReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, gin.H{"message": "Invalid request: " + err.Error()})
		return
	}

	_, err := h.Clients.Review.CreateReview(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Create review successfully"})

}

// @Summary       Get Product Rating
// @Description   Retrieve the average rating and total reviews for a product
// @Tags          Review
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// Param		  productId query int false "Product ID"
// @Success       200  {object} pb.ProductRatingRes "Product rating response"
// @Failure       400  {string}  string "Invalid request"
// @Failure       404  {string}  string "Product not found"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/reviews/{productId}/rating [get]
func (h *Handler) GetProductRating(c *gin.Context) {

	req := &pb.GetProductRatingReq{}
	id := c.Param("productId")
	req.ProductId = id
	res, err := h.Clients.Review.GetProductRating(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
