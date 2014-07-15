A reader for Microsoft's OLE Property Set Format.

Example usage:

    file, _ := os.Open("test/test.doc")
    defer file.Close()
    doc, err := mscfb.NewReader(file)
    if err != nil {
      log.Fatal(err)
    }
    for entry, err := doc.Next(); err == nil; entry, err = doc.Next() {
      if msoleps.IsMSOLEPS(doc.Initial()) {
        for prop, oerr = entry.Read(); oerr == nil; prop, oerr = entry.Read() {
          fmt.Println(prop.Name)
        }
      }
    }

Install with `go get github.com/richardlehane/msoleps`

*I'm being developed and am not yet ready...*

[![Build Status](https://travis-ci.org/richardlehane/msoleps.png?branch=master)](https://travis-ci.org/richardlehane/msoleps)