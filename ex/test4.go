package main

import (
	"github.com/GUMI-golang/giame"
	"bytes"
	"encoding/gob"
	"github.com/GUMI-golang/gumi/gcore"
	"image/color"
	"encoding/base64"
	"fmt"
)

func main() {
	var f = giame.NewUniformFiller(color.RGBA{255, 0, 0, 255})
	var buf =bytes.NewBuffer(nil)
	//
	gcore.Must(gob.NewEncoder(buf).Encode(f))
	result := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println(result)
	var de giame.Filler
	gcore.Must(gob.NewDecoder(bytes.NewReader(gcore.MustValue(base64.StdEncoding.DecodeString(result)).([]byte))).Decode(de))
	fmt.Println(de)

}