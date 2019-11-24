// 由GOVCL UI设计器自动生成，不要编辑
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TJsForm struct {
    *vcl.TForm
    Panel1  *vcl.TPanel
    Panel2  *vcl.TPanel
    Button1 *vcl.TButton
    Button2 *vcl.TButton
    Memo1   *vcl.TMemo

    //::private::
    TJsFormFields
}

var JsForm *TJsForm




// 以字节形式加载
// vcl.Application.CreateForm(jsFormBytes, &JsForm)

var (
    jsFormBytes = []byte {
        0x54, 0x50, 0x46, 0x30, 0x05, 0x54, 0x46, 0x6F, 0x72, 0x6D, 0x06, 0x4A, 
        0x73, 0x46, 0x6F, 0x72, 0x6D, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x00, 
        0x03, 0x54, 0x6F, 0x70, 0x02, 0x00, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69, 
        0x6F, 0x6E, 0x12, 0x04, 0x00, 0x00, 0x00, 0x4A, 0x00, 0x73, 0x00, 0x16, 
        0x7F, 0x91, 0x8F, 0x0C, 0x43, 0x6C, 0x69, 0x65, 0x6E, 0x74, 0x48, 0x65, 
        0x69, 0x67, 0x68, 0x74, 0x03, 0x41, 0x01, 0x0B, 0x43, 0x6C, 0x69, 0x65, 
        0x6E, 0x74, 0x57, 0x69, 0x64, 0x74, 0x68, 0x03, 0xA6, 0x02, 0x05, 0x43, 
        0x6F, 0x6C, 0x6F, 0x72, 0x07, 0x09, 0x63, 0x6C, 0x42, 0x74, 0x6E, 0x46, 
        0x61, 0x63, 0x65, 0x0C, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x43, 0x68, 0x61, 
        0x72, 0x73, 0x65, 0x74, 0x07, 0x0F, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4C, 
        0x54, 0x5F, 0x43, 0x48, 0x41, 0x52, 0x53, 0x45, 0x54, 0x0A, 0x46, 0x6F, 
        0x6E, 0x74, 0x2E, 0x43, 0x6F, 0x6C, 0x6F, 0x72, 0x07, 0x0C, 0x63, 0x6C, 
        0x57, 0x69, 0x6E, 0x64, 0x6F, 0x77, 0x54, 0x65, 0x78, 0x74, 0x0B, 0x46, 
        0x6F, 0x6E, 0x74, 0x2E, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0xF3, 
        0x09, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x4E, 0x61, 0x6D, 0x65, 0x06, 0x06, 
        0x54, 0x61, 0x68, 0x6F, 0x6D, 0x61, 0x0A, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 
        0x53, 0x74, 0x79, 0x6C, 0x65, 0x0B, 0x00, 0x0E, 0x4F, 0x6C, 0x64, 0x43, 
        0x72, 0x65, 0x61, 0x74, 0x65, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x08, 0x08, 
        0x50, 0x6F, 0x73, 0x69, 0x74, 0x69, 0x6F, 0x6E, 0x07, 0x0A, 0x70, 0x6F, 
        0x44, 0x65, 0x73, 0x69, 0x67, 0x6E, 0x65, 0x64, 0x06, 0x53, 0x63, 0x61, 
        0x6C, 0x65, 0x64, 0x08, 0x0D, 0x50, 0x69, 0x78, 0x65, 0x6C, 0x73, 0x50, 
        0x65, 0x72, 0x49, 0x6E, 0x63, 0x68, 0x02, 0x60, 0x0A, 0x54, 0x65, 0x78, 
        0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x10, 0x00, 0x06, 0x54, 
        0x50, 0x61, 0x6E, 0x65, 0x6C, 0x06, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x31, 
        0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x00, 0x03, 0x54, 0x6F, 0x70, 0x03, 
        0x18, 0x01, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x03, 0xA6, 0x02, 0x06, 
        0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x29, 0x05, 0x41, 0x6C, 0x69, 
        0x67, 0x6E, 0x07, 0x08, 0x61, 0x6C, 0x42, 0x6F, 0x74, 0x74, 0x6F, 0x6D, 
        0x0A, 0x42, 0x65, 0x76, 0x65, 0x6C, 0x4F, 0x75, 0x74, 0x65, 0x72, 0x07, 
        0x06, 0x62, 0x76, 0x4E, 0x6F, 0x6E, 0x65, 0x08, 0x54, 0x61, 0x62, 0x4F, 
        0x72, 0x64, 0x65, 0x72, 0x02, 0x00, 0x00, 0x06, 0x54, 0x50, 0x61, 0x6E, 
        0x65, 0x6C, 0x06, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x32, 0x04, 0x4C, 0x65, 
        0x66, 0x74, 0x03, 0xED, 0x01, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x00, 0x05, 
        0x57, 0x69, 0x64, 0x74, 0x68, 0x03, 0xB9, 0x00, 0x06, 0x48, 0x65, 0x69, 
        0x67, 0x68, 0x74, 0x02, 0x29, 0x05, 0x41, 0x6C, 0x69, 0x67, 0x6E, 0x07, 
        0x07, 0x61, 0x6C, 0x52, 0x69, 0x67, 0x68, 0x74, 0x0A, 0x42, 0x65, 0x76, 
        0x65, 0x6C, 0x4F, 0x75, 0x74, 0x65, 0x72, 0x07, 0x06, 0x62, 0x76, 0x4E, 
        0x6F, 0x6E, 0x65, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72, 
        0x02, 0x00, 0x00, 0x07, 0x54, 0x42, 0x75, 0x74, 0x74, 0x6F, 0x6E, 0x07, 
        0x42, 0x75, 0x74, 0x74, 0x6F, 0x6E, 0x31, 0x04, 0x4C, 0x65, 0x66, 0x74, 
        0x02, 0x0E, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x09, 0x05, 0x57, 0x69, 0x64, 
        0x74, 0x68, 0x02, 0x4B, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 
        0x19, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6F, 0x6E, 0x12, 0x02, 0x00, 
        0x00, 0x00, 0x6E, 0x78, 0x9A, 0x5B, 0x0B, 0x4D, 0x6F, 0x64, 0x61, 0x6C, 
        0x52, 0x65, 0x73, 0x75, 0x6C, 0x74, 0x02, 0x01, 0x08, 0x54, 0x61, 0x62, 
        0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 0x00, 0x00, 0x00, 0x07, 0x54, 0x42, 
        0x75, 0x74, 0x74, 0x6F, 0x6E, 0x07, 0x42, 0x75, 0x74, 0x74, 0x6F, 0x6E, 
        0x32, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x5F, 0x03, 0x54, 0x6F, 0x70, 
        0x02, 0x09, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x02, 0x4B, 0x06, 0x48, 
        0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x19, 0x07, 0x43, 0x61, 0x70, 0x74, 
        0x69, 0x6F, 0x6E, 0x12, 0x02, 0x00, 0x00, 0x00, 0xD6, 0x53, 0x88, 0x6D, 
        0x0B, 0x4D, 0x6F, 0x64, 0x61, 0x6C, 0x52, 0x65, 0x73, 0x75, 0x6C, 0x74, 
        0x02, 0x02, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 
        0x01, 0x00, 0x00, 0x00, 0x00, 0x05, 0x54, 0x4D, 0x65, 0x6D, 0x6F, 0x05, 
        0x4D, 0x65, 0x6D, 0x6F, 0x31, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x00, 
        0x03, 0x54, 0x6F, 0x70, 0x02, 0x00, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 
        0x03, 0xA6, 0x02, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x03, 0x18, 
        0x01, 0x05, 0x41, 0x6C, 0x69, 0x67, 0x6E, 0x07, 0x08, 0x61, 0x6C, 0x43, 
        0x6C, 0x69, 0x65, 0x6E, 0x74, 0x0C, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x43, 
        0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x07, 0x0F, 0x44, 0x45, 0x46, 0x41, 
        0x55, 0x4C, 0x54, 0x5F, 0x43, 0x48, 0x41, 0x52, 0x53, 0x45, 0x54, 0x0A, 
        0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x43, 0x6F, 0x6C, 0x6F, 0x72, 0x07, 0x0C, 
        0x63, 0x6C, 0x57, 0x69, 0x6E, 0x64, 0x6F, 0x77, 0x54, 0x65, 0x78, 0x74, 
        0x0B, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 
        0x02, 0xF3, 0x09, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x4E, 0x61, 0x6D, 0x65, 
        0x06, 0x0B, 0x43, 0x6F, 0x75, 0x72, 0x69, 0x65, 0x72, 0x20, 0x4E, 0x65, 
        0x77, 0x0A, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x53, 0x74, 0x79, 0x6C, 0x65, 
        0x0B, 0x00, 0x0D, 0x4C, 0x69, 0x6E, 0x65, 0x73, 0x2E, 0x53, 0x74, 0x72, 
        0x69, 0x6E, 0x67, 0x73, 0x01, 0x06, 0x12, 0x66, 0x75, 0x6E, 0x63, 0x74, 
        0x69, 0x6F, 0x6E, 0x20, 0x54, 0x65, 0x73, 0x74, 0x31, 0x28, 0x29, 0x20, 
        0x7B, 0x14, 0x1D, 0x00, 0x00, 0x00, 0x20, 0x20, 0x72, 0x65, 0x74, 0x75, 
        0x72, 0x6E, 0x20, 0x22, 0xE8, 0xBF, 0x99, 0xE6, 0x98, 0xAF, 0xE4, 0xB8, 
        0x80, 0xE4, 0xB8, 0xAA, 0xE6, 0xB5, 0x8B, 0xE8, 0xAF, 0x95, 0x22, 0x06, 
        0x01, 0x7D, 0x12, 0x00, 0x00, 0x00, 0x00, 0x06, 0x0E, 0x61, 0x6C, 0x65, 
        0x72, 0x74, 0x28, 0x54, 0x65, 0x73, 0x74, 0x31, 0x28, 0x29, 0x29, 0x12, 
        0x00, 0x00, 0x00, 0x00, 0x06, 0x0E, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6E, 
        0x20, 0x54, 0x65, 0x73, 0x74, 0x31, 0x28, 0x29, 0x00, 0x0A, 0x50, 0x61, 
        0x72, 0x65, 0x6E, 0x74, 0x46, 0x6F, 0x6E, 0x74, 0x08, 0x08, 0x54, 0x61, 
        0x62, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 0x01, 0x00, 0x00, 0x00}
)
