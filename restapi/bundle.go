package main

import (
	"bytes"
	"compress/gzip"
	"github.com/FactomProject/gobundle"
)

var _data = []byte{
	0x1F, 0x8B, 0x08, 0x00,    0x00, 0x09, 0x6E, 0x88,    0x02, 0xFF, 0x4A, 0x49,    0x2C, 0x49, 0xD4, 0xCF, 
	0x28, 0xC9, 0xCD, 0xD1,    0x4B, 0x2F, 0x2F, 0x60,    0xA0, 0x0D, 0x30, 0x80,    0x00, 0x5C, 0xB4, 0x81, 
	0xB9, 0x21, 0x82, 0x0D,    0x16, 0x37, 0x34, 0x34,    0x35, 0x30, 0x64, 0x50,    0x60, 0xA0, 0x07, 0x28, 
	0x2D, 0x2E, 0x49, 0x2C,    0x02, 0x5A, 0x4F, 0x6D,    0x4F, 0x0E, 0x11, 0x60,    0x03, 0x8A, 0x7B, 0x3B, 
	0x2E, 0x4E, 0x9B, 0x8C,    0xD4, 0xC4, 0x14, 0x20,    0xCD, 0x69, 0x53, 0x92,    0x59, 0x92, 0x93, 0x6A, 
	0xE7, 0x97, 0x0F, 0x0C,    0x96, 0x4A, 0xE7, 0x8C,    0xC4, 0xCC, 0x3C, 0x85,    0x20, 0xD7, 0xE0, 0x10, 
	0x05, 0xC7, 0x00, 0x4F,    0x1B, 0x7D, 0x88, 0x1C,    0x48, 0x55, 0x71, 0x72,    0x51, 0x66, 0x41, 0x09, 
	0x88, 0xC9, 0x99, 0x56,    0x9A, 0x97, 0x5C, 0x92,    0x99, 0x9F, 0xA7, 0x50,    0x52, 0x94, 0x9A, 0xAA, 
	0x01, 0x4A, 0x50, 0x9A,    0x0A, 0xD5, 0x20, 0x09,    0x05, 0x20, 0xC8, 0x4C,    0x53, 0xD0, 0x28, 0xA9, 
	0x2C, 0x48, 0xCD, 0x4F,    0x83, 0xCA, 0xD8, 0xDA,    0x2A, 0xA8, 0xE7, 0x27,    0x65, 0xA5, 0x26, 0x97, 
	0xA8, 0x23, 0x29, 0x03,    0x81, 0x94, 0xFC, 0xE4,    0xD2, 0xDC, 0xD4, 0xBC,    0x12, 0xBD, 0xF2, 0xA2, 
	0xCC, 0x92, 0x54, 0x0D,    0x75, 0x9B, 0xD2, 0x1C,    0x3B, 0x75, 0x4D, 0x6B,    0x64, 0x25, 0x69, 0xF9, 
	0x45, 0x0A, 0x1A, 0x65,    0x89, 0x45, 0x0A, 0x99,    0x0A, 0x40, 0x77, 0xA1,    0x59, 0x85, 0xD3, 0x9C, 
	0x9C, 0x4C, 0x3B, 0x75,    0x05, 0x6D, 0x85, 0x4C,    0x54, 0xB3, 0x40, 0x00,    0xEE, 0xE2, 0xE8, 0xCC, 
	0x58, 0x54, 0xD9, 0x5A,    0xBC, 0x2E, 0xD3, 0x47,    0x75, 0x5A, 0xAD, 0x42,    0x6A, 0x4E, 0x71, 0x2A, 
	0x7E, 0xDF, 0x28, 0xD8,    0xDA, 0x29, 0x80, 0x5C,    0x01, 0x76, 0x34, 0x42,    0x27, 0x88, 0x01, 0x22, 

	0x6C, 0xF4, 0xE1, 0x21,    0x6A, 0xA3, 0x0F, 0x8D,    0x0B, 0x9B, 0xA4, 0xFC,    0x94, 0x4A, 0x85, 0xFC, 
	0xBC, 0x9C, 0xFC, 0xC4,    0x14, 0x5B, 0x75, 0xB0,    0x5B, 0xAB, 0xAB, 0xF5,    0x6A, 0x6B, 0x35, 0xD5, 
	0xC1, 0xAA, 0x40, 0xB2,    0x76, 0x5C, 0x36, 0xE0,    0xDC, 0x6B, 0x47, 0x6E,    0xFC, 0x83, 0xF3, 0x7F, 
	0x71, 0x49, 0x7E, 0x51,    0xAA, 0x9E, 0x81, 0x5E,    0x52, 0x4E, 0x7E, 0x72,    0x36, 0xFD, 0xF3, 0xBF, 
	0xA1, 0x91, 0x11, 0x5A,    0xFE, 0x37, 0x32, 0x36,    0x1A, 0xCD, 0xFF, 0x74,    0x05, 0x0A, 0x54, 0x57, 
	0x38, 0x0A, 0x86, 0x02,    0x40, 0xCA, 0xFF, 0x86,    0x03, 0x95, 0xFF, 0xCD,    0xCC, 0x0C, 0x31, 0xF2, 
	0xBF, 0xB1, 0xD1, 0x68,    0xFE, 0xA7, 0x23, 0x60,    0x54, 0xD8, 0x76, 0x6E,    0x0E, 0xDF, 0xAD, 0x49, 
	0x53, 0xEB, 0xFF, 0xDC,    0xB1, 0x5C, 0xEC, 0xF7,    0xEA, 0xF6, 0xC5, 0x93,    0xEB, 0xAE, 0xD8, 0xCC, 
	0xD9, 0xB1, 0x4A, 0x43,    0xC0, 0xF7, 0x4C, 0x71,    0xE6, 0xEA, 0x77, 0xD5,    0x30, 0x85, 0x50, 0x0C, 
	0x06, 0xC1, 0x8B, 0x7D,    0x42, 0xA0, 0x4C, 0x16,    0x01, 0x41, 0x21, 0x61,    0xB8, 0x1A, 0x45, 0xA6, 
	0xFF, 0x40, 0x80, 0xAC,    0x14, 0x04, 0xFE, 0x23,    0x01, 0xB8, 0x0A, 0xB8,    0xE8, 0x9E, 0x67, 0xBF, 
	0xD6, 0x2E, 0x17, 0x9F,    0xD7, 0xF2, 0x79, 0xE7,    0xA9, 0x43, 0x7F, 0x92,    0x55, 0x03, 0x15, 0x99, 
	0xA2, 0x8E, 0x99, 0xDE,    0x58, 0x65, 0x35, 0xF9,    0xF9, 0xE6, 0xD7, 0x7B,    0x43, 0xCB, 0x66, 0xB4, 
	0xED, 0x49, 0x95, 0x65,    0xDB, 0x70, 0x26, 0x78,    0xC3, 0x37, 0xEB, 0x73,    0x36, 0x76, 0xEA, 0x97, 
	0x12, 0xBC, 0x15, 0x99,    0xB2, 0xC5, 0x2F, 0x7E,    0x7A, 0xA8, 0xE3, 0xE4,    0xFE, 0x63, 0xCF, 0xB3, 
	0xA7, 0xC9, 0x4B, 0x1C,    0x3E, 0x95, 0x33, 0xD7,    0x36, 0xEA, 0xBE, 0x36,    0x5E, 0xF0, 0x65, 0xA1, 

	0xA5, 0xEB, 0x8D, 0x19,    0x87, 0xA6, 0x29, 0x32,    0xF9, 0x3F, 0x76, 0x7A,    0xF4, 0x4F, 0xAA, 0x7E, 
	0x76, 0xDF, 0xF3, 0xD7,    0x5E, 0x35, 0xFC, 0xF3,    0xC4, 0xB4, 0xCF, 0x19,    0x87, 0x67, 0x1B, 0xC6, 
	0x9D, 0x3B, 0xBD, 0xCD,    0x21, 0xC3, 0x7C, 0x7F,    0xE0, 0x57, 0x84, 0x8B,    0xDF, 0x58, 0x78, 0xD6, 
	0xFC, 0x9A, 0x1D, 0xEC,    0xF4, 0x3A, 0x57, 0xF2,    0xC0, 0xE5, 0xCC, 0x3F,    0x6B, 0x7C, 0x2F, 0xE6, 
	0xFE, 0x63, 0xF4, 0xCB,    0xD9, 0xFF, 0x55, 0x2E,    0xDF, 0xE0, 0x74, 0xE5,    0xD2, 0x3A, 0x45, 0xA6, 
	0xAD, 0xEF, 0x3F, 0xEE,    0x89, 0xF2, 0xAD, 0xEF,    0x55, 0xCA, 0x6F, 0x31,    0x5D, 0x3A, 0x89, 0xA5, 
	0x47, 0x3C, 0xF2, 0xE4,    0xD1, 0x25, 0x96, 0xF1,    0x19, 0xC9, 0x4D, 0xFC,    0xBA, 0xEF, 0x02, 0xCB, 
	0xEC, 0x15, 0x99, 0x44,    0xA2, 0xBF, 0xD9, 0xBA,    0xBE, 0x11, 0xEE, 0xD7,    0xFF, 0x70, 0x3B, 0x78, 
	0xB9, 0xE5, 0xEE, 0xDB,    0xEC, 0x11, 0xFE, 0x4B,    0x16, 0x7A, 0x6C, 0xC9,    0x4F, 0xCF, 0xF7, 0x0A, 
	0x52, 0x36, 0x6A, 0x9B,    0xA3, 0xC8, 0x54, 0x14,    0x75, 0xB6, 0x56, 0x45,    0x32, 0xCB, 0xE4, 0xFE, 
	0x7C, 0xCF, 0xA3, 0xCD,    0x2D, 0xB9, 0xA6, 0x9F,    0xE7, 0x25, 0x38, 0x14,    0x36, 0xEE, 0x34, 0x58, 
	0xF1, 0x64, 0x4E, 0xC1,    0xEA, 0xCB, 0x7B, 0x77,    0x2B, 0x78, 0xAC, 0xBD,    0x3C, 0xAB, 0x42, 0x81, 
	0xA7, 0x32, 0xE4, 0x53,    0xF5, 0x33, 0xD6, 0x6C,    0x5E, 0x15, 0x51, 0xFE,    0xD9, 0x1B, 0x37, 0x54, 
	0xFF, 0x98, 0xBF, 0xF8,    0xFB, 0x0D, 0x2B, 0xA5,    0xBF, 0x77, 0xD7, 0xD3,    0x2E, 0xFF, 0x67, 0x15, 
	0xE7, 0xE7, 0x0D, 0x48,    0xFB, 0xDF, 0xCC, 0xC4,    0x04, 0x3D, 0xFF, 0x1B,    0x98, 0x1A, 0x8F, 0xE6, 
	0x7F, 0x7A, 0x80, 0x68,    0x60, 0xA3, 0x55, 0x09,    0x5C, 0xEE, 0x7B, 0xBA,    0x28, 0x59, 0x29, 0x18, 

	0xE8, 0x00, 0xB9, 0x05,    0x45, 0xA9, 0x65, 0x99,    0xF9, 0xA5, 0xC5, 0x1E,    0x89, 0xC5, 0x19, 0x40, 
	0xB1, 0xBC, 0xD2, 0x9C,    0x1C, 0x90, 0x30, 0xB0,    0x31, 0x5B, 0x94, 0x99,    0x5A, 0x0C, 0x14, 0x89, 
	0x8E, 0x05, 0xF1, 0x8B,    0x13, 0x73, 0x4A, 0x80,    0x1C, 0x50, 0xAB, 0x57,    0x29, 0xA9, 0xB2, 0x04, 
	0x2C, 0xA3, 0xE4, 0x48,    0x3C, 0xB0, 0x55, 0xE2,    0x02, 0x36, 0x7D, 0x6B,    0x75, 0xB8, 0x50, 0x9D, 
	0x60, 0x88, 0xC5, 0x09,    0xA8, 0x76, 0xF8, 0xB8,    0x84, 0xF8, 0x7B, 0xA4,    0x79, 0x9A, 0xA7, 0x9B, 
	0x54, 0x24, 0x96, 0x55,    0x16, 0xE8, 0x47, 0xB9,    0xA4, 0xF9, 0xE6, 0x5B,    0xFA, 0xBA, 0xA7, 0x95, 
	0x84, 0x7A, 0xE9, 0x57,    0x46, 0x54, 0x78, 0xA4,    0x25, 0x56, 0xB8, 0x1B,    0xA6, 0x17, 0x85, 0xA6, 
	0x16, 0x83, 0xED, 0x40,    0x73, 0x3C, 0xD0, 0x2C,    0x70, 0x4B, 0x1D, 0x2C,    0x56, 0x19, 0x02, 0xEC, 
	0x9C, 0x40, 0x3D, 0x0E,    0x14, 0x2A, 0x2E, 0x29,    0x2A, 0x4D, 0x2E, 0x29,    0x2D, 0x4A, 0x4D, 0x71, 
	0x01, 0x66, 0x0C, 0x90,    0x65, 0xAE, 0x4E, 0xAE,    0xC1, 0xAE, 0xE5, 0xB6,    0xB6, 0x4A, 0x50, 0x05, 
	0x99, 0xE9, 0x79, 0x89,    0x20, 0x05, 0xF0, 0x70,    0x00, 0x0A, 0x96, 0x64,    0xE6, 0xA6, 0x06, 0x27, 
	0xE6, 0x16, 0x80, 0xCC,    0xE1, 0x02, 0xB7, 0xE7,    0x71, 0x07, 0x90, 0x87,    0x97, 0xB9, 0x7F, 0x65, 
	0xA0, 0x49, 0xBE, 0x41,    0x76, 0x64, 0x78, 0xA8,    0xAB, 0xBB, 0xBB, 0x5F,    0x64, 0xAA, 0x77, 0x84, 
	0x97, 0x47, 0x76, 0x9E,    0x41, 0x4A, 0x64, 0x52,    0xB1, 0xA5, 0xA1, 0x45,    0x91, 0x8B, 0x4F, 0x69, 
	0xA8, 0x99, 0x57, 0x72,    0x90, 0x27, 0x34, 0x80,    0x62, 0x47, 0x5B, 0x29,    0xA3, 0x60, 0x14, 0x8C, 
	0x02, 0x5A, 0x00, 0x40,    0x00, 0x00, 0x00, 0xFF,    0xFF, 0x05, 0x19, 0x2A,    0x47, 0x00, 0x14, 0x00, 

	0x00, 
}

func init() {
	var err error
	gobundle.Setup.Application.ResourceTar, err = gzip.NewReader(bytes.NewReader(_data))
	if err != nil { panic(err) }
}
