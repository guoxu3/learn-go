package upload

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"code.google.com/p/graphics-go/graphics"
	"strings"
	"fmt"
)


// 计算按比例缩放之后较长边的长度,需要时一个可以被2整除的整数
func long_side(short, long int) int  {
	tmp := float64(long) * ( 200.0 / float64(short))
	tmp1 := int(tmp)
	if tmp1 % 2 == 0 {
		return tmp1
	} else {
		return (tmp1 + 1)
	}
}

// 获取缩放的比例
func new_size(hig, wid int) (new_hig, new_wid int) {
	if hig < wid {
		new_hig = 200
		new_wid = tmp(hig, wid)
	} else if wid < hig {
		new_wid = 200
		new_hig = tmp(wid, hig)
	} else {
		new_hig, new_wid = 200, 200
	}
	return new_hig, new_wid
}

// 根据压缩后的图片大小,获取裁切的两个点坐标
func sub_size(height, width int) (x0, y0, x1, y1 int)  {
	if height == 200 {
		y0, y1 = 0, 200
		x0 = (width - 200) / 2
		x1 = 200 + x0
	} else {
		x0, x1 = 0, 200
		y0 = (height - 200) / 2
		y1 = 200 + y0
	}
	return x0, y0, x1, y1
}


// 将图片等比例缩放之后,裁切200x200大小
func Resize(dir, image_name string) bool {
	// 打开源文件
	src0, err := os.Open(dir + image_name)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer src0.Close()

	// 获取原始图片的高度和宽度
	conf, img_type, err := image.DecodeConfig(src0)
	if err != nil {
		return false
	}
	var suffix string
	if img_type == "png" {
		suffix = ".png"
	} else if img_type == "jpeg" {
		suffix = ".jpeg"
	} else {
		return false
	}

	file_suffix := path.Ext(image_name)
	filename := strings.TrimSuffix(image_name, file_suffix)
	tmp_name := filename + "_tmp" + suffix
	tmp_name1 := filename + "_tmp1" + suffix

	height, width := conf.Height, conf.Width
	// 获取等比例缩放后的高度和宽度
	new_hig, new_wid := new_size(height, width)

	src, err := os.Open(dir + image_name)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer src.Close()

	// 打开目标文件
	dst, err := os.Create(dir + tmp_name)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer dst.Close()

	// 解码
	var (
		src1 image.Image
		err1 error
	)

	if img_type == "png" {
		src1, err1 = png.Decode(src)
		if err1 != nil {
			fmt.Println(err)
			return false
		}
	} else if img_type == "jpeg"{
		src1, err1 = jpeg.Decode(src)
		if err1 != nil {
			fmt.Println(err)
			return false
		}
	} else {
		return false
	}

	// 按比例缩放,保证最小边为200
	bounds := image.Rect(0, 0, new_wid, new_hig)
	src2 := image.NewRGBA(bounds)
   	graphics.Scale(src2, src1)

	if img_type == "png" {
		err = png.Encode(dst, src2)
   		if err != nil {
			fmt.Println(1111)
			return false
		}
	} else if img_type == "jpeg" {
		err = jpeg.Encode(dst, src2, &jpeg.Options{90})
   		if err != nil {
			fmt.Println(err)
			return false
		}
	} else {
		return  false
	}

	// 裁切出一个200x200 的图片
	if new_hig != new_wid {
		// ===========================
		src3, err := os.Open(dir + tmp_name)
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer src3.Close()

		// 打开目标文件
		dst1, err := os.Create(dir + tmp_name1)
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer dst1.Close()

		// 图片文件解码
		src4, _, _ := image.Decode(src3)
		if err != nil {
			fmt.Println(err)
			return false
		}
		rgbImg := src4.(*image.YCbCr)
		x0, y0, x1, y1 := sub_size(new_hig, new_wid)
		fmt.Println(x0,y0,x1,y1)
		src5 := rgbImg.SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
		if img_type == "png" {
			err = png.Encode(dst1, src5)
			if err != nil {
				fmt.Println(err)
				return false
			}
		} else if img_type == "jpeg" {
			err = jpeg.Encode(dst1, src5, &jpeg.Options{90})
			if err != nil {
				fmt.Println(err)
				return false
			}
		} else {
			return  false
		}
	}
	return true
}
