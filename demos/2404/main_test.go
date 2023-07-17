package main

import (
	"testing"
)

type Example struct {
	A *int `json: A`
}

func TestMain(t *testing.T) {
	// r := gin.Default()
	// r.POST("/test", func(ctx *gin.Context) {
	// 	tt := Example{}
	// 	ctx.BindJSON(&tt)
	// 	ctx.JSON(http.StatusOK, tt)
	// })
	// payload := `{"A":1}`
	// req := httptest.NewRequest("POST", "/test", strings.NewReader(payload))
	// req.Header.Set("Content-Type", "application/json")
	// w := httptest.NewRecorder()
	// r.ServeHTTP(w, req)
	// fmt.Println(w.Body.String())
	// ���� Redis �ͻ���
}
