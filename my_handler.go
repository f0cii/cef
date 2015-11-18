package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_v8_capi.h"
#include "cef_app.h"
#include "cef_client.h"
#include "cef_helpers.h"
#include "cef_list_value.h"
*/
import "C"

import (
	"crypto/rand"
	"fmt"
	"image"
	"image/color" // for color.Alpha{a}
	"image/jpeg"
	"math"
	"os"
	"strconv"
	"strings"
)

func init() {
	// without this register .. At(), Bounds() functions will
	// caused memory pointer error!!
	//image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	//image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	//image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
}

func init_cef_handlers() {
	RegisterV8Handler("setResult", cef_handler_cefResult)
	RegisterV8Handler("callback", cef_handler_cefCallback)
	RegisterV8Handler("renderImage", cef_handler_renderImage)
	RegisterV8Handler("openWindow", cef_handler_openWindow)
}

// Cef调用
func cef_handler_cefCallback(browser *Browser, args []V8Value) (result interface{}) {
	fmt.Println("cef_handler_cefCallback")
	message := CefProcessMessageCreate(kV8CallbackMessage)
	valueList := message.GetArgumentList()
	//callbackName := V8ValueToString(args[0])
	//valueList.SetString(0, callbackName)
	for i := 0; i < len(args); i++ {
		v := args[i]
		switch {
		case C.cef_v8value_t_is_bool(v) == 1:
			logger.Printf("cef_handler_cefCallback [%v] bool\n", i)
			C.cef_list_value_t_set_bool(valueList.CStruct, C.int(i), C.v8ValueToBool(v))
		case C.cef_v8value_t_is_string(v) == 1:
			logger.Printf("cef_handler_cefCallback [%v] string\n", i)
			C.cef_list_value_t_set_string(valueList.CStruct, C.int(i), C.v8ValueToString(v))
		//case C.cef_v8value_t_is_null(v) == 1:
		//C.cef_list_value_t_set_null(valueList.CStruct, C.int(i), v)
		case C.cef_v8value_t_is_int(v) == 1:
			logger.Printf("cef_handler_cefCallback [%v] int\n", i)
			C.cef_list_value_t_set_int(valueList.CStruct, C.int(i), C.int(C.v8ValueToInt32(v)))
		case C.cef_v8value_t_is_uint(v) == 1:
			logger.Printf("cef_handler_cefCallback [%v] uint\n", i)
			C.cef_list_value_t_set_int(valueList.CStruct, C.int(i), C.int(C.v8ValueToUInt32(v)))
		case C.cef_v8value_t_is_double(v) == 1:
			logger.Printf("cef_handler_cefCallback [%v] double\n", i)
			C.cef_list_value_t_set_double(valueList.CStruct, C.int(i), C.double(C.v8ValueToDouble(v)))
		default:
			logger.Printf("cef_handler_cefCallback fail. [%v]\n", i)
		}
	}
	browser.SendProcessMessage(C.PID_BROWSER, message.CStruct)
	return
}

// 创建窗口
func cef_handler_cefResult(browser *Browser, args []V8Value) (result interface{}) {
	//n := len(args)
	//for i := 0; i < n; i++ {
	//v := V8ValueToString(args[i])
	//logger.Printf("cef_handler_set_result val_%v=%v\n", i, v)
	//}
	val := V8ValueToString(args[0])
	//logger.Printf("cef_handler_set_result val=%v\n", val)

	message := CefProcessMessageCreate(kV8EvalResultMessage)
	valueList := message.GetArgumentList()
	//val = `15`
	//valueList.SetInt(0, 10)
	valueList.SetString(0, val)
	browser.SendProcessMessage(C.PID_BROWSER, message.CStruct)

	return
}

// http://www.magpcss.org/ceforum/viewtopic.php?f=6&t=12586
// http://stackoverflow.com/questions/19414814/how-can-i-convert-canvas-image-data-into-an-array-of-per-pixel-text-values-with
func cef_handler_renderImage(browser *Browser, args []V8Value) (result interface{}) {
	//context.getImageData(0, 0, width, height)

	m := len(args)
	for i := 0; i < m; i++ {
		v := C.v8ValueToInt32(args[i])
		logger.Printf("cef_handler_set_image_data val_%v=%v\n", i, v)
	}

	//data := args[0]
	strImageData := V8ValueToString(args[0])
	width := int(C.v8ValueToInt32(args[1]))
	logger.Printf("cef_handler_set_image_data width=%v\n", width)
	height := int(C.v8ValueToInt32(args[2]))
	logger.Printf("cef_handler_set_image_data height=%v\n", height)

	data := strings.Split(strImageData, ",")
	n := len(data)
	logger.Printf("cef_handler_set_image_data n=%v\n", n)

	rgba := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < n; i += 4 {
		var x = math.Floor(float64((i / 4) % width))
		var y = math.Floor(float64((i / 4) / width))
		r, _ := strconv.Atoi(data[i])   //data.C.v8ValueToUInt32(C.v8ValueGetValueByIndex(data, C.int(i)))
		g, _ := strconv.Atoi(data[i+1]) //C.v8ValueToUInt32(C.v8ValueGetValueByIndex(data, C.int(i+1)))
		b, _ := strconv.Atoi(data[i+2]) //C.v8ValueToUInt32(C.v8ValueGetValueByIndex(data, C.int(i+2)))
		a, _ := strconv.Atoi(data[i+3]) //C.v8ValueToUInt32(C.v8ValueGetValueByIndex(data, C.int(i+3)))
		if int(x) <= width && int(y) <= height {
			rgba.Set(int(x), int(y), color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	//生成Guid字串
	filename := pseudo_uuid() + ".jpeg"
	file, err := os.Create(filename)
	if err != nil {
		logger.Fatal(err)
		return ""
	}
	defer file.Close()

	logger.Println(rgba.Bounds()) //(0,0)-(500,200)，图片边界
	logger.Println(rgba.Opaque()) //false，是否图片完全透明
	jpeg.Encode(file, rgba, nil)  //将image信息写入文件中
	return filename
}

// 打开窗口
func cef_handler_openWindow(browser *Browser, args []V8Value) (result interface{}) {
	url := V8ValueToString(args[0])
	message := CefProcessMessageCreate(kV8OpenWindowMessage)
	valueList := message.GetArgumentList()
	valueList.SetString(0, url)
	browser.SendProcessMessage(C.PID_BROWSER, message.CStruct)
	return
}

// 生成UUID
func pseudo_uuid() (uuid string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}
