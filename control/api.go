package control

import (
	"TTGM/modsls"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取品质对应的颜色
func getQualityColor(q modsls.Quality) string {
	switch q {
	case modsls.Red:
		return "#FF0000" // 红色
	case modsls.Gold:
		return "#FFD700" // 金色
	case modsls.Purple:
		return "#800080" // 紫色
	case modsls.Blue:
		return "#0000FF" // 蓝色
	case modsls.Green:
		return "#008000" // 绿色
	case modsls.White:
		return "#FFFFFF" // 白色
	default:
		return "#000000"
	}
}

// RandomArticle 随机抽取物品
func RandomArticle(c *gin.Context) {
	var articles []modsls.Articles
	if err := modsls.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if len(articles) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no articles found"})
		return
	}

	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 计算总权重
	totalWeight := 0.0
	for _, article := range articles {
		totalWeight += modsls.GetQualityWeight(article.Quality)
	}

	// 生成随机数
	randomValue := rand.Float64() * totalWeight

	// 根据权重选择物品
	currentWeight := 0.0
	var selectedArticle modsls.Articles
	for _, article := range articles {
		currentWeight += modsls.GetQualityWeight(article.Quality)
		if randomValue <= currentWeight {
			selectedArticle = article
			break
		}
	}

	// 返回结果，包含物品信息和对应的颜色
	c.JSON(http.StatusOK, gin.H{
		"name":    selectedArticle.Name,
		"quality": selectedArticle.Quality,
		"color":   getQualityColor(selectedArticle.Quality),
		"space":   gin.H{"width": selectedArticle.Space.Width, "height": selectedArticle.Space.Height},
		"price":   selectedArticle.Price,
		"type":    selectedArticle.Type,
	})
}
