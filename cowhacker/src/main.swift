import Foundation
import Darwin

enum Mode {
    case Patch
    case Revert
}

func parseBasePath (file: NSFileHandle) -> String {
    file.seekToFileOffset(0)
    let data = file.readDataOfLength(1024)
    let pathOffset = 0xA0 + 24
    let pathMaxLen = 512
    let readBuffer = UnsafeMutablePointer<UInt8>.alloc(pathMaxLen)
    data.getBytes(readBuffer, range: NSMakeRange(pathOffset, pathMaxLen))

    let buf = UnsafeMutablePointer<CChar>(readBuffer)
    buf[pathMaxLen - 1] = 0
    var str = ""
    if let utf8String = String.fromCString(buf) {
        str = utf8String
    }
    readBuffer.dealloc(pathMaxLen)
    return str
}

func updateBasePath (file: NSFileHandle, newPath: String) -> Int {
    if count(newPath) > 256 {
        println("Path too long")
        return 1
    }
    // Write length byte
    let lengthByteOffset: UInt64 = 0x13
    file.seekToFileOffset(lengthByteOffset)
    let lengthByte = NSData(bytes: [ UInt8(count(newPath)) ], length: 1)
    file.writeData(lengthByte)

    // Write path
    let pathOffset: UInt64 = 0xA0 + 24
    file.seekToFileOffset(pathOffset)
    let strData = NSMutableData(data: newPath.dataUsingEncoding(NSUTF8StringEncoding)!)
    strData.appendBytes([ UInt8(0) ], length: 1)
    file.writeData(strData)

    return 0
}

// Entry point

if Process.arguments.count < 3 {
    println("Usage: \(Process.arguments[0]) [patch|revert] qcow_path [orig_path]")
    println("   1) '\(Process.arguments[0]) patch qcow_path' outputs original path to backing image")
    println("   2) '\(Process.arguments[0]) revert qcow_path orig_path' sets path to backing image back to orig_path")
    exit(1)
}

let mode: Mode

if let modeStr = String.fromCString(Process.arguments[1]) {
    if modeStr == "patch" {
        mode = Mode.Patch
    } else if modeStr == "revert" {
        mode = Mode.Revert
    } else {
        println("Invalid mode")
        exit(1)
    }
} else {
    println("Invalid mode")
    exit(1)
}

func main () -> Int32 {
    if let path = String.fromCString(Process.arguments[2]) {
        if let file = NSFileHandle(forUpdatingAtPath: path) {
            // defer {
            //     file.closeFile()
            // }
            let error: Int
            if mode == Mode.Patch {
                println(parseBasePath(file))
                error = updateBasePath(file, "C:\\ece391\\devel\\ece391.qcow")
            } else {
                if let newPath = Process.arguments.count >= 4 ? String.fromCString(Process.arguments[3]) : nil {
                    error = updateBasePath(file, newPath)
                } else {
                    println("Please specify new backing image path")
                    file.closeFile()
                    return 1
                }
            }
            if error != 0 {
                println("Error occurred")
                file.closeFile()
                return 1
            }
        } else {
            println("Cannot open file")
            return 1
        }
    } else {
        println("Invalid path")
        return 1
    }
    return 0
}

exit(main())
