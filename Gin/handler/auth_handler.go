package handler

import (
	"gin-project/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数无效"})
		return
	}

	if req.Username != "admin" || req.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	// --- 💡 修改在这里 ---
	// 调用 util 包中的 GenToken
	token, err := util.GenToken(req.Username)
	// --- 结束修改 ---

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token 生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
	})
}

func HelloHandler(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		// 理论上，如果中间件配置对了，这里总会 "exists"
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户信息"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"hello": "欢迎" + username.(string),
	})
}
