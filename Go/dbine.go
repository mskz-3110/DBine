package dbine

import(
  "os/exec"
  "bytes"
)

func DotToPdfBuffer(dot string)([]byte){
  cmd := exec.Command("dot", "-Tpdf")
  cmd.Stdin = bytes.NewReader([]byte(dot))
  pdfBuffer, _ := cmd.Output()
  return pdfBuffer
}
