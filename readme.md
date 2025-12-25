# CSV To Markdown Table Converter

This is a utility tool used to convert [Comma-separated values (CSV)](https://en.wikipedia.org/wiki/Comma-separated_values) files to a [Markdown table](https://www.markdownguide.org/extended-syntax/#tables).

## Table Of Contents

- [CSV To Markdown Table Converter](#csv-to-markdown-table-converter)
  - [Table Of Contents](#table-of-contents)
  - [Usage](#usage)
  - [Configuration Options](#configuration-options)
  - [Performance](#performance)

## Usage

An example conversion can look like this:

```go
package main

import (
  "fmt"

  csv2md "github.com/phamduylong/csv-to-md"
)

func main() {
  var cfg csv2md.Config
  cfg.Align = csv2md.Center
  cfg.VerboseLogging = true

  res, convertErr := csv2md.Convert(cfg)

  if convertErr != nil {
    fmt.Println(convertErr)
  }

  fmt.Println(res)
}
```

## Configuration Options

The program offers a range of different configuration options to customize the tool to best fit your use case.

| Option                           | Type            | What does it do? |
|----------------------------------|-----------------|------------------|
| Align                            | Align           | Align the text on the rendered table. Visual feedback on the markdown syntax is also provided. |
| CSVReaderConfig                  | CSVReaderConfig | Config options to be passed into CSV reader object. See type Reader in the encoding/csv module. |
| CSVReaderConfig.Comma            | Rune            | Set the delimiter of the CSV reader |
| CSVReaderConfig.Comment          | Rune            | Set the comment character for the CSV reader. |
| CSVReaderConfig.FieldsPerRecord  | int             | Set the amount of fields per CSV row. |
| CSVReaderConfig.LazyQuotes       | bool            | Set whether lazy quotes are allowed. If lazy quotes are allowed, a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field. |
| CSVReaderConfig.TrimLeadingSpace | bool            | Set whether leading space before the fields' values should be ignored. |
| CSVReaderConfig.ReuseRecord      | bool            | Set whether calls to Read may return a slice sharing the backing array of the previous call's returned slice for performance. By default, each call to Read returns newly allocated memory owned by the caller. |
| VerboseLogging                   | bool            | Log detailed diagnostic messages when running the program. |

## Performance

*I did not create a very proper setup to measure the performance. Ran it with my own PC so take it with a grain of salt.*

| Rows    | Columns | Average Execution Time (5 runs) |
| ------- | ------- | ------------------------------- |
| 1.000   | 10      | 29ms                            |
| 10.000  | 10      | 2,1s                            |
| 50.000  | 10      | 48,2s                           |
| 100.000 | 10      | 3 minutes                       |
