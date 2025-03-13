package home

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type home struct {
	Banners []string `json:"banners"`
	Image1  []string `json:"image1"`
	Image2  []string `json:"image2"`
}

var bannerDataList = []string{
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/20231108130714169942003451044.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/20241101033654173040341496112.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/83596a78-c842-4ca6-b468-abe1065738de.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/DALL%C2%B7E%202025-02-23%2020.44.41%20-%20A%20minimalist%20sple.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/DALL%C2%B7E%202025-02-23%2020.45.24%20-%20A%20minimalist%20.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-20-10.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-21-12.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-34-18.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-38-07.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-38-50.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-39-11.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-40-08.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-41-19.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-45-00.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-45-16.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-45-40.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-46-36.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-46-52.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-47-07.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-48-27.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-49-24.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-50-17.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-50-28.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-52-24.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-53-32.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_22-11-19.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/bizhihui_com_20231109132446169950748680704.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/bizhihui_com_20231109194252169953017213565.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/bizhihui_com_20231109211254169953557454014.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/vecteezy_silhouette-of-coral-reef-with-fish-on-blue-sea-background_14212401.jpg",
}

var imgDataList = []string{
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/20231108130714169942003451044.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/20241101033654173040341496112.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/83596a78-c842-4ca6-b468-abe1065738de.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/DALL%C2%B7E%202025-02-23%2020.44.41%20-%20A%20minimalist%20sple.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/DALL%C2%B7E%202025-02-23%2020.45.24%20-%20A%20minimalist%20.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-20-10.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-21-12.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-34-18.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-38-07.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-38-50.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-39-11.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-40-08.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-41-19.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-45-00.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-45-16.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-45-40.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-46-36.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-46-52.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-47-07.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-48-27.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-49-24.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-50-17.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-50-28.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-52-24.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_21-53-32.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/Snipaste_2025-02-23_22-11-19.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/bizhihui_com_20231109132446169950748680704.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/bizhihui_com_20231109194252169953017213565.jpg",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/bizhihui_com_20231109211254169953557454014.png",
	"https://bbucket-1253575676.cos.ap-guangzhou.myqcloud.com/public/vecteezy_silhouette-of-coral-reef-with-fish-on-blue-sea-background_14212401.jpg",
}

func Home(c *gin.Context) {
	h := &home{}
	bannerNum := 3
	indexList := make([]int, len(bannerDataList))
	for i, _ := range bannerDataList {
		indexList[i] = i
	}
	list := randList(indexList, bannerNum)
	bannerList := make([]string, bannerNum)
	for i := 0; i < bannerNum; i++ {
		bannerList[i] = bannerDataList[list[i]]
	}

	imgNum := 10
	indexList = make([]int, len(imgDataList))
	for i, _ := range imgDataList {
		indexList[i] = i
	}
	list = randList(indexList, imgNum)
	imgList := make([]string, imgNum)
	for i := 0; i < imgNum; i++ {
		imgList[i] = imgDataList[list[i]]
	}

	h.Banners = bannerList
	h.Image1 = imgList[:5]
	h.Image2 = imgList[5:]

	c.JSON(http.StatusOK, h)
}

func randList(indexList []int, num int) []int {
	list := make([]int, num)
	for i := 0; i < num; i++ {
		l := len(indexList)
		index := rand.Intn(l)
		list[i] = indexList[index]
		indexList = append(indexList[:index], indexList[index+1:]...)
	}
	return list
}
