package main

import (
	"image"
	"image/color"
	"os"
	"fmt"
	"image/png"
	"log"
)

 type Putpixel func(x, y int)

 // 求绝对值
 func abs(x int) int {
     if x >= 0 {
	         return x
	     }
	    return -x
	}

func drawline(x0, y0, x1, y1 int, brush Putpixel) {
	     dx := abs(x1 - x0)
	    dy := abs(y1 - y0)
	     sx, sy := 1, 1
	    if x0 >= x1 {
	         sx = -1
		    }
    if y0 >= y1 {
	        sy = -1
	    }
	    err := dx - dy

    for {
	       brush(x0, y0)
	        if x0 == x1 && y0 == y1 {
		            return
		       }
	       e2 := err * 2
	        if e2 > -dy {
			             err -= dy
			            x0 += sx
			       }
		         if e2 < dx {
		           err += dx
		            y0 += sy
		        }
	     }
 }

func main() {

	    const (
        dx = 300
        dy = 500
	     )

	    // 需要保存的文件

    // 新建一个 指定大小的 RGBA位图
	    img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	    drawline(5, 5, dx-8, dy-10, func(x, y int) {
	        img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
	    })

    // 左右都画一条竖线
     for i := 0; i < dy; i++ {
	        img.Set(0, i, color.Black)
	       img.Set(dx-1, i, color.Black)
	    }

    imgcounter := 1011101010
     imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
     defer imgfile.Close()

    // 以PNG格式保存文件
	   err := png.Encode(imgfile, img)
	    if err != nil {
	        log.Fatal(err)
		     }
	 }

