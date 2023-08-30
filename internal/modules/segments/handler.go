package segments

import (
	"net/http"

	"github.com/boichique/avito-test-task/internal/echox"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// @Summary Create a new segment
// @Description Create a new segment with the given slug
// @Accept json
// @Produce json
// @Param request body CreateSegmentRequest true "Segment request payload"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /segments [post]
func (h *Handler) Create(c echo.Context) error {
	req, err := echox.Bind[CreateSegmentRequest](c)
	if err != nil {
		return err
	}

	segment := &Segment{
		Slug: req.Slug,
	}

	return h.service.Create(c.Request().Context(), segment)
}

// @Summary Delete a segment
// @Description Delete a segment by slug
// @Accept json
// @Produce json
// @Param request body DeleteSegmentRequest true "Segment request payload"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /segments [delete]
func (h *Handler) Delete(c echo.Context) error {
	req, err := echox.Bind[DeleteSegmentRequest](c)
	if err != nil {
		return err
	}

	return h.service.Delete(c.Request().Context(), req.Slug)
}

// @Summary Get segments by user ID
// @Description Get segments by user ID
// @Accept json
// @Produce json
// @Param request body GetUserSegmentsRequest true "User segments request payload"
// @Success 200 {object} GetUserSegmentsResponse "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /segments [get]
func (h *Handler) GetByUserID(c echo.Context) error {
	req, err := echox.Bind[GetUserSegmentsRequest](c)
	if err != nil {
		return err
	}

	segments, err := h.service.GetByUserID(c.Request().Context(), req.UserID)
	if err != nil {
		return err
	}

	res := &GetUserSegmentsResponse{
		Segments: segments,
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Update segments by user ID
// @Description Update segments by user ID
// @Accept json
// @Produce json
// @Param request body UpdateUserSegmentsRequest true "User segments request payload"
// @Success 200 {object} GetUserSegmentsResponse "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /segments [post]
func (h *Handler) UpdateSegmentsByUserID(c echo.Context) error {
	req, err := echox.Bind[UpdateUserSegmentsRequest](c)
	if err != nil {
		return err
	}

	return h.service.UpdateSegmentsByUserID(c.Request().Context(), req.UserID, req.AddSegments, req.DeleteSegments)
}
