package util

import (
    "archive/tar"
    "compress/gzip"
    "fmt"
    "io"
    "os"
    "strings"
    "path/filepath"
)

func Untar(sourcefile string, prefix string) error {

    file, err := os.Open(sourcefile)

    if err != nil { return err }

    defer file.Close()

    var fileReader io.ReadCloser = file

    // just in case we are reading a tar.gz file, add a filter to handle gzipped file
    if strings.HasSuffix(sourcefile, ".gz") {
        if fileReader, err = gzip.NewReader(file); err != nil { return err }
        defer fileReader.Close()
    }

    tarBallReader := tar.NewReader(fileReader)

    // Extracting tarred files

    for {
        header, err := tarBallReader.Next()
        if err != nil {
            if err == io.EOF {
                break
            }
            return err
        }

        // get the individual filename and extract to the current directory
        filename := header.Name
        filename = filepath.Join(prefix, filename)
        fmt.Printf("\r + Creating %s", filename)

        switch header.Typeflag {
        case tar.TypeDir:
            // handle directory
            if _, err := os.Stat(filename); os.IsNotExist(err) {
                err = os.MkdirAll(filename, os.FileMode(header.Mode | 0600 )) // or use 0755 if you prefer
            } else {
                // If the directory exists, update its mode
                err = os.Chmod(filename, os.FileMode(header.Mode))
            }

            if err != nil { return err }

        case tar.TypeReg:
            // handle normal file
            // Sometimes the directory is specified after the file
            if err := os.MkdirAll(filepath.Dir(filename), 0777); err != nil { return err }
            writer, err := os.Create(filename)
            if err != nil { return err }

            io.Copy(writer, tarBallReader)
            // Make it at least read/writable by the current user
            err = os.Chmod(filename, os.FileMode(header.Mode | 0600))
            if err != nil { return err }

            writer.Close()
        default:
            return fmt.Errorf("Unable to untar type : %c in file %s", header.Typeflag, filename)
        }
    }

    return nil
}
