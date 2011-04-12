package common

import "os"


func PanicIf(error os.Error) {
    if error != nil {
        panic(error)
    }
}


func GetFileSize(fobj *os.File) int64 {
    fileinfo, error := fobj.Stat()
    PanicIf(error)
    return fileinfo.Size
}


func SafeSeek(fobj *os.File, offset int64, whence int) int64 {
    pos, error := fobj.Seek(offset, whence)
    PanicIf(error)
    return pos
}


func GetFilePos(fobj *os.File) int64 {
    return SafeSeek(fobj, 0, 1)
}
