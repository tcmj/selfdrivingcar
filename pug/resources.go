package pug

import (
    "bytes"
    "github.com/hajimehoshi/ebiten/v2"
    "image"
)

var DATA_GOPHER = []byte("" +
    "\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x19\x00\x00\x00 \b\x06\x00\x00\x00\xe7\x9c\xd3\x06\x00\x00\x00" +
    "\x01sRGB\x00\xae\xce\x1c\xe9\x00\x00\x01BIDATH\x89\xc5V\xc1\x8d\xc30\f\xa3\x8a{\x1d\xd0Un\xa4ۣ\x13d\x8f\x8e" +
    "\xd4U\x02\xf4\xab~\"CVdYJ\x8b\x94@`'2͈\xb6\x9c\x10z\xb0\xea\x13\xeap\xf9z\"^xm77\xbaV\xc4\x18\x00\x1c~\xc7\xe5" +
    "\x85W^x\xe5\x8d\xd0\xf53\x97\xe5\xaa\xfb\xa6\xc4\v\xaf\xfa\xed\x1b\xf4ۍ0\xe2I6?#bfr;\xd6\x13\xb3h)Z\xeb`l" +
    "\x88b\xe8\xadۋ\xc0\xf8k\x91\x8d\xa9\v\x00pQ\"\xe4\xd9D4\xde\\^L\xf1[\xf0\xb2\x1be\xd3c7\xebiLc*\xf2\t\x84\xbb" +
    "\xcb\xda!VD\xb1\x92ȌX\xd9⡈\x85\xae\x83\x8aHzMl\xa1e\n\xaf$r\xa3+\x9e\x8f{\xf7\xec\xf9\xb8\xa7\x85JvY\xa1,\xd2" +
    "\"\xbf\u007f\xff\x87\x04\x80ޮae\xc9\"K\xf1I;Y\xfcݱ\xc23+\xa46t\x1ba\x9b\x8fEd*\xa0\x85t;\x83\b5\xbb\xde\xf1܃\x9e" +
    "\xaf}\x19\xb7>W\x8al\x04\xf5}g\x00$\x99\x1c\xf93ɀ\x80\x93N\xe1SD<\x9b\xdeZ\x17\xef\u007f\xebk\x99\x00\a\xb3\xf1" +
    "\xb2\bE\xca\n\xc1\x9c\xd1\xd6m'\xc1VX\xee\xfa%ƌ3ɐ\x0f\x8cݓ\x91\xb7-\x1c\xfb\x02f5\x10\xc8q}\xfd(\x00\x00\x00" +
    "\x00IEND\xaeB`\x82",
)

var DATA_PINGUIN = []byte("" +
    "\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x02\x00\x00\x00\x02\x08\x02\x00\x00\x00\xfd\xd4\x9as\x00\x00\x00\x19IDAT\x08\x1d\x01\x0e\x00\xf1\xff\x00\x00\x00\x00\x00\xff\xff\x01\x00\xff\xff\x13\x90\x90\x1b\xe4 \x0510O\xffC \x00\x00\x00\x00IEND\xaeB`\x82",
)

func GetIconGopher() []image.Image {
    return GetIcon(DATA_GOPHER)
}

func GetIcon(b []byte) []image.Image {
    m, _, err := image.Decode(bytes.NewReader(b))
    if err != nil {
        return nil
    }
    tmp := []image.Image{ebiten.NewImageFromImage(m)}
    return tmp
}
