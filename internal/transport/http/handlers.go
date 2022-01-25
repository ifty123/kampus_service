package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/ifty123/kampus_service/internal/services"
	mhsConst "github.com/ifty123/kampus_service/pkg/common/const"
	"github.com/ifty123/kampus_service/pkg/dto"
	mhsErrors "github.com/ifty123/kampus_service/pkg/errors"

	"github.com/labstack/echo"
)

type HttpHandler struct {
	service services.Services
}

func NewHttpHandler(e *echo.Echo, srv services.Services) {
	handler := &HttpHandler{
		srv,
	}
	e.GET("api/v1/latihan/ping", handler.Ping)
	e.POST("api/v1/latihan/mahasiswa-alamat", handler.SaveMahasiswaAlamat)
	e.PATCH("api/v1/latihan/mahasiswa", handler.UpdateMhs)
	e.GET("api/v1/latihan/mahasiswa", handler.TampilMahasiswa)
	e.GET("api/v1/latihan/mahasiswa-alamat", handler.TampilMahasiswaAlamat)
	e.POST("api/v1/latihan/mahasiswa-nama", handler.SaveMahasiswa)
	e.POST("api/v1/latihan/mahasiswa-alamatID", handler.SaveAlamatWithID)
	e.PUT("api/v1/latihan/alamat", handler.UpdateAlamatMahasiswa)
	e.GET("api/v1/latihan/mahasiswa-GetID", handler.GetMahasiswaByID)
	e.GET("api/v1/latihan/mahasiswa-GetNama", handler.GetMahasiswaByName)
	e.GET("api/v1/latihan/mahasiswa-Get-OnlyAlamat", handler.GetOnlyAlamat)
	e.GET("api/v1/latihan/mahasiswa-Get-Nama", handler.GetNamaIfNotNull)

	e.POST("api/v1/latihan/dosen-simpan", handler.SaveDosens)
	e.POST("api/v1/latihan/dosen-simpan-alamat", handler.SaveDosensAlamat)
	e.POST("api/v1/latihan/dosen-simpan-dosenAlamat", handler.SaveDosenAndAlamat)
	e.GET("api/v1/latihan/dosen-tampil", handler.TampilDosens)
	e.GET("api/v1/latihan/dosenAlamat-tampil", handler.TampilDosenAlamat)
	e.PATCH("api/v1/latihan/dosen-update", handler.UpdateDosens)
	e.POST("api/v1/latihan/dosenById-tampil", handler.TampilDosenByID)
	e.PATCH("api/v1/latihan/dosenAlamat-update", handler.UpdateDosenAlamat)

	//baru
	e.GET("api/v1/latihan/dad-jokes", handler.GetRandomDadJokes)

}

func (h *HttpHandler) Ping(c echo.Context) error {

	version := os.Getenv("VERSION")
	if version == "" {
		version = "pong"
	}

	data := version

	return c.JSON(http.StatusOK, data)

}

//Simpan mahasiswa alamat
func (h *HttpHandler) SaveMahasiswaAlamat(c echo.Context) error {
	postDTO := dto.MahasiswaAlamatReqDTO{}
	//byte ke json
	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.SaveMahasiswaAlamat(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

//Simpan mahasiswa
func (h *HttpHandler) SaveMahasiswa(c echo.Context) error {
	postDTO := dto.MahasiswaReqDTO{}
	//byte ke json
	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.SaveMahasiswa(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

//Simpan mahasiswa
func (h *HttpHandler) SaveAlamatWithID(c echo.Context) error {
	postDTO := dto.AlamatReqWithIDDTO{}
	//byte ke json
	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	fmt.Println(postDTO)
	fmt.Println("Ini masuk data")
	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.SaveAlamatFromMahasiswa(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

//update mahasiswa dengan ID ke sekian
func (h *HttpHandler) UpdateMhs(c echo.Context) error {
	postDTO := dto.UpadeMahasiswaNamaReqDTO{}
	//byte ke json
	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.UpdateMahasiswaNama(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

//update alamats
//update mahasiswa dengan ID ke sekian
func (h *HttpHandler) UpdateAlamatMahasiswa(c echo.Context) error {
	postDTO := dto.UpdateAlamatMahasiswaReq{}
	//byte ke json
	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.UpdateAlamatMahasiswa(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.UpdateSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) TampilMahasiswa(c echo.Context) error {
	data, err := h.service.TampilMahasiswa()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) GetMahasiswaByID(c echo.Context) error {
	postDTO := dto.GetMahasiswaByIDReqDTO{}

	postDTO.Authorization = c.Request().Header.Get("Authorization")
	postDTO.DateTime = c.Request().Header.Get("datetime")
	postDTO.Signature = c.Request().Header.Get("signature")

	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	data, err := h.service.GetMahasiswaByID(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

//getMahasiswa By Name
func (h *HttpHandler) GetMahasiswaByName(c echo.Context) error {
	dataFromDTO := dto.GetMahasiswaByNamaReqDTO{}

	//fungsi bind : menampung payload yg dikirim, dari struck
	if err := c.Bind(&dataFromDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	//validasi
	fmt.Println("data nama : ", dataFromDTO)
	err := dataFromDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	//masuk ke service
	data, err := h.service.GetMahasiswaByName(&dataFromDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: mhsErrors.ErrorDataNotFound,
			Data:    nil,
		})
	}

	respon := dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, respon)

}

//tampil alamat
func (h *HttpHandler) TampilMahasiswaAlamat(c echo.Context) error {
	data, err := h.service.TampilMahasiswaAlamats()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) GetOnlyAlamat(c echo.Context) error {
	data, err := h.service.GetOnlyAlamat()

	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// jika berhasil
	var respon = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, respon)
}

func (h *HttpHandler) GetNamaIfNotNull(c echo.Context) error {
	dataFromDTO := dto.GetNamaIfNotNull{}

	//fungsi bind : menampung payload yg dikirim, dari struck
	if err := c.Bind(&dataFromDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	//masuk ke service
	data, err := h.service.GetAllIfNull(&dataFromDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: mhsErrors.ErrorDataNotFound,
			Data:    nil,
		})
	}

	respon := dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, respon)

}

func (h *HttpHandler) GetRandomDadJokes(c echo.Context) error {
	postDTO := dto.GetDadJokesInternalReqDTO{}

	postDTO.Authorization = c.Request().Header.Get("Authorization")
	postDTO.DateTime = c.Request().Header.Get("datetime")
	postDTO.Signature = c.Request().Header.Get("signature")

	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	data, err := h.service.GetIntegDadJoke(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

//tampil mahasiswa
func (h *HttpHandler) GetTampilMahasiswa(c echo.Context) error {
	postDTO := dto.GetMahasiswaByIDReqDTO{}

	postDTO.Authorization = c.Request().Header.Get("Authorization")
	postDTO.DateTime = c.Request().Header.Get("datetime")
	postDTO.Signature = c.Request().Header.Get("signature")

	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	data, err := h.service.GetMahasiswaByID(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

//status error
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	//error
	switch err {
	case mhsErrors.ErrInternalServerError:
		return http.StatusInternalServerError
	case mhsErrors.ErrNotFound:
		return http.StatusNotFound
	case mhsErrors.ErrConflict:
		return http.StatusConflict
	case mhsErrors.ErrInvalidRequest:
		return http.StatusBadRequest
	case mhsErrors.ErrFailAuth:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
