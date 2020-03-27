package skyphotos

import (
	"log"
	"os"
)

var (
	sourceFolder   string
	filteredFolder string
	minHeight      float64
)

const (
	jpgRegexp = `.*\.(jpg|JPG)$`
)
//check existance of folder structure and create if it does not exist
func SetupFileStructure (sourceFolder string, outputFolder string)(sF,oF string){
  //create outputFolder
if dirExists(sourceFolder){
  //check output outputFolder
  if !dirExists(outputFolder){
    //create output outputFolder
    if os.MkdirAll(outputFolder,0777)!=nil {
      log.Fatal ("Cant create Output Ditectory")
    }
  }
} else{
  log.Fatal("Source Directory does not exist")
}
  sF=sourceFolder
  oF=outputFolder
  return sF,oF
}
//=================================
//check weather Directory is exists
 func dirExists(dir string) bool {
   info, err:=os.Stat(dir)
   if os.IsNotExist(err){
     return false
   }
   return info.IsDir()
 }
